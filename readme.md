# bsmi-kb


bsmi kb 知识库,  

demo 预览地址： [https://kb.bsmi.info/#](https://kb.bsmi.info/#)

文档地址：[https://kb.bsmi.info/#view/3](https://kb.bsmi.info/#view/3)


Buy me a cup of coffee for $3

[![ko-fi](https://ko-fi.com/img/githubbutton_sm.svg)](https://ko-fi.com/M4M54KKIF)



如果你想安装，使用，或者是参加开发，可以读一下文档

## 起源

这个项目诞生的原始驱动，是因为对“语雀” 知识库的不满。一直不喜欢语雀的目录，发布流程。所以在创建bsmi-kb的时候，痛下苦功，把繁琐的操作简化了。

让任何人，可以自由的编辑，体验，组织，修改自己的知识库，帮助建立自己的知识库体系。


## 开发准备

如果你想开发和体验，请阅读这段文档说明。

您需要准备好golang的开发环境，建议安装最新的golang 1.16.5或者以上版本。

推荐在x64架构的cpu的电脑上运行，理论上windows, Linux 和mac都可以用

那么编译命令如下， 是在Linux上面运行的。

```go 
go generate

/usr/local/go/bin/go build  --tags "json1 fts5 secure_delete" -v .
```

需要加上sqlite 扩展json1

这样开发环境就可以用了，可以开始体验了


## demo 效果图

![Snipaste_2021-07-01_17-37-14](https://user-images.githubusercontent.com/278153/124104252-a7ca4c80-da94-11eb-97fd-5784a81b5ce3.png)


## 许可证

    Copyright (C) 2000-2021  cnmade

    This program is free software: you can redistribute it and/or modify
    it under the terms of the GNU Affero General Public License as published
    by the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.

    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU Affero General Public License for more details.

    You should have received a copy of the GNU Affero General Public License
    along with this program.  If not, see <https://www.gnu.org/licenses/>.
