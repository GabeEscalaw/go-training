FROM alpine:1.16-alpine
WORKDIR /prog19
COPY main.go /app
COPY go.mod /app
RUN go build -o server
EXPOSE 8080
CMD ["./server"]