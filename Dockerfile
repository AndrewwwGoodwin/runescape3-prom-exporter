FROM golang:1.22 AS build-stage
LABEL authors="Rivista"

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /rs3-prom-export

# Deploy the application binary into a lean image
FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY --from=build-stage /rs3-prom-export /rs3-prom-export

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/rs3-prom-export"]
