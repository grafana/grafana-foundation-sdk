{{- $panelPackages := .Context.PackagesForVariant "panelcfg" -}}
{{- $dataqueryPackages := .Context.PackagesForVariant "dataquery" -}}
{{- range $pkg := $panelPackages }}
from ..models import {{ $pkg }}
{{- end }}
{{- range $pkg := $dataqueryPackages }}
from ..models import {{ $pkg }}
{{- end }}
from . import runtime as cogruntime


def register_default_plugins():
    # Panelcfg variants
{{- range $pkg := $panelPackages }}
    cogruntime.register_panelcfg_variant({{ $pkg }}.variant_config())
{{- end }}

    # Dataquery variants
{{- range $pkg := $dataqueryPackages }}
    cogruntime.register_dataquery_variant({{ $pkg }}.variant_config())
{{- end }}
