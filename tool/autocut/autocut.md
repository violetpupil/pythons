# [autocut](https://github.com/mli/autocut) 通过字幕来剪切视频

```bash
# 安装
pip install git+https://github.com/mli/autocut.git
# 转录某个视频生成 .srt 和 .md 结果
autocut -t short.mp4
# 修改 md 文件，去除无声片段
go run . -f short.md
# 剪切某个视频
autocut -c short.mp4 short.md short.srt
```
