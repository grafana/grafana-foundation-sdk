<?php

namespace Grafana\Foundation\Prometheus;

class Scopes implements \JsonSerializable
{
    /**
     * @var array<string>|null
     */
    public ?array $defaultPath;

    /**
     * @var array<\Grafana\Foundation\Prometheus\ScopesFilters>|null
     */
    public ?array $filters;

    public string $title;

    /**
     * @param array<string>|null $defaultPath
     * @param array<\Grafana\Foundation\Prometheus\ScopesFilters>|null $filters
     * @param string|null $title
     */
    public function __construct(?array $defaultPath = null, ?array $filters = null, ?string $title = null)
    {
        $this->defaultPath = $defaultPath;
        $this->filters = $filters;
        $this->title = $title ?: "";
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{defaultPath?: array<string>, filters?: array<mixed>, title?: string} $inputData */
        $data = $inputData;
        return new self(
            defaultPath: $data["defaultPath"] ?? null,
            filters: array_filter(array_map((function($input) {
    	/** @var array{key?: string, operator?: string, value?: string, values?: array<string>} */
    $val = $input;
    	return \Grafana\Foundation\Prometheus\ScopesFilters::fromArray($val);
    }), $data["filters"] ?? [])),
            title: $data["title"] ?? null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->title = $this->title;
        if (isset($this->defaultPath)) {
            $data->defaultPath = $this->defaultPath;
        }
        if (isset($this->filters)) {
            $data->filters = $this->filters;
        }
        return $data;
    }
}
