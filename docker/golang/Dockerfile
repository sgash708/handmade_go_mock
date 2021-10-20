FROM golang:1.16

WORKDIR /api

COPY src/api .
ENV GO111MODULE=on

RUN apt update \
  && go get github.com/jinzhu/gorm \ 
  && go get github.com/go-sql-driver/mysql


# API HotReload
RUN go get github.com/pilu/fresh
CMD ["fresh"]

EXPOSE 8080