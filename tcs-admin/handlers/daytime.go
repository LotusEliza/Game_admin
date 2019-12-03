package handlers

import (
	"github.com/finalist736/gokit/response"
	"github.com/gocraft/web"
	"tcs/machanics/daytime"
	"tcs/tcs-admin/tcsctx"
)

type DTResponse struct {
	DayTime string
	Clocks  string
}

func DayTimeGet(ctx *tcsctx.Ctx, rw web.ResponseWriter, req *web.Request) {
	resp := &DTResponse{
		DayTime: daytime.Get().String(),
		Clocks:  "for example 17:59",
	}
	response.Json(resp, rw)
}
