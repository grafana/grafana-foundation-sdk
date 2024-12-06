<?php

namespace Grafana\Foundation\Dashboard;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboard\DashboardRegexMapOptions>
 */
class DashboardRegexMapOptionsBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboard\DashboardRegexMapOptions $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Dashboard\DashboardRegexMapOptions();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Dashboard\DashboardRegexMapOptions
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
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboard\ValueMappingResult> $result
     */
    public function result(\Grafana\Foundation\Cog\Builder $result): static
    {
        $resultResource = $result->build();
        $this->internal->result = $resultResource;
    
        return $this;
    }

}
