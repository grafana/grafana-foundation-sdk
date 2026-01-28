<?php

namespace Grafana\Foundation\Dashboardv2beta1;

class ConditionalRenderingGroupSpec implements \JsonSerializable
{
    public \Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingGroupSpecVisibility $visibility;

    public \Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingGroupSpecCondition $condition;

    /**
     * @var array<\Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingVariableKind|\Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingDataKind|\Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingTimeRangeSizeKind>
     */
    public array $items;

    /**
     * @param \Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingGroupSpecVisibility|null $visibility
     * @param \Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingGroupSpecCondition|null $condition
     * @param array<\Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingVariableKind|\Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingDataKind|\Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingTimeRangeSizeKind>|null $items
     */
    public function __construct(?\Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingGroupSpecVisibility $visibility = null, ?\Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingGroupSpecCondition $condition = null, ?array $items = null)
    {
        $this->visibility = $visibility ?: \Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingGroupSpecVisibility::Show();
        $this->condition = $condition ?: \Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingGroupSpecCondition::And();
        $this->items = $items ?: [];
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{visibility?: string, condition?: string, items?: array<mixed|mixed|mixed>} $inputData */
        $data = $inputData;
        return new self(
            visibility: isset($data["visibility"]) ? (function($input) { return \Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingGroupSpecVisibility::fromValue($input); })($data["visibility"]) : null,
            condition: isset($data["condition"]) ? (function($input) { return \Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingGroupSpecCondition::fromValue($input); })($data["condition"]) : null,
            items: !empty($data["items"]) ? array_map((function($input) {
        \assert(is_array($input), 'expected disjunction value to be an array');
        /** @var array<string, mixed> $input */
        switch ($input["kind"]) {
        case "ConditionalRenderingData":
            return ConditionalRenderingDataKind::fromArray($input);
        case "ConditionalRenderingTimeRangeSize":
            return ConditionalRenderingTimeRangeSizeKind::fromArray($input);
        case "ConditionalRenderingVariable":
            return ConditionalRenderingVariableKind::fromArray($input);
        default:
            throw new \ValueError('can not parse disjunction from array');
    }
    }), $data["items"]) : null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->visibility = $this->visibility;
        $data->condition = $this->condition;
        $data->items = $this->items;
        return $data;
    }
}
