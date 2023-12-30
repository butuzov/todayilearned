package testing

import (
	"net"
	"net/rpc"
	"testing"
)

func TestConn(t testing.T) (net.Conn, net.Conn) {
	l, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	// Start a goroutine to accept our client connection
	var serverConn net.Conn
	doneCh := make(chan struct{})
	go func() {
		defer close(doneCh)
		defer l.Close()
		var err error
		serverConn, err = l.Accept()
		if err != nil {
			t.Fatalf("err: %s", err)
		}
	}()

	// Connect to the server
	clientConn, err := net.Dial("tcp", l.Addr().String())
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	// Wait for the server side to acknowledge it has connected
	<-doneCh

	return clientConn, serverConn
}

// TestRPCConn returns a rpc client and server connected to each other.
func TestRPCConn(t testing.T) (*rpc.Client, *rpc.Server) {
	t.Helper()

	clientConn, serverConn := TestConn(t)

	server := rpc.NewServer()
	go server.ServeConn(serverConn)

	client := rpc.NewClient(clientConn)
	return client, server
}
