# syntax=docker/dockerfile:1

FROM golang:1.16-alpine as build-env

RUN mkdir /app
WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /go/bin/docker-grpc-helloworld

FROM scratch
EXPOSE 8080
COPY --from=build-env /go/bin/docker-grpc-helloworld /go/bin/docker-grpc-helloworld
CMD [ "/go/bin/docker-grpc-helloworld" ]