<?php

namespace Grafana\Foundation\Dashboardv2beta1;

class AnnotationQuerySpec implements \JsonSerializable
{
    public \Grafana\Foundation\Dashboardv2beta1\DataQueryKind $query;

    public bool $enable;

    public bool $hide;

    public string $iconColor;

    public string $name;

    public ?bool $builtIn;

    public ?\Grafana\Foundation\Dashboardv2beta1\AnnotationPanelFilter $filter;

    /**
     * Placement can be used to display the annotation query somewhere else on the dashboard other than the default location.
     */
    public ?string $placement;

    /**
     * Mappings define how to convert data frame fields to annotation event fields.
     * @var array<string, \Grafana\Foundation\Dashboardv2beta1\AnnotationEventFieldMapping>|null
     */
    public ?array $mappings;

    /**
     * Catch-all field for datasource-specific properties. Should not be available in as code tooling.
     * @var array<string, mixed>|null
     */
    public ?array $legacyOptions;

    /**
     * @param \Grafana\Foundation\Dashboardv2beta1\DataQueryKind|null $query
     * @param bool|null $enable
     * @param bool|null $hide
     * @param string|null $iconColor
     * @param string|null $name
     * @param bool|null $builtIn
     * @param \Grafana\Foundation\Dashboardv2beta1\AnnotationPanelFilter|null $filter
     * @param array<string, \Grafana\Foundation\Dashboardv2beta1\AnnotationEventFieldMapping>|null $mappings
     * @param array<string, mixed>|null $legacyOptions
     */
    public function __construct(?\Grafana\Foundation\Dashboardv2beta1\DataQueryKind $query = null, ?bool $enable = null, ?bool $hide = null, ?string $iconColor = null, ?string $name = null, ?bool $builtIn = null, ?\Grafana\Foundation\Dashboardv2beta1\AnnotationPanelFilter $filter = null, ?array $mappings = null, ?array $legacyOptions = null)
    {
        $this->query = $query ?: new \Grafana\Foundation\Dashboardv2beta1\DataQueryKind();
        $this->enable = $enable ?: false;
        $this->hide = $hide ?: false;
        $this->iconColor = $iconColor ?: "";
        $this->name = $name ?: "";
        $this->builtIn = $builtIn;
        $this->filter = $filter;
        $this->placement = \Grafana\Foundation\Dashboardv2beta1\Constants::ANNOTATION_QUERY_PLACEMENT;
        $this->mappings = $mappings;
        $this->legacyOptions = $legacyOptions;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{query?: mixed, enable?: bool, hide?: bool, iconColor?: string, name?: string, builtIn?: bool, filter?: mixed, placement?: "inControlsMenu", mappings?: array<string, mixed>, legacyOptions?: array<string, mixed>} $inputData */
        $data = $inputData;
        return new self(
            query: isset($data["query"]) ? (function($input) {
    	/** @var array{kind?: string, group?: string, version?: string, datasource?: mixed, spec?: mixed} */
    $val = $input;
    	return \Grafana\Foundation\Dashboardv2beta1\DataQueryKind::fromArray($val);
    })($data["query"]) : null,
            enable: $data["enable"] ?? null,
            hide: $data["hide"] ?? null,
            iconColor: $data["iconColor"] ?? null,
            name: $data["name"] ?? null,
            builtIn: $data["builtIn"] ?? null,
            filter: isset($data["filter"]) ? (function($input) {
    	/** @var array{exclude?: bool, ids?: array<int>} */
    $val = $input;
    	return \Grafana\Foundation\Dashboardv2beta1\AnnotationPanelFilter::fromArray($val);
    })($data["filter"]) : null,
            mappings: isset($data["mappings"]) ? (function($input) {
        /** @var array<string, \Grafana\Foundation\Dashboardv2beta1\AnnotationEventFieldMapping> $results */
        $results = [];
        foreach ($input as $key => $val) {
            $results[$key] = isset($val) ? (function($input) {
    	/** @var array{source?: string, value?: string, regex?: string} */
    $val = $input;
    	return \Grafana\Foundation\Dashboardv2beta1\AnnotationEventFieldMapping::fromArray($val);
    })($val) : null;
        }
        return array_filter($results);
    })($data["mappings"]) : null,
            legacyOptions: $data["legacyOptions"] ?? null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->query = $this->query;
        $data->enable = $this->enable;
        $data->hide = $this->hide;
        $data->iconColor = $this->iconColor;
        $data->name = $this->name;
        if (isset($this->builtIn)) {
            $data->builtIn = $this->builtIn;
        }
        if (isset($this->filter)) {
            $data->filter = $this->filter;
        }
        if (isset($this->placement)) {
            $data->placement = $this->placement;
        }
        if (isset($this->mappings)) {
            $data->mappings = $this->mappings;
        }
        if (isset($this->legacyOptions)) {
            $data->legacyOptions = $this->legacyOptions;
        }
        return $data;
    }
}
