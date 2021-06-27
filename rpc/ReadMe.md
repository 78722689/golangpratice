1. 编译
    go build -o example.exe .

2. 运行服务端
    2.1 TLS服务端
        example.exe rpc server tls
    2.2 None TLS 服务端
        example.exe rpc server
3. 运行客户端
    3.1 TLS客户端调用第n个RPC API， 如第2个
        example.exe rpc client tls 2
    3.2 None TLS客户端调用第n个RPC API， 如第3个
        example.exe rpc client ntls 3
     