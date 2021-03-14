FROM alpine:latest as builder
RUN apk --update add ca-certificates

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY ./poolscore /
COPY ./config.yaml /
COPY ./google-key.json /
ENV GOOGLE_APPLICATION_CREDENTIALS=./google-key.json 
ENTRYPOINT ["/poolscore"]
