FROM golang:1.19-alpine

RUN addgroup -S cinema && adduser -S cinema -G cinema
USER cinema

WORKDIR /src/app
COPY . .

RUN go get
RUN go install
