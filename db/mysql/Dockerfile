FROM mysql:8.0

RUN apt-get update
RUN apt-get -y install locales-all

ENV LANG ja_JP.UTF-8
ENV LANGUAGE ja_JP:ja
ENV LC_ALL ja_JP.UTF-8

# COPY ../app/database/my.cnf /etc/mysql/conf.d/my.cnf

EXPOSE 3366
