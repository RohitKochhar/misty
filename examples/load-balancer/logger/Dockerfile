# Must be run from project root (misty/) to get misty-utils and client packages
# docker build -t misty-entrypoint -f examples/load-balancer/logger/Dockerfile .

# Stage 1: build
FROM golang:1.16@sha256:d388153691a825844ebb3586dd04d1c60a2215522cc445701424205dffc8a83e as build

WORKDIR /src

# Load in the module packages first to cache them
COPY  ./examples/load-balancer/logger/go.mod ./examples/load-balancer/logger/go.sum /src/

COPY ./misty-utils /misty-utils

COPY ./clients/misty-go/ /clients

RUN go mod download

# Load in the rest of the packages
COPY ./examples/load-balancer/logger/main.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o logger


# Stage 2: Runs broker
FROM scratch

COPY --from=build /src/logger /

EXPOSE 2396

CMD ["/logger"]
