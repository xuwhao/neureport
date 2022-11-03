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

目前只支持非首次上报，且不更换地区。意思就是，如果你一直待在某个城市没变，直接用下方这个配置文件即可。

若位置变换，请手动上报一次，再使用这个配置文件。带 '*' 号的需要修改为自己的信息。

配置文件中的填报地址是杭州的，需要改成你所在的位置。等我返校了上传一版沈阳的。或者可以自己抓包。

前往 https://e-report.neu.edu.cn/mobile/notes/create，按正常填报方式填好信息。然后 F12 再断网点提交。 将 https://e-report.neu.edu.cn/api/notes 这个地址的请求 Json 直接复制到配置文件的 info 项中即可。

### 配置项说明

|配置项|说明|
|---|---|
|StudentID|学号|
|Password|统一身份认证密码|
|ProvinceCode|省份区划码|
|CityCode|城市区划码|
|credits|暂时不清楚，填默认的3，可以抓包看看自己的|
|bmap_position 相关|上报时的地址信息，应该是根据 IP 来的，猜测调用的百度地图的 API|

### Json 文件

```json
{
  "StudentID": "21*****",
  "password": "******",
  "info": {
    "_token": "",
    "jibenxinxi_shifoubenrenshangbao": "1",
    "profile": {
      "xuegonghao": "21*****",
      "xingming": "***",
      "suoshubanji": "计硕****"
    },
    "jiankangxinxi_muqianshentizhuangkuang": "正常",
    "xingchengxinxi_weizhishifouyoubianhua": "0",
    "cross_city": "无",
    "qitashixiang_qitaxuyaoshuomingdeshixiang": "",
    "credits": "3",
    "bmap_position": "{\"accuracy\":118,\"altitude\":null,\"altitudeAccuracy\":null,\"heading\":null,\"latitude\":30.18732056999,\"longitude\":120.27302702919,\"speed\":null,\"timestamp\":null,\"point\":{\"lng\":120.27302702919,\"lat\":303.18732056999,\"of\":\"inner\"},\"address\":{\"city\":\"杭州市\",\"city_code\":0,\"district\":\"滨江区\",\"province\":\"浙江省\",\"street\":\"滨文路\",\"street_number\":\"528号\"}}",
    "bmap_position_latitude": "20.18732056999",
    "bmap_position_longitude": "220.27302702919",
    "bmap_position_address": "浙江省,杭州市",
    "bmap_position_status": "0",
    "ProvinceCode": "330000",
    "CityCode": "330101",
    "travels": []
  }
}
```
