FROM golang:1.23

WORKDIR /app
COPY go.* ./

RUN go mod download
COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /main

COPY challenge.txt ./

CMD ["/main"]
