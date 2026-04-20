<?php

namespace {{ .Data.NamespaceRoot }}\Cog;

final class DataqueryConfig
{
    public readonly string $identifier;

    /**
     * @var callable(array<string, mixed>): Dataquery
     */
    public $fromArray;

    /**
     * @var (callable(Dataquery): string)|null
     */
    public $convert;
    
    {{- $hasBeta := objectExists "dashboardv2beta1" "DataQueryKind" -}}
    {{- $hasV2 := objectExists "dashboardv2" "DataQueryKind" -}}
    {{- if or $hasBeta $hasV2 }}
    /**
     * @var (callable(
     {{- if and $hasBeta $hasV2 }}
        \Grafana\Foundation\Dashboardv2beta1\DataQueryKind|
        \Grafana\Foundation\Dashboardv2\DataQueryKind
     {{- else if $hasBeta }}
        \Grafana\Foundation\Dashboardv2beta1\DataQueryKind
     {{- else }}
        \Grafana\Foundation\Dashboardv2\DataQueryKind
     {{- end }}
     ): string)|null
     */
    public $convertv2;
    {{- end }}

    /**
     * @param callable(array<string, mixed>): Dataquery $fromArray
     * @param (callable(Dataquery): string)|null $convert
    {{- if or $hasBeta $hasV2 }}
     * @param (callable(
    {{- if and $hasBeta $hasV2 }}
        \Grafana\Foundation\Dashboardv2beta1\DataQueryKind|
        \Grafana\Foundation\Dashboardv2\DataQueryKind
    {{- else if $hasBeta }}
        \Grafana\Foundation\Dashboardv2beta1\DataQueryKind
    {{- else }}
        \Grafana\Foundation\Dashboardv2\DataQueryKind
    {{- end }}
    ): string)|null $convertv2
    {{- end }}
     */
    public function __construct(string $identifier, callable $fromArray, ?callable $convert = null{{ if or (objectExists "dashboardv2beta1" "DataQueryKind") (objectExists "dashboardv2" "DataQueryKind") }}, ?callable $convertv2 = null{{ end }})
    {
        $this->identifier = $identifier;
        $this->fromArray = $fromArray;
        $this->convert = $convert;
        {{- if or (objectExists "dashboardv2beta1" "DataQueryKind") (objectExists "dashboardv2" "DataQueryKind") }}
        $this->convertv2 = $convertv2;
        {{- end }}
    }
}
