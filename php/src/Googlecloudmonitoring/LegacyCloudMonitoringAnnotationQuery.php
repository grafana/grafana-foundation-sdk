<?php

namespace Grafana\Foundation\Googlecloudmonitoring;

/**
 * @deprecated Use TimeSeriesList instead. Legacy annotation query properties for migration purposes.
 */
class LegacyCloudMonitoringAnnotationQuery implements \JsonSerializable
{
    /**
     * GCP project to execute the query against.
     */
    public string $projectName;

    public string $metricType;

    /**
     * Query refId.
     */
    public string $refId;

    /**
     * Array of filters to query data by. Labels that can be filtered on are defined by the metric.
     * @var array<string>
     */
    public array $filters;

    public \Grafana\Foundation\Googlecloudmonitoring\MetricKind $metricKind;

    public string $valueType;

    /**
     * Annotation title.
     */
    public string $title;

    /**
     * Annotation text.
     */
    public string $text;

    /**
     * @param string|null $projectName
     * @param string|null $metricType
     * @param string|null $refId
     * @param array<string>|null $filters
     * @param \Grafana\Foundation\Googlecloudmonitoring\MetricKind|null $metricKind
     * @param string|null $valueType
     * @param string|null $title
     * @param string|null $text
     */
    public function __construct(?string $projectName = null, ?string $metricType = null, ?string $refId = null, ?array $filters = null, ?\Grafana\Foundation\Googlecloudmonitoring\MetricKind $metricKind = null, ?string $valueType = null, ?string $title = null, ?string $text = null)
    {
        $this->projectName = $projectName ?: "";
        $this->metricType = $metricType ?: "";
        $this->refId = $refId ?: "";
        $this->filters = $filters ?: [];
        $this->metricKind = $metricKind ?: \Grafana\Foundation\Googlecloudmonitoring\MetricKind::METRICKINDUNSPECIFIED();
        $this->valueType = $valueType ?: "";
        $this->title = $title ?: "";
        $this->text = $text ?: "";
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{projectName?: string, metricType?: string, refId?: string, filters?: array<string>, metricKind?: string, valueType?: string, title?: string, text?: string} $inputData */
        $data = $inputData;
        return new self(
            projectName: $data["projectName"] ?? null,
            metricType: $data["metricType"] ?? null,
            refId: $data["refId"] ?? null,
            filters: $data["filters"] ?? null,
            metricKind: isset($data["metricKind"]) ? (function($input) { return \Grafana\Foundation\Googlecloudmonitoring\MetricKind::fromValue($input); })($data["metricKind"]) : null,
            valueType: $data["valueType"] ?? null,
            title: $data["title"] ?? null,
            text: $data["text"] ?? null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
            "projectName" => $this->projectName,
            "metricType" => $this->metricType,
            "refId" => $this->refId,
            "filters" => $this->filters,
            "metricKind" => $this->metricKind,
            "valueType" => $this->valueType,
            "title" => $this->title,
            "text" => $this->text,
        ];
        return $data;
    }
}
