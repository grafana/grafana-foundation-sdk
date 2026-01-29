<?php

namespace Grafana\Foundation\Dashboardv2beta1;

class RowsLayoutRowSpec implements \JsonSerializable
{
    public ?string $title;

    public ?bool $collapse;

    public ?bool $hideHeader;

    public ?bool $fillScreen;

    public ?\Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingGroupKind $conditionalRendering;

    public ?\Grafana\Foundation\Dashboardv2beta1\RowRepeatOptions $repeat;

    /**
     * @var \Grafana\Foundation\Dashboardv2beta1\GridLayoutKind|\Grafana\Foundation\Dashboardv2beta1\AutoGridLayoutKind|\Grafana\Foundation\Dashboardv2beta1\TabsLayoutKind|\Grafana\Foundation\Dashboardv2beta1\RowsLayoutKind
     */
    public $layout;

    /**
     * @param string|null $title
     * @param bool|null $collapse
     * @param bool|null $hideHeader
     * @param bool|null $fillScreen
     * @param \Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingGroupKind|null $conditionalRendering
     * @param \Grafana\Foundation\Dashboardv2beta1\RowRepeatOptions|null $repeat
     * @param \Grafana\Foundation\Dashboardv2beta1\GridLayoutKind|\Grafana\Foundation\Dashboardv2beta1\AutoGridLayoutKind|\Grafana\Foundation\Dashboardv2beta1\TabsLayoutKind|\Grafana\Foundation\Dashboardv2beta1\RowsLayoutKind|null $layout
     */
    public function __construct(?string $title = null, ?bool $collapse = null, ?bool $hideHeader = null, ?bool $fillScreen = null, ?\Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingGroupKind $conditionalRendering = null, ?\Grafana\Foundation\Dashboardv2beta1\RowRepeatOptions $repeat = null,  $layout = null)
    {
        $this->title = $title;
        $this->collapse = $collapse;
        $this->hideHeader = $hideHeader;
        $this->fillScreen = $fillScreen;
        $this->conditionalRendering = $conditionalRendering;
        $this->repeat = $repeat;
        $this->layout = $layout ?: new \Grafana\Foundation\Dashboardv2beta1\GridLayoutKind();
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{title?: string, collapse?: bool, hideHeader?: bool, fillScreen?: bool, conditionalRendering?: mixed, repeat?: mixed, layout?: mixed|mixed|mixed|mixed} $inputData */
        $data = $inputData;
        return new self(
            title: $data["title"] ?? null,
            collapse: $data["collapse"] ?? null,
            hideHeader: $data["hideHeader"] ?? null,
            fillScreen: $data["fillScreen"] ?? null,
            conditionalRendering: isset($data["conditionalRendering"]) ? (function($input) {
    	/** @var array{kind?: string, spec?: mixed} */
    $val = $input;
    	return \Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingGroupKind::fromArray($val);
    })($data["conditionalRendering"]) : null,
            repeat: isset($data["repeat"]) ? (function($input) {
    	/** @var array{mode?: "variable", value?: string} */
    $val = $input;
    	return \Grafana\Foundation\Dashboardv2beta1\RowRepeatOptions::fromArray($val);
    })($data["repeat"]) : null,
            layout: isset($data["layout"]) ? (function($input) {
        \assert(is_array($input), 'expected disjunction value to be an array');
        /** @var array<string, mixed> $input */
        switch ($input["kind"]) {
        case "AutoGridLayout":
            return AutoGridLayoutKind::fromArray($input);
        case "GridLayout":
            return GridLayoutKind::fromArray($input);
        case "RowsLayout":
            return RowsLayoutKind::fromArray($input);
        case "TabsLayout":
            return TabsLayoutKind::fromArray($input);
        default:
            throw new \ValueError('can not parse disjunction from array');
    }
    })($data["layout"]) : null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->layout = $this->layout;
        if (isset($this->title)) {
            $data->title = $this->title;
        }
        if (isset($this->collapse)) {
            $data->collapse = $this->collapse;
        }
        if (isset($this->hideHeader)) {
            $data->hideHeader = $this->hideHeader;
        }
        if (isset($this->fillScreen)) {
            $data->fillScreen = $this->fillScreen;
        }
        if (isset($this->conditionalRendering)) {
            $data->conditionalRendering = $this->conditionalRendering;
        }
        if (isset($this->repeat)) {
            $data->repeat = $this->repeat;
        }
        return $data;
    }
}
