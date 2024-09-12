<?php

namespace Grafana\Foundation\Azuremonitor;

class BaseGrafanaTemplateVariableQuery implements \JsonSerializable
{
    public ?string $rawQuery;

    /**
     * @param string|null $rawQuery
     */
    public function __construct(?string $rawQuery = null)
    {
        $this->rawQuery = $rawQuery;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{rawQuery?: string} $inputData */
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
        ];
        if (isset($this->rawQuery)) {
            $data["rawQuery"] = $this->rawQuery;
        }
        return $data;
    }
}
