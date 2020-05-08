FROM golang
COPY . /go/src/github.com/eyedeekay/echo
WORKDIR /go/src/github.com/eyedeekay/echo
RUN go build
ENTRYPOINT [ "/go/src/github.com/eyedeekay/echo/echo" ]