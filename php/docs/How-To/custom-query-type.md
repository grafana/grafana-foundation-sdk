# Defining a custom query type

While the SDK ships with support for all core datasources and their query types,
it can be extended for private/third-party plugins.

To do so, define a type for the custom query:

```php
<?php

use Grafana\Foundation\Cog;

class CustomQuery implements \JsonSerializable, Cog\Dataquery
{
    // RefId and Hide are expected on all queries
    public string $refId;
    public ?bool $hide;

    // Query is specific to the CustomQuery type
    public string $expr;

    /**
     * @param string|null $expr
     * @param string|null $refId
     * @param bool|null $hide
     */
    public function __construct(?string $expr = null, ?string $refId = null, ?bool $hide = null)
    {
        $this->expr = $expr ?: "";
        $this->refId = $refId ?: "";
        $this->hide = $hide;
    }

    /**
     * @param array{expr?: string, refId?: string, hide?: bool} $data
     */
    public static function fromArray(array $data): self
    {
        return new self(
            expr: $data["expr"] ?? null,
            refId: $data["refId"] ?? null,
            hide: $data["hide"] ?? null,
        );
    }

    public function jsonSerialize(): array
    {
        $data = [
            "expr" => $this->expr,
            "refId" => $this->refId,
        ];
        if (isset($this->hide)) {
            $data["hide"] = $this->hide;
        }
        return $data;
    }

    public function dataqueryType(): string
    {
        return "custom";
    }
}
```

Now, let's define a builder for that type:

```php
<?php

use Grafana\Foundation\Cog;

/**
 * @implements Cog\Builder<CustomQuery>
 */
class CustomQueryBuilder implements Cog\Builder
{
    protected CustomQuery $internal;

    public function __construct(string $query)
    {
    	$this->internal = new CustomQuery(expr: $query);
    }

    /**
     * @return CustomQuery
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * The actual expression/query that will be evaluated by Prometheus
     */
    public function expr(string $expr): static
    {
        $this->internal->expr = $expr;
        return $this;
    }

    /**
     * A unique identifier for the query within the list of targets.
     * In server side expressions, the refId is used as a variable name to identify results.
     * By default, the UI will assign A->Z; however setting meaningful names may be useful.
     */
    public function refId(string $refId): static
    {
        $this->internal->refId = $refId;
        return $this;
    }

    /**
     * If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.
     */
    public function hide(bool $hide): static
    {
        $this->internal->hide = $hide;
        return $this;
    }
}
```

Register the type with cog, and use it as usual to build a dashboard:

```php
<?php

use Grafana\Foundation\Cog;
use Grafana\Foundation\Dashboard\DashboardBuilder;
use Grafana\Foundation\Dashboard\RowBuilder;
use Grafana\Foundation\Timeseries;

require_once __DIR__.'/vendor/autoload.php';

// This lets cog know about the newly created query type and how to unmarshal it.
Cog\Runtime::get()->registerDataqueryVariant(new Cog\DataqueryConfig(
    identifier: 'custom', // datasource plugin ID
    fromArray: [CustomQuery::class, 'fromArray'],
));

$builder = (new DashboardBuilder(title: 'Custom query type'))
    ->uid('test-custom-query-type')
    ->refresh('1m')
    ->time('now-30m', 'now')
    ->withRow(new RowBuilder('Overview'))
    ->withPanel(
        (new Timeseries\PanelBuilder())
            ->title('Sample panel')
            ->withTarget(new CustomQueryBuilder('query here'))
    )
;

var_dump($builder->build());
```
