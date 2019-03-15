FROM golang
ADD . /go/src/github.com/apmath-web/expenses
WORKDIR /go/src
RUN go get -v github.com/gin-gonic/gin
RUN mkdir build
RUN go build -i -o ./build/server ./github.com/apmath-web/expenses/application.go
CMD ["./build/server"]