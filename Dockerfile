FROM golang:1.16 as builder

WORKDIR /home/app
COPY go.mod go.sum ./

RUN go mod download

COPY . .
RUN make build


FROM orvice/go-runtime:lite

COPY --from=builder /home/app/bin/deploy-cli .
