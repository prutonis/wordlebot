# The base go-image
FROM golang:1.20.4-buster
 
# Create a directory for the app
RUN mkdir /app
 
# Copy all files needed for Wordlebot to the app directory
COPY configs/ /app/configs
COPY pkg/ /app/pkg
COPY main.go /app
COPY go.mod /app
COPY go.sum /app
 
# Set working directory
WORKDIR /app
 
# Run command as described:
# go build will build an executable file named tbot in the current directory
RUN go build -o tbot . 
 
# Run the tbot executable
CMD [ "/app/tbot" ]