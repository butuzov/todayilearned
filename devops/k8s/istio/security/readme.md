<!-- weight: 50 -->
# Security

### Enable `service-to-service` communication inside cluster by using `mTLS`

```bash
# We deploying curl container so we can run `curl` commands within cluster
> k apply -f deploy/curl.yaml
# and exec into pod
> k exec -it curl sh
```

We also deploying httpbin as our test client

```bash
> k create deployment httpbin --image kennethreitz/httpbin  
> k expose deployment httpbin --port 80 --type=NodePort
```

Lets run command to ensure envoy injects sertificate information

```bash
/ $ curl -s httpbin/headers | grep -I X-Forwarded-Client-Cert
    "X-Forwarded-Client-Cert": "By=spiffe://cluster.local/ns/istio-course/sa/default;Hash=f89c461aa7d2d4d5c3295fb2a7efebffee2f21c991c235810db5937a67cd878e;Subject=\"\";URI=spiffe://cluster.local/ns/istio-course/sa/default"
```


This is because by default mutual TLS defined by `PeerAuthentication` CRD

```
apiVersion: security.istio.io/v1beta1
kind: PeerAuthentication
metadata:
  name: default
  namespace: istio-system 
spec:
  mtls:
    mode: STRICT  
```

#### Reading:

* https://banzaicloud.com/blog/istio-mtls/

### Integrate istio cluster with `Auth0` system

We going to create `RequestAuthentication` and `AuthorizationPolicy` that require JWT for any service with label `app: httpbin`

```yaml
apiVersion: security.istio.io/v1beta1
kind: RequestAuthentication
metadata:
  name: jwt-example
  namespace: istio-course
spec:
  selector:
    matchLabels:
      app: httpbin
  jwtRules:
  - issuer: test-issuer-1@istio.io
    jwksUri: https://example.com/demo.json

---

apiVersion: security.istio.io/v1beta1
kind: AuthorizationPolicy
metadata:
  name: httpbin
  namespace: istio-course
spec:
  selector:
    matchLabels:
      app: httpbin
  action: DENY
  rules:
  - from:
    - source:
        notRequestPrincipals: [ "*" ]

```

```shell
> curl -s $HTTPBIN_URL/headers
{
  "headers": {
    "Accept": "*/*",
    "Host": "192.168.99.101:31220",
    "User-Agent": "curl/7.64.1",
    "X-B3-Sampled": "1",
    "X-B3-Spanid": "a2239db6fe012a88",
    "X-B3-Traceid": "7d0630250efacee0a2239db6fe012a88"
  }
}
```

Applying `RequestAuthentication` and `AuthorizationPolicy`

```shell
k apply -f deploy/auth.yaml
```

```shell
> k apply -f deploy/auth.yaml
requestauthentication.security.istio.io/jwt-example configured
authorizationpolicy.security.istio.io/httpbin configured
```

### No Token Test

```shell
> curl http://192.168.99.101:31220/headers -v

RBAC: access denied
```

### Bad Token Test

```shell
 > curl --header "Authorization: Bearer foobar"  http://192.168.99.101:31220/headers -v

Jwt is not in the form of Header.Payload.Signature with two dots and 3 sections 
```

### Valid Token

```shell
TOKEN=$(curl https://example.com/data.jwt -s)
> curl --header "Authorization: Bearer $TOKEN"  http://192.168.99.101:31220/headers -v
{
  "headers": {
    "Accept": "*/*",
    "Host": "192.168.99.101:31220",
    "User-Agent": "curl/7.64.1",
    "X-B3-Sampled": "1",
    "X-B3-Spanid": "a2239db6fe012a88",
    "X-B3-Traceid": "7d0630250efacee0a2239db6fe012a88"
  }
}
```

#### Reading:

- https://github.com/istio/istio/blob/6a101c8939f841e8e3833a6d5e62d643c0be9185/security/tools/jwt/samples/README.md
- https://auth0.com/blog/securing-kubernetes-clusters-with-istio-and-auth0/
- https://rinormaloku.com/authorization-in-istio/
- https://medium.com/google-cloud/back-to-microservices-with-istio-part-2-authentication-authorization-b079f77358ac
