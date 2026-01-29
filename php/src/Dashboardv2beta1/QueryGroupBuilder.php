<?php

namespace Grafana\Foundation\Dashboardv2beta1;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\QueryGroupKind>
 */
class QueryGroupBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboardv2beta1\QueryGroupKind $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Dashboardv2beta1\QueryGroupKind();
    $this->internal->kind = "QueryGroup";
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Dashboardv2beta1\QueryGroupKind
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * @param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\PanelQueryKind>> $queries
     */
    public function targets(array $queries): static
    {
            $queriesResources = [];
            foreach ($queries as $r1) {
                    $queriesResources[] = $r1->build();
            }
        $this->internal->spec->queries = $queriesResources;
    
        return $this;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\PanelQueryKind> $querie
     */
    public function target(\Grafana\Foundation\Cog\Builder $querie): static
    {
        $querieResource = $querie->build();
        $this->internal->spec->queries[] = $querieResource;
    
        return $this;
    }

    /**
     * @param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\TransformationKind>> $transformations
     */
    public function transformations(array $transformations): static
    {
            $transformationsResources = [];
            foreach ($transformations as $r1) {
                    $transformationsResources[] = $r1->build();
            }
        $this->internal->spec->transformations = $transformationsResources;
    
        return $this;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\TransformationKind> $transformation
     */
    public function transformation(\Grafana\Foundation\Cog\Builder $transformation): static
    {
        $transformationResource = $transformation->build();
        $this->internal->spec->transformations[] = $transformationResource;
    
        return $this;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\QueryOptionsSpec> $queryOptions
     */
    public function queryOptions(\Grafana\Foundation\Cog\Builder $queryOptions): static
    {
        $queryOptionsResource = $queryOptions->build();
        $this->internal->spec->queryOptions = $queryOptionsResource;
    
        return $this;
    }

}
