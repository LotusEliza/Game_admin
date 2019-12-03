package tcsctx

import "github.com/gocraft/web"

type Ctx struct {
}

func (s *Ctx) Cors(rw web.ResponseWriter, req *web.Request, next web.NextMiddlewareFunc) {
	if req.Method == "OPTIONS" {
		rw.Header().Set("Access-Control-Allow-Origin", "*")
		rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		rw.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Requested-With")
		return
	} else {
		rw.Header().Set("Access-Control-Allow-Origin", "*")
	}
	next(rw, req)
}
