FROM golang:1.8.3 AS build-env
WORKDIR /go
COPY main.go .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o plugin

FROM docker:dind
COPY --from=build-env /go/plugin /bin/
ENTRYPOINT /bin/plugin