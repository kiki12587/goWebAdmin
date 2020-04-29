FROM golang:alpine

# 为我们的镜像设置必要的环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# 移动到工作目录：/home/www/newgoadmin 这个目录 是你项目代码 放在linux上 哪个位置的代码 这是我的目录 你们得修改成自己的
WORKDIR /home/www/newgoadmin

# 将代码复制到容器中
COPY . .

# 将我们的代码编译成二进制可执行文件app
RUN go build -o app .

# 移动到用于存放生成的二进制文件的 /dist 目录
WORKDIR /dist

# 将二进制文件从 /home/www/newgoadmin 目录复制到这里
RUN cp /home/www/newgoadmin/app .

# 声明服务端口
EXPOSE 9000

# 启动容器时运行的命令
CMD ["/dist/app"]