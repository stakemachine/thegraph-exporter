FROM golang:1.19.5-alpine as builder

RUN apk update && apk add --no-cache git ca-certificates tzdata && update-ca-certificates

RUN wget https://github.com/upx/upx/releases/download/v4.0.1/upx-4.0.1-amd64_linux.tar.xz -O /tmp/upx-4.0.1-amd64_linux.tar.xz && \
    tar -xJOf /tmp/upx-4.0.1-amd64_linux.tar.xz upx-4.0.1-amd64_linux/upx > /usr/local/bin/upx && \
    chmod +x /usr/local/bin/upx

ENV USER=thegraph
ENV UID=10001

ARG git_commit
ENV GIT_COMMIT=$git_commit

ARG git_branch
ENV GIT_BRANCH=$git_branch

ARG git_tag
ENV GIT_TAG=$git_tag

RUN adduser --disabled-password --gecos "" --home "/nonexistent" --shell "/sbin/nologin" --no-create-home --uid "${UID}" "${USER}"

# Copy the local package files to the container's workspace.
COPY . /go/src/github.com/akme/thegraph-exporter
WORKDIR /go/src/github.com/akme/thegraph-exporter

RUN go get -d -v ./...

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -trimpath -ldflags="-X main.GitCommit=${GIT_COMMIT} -X main.GitBranch=${GIT_BRANCH} -X main.GitTag=${GIT_TAG} -w -s" -a -o /go/bin/thegraph-exporter .
RUN /usr/local/bin/upx --overlay=strip --best /go/bin/thegraph-exporter

FROM scratch
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
COPY --from=builder /go/bin/thegraph-exporter /go/bin/thegraph-exporter

USER thegraph:thegraph
WORKDIR /
EXPOSE 8080

ENTRYPOINT ["/go/bin/thegraph-exporter"]
