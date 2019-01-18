FROM ubuntu:latest

RUN apt-get update

RUN mkdir /var/services/
ADD ./delta_update /var/services/delta_update

CMD ["/var/services/delta_update"]