<?php

namespace Grafana\Foundation\Dashboardv2;

class RowsLayoutRowSpec implements \JsonSerializable
{
    public ?string $title;

    public ?bool $collapse;

    public ?bool $hideHeader;

    public ?bool $fillScreen;

    public ?\Grafana\Foundation\Dashboardv2\ConditionalRenderingGroupKind $conditionalRendering;

    public ?\Grafana\Foundation\Dashboardv2\RowRepeatOptions $repeat;

    /**
     * @var \Grafana\Foundation\Dashboardv2\GridLayoutKind|\Grafana\Foundation\Dashboardv2\AutoGridLayoutKind|\Grafana\Foundation\Dashboardv2\TabsLayoutKind|\Grafana\Foundation\Dashboardv2\RowsLayoutKind
     */
    public $layout;

    /**
     * @var array<\Grafana\Foundation\Dashboardv2\QueryVariableKind|\Grafana\Foundation\Dashboardv2\TextVariableKind|\Grafana\Foundation\Dashboardv2\ConstantVariableKind|\Grafana\Foundation\Dashboardv2\DatasourceVariableKind|\Grafana\Foundation\Dashboardv2\IntervalVariableKind|\Grafana\Foundation\Dashboardv2\CustomVariableKind|\Grafana\Foundation\Dashboardv2\GroupByVariableKind|\Grafana\Foundation\Dashboardv2\AdhocVariableKind|\Grafana\Foundation\Dashboardv2\SwitchVariableKind>|null
     */
    public ?array $variables;

    /**
     * @param string|null $title
     * @param bool|null $collapse
     * @param bool|null $hideHeader
     * @param bool|null $fillScreen
     * @param \Grafana\Foundation\Dashboardv2\ConditionalRenderingGroupKind|null $conditionalRendering
     * @param \Grafana\Foundation\Dashboardv2\RowRepeatOptions|null $repeat
     * @param \Grafana\Foundation\Dashboardv2\GridLayoutKind|\Grafana\Foundation\Dashboardv2\AutoGridLayoutKind|\Grafana\Foundation\Dashboardv2\TabsLayoutKind|\Grafana\Foundation\Dashboardv2\RowsLayoutKind|null $layout
     * @param array<\Grafana\Foundation\Dashboardv2\QueryVariableKind|\Grafana\Foundation\Dashboardv2\TextVariableKind|\Grafana\Foundation\Dashboardv2\ConstantVariableKind|\Grafana\Foundation\Dashboardv2\DatasourceVariableKind|\Grafana\Foundation\Dashboardv2\IntervalVariableKind|\Grafana\Foundation\Dashboardv2\CustomVariableKind|\Grafana\Foundation\Dashboardv2\GroupByVariableKind|\Grafana\Foundation\Dashboardv2\AdhocVariableKind|\Grafana\Foundation\Dashboardv2\SwitchVariableKind>|null $variables
     */
    public function __construct(?string $title = null, ?bool $collapse = null, ?bool $hideHeader = null, ?bool $fillScreen = null, ?\Grafana\Foundation\Dashboardv2\ConditionalRenderingGroupKind $conditionalRendering = null, ?\Grafana\Foundation\Dashboardv2\RowRepeatOptions $repeat = null,  $layout = null, ?array $variables = null)
    {
        $this->title = $title;
        $this->collapse = $collapse;
        $this->hideHeader = $hideHeader;
        $this->fillScreen = $fillScreen;
        $this->conditionalRendering = $conditionalRendering;
        $this->repeat = $repeat;
        $this->layout = $layout ?: new \Grafana\Foundation\Dashboardv2\GridLayoutKind();
        $this->variables = $variables;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{title?: string, collapse?: bool, hideHeader?: bool, fillScreen?: bool, conditionalRendering?: mixed, repeat?: mixed, layout?: mixed|mixed|mixed|mixed, variables?: array<mixed|mixed|mixed|mixed|mixed|mixed|mixed|mixed|mixed>} $inputData */
        $data = $inputData;
        return new self(
            title: $data["title"] ?? null,
            collapse: $data["collapse"] ?? null,
            hideHeader: $data["hideHeader"] ?? null,
            fillScreen: $data["fillScreen"] ?? null,
            conditionalRendering: isset($data["conditionalRendering"]) ? (function($input) {
    	/** @var array{kind?: string, spec?: mixed} */
    $val = $input;
    	return \Grafana\Foundation\Dashboardv2\ConditionalRenderingGroupKind::fromArray($val);
    })($data["conditionalRendering"]) : null,
            repeat: isset($data["repeat"]) ? (function($input) {
    	/** @var array{mode?: "variable", value?: string} */
    $val = $input;
    	return \Grafana\Foundation\Dashboardv2\RowRepeatOptions::fromArray($val);
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
            variables: !empty($data["variables"]) ? array_map((function($input) {
        \assert(is_array($input), 'expected disjunction value to be an array');
        /** @var array<string, mixed> $input */
        switch ($input["kind"]) {
        case "AdhocVariable":
            return AdhocVariableKind::fromArray($input);
        case "ConstantVariable":
            return ConstantVariableKind::fromArray($input);
        case "CustomVariable":
            return CustomVariableKind::fromArray($input);
        case "DatasourceVariable":
            return DatasourceVariableKind::fromArray($input);
        case "GroupByVariable":
            return GroupByVariableKind::fromArray($input);
        case "IntervalVariable":
            return IntervalVariableKind::fromArray($input);
        case "QueryVariable":
            return QueryVariableKind::fromArray($input);
        case "SwitchVariable":
            return SwitchVariableKind::fromArray($input);
        case "TextVariable":
            return TextVariableKind::fromArray($input);
        default:
            throw new \ValueError('can not parse disjunction from array');
    }
    }), $data["variables"]) : null,
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
        if (isset($this->variables)) {
            $data->variables = $this->variables;
        }
        return $data;
    }
}
