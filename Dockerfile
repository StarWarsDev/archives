#FROM golang:1.15 as test
#ENV CGO_ENABLED=0
##RUN apk add --no-cache gcc
#WORKDIR /src
#COPY . .
#RUN go test ./...

FROM golang:1.15 as build
#RUN apk add --no-cache gcc build-essential
WORKDIR /src
COPY . .
RUN go build -o bin/archives main.go

FROM ubuntu
WORKDIR /archives
COPY --from=build /src/bin/archives .

#RUN apk update \
#  && apk upgrade \
#  && apk add --no-cache ca-certificates \
#  && update-ca-certificates 2>/dev/null || true

RUN apt update -y \
    && apt install ca-certificates -y \
    && update-ca-certificates 2>/dev/null || true

EXPOSE 3001

CMD ["./archives"]