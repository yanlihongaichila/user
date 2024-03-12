#用于拉取镜像
FROM golang:1.20

#ENV GOPROXY https://goproxy.cn

RUN mkdir /app

COPY ./ /app

WORKDIR /app

RUN go mod tidy

RUN go build main.go


#FROM scratch
#
#COPY --from=build /app/main /main

CMD ["./main"]


#FROM centos:centos7.9.2009
#
#COPY ./config /config
#
#COPY ./main.exe .
#
#CMD ["./main.exe"]
