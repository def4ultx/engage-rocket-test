# build stage
FROM golang:1.14 AS build-env

WORKDIR /src

ADD go.mod /src
ADD go.sum /src
RUN go mod download

COPY . /src

RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o server

# final stage
FROM alpine
WORKDIR /app
COPY --from=build-env /src/server /app/server

EXPOSE 8080
ENTRYPOINT /app/server