# Build Go Binary
FROM golang:1.26.11@sha256:6e59e17a20f13c102e49d82dffab66d41cffe6c266b225d82794328aea859ab5 AS build

WORKDIR /app
COPY ["go.mod", "go.sum", "./"]
RUN go mod download

COPY . .

ENV GOCACHE=/go/pkg/mod/
RUN --mount=type=cache,target="/go/pkg/mod/" go build -o dashboard-converter .

# Build Image
FROM golang:1.26.11-alpine@sha256:ecb67f987aa1a57794c91d19ade7ae0209f2f0f5496aaef9e6ae85ddf92d02ab
RUN apk add --no-cache ca-certificates

WORKDIR /root

COPY --from=build /app/dashboard-converter ./
ENTRYPOINT ["/root/dashboard-converter"]
