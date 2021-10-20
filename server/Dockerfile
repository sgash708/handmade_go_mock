FROM golang:1.16

WORKDIR /api

COPY . .
ENV GO111MODULE=on

# API HotReload
RUN go get github.com/pilu/fresh
CMD ["fresh"]