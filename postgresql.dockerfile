FROM postgres:14

COPY /db/migrations/*.up.sql /docker-entrypoint-initdb.d/

ENV POSTGRES_USER root
ENV POSTGRES_PASSWORD root
ENV POSTGRES_DB simple_bank

EXPOSE 5432