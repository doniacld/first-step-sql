FROM golang:1.14

WORKDIR /app
EXPOSE 8080
COPY ./bin/app.out /app/
RUN chmod +x /app/app.out
CMD ["/app/app.out"]
