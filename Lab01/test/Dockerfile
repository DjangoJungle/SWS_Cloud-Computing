# 使用官方 Python 运行时作为父镜像
FROM python:3.8-slim

# 将工作目录设置为 /app
WORKDIR /app

# 将当前目录内容复制到位于 /app 的容器中
COPY . /app

# 使用 pip 安装任何需要的包
RUN pip install flask

# 让端口 8080 可用于外界访问
EXPOSE 8080

# 定义环境变量
ENV NAME World

# 在容器启动时运行 app.py
CMD ["python", "app.py"]
