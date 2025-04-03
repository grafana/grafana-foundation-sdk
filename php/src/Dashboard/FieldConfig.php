<?php

namespace Grafana\Foundation\Dashboard;

/**
 * The data model used in Grafana, namely the data frame, is a columnar-oriented table structure that unifies both time series and table query results.
 * Each column within this structure is called a field. A field can represent a single time series or table column.
 * Field options allow you to change how the data is displayed in your visualizations.
 */
class FieldConfig implements \JsonSerializable
{
    /**
     * The display value for this field.  This supports template variables blank is auto
     */
    public ?string $displayName;

    /**
     * This can be used by data sources that return and explicit naming structure for values and labels
     * When this property is configured, this value is used rather than the default naming strategy.
     */
    public ?string $displayNameFromDS;

    /**
     * Human readable field metadata
     */
    public ?string $description;

    /**
     * An explicit path to the field in the datasource.  When the frame meta includes a path,
     * This will default to `${frame.meta.path}/${field.name}
     * 
     * When defined, this value can be used as an identifier within the datasource scope, and
     * may be used to update the results
     */
    public ?string $path;

    /**
     * True if data source can write a value to the path. Auth/authz are supported separately
     */
    public ?bool $writeable;

    /**
     * True if data source field supports ad-hoc filters
     */
    public ?bool $filterable;

    /**
     * Unit a field should use. The unit you select is applied to all fields except time.
     * You can use the units ID availables in Grafana or a custom unit.
     * Available units in Grafana: https://github.com/grafana/grafana/blob/main/packages/grafana-data/src/valueFormats/categories.ts
     * As custom unit, you can use the following formats:
     * `suffix:<suffix>` for custom unit that should go after value.
     * `prefix:<prefix>` for custom unit that should go before value.
     * `time:<format>` For custom date time formats type for example `time:YYYY-MM-DD`.
     * `si:<base scale><unit characters>` for custom SI units. For example: `si: mF`. This one is a bit more advanced as you can specify both a unit and the source data scale. So if your source data is represented as milli (thousands of) something prefix the unit with that SI scale character.
     * `count:<unit>` for a custom count unit.
     * `currency:<unit>` for custom a currency unit.
     */
    public ?string $unit;

    /**
     * Specify the number of decimals Grafana includes in the rendered value.
     * If you leave this field blank, Grafana automatically truncates the number of decimals based on the value.
     * For example 1.1234 will display as 1.12 and 100.456 will display as 100.
     * To display all decimals, set the unit to `String`.
     */
    public ?float $decimals;

    /**
     * The minimum value used in percentage threshold calculations. Leave blank for auto calculation based on all series and fields.
     */
    public ?float $min;

    /**
     * The maximum value used in percentage threshold calculations. Leave blank for auto calculation based on all series and fields.
     */
    public ?float $max;

    /**
     * Convert input values into a display string
     * @var array<\Grafana\Foundation\Dashboard\ValueMap|\Grafana\Foundation\Dashboard\RangeMap|\Grafana\Foundation\Dashboard\RegexMap|\Grafana\Foundation\Dashboard\SpecialValueMap>|null
     */
    public ?array $mappings;

    /**
     * Map numeric values to states
     */
    public ?\Grafana\Foundation\Dashboard\ThresholdsConfig $thresholds;

    /**
     * Panel color configuration
     */
    public ?\Grafana\Foundation\Dashboard\FieldColor $color;

    /**
     * The behavior when clicking on a result
     * @var array<\Grafana\Foundation\Dashboard\DashboardLink>|null
     */
    public ?array $links;

    /**
     * Alternative to empty string
     */
    public ?string $noValue;

    /**
     * custom is specified by the FieldConfig field
     * in panel plugin schemas.
     * @var mixed|null
     */
    public $custom;

