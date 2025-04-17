<?php

namespace Grafana\Foundation\Elasticsearch;

class ElasticsearchRawDocumentSettings implements \JsonSerializable
{
    public ?string $size;

    /**
     * @param string|null $size
     */
    public function __construct(?string $size = null)
    {
        $this->size = $size;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{size?: string} $inputData */
        $data = $inputData;
        return new self(
            size: $data["size"] ?? null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        if (isset($this->size)) {
            $data->size = $this->size;
        }
        return $data;
    }
}
