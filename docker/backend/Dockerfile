FROM golang:1.19.4 as dev
WORKDIR /go/src/app
RUN go install github.com/cosmtrek/air@latest
# CMD ["air", "-c", ".air.toml"]