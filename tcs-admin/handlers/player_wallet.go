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

func PlayersWallet(ctx *tcsctx.Ctx, rw web.ResponseWriter, req *web.Request) {

	var (
		err         error
		playerIDStr string
		playerID    int64
	)

	playerIDStr = req.FormValue("id")

	playerID, err = strconv.ParseInt(playerIDStr, 10, 64)
	if err != nil || playerID == 0 {
		logger.StdOut().Warnf(
			"PlayersWallet parse playerID error: %s string: %s",
			err, playerIDStr)
		response.ErrorBadRequest(response.ERROR_REQUEST_DATA, rw)
		return
	}

	session := database.GetDefaultSession()
	resp := &models.PlayerWallet{}
	_, err = session.
		Select("*").
		From(models.PlayerWalletTableName).
		Where("tg=?", playerID).Load(resp)
	if err != nil {
		logger.StdOut().Warnf(
			"PlayersWallet Select player wallet error: %s string: %s",
			err, playerIDStr)
		response.ErrorInternalServer(response.ERROR_DBERROR, rw)
		return
	}
	response.JsonIntent(resp, rw)
}

func PlayersWalletSet(ctx *tcsctx.Ctx, rw web.ResponseWriter, req *web.Request) {
	// TODO call bot RPC for clearing cache
	var (
		err        error
		itemWallet *models.PlayerWallet
	)

	defer req.Body.Close()

	itemWallet = new(models.PlayerWallet)
	ba, err := ioutil.ReadAll(req.Body)
	if err != nil {
		logger.StdErr().Errorf("PlayersWalletSet read body error: %s", err)
		response.ErrorBadRequest(response.ERROR_NO_CONTENT, rw)
		return
	}

	err = json.Unmarshal(ba, itemWallet)
	if err != nil {
		logger.StdErr().Errorf("PlayersWalletSet json error: %s", err)
		response.ErrorBadRequest(response.ERROR_REQUEST_DATA, rw)
		return
	}

	session := database.GetDefaultSession()
	_, err = session.
		Update(models.PlayerWalletTableName).
		Set("credits", itemWallet.Credits).
		Set("gold", itemWallet.Gold).
		Where("tg=?", itemWallet.Tg).
		Exec()
	if err != nil {
		logger.StdErr().Errorf("PlayersWalletSet update error: %s", err)
		response.ErrorInternalServer(response.ERROR_DBERROR, rw)
		return
	}
	response.Error("", rw)
}
