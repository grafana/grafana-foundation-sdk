<?php

namespace Grafana\Foundation\Elasticsearch;

class ElasticsearchInlineScript implements \JsonSerializable
{
    public ?string $inline;

    /**
     * @param string|null $inline
     */
    public function __construct(?string $inline = null)
    {
        $this->inline = $inline;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{inline?: string} $inputData */
        $data = $inputData;
        return new self(
            inline: $data["inline"] ?? null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
        ];
        if (isset($this->inline)) {
            $data["inline"] = $this->inline;
        }
        return $data;
    }
}
