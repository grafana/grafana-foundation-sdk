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
}