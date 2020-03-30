package handler

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
	"github.com/tidwall/gjson"

	"github.com/day-dreams/vshare.zhangnan.xyz/utils"
)

type action = string

const (
	chat  = action("chat")  // 聊天类消息
	play  = action("play")  // 控制类消息-播放
	pause = action("pause") // 控制类消息-暂停
	jump  = action("jump")  // 控制类消息-快进或者跳转
)

type message struct {
	Action action

	Ctime int64  // 消息生成时间
	From  string // cid

	// chat
	Text string // 聊天文字
	// play / pause / jump
	Offset int // 目标进度的时间偏移 秒
}

type Client struct {
	Cid    string     // 客户端id
	Addr   string     // 客户端ip
	JoinAt time.Time  // 加入时间
	Box    []*message // 未读消息盒
}

type status struct {
	// 客户端
	Clients []*Client

	// 消息列表
	Msgs []*message

	// 当前进度
	Current int    // todo 需要一个主从机制 需要有人来同步这个状态，让新进来的人能够自动播放
	State   action // play or pause
}

type Room struct {
	RID    string
	VID    string
	Video  video
	Status *status
}

var (
	mutex      sync.RWMutex // 一张大锁cover所有问题
	rooms      []*Room
	rid2room   = map[string]*Room{}
	cid2client = map[string]*Client{}
)

func setupRoom() {
	bytes, err := ioutil.ReadFile(utils.VFile)
	if err != nil {
		panic(err)
	}
	data := string(bytes)

	gjson.Get(data, "rooms").ForEach(func(key, value gjson.Result) bool {
		vid := value.Get("vid").String()
		rid := value.Get("rid").String()

		r := Room{
			RID:   rid,
			VID:   vid,
			Video: vid2video[vid],
			Status: &status{
				Clients: nil,
				Msgs:    nil,
				Current: 0,
			},
		}
		rooms = append(rooms, &r)
		rid2room[rid] = &r
		return true
	})
}

func CidAddToCookie(c *gin.Context, cid string) {
	c.SetCookie("vsharecid", cid, int(time.Hour*24*7),
		"/", c.Request.Host, false, false)
}
func CidFromCookie(c *gin.Context) string {
	rid, err := c.Cookie("vsharecid")
	if err != nil {
		return ""
	}
	return rid
}

func genCid() string {
	return xid.New().String()
}

func RegisterAsClient(c *gin.Context, rid string) (cid string, err error) {
	mutex.Lock()
	defer mutex.Unlock()

	cid = CidFromCookie(c)

	room, ok := rid2room[rid]
	if !ok {
		return "", utils.ErrParamInvalid
	}

	client := cid2client[cid]
	if cid == "" || client == nil {
		cid = genCid()
		client := &Client{
			Cid:    cid,
			Addr:   c.Request.RemoteAddr,
			JoinAt: time.Now(),
			Box:    nil,
		}
		cid2client[cid] = client
		room.Status.Clients = append(room.Status.Clients, client)
		CidAddToCookie(c, cid)
	} else {
		c := cid2client[cid]
		if c == nil {
			return "", utils.ErrParamInvalid
		}
	}

	return cid, nil
}

func RoomEnter() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Printf("rid2room: [%s]\n", utils.ToPrettyJSON(rid2room))

		rid := c.Query("rid")
		room, ok := rid2room[rid]
		if !ok {
			utils.GinJson(c, nil, utils.ErrParamInvalid)
			return
		}

		cid, err := RegisterAsClient(c, rid)
		if err != nil {
			utils.GinJson(c, nil, err)
			return
		}

		data := map[string]interface{}{
			"online": len(room.Status.Clients),
			"cid":    cid,
			"at":     room.Status.Current,
		}
		utils.GinJson(c, data, nil)
	}
}

type paramReadStatus struct {
	Rid string `json:"rid" validate:"required"`
	Cid string `json:"cid"`
}

