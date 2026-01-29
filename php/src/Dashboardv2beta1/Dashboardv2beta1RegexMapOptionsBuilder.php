<?php

namespace Grafana\Foundation\Dashboardv2beta1;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\Dashboardv2beta1RegexMapOptions>
 */
class Dashboardv2beta1RegexMapOptionsBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboardv2beta1\Dashboardv2beta1RegexMapOptions $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Dashboardv2beta1\Dashboardv2beta1RegexMapOptions();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Dashboardv2beta1\Dashboardv2beta1RegexMapOptions
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * Regular expression to match against
     */
    public function pattern(string $pattern): static
    {
        $this->internal->pattern = $pattern;
    
        return $this;
    }

    /**
     * Config to apply when the value matches the regex
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\ValueMappingResult> $result
     */
    public function result(\Grafana\Foundation\Cog\Builder $result): static
    {
        $resultResource = $result->build();
        $this->internal->result = $resultResource;
    
        return $this;
    }

}
