package handler

import (
	"net/http"
	"strconv"

	"github.com/yusukemisa/go-vue-pay-grpc/backend-api/db"
)

// GetLists - get all items
func GetLists(ctx Context) {
	items, err := db.SelectAllItems()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}
	ctx.JSON(http.StatusOK, items)
}

// GetItem - get item by id
func GetItem(ctx Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}
	item, err := db.SelectItem(int64(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}
	ctx.JSON(http.StatusOK, item)
}
