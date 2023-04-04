FROM golang:1.18-alpine

WORKDIR /app

COPY . .
RUN go mod download

RUN apk update && apk upgrade && apk add --no-cache make

# Build
RUN make build

ENV HOST 0.0.0.0
EXPOSE 80
CMD ["./bin/server"]