ARG GO_VERSION=1.24

# Start from base golang image
FROM golang:${GO_VERSION}-alpine AS builder

# Install the Certificate-Authority certificates for the app to be able to make
# calls to HTTPS endpoints.
RUN apk add --no-cache git
# Set the working directory outside $GOPATH to enable the support for modules.
WORKDIR /src

# Import the code from the context.
COPY ./ ./

# Build the Go app
RUN CGO_ENABLED=0 go build -mod=vendor -o /app .

######## Start a new stage from scratch #######
# Final stage: the running container.
FROM alpine:3.22.0 AS final

RUN echo 'nobody:x:65534:65534:nobody:/:' > /etc/passwd && \
    echo 'nobody:x:65534:' > /etc/group

RUN apk add --no-cache ca-certificates tzdata libwebp libwebp-tools

# Import the compiled executable from the first stage.
COPY --from=builder /app /app

# As we're going to run the executable as an unprivileged user, we can't bind
# to ports below 1024.
EXPOSE 8000

# Perform any further action as an unprivileged user.
USER nobody:nobody

# Run the compiled binary.
ENTRYPOINT ["/app"]
