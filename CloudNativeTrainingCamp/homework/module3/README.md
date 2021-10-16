# job

构建本地镜像。

`docker build -t caiyiluo/httpserver:v1.0 .`

编写 Dockerfile 将练习 2.2 编写的 httpserver 容器化（请思考有哪些最佳实践可以引入到 Dockerfile 中来）。

```text
Dockerfile in httpserver/Dockerfile
Best practices references:
- Use multi-stage builds
- Use Alpine image 
- Use .dockerignore files
```

将镜像推送至 Docker 官方镜像仓库。

`docker push caiyiluo/httpserver:v1.0`

通过 Docker 命令本地启动 httpserver。

`docker run -d --name caiyiluohttpserver -p 80:80 caiyiluo/httpserver:v1.0`

通过 nsenter 进入容器查看 IP 配置。

`nsenter -t $(docker inspect -f "{{.State.Pid}}" caiyiluohttpserver1) -n ip a`
