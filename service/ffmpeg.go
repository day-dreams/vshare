package service

import (
	"context"
	"os/exec"
	"strconv"

	"github.com/tidwall/gjson"

	"github.com/day-dreams/vshare.zhangnan.xyz/utils"
)

type VideoInfo struct {
	Duration float64
	Size     uint64
}

func VideoInfoGet(ctx context.Context, path string) (*VideoInfo, error) {

	cmd := exec.CommandContext(ctx, "ffprobe",
		"-v", "quiet",
		"-print_format", "json",
		"-show_format", path)

	if bytes, err := cmd.Output(); err != nil {
		utils.Logger().Errorf("cmd.Output() failed. %v", err)
		return nil, err
	} else {

		utils.Logger().Debugf("data: %s ", string(bytes))
		ret := &VideoInfo{}

		duration := gjson.Get(string(bytes), "format.duration").String()
		ret.Duration, err = strconv.ParseFloat(duration, 64)
		if err != nil {
			utils.Logger().Errorf("strconv.ParseFloat failed. %v", err)
			return nil, err
		}

		size := gjson.Get(string(bytes), "format.size").String()
		ret.Size, err = strconv.ParseUint(size, 10, 64)
		if err != nil {
			utils.Logger().Errorf("strconv.ParseFloat failed. %v", err)
			return nil, err
		}
		return ret, nil
	}

}
