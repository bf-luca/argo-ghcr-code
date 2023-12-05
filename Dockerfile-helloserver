# Specifies a parent image
FROM golang:1.19.2-bullseye
 
# Creates an app directory to hold your app’s source code
WORKDIR /app

COPY ./helloserver /app
 
# Builds your app with optional configuration
RUN go mod init github.com/bf-luca/argo-ghcr-code/helloserver && go mod tidy
RUN go build .
 
# Tells Docker which network port your container listens on
EXPOSE 8080
 
# Specifies the executable command that runs when the container starts
ENTRYPOINT [ "./helloserver" ]
