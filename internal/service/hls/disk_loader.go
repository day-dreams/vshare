package hls

import (
	bytes2 "bytes"
	"context"
	"fmt"
	"github.com/day-dreams/vshare.zhangnan.xyz/internal/utils"
	"io"
	"io/ioutil"
	"path/filepath"
	"strings"
)

type DiskFfmpegLoader struct {
}

func (d DiskFfmpegLoader) M3u8PlayList(ctx context.Context, req *ReqM3u8PlayList) (*ResM3u8PlayList, error) {
	dir := GetVidM3u8Path(req.Vid)

	index := filepath.Join(dir, "index.m3u8")
	bytes, err := ioutil.ReadFile(index)
	if err != nil {
		utils.Logger().Errorf("fail to read file. %s,%v", index, err)
		return nil, err
	}
	playlist := bytes2.NewBuffer(bytes)

	res := &ResM3u8PlayList{PlayListContent: ""}

	for {

		line, err := playlist.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			utils.Logger().Errorf("read string failed. %v", err)
			return nil, err
		}

		pos := strings.Index(line, "#")
		if pos == 0 {
			// # xxxxxx
			res.PlayListContent += line

		} else if pos < 0 {
			// indexN.ts
			num := strings.TrimSuffix(strings.TrimPrefix(line, "index"), ".ts\n")
			res.PlayListContent += fmt.Sprintf("%s?vid=%s&segment=%s\n", req.Path, req.Vid, num)

		} else {
			err := fmt.Errorf("invalid line in m3u8.index. [%s]", line)
			utils.Logger().Error(err)
			return nil, err
		}

	}

	return res, nil
}

func (d DiskFfmpegLoader) M3u8Segment(ctx context.Context, req *ReqM3u8Segment) (*ResM3u8Segment, error) {
	dir := GetVidM3u8Path(req.Vid)
	aim := filepath.Join(dir, fmt.Sprintf("index%d.ts", req.Index))

	bytes, err := ioutil.ReadFile(aim)
	if err != nil {
		utils.Logger().Errorf("fail to read file. %s,%v", aim, err)
		return nil, err
	}

	return &ResM3u8Segment{Content: bytes}, nil
}
