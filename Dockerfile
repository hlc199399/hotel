FROM registry.cn-shenzhen.aliyuncs.com/lelelailele/golang:v1.0 AS builder


ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOPROXY https://goproxy.cn,direct

WORKDIR /go/src/hotel

COPY . .
COPY ./etc/hotel-api.yaml ./hotel-api.yaml

RUN CGO_ENABLED=0 GOOS-linux go build -a -installsuffix cgo -o app .

FROM registry.cn-shenzhen.aliyuncs.com/lelelailele/alpine:v1.0

RUN echo "http://mirrors.ustc.edu.cn/alpine/V3.10/main" > /etc/apk/repositories
RUN echo "http://mirrors.ustc.edu.cn/alpine/v3.10/community" >> /etc/apk/repositories
RUN apk - -no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/hotel/app .

EXPOSE 8080
ENTRYPOINT ["./app"]