搭建节点手册
系统建议：
配置：4核(CPU) 内存8 GiB  硬盘2TB及以上
操作系统：Ubuntu



##### 1. 首先拿到压缩包，里面包含

```
geth进程

init.sh初始化脚本
start.sh启动脚本
stop.sh终止脚本


sha1sum geth
2106b9a0f1bc21ec6091b0f4b3a3b3cd0bb90af0
```



##### 2. 首先第一步创建账户

```
./geth account new --datadir /root/.kad/data

输入密码
自己记住密码,在文件中保存密码,挖矿程序启动的时候需要
在geth进程目录执行
echo "自己的密码" >  password.txt

打开start.sh，将start.sh中的挖矿账户改成自己的账号
```



#### 3.初始化

```
sh init.sh
```



#### 4.启动节点程序

一定要先把start.sh脚本里面的地址先换成自己的

```
sh start.sh
```



##### 5、  同步数据，xx状态时，表示数据同步完成

```
在终端中执行
./geth attach ipc:/root/.kad/data/geth.ipc

> eth.syncing
false

如果输出为 false，则表示 Geth 已完成同步
```



##### 6、  完成以上步骤   你已成为一个全节点。 



##### 7.成为候选节点

如果想参与验证节点获得出块和交易收益
节点地址必须持有10000KAD，并质押10000KAD  即可成为候选节点

- 质押流程

1. 先把服务器上生成的密钥文件下载下来导入到metamask(文件目录为/root/.kad/data//keystore目录下)，导入到小狐狸中

2.  再调用合约自己给自己进行质押



##### 8、  开放验证节点竞选

在验证节点投票开启期间，参与投票质押，排名前21名候选节点，成为验证节点进行出块。



