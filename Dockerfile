FROM harbor.cloudminds.com/library/docker:bionic.CM-v1.4

ENV PROJECT=smartvoice-platform
ENV MODULE=applet-server
ENV LOGPATH=/app/runtime/logs

WORKDIR /app

EXPOSE 8300
EXPOSE 9300

RUN apt-get update && apt-get install -y curl

COPY internal/vad/libs  libs
COPY voiceText voiceText

COPY  ./bin/applet-server /app/applet-server
COPY run.sh /etc/services.d/applet/run