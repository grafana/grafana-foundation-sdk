<?php

namespace Grafana\Foundation\Prometheus;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Prometheus\Scopes>
 */
class ScopesBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Prometheus\Scopes $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Prometheus\Scopes();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Prometheus\Scopes
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * @param array<string> $defaultPath
     */
    public function defaultPath(array $defaultPath): static
    {
        $this->internal->defaultPath = $defaultPath;
    
        return $this;
    }

    /**
     * @param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Prometheus\ScopesFilters>> $filters
     */
    public function filters(array $filters): static
    {
            $filtersResources = [];
            foreach ($filters as $r1) {
                    $filtersResources[] = $r1->build();
            }
        $this->internal->filters = $filtersResources;
    
        return $this;
    }

    public function title(string $title): static
    {
        $this->internal->title = $title;
    
        return $this;
    }

}
