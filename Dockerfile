FROM golang:latest

RUN go version
ENV GOPATH=/

COPY ./ ./

COPY --from=migrate/migrate ./ ./

RUN go mod download
RUN go build -o todo ./cmd/app/main.go

CMD ["./todo"]


