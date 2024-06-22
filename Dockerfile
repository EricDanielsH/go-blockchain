# Downloand base Golang image
FROM golang:1.22

# Choose the working directory in the container
WORKDIR /app

# Copy packege files to the workspace
COPY go.mod go.sum /app

# Download the package dependencies
RUN go mod download

# Copy program source code files to workspace
COPY . . 

# Build the go app
RUN go build -o program .

# Command to run the executable (in an array of strings)
CMD ["./program"]

# Build image:
# docker build -t ericdanielsh/go-blockchain:1.0 .

# Run the image:
# docker run --rm -it ericdanielsh/go-blockchain:1.0 ./program printchain
# docker run --rm -it ericdanielsh/go-blockchain:1.0 ./program addblock -data "Send 1BTC to someone"



