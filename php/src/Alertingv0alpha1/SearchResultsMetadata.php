<?php

namespace Grafana\Foundation\Alertingv0alpha1;

/**
 * #SearchResultsMetadata carries the paging token and total authorised match
 * count.
 */
class SearchResultsMetadata implements \JsonSerializable
{
    public ?string $continue;

    public ?int $totalHits;

    /**
     * Always read totalHits together with totalHitsRelation.
     */
    public ?\Grafana\Foundation\Alertingv0alpha1\TotalHitsRelation $totalHitsRelation;

    /**
     * @param string|null $continue
     * @param int|null $totalHits
     * @param \Grafana\Foundation\Alertingv0alpha1\TotalHitsRelation|null $totalHitsRelation
     */
    public function __construct(?string $continue = null, ?int $totalHits = null, ?\Grafana\Foundation\Alertingv0alpha1\TotalHitsRelation $totalHitsRelation = null)
    {
        $this->continue = $continue;
        $this->totalHits = $totalHits;
        $this->totalHitsRelation = $totalHitsRelation;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{continue?: string, totalHits?: int, totalHitsRelation?: string} $inputData */
        $data = $inputData;
        return new self(
            continue: $data["continue"] ?? null,
            totalHits: $data["totalHits"] ?? null,
            totalHitsRelation: isset($data["totalHitsRelation"]) ? (function($input) { return \Grafana\Foundation\Alertingv0alpha1\TotalHitsRelation::fromValue($input); })($data["totalHitsRelation"]) : null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        if (isset($this->continue)) {
            $data->continue = $this->continue;
        }
        if (isset($this->totalHits)) {
            $data->totalHits = $this->totalHits;
        }
        if (isset($this->totalHitsRelation)) {
            $data->totalHitsRelation = $this->totalHitsRelation;
        }
        return $data;
    }
}
