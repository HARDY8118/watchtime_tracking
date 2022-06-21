FROM docker.elastic.co/logstash/logstash:8.2.0

COPY ./logstash.conf /usr/share/logstash/

WORKDIR /usr/share/logstash

CMD ["bin/logstash", "-f", "logstash.conf"]