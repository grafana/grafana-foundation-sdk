<?php

namespace Grafana\Foundation\Dashboard;

/**
 * TODO: this should be a regular DataQuery that depends on the selected dashboard
 * these match the properties of the "grafana" datasouce that is default in most dashboards
 */
class AnnotationTarget implements \JsonSerializable
{
    /**
     * Only required/valid for the grafana datasource...
     * but code+tests is already depending on it so hard to change
     */
    public int $limit;

    /**
     * Only required/valid for the grafana datasource...
     * but code+tests is already depending on it so hard to change
     */
    public bool $matchAny;

    /**
     * Only required/valid for the grafana datasource...
     * but code+tests is already depending on it so hard to change
     * @var array<string>
     */
    public array $tags;

    /**
     * Only required/valid for the grafana datasource...
     * but code+tests is already depending on it so hard to change
     */
    public string $type;

    /**
     * @param int|null $limit
     * @param bool|null $matchAny
     * @param array<string>|null $tags
     * @param string|null $type
     */
    public function __construct(?int $limit = null, ?bool $matchAny = null, ?array $tags = null, ?string $type = null)
    {
        $this->limit = $limit ?: 0;
        $this->matchAny = $matchAny ?: false;
        $this->tags = $tags ?: [];
        $this->type = $type ?: "";
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{limit?: int, matchAny?: bool, tags?: array<string>, type?: string} $inputData */
        $data = $inputData;
        return new self(
            limit: $data["limit"] ?? null,
            matchAny: $data["matchAny"] ?? null,
            tags: $data["tags"] ?? null,
            type: $data["type"] ?? null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
            "limit" => $this->limit,
            "matchAny" => $this->matchAny,
            "tags" => $this->tags,
            "type" => $this->type,
        ];
        return $data;
    }
}
