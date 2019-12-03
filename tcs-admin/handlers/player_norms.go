package handlers

import (
	"encoding/json"
	"github.com/finalist736/gokit/database"
	"github.com/finalist736/gokit/logger"
	"github.com/finalist736/gokit/response"
	"github.com/gocraft/web"
	"io/ioutil"
	"strconv"
	"tcs/models"
	"tcs/tcs-admin/tcsctx"
)

type (
	PlayerNormsResponse struct {
		Item []*models.PlayerNorms `json:"Item"`
	}
)

func PlayerNormsGet(ctx *tcsctx.Ctx, rw web.ResponseWriter, req *web.Request) {

	var (
		err         error
		playerTgStr string
		playerTg    int64
	)

	playerTgStr = req.FormValue("tg")

	playerTg, err = strconv.ParseInt(playerTgStr, 10, 64)
	if err != nil || playerTg == 0 {
		logger.StdOut().Warnf(
			"PlayerNorms parse playerTg error: %s string: %s",
			err, playerTgStr)
		response.Error(response.ERROR_REQUEST_DATA, rw)
		return
	}

	session := database.GetDefaultSession()
	var resp PlayerNormsResponse
	_, err = session.
		Select("*").
		From(models.PlayerNormsTableName).
		Where("tg=?", playerTg).
		Load(&resp.Item)
	if err != nil {
		logger.StdOut().Warnf(
			"PlayerNormsGet Select norms error: %s string: %s",
			err)
		response.Error(response.ERROR_DBERROR, rw)
		return
	}
	response.JsonIntent(resp, rw)
}

func PlayerNormGet(ctx *tcsctx.Ctx, rw web.ResponseWriter, req *web.Request) {

	var (
		err         error
		playerTgStr string
		playerTg    int64
		dateStr string
	)

	playerTgStr = req.FormValue("tg")

	playerTg, err = strconv.ParseInt(playerTgStr, 10, 64)
	if err != nil || playerTg == 0 {
		logger.StdOut().Warnf(
			"PlayerNorm parse playerTg error: %s string: %s",
			err, playerTgStr)
		response.Error(response.ERROR_REQUEST_DATA, rw)
		return
	}

	dateStr = req.FormValue("date")

	session := database.GetDefaultSession()
	var resp PlayerNormsResponse
	_, err = session.
		Select("*").
		From(models.PlayerNormsTableName).
		Where("tg=? AND date=?", playerTg, dateStr).
		Load(&resp.Item)
	if err != nil {
		logger.StdOut().Warnf(
			"PlayerNormGet Select norms error: %s string: %s",
			err)
		response.Error(response.ERROR_DBERROR, rw)
		return
	}
	response.JsonIntent(resp, rw)
}


func PlayerNormSet(ctx *tcsctx.Ctx, rw web.ResponseWriter, req *web.Request) {
	// TODO call bot RPC for clearing cache
	var (
		err            error
		itemPlayerNorm *models.PlayerNorms
	)

	defer req.Body.Close()

	itemPlayerNorm = new(models.PlayerNorms)
	ba, err := ioutil.ReadAll(req.Body)
	if err != nil {
		logger.StdErr().Errorf("PlayerNormSet read body error: %s", err)
		response.ErrorBadRequest(response.ERROR_NO_CONTENT, rw)
		return
	}

	err = json.Unmarshal(ba, itemPlayerNorm)
	if err != nil {
		logger.StdErr().Errorf("PlayerNormSet json error: %s", err)
		response.ErrorBadRequest(response.ERROR_REQUEST_DATA, rw)
		return
	}

	session := database.GetDefaultSession()
	// session.InsertInto(«table»).Columns(models.PlayerColumns…).Record(itemPlayer).Exec()
	_, err = session.
		InsertInto(models.PlayerNormsTableName).
		Columns(models.PlayerNormsColumns...).
		Record(itemPlayerNorm).
		Exec()
	if err != nil {
		logger.StdErr().Errorf("PlayerNormSet update error: %s", err)
		response.ErrorInternalServer(response.ERROR_DBERROR, rw)
		return
	}
	response.Error("", rw)
}

func PlayerNormRemove(ctx *tcsctx.Ctx, rw web.ResponseWriter, req *web.Request) {
	// TODO call bot RPC for clearing cache
	var (
		err            error
		itemPlayerNorm *models.PlayerNorms
	)

	defer req.Body.Close()

	itemPlayerNorm = new(models.PlayerNorms)
	ba, err := ioutil.ReadAll(req.Body)
	if err != nil {
		logger.StdErr().Errorf("PlayerNormsRemove read body error: %s", err)
		response.ErrorBadRequest(response.ERROR_NO_CONTENT, rw)
		return
	}

	err = json.Unmarshal(ba, itemPlayerNorm)
	if err != nil {
		logger.StdErr().Errorf("PlayerNormsRemove json error: %s", err)
		response.ErrorBadRequest(response.ERROR_REQUEST_DATA, rw)
		return
	}

	session := database.GetDefaultSession()
	_, err = session.
		DeleteFrom(models.PlayerNormsTableName).
		Where("tg=? AND resource=? AND date=?",
			itemPlayerNorm.Tg, itemPlayerNorm.Resource, itemPlayerNorm.Date).
		Exec()
	if err != nil {
		logger.StdErr().Errorf("PlayerNormsRemove remove error: %s", err)
		response.ErrorInternalServer(response.ERROR_DBERROR, rw)
		return
	}
	response.Error("", rw)
}

func PlayerNormUpdate(ctx *tcsctx.Ctx, rw web.ResponseWriter, req *web.Request) {
	// TODO call bot RPC for clearing cache
	var (
		err            error
		itemPlayerNorm *models.PlayerNorms
	)

	defer req.Body.Close()

	itemPlayerNorm = new(models.PlayerNorms)
	ba, err := ioutil.ReadAll(req.Body)
	if err != nil {
		logger.StdErr().Errorf("PlayerNormUpdate read body error: %s", err)
		response.ErrorBadRequest(response.ERROR_NO_CONTENT, rw)
		return
	}

	err = json.Unmarshal(ba, itemPlayerNorm)
	if err != nil {
		logger.StdErr().Errorf("PlayerNormUpdate json error: %s", err)
		response.ErrorBadRequest(response.ERROR_REQUEST_DATA, rw)
		return
	}

	session := database.GetDefaultSession()
	_, err = session.
		Update(models.PlayerNormsTableName).
		Set("resource", itemPlayerNorm.Resource).
		Set("amount", itemPlayerNorm.Amount).
		Where("tg=? AND date=?", itemPlayerNorm.Tg, itemPlayerNorm.Date).
		Exec()
	if err != nil {
		logger.StdErr().Errorf("PlayerNormUpdate update error: %s", err)
		response.ErrorInternalServer(response.ERROR_DBERROR, rw)
		return
	}
	response.Error("", rw)
}
