FROM golang:1.19-alpine

RUN addgroup -S cinema && adduser -S cinema -G cinema

RUN mkdir -p /var/cinema && chown cinema /var/cinema
VOLUME /var/cinema
RUN chown cinema:cinema /var/cinema -R

WORKDIR /src/app
COPY --chown=cinema:cinema . .

USER cinema

RUN go get
RUN go install
