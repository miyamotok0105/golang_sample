#platform=linux/amd64の部分はm1 macの場合の処理
FROM --platform=linux/amd64 library/mysql:8.0.25

ENV MYSQL_DATABASE todo

COPY custom.cnf /etc/mysql/conf.d/

COPY sql /docker-entrypoint-initdb.d
