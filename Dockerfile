FROM golang:1.19-alpine AS builder
RUN apk update && apk add --no-cache git
WORKDIR /btc-billionaire
COPY . /btc-billionaire/
RUN go mod tidy
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o btc-billionaire cmd/main.go

FROM alpine
RUN mkdir /server
WORKDIR /server
COPY --from=builder /btc-billionaire/btc-billionaire /server/btc-billionaire
COPY --from=builder /btc-billionaire/app.env /server/app.env
EXPOSE 8089
CMD [ "/server/btc-billionaire" ]