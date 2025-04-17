<?php

namespace Grafana\Foundation\Googlecloudmonitoring;

/**
 * Time Series sub-query properties.
 */
class TimeSeriesQuery implements \JsonSerializable
{
    /**
     * GCP project to execute the query against.
     */
    public string $projectName;

    /**
     * MQL query to be executed.
     */
    public string $query;

    /**
     * To disable the graphPeriod, it should explictly be set to 'disabled'.
     */
    public ?string $graphPeriod;

    /**
     * @param string|null $projectName
     * @param string|null $query
     * @param string|null $graphPeriod
     */
    public function __construct(?string $projectName = null, ?string $query = null, ?string $graphPeriod = null)
    {
        $this->projectName = $projectName ?: "";
        $this->query = $query ?: "";
        $this->graphPeriod = $graphPeriod;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{projectName?: string, query?: string, graphPeriod?: string} $inputData */
        $data = $inputData;
        return new self(
            projectName: $data["projectName"] ?? null,
            query: $data["query"] ?? null,
            graphPeriod: $data["graphPeriod"] ?? null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->projectName = $this->projectName;
        $data->query = $this->query;
        if (isset($this->graphPeriod)) {
            $data->graphPeriod = $this->graphPeriod;
        }
        return $data;
    }
}
