{{- if not .Data.Debug -}}
module {{ .Data.PackageRoot }}

go 1.21

{{- end -}}