package plugins

{{- $panelPackages := .Context.PackagesForVariant "panelcfg" -}}
{{- $dataqueryPackages := .Context.PackagesForVariant "dataquery" -}}

{{- if or (ne (len $panelPackages) 0) (ne (len $dataqueryPackages) 0) }}
import (
	cog "{{ .Data.PackageRoot }}/cog"
	{{- range $pkg := $panelPackages }}
	{{ $pkg }} "{{ $.Data.PackageRoot }}/{{ $pkg }}"
	{{- end }}
	{{- range $pkg := $dataqueryPackages }}
	{{ $pkg }} "{{ $.Data.PackageRoot }}/{{ $pkg }}"
	{{- end }}
)
{{- end }}

func RegisterDefaultPlugins() {
	{{- if or (ne (len $panelPackages) 0) (ne (len $dataqueryPackages) 0) }}
	runtime := cog.NewRuntime()
	{{- end }}

    // Panelcfg variants
{{- range $pkg := $panelPackages }}
	runtime.RegisterPanelcfgVariant({{ $pkg | formatPackageName }}.VariantConfig())
{{- end }}

    // Dataquery variants
{{- range $pkg := $dataqueryPackages }}
	runtime.RegisterDataqueryVariant({{ $pkg | formatPackageName }}.VariantConfig())
{{- end }}
}
