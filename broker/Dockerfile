# Stage 1: Test and build (Runs unit tests, creats binary)
FROM golang:1.16@sha256:d388153691a825844ebb3586dd04d1c60a2215522cc445701424205dffc8a83e as testing

WORKDIR /src

# Load in the module packages first to cache them
COPY  ./broker/go.mod ./broker/go.sum  /src/

COPY ./misty-utils /misty-utils

RUN go mod download

# Load in the rest of the packages
COPY ./broker/repository /src/repository
COPY ./broker/service /src/service

# Load in the configuration file
COPY ./broker/.misty.yaml /src/.misty.yaml

# Load in the entrypoint
COPY ./broker/main.go /src/main.go

RUN go test -v ./repository

RUN CGO_ENABLED=0 GOOS=linux go build -o misty-broker


# Stage 2: Runs broker
FROM scratch

COPY --from=testing /src/.misty.yaml .
COPY --from=testing /src/misty-broker .

EXPOSE 1315

CMD ["/misty-broker"]
