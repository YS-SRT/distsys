FROM golang:1.22
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct

WORKDIR /home/workspace
ADD . /home/workspace
RUN rm /home/workspace/application.toml && mv /home/workspace/application-docker.toml /home/workspace/application.toml

# RUN go mod init distsys
RUN go mod tidy
RUN go build main.go

EXPOSE 8090
ENTRYPOINT ["./main"]
