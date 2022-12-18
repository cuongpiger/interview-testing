FROM postgres:11.5-alpine

ENV POSTGRES_PASSWORD=my-secret
ENV POSTGRES_USER=postgreuser
ENV POSTGRES_DB=cylo

WORKDIR /app

EXPOSE 5432

COPY init_cylo.sql /docker-entrypoint-initdb.d/
