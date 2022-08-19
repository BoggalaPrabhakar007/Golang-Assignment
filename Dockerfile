# Start from a Debian image with the latest version of Go installed and a workspace (GOPATH) configured at /go.
FROM golang:latest
#Maintainer details
MAINTAINER reddyprabhakar528@gmail.com

#Creates the directory called build
RUN mkdir /build
#Make the build directory as current directory
WORKDIR /build

#Copy the go project to the current directory
COPY . .

#Enable the go module
RUN export GO111MODULE=on
#Get the packages
RUN go get github.com/BoggalaPrabhakar007/golang-assignment/cmd
#Clone the project from the git repo every time
RUN cd /build && git clone https://github.com/BoggalaPrabhakar007/golang-assignment.git

#Build the go project
RUN cd /build/golang-assignment/cmd && go build

#Expose the port on 8080
EXPOSE 8080

#Run the go service
ENTRYPOINT ["/build/golang-assignment/cmd/cmd"]