FROM golang:latest AS build-stage
ENV \
DBUSER=root \
DBPASS=12345678 \
DBADDR=127.0.0.1:3306
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./
# Check what CGO does
RUN CGO_ENABLED=0 GOOS=linux go build -o /dockerized-simple-api

FROM gcr.io/distroless/base-debian11 AS build-release-stage
WORKDIR /
COPY --from=build-stage /dockerized-simple-api /dockerized-simple-api
EXPOSE 8080
USER nonroot:nonroot
ENTRYPOINT [ "/dockerized-simple-api" ]