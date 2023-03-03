FROM golang:latest
WORKDIR /go/src/TechRead
COPY . /go/src/TechRead
RUN go mod tidy
# Airをインストール
RUN go install github.com/cosmtrek/air@latest
