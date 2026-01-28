<?php

namespace Grafana\Foundation\Piechart;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Piechart\FieldConfig>
 */
class FieldConfigBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Piechart\FieldConfig $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Piechart\FieldConfig();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Piechart\FieldConfig
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\HideSeriesConfig> $hideFrom
     */
    public function hideFrom(\Grafana\Foundation\Cog\Builder $hideFrom): static
    {
        $hideFromResource = $hideFrom->build();
        $this->internal->hideFrom = $hideFromResource;
    
        return $this;
    }

}
