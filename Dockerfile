#REDO
FROM golang:1.19-alpine

LABEL author="ax-i-om"
LABEL github="https://github.com/ax-i-om/bitcrook"

RUN mkdir /app
COPY . /app
WORKDIR /app
RUN go mod download
RUN go build -o main .

ENTRYPOINT ["/app/main"]