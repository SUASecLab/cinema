FROM golang:1.23-alpine AS golang-builder

RUN addgroup -S cinema && adduser -S cinema -G cinema

WORKDIR /src/app
COPY --chown=cinema:cinema . .

RUN go get
RUN go build

FROM scratch
COPY --from=golang-builder /src/app/cinema /cinema
COPY --from=golang-builder /etc/passwd /etc/passwd

USER cinema
CMD [ "/cinema" ]
