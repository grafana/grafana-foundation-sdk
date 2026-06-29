# Build Go Binary
FROM golang:1.26.3@sha256:2d6c80227255c3112a4d08e67ba98e58efd3846daf15d9d7d4c389565d881b1a AS build

WORKDIR /app
COPY ["go.mod", "go.sum", "./"]
RUN go mod download

COPY . .

ENV GOCACHE=/go/pkg/mod/
RUN --mount=type=cache,target="/go/pkg/mod/" go build -o dashboard-converter .

# Build Image
FROM golang:1.26.4-alpine@sha256:3ad57304ad93bbec8548a0437ad9e06a455660655d9af011d58b993f6f615648
RUN apk add --no-cache ca-certificates

WORKDIR /root

COPY --from=build /app/dashboard-converter ./
ENTRYPOINT ["/root/dashboard-converter"]
