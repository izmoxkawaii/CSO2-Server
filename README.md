## Counter-Strike Online 2 Server 

[![Build status](https://ci.appveyor.com/api/projects/status/a4pj1il9li5s08k5?svg=true)](https://ci.appveyor.com/project/KouKouChan/cso2-server)
[![](https://img.shields.io/badge/license-MIT-green)](./LICENSE)
[![](https://img.shields.io/badge/version-v0.6.0-blue)](https://github.com/KouKouChan/CSO2-Server/releases)

[English](./README.en.md) | [Turkish](./README.tr.md) | [CodePage](./CodePage.md)

*声明：Counter-Strike Online 2 归 NEXON 所有 ，本程序仅用于学习之用*

### 一、介绍

CSOL2 服务器 v0.6.0

默认需要浏览器打开**localhost:1314**网页注册账号，不需要输入邮箱和验证码

用于 **2017年国服 Counter-Strike Online 2** 和 **2018年韩服 Counter-Strike Online 2**  

目前客户端请使用L-Leite的启动器,**韩服端竞技模式有问题的请下载最新的汉化包**。

基于L-Leite的[cso2-master-server](https://github.com/L-Leite/cso2-master-server)。

如果大家有什么建议或问题，欢迎提出。

欢迎大家帮忙本地化文件，具体见下方第三部分。

**使用该服务端的服务器：**

CSO2-CN QQ-group 948725573

CSO2-TW https://discord.gg/CBqcPP5

CSO2-TUR/ENG/KOR : https://discord.gg/DC3n6nX89D

### 二、已知问题

    1.房主离开后，其余玩家会卡住直到炸出房间
    2.服务端目前还未适配2017国服端的部分数据包，所以结算界面数据显示存在错误

### 三、部分数据本地化方法

```
1.打开server.conf
2.修改LocaleFile选项，将其改为你的语种的文件名，比如 zh-tw.ini
3.进入 CSO2-Server\locales\ 目录
4.创建相同文件名文件，比如 zh-tw.ini
5.根据 zh-cn.ini 中的内容相应修改你的zh-tw.ini文件
```

### 四、客户端下载

  [2018年国服端](https://pan.baidu.com/s/1KD5ZIEsdTogUJ4WY_b-OjA) 提取码：52p7 

  [启动器(必下)](https://pan.baidu.com/s/1QGyRmjw24eJ5ycrFjorv_g)  提取码：amys

### 五、使用方法

1.需要有CSOL2客户端，同时使用第三方启动器

2.进入本项目的release页面下载最新版本的程序（ https://github.com/KouKouChan/CSO2-Server/releases ）

3 .(国服端请跳过该步骤) 建立bat文件，和游戏的bin目录同级，里面写入：

```shell
START ./bin/launcher.exe -masterip IP地址 -enablecustom -username 用户名 -password 密码
```

4.IP地址指的是你的服务端IP，如果是本地那么就填127.0.0.1（仅单人情况下），如果你要连接局域网别人的服务端那么就填别人的IP地址，如果你安装了汉化包，也可以再加上以下语句：

```shell
-lang schinese
```

5.先运行本项目的exe文件启动服务器，然后打开bat文件启动客户端即可**（国服端可能启动稍慢）**

- 密码最好使用数字+字母形式，不然可能闪退！
- 从网盘里面下载得到的start-cso2.bat文件需要修改里面的IP地址和用户名！
- 默认情况下，你可以使用浏览器打开 localhost:1314 ，点击右上角register来注册，且不需要输入邮箱和验证码。
- 如果你想开启邮箱注册，那么你需要一个邮箱账号并且申请到了密钥，将密钥填入配置文件，同时开启EnableMail。
- 如果你是和别人联机玩，那么即使你的电脑运行着服务端也**不能**在bat文件里填127.0.0.1，不然对方无法通过你的ip连接你。
- 浏览器里可以输入localhost:1314进入web页面，点击右上角register进行注册。
- 当你要架设局域网或外网时，请打开防火墙的端口。30001-TCP类型端口、30002-UDP类型端口
- 建立互联网服务器需要双方玩家都能内网穿透，实测局域网能够连接，互联网无法房间内加入主机，可能需要架设虚拟局域网。

### 六、Console使用方法

CSO2-Server自带管理员功能，可通过命令行参数打开console功能管理服务器，前提需要服务器已经在运行。

1.运行服务器。

2.使用local-console.bat连接本地服务器或者使用如下命令连接服务器：

```
CSO2-Server.exe -console -ip YOURIP -port YOURPORT -username GMNAME -password PASSWORD
```

默认参数如下:

```
Usage of CSO2-Server.exe:
  -ip string
        主机名，默认为localhost (default "localhost")
  -password string
        密码，默认为cso2server123 (default "cso2server123")
  -port string
        端口号，默认为1315 (default "1315")
  -username string
        账号，默认为admin (default "admin")
```

3.连接成功后可以使用命令管理服务器了，你可以踢出玩家，或者给予玩家物品等。

### 七、自定义文件方法

1.下载CSOL2解包工具，[点击这里下载](https://pan.baidu.com/s/14q1SoIdHwp1casMWG2OS-w) 提取码：41bs

2.解压后，打开工具，点击左上角File选项，点击Open folder，选中csol2的data文件夹即可

3.解压你需要的文件，并且将解压后的文件按你的想法进行修改

4.将文件放入csol2目录的custom文件夹下，打开游戏

### 八、Docker下使用方法

1.首先你需要拥有Docker,请下载并安装Docker,同时配置好Docker,比如Docker源

2.输入以下命令拉取最新版的服务端:

```shell
docker pull koukouchan/cso2-server:latest
```

3.运行服务端

```shell
docker run -p 30001:30001 -p 30002:30002 -p 1314:1314 -p 1315:1315 koukouchan/cso2-server:latest
```

4.接下来打开客户端，连接服务器

5.建议关闭docker时将容器数据导出，否则将丢失玩家数据！

### 九、编译环境

*Go 1.15.8*

### 十、图片

![Image](./photos/main.png)

![Image](./photos/intro.png)

![Image](./photos/channel.png)

![Image](./photos/ingame.jpg)

![Image](./photos/result.jpg)
