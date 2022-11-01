# 东北大学每日健康上报

## 声明

本项旨在避免同学们漏报信息，使用请前检查配置文件信息，切勿错报。因错报引起的任何后果，本项目及其作者不承担任何责任。

请严格遵守防疫政策，切勿瞒报错报漏报。

## 下载

在 release 页面中下载对应版本的可执行文件，或 clone 本项目自行编译。

Linux 环境下赋予可执行权限

```shell
sudo chmod +x neuport
```

## 使用

### Usage

```
 neureport -fpath <path-without-filename> -fname <filename>
```

### 使用默认路径

路径在可执行文件同级目录，文件名为 config.json. 结构如下。

```tree
.
|-- neureport
|-- config.json
```

直接执行即可。

```shell
./neureport
```

### 指定配置文件路径

```shell
./neureport -fpath ./ -fname "config.json"
```

### 日志写入到文件

可重定向将日志写入到文件, 配合定时任务使用。

```
./neureport -fpath ./ -fname "config.json" > neureport.out 2>&1
```


## 配置文件格式

目前只支持非首次上报，切不更换地区。意思就是，如果你一直待在某个城市没变，直接用下方这个配置文件即可。

若位置变换，请手动上报一次，再使用这个配置文件。

### 非首次上报

|配置项|说明|
|---|---|
|StudentID|学号|
|Password|统一身份认证密码|
```json
{
    "StudentID":"",
    "Password":"",
    "info": {
        "jibenxinxi_shifoubenrenshangbao": "1", 
        "profile[xuegonghao]": "2171960",
        "profile[xingming]": "",
        "profile[suoshubanji]": "",
        "jiankangxinxi_muqianshentizhuangkuang": "正常",
        "xingchengxinxi_weizhishifouyoubianhua": "0",
        "cross_city": "无", 
        "qitashixiang_qitaxuyaoshuomingdeshixiang": ""
      }
}
```

### [WIP] RoadMap

所有参数，未完成。

```json
{
    "StudentID":"",
    "password":"",
    "info": {
        "jibenxinxi_shifoubenrenshangbao": "1", 
        "profile[xuegonghao]": "2171960",
        "profile[xingbie]": "1",
        "profile[zhengjianleixing]": "",
        "profile[lianxidianhua]": "",
        "profile[shenfenleixing]": "",
        "profile[jinjilianxirenxingming]": "",
        "profile[xingming]": "",
        "profile[chushengriqi]": "",
        "profile[zhengjianhaoma]": "",
        "profile[suoshudanwei]": "",
        "profile[suoshubanji]": "",
        "profile[jinjilianxirendianhua]": "",
        "jiankangxinxi_muqianshentizhuangkuang": "正常",
        "xingchengxinxi_weizhishifouyoubianhua": "1",
        "xingchengxinxi_guojia": "中国",
        "xingchengxinxi_shengfen": "浙江省",
        "xingchengxinxi_chengshi": "xxx市",
        "xingchengxinxi_quxian": "",
        "cross_city": "是/", // 无
        "qitashixiang_qitaxuyaoshuomingdeshixiang": "",
        "travels[0][chufadi]": "中国,北京市,市辖区",
        "travels[0][likaishijian]": "2022-11-24",
        "travels[0][mudidi]": "中国,天津市,市辖区",
        "travels[0][didashijian]": "2022-11-30",
        "travels[0][jiaotonggongju]": "火车",
        "travels[0][checi]": "3434"
      }
}
```
