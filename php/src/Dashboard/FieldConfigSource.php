<?php

namespace Grafana\Foundation\Dashboard;

/**
 * The data model used in Grafana, namely the data frame, is a columnar-oriented table structure that unifies both time series and table query results.
 * Each column within this structure is called a field. A field can represent a single time series or table column.
 * Field options allow you to change how the data is displayed in your visualizations.
 */
class FieldConfigSource implements \JsonSerializable
{
    /**
     * Defaults are the options applied to all fields.
     */
    public \Grafana\Foundation\Dashboard\FieldConfig $defaults;

    /**
     * Overrides are the options applied to specific fields overriding the defaults.
     * @var array<\Grafana\Foundation\Dashboard\DashboardFieldConfigSourceOverrides>
     */
    public array $overrides;

    /**
     * @param \Grafana\Foundation\Dashboard\FieldConfig|null $defaults
     * @param array<\Grafana\Foundation\Dashboard\DashboardFieldConfigSourceOverrides>|null $overrides
     */
    public function __construct(?\Grafana\Foundation\Dashboard\FieldConfig $defaults = null, ?array $overrides = null)
    {
        $this->defaults = $defaults ?: new \Grafana\Foundation\Dashboard\FieldConfig();
        $this->overrides = $overrides ?: [];
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{defaults?: mixed, overrides?: array<mixed>} $inputData */
        $data = $inputData;
        return new self(
            defaults: isset($data["defaults"]) ? (function($input) {
    	/** @var array{displayName?: string, displayNameFromDS?: string, description?: string, path?: string, writeable?: bool, filterable?: bool, unit?: string, decimals?: float, min?: float, max?: float, mappings?: array<mixed|mixed|mixed|mixed>, thresholds?: mixed, color?: mixed, links?: array<mixed>, noValue?: string, custom?: mixed} */
    $val = $input;
    	return \Grafana\Foundation\Dashboard\FieldConfig::fromArray($val);
    })($data["defaults"]) : null,
            overrides: array_filter(array_map((function($input) {
    	/** @var array{matcher?: mixed, properties?: array<mixed>} */
    $val = $input;
    	return \Grafana\Foundation\Dashboard\DashboardFieldConfigSourceOverrides::fromArray($val);
    }), $data["overrides"] ?? [])),
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->defaults = $this->defaults;
        $data->overrides = $this->overrides;
        return $data;
    }
}
