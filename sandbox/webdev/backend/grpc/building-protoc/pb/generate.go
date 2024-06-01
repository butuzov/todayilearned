package pb

//go:generate protoc test.proto --go_out=. --go-grpc_out=. --go_opt=Mtest.proto=github.com/butuzov/exampleproto/pb --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative --go-grpc_opt=Mtest.proto=github.com/butuzov/exampleproto/pb
