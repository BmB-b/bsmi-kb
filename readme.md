# bsmi-kb


bsmi kb 知识库

## 准备

```go 

/usr/local/go/bin/go build -ldflags "-X 'github.com/cnmade/bsmi-kb/pkg/version.BuildTag=$(git describe --tags --abbrev=0)' -X 'github.com/cnmade/bsmi-kb/pkg/version.BuildNum=$(date  '+%Y%m%d%H%M%S')'" --tags "json1 fts5 secure_delete" -v .
```

需要加上sqlite 扩展json1