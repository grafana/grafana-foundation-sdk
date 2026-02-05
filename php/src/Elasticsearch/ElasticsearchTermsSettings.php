<?php

namespace Grafana\Foundation\Elasticsearch;

class ElasticsearchTermsSettings implements \JsonSerializable
{
    public ?\Grafana\Foundation\Elasticsearch\TermsOrder $order;

    public ?string $size;

    public ?string $minDocCount;

    public ?string $orderBy;

    public ?string $missing;

    /**
     * @param \Grafana\Foundation\Elasticsearch\TermsOrder|null $order
     * @param string|null $size
     * @param string|null $minDocCount
     * @param string|null $orderBy
     * @param string|null $missing
     */
    public function __construct(?\Grafana\Foundation\Elasticsearch\TermsOrder $order = null, ?string $size = null, ?string $minDocCount = null, ?string $orderBy = null, ?string $missing = null)
    {
        $this->order = $order;
        $this->size = $size;
        $this->minDocCount = $minDocCount;
        $this->orderBy = $orderBy;
        $this->missing = $missing;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{order?: string, size?: string, min_doc_count?: string, orderBy?: string, missing?: string} $inputData */
        $data = $inputData;
        return new self(
            order: isset($data["order"]) ? (function($input) { return \Grafana\Foundation\Elasticsearch\TermsOrder::fromValue($input); })($data["order"]) : null,
            size: $data["size"] ?? null,
            minDocCount: $data["min_doc_count"] ?? null,
            orderBy: $data["orderBy"] ?? null,
            missing: $data["missing"] ?? null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        if (isset($this->order)) {
            $data->order = $this->order;
        }
        if (isset($this->size)) {
            $data->size = $this->size;
        }
        if (isset($this->minDocCount)) {
            $data->min_doc_count = $this->minDocCount;
        }
        if (isset($this->orderBy)) {
            $data->orderBy = $this->orderBy;
        }
        if (isset($this->missing)) {
            $data->missing = $this->missing;
        }
        return $data;
    }
}
