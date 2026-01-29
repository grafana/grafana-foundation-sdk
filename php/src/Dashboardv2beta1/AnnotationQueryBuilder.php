<?php

namespace Grafana\Foundation\Dashboardv2beta1;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\AnnotationQueryKind>
 */
class AnnotationQueryBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboardv2beta1\AnnotationQueryKind $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Dashboardv2beta1\AnnotationQueryKind();
    $this->internal->kind = "AnnotationQuery";
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Dashboardv2beta1\AnnotationQueryKind
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\AnnotationQuerySpec> $spec
     */
    public function spec(\Grafana\Foundation\Cog\Builder $spec): static
    {
        $specResource = $spec->build();
        $this->internal->spec = $specResource;
    
        return $this;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\DataQueryKind> $query
     */
    public function query(\Grafana\Foundation\Cog\Builder $query): static
    {
        $queryResource = $query->build();
        $this->internal->spec->query = $queryResource;
    
        return $this;
    }

    public function enable(bool $enable): static
    {
        $this->internal->spec->enable = $enable;
    
        return $this;
    }

    public function hide(bool $hide): static
    {
        $this->internal->spec->hide = $hide;
    
        return $this;
    }

    public function iconColor(string $iconColor): static
    {
        $this->internal->spec->iconColor = $iconColor;
    
        return $this;
    }

    public function name(string $name): static
    {
        $this->internal->spec->name = $name;
    
        return $this;
    }

    public function builtIn(bool $builtIn): static
    {
        $this->internal->spec->builtIn = $builtIn;
    
        return $this;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\AnnotationPanelFilter> $filter
     */
    public function filter(\Grafana\Foundation\Cog\Builder $filter): static
    {
        $filterResource = $filter->build();
        $this->internal->spec->filter = $filterResource;
    
        return $this;
    }

    /**
     * Placement can be used to display the annotation query somewhere else on the dashboard other than the default location.
     */
    public function placement(string $placement): static
    {
        $this->internal->spec->placement = $placement;
    
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
        $this->internal->spec->mappings = $mappingsResources;
    
        return $this;
    }

    /**
     * Catch-all field for datasource-specific properties. Should not be available in as code tooling.
     * @param array<string, mixed> $legacyOptions
     */
    public function legacyOptions(array $legacyOptions): static
    {
        $this->internal->spec->legacyOptions = $legacyOptions;
    
        return $this;
    }

}
