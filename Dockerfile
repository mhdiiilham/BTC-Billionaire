FROM golang:1.19-alpine AS builder
RUN apk update && apk add --no-cache git
WORKDIR /btc-billionaire
COPY . /btc-billionaire/
RUN go mod tidy
RUN go build -o btc-billionaire cmd/main.go

FROM scratch
RUN mkdir /server
WORKDIR /server
COPY --from=builder /btc-billionaire/btc-billionaire .
COPY --from=builder /btc-billionaire/app.env .
EXPOSE 8089
CMD [ "/btc-billionaire" ]
