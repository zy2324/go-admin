FROM golang
 
MAINTAINER "zhaoyi"
 
WORKDIR /home/go-admin
 
ADD . /home/go-admin

ENV GOPROXY=https://goproxy.cn,https://goproxy.io,direct 

RUN go build main.go
 
EXPOSE 5000
 
ENTRYPOINT ["./main"]
