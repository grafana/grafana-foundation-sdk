<?php

namespace Grafana\Foundation\Elasticsearch;

class ElasticsearchCumulativeSumSettings implements \JsonSerializable
{
    public ?string $format;

    /**
     * @param string|null $format
     */
    public function __construct(?string $format = null)
    {
        $this->format = $format;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{format?: string} $inputData */
        $data = $inputData;
        return new self(
            format: $data["format"] ?? null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        if (isset($this->format)) {
            $data->format = $this->format;
        }
        return $data;
    }
}
