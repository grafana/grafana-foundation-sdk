<?php

namespace Grafana\Foundation\Dashboard;

/**
 * Transformations allow to manipulate data returned by a query before the system applies a visualization.
 * Using transformations you can: rename fields, join time series data, perform mathematical operations across queries,
 * use the output of one transformation as the input to another transformation, etc.
 */
class DataTransformerConfig implements \JsonSerializable
{
    /**
     * Unique identifier of transformer
     */
    public string $id;

    /**
     * Disabled transformations are skipped
     */
    public ?bool $disabled;

    /**
     * Optional frame matcher. When missing it will be applied to all results
     */
    public ?\Grafana\Foundation\Dashboard\MatcherConfig $filter;

    /**
     * Options to be passed to the transformer
     * Valid options depend on the transformer id
     * @var mixed
     */
    public $options;

    /**
     * @param string|null $id
     * @param bool|null $disabled
     * @param \Grafana\Foundation\Dashboard\MatcherConfig|null $filter
     * @param mixed|null $options
     */
    public function __construct(?string $id = null, ?bool $disabled = null, ?\Grafana\Foundation\Dashboard\MatcherConfig $filter = null,  $options = null)
    {
        $this->id = $id ?: "";
        $this->disabled = $disabled;
        $this->filter = $filter;
        $this->options = $options ?: null;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{id?: string, disabled?: bool, filter?: mixed, options?: mixed} $inputData */
        $data = $inputData;
        return new self(
            id: $data["id"] ?? null,
            disabled: $data["disabled"] ?? null,
            filter: isset($data["filter"]) ? (function($input) {
    	/** @var array{id?: string, options?: mixed} */
    $val = $input;
    	return \Grafana\Foundation\Dashboard\MatcherConfig::fromArray($val);
    })($data["filter"]) : null,
            options: $data["options"] ?? null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
            "id" => $this->id,
            "options" => $this->options,
        ];
        if (isset($this->disabled)) {
            $data["disabled"] = $this->disabled;
        }
        if (isset($this->filter)) {
            $data["filter"] = $this->filter;
        }
        return $data;
    }
}
