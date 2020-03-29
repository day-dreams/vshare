package utils

import (
	"io/ioutil"
)

func Index() []byte {
	bytes, err := ioutil.ReadFile(VIndex)
	if err != nil {
		panic(err)
	}
	return bytes
}
