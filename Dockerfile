FROM golang:1.25-bookworm

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy && go mod download

COPY . .

RUN go build -o /usr/local/bin/hypertube-serve ./main.go

CMD ["hypertube-serve"]
