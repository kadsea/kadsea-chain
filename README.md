首先拿到压缩包，里面包含

```
geth进程

init.sh初始化脚本
start.sh启动脚本
stop.sh终止脚本


sha1sum geth
0e19543822abe90d43017f074a1181b17042cd31
```



##### 首先第一步创建账户

```
./geth account new --datadir /root/.kad/data

输入密码
自己记住密码,在文件中保存密码,挖矿程序启动的时候需要
在geth进程目录执行
echo "自己的密码" >  password.txt

打开start.sh，将start.sh中的挖矿账户改成自己的账号
```



#### 初始化

```
sh init.sh
```



#### 启动节点程序

一定要先把start.sh脚本里面的地址先换成自己的

```
sh start.sh
```

