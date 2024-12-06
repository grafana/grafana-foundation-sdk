<?php

namespace Grafana\Foundation\Tempo;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Tempo\TempoQuery>
 */
class TempoQueryBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Tempo\TempoQuery $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Tempo\TempoQuery();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Tempo\TempoQuery
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * A unique identifier for the query within the list of targets.
     * In server side expressions, the refId is used as a variable name to identify results.
     * By default, the UI will assign A->Z; however setting meaningful names may be useful.
     */
    public function refId(string $refId): static
    {
        $this->internal->refId = $refId;
    
        return $this;
    }
    /**
     * If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.
     */
    public function hide(bool $hide): static
    {
        $this->internal->hide = $hide;
    
        return $this;
    }
    /**
     * Specify the query flavor
     * TODO make this required and give it a default
     */
    public function queryType(string $queryType): static
    {
        $this->internal->queryType = $queryType;
    
        return $this;
    }
    /**
     * TraceQL query or trace ID
     */
    public function query(string $query): static
    {
        $this->internal->query = $query;
    
        return $this;
    }
    /**
     * @deprecated Logfmt query to filter traces by their tags. Example: http.status_code=200 error=true
     */
    public function search(string $search): static
    {
        $this->internal->search = $search;
    
        return $this;
    }
    /**
     * @deprecated Query traces by service name
     */
    public function serviceName(string $serviceName): static
    {
        $this->internal->serviceName = $serviceName;
    
        return $this;
    }
    /**
     * @deprecated Query traces by span name
     */
    public function spanName(string $spanName): static
    {
        $this->internal->spanName = $spanName;
    
        return $this;
    }
    /**
     * @deprecated Define the minimum duration to select traces. Use duration format, for example: 1.2s, 100ms
     */
    public function minDuration(string $minDuration): static
    {
        $this->internal->minDuration = $minDuration;
    
        return $this;
    }
    /**
     * @deprecated Define the maximum duration to select traces. Use duration format, for example: 1.2s, 100ms
     */
    public function maxDuration(string $maxDuration): static
    {
        $this->internal->maxDuration = $maxDuration;
    
        return $this;
    }
    /**
     * Filters to be included in a PromQL query to select data for the service graph. Example: {client="app",service="app"}. Providing multiple values will produce union of results for each filter, using PromQL OR operator internally.
     * @param string|array<string> $serviceMapQuery
     */
    public function serviceMapQuery( $serviceMapQuery): static
    {
        $this->internal->serviceMapQuery = $serviceMapQuery;
    
        return $this;
    }
    /**
     * Use service.namespace in addition to service.name to uniquely identify a service.
     */
    public function serviceMapIncludeNamespace(bool $serviceMapIncludeNamespace): static
    {
        $this->internal->serviceMapIncludeNamespace = $serviceMapIncludeNamespace;
    
        return $this;
    }
    /**
     * Defines the maximum number of traces that are returned from Tempo
     */
    public function limit(int $limit): static
    {
        $this->internal->limit = $limit;
    
        return $this;
    }
    /**
     * Defines the maximum number of spans per spanset that are returned from Tempo
     */
    public function spss(int $spss): static
    {
        $this->internal->spss = $spss;
    
        return $this;
    }
    /**
     * @param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Tempo\TraceqlFilter>> $filters
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
    /**
     * Filters that are used to query the metrics summary
     * @param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Tempo\TraceqlFilter>> $groupBy
     */
    public function groupBy(array $groupBy): static
    {
            $groupByResources = [];
            foreach ($groupBy as $r1) {
                    $groupByResources[] = $r1->build();
            }
        $this->internal->groupBy = $groupByResources;
    
        return $this;
    }
    /**
     * The type of the table that is used to display the search results
     */
    public function tableType(\Grafana\Foundation\Tempo\SearchTableType $tableType): static
    {
        $this->internal->tableType = $tableType;
    
        return $this;
    }
    /**
     * For mixed data sources the selected datasource is on the query level.
     * For non mixed scenarios this is undefined.
     * TODO find a better way to do this ^ that's friendly to schema
     * TODO this shouldn't be unknown but DataSourceRef | null
     */
    public function datasource(\Grafana\Foundation\Dashboard\DataSourceRef $datasource): static
    {
        $this->internal->datasource = $datasource;
    
        return $this;
    }
    /**
     * For metric queries, the step size to use
     */
    public function step(string $step): static
    {
        $this->internal->step = $step;
    
        return $this;
    }

}
