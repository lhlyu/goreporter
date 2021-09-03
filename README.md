# GoReporter

本项目是克隆了[https://github.com/qax-os/goreporter](https://github.com/qax-os/goreporter)进行改造，
原项目不是go module, 直接安装会有些依赖问题，
我将其改造成go module项目

## 安装

### 要求

- go 1.16+
- [Graphviz](http://www.graphviz.org/download/)

`go get -u github.com/lhlyu/goreporter`

## 使用

跟原项目一样，[文档](https://github.com/qax-os/goreporter#run-it)


## 遇到的问题

`linters/simplecode/lint/lint.go` 内存在一个变量 `var gcImporter` ,
需要引用 `golang.org/x/tools/go/internal/gcimporter`, 因为go module无法直接使用第三方的internal内容,
所以使用了 `go:linkname`

```go
var gcImporter = gcimporter.Import
```

改造后

```go
//go:linkname gcImporter golang.org/x/tools/go/internal/gcimporter.Import
var gcImporter func(packages map[string]*types.Package, path, srcDir string) (*types.Package, error)
```
