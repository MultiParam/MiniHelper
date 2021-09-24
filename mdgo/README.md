## Motivation

相信大家都有使用过 markdown 做笔记的经历，markdown 对程序员等人群来说超级友好。在实际使用的时候，我们往往也会在 markdown 里引用一些图片，对于这些图片的引用可以使用本地路径的方式，也可以使用网络 URL 的方式。为了方便笔记的共享或者为了发博客，则需要采用网络 URL 的方式来引用这些图片。

为了让这些图片可以在任何地方被引用和图片的统一管理，一般使用 markdown 的人都会有一个图床，或者是基于阿里云搭建图床，或者是基于腾讯云搭建图床，或者是使用其他开源的图床。有了图床之后，将会借助一些工具（比如 PicGo）方便快速地将图片上传到图床中。更进一步的，在使用 typora 做笔记的时候，假如结合使用 PicGo，那么这个时候不需要手动键入引用图片的语句，而是直接以粘贴的方式在 typora 中粘贴图片，那么就可以将图片自动上传到图床中并自动键入对该图片的引用。

然而，在使用的过程中，遇到了以下几个问题：
1. PicGo 经常性上传图片失败；
2. PicGo 和 typora 结合使用的时候，无论打开哪个文件，这些文件里的图片都将被上传到图床的同一个目录中，不方便分类；
3. 对于之前没有使用 PicGo+typora 做的笔记，picgo 不方便将它们使用的本地图片或者网络上的其他图片统一上传到自己的图床中（里头有个插件是可以用的，但是我是真的没用成功过，安装都没安装成功过）；
4. markdown 有时候引用的图片地址比较杂，想要将这些图片全都统一到自己的图床中，然后引用它们；
5. 希望将 markdown 引用的图片进行备份，统一保存到本地目录中。

## Design
针对以上这些问题，准备设计和实现一个小工具 mdgo。mdgo 可以一键就将 markdown 文件引用的图片上传到图床中，并且更新文档中对这些图片的引用。也可以一条命令就将 markdown 引用的图片统一保存到本地目录中，可以根据选项是否更改引用。

## Use cases
```bash
# 根据 config.yaml 中的配置信息将 demo.md 中引用的图片统一上传到图床中，并且更新文档中对图片的引用为图床中的 URL
mdgo push demo.md demo2.md -c config.yaml

# 将 deme.md 中引用的图片统一 pull 到 ./images 目录
mdgo pull demo.md demo2.md --path ./images

# 下载并且更新引用
mdgo pull demo.md --path ./images --update

## ToDo
# 考虑，可以上传/下载整个目录中的 markdown 文件所使用的图片
# 在 markdown 引用了本地图片的情况下，后续准备考虑一键清理本地未引用的链接；
# 同理，在全都 push 到图床之后，可以考虑一键清理图床中未引用的图片；
```

```yaml
# config.yaml
pushOptions: 
  endPoint: ""
  bucketName: ""
  accessID: ""
  accessKey: ""
  path: ""
  domain: "https://blog.dawnguo.cn"
```