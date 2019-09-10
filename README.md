## 将你的掘金小册装进kindle

1. 下载掘金小册保存为MarkDown的格式
2. 下载掘金小册保存为mobi格式

### 注意：转成mobi用的是[这个项目](https://github.com/Ares002/mobi)，只支持一级目录和二级目录，导致目前mobi格式比较乱，里面的图片没有处理，如果你有好用的转mobi的工具或者代码，一定提issue告诉我，或者直接contribute.

### 使用方法

- 从掘金小册的站点获取client_id,uid,token等信息，填入./config/config.json中
- go run main.go
- 对于没有Go环境的小伙伴，在tool文件夹中已经编译好了Linux&Mac版本的可执行文件，（在跟main.go同一目录下，因为要读取config）./juejinxiaoceToolForLinux或者./juejinxiaoceToolForMac即可
- 等待你的小册出炉🍺

### 小白操作步骤

1. 登录掘金网站，打开你的一本小册，如下图：

![](https://ws1.sinaimg.cn/large/006tKfTcgy1g0d2ooxxe4j326g0h8q8p.jpg)

2. windows按F12,Mac使用快捷键conmand+shift+i,进入开发者工具，如下图：

![](https://ws3.sinaimg.cn/large/006tKfTcgy1g0d2rgef4lj31eg0q846g.jpg)

3. 选择XHR（意思是过滤Ajax的异步请求），刷新一下网页，找到get?uid=xxx，上图中红色圈出部分。点击get?uid=xxx部分，可以看到该API的请求与返回等信息，其实就是我们需要的uid,id,token等，如下图所示：

![](https://ws3.sinaimg.cn/large/006tKfTcgy1g0d3xyiy8gj30k612kgrv.jpg)

4. 可以点击response查看该API的返回信息，对于golang来说也是构造结构体的关键，具体可查看代码；
5. 通过这个API我们可以获取sectionId的一个数组，每一个sectionId就对应文章中的一节，在第filter中搜索`getListSection`，该API就是获取指定sectionId的内容；

![](https://ws3.sinaimg.cn/large/006tKfTcgy1g0d430l8vzj30pg0cqmz7.jpg)

上面所用的URL我都放到了config/config.json里面，可以直接使用。

以上是基本操作，搞定之后，就执行`go run main.go`等待你的小册出炉🍺。

![](https://ws1.sinaimg.cn/large/006tKfTcgy1g0d48s2jflj313q06udgy.jpg)

可以看到MarkDown格式和mobi格式的书籍都已经下载。
