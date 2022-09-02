FROM golang

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . ./

RUN go build -o /app

CMD ["/app"]

#FROM golang:alpine as builder
#WORKDIR /app

#COPY . .

#RUN go mod download
#RUN go build -o main .

#CMD ["./main"]