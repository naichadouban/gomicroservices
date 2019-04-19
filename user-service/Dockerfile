FROM alpine:latest

RUN mkdir -p /app
WORKDIR /app
ADD user-service /app/user-service

CMD [ "./user-service" ]