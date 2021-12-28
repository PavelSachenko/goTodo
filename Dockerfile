FROM golang:latest

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN go mod download
RUN go build -o todo ./cmd/app/main.go

CMD ["./todo"]


