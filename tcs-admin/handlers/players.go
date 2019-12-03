package handlers

import (
	"github.com/finalist736/gokit/database"
	"github.com/finalist736/gokit/logger"
	"github.com/finalist736/gokit/response"
	"github.com/gocraft/web"
	"strconv"
	"tcs/models"
	"tcs/tcs-admin/tcsctx"
)

const playersPerPage = 10

type playersGetsResponse struct {
	Players      []*models.Player `json:"Players"`
	Pages        int              `json:"Pages"`
	TotalPlayers int              `json:"TotalPlayers"`
}

func PlayersGet(ctx *tcsctx.Ctx, rw web.ResponseWriter, req *web.Request) {

	var (
		page uint64 = 1
		err  error
	)

	page, err = strconv.ParseUint(req.FormValue("page"), 10, 64)
	if err != nil {
		page = 1
	}
	if page == 0 {
		page = 1
	}

	totalPages := 1
	offset := page*playersPerPage - playersPerPage

	session := database.GetDefaultSession()
	resp := &playersGetsResponse{}
	resp.Players = make([]*models.Player, 0, 100)
	_, err = session.
		Select("*").
		From("players").
		Limit(playersPerPage).
		Offset(offset).
		Load(&resp.Players)
	if err != nil {
		logger.StdErr().Errorf("select players error: %s", err)
		response.Error(response.ERROR_DBERROR, rw)
		return
	}
	rows, err := session.Query("select count(*) from players")
	if err != nil {
		logger.StdErr().Errorf("select count(*) error: %s", err)
		totalPages = 1
	} else {
		if rows.Next() {
			_ = rows.Scan(&totalPages)
		}
		resp.TotalPlayers = totalPages
		if totalPages%playersPerPage == 0 {
			totalPages = totalPages / playersPerPage
		} else {
			totalPages = totalPages/playersPerPage + 1
		}
		_ = rows.Close()
	}
	resp.Pages = totalPages
	response.JsonIntent(resp, rw)
}
