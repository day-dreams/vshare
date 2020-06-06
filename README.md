```text
           _
__   _____| |__   __ _ _ __ ___
\ \ / / __| '_ \ / _` | '__/ _ \
 \ V /\__ \ | | | (_| | | |  __/
  \_/ |___/_| |_|\__,_|_|  \___|

```
视频分享/同步观看

![CI](https://github.com/day-dreams/vshare.zhangnan.xyz/workflows/CI/badge.svg?branch=master)

- 配置视频，修改`build/config.json`
- 编译运行，`make docker-compose`
- 浏览器访问，`http://localhost:8080/?vid=demovid`
- Reference
    - HLS标准，https://www.rfc-editor.org/pdfrfc/rfc8216.txt.pdf
    - ffmpeg命令，MP5转m4u8，`ffmpeg -i input.mp4 -profile:v baseline -level 3.0 -s 640x360 -start_number 0 -hls_time 10 -hls_list_size 0 -f hls index.m3u8`
- todo
    - m3u8视频片段缓存
    - 聊天交互功能

运行示例

![demo.png](docs/demo.png)
