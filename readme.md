# bsmi-kb


bsmi kb 知识库

## 准备

```go 

/usr/local/go/bin/go build -ldflags "-X 'github.com/cnmade/bsmi-kb/pkg/version.BuildTag=$(git describe --tags --abbrev=0)' -X 'github.com/cnmade/bsmi-kb/pkg/version.BuildNum=$(date  '+%Y%m%d%H%M%S')'" --tags "json1 fts5 secure_delete" -v .
```

需要加上sqlite 扩展json1


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
