package httphandler

import (
	"crypto/rsa"
	"io/ioutil"
	"log"

	"github.com/dgrijalva/jwt-go"
)

const (
	privKeyPath = "keys/app.rsa"
	pubKeyPath  = "keys/app.rsa.pub"
)

var (
	verifyKey  *rsa.PublicKey
	signKey    *rsa.PrivateKey
	serverPort int

	// storing sample username/password pairs
	// don't do this on a real server
	users = map[string]string{
		"user": "pass",
		"ford": "betelgeuse7",
	}
)

type TokenClaim struct {
	*jwt.StandardClaims
	TokenType string
	Name      string
}

func fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// read the key files before starting http handlers
func init() {
	signBytes, err := ioutil.ReadFile(privKeyPath)
	fatal(err)

	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	fatal(err)

	verifyBytes, err := ioutil.ReadFile(pubKeyPath)
	fatal(err)

	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	fatal(err)
}
