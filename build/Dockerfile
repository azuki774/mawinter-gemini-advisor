# syntax=docker/dockerfile:1

FROM golang:1.24.2 AS builder
LABEL description="mawinter-gemini-advisor"

WORKDIR /app

RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,source=go.mod,target=go.mod \
    --mount=type=bind,target=. \
    go mod download -x

RUN --mount=type=bind,target=. \
    go build -a -tags "netgo" -installsuffix netgo  -ldflags="-s -w -extldflags \"-static\" \
    -X main.version=$(git describe --tag --abbrev=0) \
    -X main.revision=$(git rev-list -1 HEAD) \
    -X main.build=$(git describe --tags)" \
    -o /bin/ ./...

FROM gcr.io/distroless/base-debian12
ENV TZ=Asia/Tokyo
COPY --from=builder /bin/mawinter-gemini-advisor /usr/local/bin/mawinter-gemini-advisor
ENTRYPOINT ["/usr/local/bin/mawinter-gemini-advisor", "start"]
