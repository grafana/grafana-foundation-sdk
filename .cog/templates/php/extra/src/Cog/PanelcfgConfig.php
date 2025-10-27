<?php

namespace {{ .Data.NamespaceRoot }}\Cog;
{{- $convertReturn := "\\Grafana\\Foundation\\Dashboard\\Panel" -}}
{{- if (objectExists "dashboardv2beta1" "VizConfigKind") -}}
  {{- $convertReturn = "mixed" -}}
{{- end -}}

final class PanelcfgConfig
{
    public readonly string $identifier;

    /**
     * @var (callable({{ $convertReturn }}): string)|null
     */
    public $convert;

    /**
     * @var (callable(array<string, mixed>): object)|null
     */
    public $optionsFromArray;

    /**
     * @var (callable(array<string, mixed>): object)|null
     */
    public $fieldConfigFromArray;

    /**
     * @param (callable({{ $convertReturn }}): string)|null $convert
     * @param (callable(array<string, mixed>): object)|null $optionsFromArray
     * @param (callable(array<string, mixed>): object)|null $fieldConfigFromArray
     */
    public function __construct(string $identifier, ?callable $convert = null, ?callable $optionsFromArray = null, ?callable $fieldConfigFromArray = null)
    {
        $this->identifier = $identifier;
        $this->convert = $convert;
        $this->optionsFromArray = $optionsFromArray;
        $this->fieldConfigFromArray = $fieldConfigFromArray;
    }
}
