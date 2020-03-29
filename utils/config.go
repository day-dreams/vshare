package utils

import (
	"flag"
	"fmt"
	"os"
)

var (
	VFile string
)

func init() {
	flag.StringVar(&VFile, "vfile", "", "video file path")
	flag.Parse()
}

func Config(key string) string {
	val := os.Getenv(key)
	if val == "" {
		panic(fmt.Errorf("%s  not exists", key))
	}
	return val
}
