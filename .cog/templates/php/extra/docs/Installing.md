# Installing

```shell
composer require "grafana/foundation-sdk:{{ if ne .Extra.ReleaseTag ""}}{{ .Extra.ReleaseTag }}{{ else }}dev-{{ .Extra.ReleaseBranch }}{{ end }}"
```
