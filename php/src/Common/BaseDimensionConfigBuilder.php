<?php

namespace Grafana\Foundation\Common;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\BaseDimensionConfig>
 */
class BaseDimensionConfigBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Common\BaseDimensionConfig $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Common\BaseDimensionConfig();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Common\BaseDimensionConfig
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * fixed: T -- will be added by each element
     */
    public function field(string $field): static
    {
        $this->internal->field = $field;
    
        return $this;
    }

}
