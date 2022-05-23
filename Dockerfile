# Build the Go API
FROM golang:latest AS go_builder
ADD . /app
WORKDIR /app
RUN go mod download
WORKDIR /app/backend
RUN ls
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /main ./cmd/main.go
# Build the React application
FROM node:alpine AS node_builder
WORKDIR /app
COPY --from=go_builder /app/frontend ./
RUN npm install
RUN npm run build
# Final stage build, this will be the container
# that we will deploy to production
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=go_builder /main ./
COPY --from=node_builder /app/build ./web
RUN chmod +x ./main
EXPOSE 8080
CMD ./main