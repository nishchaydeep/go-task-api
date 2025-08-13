FROM golang:1.23-alpine
WORKDIR /go-task-api
COPY . .
RUN go build -o main .
CMD ["./main"]

