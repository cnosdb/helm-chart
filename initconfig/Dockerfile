FROM --platform=$TARGETPLATFORM ubuntu:jammy
WORKDIR /app
ARG TARGETARCH
COPY initconfig_${TARGETARCH} /app/initconfig
RUN chmod +x /app/initconfig
ENTRYPOINT ["/app/initconfig"]

