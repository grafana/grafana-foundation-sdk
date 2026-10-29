<?php

namespace Grafana\Foundation\Alertingv0alpha1;

/**
 * #SearchResults is the response envelope, mirroring
 * search.grafana.app SearchResults.
 */
class SearchResults implements \JsonSerializable
{
    public \Grafana\Foundation\Alertingv0alpha1\SearchResultsMetadata $metadata;

    /**
     * @var array<\Grafana\Foundation\Alertingv0alpha1\SearchResultHit>
     */
    public array $items;

    /**
     * @var array<string, array<\Grafana\Foundation\Alertingv0alpha1\FacetValue>>|null
     */
    public ?array $facets;

    /**
     * @param \Grafana\Foundation\Alertingv0alpha1\SearchResultsMetadata|null $metadata
     * @param array<\Grafana\Foundation\Alertingv0alpha1\SearchResultHit>|null $items
     * @param array<string, array<\Grafana\Foundation\Alertingv0alpha1\FacetValue>>|null $facets
     */
    public function __construct(?\Grafana\Foundation\Alertingv0alpha1\SearchResultsMetadata $metadata = null, ?array $items = null, ?array $facets = null)
    {
        $this->metadata = $metadata ?: new \Grafana\Foundation\Alertingv0alpha1\SearchResultsMetadata();
        $this->items = $items ?: [];
        $this->facets = $facets;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{metadata?: mixed, items?: array<mixed>, facets?: array<string, array<mixed>>} $inputData */
        $data = $inputData;
        return new self(
            metadata: isset($data["metadata"]) ? (function($input) {
    	/** @var array{continue?: string, totalHits?: int, totalHitsRelation?: string} */
    $val = $input;
    	return \Grafana\Foundation\Alertingv0alpha1\SearchResultsMetadata::fromArray($val);
    })($data["metadata"]) : null,
            items: array_filter(array_map((function($input) {
    	/** @var array{resource?: mixed, score?: float, fields?: mixed} */
    $val = $input;
    	return \Grafana\Foundation\Alertingv0alpha1\SearchResultHit::fromArray($val);
    }), $data["items"] ?? [])),
            facets: isset($data["facets"]) ? (function($input) {
        /** @var array<string, array<\Grafana\Foundation\Alertingv0alpha1\FacetValue>> $results */
        $results = [];
        foreach ($input as $key => $val) {
            $results[$key] = array_filter(array_map((function($input) {
    	/** @var array{value?: string, count?: int} */
    $val = $input;
    	return \Grafana\Foundation\Alertingv0alpha1\FacetValue::fromArray($val);
    }), $val ?? []));
        }
        return array_filter($results);
    })($data["facets"]) : null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->metadata = $this->metadata;
        $data->items = $this->items;
        if (isset($this->facets)) {
            $data->facets = $this->facets;
        }
        return $data;
    }
}
