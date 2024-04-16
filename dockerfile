FROM golang:latest as build

WORKDIR /app

COPY go.mod go.sum ./

RUN --mount=type=cache,target=go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    go mod download

COPY . .

RUN go build \
    -ldflags="-linkmode external -extldflags -static" \
    -tags netgo \
    -o lru-cache


FROM scratch

COPY --from=build /etc/ssl/certs /etc/ssl/certs

COPY --from=build /app/lru-cache lru-cache

EXPOSE 8080

CMD [ "/lru-cache" ]