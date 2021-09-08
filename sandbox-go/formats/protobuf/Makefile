GOPATH:=$(shell go env GOPATH)

proto-general:
	protoc \
		-I . \
		-I /usr/local/include \
		--go_out=. \
		example/example.proto \
		--experimental_allow_proto3_optional


proto-validation:
	protoc \
		-I . \
		-I /usr/local/include \
		-I ${GOPATH}/src \
		-I ${GOPATH}/src/include \
		--go_out=. \
		--validate_out="lang=go:." \
		validation/validation.proto \

proto: proto-general proto-validation
