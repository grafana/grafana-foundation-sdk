<?php

namespace Grafana\Foundation\Dashboardv2beta1;

class QueryGroupSpec implements \JsonSerializable
{
    /**
     * @var array<\Grafana\Foundation\Dashboardv2beta1\PanelQueryKind>
     */
    public array $queries;

    /**
     * @var array<\Grafana\Foundation\Dashboardv2beta1\TransformationKind>
     */
    public array $transformations;

    public \Grafana\Foundation\Dashboardv2beta1\QueryOptionsSpec $queryOptions;

    /**
     * @param array<\Grafana\Foundation\Dashboardv2beta1\PanelQueryKind>|null $queries
     * @param array<\Grafana\Foundation\Dashboardv2beta1\TransformationKind>|null $transformations
     * @param \Grafana\Foundation\Dashboardv2beta1\QueryOptionsSpec|null $queryOptions
     */
    public function __construct(?array $queries = null, ?array $transformations = null, ?\Grafana\Foundation\Dashboardv2beta1\QueryOptionsSpec $queryOptions = null)
    {
        $this->queries = $queries ?: [];
        $this->transformations = $transformations ?: [];
        $this->queryOptions = $queryOptions ?: new \Grafana\Foundation\Dashboardv2beta1\QueryOptionsSpec();
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{queries?: array<mixed>, transformations?: array<mixed>, queryOptions?: mixed} $inputData */
        $data = $inputData;
        return new self(
            queries: array_filter(array_map((function($input) {
    	/** @var array{kind?: string, spec?: mixed} */
    $val = $input;
    	return \Grafana\Foundation\Dashboardv2beta1\PanelQueryKind::fromArray($val);
    }), $data["queries"] ?? [])),
            transformations: array_filter(array_map((function($input) {
    	/** @var array{kind?: string, spec?: mixed} */
    $val = $input;
    	return \Grafana\Foundation\Dashboardv2beta1\TransformationKind::fromArray($val);
    }), $data["transformations"] ?? [])),
            queryOptions: isset($data["queryOptions"]) ? (function($input) {
    	/** @var array{timeFrom?: string, maxDataPoints?: int, timeShift?: string, queryCachingTTL?: int, interval?: string, cacheTimeout?: string, hideTimeOverride?: bool, timeCompare?: string} */
    $val = $input;
    	return \Grafana\Foundation\Dashboardv2beta1\QueryOptionsSpec::fromArray($val);
    })($data["queryOptions"]) : null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->queries = $this->queries;
        $data->transformations = $this->transformations;
        $data->queryOptions = $this->queryOptions;
        return $data;
    }
}
