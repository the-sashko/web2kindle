FROM ubuntu:22.04

MAINTAINER inbox@sashko.me

ENV TZ=Europe/Kiev

RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

RUN apt-get update -y

RUN apt-get upgrade -y

RUN apt-get dist-upgrade -y

RUN apt-get autoremove

RUN apt-get install -y git bash golang calibre

RUN mkdir /storage

RUN git clone https://github.com/the-sashko/web2kindle.git /storage/web2kindle

RUN cd /storage/web2kindle && git checkout -f master && git pull origin master

RUN mkdir /storage/web2kindle/bin

RUN chmod -R 775 /storage/web2kindle/bin

RUN cd /storage/web2kindle/src && go build -o /storage/web2kindle/bin/app

RUN chmod -x /storage/web2kindle/bin/app

RUN chmod 775 /storage/web2kindle/bin/app

RUN rm -rf /storage/web2kindle/config

EXPOSE 80

CMD ["./storage/web2kindle/scripts/deamon.sh"]
