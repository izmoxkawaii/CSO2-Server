#镜像
FROM ubuntu:18.04
RUN apt-get update -y -q && apt-get upgrade -y -q 
#安装工具
RUN apt-get install gcc -y -q 
RUN gcc --version
RUN apt-get install wget -y -q 
#安装go
RUN wget https://studygolang.com/dl/golang/go1.15.4.linux-amd64.tar.gz
RUN tar xfz go1.15.4.linux-amd64.tar.gz -C /usr/local
RUN rm -f go1.15.4.linux-amd64.tar.gz
#设置环境变量
ENV PATH $PATH:/usr/local/go/bin
ENV GO111MODULE on
ENV GOPROXY=https://goproxy.cn,direct
RUN go version
#设置工作目录
WORKDIR $GOPATH/src/github.com/KouKouChan/CSO2-Server
#下载mod
COPY go.mod .
COPY go.sum .
RUN go mod download
#复制项目文件
COPY . .
#构建项目
RUN GOOS=linux GOARCH=amd64 go build -o CSO2-Server-docker .
#设置工作目录
WORKDIR $GOPATH/src/github.com/KouKouChan/
#切换可执行文件位置
RUN mv ./CSO2-Server/CSO2-Server-docker ./CSO2-Server-docker
#暴露端口
EXPOSE 1314
EXPOSE 1315
EXPOSE 30001
EXPOSE 30002
#最终运行docker的命令
#USER app-runner
ENTRYPOINT  ["./CSO2-Server-docker"]