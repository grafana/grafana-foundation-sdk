<?php

namespace Grafana\Foundation\Dashboardv2beta1;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\V2QueryGroupSpec>
 */
class V2QueryGroupSpecBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboardv2beta1\V2QueryGroupSpec $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Dashboardv2beta1\V2QueryGroupSpec();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Dashboardv2beta1\V2QueryGroupSpec
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * @param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\PanelQueryKind>> $queries
     */
    public function queries(array $queries): static
    {
            $queriesResources = [];
            foreach ($queries as $r1) {
                    $queriesResources[] = $r1->build();
            }
        $this->internal->queries = $queriesResources;
    
        return $this;
    }

    /**
     * @param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\V2TransformationKind>> $transformations
     */
    public function transformations(array $transformations): static
    {
            $transformationsResources = [];
            foreach ($transformations as $r1) {
                    $transformationsResources[] = $r1->build();
            }
        $this->internal->transformations = $transformationsResources;
    
        return $this;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\QueryOptionsSpec> $queryOptions
     */
    public function queryOptions(\Grafana\Foundation\Cog\Builder $queryOptions): static
    {
        $queryOptionsResource = $queryOptions->build();
        $this->internal->queryOptions = $queryOptionsResource;
    
        return $this;
    }

}
