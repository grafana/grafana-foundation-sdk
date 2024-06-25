<?php

namespace Grafana\Foundation\Googlecloudmonitoring;

/**
 * @deprecated This type is for migration purposes only. Replaced by TimeSeriesList Metric sub-query properties.
 */
class MetricQuery implements \JsonSerializable
{
    /**
     * GCP project to execute the query against.
     */
    public string $projectName;

    /**
     * Alignment function to be used. Defaults to ALIGN_MEAN.
     */
    public ?string $perSeriesAligner;

    /**
     * Alignment period to use when regularizing data. Defaults to cloud-monitoring-auto.
     */
    public ?string $alignmentPeriod;

    /**
     * Aliases can be set to modify the legend labels. e.g. {{metric.label.xxx}}. See docs for more detail.
     */
    public ?string $aliasBy;

    public string $editorMode;

    public string $metricType;

    /**
     * Reducer applied across a set of time-series values. Defaults to REDUCE_NONE.
     */
    public string $crossSeriesReducer;

    /**
     * Array of labels to group data by.
     * @var array<string>|null
     */
    public ?array $groupBys;

    /**
     * Array of filters to query data by. Labels that can be filtered on are defined by the metric.
     * @var array<string>|null
     */
    public ?array $filters;

    public ?\Grafana\Foundation\Googlecloudmonitoring\MetricKind $metricKind;

    public ?string $valueType;

    public ?string $view;

    /**
     * MQL query to be executed.
     */
    public string $query;

    /**
     * Preprocessor is not part of the API, but is used to store the preprocessor and not affect the UI for the rest of parameters
     */
    public ?\Grafana\Foundation\Googlecloudmonitoring\PreprocessorType $preprocessor;

    /**
     * To disable the graphPeriod, it should explictly be set to 'disabled'.
     */
    public ?string $graphPeriod;

    /**
     * @param string|null $projectName
     * @param string|null $perSeriesAligner
     * @param string|null $alignmentPeriod
     * @param string|null $aliasBy
     * @param string|null $editorMode
     * @param string|null $metricType
     * @param string|null $crossSeriesReducer
     * @param array<string>|null $groupBys
     * @param array<string>|null $filters
     * @param \Grafana\Foundation\Googlecloudmonitoring\MetricKind|null $metricKind
     * @param string|null $valueType
     * @param string|null $view
     * @param string|null $query
     * @param \Grafana\Foundation\Googlecloudmonitoring\PreprocessorType|null $preprocessor
     * @param string|null $graphPeriod
     */
    public function __construct(?string $projectName = null, ?string $perSeriesAligner = null, ?string $alignmentPeriod = null, ?string $aliasBy = null, ?string $editorMode = null, ?string $metricType = null, ?string $crossSeriesReducer = null, ?array $groupBys = null, ?array $filters = null, ?\Grafana\Foundation\Googlecloudmonitoring\MetricKind $metricKind = null, ?string $valueType = null, ?string $view = null, ?string $query = null, ?\Grafana\Foundation\Googlecloudmonitoring\PreprocessorType $preprocessor = null, ?string $graphPeriod = null)
    {
        $this->projectName = $projectName ?: "";
        $this->perSeriesAligner = $perSeriesAligner;
        $this->alignmentPeriod = $alignmentPeriod;
        $this->aliasBy = $aliasBy;
        $this->editorMode = $editorMode ?: "";
        $this->metricType = $metricType ?: "";
        $this->crossSeriesReducer = $crossSeriesReducer ?: "";
        $this->groupBys = $groupBys;
        $this->filters = $filters;
        $this->metricKind = $metricKind;
        $this->valueType = $valueType;
        $this->view = $view;
        $this->query = $query ?: "";
        $this->preprocessor = $preprocessor;
        $this->graphPeriod = $graphPeriod;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{projectName?: string, perSeriesAligner?: string, alignmentPeriod?: string, aliasBy?: string, editorMode?: string, metricType?: string, crossSeriesReducer?: string, groupBys?: array<string>, filters?: array<string>, metricKind?: string, valueType?: string, view?: string, query?: string, preprocessor?: string, graphPeriod?: string} $inputData */
        $data = $inputData;
        return new self(
            projectName: $data["projectName"] ?? null,
            perSeriesAligner: $data["perSeriesAligner"] ?? null,
            alignmentPeriod: $data["alignmentPeriod"] ?? null,
            aliasBy: $data["aliasBy"] ?? null,
            editorMode: $data["editorMode"] ?? null,
            metricType: $data["metricType"] ?? null,
            crossSeriesReducer: $data["crossSeriesReducer"] ?? null,
            groupBys: $data["groupBys"] ?? null,
            filters: $data["filters"] ?? null,
            metricKind: isset($data["metricKind"]) ? (function($input) { return \Grafana\Foundation\Googlecloudmonitoring\MetricKind::fromValue($input); })($data["metricKind"]) : null,
            valueType: $data["valueType"] ?? null,
            view: $data["view"] ?? null,
            query: $data["query"] ?? null,
            preprocessor: isset($data["preprocessor"]) ? (function($input) { return \Grafana\Foundation\Googlecloudmonitoring\PreprocessorType::fromValue($input); })($data["preprocessor"]) : null,
            graphPeriod: $data["graphPeriod"] ?? null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
            "projectName" => $this->projectName,
            "editorMode" => $this->editorMode,
            "metricType" => $this->metricType,
            "crossSeriesReducer" => $this->crossSeriesReducer,
            "query" => $this->query,
        ];
        if (isset($this->perSeriesAligner)) {
            $data["perSeriesAligner"] = $this->perSeriesAligner;
        }
        if (isset($this->alignmentPeriod)) {
            $data["alignmentPeriod"] = $this->alignmentPeriod;
        }
        if (isset($this->aliasBy)) {
            $data["aliasBy"] = $this->aliasBy;
        }
        if (isset($this->groupBys)) {
            $data["groupBys"] = $this->groupBys;
        }
        if (isset($this->filters)) {
            $data["filters"] = $this->filters;
        }
        if (isset($this->metricKind)) {
            $data["metricKind"] = $this->metricKind;
        }
        if (isset($this->valueType)) {
            $data["valueType"] = $this->valueType;
        }
        if (isset($this->view)) {
            $data["view"] = $this->view;
        }
        if (isset($this->preprocessor)) {
            $data["preprocessor"] = $this->preprocessor;
        }
        if (isset($this->graphPeriod)) {
            $data["graphPeriod"] = $this->graphPeriod;
        }
        return $data;
    }
}
