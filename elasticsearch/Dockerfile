FROM docker.elastic.co/elasticsearch/elasticsearch:8.2.0

COPY ./elasticsearch.yml /usr/share/elasticsearch/config/

WORKDIR /usr/share/elasticsearch

CMD ["bin/elasticsearch", "-f", "/usr/share/elasticsearch/config/elasticsearch.yml"]