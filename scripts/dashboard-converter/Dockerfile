# Build Go Binary
FROM golang:1.26.10@sha256:cfa27959837144bf4649802ba38e59daffd0dd2a7994aa83daaf113285b0f551 AS build

WORKDIR /app
COPY ["go.mod", "go.sum", "./"]
RUN go mod download

COPY . .

ENV GOCACHE=/go/pkg/mod/
RUN --mount=type=cache,target="/go/pkg/mod/" go build -o dashboard-converter .

# Build Image
FROM golang:1.26.10-alpine@sha256:87c65a688975c0b1d4585d204874c45ca78354670e3a5f11368c5a492303d696
RUN apk add --no-cache ca-certificates

WORKDIR /root

COPY --from=build /app/dashboard-converter ./
ENTRYPOINT ["/root/dashboard-converter"]
