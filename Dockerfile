FROM alpine
MAINTAINER joe_ke
RUN mkdir /config && mkdir /app
WORKDIR /app
COPY proxy /app

EXPOSE 8080

ENTRYPOINT ["./proxy"]