FROM golang:1.16.3-alpine3.12

LABEL author="Britton Hayes"
LABEL github="https://github.com/brittonhayes/pillager"

RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o main .

CMD ["/app/main"]