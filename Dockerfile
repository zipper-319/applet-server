FROM golang:1.19 as build-stage

WORKDIR /app

COPY . .

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
ENV GOARM=6

RUN go env -w GOPROXY=https://goproxy.io,direct
RUN go mod download && go mod verify && go build -ldflags="-s -w" -o ./bin/appletBackend ./cmd/$(PROJECT)/...

FROM harbor.cloudminds.com/library/alpine:3.CM-Beta-1.3

ENV PROJECT=speech-tts
ENV MODULE=applet-server
ENV LOGPATH=/app/runtime/logs

WORKDIR /app

EXPOSE 8300
EXPOSE 9300

COPY --from=build-stage /app/bin/appletBackend .
COPY run.sh /etc/services.d/applet/run