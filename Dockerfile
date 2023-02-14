FROM golang:1.20
WORKDIR /go/src/TechRead
COPY . /go/src/TechRead
RUN go mod tidy
# Airをインストール
RUN go install github.com/cosmtrek/air@v1.27.3
# RUN chmod +x /go/src/TechRead/start.sh
CMD ["chmod","+x","/go/src/TechRead/start.sh"]

# airコマンドでGoファイルを起動
# CMD ["/go/src/TechRead/start.sh"]