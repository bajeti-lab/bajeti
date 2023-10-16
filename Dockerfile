FROM golang:1.21.2

# add maintainer
LABEL maintainer="Kagirim + sarah"

# install git
RUN apt-get update && apt-get install -y git

# set working directory
RUN mkdir /app
WORKDIR /app

# copy the source from the current directory to the working directory
COPY . .
COPY .env .

# download dependencies
RUN go get -d -v ./...

# build the application
RUN go install -v ./...

# expose port 8080 to the outside world
EXPOSE 8080

# command to run the application
CMD ["/build"]