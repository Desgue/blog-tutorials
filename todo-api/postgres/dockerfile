FROM postgres

ENV POSTGRES_DB todos
ENV POSTGRES_USER develop
ENV POSTGRES_PASSWORD develop

COPY init.sql /docker-entrypoint-initdb.d/ 