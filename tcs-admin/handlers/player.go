package handlers

import (
	"encoding/json"
	"github.com/finalist736/gokit/database"
	"github.com/finalist736/gokit/logger"
	"github.com/finalist736/gokit/response"
	"github.com/gocraft/dbr"
	"github.com/gocraft/web"
	"io/ioutil"
	"strconv"
	"tcs/models"
	"tcs/tcs-admin/tcsctx"
)

func PlayerGet(ctx *tcsctx.Ctx, rw web.ResponseWriter, req *web.Request) {

	var (
		err         error
		playerIDStr string
		playerID    int64
	)

	playerIDStr = req.FormValue("id")

	playerID, err = strconv.ParseInt(playerIDStr, 10, 64)
	if err != nil || playerID == 0 {
		logger.StdOut().Warnf(
			"Player parse playerID error: %s string: %s",
			err, playerIDStr)
		response.Error(response.ERROR_REQUEST_DATA, rw)
		return
	}

	session := database.GetDefaultSession()
	resp := &models.Player{}
	err = session.
		Select("*").
		From("players").
		Where("tg=?", playerID).LoadOne(resp)
	if err != nil {
		if err == dbr.ErrNotFound {
			response.Error(response.ERROR_NO_SUCH_USER, rw)
			return
		}
		logger.StdOut().Warnf(
			"Select player error: %s string: %s",
			err, playerIDStr)
		response.Error(response.ERROR_DBERROR, rw)
		return
	}

	response.JsonIntent(resp, rw)
}

func PlayerNameSet(ctx *tcsctx.Ctx, rw web.ResponseWriter, req *web.Request) {
	// TODO call bot RPC for clearing cache
	var (
		err        error
		itemPlayer *models.Player
	)

	defer req.Body.Close()

	itemPlayer = new(models.Player)
	ba, err := ioutil.ReadAll(req.Body)
	if err != nil {
		logger.StdErr().Errorf("PlayerNameSet read body error: %s", err)
		response.ErrorBadRequest(response.ERROR_NO_CONTENT, rw)
		return
	}

	err = json.Unmarshal(ba, itemPlayer)
	if err != nil {
		logger.StdErr().Errorf("PlayerNameSet json error: %s", err)
		response.ErrorBadRequest(response.ERROR_REQUEST_DATA, rw)
		return
	}

	session := database.GetDefaultSession()
	_, err = session.
		Update(models.PlayerTableName).
		Set("name", itemPlayer.Name).
		Where("tg=?", itemPlayer.Tg).
		Exec()
	if err != nil {
		logger.StdErr().Errorf("PlayerNameSet update error: %s", err)
		response.ErrorInternalServer(response.ERROR_DBERROR, rw)
		return
	}
	response.Error("", rw)
}

func PlayerLocationSet(ctx *tcsctx.Ctx, rw web.ResponseWriter, req *web.Request) {
	// TODO call bot RPC for clearing cache
	var (
		err          error
		itemLocation *models.Player
	)

	defer req.Body.Close()

	itemLocation = new(models.Player)
	ba, err := ioutil.ReadAll(req.Body)
	if err != nil {
		logger.StdErr().Errorf("PlayerLocationSet read body error: %s", err)
		response.ErrorBadRequest(response.ERROR_NO_CONTENT, rw)
		return
	}

	err = json.Unmarshal(ba, itemLocation)
	if err != nil {
		logger.StdErr().Errorf("PlayerLocationSet json error: %s", err)
		response.ErrorBadRequest(response.ERROR_REQUEST_DATA, rw)
		return
	}

	session := database.GetDefaultSession()
	_, err = session.
		Update(models.PlayerTableName).
		Set("location", itemLocation.Location).
		Set("posx", itemLocation.PosX).
		Set("posy", itemLocation.PosY).
		Where("tg=?", itemLocation.Tg).
		Exec()
	if err != nil {
		logger.StdErr().Errorf("PlayerLocationSet update error: %s", err)
		response.ErrorInternalServer(response.ERROR_DBERROR, rw)
		return
	}
	response.Error("", rw)
}

