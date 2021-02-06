FROM golang as builder

WORKDIR /usr/src/peche
COPY . .

RUN CGO_ENABLED=0 go build -o /bin/peche


FROM scratch

WORKDIR /bin/peche

COPY --from=builder /bin/peche .

CMD ["./peche"]
