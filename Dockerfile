FROM golang:1.20-bullseye
WORKDIR /app
COPY go.mod ./
COPY *.go ./
RUN go build -o /go-docker-demo
EXPOSE 8080
CMD [ "/go-docker-demo" ]