# permset
Simple Go-based permission setter for containers running as non root users

## Usage
When this binary is called with setuid root permissions it will attempt to recursively change ownership
of the directory specified at compile time to the user calling the binary. This is intended for docker containers
that you want to run as a non root user but may need to ensure their data folder is owned by the user in the container

## Example docker file
```
FROM golang:1.17-bullseye as permset
WORKDIR /src
RUN git clone https://github.com/jacobalberty/permset.git /src && \
    mkdir -p /out && \
    go build -ldflags "-X main.chownDir=/data" -o /out/permset

FROM debian:bullseye
COPY --from=permset /out/permset /usr/local/bin/permset
RUN chown 0.0 /usr/local/bin/permset && \
    chmod +s /usr/local/bin/permset  && \
    mkdir /data

USER nobody
```

if you build then run this container then `/data` will initially be owned by `root`. If you then call `/usr/local/data/permset` inside the container
then `/data` will be owned by `nobody`

## Security
chownDir MUST be an absolute path, not a relative one and any symbolic links in the path will be ignored.
