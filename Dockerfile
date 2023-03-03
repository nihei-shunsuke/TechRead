FROM golang:1.20
WORKDIR /go/src/TechRead
COPY . /go/src/TechRead
RUN go mod tidy
# Airをインストール
RUN go install github.com/cosmtrek/air@v1.27.3
