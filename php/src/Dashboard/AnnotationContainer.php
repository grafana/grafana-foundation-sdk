<?php

namespace Grafana\Foundation\Dashboard;

/**
 * Contains the list of annotations that are associated with the dashboard.
 * Annotations are used to overlay event markers and overlay event tags on graphs.
 * Grafana comes with a native annotation store and the ability to add annotation events directly from the graph panel or via the HTTP API.
 * See https://grafana.com/docs/grafana/latest/dashboards/build-dashboards/annotate-visualizations/
 */
class AnnotationContainer implements \JsonSerializable
{
    /**
     * List of annotations
     * @var array<\Grafana\Foundation\Dashboard\AnnotationQuery>|null
     */
    public ?array $list;

    /**
     * @param array<\Grafana\Foundation\Dashboard\AnnotationQuery>|null $list
     */
    public function __construct(?array $list = null)
    {
        $this->list = $list;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{list?: array<mixed>} $inputData */
        $data = $inputData;
        return new self(
            list: array_filter(array_map((function($input) {
    	/** @var array{name?: string, datasource?: mixed, enable?: bool, hide?: bool, iconColor?: string, filter?: mixed, target?: mixed, type?: string, builtIn?: float, placement?: "inControlsMenu", expr?: string} */
    $val = $input;
    	return \Grafana\Foundation\Dashboard\AnnotationQuery::fromArray($val);
    }), $data["list"] ?? [])),
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        if (isset($this->list)) {
            $data->list = $this->list;
        }
        return $data;
    }
}
