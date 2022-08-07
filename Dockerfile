FROM golang:1.15 as builder

WORKDIR /usr/src/peche
COPY . .

RUN go mod tidy

RUN CGO_ENABLED=0 go build -o /bin/peche


FROM alpine

WORKDIR /bin/peche

COPY --from=builder /bin/peche .

CMD ["./peche"]
