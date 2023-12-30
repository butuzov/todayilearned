package testing

import (
	"bytes"
	"io"
	"os"
)

// Error checking omitted for brevity

func readFile(location string) []byte {
	var buf bytes.Buffer
	f, err := os.Open(location)
	if err != nil {
		return nil
	}
	defer f.Close()

	if _, err := io.Copy(&buf, f); err != nil {
		return nil
	}

	return buf.Bytes()
}

func writeFile(location string, r io.Reader) error {
	f, err := os.OpenFile(location, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err := io.Copy(f, r); err != nil {
		return err
	}

	return nil
}
