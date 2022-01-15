FROM golang:1.17

WORKDIR /myapp
COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .

RUN go build -o ./bin/task ./cmd/main.go


EXPOSE 8080

ENTRYPOINT [ "./bin/task" ]