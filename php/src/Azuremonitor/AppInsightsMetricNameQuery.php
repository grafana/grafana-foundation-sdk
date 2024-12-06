<?php

namespace Grafana\Foundation\Azuremonitor;

class AppInsightsMetricNameQuery implements \JsonSerializable
{
    public ?string $rawQuery;

    public string $kind;

    /**
     * @param string|null $rawQuery
     */
    public function __construct(?string $rawQuery = null)
    {
        $this->rawQuery = $rawQuery;
        $this->kind = "AppInsightsMetricNameQuery";
    
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{rawQuery?: string, kind?: string} $inputData */
        $data = $inputData;
        return new self(
            rawQuery: $data["rawQuery"] ?? null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
            "kind" => $this->kind,
        ];
        if (isset($this->rawQuery)) {
            $data["rawQuery"] = $this->rawQuery;
        }
        return $data;
    }
}
