<?php

namespace Grafana\Foundation\Dashboard;

/**
 * Maps text values to a color or different display text and color.
 * For example, you can configure a value mapping so that all instances of the value 10 appear as Perfection! rather than the number.
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboard\ValueMap>
 */
class ValueMapBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboard\ValueMap $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Dashboard\ValueMap();
    $this->internal->type = "value";
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Dashboard\ValueMap
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * Map with <value_to_match>: ValueMappingResult. For example: { "10": { text: "Perfection!", color: "green" } }
     * @param array<string, \Grafana\Foundation\Dashboard\ValueMappingResult> $options
     */
    public function options(array $options): static
    {
        $this->internal->options = $options;
    
        return $this;
    }

}
