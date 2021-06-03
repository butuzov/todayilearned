
# Methods ======================================================================

gawk_exists:
	@if [ -z "$(shell which gawk)" ]; then\
		echo  "GNU Awk Required";\
		exit 1;\
	fi

help: gawk_exists
	@echo "===================================================================="
	@echo " Makefile:  gRPC examples                                           "
	@echo "===================================================================="
	@cat $(MAKEFILE_LIST) | \
		grep -E '^# ~~~ .*? [~]+$$|^[a-zA-Z0-9_-]+:.*?## .*$$' | \
		gawk '{if ( $$1=="#" ) { \
			match($$0, /^# ~~~ (.+?) [~]+$$/, a);\
			{print "\n", a[1], ""}\
		} else { \
			match($$0, /^([a-zA-Z/_-]+):.*?## (.*)$$/, a); \
			{printf "  - \033[32m%-20s\033[0m %s\n",   a[1], a[2]} \
 		}}'
	@echo ""


# ~~~ Hello World ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

hello-world-proto: ## Generate Protobuf File
	protoc \
		-I . \
		-I /usr/local/include \
		-I ${GOPATH}/src \
		--go_out=helloworld/. \
		--go-grpc_out=helloworld/. \
		helloworld/proto/example.proto \
		--experimental_allow_proto3_optional


hello-world: hello-world-proto ## Run `Hello World` Example
	go run ./helloworld


# ~~~ Seample RPC (Unary) ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

unary-proto-generate: ## Generate Protobuf File
	@protoc \
		-I . \
		-I /usr/local/include \
		-I ${GOPATH}/src \
		--go_out=simple-rpc/. \
		--go-grpc_out=simple-rpc/. \
		simple-rpc/proto/unary.proto \
		--experimental_allow_proto3_optional

unary-client-build: unary-proto-generate ## Build Client
	go build -o simple-rpc/client simple-rpc/client.go

unary-client-run: unary-client-build ## Run Client
	@./simple-rpc/client &

unary-server-build: unary-proto-generate ## Build Server
	go build -o simple-rpc/server simple-rpc/server.go

unary-server-run: unary-server-build  ## Run Server
	@./simple-rpc/server &

unary-demo: unary-server-run unary-client-run ## Run Simple RPC Demo
	@sleep 2
	kill -9 `lsof -t -i :4772`
	unlink simple-rpc/server
	unlink simple-rpc/client


# ~~~ Server Streaming ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

ss-proto: ## Generate Protobuf
	protoc \
		-I . \
		-I /usr/local/include \
		-I ${GOPATH}/src \
		--go_out=server-streaming/. \
		--go-grpc_out=server-streaming/. \
		server-streaming/proto/serverstream.proto \
		--experimental_allow_proto3_optional

ss-client-build: ss-proto ## Build Client
	go build -o server-streaming/client server-streaming/client.go

ss-server-build: ss-proto ## Build Server
	go build -o server-streaming/server server-streaming/server.go

ss-client-run: ss-client-build ## Run Client
	@./server-streaming/client &

ss-server-run: ss-server-build ## Run Server
	@./server-streaming/server &

server-streaming: ss-server-run ss-client-run ## Run Demo
	@sleep 2
	@kill -9 `lsof -t -i :4772`
	@unlink server-streaming/server
	@unlink server-streaming/client

# ~~~ Clinet Streaming ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
cs-proto: ## Generate Protobuf
	protoc \
		-I . \
		-I /usr/local/include \
		-I ${GOPATH}/src \
		--go_out=client-streaming/. \
		--go-grpc_out=client-streaming/. \
		client-streaming/proto/clientstreaming.proto \
		--experimental_allow_proto3_optional

cs-client-build: cs-proto ## Build Client
	go build -o client-streaming/client client-streaming/client.go

cs-server-build: cs-proto ## Build Server
	go build -o client-streaming/server client-streaming/server.go

cs-client-run: cs-client-build ## Run Client
	@./client-streaming/client &

cs-server-run: cs-server-build ## Run Server
	@./client-streaming/server &

client-streaming: cs-server-run cs-client-run ## Run Demo
	@sleep 2
	@kill -9 `lsof -t -i :4772`
	@unlink client-streaming/server
	@unlink client-streaming/client

# ~~~ Bi-Directional Streaming ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

bd-proto: ## Generate Protobuf File
	protoc \
		-I . \
		-I /usr/local/include \
		-I ${GOPATH}/src \
		--go_out=bidirectional-streaming/. \
		--go-grpc_out=bidirectional-streaming/. \
		bidirectional-streaming/proto/bidirectional-streaming.proto \
		--experimental_allow_proto3_optional


bd-client-build: bd-proto ## Build Client
	go build -o bidirectional-streaming/client bidirectional-streaming/client.go

bd-server-build: bd-proto ## Build Server
	go build -o bidirectional-streaming/server bidirectional-streaming/server.go

bd-client-run: bd-client-build ## Run Client
	@./bidirectional-streaming/client &

bd-server-run: bd-server-build ## Run Server
	@./bidirectional-streaming/server &

bd-streaming: bd-server-run bd-client-run ## Run Demo
	@sleep 2
	@kill -9 `lsof -t -i :4772`
	@unlink bidirectional-streaming/server
	@unlink bidirectional-streaming/client
