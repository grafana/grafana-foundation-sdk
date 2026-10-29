<?php

namespace Grafana\Foundation\Alertingv0alpha1;

/**
 * #SearchResultHit is a single match: its identity, an optional relevance
 * score (present only when the query included free text), and the requested
 * fields.
 */
class SearchResultHit implements \JsonSerializable
{
    public \Grafana\Foundation\Alertingv0alpha1\SearchResultResource $resource;

    public ?float $score;

    public \Grafana\Foundation\Alertingv0alpha1\RuleSearchHitFields $fields;

    /**
     * @param \Grafana\Foundation\Alertingv0alpha1\SearchResultResource|null $resource
     * @param float|null $score
     * @param \Grafana\Foundation\Alertingv0alpha1\RuleSearchHitFields|null $fields
     */
    public function __construct(?\Grafana\Foundation\Alertingv0alpha1\SearchResultResource $resource = null, ?float $score = null, ?\Grafana\Foundation\Alertingv0alpha1\RuleSearchHitFields $fields = null)
    {
        $this->resource = $resource ?: new \Grafana\Foundation\Alertingv0alpha1\SearchResultResource();
        $this->score = $score;
        $this->fields = $fields ?: new \Grafana\Foundation\Alertingv0alpha1\RuleSearchHitFields();
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{resource?: mixed, score?: float, fields?: mixed} $inputData */
        $data = $inputData;
        return new self(
            resource: isset($data["resource"]) ? (function($input) {
    	/** @var array{group?: string, resource?: string, kind?: string, name?: string} */
    $val = $input;
    	return \Grafana\Foundation\Alertingv0alpha1\SearchResultResource::fromArray($val);
    })($data["resource"]) : null,
            score: $data["score"] ?? null,
            fields: isset($data["fields"]) ? (function($input) {
    	/** @var array{title?: string, folder?: string, type?: string, interval?: string, paused?: bool, labels?: array<string, string>, datasourceUIDs?: array<string>, annotations?: array<string, string>, for?: string, keepFiringFor?: string, dashboardUID?: string, panelID?: int, receiver?: string, notificationType?: string, routingTree?: string, metric?: string, targetDatasourceUID?: string} */
    $val = $input;
    	return \Grafana\Foundation\Alertingv0alpha1\RuleSearchHitFields::fromArray($val);
    })($data["fields"]) : null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->resource = $this->resource;
        $data->fields = $this->fields;
        if (isset($this->score)) {
            $data->score = $this->score;
        }
        return $data;
    }
}
