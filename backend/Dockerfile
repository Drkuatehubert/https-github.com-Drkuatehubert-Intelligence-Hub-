FROM golang:1.23-alpine AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o /app/hub ./cmd/pentagi

FROM alpine:latest
WORKDIR /app
COPY --from=build /app/hub /app/hub
COPY --from=build /app/migrations /app/migrations
COPY --from=build /app/pkg/templates/prompts /app/pkg/templates/prompts
EXPOSE 8080
CMD ["/app/hub"]
