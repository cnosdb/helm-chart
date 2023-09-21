FROM ubuntu:jammy
WORKDIR /app
COPY initconfig /app/initconfig
RUN chmod +x /app/initconfig
ENTRYPOINT ["/app/initconfig"]

