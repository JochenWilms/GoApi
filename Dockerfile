# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from the latest golang base image
FROM golang:latest

# Add Maintainer Info
LABEL maintainer="Jochen Wilms <jochen@mariekerke.be>"

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

#set GOPATH to packages folder
ENV GOPATH=/app/packages
RUN mkdir $GOPATH/bin
ENV GOBIN=$GOPATH/bin
# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go get ./


# Build the Go app
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]