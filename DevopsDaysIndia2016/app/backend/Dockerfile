FROM centos:7

COPY server /opt/api/


EXPOSE 3000

ENV HOST 0.0.0.0
ENV DATA_FILE /opt/api/data/comments.json

RUN mkdir -p /opt/api/data/
RUN chmod a+w  /opt/api/data/

ENTRYPOINT ["/opt/api/server"]
