{{- $panelPackages := .Context.PackagesForVariant "panelcfg" -}}
{{- $dataqueryPackages := .Context.PackagesForVariant "dataquery" -}}
<?php

namespace {{ .Data.NamespaceRoot }}\Cog;

final class Runtime
{
    /**
     * @var array<string, PanelcfgConfig>
     */
    private $panelcfgVariants = [];

    /**
     * @var array<string, DataqueryConfig>
     */
    private $dataqueryVariants = [];

    private static ?self $instance = null;

    private function __construct()
    {
{{- range $pkg := $panelPackages }}
        $this->registerPanelcfgVariant(\{{ $.Data.NamespaceRoot }}\{{ $pkg|formatPackageName }}\VariantConfig::get());
{{- end }}

{{- range $pkg := $dataqueryPackages }}
        $this->registerDataqueryVariant(\{{ $.Data.NamespaceRoot }}\{{ $pkg|formatPackageName }}\VariantConfig::get());
{{- end }}
    }

    public static function get(): self
    {
        if (self::$instance === null) {
            self::$instance = new self();
        }

        return self::$instance;
    }

    public function registerPanelcfgVariant(PanelcfgConfig $variantConfig): void
    {
        $this->panelcfgVariants[$variantConfig->identifier] = $variantConfig;
    }

    public function registerDataqueryVariant(DataqueryConfig $variantConfig): void
    {
        $this->dataqueryVariants[$variantConfig->identifier] = $variantConfig;
    }

    public function panelcfgVariantExists(string $identifier): bool
    {
        return isset($this->panelcfgVariants[$identifier]);
    }

    public function panelcfgVariantConfig(string $identifier): PanelcfgConfig
    {
        if (!$this->panelcfgVariantExists($identifier)) {
            throw new \ValueError("$identifier panelcfg does not exist");
        }

        return $this->panelcfgVariants[$identifier];
    }

    /**
     * @param array<string, mixed> $data
     */
    public function dataqueryFromArray(array $data, string $dataqueryTypeHint): Dataquery
    {
        // Dataqueries might reference the datasource to use, and its type. Let's use that.
        if (empty($dataqueryTypeHint) && !empty($data['datasource']) && !empty($data['datasource']['type'])) {
            $dataqueryTypeHint = $data['datasource']['type'];
        }

        // A hint tells us the dataquery type: let's use it.
        if (!empty($dataqueryTypeHint) && isset($this->dataqueryVariants[$dataqueryTypeHint])) {
            $fromArray = $this->dataqueryVariants[$dataqueryTypeHint]->fromArray;

            return $fromArray($data);
        }

        // We have no idea what type the dataquery is: use our `UnknownDataquery` bag to not lose data.
        return new UnknownDataquery($data);
    }

    /**
     * @param array<array<string, mixed>> $data
     * @return Dataquery[]
     */
    public function dataqueriesFromArray(array $data, string $dataqueryTypeHint): array
    {
        $queries = [];
        foreach ($data as $query) {
            $queries[] = $this->dataqueryFromArray($query, $dataqueryTypeHint);
        }
        return $queries;
    }

    {{ $panelArgument := "\\Grafana\\Foundation\\Dashboard\\Panel" }}
    {{- if objectExists "dashboardv2beta1" "VizConfigKind" }}
    {{- $panelArgument = "mixed" }}
    {{- end }}

    public function convertPanelToCode({{ $panelArgument }} $panel, string $panelType): string
    {
        if (!$this->panelcfgVariantExists($panelType)) {
            return '/* could not convert panel to PHP */';
        }

        $convert = $this->panelcfgVariants[$panelType]->convert;
        if ($convert === null) {
            return '/* could not convert panel to PHP */';
        }

        return $convert($panel);
    }

    public function convertDataqueryToCode(Dataquery $dataquery): string
    {
        if ($dataquery instanceof UnknownDataquery) {
            return '(new \{{ .Data.NamespaceRoot }}\Cog\UnknownDataqueryBuilder(new \{{ .Data.NamespaceRoot }}\Cog\UnknownDataquery('.var_export($dataquery->toArray(), true).')))';
        }

        if (!isset($this->dataqueryVariants[$dataquery->dataqueryType()])) {
            return '/* could not convert dataquery to PHP */';
        }

        $convert = $this->dataqueryVariants[$dataquery->dataqueryType()]->convert;
        if ($convert === null) {
            return '/* could not convert dataquery to PHP */';
        }

        return $convert($dataquery);
    }

    {{- if objectExists "dashboardv2beta1" "DataQueryKind" }}
    public function convertDataQueryKindToCode(\Grafana\Foundation\Dashboardv2beta1\DataQueryKind $dataquery, string $group): string 
    {
    	if (!isset($this->dataqueryVariants[$group])) {
            return '/* could not convert DataQueryKind to PHP */';
        }
    
        $convert = $this->dataqueryVariants[$group]->convertv2;
        if ($convert === null) {
            return '/* could not convert DataQueryKind to PHP */';
        }
        
        return $convert($dataquery);
    }
    {{- end }}
}
