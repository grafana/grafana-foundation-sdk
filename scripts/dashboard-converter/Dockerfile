# Build Go Binary
FROM golang:1.26.10@sha256:672c944cfd07599b5f1bb179869270cdc94c6254e41db4d431259693cb286b82 AS build

WORKDIR /app
COPY ["go.mod", "go.sum", "./"]
RUN go mod download

COPY . .

ENV GOCACHE=/go/pkg/mod/
RUN --mount=type=cache,target="/go/pkg/mod/" go build -o dashboard-converter .

# Build Image
FROM golang:1.26.10-alpine@sha256:27f2d3b83cfc5803c8078f8e730af0f7c431bf067afaec8f60a36fb763187cba
RUN apk add --no-cache ca-certificates

WORKDIR /root

COPY --from=build /app/dashboard-converter ./
ENTRYPOINT ["/root/dashboard-converter"]
