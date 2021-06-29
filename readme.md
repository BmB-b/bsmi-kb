# GoSense




Gosense 是一个用golang写的个人博客系统
  

## 准备

```go 
go get -u github.com/valyala/quicktemplate/qtc
ln -s $HOME/go/bin/qtc /usr/bin/qtc

/usr/local/go/bin/go build -ldflags "-X 'code.aliyun.com/netroby/gosense/pkg/version.BuildTag=$(git describe --tags --abbrev=0)' -X 'code.aliyun.com/netroby/gosense/pkg/version.BuildNum=$(date  '+%Y%m%d%H%M%S')'" --tags "json1 fts5 secure_delete" -v .
```

需要加上sqlite 扩展json1