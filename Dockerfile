FROM golang:1.19-alpine AS build

WORKDIR /app
COPY . .
COPY cmd/api/main.go ./
COPY go.mod go.sum ./

RUN go mod vendor
RUN CGO_ENABLED=0 go build -o /server

FROM scratch
COPY --from=build /server /server
COPY --from=build /app/.env ./

ENTRYPOINT ["/server"]