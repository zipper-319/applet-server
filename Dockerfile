FROM golang:1.19 as build-stage

WORKDIR /app

COPY . .

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
ENV GOARM=6

RUN make build

FROM harbor.cloudminds.com/library/alpine:3.CM-Beta-1.3

ENV PROJECT=smartvoice-platform
ENV MODULE=applet-server
ENV LOGPATH=/app/runtime/logs

WORKDIR /app

EXPOSE 8300
EXPOSE 9300

COPY internal/vad/libs /app/libs

COPY --from=build-stage  /app/applet-server .
COPY run.sh /etc/services.d/applet/run