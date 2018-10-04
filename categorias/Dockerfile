# golang image where workspace (GOPATH) configured at /go.
FROM golang:1.6

# Install dependencies
RUN go get gopkg.in/mgo.v2
RUN go get github.com/gorilla/mux

# copy the local package files to the container workspace
ADD . /go/src/foro/categorias/

# Setting up working directory
WORKDIR /go/src/foro/categorias/

# Build the categorias command inside the container.
RUN go install foro/categorias/

# Run the categorias microservice when the container starts.
ENTRYPOINT /go/bin/categorias

# Service listens on port 8080.
EXPOSE 8080
