FROM golang:1.14

WORKDIR /app
EXPOSE 8080
COPY ./bin/app.out /app/
CMD ["/app/app.out"]