FROM golang:1.20.6 AS build

WORKDIR /go/src/wawan/endeus

COPY . .

RUN go get .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

# deployment image
FROM alpine:latest
RUN apk --no-cache add ca-certificates

LABEL author="Wawan Randianto"

WORKDIR /root/
COPY --from=build /go/src/wawan/endeus/app .

CMD [ "./app" ]

EXPOSE 8080