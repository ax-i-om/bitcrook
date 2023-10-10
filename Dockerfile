FROM golang:1.20-alpine

LABEL author="ax-i-om"
LABEL email="addressaxiom@pm.me"
LABEL github="https://github.com/ax-i-om/bitcrook"

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . ./

RUN go build -o /bitcrook .

EXPOSE 6174

CMD [ "/bitcrook", "server" ]