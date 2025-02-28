FROM golang:alpine

WORKDIR /app

COPY main.go /app/main.go

RUN go mod init pathfinder-app
RUN go mod tidy
RUN go get github.com/gofiber/fiber/v2
RUN go get github.com/gofiber/fiber/v2/middleware/cors
RUN go build -o /app/main main.go

EXPOSE 4000

CMD ["/app/main"]