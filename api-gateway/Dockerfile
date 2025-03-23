FROM golang:1.23-alpine AS buildphase

#set the working directory from here all commands execute inside this dir
WORKDIR /app

#copy go.mod and go.sum then download the dependencies
COPY go.mod go.sum ./
RUN go mod download


# COPY all files to container
COPY . .

RUN go mod tidy

#compile/build the application
RUN go build -o main .

#now our final stage 
FROM alpine:latest

WORKDIR /root/

#copy the compiled go binary from buildphase to here 
COPY --from=buildphase /app/main .

#expose the application port
EXPOSE 8086

#run the go application
CMD ["./main"]

