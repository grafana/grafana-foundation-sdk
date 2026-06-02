<?php

namespace Grafana\Foundation\Dashboard;

/**
 * TODO docs
 * FROM: AnnotationQuery in grafana-data/src/types/annotations.ts
 */
class AnnotationQuery implements \JsonSerializable
{
    /**
     * Name of annotation.
     */
    public string $name;

    /**
     * Datasource where the annotations data is
     */
    public ?\Grafana\Foundation\Common\DataSourceRef $datasource;

    /**
     * When enabled the annotation query is issued with every dashboard refresh
     */
    public bool $enable;

    /**
     * Annotation queries can be toggled on or off at the top of the dashboard.
     * When hide is true, the toggle is not shown in the dashboard.
     */
    public ?bool $hide;

    /**
     * Color to use for the annotation event markers
     */
    public string $iconColor;

    /**
     * Filters to apply when fetching annotations
     */
    public ?\Grafana\Foundation\Dashboard\AnnotationPanelFilter $filter;

    /**
     * TODO.. this should just be a normal query target
     * @var \Grafana\Foundation\Cog\Dataquery|null
     */
    public ?\Grafana\Foundation\Cog\Dataquery $target;

    /**
     * TODO -- this should not exist here, it is based on the --grafana-- datasource
     */
    public ?string $type;

    /**
     * Set to 1 for the standard annotation query all dashboards have by default.
     */
    public ?float $builtIn;

    /**
     * Placement can be used to display the annotation query somewhere else on the dashboard other than the default location.
     */
    public ?\Grafana\Foundation\Dashboard\AnnotationQueryPlacement $placement;

    public ?string $expr;

    /**
     * Format for Prometheus annotation text. Label values can be interpolated with templates like {{instance}}.
     */
    public ?string $textFormat;

    /**
     * Format for Prometheus and Loki annotation titles. Label values can be interpolated with templates like {{instance}}.
     */
    public ?string $titleFormat;

    /**
     * Comma-separated label keys used as annotation tags.
     */
    public ?string $tagKeys;

    /**
     * Legacy Prometheus annotation query step interval.
     */
    public ?string $step;

    /**
     * Use the Prometheus series value as the annotation timestamp.
     */
    public ?bool $useValueForTime;

    /**
     * Mappings define how to convert data frame fields to annotation event fields.
     * @var array<string, \Grafana\Foundation\Dashboard\AnnotationEventFieldMapping>|null
     */
    public ?array $mappings;

