# webhook 예제

### **📌 빌드 & 실행 방법**
**1. Docker 이미지 빌드**
```sh
docker build -t my-webhook:v1 .
```

**2. 테스트 실행**
```sh
docker run --rm -p 8080:8080 my-webhook:v1
```

**3. 이미지 푸시 (Harbor 사용 시)**
```sh
docker tag my-webhook:v1 docker.local/my-webhook:v1
docker push docker.local/my-webhook:v1
```

### webhook 적용

**1. 인증서 생성**
```shell
openssl req -x509 -newkey rsa:4096 -keyout tls.key -out tls.crt -days 365 -nodes \
  -subj "/CN=webhook-service.default.svc" \
  -addext "subjectAltName=DNS:webhook-service.default.svc"

kubectl create secret tls webhook-tls --cert=tls.crt --key=tls.key -n default
```

**2. webhook 서비스 배포**
```shell
kubectl apply -f webhook-deploy.yaml
```

**3. 테스트**
```shell
kubectl apply -f webhook-test.yaml
```

