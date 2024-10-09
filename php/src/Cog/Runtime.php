<?php

namespace Grafana\Foundation\Cog;

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
        $this->registerPanelcfgVariant(\Grafana\Foundation\Alertgroups\VariantConfig::get());
        $this->registerPanelcfgVariant(\Grafana\Foundation\Annotationslist\VariantConfig::get());
        $this->registerPanelcfgVariant(\Grafana\Foundation\Barchart\VariantConfig::get());
        $this->registerPanelcfgVariant(\Grafana\Foundation\Bargauge\VariantConfig::get());
        $this->registerPanelcfgVariant(\Grafana\Foundation\Candlestick\VariantConfig::get());
        $this->registerPanelcfgVariant(\Grafana\Foundation\Canvas\VariantConfig::get());
        $this->registerPanelcfgVariant(\Grafana\Foundation\Dashboardlist\VariantConfig::get());
        $this->registerPanelcfgVariant(\Grafana\Foundation\Datagrid\VariantConfig::get());
        $this->registerPanelcfgVariant(\Grafana\Foundation\Debug\VariantConfig::get());
        $this->registerPanelcfgVariant(\Grafana\Foundation\Gauge\VariantConfig::get());
        $this->registerPanelcfgVariant(\Grafana\Foundation\Geomap\VariantConfig::get());
        $this->registerPanelcfgVariant(\Grafana\Foundation\Heatmap\VariantConfig::get());
        $this->registerPanelcfgVariant(\Grafana\Foundation\Histogram\VariantConfig::get());
        $this->registerPanelcfgVariant(\Grafana\Foundation\Logs\VariantConfig::get());
        $this->registerPanelcfgVariant(\Grafana\Foundation\News\VariantConfig::get());
        $this->registerPanelcfgVariant(\Grafana\Foundation\Nodegraph\VariantConfig::get());
        $this->registerPanelcfgVariant(\Grafana\Foundation\Piechart\VariantConfig::get());
        $this->registerPanelcfgVariant(\Grafana\Foundation\Stat\VariantConfig::get());
        $this->registerPanelcfgVariant(\Grafana\Foundation\Statetimeline\VariantConfig::get());
        $this->registerPanelcfgVariant(\Grafana\Foundation\Statushistory\VariantConfig::get());
        $this->registerPanelcfgVariant(\Grafana\Foundation\Table\VariantConfig::get());
        $this->registerPanelcfgVariant(\Grafana\Foundation\Text\VariantConfig::get());
        $this->registerPanelcfgVariant(\Grafana\Foundation\Timeseries\VariantConfig::get());
        $this->registerPanelcfgVariant(\Grafana\Foundation\Trend\VariantConfig::get());
        $this->registerPanelcfgVariant(\Grafana\Foundation\Xychart\VariantConfig::get());
        $this->registerDataqueryVariant(\Grafana\Foundation\Azuremonitor\VariantConfig::get());
        $this->registerDataqueryVariant(\Grafana\Foundation\Cloudwatch\VariantConfig::get());
        $this->registerDataqueryVariant(\Grafana\Foundation\Elasticsearch\VariantConfig::get());
        $this->registerDataqueryVariant(\Grafana\Foundation\Expr\VariantConfig::get());
        $this->registerDataqueryVariant(\Grafana\Foundation\Googlecloudmonitoring\VariantConfig::get());
        $this->registerDataqueryVariant(\Grafana\Foundation\Grafanapyroscope\VariantConfig::get());
        $this->registerDataqueryVariant(\Grafana\Foundation\Loki\VariantConfig::get());
        $this->registerDataqueryVariant(\Grafana\Foundation\Parca\VariantConfig::get());
        $this->registerDataqueryVariant(\Grafana\Foundation\Prometheus\VariantConfig::get());
        $this->registerDataqueryVariant(\Grafana\Foundation\Tempo\VariantConfig::get());
        $this->registerDataqueryVariant(\Grafana\Foundation\Testdata\VariantConfig::get());
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

    public function convertPanelToCode(\Grafana\Foundation\Dashboard\Panel $panel, string $panelType): string
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
            return '(new \Grafana\Foundation\Cog\UnknownDataqueryBuilder(new \Grafana\Foundation\Cog\UnknownDataquery('.var_export($dataquery->toArray(), true).')))';
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
}