    /**
     * @param string|null $name
     * @param \Grafana\Foundation\Common\DataSourceRef|null $datasource
     * @param bool|null $enable
     * @param bool|null $hide
     * @param string|null $iconColor
     * @param \Grafana\Foundation\Dashboard\AnnotationPanelFilter|null $filter
     * @param \Grafana\Foundation\Cog\Dataquery|null $target
     * @param string|null $type
     * @param float|null $builtIn
     * @param string|null $expr
     * @param string|null $textFormat
     * @param string|null $titleFormat
     * @param string|null $tagKeys
     * @param string|null $step
     * @param bool|null $useValueForTime
     * @param array<string, \Grafana\Foundation\Dashboard\AnnotationEventFieldMapping>|null $mappings
     */
    public function __construct(?string $name = null, ?\Grafana\Foundation\Common\DataSourceRef $datasource = null, ?bool $enable = null, ?bool $hide = null, ?string $iconColor = null, ?\Grafana\Foundation\Dashboard\AnnotationPanelFilter $filter = null, ?\Grafana\Foundation\Cog\Dataquery $target = null, ?string $type = null, ?float $builtIn = null, ?string $expr = null, ?string $textFormat = null, ?string $titleFormat = null, ?string $tagKeys = null, ?string $step = null, ?bool $useValueForTime = null, ?array $mappings = null)
    {
        $this->name = $name ?: "";
        $this->datasource = $datasource;
        $this->enable = $enable ?: true;
        $this->hide = $hide;
        $this->iconColor = $iconColor ?: "";
        $this->filter = $filter;
        $this->target = $target;
        $this->type = $type;
        $this->builtIn = $builtIn;
        $this->placement = \Grafana\Foundation\Dashboard\AnnotationQueryPlacement::inControlsMenu();
        $this->expr = $expr;
        $this->textFormat = $textFormat;
        $this->titleFormat = $titleFormat;
        $this->tagKeys = $tagKeys;
        $this->step = $step;
        $this->useValueForTime = $useValueForTime;
        $this->mappings = $mappings;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{name?: string, datasource?: mixed, enable?: bool, hide?: bool, iconColor?: string, filter?: mixed, target?: mixed, type?: string, builtIn?: float, placement?: "inControlsMenu", expr?: string, textFormat?: string, titleFormat?: string, tagKeys?: string, step?: string, useValueForTime?: bool, mappings?: array<string, mixed>} $inputData */
        $data = $inputData;
        return new self(
            name: $data["name"] ?? null,
            datasource: isset($data["datasource"]) ? (function($input) {
    	/** @var array{type?: string, uid?: string} */
    $val = $input;
    	return \Grafana\Foundation\Common\DataSourceRef::fromArray($val);
    })($data["datasource"]) : null,
            enable: $data["enable"] ?? null,
            hide: $data["hide"] ?? null,
            iconColor: $data["iconColor"] ?? null,
            filter: isset($data["filter"]) ? (function($input) {
    	/** @var array{exclude?: bool, ids?: array<int>} */
    $val = $input;
    	return \Grafana\Foundation\Dashboard\AnnotationPanelFilter::fromArray($val);
    })($data["filter"]) : null,
            target: isset($data["target"]) ? (function ($in) {
    	/** @var array{datasource?: array{type?: mixed}} $in */
        $hint = (isset($in["datasource"], $in["datasource"]["type"]) && is_string($in["datasource"]["type"])) ? $in["datasource"]["type"] : "";
    
        /** @var array<string, mixed> $in */
        return \Grafana\Foundation\Cog\Runtime::get()->dataqueryFromArray($in, $hint);
    })($data["target"]) : null,
            type: $data["type"] ?? null,
            builtIn: $data["builtIn"] ?? null,
            expr: $data["expr"] ?? null,
            textFormat: $data["textFormat"] ?? null,
            titleFormat: $data["titleFormat"] ?? null,
            tagKeys: $data["tagKeys"] ?? null,
            step: $data["step"] ?? null,
            useValueForTime: $data["useValueForTime"] ?? null,
            mappings: isset($data["mappings"]) ? (function($input) {
        /** @var array<string, \Grafana\Foundation\Dashboard\AnnotationEventFieldMapping> $results */
        $results = [];
        foreach ($input as $key => $val) {
            $results[$key] = isset($val) ? (function($input) {
    	/** @var array{source?: string, value?: string, regex?: string} */
    $val = $input;
    	return \Grafana\Foundation\Dashboard\AnnotationEventFieldMapping::fromArray($val);
    })($val) : null;
        }
        return array_filter($results);
    })($data["mappings"]) : null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->name = $this->name;
        $data->enable = $this->enable;
        $data->iconColor = $this->iconColor;
        if (isset($this->datasource)) {
            $data->datasource = $this->datasource;
        }
        if (isset($this->hide)) {
            $data->hide = $this->hide;
        }
        if (isset($this->filter)) {
            $data->filter = $this->filter;
        }
        if (isset($this->target)) {
            $data->target = $this->target;
        }
        if (isset($this->type)) {
            $data->type = $this->type;
        }
        if (isset($this->builtIn)) {
            $data->builtIn = $this->builtIn;
        }
        if (isset($this->placement)) {
            $data->placement = $this->placement;
        }
        if (isset($this->expr)) {
            $data->expr = $this->expr;
        }
        if (isset($this->textFormat)) {
            $data->textFormat = $this->textFormat;
        }
        if (isset($this->titleFormat)) {
            $data->titleFormat = $this->titleFormat;
        }
        if (isset($this->tagKeys)) {
            $data->tagKeys = $this->tagKeys;
        }
        if (isset($this->step)) {
            $data->step = $this->step;
        }
        if (isset($this->useValueForTime)) {
            $data->useValueForTime = $this->useValueForTime;
        }
        if (isset($this->mappings)) {
            $data->mappings = $this->mappings;
        }
        return $data;
    }
}
