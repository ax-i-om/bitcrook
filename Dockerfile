FROM golang:1.16.4-alpine3.12

LABEL author="bitcrook"
LABEL github="https://github.com/bitcrook/goseek"

RUN mkdir /app
COPY . /app
WORKDIR /app
RUN go mod download
RUN go build -o main .

ENTRYPOINT ["/app/main"]