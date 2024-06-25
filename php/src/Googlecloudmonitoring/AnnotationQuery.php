<?php

namespace Grafana\Foundation\Googlecloudmonitoring;

/**
 * Annotation sub-query properties.
 */
class AnnotationQuery implements \JsonSerializable
{
    /**
     * GCP project to execute the query against.
     */
    public string $projectName;

    /**
     * Reducer applied across a set of time-series values. Defaults to REDUCE_NONE.
     */
    public string $crossSeriesReducer;

    /**
     * Alignment period to use when regularizing data. Defaults to cloud-monitoring-auto.
     */
    public ?string $alignmentPeriod;

    /**
     * Alignment function to be used. Defaults to ALIGN_MEAN.
     */
    public ?string $perSeriesAligner;

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

    /**
     * Data view, defaults to FULL.
     */
    public ?string $view;

    /**
     * Only present if a preprocessor is selected. Reducer applied across a set of time-series values. Defaults to REDUCE_NONE.
     */
    public ?string $secondaryCrossSeriesReducer;

    /**
     * Only present if a preprocessor is selected. Alignment period to use when regularizing data. Defaults to cloud-monitoring-auto.
     */
    public ?string $secondaryAlignmentPeriod;

    /**
     * Only present if a preprocessor is selected. Alignment function to be used. Defaults to ALIGN_MEAN.
     */
    public ?string $secondaryPerSeriesAligner;

    /**
     * Only present if a preprocessor is selected. Array of labels to group data by.
     * @var array<string>|null
     */
    public ?array $secondaryGroupBys;

    /**
     * Annotation title.
     */
    public ?string $title;

    /**
     * Preprocessor is not part of the API, but is used to store the preprocessor and not affect the UI for the rest of parameters
     */
    public ?\Grafana\Foundation\Googlecloudmonitoring\PreprocessorType $preprocessor;

    /**
     * Annotation text.
     */
    public ?string $text;

    /**
     * @param string|null $projectName
     * @param string|null $crossSeriesReducer
     * @param string|null $alignmentPeriod
     * @param string|null $perSeriesAligner
     * @param array<string>|null $groupBys
     * @param array<string>|null $filters
     * @param string|null $view
     * @param string|null $secondaryCrossSeriesReducer
     * @param string|null $secondaryAlignmentPeriod
     * @param string|null $secondaryPerSeriesAligner
     * @param array<string>|null $secondaryGroupBys
     * @param string|null $title
     * @param \Grafana\Foundation\Googlecloudmonitoring\PreprocessorType|null $preprocessor
     * @param string|null $text
     */
    public function __construct(?string $projectName = null, ?string $crossSeriesReducer = null, ?string $alignmentPeriod = null, ?string $perSeriesAligner = null, ?array $groupBys = null, ?array $filters = null, ?string $view = null, ?string $secondaryCrossSeriesReducer = null, ?string $secondaryAlignmentPeriod = null, ?string $secondaryPerSeriesAligner = null, ?array $secondaryGroupBys = null, ?string $title = null, ?\Grafana\Foundation\Googlecloudmonitoring\PreprocessorType $preprocessor = null, ?string $text = null)
    {
        $this->projectName = $projectName ?: "";
        $this->crossSeriesReducer = $crossSeriesReducer ?: "";
        $this->alignmentPeriod = $alignmentPeriod;
        $this->perSeriesAligner = $perSeriesAligner;
        $this->groupBys = $groupBys;
        $this->filters = $filters;
        $this->view = $view;
        $this->secondaryCrossSeriesReducer = $secondaryCrossSeriesReducer;
        $this->secondaryAlignmentPeriod = $secondaryAlignmentPeriod;
        $this->secondaryPerSeriesAligner = $secondaryPerSeriesAligner;
        $this->secondaryGroupBys = $secondaryGroupBys;
        $this->title = $title;
        $this->preprocessor = $preprocessor;
        $this->text = $text;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{projectName?: string, crossSeriesReducer?: string, alignmentPeriod?: string, perSeriesAligner?: string, groupBys?: array<string>, filters?: array<string>, view?: string, secondaryCrossSeriesReducer?: string, secondaryAlignmentPeriod?: string, secondaryPerSeriesAligner?: string, secondaryGroupBys?: array<string>, title?: string, preprocessor?: string, text?: string} $inputData */
        $data = $inputData;
        return new self(
            projectName: $data["projectName"] ?? null,
            crossSeriesReducer: $data["crossSeriesReducer"] ?? null,
            alignmentPeriod: $data["alignmentPeriod"] ?? null,
            perSeriesAligner: $data["perSeriesAligner"] ?? null,
            groupBys: $data["groupBys"] ?? null,
            filters: $data["filters"] ?? null,
            view: $data["view"] ?? null,
            secondaryCrossSeriesReducer: $data["secondaryCrossSeriesReducer"] ?? null,
            secondaryAlignmentPeriod: $data["secondaryAlignmentPeriod"] ?? null,
            secondaryPerSeriesAligner: $data["secondaryPerSeriesAligner"] ?? null,
            secondaryGroupBys: $data["secondaryGroupBys"] ?? null,
            title: $data["title"] ?? null,
            preprocessor: isset($data["preprocessor"]) ? (function($input) { return \Grafana\Foundation\Googlecloudmonitoring\PreprocessorType::fromValue($input); })($data["preprocessor"]) : null,
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
            "crossSeriesReducer" => $this->crossSeriesReducer,
        ];
        if (isset($this->alignmentPeriod)) {
            $data["alignmentPeriod"] = $this->alignmentPeriod;
        }
        if (isset($this->perSeriesAligner)) {
            $data["perSeriesAligner"] = $this->perSeriesAligner;
        }
        if (isset($this->groupBys)) {
            $data["groupBys"] = $this->groupBys;
        }
        if (isset($this->filters)) {
            $data["filters"] = $this->filters;
        }
        if (isset($this->view)) {
            $data["view"] = $this->view;
        }
        if (isset($this->secondaryCrossSeriesReducer)) {
            $data["secondaryCrossSeriesReducer"] = $this->secondaryCrossSeriesReducer;
        }
        if (isset($this->secondaryAlignmentPeriod)) {
            $data["secondaryAlignmentPeriod"] = $this->secondaryAlignmentPeriod;
        }
        if (isset($this->secondaryPerSeriesAligner)) {
            $data["secondaryPerSeriesAligner"] = $this->secondaryPerSeriesAligner;
        }
        if (isset($this->secondaryGroupBys)) {
            $data["secondaryGroupBys"] = $this->secondaryGroupBys;
        }
        if (isset($this->title)) {
            $data["title"] = $this->title;
        }
        if (isset($this->preprocessor)) {
            $data["preprocessor"] = $this->preprocessor;
        }
        if (isset($this->text)) {
            $data["text"] = $this->text;
        }
        return $data;
    }
}
