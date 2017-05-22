FROM golang:1.7

# create container directory
RUN mkdir /go/src/app

ENV AMQP_HOST "amqp://guest:guest@192.168.99.100:5672/"

# make container directory workdir
WORKDIR /go/src/app

# copy source code to container
COPY . /go/src/app

# install packages
RUN go get



# build application
RUN go build ./main.go

CMD ["./main"]