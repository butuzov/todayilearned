# JSON Web Tokens

JSON Web Tokens are an open, industry standard [RFC 7519](https://tools.ietf.org/html/rfc7519) method for representing claims securely between two parties.

## Videos

* [Dan Moore's talk at GO Berlin](https://www.youtube.com/watch?v=rArCF7nUcvY)
* [JWS & REfresh tokens (`rus`)](https://www.youtube.com/watch?v=Y41nrrMcOew)

```go
$go get github.com/golang-jwt/jwt/v4
```

```go
import (
    jwt "github.com/golang-jwt/jwt/v4"
    "time"
    "fmt"
)
```

# `Go`'s Transcript of intro tutorial

https://jwt.io/introduction

In its compact form, JSON Web Tokens consist of three parts separated by dots (.), which are:

* Header
* Payload
* Signature

Therefore, a JWT typically looks like the following.

`xxxxx.yyyyy.zzzzz`

Let's break down the different parts.

### Header

The header typically consists of two parts: the type of the token, which is JWT, and the signing algorithm being used, such as HMAC SHA256 or RSA.

For example:

```json
{
  "alg": "HS256",
  "typ": "JWT"
}
```

Then, this JSON is Base64Url encoded to form the first part of the JWT.

### Payload 

The second part of the token is the payload, which contains the claims. Claims are statements about an entity (typically, the user) and additional data. There are three types of claims: registered, public, and private claims.

### Signature

To create the signature part you have to take the encoded header, the encoded payload, a secret, the algorithm specified in the header, and sign that.

For example if you want to use the HMAC SHA256 algorithm, the signature will be created in the following way:

```
HMACSHA256(
  base64UrlEncode(header) + "." +
  base64UrlEncode(payload),
  secret)
```

The signature is used to verify the message wasn't changed along the way, and, in the case of tokens signed with a private key, it can also verify that the sender of the JWT is who it says it is.

https://jwt.io/#debugger-io

### HMAC

```go
token := jwt.New(jwt.SigningMethodHS256)
claims := token.Claims.(jwt.MapClaims)
```

```go
claims["iss"] = "made.ua"
claims["exp"] = time.Now().Add(time.Minute * 5).Unix()
claims["name"] = "Oleg"
claims["roles"] = []string{"CAN_LIFT_THINGS"}
```

```go
claims

result >>> map[exp:1632051051 iss:made.ua name:Oleg roles:[CAN_LIFT_THINGS]]
```

```go
var secret = []byte("foobar")
```

```go
tokenString, _ := token.SignedString(secret)
```

```go
tokenString

result >>> eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MzIwNTEwNTEsImlzcyI6Im1hZGUudWEiLCJuYW1lIjoiT2xlZyIsInJvbGVzIjpbIkNBTl9MSUZUX1RISU5HUyJdfQ.Az2wz1I3jjOdlY1s0IVrzo_v0TxJ7otJcVMI3udei3s
```

Decoding `token`

```go
decodedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
    fmt.Printf("Before: %#v\n", token)
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
        return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
    }
    return secret, nil
});

stdout >>> Before: &jwt.Token{Raw:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MzIwNTEwNTEsImlzcyI6Im1hZGUudWEiLCJuYW1lIjoiT2xlZyIsInJvbGVzIjpbIkNBTl9MSUZUX1RISU5HUyJdfQ.Az2wz1I3jjOdlY1s0IVrzo_v0TxJ7otJcVMI3udei3s", Method:(*jwt.SigningMethodHMAC)(0xc000363a58), Header:map[string]interface {}{"alg":"HS256", "typ":"JWT"}, Claims:jwt.MapClaims{"exp":1.632051051e+09, "iss":"made.ua", "name":"Oleg", "roles":[]interface {}{"CAN_LIFT_THINGS"}}, Signature:"", Valid:false}
```

```go
fmt.Sprintf("After: %#v\n", decodedToken)

result >>> After: &jwt.Token{Raw:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MzIwNTEwNTEsImlzcyI6Im1hZGUudWEiLCJuYW1lIjoiT2xlZyIsInJvbGVzIjpbIkNBTl9MSUZUX1RISU5HUyJdfQ.Az2wz1I3jjOdlY1s0IVrzo_v0TxJ7otJcVMI3udei3s", Method:(*jwt.SigningMethodHMAC)(0xc000363a58), Header:map[string]interface {}{"alg":"HS256", "typ":"JWT"}, Claims:jwt.MapClaims{"exp":1.632051051e+09, "iss":"made.ua", "name":"Oleg", "roles":[]interface {}{"CAN_LIFT_THINGS"}}, Signature:"Az2wz1I3jjOdlY1s0IVrzo_v0TxJ7otJcVMI3udei3s", Valid:true}
```

### How decodede token looks like?

```go
decodedToken.Claims

result >>> map[exp:1.632051051e+09 iss:made.ua name:Oleg roles:[CAN_LIFT_THINGS]]
```

```go
decodedToken.Valid

result >>> true
```

```go
decodedToken.Signature

result >>> Az2wz1I3jjOdlY1s0IVrzo_v0TxJ7otJcVMI3udei3s
```

```go
decodedToken.Method

result >>> &{HS256 SHA-256}
```

```go
decodedToken.Raw

result >>> eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MzIwNTEwNTEsImlzcyI6Im1hZGUudWEiLCJuYW1lIjoiT2xlZyIsInJvbGVzIjpbIkNBTl9MSUZUX1RISU5HUyJdfQ.Az2wz1I3jjOdlY1s0IVrzo_v0TxJ7otJcVMI3udei3s
```

## RSA

```go
import (
	"crypto/rsa"
	"crypto/rand"
)
```

```go
privateKey, _ := rsa.GenerateKey(rand.Reader, 1024)
publicKey := privateKey.PublicKey
```

```go
token  := jwt.New(jwt.SigningMethodRS256)
claims := token.Claims.(jwt.MapClaims)
claims["iss"] = "made.ua"
claims["exp"] = time.Now().Add(time.Minute * 5).Unix()
claims["name"] = "Oleg"
claims["roles"] = []string{"CAN_LIFT_THINGS"}
```

```go
tokenString, err := token.SignedString(privateKey)
```

```go
tokenString

result >>> eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MzIwNTEwNTEsImlzcyI6Im1hZGUudWEiLCJuYW1lIjoiT2xlZyIsInJvbGVzIjpbIkNBTl9MSUZUX1RISU5HUyJdfQ.X7uanrVtDx3SkfZcJnM-vVPo_B5LeJIqAuQS6qjfg6dNu1nqy3CVuDTh_Arac1faOGTRNj3_Zx5eZUVCGEz7s-Og4VlLyFbENZfaeMzdXvzc3G7nJ87zdL6I6YfzTPgev_TXAVOgllvQdVULMEiS5gREqNrtgNj5Kt7G47Lgetc
```

```go
decodedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
    fmt.Printf("Before: %#v\n", token)
    if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
        return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
     }
    return &publicKey, nil
});

stdout >>> Before: &jwt.Token{Raw:"eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MzIwNTEwNTEsImlzcyI6Im1hZGUudWEiLCJuYW1lIjoiT2xlZyIsInJvbGVzIjpbIkNBTl9MSUZUX1RISU5HUyJdfQ.X7uanrVtDx3SkfZcJnM-vVPo_B5LeJIqAuQS6qjfg6dNu1nqy3CVuDTh_Arac1faOGTRNj3_Zx5eZUVCGEz7s-Og4VlLyFbENZfaeMzdXvzc3G7nJ87zdL6I6YfzTPgev_TXAVOgllvQdVULMEiS5gREqNrtgNj5Kt7G47Lgetc", Method:(*jwt.SigningMethodRSA)(0xc000363aa0), Header:map[string]interface {}{"alg":"RS256", "typ":"JWT"}, Claims:jwt.MapClaims{"exp":1.632051051e+09, "iss":"made.ua", "name":"Oleg", "roles":[]interface {}{"CAN_LIFT_THINGS"}}, Signature:"", Valid:false}
```

```go
fmt.Sprintf("After: %#v\n", decodedToken)

result >>> After: &jwt.Token{Raw:"eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MzIwNTEwNTEsImlzcyI6Im1hZGUudWEiLCJuYW1lIjoiT2xlZyIsInJvbGVzIjpbIkNBTl9MSUZUX1RISU5HUyJdfQ.X7uanrVtDx3SkfZcJnM-vVPo_B5LeJIqAuQS6qjfg6dNu1nqy3CVuDTh_Arac1faOGTRNj3_Zx5eZUVCGEz7s-Og4VlLyFbENZfaeMzdXvzc3G7nJ87zdL6I6YfzTPgev_TXAVOgllvQdVULMEiS5gREqNrtgNj5Kt7G47Lgetc", Method:(*jwt.SigningMethodRSA)(0xc000363aa0), Header:map[string]interface {}{"alg":"RS256", "typ":"JWT"}, Claims:jwt.MapClaims{"exp":1.632051051e+09, "iss":"made.ua", "name":"Oleg", "roles":[]interface {}{"CAN_LIFT_THINGS"}}, Signature:"X7uanrVtDx3SkfZcJnM-vVPo_B5LeJIqAuQS6qjfg6dNu1nqy3CVuDTh_Arac1faOGTRNj3_Zx5eZUVCGEz7s-Og4VlLyFbENZfaeMzdXvzc3G7nJ87zdL6I6YfzTPgev_TXAVOgllvQdVULMEiS5gREqNrtgNj5Kt7G47Lgetc", Valid:true}
```

### How decodede token looks like?

```go
decodedToken.Claims

result >>> map[exp:1.632051051e+09 iss:made.ua name:Oleg roles:[CAN_LIFT_THINGS]]
```

```go
decodedToken.Valid

result >>> true
```

```go
decodedToken.Signature

result >>> X7uanrVtDx3SkfZcJnM-vVPo_B5LeJIqAuQS6qjfg6dNu1nqy3CVuDTh_Arac1faOGTRNj3_Zx5eZUVCGEz7s-Og4VlLyFbENZfaeMzdXvzc3G7nJ87zdL6I6YfzTPgev_TXAVOgllvQdVULMEiS5gREqNrtgNj5Kt7G47Lgetc
```

```go
decodedToken.Method

result >>> &{RS256 SHA-256}
```

```go
decodedToken.Raw

result >>> eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MzIwNTEwNTEsImlzcyI6Im1hZGUudWEiLCJuYW1lIjoiT2xlZyIsInJvbGVzIjpbIkNBTl9MSUZUX1RISU5HUyJdfQ.X7uanrVtDx3SkfZcJnM-vVPo_B5LeJIqAuQS6qjfg6dNu1nqy3CVuDTh_Arac1faOGTRNj3_Zx5eZUVCGEz7s-Og4VlLyFbENZfaeMzdXvzc3G7nJ87zdL6I6YfzTPgev_TXAVOgllvQdVULMEiS5gREqNrtgNj5Kt7G47Lgetc
```

```go
out, ok := decodedToken.Claims.(jwt.MapClaims);
if ok && decodedToken.Valid {
    for k, i:= range out {
        fmt.Println(k, i, "vs", claims[k])
    }
}

stdout >>> exp 1.632051051e+09 vs 1632051051
stdout >>> iss made.ua vs made.ua
stdout >>> name Oleg vs Oleg
stdout >>> roles [CAN_LIFT_THINGS] vs [CAN_LIFT_THINGS]
```

```go
[]bool{
    out.VerifyIssuer("made.ua", true), 
    out.VerifyIssuer("made.de", true),
}

result >>> [true false]
```

* (`aud`) `VerifyAudience`
* (`exp`) `VerifyExpiresAt`
* (`iat`) `VerifyIssuedAt`
* (`iss`) `VerifyIssuer`
* (`nbf`) `VerifyNotBefore`