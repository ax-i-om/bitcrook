#REDO
FROM golang:1.19-alpine

LABEL author="audioo"
LABEL github="https://github.com/audioo/bitcrook"

RUN mkdir /app
COPY . /app
WORKDIR /app
RUN go mod download
RUN go build -o main .

ENTRYPOINT ["/app/main"]