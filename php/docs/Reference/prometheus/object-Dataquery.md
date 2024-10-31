---
title: <span class="badge object-type-class"></span> Dataquery
---
# <span class="badge object-type-class"></span> Dataquery

## Definition

```php
class Dataquery implements \JsonSerializable, \Grafana\Foundation\Cog\Dataquery
{
    /**
     * The actual expression/query that will be evaluated by Prometheus
     */
    public string $expr;

    /**
     * Returns only the latest value that Prometheus has scraped for the requested time series
     */
    public ?bool $instant;

    /**
     * Returns a Range vector, comprised of a set of time series containing a range of data points over time for each time series
     */
    public ?bool $range;

    /**
     * Execute an additional query to identify interesting raw samples relevant for the given expr
     */
    public ?bool $exemplar;

    /**
     * Specifies which editor is being used to prepare the query. It can be "code" or "builder"
     */
    public ?\Grafana\Foundation\Prometheus\QueryEditorMode $editorMode;

    /**
     * Query format to determine how to display data points in panel. It can be "time_series", "table", "heatmap"
     */
    public ?\Grafana\Foundation\Prometheus\PromQueryFormat $format;

    /**
     * Series name override or template. Ex. {{hostname}} will be replaced with label value for hostname
     */
    public ?string $legendFormat;

    /**
     * @deprecated Used to specify how many times to divide max data points by. We use max data points under query options
     * See https://github.com/grafana/grafana/issues/48081
     */
    public ?float $intervalFactor;

    /**
     * A unique identifier for the query within the list of targets.
     * In server side expressions, the refId is used as a variable name to identify results.
     * By default, the UI will assign A->Z; however setting meaningful names may be useful.
     */
    public string $refId;

    /**
     * If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.
     */
    public ?bool $hide;

    /**
     * Specify the query flavor
     * TODO make this required and give it a default
     */
    public ?string $queryType;

    /**
     * For mixed data sources the selected datasource is on the query level.
     * For non mixed scenarios this is undefined.
     * TODO find a better way to do this ^ that's friendly to schema
     * TODO this shouldn't be unknown but DataSourceRef | null
     */
    public ?\Grafana\Foundation\Dashboard\DataSourceRef $datasource;

    /**
     * An additional lower limit for the step parameter of the Prometheus query and for the
     * `$__interval` and `$__rate_interval` variables.
     */
    public ?string $interval;

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

### <span class="badge object-method"></span> dataqueryType

Returns the type of this dataquery object.

```php
dataqueryType()
```

## See also

 * <span class="badge builder"></span> [DataqueryBuilder](./builder-DataqueryBuilder.md)
