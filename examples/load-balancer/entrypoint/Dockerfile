# Must be run from project root (misty/) to get misty-utils and client packages
# docker build -t misty-entrypoint -f examples/load-balancer/entrypoint/Dockerfile .

# Stage 1: build
FROM golang:1.16@sha256:d388153691a825844ebb3586dd04d1c60a2215522cc445701424205dffc8a83e as build

WORKDIR /src

# Load in the module packages first to cache them
COPY  ./examples/load-balancer/entrypoint/go.mod ./examples/load-balancer/entrypoint/go.sum  /src/

COPY ./clients/misty-go/ /clients

RUN go mod download

# Load in the rest of the packages
COPY ./examples/load-balancer/entrypoint/main.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o entrypoint


# Stage 2: Runs broker
FROM scratch

COPY --from=build /src/entrypoint /

EXPOSE 2395

CMD ["/entrypoint"]
