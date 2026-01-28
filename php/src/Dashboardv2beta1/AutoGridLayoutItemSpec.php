<?php

namespace Grafana\Foundation\Dashboardv2beta1;

class AutoGridLayoutItemSpec implements \JsonSerializable
{
    public \Grafana\Foundation\Dashboardv2beta1\ElementReference $element;

    public ?\Grafana\Foundation\Dashboardv2beta1\AutoGridRepeatOptions $repeat;

    public ?\Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingGroupKind $conditionalRendering;

    /**
     * @param \Grafana\Foundation\Dashboardv2beta1\ElementReference|null $element
     * @param \Grafana\Foundation\Dashboardv2beta1\AutoGridRepeatOptions|null $repeat
     * @param \Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingGroupKind|null $conditionalRendering
     */
    public function __construct(?\Grafana\Foundation\Dashboardv2beta1\ElementReference $element = null, ?\Grafana\Foundation\Dashboardv2beta1\AutoGridRepeatOptions $repeat = null, ?\Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingGroupKind $conditionalRendering = null)
    {
        $this->element = $element ?: new \Grafana\Foundation\Dashboardv2beta1\ElementReference();
        $this->repeat = $repeat;
        $this->conditionalRendering = $conditionalRendering;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{element?: mixed, repeat?: mixed, conditionalRendering?: mixed} $inputData */
        $data = $inputData;
        return new self(
            element: isset($data["element"]) ? (function($input) {
    	/** @var array{kind?: string, name?: string} */
    $val = $input;
    	return \Grafana\Foundation\Dashboardv2beta1\ElementReference::fromArray($val);
    })($data["element"]) : null,
            repeat: isset($data["repeat"]) ? (function($input) {
    	/** @var array{mode?: "variable", value?: string} */
    $val = $input;
    	return \Grafana\Foundation\Dashboardv2beta1\AutoGridRepeatOptions::fromArray($val);
    })($data["repeat"]) : null,
            conditionalRendering: isset($data["conditionalRendering"]) ? (function($input) {
    	/** @var array{kind?: string, spec?: mixed} */
    $val = $input;
    	return \Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingGroupKind::fromArray($val);
    })($data["conditionalRendering"]) : null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->element = $this->element;
        if (isset($this->repeat)) {
            $data->repeat = $this->repeat;
        }
        if (isset($this->conditionalRendering)) {
            $data->conditionalRendering = $this->conditionalRendering;
        }
        return $data;
    }
}
