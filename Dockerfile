FROM harbor.cloudminds.com/library/debian9:slim.CM-v1.4

ENV PROJECT=smartvoice-platform
ENV MODULE=applet-server
ENV LOGPATH=/app/runtime/logs

WORKDIR /app

EXPOSE 8300
EXPOSE 9300

RUN apt-get update && apt-get install -y curl

COPY internal/vad/libs /app/libs

COPY  ./bin/applet-server /app/applet-server
COPY run.sh /etc/services.d/applet/run