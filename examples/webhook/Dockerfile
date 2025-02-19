# 1. Go 빌드 단계
FROM golang:1.23 AS builder

WORKDIR /app

# 필요한 모듈 다운로드
COPY go.mod go.sum ./
RUN go mod download

# 소스 코드 복사 및 빌드
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o webhook-server main.go

# 2. 런타임 컨테이너 (최소한의 크기로 실행)
FROM alpine:latest

WORKDIR /root/

# SSL 인증서 폴더 생성
#RUN mkdir -p /etc/certs

# 인증서 및 실행 파일 복사
COPY --from=builder /app/webhook-server .
#COPY tls.crt tls.key /etc/certs/

# 실행 포트 설정
EXPOSE 8080

# Webhook 서버 실행
CMD ["./webhook-server"]