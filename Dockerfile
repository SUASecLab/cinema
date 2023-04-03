FROM golang:1.20-alpine

RUN addgroup -S cinema && adduser -S cinema -G cinema

WORKDIR /src/app
COPY --chown=cinema:cinema . .

USER cinema

RUN go get
RUN go install
