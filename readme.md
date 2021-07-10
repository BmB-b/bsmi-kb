# bsmi-kb


bsmi kb 知识库,  

demo 预览地址： [https://kb.bsmi.info/#](https://kb.bsmi.info/#)

文档地址：[https://kb.bsmi.info/#view/3](https://kb.bsmi.info/#view/3)

如果你想安装，使用，或者是参加开发，可以读一下文档

## 开发准备

如果你想开发和体验，请阅读这段文档说明。

您需要准备好golang的开发环境，建议安装最新的golang 1.16.5或者以上版本。

推荐在x64架构的cpu的电脑上运行，理论上windows, Linux 和mac都可以用

那么编译命令如下， 是在Linux上面运行的。

```go 

/usr/local/go/bin/go build -ldflags "-X 'github.com/cnmade/bsmi-kb/pkg/version.BuildTag=$(git describe --tags --abbrev=0)' -X 'github.com/cnmade/bsmi-kb/pkg/version.BuildNum=$(date  '+%Y%m%d%H%M%S')'" --tags "json1 fts5 secure_delete" -v .
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
