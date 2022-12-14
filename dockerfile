# https://dev.to/chseki/build-a-super-minimalistic-docker-image-to-run-your-golang-app-33j0

FROM golang:alpine as builder

WORKDIR /app 

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" .

FROM scratch

WORKDIR /app

COPY --from=builder /app/dev-to /usr/bin/

ENTRYPOINT ["dev-to"]