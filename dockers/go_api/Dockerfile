FROM ubuntu:latest

RUN apt-get update

RUN mkdir /var/www/
ADD ./go_api /var/www/go_api

CMD ["/var/www/go_api"]