BACK_BUILD = go build -o backend_api ./backend-api/main.go
BACK_PID = ./backend_api.pid

SERVER_BUILD = go build -o grpc_server ./payment-service/server/server.go
SERVER_PID = ./grpc_server.pid

FRONT_PID = ./front.pid

.PHONY: test
test: itemdb_mysql
	go test -v github.com/yusukemisa/go-vue-pay-grpc/backend-api/db/...

is_itemdb_mysql_launch=`docker ps | grep "itemdb_mysql" | wc -l`

.PHONY: itemdb_mysql
itemdb_mysql:
	@if [  -z "$(is_itemdb_mysql_launch)" ]; then \
		@echo "itemdb_mysql start"; \
		docker run -p 3306:3306 --name=itemdb_mysql -d itemdb_mysql; \
	fi
	@if [ ! -z "$(is_itemdb_mysql_launch)" ]; then echo "itemdb_mysql already started"; fi

.PHONY: start
start: stop front_start backend_api_build grpc_server
	./backend_api & echo $$! > $(BACK_PID)
	./grpc_server & echo $$! > $(SERVER_PID)

.PHONY: front_start
front_start:
	cd ./frontend-spa; npm start & echo $$! > ../$(FRONT_PID) ; cd .. ;\

backend_api_build:
	@$(BACK_BUILD)

grpc_server:
	@$(SERVER_BUILD)

.PHONY: stop
stop:
	@if [ -e $(BACK_PID) ]; then\
		kill `cat $(BACK_PID)` 2> /dev/null || true; rm -f ${BACK_PID} ./backend_api ;\
	fi
	@if [ -e $(SERVER_PID) ]; then\
		kill `cat $(SERVER_PID)` 2> /dev/null || true; rm -f ${SERVER_PID} ./grpc_server ;\
	fi
	@if [ -e $(FRONT_PID) ]; then\
		kill `cat $(FRONT_PID)` 2> /dev/null || true; rm -f ${FRONT_PID} ;\
	fi

# go test -v hoge/...　 hoge/以下のテストを実行
# go test packagePath  # cf. go test github.com/hoge/moge パッケージを指定
# go test -run ''      # カレントディレクトリ内の全てのテストを実行
# go test -run Foo     # Fooを含むテスト関数を実行
# 
