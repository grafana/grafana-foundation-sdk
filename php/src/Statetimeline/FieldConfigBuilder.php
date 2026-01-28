<?php

namespace Grafana\Foundation\Statetimeline;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Statetimeline\FieldConfig>
 */
class FieldConfigBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Statetimeline\FieldConfig $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Statetimeline\FieldConfig();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Statetimeline\FieldConfig
     */
    public function build()
    {
        return $this->internal;
    }

    public function lineWidth(int $lineWidth): static
    {
        if (!($lineWidth <= 10)) {
            throw new \ValueError('$lineWidth must be <= 10');
        }
        $this->internal->lineWidth = $lineWidth;
    
        return $this;
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

    public function fillOpacity(int $fillOpacity): static
    {
        if (!($fillOpacity <= 100)) {
            throw new \ValueError('$fillOpacity must be <= 100');
        }
        $this->internal->fillOpacity = $fillOpacity;
    
        return $this;
    }

}
