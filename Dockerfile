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

# sep hot reload
RUN go get github.com/githubnemo/CompileDaemon
RUN go get -v golang.org/x/tools/gopls

ENTRYPOINT CompileDaemon --build="go build -a -installsuffix cgo -o main." --command=./main
