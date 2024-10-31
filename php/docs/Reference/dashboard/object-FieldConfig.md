---
title: <span class="badge object-type-class"></span> FieldConfig
---
# <span class="badge object-type-class"></span> FieldConfig

The data model used in Grafana, namely the data frame, is a columnar-oriented table structure that unifies both time series and table query results.

Each column within this structure is called a field. A field can represent a single time series or table column.

Field options allow you to change how the data is displayed in your visualizations.

## Definition

```php
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
     * @var array<mixed>|null
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

}
```
## Methods

### <span class="badge object-method"></span> fromArray

Builds this object from an array.

This function is meant to be used with the return value of `json_decode($json, true)`.

```php
static fromArray(array $inputData)
```

### <span class="badge object-method"></span> jsonSerialize

Returns the data representing this object, preparing it for JSON serialization with `json_encode()`.

```php
jsonSerialize()
```

