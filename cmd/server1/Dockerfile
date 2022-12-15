FROM alpine
ARG TARGETARCH
WORKDIR /app
COPY server-linux-${TARGETARCH} /app/server
ENTRYPOINT [ "/app/server" ]
