# Kube Apiserver

### REST 호출하기 - api 접근제어

> https://kubernetes.io/ko/docs/concepts/security/controlling-access/
> https://kubernetes.io/docs/reference/access-authn-authz/authentication/

방법1. x509 인증방식
```shell
curl --cacert /var/lib/rancher/rke2/server/tls/server-ca.crt \
     --cert /var/lib/rancher/rke2/server/tls/client-kube-apiserver.crt \
     --key /var/lib/rancher/rke2/server/tls/client-kube-apiserver.key \
     -X GET "https://127.0.0.1:6443/api/v1/nodes"
```

방법2. kubectl proxy
```shell
kubectl proxy --port 8080
curl http://127.0.0.1:8080/api/v1/nodes
```

방법3. ServiceAccount Token 호출
```shell
kubectl create serviceaccount test
curl -k -H "Authorization: Bearer $(kubectl create token test)" \
     https://127.0.0.1:6443/api/v1/nodes
```
- jwt 토큰을 동적으로 만들어 사용한다.

방법4. OIDC 연계
- OIDC 서비스 연동하여 Bearer 토큰으로 인증제공

