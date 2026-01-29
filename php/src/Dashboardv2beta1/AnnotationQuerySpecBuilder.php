<?php

namespace Grafana\Foundation\Dashboardv2beta1;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\AnnotationQuerySpec>
 */
class AnnotationQuerySpecBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboardv2beta1\AnnotationQuerySpec $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Dashboardv2beta1\AnnotationQuerySpec();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Dashboardv2beta1\AnnotationQuerySpec
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\DataQueryKind> $query
     */
    public function query(\Grafana\Foundation\Cog\Builder $query): static
    {
        $queryResource = $query->build();
        $this->internal->query = $queryResource;
    
        return $this;
    }

    public function enable(bool $enable): static
    {
        $this->internal->enable = $enable;
    
        return $this;
    }

    public function hide(bool $hide): static
    {
        $this->internal->hide = $hide;
    
        return $this;
    }

    public function iconColor(string $iconColor): static
    {
        $this->internal->iconColor = $iconColor;
    
        return $this;
    }

    public function name(string $name): static
    {
        $this->internal->name = $name;
    
        return $this;
    }

    public function builtIn(bool $builtIn): static
    {
        $this->internal->builtIn = $builtIn;
    
        return $this;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\AnnotationPanelFilter> $filter
     */
    public function filter(\Grafana\Foundation\Cog\Builder $filter): static
    {
        $filterResource = $filter->build();
        $this->internal->filter = $filterResource;
    
        return $this;
    }

    /**
     * Placement can be used to display the annotation query somewhere else on the dashboard other than the default location.
     */
    public function placement(string $placement): static
    {
        $this->internal->placement = $placement;
    
        return $this;
    }

    /**
     * Mappings define how to convert data frame fields to annotation event fields.
     * @param array<string, \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\AnnotationEventFieldMapping>> $mappings
     */
    public function mappings(array $mappings): static
    {
            $mappingsResources = [];
            foreach ($mappings as $key1 => $val1) {
                    $mappingsResources[$key1] = $val1->build();
            }
        $this->internal->mappings = $mappingsResources;
    
        return $this;
    }

    /**
     * Catch-all field for datasource-specific properties. Should not be available in as code tooling.
     * @param array<string, mixed> $legacyOptions
     */
    public function legacyOptions(array $legacyOptions): static
    {
        $this->internal->legacyOptions = $legacyOptions;
    
        return $this;
    }

}
