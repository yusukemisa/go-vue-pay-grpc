package main

import (
	"os"

	"github.com/yusukemisa/go-vue-pay-grpc/infrastructure"
)

func main() {
	infrastructure.Router.Run(os.Getenv("API_SERVER_PORT"))
}
