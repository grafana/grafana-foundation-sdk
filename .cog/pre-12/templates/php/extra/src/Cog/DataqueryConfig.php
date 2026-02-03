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
    
    {{- if objectExists "dashboardv2beta1" "DataQueryKind" }}
    /**
     * @var (callable(\Grafana\Foundation\Dashboardv2beta1\DataQueryKind): string)|null
     */
    public $convertv2;
    {{- end }}

    /**
     * @param callable(array<string, mixed>): Dataquery $fromArray
     * @param (callable(Dataquery): string)|null $convert
     {{- if objectExists "dashboardv2beta1" "DataQueryKind" }} 
     * @param (callable(\Grafana\Foundation\Dashboardv2beta1\DataQueryKind): string)|null $convertv2 
     {{- end }}
     */
    public function __construct(string $identifier, callable $fromArray, ?callable $convert = null{{ if objectExists "dashboardv2beta1" "DataQueryKind" }}, ?callable $convertv2 = null{{ end }})
    {
        $this->identifier = $identifier;
        $this->fromArray = $fromArray;
        $this->convert = $convert;
        {{- if objectExists "dashboardv2beta1" "DataQueryKind" }}
        $this->convertv2 = $convertv2;
        {{- end }}
    }
}
