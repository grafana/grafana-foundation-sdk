# Installing

```shell
yarn add '@grafana/grafana-foundation-sdk@~{{ if ne .Extra.ReleaseTag ""}}{{ .Extra.ReleaseTag }}{{ else }}{{ .Extra.GrafanaVersion|registryToSemver }}-cog{{ .Extra.CogVersion }}.{{ .Extra.BuildTimestamp }}{{ end }}'
```
