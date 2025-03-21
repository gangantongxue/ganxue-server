FROM golang:1.24.1-alpine

# 创建非特权用户
RUN adduser -D -g '' sandboxuser
USER sandboxuser

WORKDIR /home/sandboxuser

# 复制用户代码到工作目录
# COPY user_code.go .

# 编译用户代码
# RUN go build -o user_code user_code.go

# 设置入口点
#ENTRYPOINT ["go","run","user_code.go"]
