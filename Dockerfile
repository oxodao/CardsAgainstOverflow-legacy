FROM debian:bullseye-20200803-slim

MAINTAINER Nathan JANCZEWSKI <nathan@janczewski.fr>

COPY ./cardsagainstoverflow /cao

ENTRYPOINT [ "/cao" ]
