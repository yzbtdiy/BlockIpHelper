# 配置文件说明

配置文件位于./data/config.yaml, 执行程序时会检测配置文件是否存在, 不存在会自动生成

```yaml
# 需要筛选分类的地址列表文件, 一行一个地址
target_file: ./target.txt     
# 白名单列表, 可以是单个地址, 也可以是CIDR, 用 / 分割掩码
white_list: ./data/whitelist.txt
# 处理完成后导出文件
export_file:
    # 白名单匹配到的地址
    in_whitelist: ./白名单IP.txt
    # 查询后归属地是国内的地址列表
    in_china: ./国内攻击IP.txt
    # 查询后归属地是国外的地址列表
    out_china: ./国外攻击IP.txt
# 导出设备黑名单模板, 目前仅支持k01和明御防火墙
template:
    # 模板名称
    - name: k01
    # 是否启用, 为 true 则生成对应设备导入文件
      enable: true
    # 生成csv文件的保存路径
      export_path: ./k01Block.csv
    - name: myfw
      enable: true
      export_path: ./myFwBlock.csv
# ip2region相关配置
ip2region:
    # 纯真IP解压得到的 qqwry.txt 文件
    cz_txt: ./data/qqwry.txt
    # 纯真IP库转化为 ip_merge.txt 文件的保存位置
    merge_file: ./data/ip_merge.txt
    # 生成的xdb文件保存位置
    xdb_file: ./data/ip2region.xdb
    # 根据以下关键词判断是否为国内地址
    cn_keys:
        - 中国
        - 北京
        - 广东
        - 山东
        - 江苏
        - 河南
        - 上海
        - 河北
        - 浙江
        - 香港
        - 陕西
        - 湖南
        - 重庆
        - 福建
        - 天津
        - 云南
        - 四川
        - 广西
        - 安徽
        - 海南
        - 江西
        - 湖北
        - 山西
        - 辽宁
        - 台湾
        - 黑龙江
        - 内蒙古
        - 澳门
        - 贵州
        - 甘肃
        - 青海
        - 新疆
        - 西藏
        - 吉林
        - 宁夏
```