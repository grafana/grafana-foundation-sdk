# Installing

```shell
go get github.com/grafana/grafana-foundation-sdk/go@{{ if ne .Extra.ReleaseTag ""}}{{ .Extra.ReleaseTag }}{{ else }}{{ .Extra.ReleaseBranch }}{{ end }}
```
