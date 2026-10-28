# Build Go Binary
FROM golang:1.26.10@sha256:670bb5aa5349ae03b8516764718486f4f16399134bdc9acd0c46a23f043ab05c AS build

WORKDIR /app
COPY ["go.mod", "go.sum", "./"]
RUN go mod download

COPY . .

ENV GOCACHE=/go/pkg/mod/
RUN --mount=type=cache,target="/go/pkg/mod/" go build -o dashboard-converter .

# Build Image
FROM golang:1.26.10-alpine@sha256:a2012b4a1d1f9f7e9f44e0d39b805a901e856b1dd80b4e6ee4338db547050308
RUN apk add --no-cache ca-certificates

WORKDIR /root

COPY --from=build /app/dashboard-converter ./
ENTRYPOINT ["/root/dashboard-converter"]
