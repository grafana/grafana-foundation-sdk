<?php

namespace Grafana\Foundation\Debug;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Debug\UpdateConfig>
 */
class UpdateConfigBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Debug\UpdateConfig $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Debug\UpdateConfig();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Debug\UpdateConfig
     */
    public function build()
    {
        return $this->internal;
    }

    public function render(bool $render): static
    {
        $this->internal->render = $render;
    
        return $this;
    }
    public function dataChanged(bool $dataChanged): static
    {
        $this->internal->dataChanged = $dataChanged;
    
        return $this;
    }
    public function schemaChanged(bool $schemaChanged): static
    {
        $this->internal->schemaChanged = $schemaChanged;
    
        return $this;
    }

}
