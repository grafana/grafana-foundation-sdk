<?php

namespace Grafana\Foundation\Azuremonitor;

class AppInsightsGroupByQuery implements \JsonSerializable
{
    public ?string $rawQuery;

    public string $kind;

    public string $metricName;

    /**
     * @param string|null $rawQuery
     * @param string|null $metricName
     */
    public function __construct(?string $rawQuery = null, ?string $metricName = null)
    {
        $this->rawQuery = $rawQuery;
        $this->kind = "AppInsightsGroupByQuery";
    
        $this->metricName = $metricName ?: "";
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{rawQuery?: string, kind?: string, metricName?: string} $inputData */
        $data = $inputData;
        return new self(
            rawQuery: $data["rawQuery"] ?? null,
            metricName: $data["metricName"] ?? null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->kind = $this->kind;
        $data->metricName = $this->metricName;
        if (isset($this->rawQuery)) {
            $data->rawQuery = $this->rawQuery;
        }
        return $data;
    }
}
