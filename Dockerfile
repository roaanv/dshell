FROM --platform=$BUILDPLATFORM golang:1.22.5-alpine AS build
WORKDIR /src
COPY ./the-app/ .

ARG TARGETOS TARGETARCH
RUN GOOS=$TARGETOS GOARCH=$TARGETARCH  go build -o /out/the-app

# Deploy stage
FROM alpine:latest
RUN apk update &&  \
  apk upgrade && \
  apk add --no-cache \
    busybox-extras \
    bind-tools \
    curl \
    aws-cli

COPY --from=build /out/the-app /bin

EXPOSE 8080
CMD ["/bin/the-app"]
