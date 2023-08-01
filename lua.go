package arp

import (
	"github.com/vela-ssoc/vela-kit/lua"
	"github.com/vela-ssoc/vela-kit/vela"
)

var xEnv vela.Environment

func arpL(L *lua.LState) int {
	tab := Table()
	ret := lua.NewMap(len(tab), false) //生成一个 lua 中的map对象

	for key, val := range tab {
		ret.Set(key, lua.S2L(val))
	}

	L.Push(ret)
	return 1
}

func WithEnv(env vela.Environment) {
	xEnv = env
	define(xEnv.R())

	xEnv.Set("arp", lua.NewExport("vela.arp.export", lua.WithFunc(arpL)))

}
