FROM mysql:8.0

ENV MYSQL_ROOT_PASSWORD=my-secret

RUN groupadd -r user-group && useradd -r -g user-group new-user

WORKDIR /app

EXPOSE 3306

COPY init.sql /app/

CMD ["mysqld", "--init-file=/app/init.sql"]