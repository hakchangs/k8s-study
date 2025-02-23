# Harbor
> https://goharbor.io/docs/2.12.0/install-config/

인증서 발급 및 등록
```shell
openssl req -x509 -newkey rsa:4096 -sha256 -days 3650 -nodes \
  -keyout harbor.key -out harbor.crt -subj '/CN=harbor.local' \
  -addext 'subjectAltName=DNS:harbor.local'

cp harbor.crt harbor.key /etc/pki/ca-trust/source/anchors/
update-ca-trust

cat << EOF >> /etc/hosts
192.168.122.111 harbor.local
EOF

# secret 에도 등록
kubectl create ns harbor
kubectl create secret tls harbor-ingress-tls --key harbor.key --cert harbor.crt -n harbor
```

설치
```shell
helm upgrade -i harbor ./harbor -n harbor --create-namespace -f values.yaml
```

```shell
# 접속확인
nerdctl login harbor.local

# private registry 사용설정
cat << EOF > /etc/rancher/rke2/registries.yaml
mirrors:
  docker.io:
    endpoint:
      - "https://harbor.local"
configs:
  "harbor.local":
    auth:
      username: admin # this is the registry username
      password: Harbor12345 # this is the registry password
    tls:
      cert_file: /etc/pki/ca-trust/source/anchors/harbor.crt
      key_file: /etc/pki/ca-trust/source/anchors/harbor.key
EOF

systemctl restart rke2-server
```


### 참고: Proxy Cache 프로젝트 이용하기
Proxy Cache 를 이용하면 Airgap 환경에서 이미지주소 변경없이도 이미지를 불러올 수 있음.

**조건**
- Airgap 환경에서 Private Registry 로 Harbor 사용
- Harbor 에서는 외부 주요 Registry 연결은 오픈 (docker hub, quay 등)
- RKE2 로 설치하여 mirror 설정을 할 수 있어야함.

**처리방법**
```yaml
# https://github.com/rancher/rke2/discussions/5950
mirrors:
  "docker.io":
    endpoint:
      - "https://harbor.example.com:443"
    rewrite:
      "(.*)": "docker/$1"
  "quay.io":
    endpoint:
      - "https://harbor.example.com:443"
    rewrite:
      "(.*)": "quay/$1"
  "ghcr.io":
    endpoint:
      - "https://harbor.example.com:443"
    rewrite:
      "(.*)": "ghcr/$1"
```
1. 외부로 통하는 주요 Registry 는 모두 Proxy Cache 프로젝트로 만든다.
2. rke2 registries.yaml 파일에 mirrors 설정으로 proxy 프로젝트를 바라보도록 설정한다.

