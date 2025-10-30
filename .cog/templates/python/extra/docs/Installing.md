# Installing

```shell
python3 -m pip install 'grafana_foundation_sdk=={{ if ne .Extra.ReleaseTag ""}}{{ .Extra.ReleaseTag }}{{ else if eq .Extra.GrafanaVersion "next" }}{{ .Extra.BuildTimestamp }}{{ else }}{{ .Extra.BuildTimestamp }}!{{ .Extra.GrafanaVersion|registryToSemver }}{{ end }}'
```
