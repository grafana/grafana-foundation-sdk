<?php

namespace Grafana\Foundation\Tempo;

class TempoQuery implements \JsonSerializable, \Grafana\Foundation\Cog\Dataquery
{
    /**
     * A unique identifier for the query within the list of targets.
     * In server side expressions, the refId is used as a variable name to identify results.
     * By default, the UI will assign A->Z; however setting meaningful names may be useful.
     */
    public string $refId;

    /**
     * If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.
     */
    public ?bool $hide;

    /**
     * Specify the query flavor
     * TODO make this required and give it a default
     */
    public ?string $queryType;

    /**
     * TraceQL query or trace ID
     */
    public ?string $query;

    /**
     * @deprecated Logfmt query to filter traces by their tags. Example: http.status_code=200 error=true
     */
    public ?string $search;

    /**
     * @deprecated Query traces by service name
     */
    public ?string $serviceName;

    /**
     * @deprecated Query traces by span name
     */
    public ?string $spanName;

    /**
     * @deprecated Define the minimum duration to select traces. Use duration format, for example: 1.2s, 100ms
     */
    public ?string $minDuration;

    /**
     * @deprecated Define the maximum duration to select traces. Use duration format, for example: 1.2s, 100ms
     */
    public ?string $maxDuration;

    /**
     * Filters to be included in a PromQL query to select data for the service graph. Example: {client="app",service="app"}. Providing multiple values will produce union of results for each filter, using PromQL OR operator internally.
     * @var string|array<string>|null
     */
    public $serviceMapQuery;

    /**
     * Use service.namespace in addition to service.name to uniquely identify a service.
     */
    public ?bool $serviceMapIncludeNamespace;

    /**
     * Defines the maximum number of traces that are returned from Tempo
     */
    public ?int $limit;

    /**
     * Defines the maximum number of spans per spanset that are returned from Tempo
     */
    public ?int $spss;

    /**
     * @var array<\Grafana\Foundation\Tempo\TraceqlFilter>
     */
    public array $filters;

    /**
     * Filters that are used to query the metrics summary
     * @var array<\Grafana\Foundation\Tempo\TraceqlFilter>|null
     */
    public ?array $groupBy;

    /**
     * The type of the table that is used to display the search results
     */
    public ?\Grafana\Foundation\Tempo\SearchTableType $tableType;

    /**
     * For mixed data sources the selected datasource is on the query level.
     * For non mixed scenarios this is undefined.
     * TODO find a better way to do this ^ that's friendly to schema
     * TODO this shouldn't be unknown but DataSourceRef | null
     */
    public ?\Grafana\Foundation\Dashboard\DataSourceRef $datasource;

    /**
     * For metric queries, the step size to use
     */
    public ?string $step;

