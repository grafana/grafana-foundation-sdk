<?php

namespace Grafana\Foundation\Dashboard;

/**
 * TODO docs
 * FROM: AnnotationQuery in grafana-data/src/types/annotations.ts
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboard\AnnotationQuery>
 */
class AnnotationQueryBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboard\AnnotationQuery $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Dashboard\AnnotationQuery();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Dashboard\AnnotationQuery
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * Name of annotation.
     */
    public function name(string $name): static
    {
        $this->internal->name = $name;
    
        return $this;
    }
    /**
     * Datasource where the annotations data is
     */
    public function datasource(\Grafana\Foundation\Dashboard\DataSourceRef $datasource): static
    {
        $this->internal->datasource = $datasource;
    
        return $this;
    }
    /**
     * When enabled the annotation query is issued with every dashboard refresh
     */
    public function enable(bool $enable): static
    {
        $this->internal->enable = $enable;
    
        return $this;
    }
    /**
     * Annotation queries can be toggled on or off at the top of the dashboard.
     * When hide is true, the toggle is not shown in the dashboard.
     */
    public function hide(bool $hide): static
    {
        $this->internal->hide = $hide;
    
        return $this;
    }
    /**
     * Color to use for the annotation event markers
     */
    public function iconColor(string $iconColor): static
    {
        $this->internal->iconColor = $iconColor;
    
        return $this;
    }
    /**
     * Filters to apply when fetching annotations
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboard\AnnotationPanelFilter> $filter
     */
    public function filter(\Grafana\Foundation\Cog\Builder $filter): static
    {
        $filterResource = $filter->build();
        $this->internal->filter = $filterResource;
    
        return $this;
    }
    /**
     * TODO.. this should just be a normal query target
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboard\AnnotationTarget> $target
     */
    public function target(\Grafana\Foundation\Cog\Builder $target): static
    {
        $targetResource = $target->build();
        $this->internal->target = $targetResource;
    
        return $this;
    }
    /**
     * TODO -- this should not exist here, it is based on the --grafana-- datasource
     */
    public function type(string $type): static
    {
        $this->internal->type = $type;
    
        return $this;
    }
    public function expr(string $expr): static
    {
        $this->internal->expr = $expr;
    
        return $this;
    }

}
