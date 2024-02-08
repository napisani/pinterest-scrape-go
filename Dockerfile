FROM golang:1.20.3
RUN mkdir /app
WORKDIR /app
COPY . .
RUN go mod download
CMD ["go","run","./cmd/main/main.go"]
