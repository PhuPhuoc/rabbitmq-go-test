FROM golang:1.25-alpine AS build

RUN apk add --no-cache git

WORKDIR /src

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN go build cmd/main.go

FROM alpine AS runtime

RUN apk add --no-cache ca-certificates

COPY --from=build /src/main /app/main

CMD ["/app/main"]
