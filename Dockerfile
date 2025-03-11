# get golang container
FROM golang:1.24.1 AS builder

# get args
ARG TibiaDataBuildBuilder=dockerfile
ARG TibiaDataBuildRelease=-
ARG TibiaDataBuildCommit=-

# create and set workingfolder
WORKDIR /go/src/

# copy go mod files and sourcecode
COPY go.mod go.sum ./
COPY src/ ./src/

# download go mods and compile the program
RUN go mod download && \
  CGO_ENABLED=0 GOOS=linux go build \
  -a -installsuffix cgo -ldflags="-w -s \
  -X 'main.TibiaDataBuildBuilder=${TibiaDataBuildBuilder}' \
  -X 'main.TibiaDataBuildRelease=${TibiaDataBuildRelease}' \
  -X 'main.TibiaDataBuildCommit=${TibiaDataBuildCommit}' \
  " -o app ./src/

# get alpine container
FROM alpine:3.21.3 AS app

# create workdir
WORKDIR /opt/app

# add packages, create nonroot user and group
RUN apk --no-cache add ca-certificates tzdata && \
  addgroup -S nonroot && \
  adduser -S nonroot -G nonroot && \
  chown -R nonroot:nonroot .

# set user to nonroot
USER nonroot:nonroot

# copy binary from builder
COPY --from=builder --chown=nonroot:nonroot --chmod=544 /go/src/app .

# expose port 8080
EXPOSE 8080

# set environment variables for Redis
ENV REDIS_ENABLED=false
ENV REDIS_ADDR=redis:6379
ENV REDIS_PASSWORD=
ENV REDIS_DB=0

# run application
CMD ["./app"]