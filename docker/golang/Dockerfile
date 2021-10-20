FROM golang:1.16

WORKDIR /api

# https://qiita.com/mk-tool/items/1c7e4929055bb3b7aeda#%E7%8F%BE%E5%9C%A8%E3%81%AE%E3%82%B3%E3%83%B3%E3%83%86%E3%82%AD%E3%82%B9%E3%83%88%E3%82%88%E3%82%8A%E3%82%82%E5%A4%96%E9%83%A8%E3%81%AE%E3%83%95%E3%82%A1%E3%82%A4%E3%83%AB%E3%81%AB%E3%82%A2%E3%82%AF%E3%82%BB%E3%82%B9%E3%81%A7%E3%81%8D%E3%81%AA%E3%81%84
COPY ./src/api /api
ENV GO111MODULE=on

RUN apt update \
  && go get -u github.com/go-delve/delve/cmd/dlv \
  && go get github.com/jinzhu/gorm \ 
  && go get github.com/go-sql-driver/mysql

# API HotReload
# RUN go get github.com/pilu/fresh
# CMD ["fresh"]
