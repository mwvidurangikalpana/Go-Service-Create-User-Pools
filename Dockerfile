FROM golang:1.19.0-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . . 

RUN go build -o /go-service-create-user-pool

EXPOSE 8088

CMD [ "/go-service-create-user-pool" ]