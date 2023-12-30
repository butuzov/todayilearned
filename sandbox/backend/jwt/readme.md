# JSON Web Tokens

<!-- todo: dddd -->

JSON Web Tokens are an open, industry standard [RFC 7519](https://tools.ietf.org/html/rfc7519) method for representing claims securely between two parties.

## Client libraries, Apps & Etc

- library https://golang-jwt.github.io/jwt/
- decoder https://jwt.io/
- decoder https://token.dev/
- decoder https://fusionauth.io/dev-tools/jwt-decoder
- decoder https://dinochiesa.github.io/jwt/

## Videos

* [Dan Moore's talk at GO Berlin](https://www.youtube.com/watch?v=rArCF7nUcvY)
* [JWS & Refresh tokens (`rus`)](https://www.youtube.com/watch?v=Y41nrrMcOew)


## Example

```go
// sample token string taken from the New example
tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIiLCJuYmYiOjE0NDQ0Nzg0MDB9.u1riaD1rW97opCoAuRCTy4w58Br-Zk-bh7vLiRIsrpU"

// Parse takes the token string and a function for looking up the key. The latter is especially
// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
// head of the token to identify which key to use, but the parsed token (head and claims) is provided
// to the callback, providing flexibility.
token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
	// Don't forget to validate the alg is what you expect:
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
	}

	// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
	return hmacSampleSecret, nil
})
if err != nil {
	log.Fatal(err)
}

if claims, ok := token.Claims.(jwt.MapClaims); ok {
	fmt.Println(claims["foo"], claims["nbf"])
} else {
	fmt.Println(err)
}
```
