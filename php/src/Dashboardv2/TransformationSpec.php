<?php

namespace Grafana\Foundation\Dashboardv2;

/**
 * Transformations allow to manipulate data returned by a query before the system applies a visualization.
 * Using transformations you can: rename fields, join time series data, perform mathematical operations across queries,
 * use the output of one transformation as the input to another transformation, etc.
 */
class TransformationSpec implements \JsonSerializable
{
    /**
     * Disabled transformations are skipped
     */
    public ?bool $disabled;

    /**
     * Optional frame matcher. When missing it will be applied to all results
     */
    public ?\Grafana\Foundation\Dashboardv2\MatcherConfig $filter;

    /**
     * Where to pull DataFrames from as input to transformation
     */
    public ?\Grafana\Foundation\Dashboardv2\DataTopic $topic;

    /**
     * Options to be passed to the transformer
     * Valid options depend on the transformer id
     * @var mixed
     */
    public $options;

    /**
     * @param bool|null $disabled
     * @param \Grafana\Foundation\Dashboardv2\MatcherConfig|null $filter
     * @param \Grafana\Foundation\Dashboardv2\DataTopic|null $topic
     * @param mixed|null $options
     */
    public function __construct(?bool $disabled = null, ?\Grafana\Foundation\Dashboardv2\MatcherConfig $filter = null, ?\Grafana\Foundation\Dashboardv2\DataTopic $topic = null,  $options = null)
    {
        $this->disabled = $disabled;
        $this->filter = $filter;
        $this->topic = $topic;
        $this->options = $options ?: null;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{disabled?: bool, filter?: mixed, topic?: string, options?: mixed} $inputData */
        $data = $inputData;
        return new self(
            disabled: $data["disabled"] ?? null,
            filter: isset($data["filter"]) ? (function($input) {
    	/** @var array{id?: string, scope?: string, options?: mixed} */
    $val = $input;
    	return \Grafana\Foundation\Dashboardv2\MatcherConfig::fromArray($val);
    })($data["filter"]) : null,
            topic: isset($data["topic"]) ? (function($input) { return \Grafana\Foundation\Dashboardv2\DataTopic::fromValue($input); })($data["topic"]) : null,
            options: $data["options"] ?? null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->options = $this->options;
        if (isset($this->disabled)) {
            $data->disabled = $this->disabled;
        }
        if (isset($this->filter)) {
            $data->filter = $this->filter;
        }
        if (isset($this->topic)) {
            $data->topic = $this->topic;
        }
        return $data;
    }
}
