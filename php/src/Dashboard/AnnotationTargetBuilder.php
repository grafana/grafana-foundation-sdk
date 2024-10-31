<?php

namespace Grafana\Foundation\Dashboard;

/**
 * TODO: this should be a regular DataQuery that depends on the selected dashboard
 * these match the properties of the "grafana" datasouce that is default in most dashboards
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboard\AnnotationTarget>
 */
class AnnotationTargetBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboard\AnnotationTarget $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Dashboard\AnnotationTarget();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Dashboard\AnnotationTarget
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * Only required/valid for the grafana datasource...
     * but code+tests is already depending on it so hard to change
     */
    public function limit(int $limit): static
    {
        $this->internal->limit = $limit;
    
        return $this;
    }
    /**
     * Only required/valid for the grafana datasource...
     * but code+tests is already depending on it so hard to change
     */
    public function matchAny(bool $matchAny): static
    {
        $this->internal->matchAny = $matchAny;
    
        return $this;
    }
    /**
     * Only required/valid for the grafana datasource...
     * but code+tests is already depending on it so hard to change
     * @param array<string> $tags
     */
    public function tags(array $tags): static
    {
        $this->internal->tags = $tags;
    
        return $this;
    }
    /**
     * Only required/valid for the grafana datasource...
     * but code+tests is already depending on it so hard to change
     */
    public function type(string $type): static
    {
        $this->internal->type = $type;
    
        return $this;
    }

}
