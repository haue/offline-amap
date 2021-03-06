# 基本思路

- 浏览器设置代理到代理程序，代理程序中dump到响应内容，即把访问过的请求缓存起来
- 修改官方的js文件使原来到外网的请求重定向至://本地服务器/offline-amap/xxxx。这个比较麻烦，除了jsloader外的js都是混淆压缩过的，修改的内容挺多，里面的坑也很多，目前修改后的文件已基本能用。
- 本地服务器设置反向代理使/offline-amap走缓存的内容。

# TODO

- js脚本自动移动地图到指定范围和zoom level

# go-forward-proxy

- 配置文件示例（.env）

  ```
  PORT=8888
  DOWNLOADDIR=download
  ```

# go-reverse-proxy

- 配置文件示例（.env）

  ```
  PORT=37706
  DOWNLOADDIR=C:/code/offline-amap/go-forward-proxy/download
  ```

## amap js

### js文件夹下的文件来源

```
1.js:
	https://webapi.amap.com/maps?callback=___onAPILoaded&v=2.0&key=6bb374bf3108300e7b2bcf0ed16a65a3&plugin=
2.js
	https://restapi.amap.com/v3/log/init?platform=JS&s=rsv3&logversion=2.0&product=JsInit&key=6bb374bf3108300e7b2bcf0ed16a65a3&t=1614219853649&sdkversion=2.0&appname=http%253A%252F%252Flocalhost%253A3000%252F&csid=B3B540CF-A323-4D5E-BCB1-77412116D65B&resolution=1920*1080&mob=0&vt=1&dpr=1&scale=1&detect=false&callback=jsonp_925333_1614219853649_
3.js
	https://webapi.amap.com/ui/1.1/main.js
4.js
	http://webapi.amap.com/ui/1.1/ui/misc/PathSimplifier.js?v=1.1.2&mt=ui&key=
5.js
	http://webapi.amap.com/ui/1.1/ui/misc/MarkerList.js?v=1.1.2&mt=ui&key=
6.js
	http://webapi.amap.com/count?type=UIInit&k=&v=1.1.2
7.js
	http://webapi.amap.com/ui/1.1/plug/ext/jquery-1.12.4.min.js?v=1.1.2
8.js
	https://webapi.amap.com/maps/plugin?v=2.0&cls=AMap.Scale&key=6bb374bf3108300e7b2bcf0ed16a65a3
9.js
	http://webapi.amap.com/ui/1.1/ui/geo/DistrictExplorer.js?v=1.1.2&mt=ui&key=
10.js
	https://webapi.amap.com/style?name=macaron&key=6bb374bf3108300e7b2bcf0ed16a65a3&callback=jsonp_823235_1614219854076_
11.js
	https://vdata.amap.com/style/2.0
12.js
	http://webapi.amap.com/ui/1.1/ui/geo/DistrictExplorer/assets/d_v2/an_410000.json?v=1.1.2&key=
13.js
	http://glyph.amap.com/static/commonWordsInfo.v2.1.json
	
```



### 相关文件内容替换

- base

  ```
    <script type="text/javascript">
      var AMapUIProtocol = 'http:';
    </script>
  ```

  

- index.js

  ```
  webapi.amap.com
  -->
  "+location.host+"/offline-amap"+"
  ```

  ```
  https
  -->
  http
  ```

- 1.js

  ```
  Lu.server
  -->
  "http://"+location.host+"/offline-amap"
  ```
obsoleted
  ```
  n+"/plugin?
  -->
  "http://"+location.host+"/offline-amap"+"/plugin?
  ```
  ```
  s+"://webapi.amap.com/style?name="
  -->
  "http://"+location.host+"/offline-amap/style?name="
  ```
  
  ```
  https://vdata.amap.com/style/2.0
  -->
  http://"+location.host+"/offline-amap/style/2.0
  ```
  
  ```
  https://
  -->
  http://
  ```
  
  ```
  poi名称
  搜索function makeFetchRequest
  函数内第一行加
  const z1=self.location.origin+'/offline-amap'
  g=g.replace(/^.*?vdata\d*?.amap.com/g,z1)
  
  //glyph.amap.com/sdf的替换
  ```
  
  
  
  
  
- 3.js

  ```
  "productWebRoot": "//webapi.amap.com/ui"
  -->
  "productWebRoot": "http://"+location.host+"/offline-amap"
  ```

  ```
  "baseUrl": "//webapi.amap.com/ui/1.1/"
  -->
  "baseUrl": "http://"+location.host+"/offline-amap/ui/1.1/"
  ```

  ```
  ns.docProtocol + "//webapi.amap.com/
  -->
  "http://"+location.host+"/offline-amap/
  ```

  