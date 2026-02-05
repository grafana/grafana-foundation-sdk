<?php

namespace Grafana\Foundation\Azuremonitor;

class AzureResourceGraphQuery implements \JsonSerializable
{
    /**
     * Azure Resource Graph KQL query to be executed.
     */
    public ?string $query;

    /**
     * Specifies the format results should be returned as. Defaults to table.
     */
    public ?string $resultFormat;

    /**
     * @param string|null $query
     * @param string|null $resultFormat
     */
    public function __construct(?string $query = null, ?string $resultFormat = null)
    {
        $this->query = $query;
        $this->resultFormat = $resultFormat;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{query?: string, resultFormat?: string} $inputData */
        $data = $inputData;
        return new self(
            query: $data["query"] ?? null,
            resultFormat: $data["resultFormat"] ?? null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        if (isset($this->query)) {
            $data->query = $this->query;
        }
        if (isset($this->resultFormat)) {
            $data->resultFormat = $this->resultFormat;
        }
        return $data;
    }
}
