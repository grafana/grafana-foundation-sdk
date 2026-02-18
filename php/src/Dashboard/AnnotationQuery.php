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

    public ?string $expr;

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
     */
    public function __construct(?string $name = null, ?\Grafana\Foundation\Common\DataSourceRef $datasource = null, ?bool $enable = null, ?bool $hide = null, ?string $iconColor = null, ?\Grafana\Foundation\Dashboard\AnnotationPanelFilter $filter = null, ?\Grafana\Foundation\Cog\Dataquery $target = null, ?string $type = null, ?float $builtIn = null, ?string $expr = null)
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
        $this->expr = $expr;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{name?: string, datasource?: mixed, enable?: bool, hide?: bool, iconColor?: string, filter?: mixed, target?: mixed, type?: string, builtIn?: float, expr?: string} $inputData */
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
        if (isset($this->expr)) {
            $data->expr = $this->expr;
        }
        return $data;
    }
}