    /**
     * @param string|null $refId
     * @param bool|null $hide
     * @param string|null $queryType
     * @param string|null $query
     * @param string|null $search
     * @param string|null $serviceName
     * @param string|null $spanName
     * @param string|null $minDuration
     * @param string|null $maxDuration
     * @param string|array<string>|null $serviceMapQuery
     * @param bool|null $serviceMapIncludeNamespace
     * @param int|null $limit
     * @param int|null $spss
     * @param array<\Grafana\Foundation\Tempo\TraceqlFilter>|null $filters
     * @param array<\Grafana\Foundation\Tempo\TraceqlFilter>|null $groupBy
     * @param \Grafana\Foundation\Tempo\SearchTableType|null $tableType
     * @param \Grafana\Foundation\Dashboard\DataSourceRef|null $datasource
     * @param string|null $step
     */
    public function __construct(?string $refId = null, ?bool $hide = null, ?string $queryType = null, ?string $query = null, ?string $search = null, ?string $serviceName = null, ?string $spanName = null, ?string $minDuration = null, ?string $maxDuration = null,  $serviceMapQuery = null, ?bool $serviceMapIncludeNamespace = null, ?int $limit = null, ?int $spss = null, ?array $filters = null, ?array $groupBy = null, ?\Grafana\Foundation\Tempo\SearchTableType $tableType = null, ?\Grafana\Foundation\Dashboard\DataSourceRef $datasource = null, ?string $step = null)
    {
        $this->refId = $refId ?: "";
        $this->hide = $hide;
        $this->queryType = $queryType;
        $this->query = $query;
        $this->search = $search;
        $this->serviceName = $serviceName;
        $this->spanName = $spanName;
        $this->minDuration = $minDuration;
        $this->maxDuration = $maxDuration;
        $this->serviceMapQuery = $serviceMapQuery;
        $this->serviceMapIncludeNamespace = $serviceMapIncludeNamespace;
        $this->limit = $limit;
        $this->spss = $spss;
        $this->filters = $filters ?: [];
        $this->groupBy = $groupBy;
        $this->tableType = $tableType;
        $this->datasource = $datasource;
        $this->step = $step;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{refId?: string, hide?: bool, queryType?: string, query?: string, search?: string, serviceName?: string, spanName?: string, minDuration?: string, maxDuration?: string, serviceMapQuery?: string|array<string>, serviceMapIncludeNamespace?: bool, limit?: int, spss?: int, filters?: array<mixed>, groupBy?: array<mixed>, tableType?: string, datasource?: mixed, step?: string} $inputData */
        $data = $inputData;
        return new self(
            refId: $data["refId"] ?? null,
            hide: $data["hide"] ?? null,
            queryType: $data["queryType"] ?? null,
            query: $data["query"] ?? null,
            search: $data["search"] ?? null,
            serviceName: $data["serviceName"] ?? null,
            spanName: $data["spanName"] ?? null,
            minDuration: $data["minDuration"] ?? null,
            maxDuration: $data["maxDuration"] ?? null,
            serviceMapQuery: isset($data["serviceMapQuery"]) ? (function($input) {
        switch (true) {
        case is_string($input):
            return $input;
        default:
            return $input;
    }
    })($data["serviceMapQuery"]) : null,
            serviceMapIncludeNamespace: $data["serviceMapIncludeNamespace"] ?? null,
            limit: $data["limit"] ?? null,
            spss: $data["spss"] ?? null,
            filters: array_filter(array_map((function($input) {
    	/** @var array{id?: string, tag?: string, operator?: string, value?: string|array<string>, valueType?: string, scope?: string} */
    $val = $input;
    	return \Grafana\Foundation\Tempo\TraceqlFilter::fromArray($val);
    }), $data["filters"] ?? [])),
            groupBy: array_filter(array_map((function($input) {
    	/** @var array{id?: string, tag?: string, operator?: string, value?: string|array<string>, valueType?: string, scope?: string} */
    $val = $input;
    	return \Grafana\Foundation\Tempo\TraceqlFilter::fromArray($val);
    }), $data["groupBy"] ?? [])),
            tableType: isset($data["tableType"]) ? (function($input) { return \Grafana\Foundation\Tempo\SearchTableType::fromValue($input); })($data["tableType"]) : null,
            datasource: isset($data["datasource"]) ? (function($input) {
    	/** @var array{type?: string, uid?: string} */
    $val = $input;
    	return \Grafana\Foundation\Dashboard\DataSourceRef::fromArray($val);
    })($data["datasource"]) : null,
            step: $data["step"] ?? null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
            "refId" => $this->refId,
            "filters" => $this->filters,
        ];
        if (isset($this->hide)) {
            $data["hide"] = $this->hide;
        }
        if (isset($this->queryType)) {
            $data["queryType"] = $this->queryType;
        }
        if (isset($this->query)) {
            $data["query"] = $this->query;
        }
        if (isset($this->search)) {
            $data["search"] = $this->search;
        }
        if (isset($this->serviceName)) {
            $data["serviceName"] = $this->serviceName;
        }
        if (isset($this->spanName)) {
            $data["spanName"] = $this->spanName;
        }
        if (isset($this->minDuration)) {
            $data["minDuration"] = $this->minDuration;
        }
        if (isset($this->maxDuration)) {
            $data["maxDuration"] = $this->maxDuration;
        }
        if (isset($this->serviceMapQuery)) {
            $data["serviceMapQuery"] = $this->serviceMapQuery;
        }
        if (isset($this->serviceMapIncludeNamespace)) {
            $data["serviceMapIncludeNamespace"] = $this->serviceMapIncludeNamespace;
        }
        if (isset($this->limit)) {
            $data["limit"] = $this->limit;
        }
        if (isset($this->spss)) {
            $data["spss"] = $this->spss;
        }
        if (isset($this->groupBy)) {
            $data["groupBy"] = $this->groupBy;
        }
        if (isset($this->tableType)) {
            $data["tableType"] = $this->tableType;
        }
        if (isset($this->datasource)) {
            $data["datasource"] = $this->datasource;
        }
        if (isset($this->step)) {
            $data["step"] = $this->step;
        }
        return $data;
    }

    public function dataqueryType(): string
    {
        return "tempo";
    }
}
