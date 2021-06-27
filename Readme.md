1. 编译
    go build -o example.exe .

2. 运行指定类型的程序
    example.exe [type] [-arguments]

    type:
        kafka
        context
        channel
        tip
        sync
        gc
        rpc
        https

    arguments根据不同的type，arguments会有所不同