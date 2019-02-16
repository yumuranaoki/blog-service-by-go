FROM golang:latest
WORKDIR /go/src/github.com/yumuranaoki/blog-service-by-go
COPY . .
RUN go get github.com/jinzhu/gorm \
  && go get github.com/lib/pq \
  && go get github.com/go-redis/redis \
  && go get github.com/gorilla/mux \
  && go get github.com/gorilla/securecookie
EXPOSE 8080