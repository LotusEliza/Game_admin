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
	PlayerInventoryResponse struct {
		Items []*models.PlayerInventory `json:"Items"`
	}
)

func PlayersInventory(ctx *tcsctx.Ctx, rw web.ResponseWriter, req *web.Request) {

	var (
		err         error
		playerIDStr string
		playerID    int64
	)

	playerIDStr = req.FormValue("tg")

	playerID, err = strconv.ParseInt(playerIDStr, 10, 64)
	if err != nil || playerID == 0 {
		logger.StdOut().Warnf(
			"PlayersInventory parse playerID error: %s string: %s",
			err, playerIDStr)
		response.Error(response.ERROR_REQUEST_DATA, rw)
		return
	}

	session := database.GetDefaultSession()
	var resp PlayerInventoryResponse
	_, err = session.
		Select("*").
		From(models.PlayerInventoryTableName).
		Where("tg=?", playerID).Load(&resp.Items)
	if err != nil {
		logger.StdOut().Warnf(
			"PlayersInventory Select player inventory error: %s string: %s",
			err, playerIDStr)
		response.Error(response.ERROR_DBERROR, rw)
		return
	}
	response.JsonIntent(resp, rw)
}

func PlayerInventorySet(ctx *tcsctx.Ctx, rw web.ResponseWriter, req *web.Request) {
	// TODO call bot RPC for clearing cache
	var (
		err           error
		itemInventory *models.PlayerInventory
	)

	defer req.Body.Close()

	itemInventory = new(models.PlayerInventory)
	ba, err := ioutil.ReadAll(req.Body)
	if err != nil {
		logger.StdErr().Errorf("PlayerInventorySet read body error: %s", err)
		response.ErrorBadRequest(response.ERROR_NO_CONTENT, rw)
		return
	}

	err = json.Unmarshal(ba, itemInventory)
	if err != nil {
		logger.StdErr().Errorf("PlayerInventorySet json error: %s", err)
		response.ErrorBadRequest(response.ERROR_REQUEST_DATA, rw)
		return
	}

	session := database.GetDefaultSession()
	// session.InsertInto(«table»).Columns(models.PlayerColumns…).Record(itemPlayer).Exec()
	_, err = session.
		InsertInto(models.PlayerInventoryTableName).
		Columns(models.PlayerInventoryColumns...).
		Record(itemInventory).
		Exec()
	if err != nil {
		logger.StdErr().Errorf("PlayerInventorySet update error: %s", err)
		response.ErrorInternalServer(response.ERROR_DBERROR, rw)
		return
	}
	response.Error("", rw)
}

func PlayerInventoryUpdate(ctx *tcsctx.Ctx, rw web.ResponseWriter, req *web.Request) {
	// TODO call bot RPC for clearing cache
	var (
		err          error
		itemInventory *models.PlayerInventory
	)

	defer req.Body.Close()

	itemInventory = new(models.PlayerInventory)
	ba, err := ioutil.ReadAll(req.Body)
	if err != nil {
		logger.StdErr().Errorf("PlayerInventoryUpdate read body error: %s", err)
		response.ErrorBadRequest(response.ERROR_NO_CONTENT, rw)
		return
	}

	err = json.Unmarshal(ba, itemInventory)
	if err != nil {
		logger.StdErr().Errorf("PlayerInventoryUpdate json error: %s", err)
		response.ErrorBadRequest(response.ERROR_REQUEST_DATA, rw)
		return
	}

	session := database.GetDefaultSession()
	_, err = session.
		Update(models.PlayerInventoryTableName).
		Set("itemvalue", itemInventory.ItemValue).
		Where("tg=? AND itemtype=?", itemInventory.Tg, itemInventory.ItemType).
		Exec()
	if err != nil {
		logger.StdErr().Errorf("PlayerInventoryUpdate update error: %s", err)
		response.ErrorInternalServer(response.ERROR_DBERROR, rw)
		return
	}
	response.Error("", rw)
}


func PlayerInventoryRemove(ctx *tcsctx.Ctx, rw web.ResponseWriter, req *web.Request) {
	// TODO call bot RPC for clearing cache
	var (
		err           error
		itemInventory *models.PlayerInventory
	)

	defer req.Body.Close()

	itemInventory = new(models.PlayerInventory)
	ba, err := ioutil.ReadAll(req.Body)
	if err != nil {
		logger.StdErr().Errorf("PlayerInventoryRemove read body error: %s", err)
		response.ErrorBadRequest(response.ERROR_NO_CONTENT, rw)
		return
	}

	err = json.Unmarshal(ba, itemInventory)
	if err != nil {
		logger.StdErr().Errorf("PlayerInventoryRemove json error: %s", err)
		response.ErrorBadRequest(response.ERROR_REQUEST_DATA, rw)
		return
	}

	session := database.GetDefaultSession()
	_, err = session.
		DeleteFrom(models.PlayerInventoryTableName).
		Where("tg=? AND itemtype=? AND itemvalue=?",
			itemInventory.Tg, itemInventory.ItemType,
			itemInventory.ItemValue).
		Exec()
	if err != nil {
		logger.StdErr().Errorf("PlayerInventoryRemove remove error: %s", err)
		response.ErrorInternalServer(response.ERROR_DBERROR, rw)
		return
	}
	response.Error("", rw)
}
