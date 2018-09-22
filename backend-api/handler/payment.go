package handler

import (
	"context"
	"net/http"
	"strconv"

	"google.golang.org/grpc"

	"github.com/yusukemisa/go-vue-pay-grpc/backend-api/db"
	"github.com/yusukemisa/go-vue-pay-grpc/backend-api/domain"
	gpay "github.com/yusukemisa/go-vue-pay-grpc/payment-service/proto"
)

const addr = "localhost:50051"

//Charge exec payment-service charge
func Charge(ctx Context) {
	//パラメータや body をうけとる
	t := domain.Payment{}
	ctx.Bind(&t)

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}
	// id から item情報所得
	item, err := db.SelectItem(int64(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}
	greq := &gpay.PayRequest{
		Id:          int64(id),
		Token:       t.Token,
		Amount:      item.Amount,
		Name:        item.Name,
		Description: item.Description,
	}

	//IPアドレス(ここではlocalhost)とポート番号(50051)を指定して、サーバーと接続する
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}
	defer conn.Close()

	client := gpay.NewPayManagerClient(conn)
	// gRPCマイクロサービスの支払い処理関数を叩く
	pres, err := client.Charge(context.Background(), greq)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}
	ctx.JSON(http.StatusSeeOther, pres)
}
