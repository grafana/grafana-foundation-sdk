<?php

namespace Grafana\Foundation\Alertgroups;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Alertgroups\Options>
 */
class OptionsBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Alertgroups\Options $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Alertgroups\Options();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Alertgroups\Options
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * Comma-separated list of values used to filter alert results
     */
    public function labels(string $labels): static
    {
        $this->internal->labels = $labels;
    
        return $this;
    }

    /**
     * Name of the alertmanager used as a source for alerts
     */
    public function alertmanager(string $alertmanager): static
    {
        $this->internal->alertmanager = $alertmanager;
    
        return $this;
    }

    /**
     * Expand all alert groups by default
     */
    public function expandAll(bool $expandAll): static
    {
        $this->internal->expandAll = $expandAll;
    
        return $this;
    }

}
