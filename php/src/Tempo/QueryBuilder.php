<?php

namespace Grafana\Foundation\Tempo;

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
    $this->internal->group = "tempo";
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
     * A unique identifier for the query within the list of targets.
     * In server side expressions, the refId is used as a variable name to identify results.
     * By default, the UI will assign A->Z; however setting meaningful names may be useful.
     */
    public function refId(string $refId): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Tempo\TempoQuery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Tempo\TempoQuery);
        $this->internal->spec->refId = $refId;
    
        return $this;
    }

    /**
     * If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.
     */
    public function hide(bool $hide): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Tempo\TempoQuery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Tempo\TempoQuery);
        $this->internal->spec->hide = $hide;
    
        return $this;
    }

    /**
     * Specify the query flavor
     * TODO make this required and give it a default
     */
    public function queryType(string $queryType): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Tempo\TempoQuery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Tempo\TempoQuery);
        $this->internal->spec->queryType = $queryType;
    
        return $this;
    }

    /**
     * TraceQL query or trace ID
     */
    public function query(string $query): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Tempo\TempoQuery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Tempo\TempoQuery);
        $this->internal->spec->query = $query;
    
        return $this;
    }

    /**
     * @deprecated Logfmt query to filter traces by their tags. Example: http.status_code=200 error=true
     */
    public function search(string $search): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Tempo\TempoQuery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Tempo\TempoQuery);
        $this->internal->spec->search = $search;
    
        return $this;
    }

    /**
     * @deprecated Query traces by service name
     */
    public function serviceName(string $serviceName): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Tempo\TempoQuery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Tempo\TempoQuery);
        $this->internal->spec->serviceName = $serviceName;
    
        return $this;
    }

    /**
     * @deprecated Query traces by span name
     */
    public function spanName(string $spanName): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Tempo\TempoQuery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Tempo\TempoQuery);
        $this->internal->spec->spanName = $spanName;
    
        return $this;
    }

    /**
     * @deprecated Define the minimum duration to select traces. Use duration format, for example: 1.2s, 100ms
     */
    public function minDuration(string $minDuration): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Tempo\TempoQuery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Tempo\TempoQuery);
        $this->internal->spec->minDuration = $minDuration;
    
        return $this;
    }

    /**
     * @deprecated Define the maximum duration to select traces. Use duration format, for example: 1.2s, 100ms
     */
    public function maxDuration(string $maxDuration): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Tempo\TempoQuery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Tempo\TempoQuery);
        $this->internal->spec->maxDuration = $maxDuration;
    
        return $this;
    }

    /**
     * Filters to be included in a PromQL query to select data for the service graph. Example: {client="app",service="app"}. Providing multiple values will produce union of results for each filter, using PromQL OR operator internally.
     * @param string|array<string> $serviceMapQuery
     */
    public function serviceMapQuery( $serviceMapQuery): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Tempo\TempoQuery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Tempo\TempoQuery);
        $this->internal->spec->serviceMapQuery = $serviceMapQuery;
    
        return $this;
    }

    /**
     * Use service.namespace in addition to service.name to uniquely identify a service.
     */
    public function serviceMapIncludeNamespace(bool $serviceMapIncludeNamespace): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Tempo\TempoQuery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Tempo\TempoQuery);
        $this->internal->spec->serviceMapIncludeNamespace = $serviceMapIncludeNamespace;
    
        return $this;
    }

    /**
     * Defines the maximum number of traces that are returned from Tempo
     */
    public function limit(int $limit): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Tempo\TempoQuery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Tempo\TempoQuery);
        $this->internal->spec->limit = $limit;
    
        return $this;
    }

    /**
     * Defines the maximum number of spans per spanset that are returned from Tempo
     */
    public function spss(int $spss): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Tempo\TempoQuery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Tempo\TempoQuery);
        $this->internal->spec->spss = $spss;
    
        return $this;
    }

    /**
     * @param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Tempo\TraceqlFilter>> $filters
     */
    public function filters(array $filters): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Tempo\TempoQuery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Tempo\TempoQuery);
            $filtersResources = [];
            foreach ($filters as $r1) {
                    $filtersResources[] = $r1->build();
            }
        $this->internal->spec->filters = $filtersResources;
    
        return $this;
    }

    /**
     * deprecated Filters that are used to query the metrics summary
     * @param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Tempo\TraceqlFilter>> $groupBy
     */
    public function groupBy(array $groupBy): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Tempo\TempoQuery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Tempo\TempoQuery);
            $groupByResources = [];
            foreach ($groupBy as $r1) {
                    $groupByResources[] = $r1->build();
            }
        $this->internal->spec->groupBy = $groupByResources;
    
        return $this;
    }

    /**
     * The type of the table that is used to display the search results
     */
    public function tableType(\Grafana\Foundation\Tempo\SearchTableType $tableType): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Tempo\TempoQuery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Tempo\TempoQuery);
        $this->internal->spec->tableType = $tableType;
    
        return $this;
    }

    /**
     * For metric queries, the step size to use
     */
    public function step(string $step): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Tempo\TempoQuery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Tempo\TempoQuery);
        $this->internal->spec->step = $step;
    
        return $this;
    }

    /**
     * For metric queries, how many exemplars to request, 0 means no exemplars
     */
    public function exemplars(int $exemplars): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Tempo\TempoQuery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Tempo\TempoQuery);
        $this->internal->spec->exemplars = $exemplars;
    
        return $this;
    }

    /**
     * For mixed data sources the selected datasource is on the query level.
     * For non mixed scenarios this is undefined.
     * TODO find a better way to do this ^ that's friendly to schema
     * TODO this shouldn't be unknown but DataSourceRef | null
     */
    public function oldDatasource(\Grafana\Foundation\Common\DataSourceRef $datasource): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Tempo\TempoQuery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Tempo\TempoQuery);
        $this->internal->spec->datasource = $datasource;
    
        return $this;
    }

    /**
     * For metric queries, whether to run instant or range queries
     */
    public function metricsQueryType(\Grafana\Foundation\Tempo\MetricsQueryType $metricsQueryType): static
    {    
        if ($this->internal->spec === null) {
            $this->internal->spec = new \Grafana\Foundation\Tempo\TempoQuery();
        }
        assert($this->internal->spec instanceof \Grafana\Foundation\Tempo\TempoQuery);
        $this->internal->spec->metricsQueryType = $metricsQueryType;
    
        return $this;
    }

}
