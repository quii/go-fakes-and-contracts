FROM golang:1.18-alpine as builder

WORKDIR /app

ARG bin_to_build

COPY go.mod ./

RUN go mod download

COPY . .

RUN go build -o svr cmd/${bin_to_build}/main.go

FROM alpine:3.16.2
COPY --from=builder /app/svr .
CMD [ "./svr" ]