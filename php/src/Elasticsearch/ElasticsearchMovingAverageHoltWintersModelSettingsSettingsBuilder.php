<?php

namespace Grafana\Foundation\Elasticsearch;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Elasticsearch\ElasticsearchMovingAverageHoltWintersModelSettingsSettings>
 */
class ElasticsearchMovingAverageHoltWintersModelSettingsSettingsBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Elasticsearch\ElasticsearchMovingAverageHoltWintersModelSettingsSettings $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Elasticsearch\ElasticsearchMovingAverageHoltWintersModelSettingsSettings();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Elasticsearch\ElasticsearchMovingAverageHoltWintersModelSettingsSettings
     */
    public function build()
    {
        return $this->internal;
    }

    public function alpha(string $alpha): static
    {
        $this->internal->alpha = $alpha;
    
        return $this;
    }
    public function beta(string $beta): static
    {
        $this->internal->beta = $beta;
    
        return $this;
    }
    public function gamma(string $gamma): static
    {
        $this->internal->gamma = $gamma;
    
        return $this;
    }
    public function period(string $period): static
    {
        $this->internal->period = $period;
    
        return $this;
    }
    public function pad(bool $pad): static
    {
        $this->internal->pad = $pad;
    
        return $this;
    }

}
