## Motivation
很多大视频不太方便用于网络训练，算力要求很大，费力不讨好。所以就想把视频切小一些。

## Design
基于 FFmpeg。基本的 shell 操作就解决了，加以对指定文件的搜索辅助。

## Use cases
```bash
python cut.py
```

需要注意的是，本工具没有拓展参数化，所以实在代码里头修改，这个可以成为一个 todo。修改内容如下
```python
video_path  = './'
output_path = '/mnt/.../1080-cut-10s'
delta_x    = 10 # s
cut(video_path, output_path, ('mp4'), delta_x)
```