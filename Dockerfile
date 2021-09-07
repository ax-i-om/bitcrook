FROM golang:1.16.4-alpine3.12

LABEL author="Maraudery"
LABEL github="https://github.com/maraudery/qualear"

RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go mod download
RUN go build -o main .

ENTRYPOINT ["/app/main"]