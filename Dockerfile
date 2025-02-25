FROM alpine:3.21.3 AS app

WORKDIR /opt/app

RUN apk --no-cache add ca-certificates tzdata && \
  addgroup -S nonroot && \
  adduser -S nonroot -G nonroot && \
  chown -R nonroot:nonroot .

USER nonroot:nonroot

COPY --from=builder --chown=nonroot:nonroot --chmod=544 /go/src/app .

EXPOSE 8080

# set environment variables for Redis
ENV REDIS_ENABLED=false
ENV REDIS_URL=redis://localhost:6379
ENV REDIS_ADDR=redis:6379
ENV REDIS_PASSWORD=
ENV REDIS_DB=0

# run application
CMD ["./app"]