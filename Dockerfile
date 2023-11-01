FROM golang:alpine3.18 as dev

RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go mod tidy
RUN go build -o villa-go .

FROM alpine:3.18
WORKDIR /root/
COPY --from=dev /app/villa-go .
COPY env.yaml /root/env.yaml
EXPOSE 8080
CMD ["./villa-go"]