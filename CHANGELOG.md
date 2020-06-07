
<a name="v0.2.0"></a>
## [v0.2.0](https://github.com/day-dreams/vshare/compare/v0.0.1...v0.2.0) (2020-05-06)

### Debug

* rm debug info

### Dev

* 1080画质；更新readme
* set default addr
* rm location url
* update index.html
* add player.go
* add cors
* 从配置文件获取vid/hls相关参数

### Docs

* 更新readme

### Feat

* 引入前端播放器
* 实现hls m3u8 segment接口
* 实现hls m3u8 playlist接口
* 使用exec.Command调用ffmpeg，获取视频基本信息
* 启用json配置文件
* 支持docker运行

### Fix

* HLS/VOD模式，需要在playlist末尾指定 '#EXT-X-ENDLIST'
* 引入build文件


<a name="v0.0.1"></a>
## v0.0.1 (2020-03-30)

### Dev

* 调整聊天框的height，适应手机
* 自动滚动聊天框到最新的消息
* 调整ui
* 调整ui
* 更新配置和vid获取方法
* 调整聊天ui和基本配置
* 同步功能调试完毕，暂时不要seek
* 调整UI和接口
* 增加room enter/read/write 接口，同步聊天和播放状态
* add axios api example; change vid/rid structure
* add index.html
* add make daemon
* add index.html
* add /api/video/list
* add framework

