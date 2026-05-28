<?php

namespace Grafana\Foundation\Dashboardv2;

class TabsLayoutTabSpec implements \JsonSerializable
{
    public ?string $title;

    /**
     * @var \Grafana\Foundation\Dashboardv2\GridLayoutKind|\Grafana\Foundation\Dashboardv2\RowsLayoutKind|\Grafana\Foundation\Dashboardv2\AutoGridLayoutKind|\Grafana\Foundation\Dashboardv2\TabsLayoutKind
     */
    public $layout;

    public ?\Grafana\Foundation\Dashboardv2\ConditionalRenderingGroupKind $conditionalRendering;

    public ?\Grafana\Foundation\Dashboardv2\TabRepeatOptions $repeat;

    /**
     * @var array<\Grafana\Foundation\Dashboardv2\QueryVariableKind|\Grafana\Foundation\Dashboardv2\TextVariableKind|\Grafana\Foundation\Dashboardv2\ConstantVariableKind|\Grafana\Foundation\Dashboardv2\DatasourceVariableKind|\Grafana\Foundation\Dashboardv2\IntervalVariableKind|\Grafana\Foundation\Dashboardv2\CustomVariableKind|\Grafana\Foundation\Dashboardv2\GroupByVariableKind|\Grafana\Foundation\Dashboardv2\AdhocVariableKind|\Grafana\Foundation\Dashboardv2\SwitchVariableKind>|null
     */
    public ?array $variables;

    /**
     * @param string|null $title
     * @param \Grafana\Foundation\Dashboardv2\GridLayoutKind|\Grafana\Foundation\Dashboardv2\RowsLayoutKind|\Grafana\Foundation\Dashboardv2\AutoGridLayoutKind|\Grafana\Foundation\Dashboardv2\TabsLayoutKind|null $layout
     * @param \Grafana\Foundation\Dashboardv2\ConditionalRenderingGroupKind|null $conditionalRendering
     * @param \Grafana\Foundation\Dashboardv2\TabRepeatOptions|null $repeat
     * @param array<\Grafana\Foundation\Dashboardv2\QueryVariableKind|\Grafana\Foundation\Dashboardv2\TextVariableKind|\Grafana\Foundation\Dashboardv2\ConstantVariableKind|\Grafana\Foundation\Dashboardv2\DatasourceVariableKind|\Grafana\Foundation\Dashboardv2\IntervalVariableKind|\Grafana\Foundation\Dashboardv2\CustomVariableKind|\Grafana\Foundation\Dashboardv2\GroupByVariableKind|\Grafana\Foundation\Dashboardv2\AdhocVariableKind|\Grafana\Foundation\Dashboardv2\SwitchVariableKind>|null $variables
     */
    public function __construct(?string $title = null,  $layout = null, ?\Grafana\Foundation\Dashboardv2\ConditionalRenderingGroupKind $conditionalRendering = null, ?\Grafana\Foundation\Dashboardv2\TabRepeatOptions $repeat = null, ?array $variables = null)
    {
        $this->title = $title;
        $this->layout = $layout ?: new \Grafana\Foundation\Dashboardv2\GridLayoutKind();
        $this->conditionalRendering = $conditionalRendering;
        $this->repeat = $repeat;
        $this->variables = $variables;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{title?: string, layout?: mixed|mixed|mixed|mixed, conditionalRendering?: mixed, repeat?: mixed, variables?: array<mixed|mixed|mixed|mixed|mixed|mixed|mixed|mixed|mixed>} $inputData */
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
    	return \Grafana\Foundation\Dashboardv2\ConditionalRenderingGroupKind::fromArray($val);
    })($data["conditionalRendering"]) : null,
            repeat: isset($data["repeat"]) ? (function($input) {
    	/** @var array{mode?: "variable", value?: string} */
    $val = $input;
    	return \Grafana\Foundation\Dashboardv2\TabRepeatOptions::fromArray($val);
    })($data["repeat"]) : null,
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
