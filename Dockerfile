FROM golang:1.22.3 AS build
WORKDIR /app
COPY . .
RUN go mod tidy
RUN GOOS=linux GOARCH=amd64 go build -o server cmd/main.go

FROM gcr.io/distroless/base
COPY --from=build /app/server /server
EXPOSE 8080
ENTRYPOINT ["/server"]