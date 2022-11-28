## Motivation
做批量处理时，需要针对一类文件进行递归搜索。

## Design
利用 os 模块即可。搜索就是简单地递归还有加一些 map 和 lambda

## Use cases
```python
from files import walk_4_files
include = ('mp4', 'avi')
videos = walk_4_files(video_path, include) 
```
