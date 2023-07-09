# STAGE 1

FROM golang:alpine as build
# Update alpine and add git
RUN apk update && apk add --no-cache git
# Set WORKDIR
WORKDIR /src
# Copy all file to docker
COPY . .
# Install and tidying the dependency
RUN go mod tidy
# Build golang project to binary file
RUN go build -o binary

# STAGE 2
FROM alpine
# Set workdir in stage 2
WORKDIR /app
# COPY file from build stage to stage 2
COPY --from=build /src/binary .
#  Set entrypoint
ENTRYPOINT ["/app/binary"]