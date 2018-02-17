## Use Go runtime
## I just searched here: https://hub.docker.com/search/?isAutomated=0&isOfficial=0&page=1&pullCount=0&q=golang+slim&starCount=0
FROM golang:1.9
WORKDIR /app
ADD . /app
EXPOSE 31337
CMD ["/app/dundering"]
