<?php

namespace Grafana\Foundation\Dashboardv2beta1;

/**
 * A notebook spec is a dashboard spec with narrative content. It has a title, optional description, tags, time
 * settings, and a map of elements (panels and cells) referenced by the layout.
 */
class NotebookSpec implements \JsonSerializable
{
    public string $title;

    public ?string $description;

    /**
     * @var array<string>
     */
    public array $tags;

    public \Grafana\Foundation\Dashboardv2beta1\TimeSettingsSpec $timeSettings;

    /**
     * @var array<string, \Grafana\Foundation\Dashboardv2beta1\CellKind|\Grafana\Foundation\Dashboardv2beta1\V2PanelKind|\Grafana\Foundation\Dashboardv2beta1\LibraryPanelKind>
     */
    public array $elements;

    public \Grafana\Foundation\Dashboardv2beta1\NotebookLayoutKind $layout;

    /**
     * @param string|null $title
     * @param string|null $description
     * @param array<string>|null $tags
     * @param \Grafana\Foundation\Dashboardv2beta1\TimeSettingsSpec|null $timeSettings
     * @param array<string, \Grafana\Foundation\Dashboardv2beta1\CellKind|\Grafana\Foundation\Dashboardv2beta1\V2PanelKind|\Grafana\Foundation\Dashboardv2beta1\LibraryPanelKind>|null $elements
     * @param \Grafana\Foundation\Dashboardv2beta1\NotebookLayoutKind|null $layout
     */
    public function __construct(?string $title = null, ?string $description = null, ?array $tags = null, ?\Grafana\Foundation\Dashboardv2beta1\TimeSettingsSpec $timeSettings = null, ?array $elements = null, ?\Grafana\Foundation\Dashboardv2beta1\NotebookLayoutKind $layout = null)
    {
        $this->title = $title ?: "";
        $this->description = $description;
        $this->tags = $tags ?: [];
        $this->timeSettings = $timeSettings ?: new \Grafana\Foundation\Dashboardv2beta1\TimeSettingsSpec();
        $this->elements = $elements ?: [];
        $this->layout = $layout ?: new \Grafana\Foundation\Dashboardv2beta1\NotebookLayoutKind();
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{title?: string, description?: string, tags?: array<string>, timeSettings?: mixed, elements?: array<string, mixed|mixed|mixed>, layout?: mixed} $inputData */
        $data = $inputData;
        return new self(
            title: $data["title"] ?? null,
            description: $data["description"] ?? null,
            tags: $data["tags"] ?? null,
            timeSettings: isset($data["timeSettings"]) ? (function($input) {
    	/** @var array{timezone?: string, from?: string, to?: string, autoRefresh?: string, autoRefreshIntervals?: array<string>, quickRanges?: array<mixed>, hideTimepicker?: bool, weekStart?: string, fiscalYearStartMonth?: int, nowDelay?: string} */
    $val = $input;
    	return \Grafana\Foundation\Dashboardv2beta1\TimeSettingsSpec::fromArray($val);
    })($data["timeSettings"]) : null,
            elements: isset($data["elements"]) ? (function($input) {
        /** @var array<string, \Grafana\Foundation\Dashboardv2beta1\CellKind|\Grafana\Foundation\Dashboardv2beta1\V2PanelKind|\Grafana\Foundation\Dashboardv2beta1\LibraryPanelKind> $results */
        $results = [];
        foreach ($input as $key => $val) {
            $results[$key] = isset($val) ? (function($input) {
        \assert(is_array($input), 'expected disjunction value to be an array');
        /** @var array<string, mixed> $input */
        switch ($input["kind"]) {
        case "Cell":
            return CellKind::fromArray($input);
        case "LibraryPanel":
            return LibraryPanelKind::fromArray($input);
        case "Panel":
            return V2PanelKind::fromArray($input);
        default:
            throw new \ValueError('can not parse disjunction from array');
    }
    })($val) : null;
        }
        return array_filter($results);
    })($data["elements"]) : null,
            layout: isset($data["layout"]) ? (function($input) {
    	/** @var array{kind?: string, spec?: mixed} */
    $val = $input;
    	return \Grafana\Foundation\Dashboardv2beta1\NotebookLayoutKind::fromArray($val);
    })($data["layout"]) : null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->title = $this->title;
        $data->tags = $this->tags;
        $data->timeSettings = $this->timeSettings;
        $data->elements = $this->elements;
        $data->layout = $this->layout;
        if (isset($this->description)) {
            $data->description = $this->description;
        }
        return $data;
    }
}
