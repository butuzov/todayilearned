package testdata

import "fmt"

func _() {
	_ = fmt.Sprintf("gopher://%s:%d", "myHost", 70) // want "host:port in url should be constructed with net.JoinHostPort and not directly with fmt.Sprintf"

	_ = fmt.Sprintf("telnet+ssl://%s:%d", "myHost", 23) // want "host:port in url should be constructed with net.JoinHostPort and not directly with fmt.Sprintf"

	_ = fmt.Sprintf("weird3.6://%s:%d", "myHost", 23) // want "host:port in url should be constructed with net.JoinHostPort and not directly with fmt.Sprintf"
}
