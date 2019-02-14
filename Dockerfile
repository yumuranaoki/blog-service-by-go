FROM golang:latest
WORKDIR /go/src/github.com/yumruanaoki/blog-service-by-go
COPY . .
RUN go get github.com/jinzhu/gorm \
  && go get github.com/lib/pq
EXPOSE 8080