<?php

namespace Grafana\Foundation\Dashboard;

/**
 * Maps special values like Null, NaN (not a number), and boolean values like true and false to a display text and color.
 * See SpecialValueMatch to see the list of special values.
 * For example, you can configure a special value mapping so that null values appear as N/A.
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboard\SpecialValueMap>
 */
class SpecialValueMapBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboard\SpecialValueMap $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Dashboard\SpecialValueMap();
    $this->internal->type = "special";
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Dashboard\SpecialValueMap
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboard\DashboardSpecialValueMapOptions> $options
     */
    public function options(\Grafana\Foundation\Cog\Builder $options): static
    {
        $optionsResource = $options->build();
        $this->internal->options = $optionsResource;
    
        return $this;
    }

}
