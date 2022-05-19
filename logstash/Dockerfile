FROM docker.elastic.co/logstash/logstash:8.2.0

COPY ./pipeline.conf /usr/share/logstash/

WORKDIR /usr/share/logstash

CMD ["bin/logstash", "-f", "pipeline.conf"]