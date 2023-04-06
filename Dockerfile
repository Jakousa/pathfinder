FROM golang:1.14 AS build

WORKDIR /usr/src/app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

FROM scratch

COPY --from=build /usr/src/app/server /usr/src/app/server

ENTRYPOINT [ "/usr/src/app/server" ]