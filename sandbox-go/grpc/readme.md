# gRPC examples

## Example 1: `HelloWorld` & `grpcurl`

  Run the server (and proto generation) with `make helloworld`

  ```shell
  export GRPC_URL=... # put address provided with make message: listening on ...

  # -- Request a list of services.
  > grpcurl -plaintext -v $GRPC_URL list

  grpc.reflection.v1alpha.ServerReflection
  helloworld.Greeter


  # -- Request a list of methods of the service `helloworld.Greeter`
  > grpcurl -plaintext -v $GRPC_URL list helloworld.Greeter

  helloworld.Greeter.SayHello


  # -- Describe the service
  > grpcurl -plaintext -v $GRPC_URL describe helloworld.Greeter

  helloworld.Greeter is a service:
  service Greeter {
    rpc SayHello ( .helloworld.HelloRequest ) returns ( .helloworld.HelloReply );
  }


  # -- Describe the service method
  > grpcurl -plaintext -v $GRPC_URL describe helloworld.Greeter.SayHello

  rpc SayHello ( .helloworld.HelloRequest ) returns ( .helloworld.HelloReply );


  # -- Describe the message
  > grpcurl -plaintext -v $GRPC_URL describe helloworld.HelloRequest

  helloworld.HelloRequest is a message:
  message HelloRequest {
    string name = 1;
  }

  # -- Query Method
  grpcurl -plaintext -d '{"name": "Світ!"}' $GRPC_URL helloworld.Greeter/SayHello

  {
    "message": "Hello Світ!"
  }

  # is same as...
  grpcurl -plaintext -d @ $GRPC_URL helloworld.Greeter/SayHello <<EOM
    {"name": "Світ!"}
  EOM
  ```


## Example 2: `HelloWorld` & `evans`

Run repl for the same address we running server on.

```shell
evans -r repl --host 127.0.0.1 --port 65095
> show package
> call SayHello
> help
... etc
```

## Example 3: Unary Streams

[Simple Ping Pong](simple-rpc)

```shell
make unary-demo
```

## Example 4: Server Streaming

[Fibonacci Sequence Generator](server-streaming)

```shell
make server-streaming
```

## Example 5: Client Streaming


[AVG for client sequence](client-streaming)

```shell
make client-streaming
```

## Todo: Example 6: Bi-Directional Streaming

[AVG of lst 10 sequence calc](bidirectional-streaming)

```shell
make bidirectional-streaming
```