    /**
     * @param string|null $displayName
     * @param string|null $displayNameFromDS
     * @param string|null $description
     * @param string|null $path
     * @param bool|null $writeable
     * @param bool|null $filterable
     * @param string|null $unit
     * @param float|null $decimals
     * @param float|null $min
     * @param float|null $max
     * @param array<\Grafana\Foundation\Dashboard\ValueMap|\Grafana\Foundation\Dashboard\RangeMap|\Grafana\Foundation\Dashboard\RegexMap|\Grafana\Foundation\Dashboard\SpecialValueMap>|null $mappings
     * @param \Grafana\Foundation\Dashboard\ThresholdsConfig|null $thresholds
     * @param \Grafana\Foundation\Dashboard\FieldColor|null $color
     * @param array<\Grafana\Foundation\Dashboard\DashboardLink>|null $links
     * @param string|null $noValue
     * @param mixed|null $custom
     */
    public function __construct(?string $displayName = null, ?string $displayNameFromDS = null, ?string $description = null, ?string $path = null, ?bool $writeable = null, ?bool $filterable = null, ?string $unit = null, ?float $decimals = null, ?float $min = null, ?float $max = null, ?array $mappings = null, ?\Grafana\Foundation\Dashboard\ThresholdsConfig $thresholds = null, ?\Grafana\Foundation\Dashboard\FieldColor $color = null, ?array $links = null, ?string $noValue = null,  $custom = null)
    {
        $this->displayName = $displayName;
        $this->displayNameFromDS = $displayNameFromDS;
        $this->description = $description;
        $this->path = $path;
        $this->writeable = $writeable;
        $this->filterable = $filterable;
        $this->unit = $unit;
        $this->decimals = $decimals;
        $this->min = $min;
        $this->max = $max;
        $this->mappings = $mappings;
        $this->thresholds = $thresholds;
        $this->color = $color;
        $this->links = $links;
        $this->noValue = $noValue;
        $this->custom = $custom;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{displayName?: string, displayNameFromDS?: string, description?: string, path?: string, writeable?: bool, filterable?: bool, unit?: string, decimals?: float, min?: float, max?: float, mappings?: array<mixed|mixed|mixed|mixed>, thresholds?: mixed, color?: mixed, links?: array<mixed>, noValue?: string, custom?: mixed} $inputData */
        $data = $inputData;
        return new self(
            displayName: $data["displayName"] ?? null,
            displayNameFromDS: $data["displayNameFromDS"] ?? null,
            description: $data["description"] ?? null,
            path: $data["path"] ?? null,
            writeable: $data["writeable"] ?? null,
            filterable: $data["filterable"] ?? null,
            unit: $data["unit"] ?? null,
            decimals: $data["decimals"] ?? null,
            min: $data["min"] ?? null,
            max: $data["max"] ?? null,
            mappings: !empty($data["mappings"]) ? array_map((function($input) {
        \assert(is_array($input), 'expected disjunction value to be an array');
        /** @var array<string, mixed> $input */
        switch ($input["type"]) {
        case "range":
            return RangeMap::fromArray($input);
        case "regex":
            return RegexMap::fromArray($input);
        case "special":
            return SpecialValueMap::fromArray($input);
        case "value":
            return ValueMap::fromArray($input);
        default:
            throw new \ValueError('can not parse disjunction from array');
    }
    }), $data["mappings"]) : null,
            thresholds: isset($data["thresholds"]) ? (function($input) {
    	/** @var array{mode?: string, steps?: array<mixed>} */
    $val = $input;
    	return \Grafana\Foundation\Dashboard\ThresholdsConfig::fromArray($val);
    })($data["thresholds"]) : null,
            color: isset($data["color"]) ? (function($input) {
    	/** @var array{mode?: string, fixedColor?: string, seriesBy?: string} */
    $val = $input;
    	return \Grafana\Foundation\Dashboard\FieldColor::fromArray($val);
    })($data["color"]) : null,
            links: array_filter(array_map((function($input) {
    	/** @var array{title?: string, type?: string, icon?: string, tooltip?: string, url?: string, tags?: array<string>, asDropdown?: bool, targetBlank?: bool, includeVars?: bool, keepTime?: bool} */
    $val = $input;
    	return \Grafana\Foundation\Dashboard\DashboardLink::fromArray($val);
    }), $data["links"] ?? [])),
            noValue: $data["noValue"] ?? null,
            custom: $data["custom"] ?? null,
        );
    }

    /**
     * @return array<string, mixed>
     */
    public function jsonSerialize(): array
    {
        $data = [
        ];
        if (isset($this->displayName)) {
            $data["displayName"] = $this->displayName;
        }
        if (isset($this->displayNameFromDS)) {
            $data["displayNameFromDS"] = $this->displayNameFromDS;
        }
        if (isset($this->description)) {
            $data["description"] = $this->description;
        }
        if (isset($this->path)) {
            $data["path"] = $this->path;
        }
        if (isset($this->writeable)) {
            $data["writeable"] = $this->writeable;
        }
        if (isset($this->filterable)) {
            $data["filterable"] = $this->filterable;
        }
        if (isset($this->unit)) {
            $data["unit"] = $this->unit;
        }
        if (isset($this->decimals)) {
            $data["decimals"] = $this->decimals;
        }
        if (isset($this->min)) {
            $data["min"] = $this->min;
        }
        if (isset($this->max)) {
            $data["max"] = $this->max;
        }
        if (isset($this->mappings)) {
            $data["mappings"] = $this->mappings;
        }
        if (isset($this->thresholds)) {
            $data["thresholds"] = $this->thresholds;
        }
        if (isset($this->color)) {
            $data["color"] = $this->color;
        }
        if (isset($this->links)) {
            $data["links"] = $this->links;
        }
        if (isset($this->noValue)) {
            $data["noValue"] = $this->noValue;
        }
        if (isset($this->custom)) {
            $data["custom"] = $this->custom;
        }
        return $data;
    }
}
