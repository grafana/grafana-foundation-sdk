<?php

namespace Grafana\Foundation\Elasticsearch;

class ElasticsearchFiltersSettings implements \JsonSerializable
{
    /**
     * @var array<\Grafana\Foundation\Elasticsearch\Filter>|null
     */
    public ?array $filters;

    /**
     * @param array<\Grafana\Foundation\Elasticsearch\Filter>|null $filters
     */
    public function __construct(?array $filters = null)
    {
        $this->filters = $filters;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{filters?: array<mixed>} $inputData */
        $data = $inputData;
        return new self(
            filters: array_filter(array_map((function($input) {
    	/** @var array{query?: string, label?: string} */
    $val = $input;
    	return \Grafana\Foundation\Elasticsearch\Filter::fromArray($val);
    }), $data["filters"] ?? [])),
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        if (isset($this->filters)) {
            $data->filters = $this->filters;
        }
        return $data;
    }
}
