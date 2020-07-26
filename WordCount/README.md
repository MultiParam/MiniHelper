这是一个超级超级超级简单的代码，希望不要被喷，说这样简单的代码都 push 上来。

## Motivation

有一天我突然被安排了一个任务，这个任务简单来说就是我需要将一份以 `缩写、全称、中文翻译` 为格式的 word 文档进行去重。文档部分内容如下所示。

```
第 1 章操作系统概述
OS			Operating System					操作系统
CPU			CentralProcessingUnit				中央处理器
I/O			Input/Output						输入/输出
API			Application Programming Interface	应用程序接口
FMS			Fortran Monitor System				FORTRAN监督系统
CTSS		CompatibleTime-SharingSystem		兼容分时系统
MULTICS		MULTiplexed Information and Computing System
UNICS		UNiplexed Information and Computing System
IEEE		Institute of Electrical and Electronics Engineers	电气和电子工程师学会
POSIX		Portable Operating System Interface	可移植操作系统接口
FSF			Free Software Foundation			自由软件基金
GLIBC		GNU C Library
GDB			GNU debugger
OS			Operating System					操作系统
......
```

去重主要是将那些重复的缩写使用进行去重，比如示例中，OS 这个缩写就有两处，那么需要将后者删除。这要是用人眼一点一点去看，不仅浪费时间而且很容易少查。

因此，我先去网上搜了一下，word 好像有自带的去重功能。但是每当我使用网上提供的 word 自带的查重功能时，整个程序就崩溃了，所以我也不知道这个到底好不好使。

走投无路，我就想到用代码写一个简单的吧。然后我就用 Python 代码撸了一个简单的，这段代码只对缩写进行统计所以超级简单快速。这样一来，我就知道哪些缩写是重复的了。然后，结合搜索方式将重复的内容删除。这边不直接删除重复的内容，一是因为 word 难以操作，如果要写会更加费时；二是因为这份文档是由别人汇总整理的，所以有些条目原本应该是一样的，但是由于格式或者拼写错写等原因可能会存在差错，从而简单的判断无法去除；三是我还得过一下文档，反馈一下缩写中存在的问题。



