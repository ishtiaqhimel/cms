# DO NOT use in production!
# Dockerfile for local development

FROM golang:1.24

# install the watcher
RUN go install github.com/githubnemo/CompileDaemon@latest
RUN apt update && apt install -y --no-install-recommends webp

WORKDIR /app
COPY ./ /app

ENTRYPOINT ["CompileDaemon", "--build=go build -buildvcs=false", "-log-prefix=false", "-exclude-dir=.git"]
