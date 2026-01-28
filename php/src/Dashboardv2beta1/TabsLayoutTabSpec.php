<?php

namespace Grafana\Foundation\Dashboardv2beta1;

class TabsLayoutTabSpec implements \JsonSerializable
{
    public ?string $title;

    /**
     * @var \Grafana\Foundation\Dashboardv2beta1\GridLayoutKind|\Grafana\Foundation\Dashboardv2beta1\RowsLayoutKind|\Grafana\Foundation\Dashboardv2beta1\AutoGridLayoutKind|\Grafana\Foundation\Dashboardv2beta1\TabsLayoutKind
     */
    public $layout;

    public ?\Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingGroupKind $conditionalRendering;

    public ?\Grafana\Foundation\Dashboardv2beta1\TabRepeatOptions $repeat;

    /**
     * @param string|null $title
     * @param \Grafana\Foundation\Dashboardv2beta1\GridLayoutKind|\Grafana\Foundation\Dashboardv2beta1\RowsLayoutKind|\Grafana\Foundation\Dashboardv2beta1\AutoGridLayoutKind|\Grafana\Foundation\Dashboardv2beta1\TabsLayoutKind|null $layout
     * @param \Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingGroupKind|null $conditionalRendering
     * @param \Grafana\Foundation\Dashboardv2beta1\TabRepeatOptions|null $repeat
     */
    public function __construct(?string $title = null,  $layout = null, ?\Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingGroupKind $conditionalRendering = null, ?\Grafana\Foundation\Dashboardv2beta1\TabRepeatOptions $repeat = null)
    {
        $this->title = $title;
        $this->layout = $layout ?: new \Grafana\Foundation\Dashboardv2beta1\GridLayoutKind();
        $this->conditionalRendering = $conditionalRendering;
        $this->repeat = $repeat;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{title?: string, layout?: mixed|mixed|mixed|mixed, conditionalRendering?: mixed, repeat?: mixed} $inputData */
        $data = $inputData;
        return new self(
            title: $data["title"] ?? null,
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
            conditionalRendering: isset($data["conditionalRendering"]) ? (function($input) {
    	/** @var array{kind?: string, spec?: mixed} */
    $val = $input;
    	return \Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingGroupKind::fromArray($val);
    })($data["conditionalRendering"]) : null,
            repeat: isset($data["repeat"]) ? (function($input) {
    	/** @var array{mode?: "variable", value?: string} */
    $val = $input;
    	return \Grafana\Foundation\Dashboardv2beta1\TabRepeatOptions::fromArray($val);
    })($data["repeat"]) : null,
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
        if (isset($this->conditionalRendering)) {
            $data->conditionalRendering = $this->conditionalRendering;
        }
        if (isset($this->repeat)) {
            $data->repeat = $this->repeat;
        }
        return $data;
    }
}