func StatusRead() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取消息列表
		param := paramReadStatus{}
		param.Rid = c.Query("rid")
		param.Cid = CidFromCookie(c)

		fmt.Printf("param:%s\n", utils.ToPrettyJSON(param))
		if err := utils.Validate(&param); err != nil {
			utils.GinJson(c, nil, err)
			return
		}
		param.Cid = CidFromCookie(c)

		client, ok := cid2client[param.Cid]
		if !ok {
			utils.GinJson(c, nil, utils.ErrParam("cid不存在"))
			return
		}

		mutex.Lock()
		defer mutex.Unlock()

		var (
			msgs    = client.Box
			ret     = map[string]interface{}{}
			chats   []*message
			control *message
		)
		client.Box = nil
		for _, msg := range msgs {
			if msg.Action == chat {
				chats = append(chats, msg)
			} else {
				// 不要下发多个播放控制命令，频繁开销
				control = msg
			}
		}
		ret["chats"] = chats
		ret["control"] = control

		utils.GinJson(c, ret, nil)
	}
}

type paramWriteStatus struct {
	Rid    string `validate:"required"`
	Action action `validate:"required"`
	Cid    string `validate:"required"`
	At     int
	Text   string
}

func atoi(src string) int {
	rv, _ := strconv.ParseFloat(src, 64)
	return int(rv)
}

func StatusWrite() gin.HandlerFunc {
	return func(c *gin.Context) {

		// 写扩散 - 每个消息都要发到各个client的msg box，client处理起来方便一些

		param := paramWriteStatus{}
		param.Rid = c.Query("rid")
		param.Action = c.Query("action")
		param.Cid = CidFromCookie(c)
		param.At = atoi(c.Query("at"))
		param.Text = c.Query("text")
		if err := utils.Validate(&param); err != nil {
			utils.GinJson(c, nil, err)
			return
		}

		room, ok := rid2room[param.Rid]
		if !ok {
			utils.GinJson(c, nil, utils.ErrParam("rid不存在"))
			return
		}
		sclient, ok := cid2client[param.Cid]
		if !ok {
			utils.GinJson(c, nil, utils.ErrParam("cid不存在"))
			return
		}

		mutex.Lock()
		defer mutex.Unlock()
		ctime := time.Now().Unix()
		switch param.Action {
		case chat:
			msg := &message{
				Action: chat,
				Ctime:  ctime,
				From:   sclient.Cid,
				Text:   param.Text,
				Offset: 0,
			}
			for _, client := range room.Status.Clients {
				if client.Cid != param.Cid {
					client.Box = append(client.Box, msg)
				}
			}
			room.Status.Msgs = append(room.Status.Msgs, msg)
		case pause:
			room.Status.State = pause
			room.Status.Current = param.At
			msg := &message{
				Action: pause,
				Ctime:  ctime,
				From:   sclient.Cid,
				Offset: param.At,
			}
			for _, client := range room.Status.Clients {
				if client.Cid != param.Cid {
					client.Box = append(client.Box, msg)
				}
			}
		case play:
			room.Status.State = play
			room.Status.Current = param.At
			msg := &message{
				Action: play,
				Ctime:  ctime,
				From:   sclient.Cid,
				Offset: param.At,
			}
			for _, client := range room.Status.Clients {
				if client.Cid != param.Cid {
					client.Box = append(client.Box, msg)
				}
			}
		case jump:
			room.Status.State = play
			room.Status.Current = param.At
			msg := &message{
				Action: jump,
				Ctime:  ctime,
				From:   sclient.Cid,
				Offset: param.At,
			}
			for _, client := range room.Status.Clients {
				if client.Cid != param.Cid {
					client.Box = append(client.Box, msg)
				}
			}
		default:
			utils.GinJson(c, nil, utils.ErrParam("action wrong"))
			return
		}

		utils.GinJson(c, nil, nil)
	}
}
