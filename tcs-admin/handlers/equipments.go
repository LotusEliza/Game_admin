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
	EquipmentsResponse struct {
		Items []*models.Equipments `json:"Items"`
	}
)

func EquipmentsGet(ctx *tcsctx.Ctx, rw web.ResponseWriter, req *web.Request) {

	var (
		err error
	)

	session := database.GetDefaultSession()
	var resp EquipmentsResponse
	_, err = session.
		Select("*").
		From(models.EquipmentsTableName).
		Load(&resp.Items)
	if err != nil {
		logger.StdOut().Warnf(
			"EquipmentsGet Select equipments error: %s string: %s",
			err)
		response.Error(response.ERROR_DBERROR, rw)
		return
	}
	response.JsonIntent(resp, rw)
}

func EquipmentGet(ctx *tcsctx.Ctx, rw web.ResponseWriter, req *web.Request) {

	var (
		err         error
		playerIDStr string
		playerID    int64
	)

	playerIDStr = req.FormValue("id")

	playerID, err = strconv.ParseInt(playerIDStr, 10, 64)
	if err != nil || playerID == 0 {
		logger.StdOut().Warnf(
			"Equipments parse playerID error: %s string: %s",
			err, playerIDStr)
		response.Error(response.ERROR_REQUEST_DATA, rw)
		return
	}

	session := database.GetDefaultSession()
	var resp EquipmentsResponse
	_, err = session.
		Select("*").
		From(models.EquipmentsTableName).
		Where("id=?", playerID).Load(&resp.Items)
	if err != nil {
		logger.StdOut().Warnf(
			"Equipments Select player inventory error: %s string: %s",
			err, playerIDStr)
		response.Error(response.ERROR_DBERROR, rw)
		return
	}
	response.JsonIntent(resp, rw)
}

func EquipmentsSet(ctx *tcsctx.Ctx, rw web.ResponseWriter, req *web.Request) {
	// TODO call bot RPC for clearing cache
	var (
		err            error
		itemEquipments *models.Equipments
	)

	defer req.Body.Close()

	itemEquipments = new(models.Equipments)
	ba, err := ioutil.ReadAll(req.Body)
	if err != nil {
		logger.StdErr().Errorf("EquipmentsSet read body error: %s", err)
		response.ErrorBadRequest(response.ERROR_NO_CONTENT, rw)
		return
	}

	err = json.Unmarshal(ba, itemEquipments)
	if err != nil {
		logger.StdErr().Errorf("EquipmentsSet json error: %s", err)
		response.ErrorBadRequest(response.ERROR_REQUEST_DATA, rw)
		return
	}

	session := database.GetDefaultSession()
	_, err = session.
		InsertInto(models.EquipmentsTableName).
		Columns(models.EquipmentsColumns...).
		Record(itemEquipments).
		Exec()
	if err != nil {
		logger.StdErr().Errorf("EquipmentsSet update error: %s", err)
		response.ErrorInternalServer(response.ERROR_DBERROR, rw)
		return
	}
	response.Error("", rw)
}

func EquipmentsRemove(ctx *tcsctx.Ctx, rw web.ResponseWriter, req *web.Request) {
	// TODO call bot RPC for clearing cache
	var (
		err            error
		itemEquipments *models.Equipments
	)

	defer req.Body.Close()

	itemEquipments = new(models.Equipments)
	ba, err := ioutil.ReadAll(req.Body)
	if err != nil {
		logger.StdErr().Errorf("EquipmentsRemove read body error: %s", err)
		response.ErrorBadRequest(response.ERROR_NO_CONTENT, rw)
		return
	}

	err = json.Unmarshal(ba, itemEquipments)
	if err != nil {
		logger.StdErr().Errorf("EquipmentsRemove json error: %s", err)
		response.ErrorBadRequest(response.ERROR_REQUEST_DATA, rw)
		return
	}

	session := database.GetDefaultSession()
	_, err = session.
		DeleteFrom(models.EquipmentsTableName).
		Where("id=?",
			itemEquipments.ID).
		Exec()
	if err != nil {
		logger.StdErr().Errorf("EquipmentsRemove remove error: %s", err)
		response.ErrorInternalServer(response.ERROR_DBERROR, rw)
		return
	}
	response.Error("", rw)
}

func EquipmentUpdate(ctx *tcsctx.Ctx, rw web.ResponseWriter, req *web.Request) {
	// TODO call bot RPC for clearing cache
	var (
		err           error
		itemEquipment *models.Equipments
	)

	defer req.Body.Close()

	itemEquipment = new(models.Equipments)
	ba, err := ioutil.ReadAll(req.Body)
	if err != nil {
		logger.StdErr().Errorf("EquipmentUpdate read body error: %s", err)
		response.ErrorBadRequest(response.ERROR_NO_CONTENT, rw)
		return
	}

	err = json.Unmarshal(ba, itemEquipment)
	if err != nil {
		logger.StdErr().Errorf("EquipmentUpdate json error: %s", err)
		response.ErrorBadRequest(response.ERROR_REQUEST_DATA, rw)
		return
	}

	session := database.GetDefaultSession()
	_, err = session.
		Update(models.EquipmentsTableName).
		Set("title", itemEquipment.Title).
		Set("type", itemEquipment.Type).
		Set("subtype", itemEquipment.SubType).
		Set("buyprice", itemEquipment.BuyPrice).
		Set("sellprice", itemEquipment.SellPrice).
		Set("reputation", itemEquipment.Reputation).
		Set("damage", itemEquipment.Damage).
		Set("armor", itemEquipment.Armor).
		Set("air", itemEquipment.Air).
		Set("mine", itemEquipment.Mine).
		Set("time", itemEquipment.Time).
		Set("socketype", itemEquipment.SocketType).
		Set("sockets", itemEquipment.Sockets).
		Where("id=?", itemEquipment.ID).
		Exec()
	if err != nil {
		logger.StdErr().Errorf("EquipmentUpdate update error: %s", err)
		response.ErrorInternalServer(response.ERROR_DBERROR, rw)
		return
	}
	response.Error("", rw)
}
