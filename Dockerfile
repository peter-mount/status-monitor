# Dockerfile used to build the application

ARG arch=amd64
ARG goos=linux

# Build container containing our pre-pulled libraries
FROM golang:alpine AS build

# The golang alpine image is missing git so ensure we have additional tools
RUN apk add --no-cache \
      curl \
      git

# Ensure we have the libraries - docker will cache these between builds
RUN go get -v \
      github.com/coreos/bbolt/... \
      github.com/gorilla/mux \
      github.com/peter-mount/golib/kernel \
      github.com/peter-mount/golib/rest \
      gopkg.in/robfig/cron.v2 \
      gopkg.in/yaml.v2

# ============================================================
# source container contains the source as it exists within the
# repository.
FROM build AS source
WORKDIR /go/src/github.com/peter-mount/status-monitor
ADD . .

# ============================================================
FROM source AS compiler
ARG goos
ARG goarch
ARG goarm

RUN CGO_ENABLED=0 \
    GOOS=${goos} \
    GOARCH=${goarch} \
    GOARM=${goarm} \
    go build \
      -o /dest/statusmonitor \
      github.com/peter-mount/status-monitor/status/bin

      # ============================================================
# Finally build the final runtime container will all required files
FROM area51/scratch-base
COPY --from=compiler /dest/ /
ENTRYPOINT [ "/statusmonitor" ]