func PlayerAirSet(ctx *tcsctx.Ctx, rw web.ResponseWriter, req *web.Request) {
	// TODO call bot RPC for clearing cache
	var (
		err     error
		itemAir *models.Player
	)

	defer req.Body.Close()

	itemAir = new(models.Player)
	ba, err := ioutil.ReadAll(req.Body)
	if err != nil {
		logger.StdErr().Errorf("PlayerAirSet read body error: %s", err)
		response.ErrorBadRequest(response.ERROR_NO_CONTENT, rw)
		return
	}

	err = json.Unmarshal(ba, itemAir)
	if err != nil {
		logger.StdErr().Errorf("PlayerAirSet json error: %s", err)
		response.ErrorBadRequest(response.ERROR_REQUEST_DATA, rw)
		return
	}

	session := database.GetDefaultSession()
	_, err = session.
		Update(models.PlayerTableName).
		Set("air", itemAir.Air).
		Where("tg=?", itemAir.Tg).
		Exec()
	if err != nil {
		logger.StdErr().Errorf("PlayerAirSet update error: %s", err)
		response.ErrorInternalServer(response.ERROR_DBERROR, rw)
		return
	}
	response.Error("", rw)
}

func PlayerBalloonSet(ctx *tcsctx.Ctx, rw web.ResponseWriter, req *web.Request) {
	// TODO call bot RPC for clearing cache
	var (
		err         error
		itemBalloon *models.Player
	)

	defer req.Body.Close()

	itemBalloon = new(models.Player)
	ba, err := ioutil.ReadAll(req.Body)
	if err != nil {
		logger.StdErr().Errorf("PlayerBalloonSet read body error: %s", err)
		response.ErrorBadRequest(response.ERROR_NO_CONTENT, rw)
		return
	}

	err = json.Unmarshal(ba, itemBalloon)
	if err != nil {
		logger.StdErr().Errorf("PlayerBalloonSet json error: %s", err)
		response.ErrorBadRequest(response.ERROR_REQUEST_DATA, rw)
		return
	}

	session := database.GetDefaultSession()
	_, err = session.
		Update(models.PlayerTableName).
		Set("equipb", itemBalloon.EquipB).
		Where("tg=?", itemBalloon.Tg).
		Exec()
	if err != nil {
		logger.StdErr().Errorf("PlayerBalloonSet update error: %s", err)
		response.ErrorInternalServer(response.ERROR_DBERROR, rw)
		return
	}
	response.Error("", rw)
}

func PlayerConsoleSet(ctx *tcsctx.Ctx, rw web.ResponseWriter, req *web.Request) {
	// TODO call bot RPC for clearing cache
	var (
		err         error
		itemConsole *models.Player
	)

	defer req.Body.Close()

	itemConsole = new(models.Player)
	ba, err := ioutil.ReadAll(req.Body)
	if err != nil {
		logger.StdErr().Errorf("PlayerConsoleSet read body error: %s", err)
		response.ErrorBadRequest(response.ERROR_NO_CONTENT, rw)
		return
	}

	err = json.Unmarshal(ba, itemConsole)
	if err != nil {
		logger.StdErr().Errorf("PlayerConsoleSet json error: %s", err)
		response.ErrorBadRequest(response.ERROR_REQUEST_DATA, rw)
		return
	}

	session := database.GetDefaultSession()
	_, err = session.
		Update(models.PlayerTableName).
		Set("equipc", itemConsole.EquipC).
		Where("tg=?", itemConsole.Tg).
		Exec()
	if err != nil {
		logger.StdErr().Errorf("PlayerConsoleSet update error: %s", err)
		response.ErrorInternalServer(response.ERROR_DBERROR, rw)
		return
	}
	response.Error("", rw)
}
