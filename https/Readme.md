1. 运行环境
    OS: Windows 10
    Golang: go version go1.15.1 windows/amd6
            go1.9.2 windows/386
    
    PS: 本程序由于一些原因目前只在Win10+go1.15.1, Win7+go1.9.2上验证过

2. run testing
    2.1 如果go >= 1.15，先设置环境变量export GODEBUG="x509ignoreCN=0"
    2.2 cd $GOPATH/httpsserver
    2.3 go test -v .

3. 本程序支持单独运行服务端和客户端在不同的进程里, 以Windows10为例。
    3.1 编译可执行文件
      cd $GOPATH/httpsserver
      go build -o example.exe .
    3.2 运行服务端
      example.exe https -strart=server
    3.3 运行客户端发送数据
      example.exe https -start=client -send="11,22,33,44,55,aa,bb,cc"

    PS: 只支持服务端和客户端运行在同一个host

4. 生成证书(可选)
    PS: 项目中已经生成了默认的证书
    工具: https://github.com/square/certstrap/releases
    使用说明: https://github.com/square/certstrap
    4.1生成 root CA
        certstrap init --common-name "rootca"
    4.2 生成client和server的CA
        certstrap request-cert --common-name  "localhost" -domain "localhost"
        certstrap request-cert --common-name  "client" -domain "localhost"

      Note: GO1.15之后的版本使用tls证书必须要包含SAN(Subject alternative name)字段，该字段是SSL标准x509中定义的一个扩展字段，用于扩展此证书支持的域名，使得一个证书可以支持多个不同域名的解析。certstrap使用 "--domain"参数指定需要绑定的域名, 也可以添加更多参数比如--ip等绑定更多的参数. 如果不指定domain参数则会报错error "x509: certificate relies on legacy Common Name field, use SANs or temporarily enable Common Name matching with GODEBUG=x509ignoreCN=0"。也可以使用WA来规避这个错误，即增加环境变量"export GODEBUG=x509ignoreCN=0"。
      如果使用openssl工具，也可以使用如下配置: 参考https://blog.csdn.net/weixin_40280629/article/details/113563351?utm_medium=distribute.pc_relevant.none-task-blog-2%7Edefault%7EBlogCommendFromMachineLearnPai2%7Edefault-1.control&depth_1-utm_source=distribute.pc_relevant.none-task-blog-2%7Edefault%7EBlogCommendFromMachineLearnPai2%7Edefault-1.control
        [req_ext]
        subjectAltName = @alt_names
        [alt_names]
        DNS.1   = www.eline.com

    4.3 签发client和server的证书
        certstrap sign localhost --CA rootca  
        certstrap sign client --CA rootca   
