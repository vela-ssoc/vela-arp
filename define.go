package arp

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
	"github.com/vela-ssoc/vela-kit/vela"
)

func define(r vela.Router) {
	r.GET("/api/v1/arr/agent/arp/list", xEnv.Then(func(ctx *fasthttp.RequestCtx) error {
		table := Table()
		chunk, err := json.Marshal(table)
		if err != nil {
			return err
		}

		ctx.Write(chunk)
		return nil
	}))
}
