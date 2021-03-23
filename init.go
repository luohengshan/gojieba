package gojieba

import (
	"github.com/luohengshan/gojieba/deps/cppjieba"
	"github.com/luohengshan/gojieba/deps/limonp"
	"github.com/luohengshan/gojieba/dict"
)

func init() {
	dict.Init()
	limonp.Init()
	cppjieba.Init()
}
