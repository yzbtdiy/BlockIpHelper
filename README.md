# BlockIpHelper

* 自用工具, 用于地址分类和生成csv文件, 无其他用途

* 地址库使用纯真IP社区版, 地址纠错请反馈给[纯真IP](https://update.cz88.net/review)


### 功能

* 可以导入白名单, 生成模板时排除白名单地址, 避免误封

* 使用本地ip2region的xdb数据库, 毫秒级查询IP归属地, 并导出分类文件

* 生成某些设备黑名单导入模板, 默认国内地址封禁30天, 国外地址永久封禁

### 目前支持的设备模板:

* k01

* 明御防火墙

### 使用说明

#### 下载二进制文件, 或修改代码自行编译

> 首次运行会生成配置文件, 保存在 `./data/config.yaml`, [配置说明](https://github.com/yzbtdiy/BlockIpHelper/blob/main/docs/config.md)

> 命令行支持 `-gen` 和 `-imp` 两个选项, 处理的文件均需要放到 `./data` 目录下

* `-imp white` 可以导入 whitelist.txt 内白名单地址

* `-gen merge` 可以将纯真IP解压得到的 qqwry.txt 转化为 ip_merge.txt

* `-gen xdb` 可以将上述生成的 ip_merge.txt 文件转化为 ip2region.xdb 文件

```cmd
.\BlockIpHelper.exe -h

flag needs an argument: -imp
Usage of D:\mycode\BlockIpHelper\BlockIpHelper.exe:
  -gen string
        merge   纯真IP解压文件(./data/qqwry.txt)生成的xdb源文件(./data/ip_merge.txt)
        xdb     xdb源文件(./data/ip_merge.txt)生成ip2region的xdb文件(./data/ip2region.xdb)
  -imp string
        white   导入白名单(./data/whitelist.txt)
```

#### 导入白名单, 支持单个地址和CIDR

```cmd
 .\BlockIpHelper.exe -imp white

2023/08/17 17:33:30 开始导入白名单 ...
2023/08/17 17:33:30 121.0.0.0/8 已添加到数据库白名单地址表中
2023/08/17 17:33:30 209.86.118.253 已添加到数据库白名单地址表中
```

#### 运行后自动分类白名单, 国内攻击, 国外攻击, 生成对应 txt 文件

使用本地 ip2region.xdb 数据库判断归属地, 归属地离线数据来源于纯真IP社区版(2023.08.23)

target.txt 为 10000 条测试数据, 1秒内极速分类

```cmd
 .\BlockIpHelper.exe

2023/08/17 17:37:20 #############################################################
2023/08/17 17:37:20 目标地址列表存在白名单地址, 保存到 ./白名单IP.txt
2023/08/17 17:37:20 #############################################################
2023/08/17 17:37:20 白名单IP , 国内攻击IP, 国外攻击IP 分类完成, 请查看对应txt文件
2023/08/17 17:37:20 #############################################################
2023/08/17 17:37:20 开始生成k01黑名单导入文件, 保存到 ./k01Block.csv
2023/08/17 17:37:20 #############################################################
2023/08/17 17:37:20 开始生成明御防火墙黑名单导入文件, 保存到 ./myFwBlock.csv
2023/08/17 17:37:20 #############################################################
2023/08/17 17:37:20 本次处理10000个地址, 耗时330.3895ms
```

### 使用的第三方库

* [Gorm](https://github.com/go-gorm/gorm)

* [Sqlite](https://github.com/glebarez/sqlite)

* [ip2region](https://github.com/lionsoul2014/ip2region)

* [纯真社区版IP库](https://cz88.net/geo-public)