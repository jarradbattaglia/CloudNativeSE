FROM golang:1.21 as build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /voterapi

FROM alpine:latest AS run-stage

WORKDIR /

COPY --from=build-stage /voterapi /voterapi

EXPOSE 8080

ENV REDIS_URL=host.docker.internal:6379

CMD ["/voterapi"]