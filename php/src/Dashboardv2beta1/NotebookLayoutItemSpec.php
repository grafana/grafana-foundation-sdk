<?php

namespace Grafana\Foundation\Dashboardv2beta1;

/**
 * One ordered item in a notebook layout. `element` references either a CellKind
 * (markdown/code content) or a V2PanelKind in the notebook's elements map. `source`
 * records who authored the cell; `collapsed` hides the body in the UI.
 */
class NotebookLayoutItemSpec implements \JsonSerializable
{
    public \Grafana\Foundation\Dashboardv2beta1\ElementReference $element;

    public \Grafana\Foundation\Dashboardv2beta1\NotebookLayoutItemSpecSource $source;

    public ?bool $collapsed;

    /**
     * @param \Grafana\Foundation\Dashboardv2beta1\ElementReference|null $element
     * @param \Grafana\Foundation\Dashboardv2beta1\NotebookLayoutItemSpecSource|null $source
     * @param bool|null $collapsed
     */
    public function __construct(?\Grafana\Foundation\Dashboardv2beta1\ElementReference $element = null, ?\Grafana\Foundation\Dashboardv2beta1\NotebookLayoutItemSpecSource $source = null, ?bool $collapsed = null)
    {
        $this->element = $element ?: new \Grafana\Foundation\Dashboardv2beta1\ElementReference();
        $this->source = $source ?: \Grafana\Foundation\Dashboardv2beta1\NotebookLayoutItemSpecSource::Assistant();
        $this->collapsed = $collapsed;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{element?: mixed, source?: string, collapsed?: bool} $inputData */
        $data = $inputData;
        return new self(
            element: isset($data["element"]) ? (function($input) {
    	/** @var array{kind?: string, name?: string} */
    $val = $input;
    	return \Grafana\Foundation\Dashboardv2beta1\ElementReference::fromArray($val);
    })($data["element"]) : null,
            source: isset($data["source"]) ? (function($input) { return \Grafana\Foundation\Dashboardv2beta1\NotebookLayoutItemSpecSource::fromValue($input); })($data["source"]) : null,
            collapsed: $data["collapsed"] ?? null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->element = $this->element;
        $data->source = $this->source;
        if (isset($this->collapsed)) {
            $data->collapsed = $this->collapsed;
        }
        return $data;
    }
}
