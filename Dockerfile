FROM golang:1.13-alpine as test
ENV CGO_ENABLED=0
#RUN apk add --no-cache make
WORKDIR /src
COPY . .
RUN go test ./...

FROM golang:1.13-alpine as build
#RUN apk add --no-cache make
WORKDIR /src
COPY . .
RUN go build -o bin/archives main.go

FROM alpine
WORKDIR /archives
COPY --from=build /src/bin/archives .

RUN apk update \
  && apk upgrade \
  && apk add --no-cache ca-certificates \
  && update-ca-certificates 2>/dev/null || true
EXPOSE 3001

CMD ["./archives"]