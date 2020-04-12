package handler

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/day-dreams/vshare.zhangnan.xyz/utils"
)

// 播放相关

type tag struct {
	payload string // #EXT-X-VERSION:3 + '\n'
}

func (t *tag) IsEXTINF() bool {
	return strings.Index(t.payload, "EXTINF:") == 0
}
func (t *tag) GetEXTINFUri() string {
	return strings.Split(t.payload, "\n")[1]
}
func (t *tag) SetEXTINFUri(uri string) {
	// fmt.Printf("split: %s -> %s\n", t.payload, strings.Split(t.payload, "\n"))
	t.payload = strings.Split(t.payload, "\n")[0] + "\n" + uri + "\n"
}

type m3u8 struct {
	tags []tag
}

func (m *m3u8) ToHttpBody(vid string) string {
	body := ""
	for _, tag := range m.tags {
		if tag.IsEXTINF() {
			// format as http uri
			uri := tag.GetEXTINFUri()
			uri = fmt.Sprintf("http://vshare.zhangnan.xyz/api/play/m3u8/ts?vid=%s&tsname=%s", vid, uri)
			tag.SetEXTINFUri(uri)
		}
		body += "#" + tag.payload
	}
	return body
}

func newM3u8ByFile(path string) (*m3u8, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	rv := &m3u8{}
	tags := strings.Split(string(content), "#")
	for _, t := range tags {
		rv.tags = append(rv.tags, tag{t})
	}

	return rv, nil
}

func M3u8PlayList() gin.HandlerFunc {
	return func(c *gin.Context) {

		param := ParamM3u8TsFile{
			Vid: c.Query("vid"),
		}
		if param.Vid == "" {
			utils.GinJson(c, nil, fmt.Errorf("need vid"))
			return
		}

		video := vid2video[param.Vid]

		// m, err := newM3u8ByFile("/Users/kakaxi/Desktop/1988/index.m3u8")
		m, err := newM3u8ByFile(video.Hls.M3u8)
		if err != nil {
			utils.GinJson(c, nil, err)
			return
		}

		c.Writer.Header().Set("Content-Type", "application/vnd.apple.mpegurl")
		body := m.ToHttpBody(param.Vid)
		c.Writer.Header().Set("Content-Length", strconv.Itoa(len(body)))
		c.Writer.WriteHeader(http.StatusOK)
		_, _ = c.Writer.Write([]byte(body))
	}
}

type ParamM3u8TsFile struct {
	Vid    string `json:"vid"`
	TsName string `json:"tsname"`
}

func M3u8TsFile() gin.HandlerFunc {
	return func(c *gin.Context) {
		param := ParamM3u8TsFile{
			Vid:    c.Query("vid"),
			TsName: c.Query("tsname"),
		}
		if param.Vid == "" || param.TsName == "" {
			utils.GinJson(c, nil, fmt.Errorf("need vid/tsname"))
			return
		}

		video, ok := vid2video[param.Vid]
		if !ok {
			utils.GinJson(c, nil, fmt.Errorf("video not exists"))
			return
		}

		path := filepath.Join(video.Hls.TsDir, param.TsName)
		bytes, err := ioutil.ReadFile(path)
		if err != nil {
			utils.GinJson(c, nil, err)
			return
		}

		c.Writer.Header().Set("Content-Type", "application/octet-stream")
		c.Writer.WriteHeader(http.StatusOK)
		_, _ = c.Writer.Write(bytes)
	}
}
