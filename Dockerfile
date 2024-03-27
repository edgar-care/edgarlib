FROM mongo:latest

RUN rm -rf /data/db/*

CMD ["mongod", "--bind_ip_all"]
