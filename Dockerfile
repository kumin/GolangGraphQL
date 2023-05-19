FROM golang:1.19 AS compiler

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN make static -j$(nproc)

FROM alpine:3.18.0

WORKDIR /

COPY --from=compiler /app/bin/* /bin/

ENTRYPOINT ["tiny", "--"]
