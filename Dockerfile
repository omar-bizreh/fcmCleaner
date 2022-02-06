FROM golang

WORKDIR /app
COPY fcmCleaner /app/

RUN ["chmod", "+x", "fcmCleaner"]

# Install fcmCleaner
#RUN go install fcmCleaner
ENTRYPOINT ["./fcmCleaner"]
# Document that the service listens on port 8080.
EXPOSE 8081
