---
title: <span class="badge object-type-class"></span> Dataquery
---
# <span class="badge object-type-class"></span> Dataquery

## Definition

```php
class Dataquery implements \JsonSerializable, \Grafana\Foundation\Cog\Dataquery
{
    public ?string $alias;

    public ?\Grafana\Foundation\Testdata\TestDataQueryType $scenarioId;

    public ?string $stringInput;

    public ?\Grafana\Foundation\Testdata\StreamingQuery $stream;

    public ?\Grafana\Foundation\Testdata\PulseWaveQuery $pulseWave;

    public ?\Grafana\Foundation\Testdata\SimulationQuery $sim;

    /**
     * @var array<\Grafana\Foundation\Testdata\CSVWave>|null
     */
    public ?array $csvWave;

    public ?string $labels;

    public ?int $lines;

    public ?bool $levelColumn;

    public ?string $channel;

    public ?\Grafana\Foundation\Testdata\NodesQuery $nodes;

    public ?string $csvFileName;

    public ?string $csvContent;

    public ?string $rawFrameContent;

    public ?int $seriesCount;

    public ?\Grafana\Foundation\Testdata\USAQuery $usa;

    public ?\Grafana\Foundation\Testdata\DataqueryErrorType $errorType;

    public ?int $spanCount;

    /**
     * @var array<array<string|int>>|null
     */
    public ?array $points;

    /**
     * Drop percentage (the chance we will lose a point 0-100)
     */
    public ?float $dropPercent;

    /**
     * A unique identifier for the query within the list of targets.
     * In server side expressions, the refId is used as a variable name to identify results.
     * By default, the UI will assign A->Z; however setting meaningful names may be useful.
     */
    public string $refId;

    /**
     * true if query is disabled (ie should not be returned to the dashboard)
     * Note this does not always imply that the query should not be executed since
     * the results from a hidden query may be used as the input to other queries (SSE etc)
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
