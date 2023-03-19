# syntax=docker/dockerfile:1
FROM golang:1.19-alpine AS builder
# From which image it was created (and which later will be used for multi-staging)
WORKDIR /app
# will create a folder app in the container in which will work
COPY . .
# copy current folder files to the same folder
RUN go build -o cmd/asciiArtWeb main.go
# create binary file of the project inside container folder cmd with name groupieTracker
FROM alpine
# used for multi-staging (just conteiner with bash which will run binary file)
WORKDIR /app
COPY --from=builder /app .
# copy builder conteiner into working directory
LABEL "authors"="nzharylk"
LABEL version="1.0"
LABEL description="Ascii art web project with docker"
# Labels for metadata
EXPOSE 8080
# container will listen to port 8080 at runtime
CMD ["cmd/asciiArtWeb"]