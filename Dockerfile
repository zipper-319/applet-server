FROM harbor.cloudminds.com/library/alpine:3.CM-Beta-1.3

ENV PROJECT=smartvoice-platform
ENV MODULE=applet-server
ENV LOGPATH=/app/runtime/logs

WORKDIR /app

EXPOSE 8300
EXPOSE 9300

COPY internal/vad/libs /app/libs

COPY bin/applet-server /app/applet-server
COPY run.sh /etc/services.d/applet/run