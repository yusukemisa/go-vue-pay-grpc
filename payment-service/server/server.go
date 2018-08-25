package main

import (
	"context"
	"log"
	"net"
	"os"

	payjp "github.com/payjp/payjp-go/v1"
	gpay "github.com/yusukemisa/go-vue-pay-grpc/payment-service/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const port = ":50051"

type server struct{}

func (s *server) Charge(ctx context.Context, req *gpay.PayRequest) (*gpay.PayResponse, error) {
	// PAI の初期化
	pay := payjp.New(os.Getenv("PAYJP_TEST_SECRET_KEY"), nil)
	// 支払い
	charge, err := pay.Charge.Create(
		int(req.Amount), //支払い金額
		payjp.Charge{ //設定
			Currency:    "jpy",
			CardToken:   req.Token,
			Capture:     true,
			Description: req.Name + ":" + req.Description,
		},
	)
	if err != nil {
		return nil, err
	}

	//支払い結果からレスポンス作成
	res := &gpay.PayResponse{
		Paid:     charge.Paid,
		Captured: charge.Captured,
		Amount:   int64(charge.Amount),
	}

	return res, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	//grpcサーバーの作成
	s := grpc.NewServer()
	//ペイメントサービスにgrpcサーバーを登録
	gpay.RegisterPayManagerServer(s, &server{})

	//?
	reflection.Register(s)
	log.Printf("gRPC Server started: localhost%s\n", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
