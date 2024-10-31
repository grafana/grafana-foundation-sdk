<?php

namespace Grafana\Foundation\Common;

/**
 * TODO docs
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\HideableFieldConfig>
 */
class HideableFieldConfigBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Common\HideableFieldConfig $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Common\HideableFieldConfig();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Common\HideableFieldConfig
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
