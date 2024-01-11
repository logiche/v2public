FROM logiche/alpine-tz:3.16
WORKDIR /run/app
COPY ./app ./
RUN chmod +x ./app
ENV V_CONFIG=''
EXPOSE 10082
USER root
CMD ["/run/app/app"]
