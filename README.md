## go-vue-pay-grpc
TODO: Write your project about

## Features
Write your project features

## How to use

### 入手
```
$ go get github.com/yusukemisa/go-vue-pay-grpc
```

### protocを公式からダウンロードして/usr/local/bin配下にコマンドをおいとく

公式：https://github.com/google/protobuf/releases

```
cp -p  ~/Downloads/protoc-3/bin/protoc /usr/local/bin
```

### 必要なパッケージ
```
$ go get -u google.golang.org/grpc
$ go get -u github.com/golang/protobuf/protoc-gen-go
```

### protoc-gen-goはprotocコマンドと同じパスにコピーしておく
```
$ cp $GOPATH/bin/protoc-gen-go /usr/local/bin/
```

### .proto作成後.pb.goファイルの作成
ちょっとわからん
--go_out=plugins=grpc:.

とりあえずprotoファイルのパスを指定すればおk
```
$ protoc --go_out=plugins=grpc:. pay.proto

yusukemisawa at MacBook-Air in ~/dev/go/src/github.com/yusukemisa/go-vue-pay-grpc/payment-service/proto 
$ ll
total 32
-rw-r--r--  1 yusukemisawa  staff   8.2K  7 23 22:34 pay.pb.go
-rw-r--r--  1 yusukemisawa  staff   444B  7 23 22:34 pay.proto
```
pay.pb.goは基本protocで生成するもので編集はしない。
内部にprotoで定義したインターフェースのコードが作成されるのでそれを満たすような実装を行う。

### gRPCのインターフェースとなるAPIサーバーを作成する

| method | path | 機能|
| ------------ | ------------- | ------------- |
| GET | /api/v1/items |商品をすべて返却 |
| GET | /api/v1/items/:id | id で指定された商品情報を返却する|
| POST | /api/v1/charge/items/:id | id で指定された商品を購入する|






