<?php

namespace Grafana\Foundation\Expr;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\DataQueryKind>
 */
class QueryBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Dashboardv2beta1\DataQueryKind $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Dashboardv2beta1\DataQueryKind();
    $this->internal->kind = "DataQuery";
    $this->internal->group = "__expr__";
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Dashboardv2beta1\DataQueryKind
     */
    public function build()
    {
        return $this->internal;
    }

    public function version(string $version): static
    {
        $this->internal->version = $version;
    
        return $this;
    }

    /**
     * New type for datasource reference
     * Not creating a new type until we figure out how to handle DS refs for group by, adhoc, and every place that uses DataSourceRef in TS.
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\Dashboardv2beta1DataQueryKindDatasource> $datasource
     */
    public function datasource(\Grafana\Foundation\Cog\Builder $datasource): static
    {
        $datasourceResource = $datasource->build();
        $this->internal->datasource = $datasourceResource;
    
        return $this;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Expr\TypeMath> $typeMath
     */
    public function typeMath(\Grafana\Foundation\Cog\Builder $typeMath): static
    {
        $typeMathResource = $typeMath->build();
        $this->internal->spec = $typeMathResource;
    
        return $this;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Expr\TypeReduce> $typeReduce
     */
    public function typeReduce(\Grafana\Foundation\Cog\Builder $typeReduce): static
    {
        $typeReduceResource = $typeReduce->build();
        $this->internal->spec = $typeReduceResource;
    
        return $this;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Expr\TypeResample> $typeResample
     */
    public function typeResample(\Grafana\Foundation\Cog\Builder $typeResample): static
    {
        $typeResampleResource = $typeResample->build();
        $this->internal->spec = $typeResampleResource;
    
        return $this;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Expr\TypeClassicConditions> $typeClassicConditions
     */
    public function typeClassicConditions(\Grafana\Foundation\Cog\Builder $typeClassicConditions): static
    {
        $typeClassicConditionsResource = $typeClassicConditions->build();
        $this->internal->spec = $typeClassicConditionsResource;
    
        return $this;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Expr\TypeThreshold> $typeThreshold
     */
    public function typeThreshold(\Grafana\Foundation\Cog\Builder $typeThreshold): static
    {
        $typeThresholdResource = $typeThreshold->build();
        $this->internal->spec = $typeThresholdResource;
    
        return $this;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Expr\TypeSql> $typeSql
     */
    public function typeSql(\Grafana\Foundation\Cog\Builder $typeSql): static
    {
        $typeSqlResource = $typeSql->build();
        $this->internal->spec = $typeSqlResource;
    
        return $this;
    }

}
