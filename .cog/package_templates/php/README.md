# Grafana Foundation SDK – PHP

A set of tools, types and *builder libraries* for building and manipulating Grafana objects in PHP.

> [!NOTE]
> This branch contains **types and builders generated for Grafana {{ .Extra.GrafanaVersion }}.**
> Other supported versions of Grafana can be found at [this repository's root](https://github.com/grafana/grafana-foundation-sdk/).

## Installing

```shell
composer require "grafana/foundation-sdk:dev-{{ .Extra.ReleaseBranch }}"
```

## Example usage

[More examples](https://github.com/grafana/grafana-foundation-sdk/tree/main/examples/php) can be found at the repository root.

### Building a dashboard

```php
<?php

use Grafana\Foundation\Common;
use Grafana\Foundation\Dashboard\DashboardBuilder;
use Grafana\Foundation\Dashboard\RowBuilder;
use Grafana\Foundation\Prometheus;
use Grafana\Foundation\Timeseries;

require_once __DIR__.'/vendor/autoload.php';

$builder = (new DashboardBuilder(title: 'Sample dashboard'))
    ->uid('generated-from-php')
    ->tags(['generated', 'from', 'php'])
    ->refresh('1m')
    ->time('now-30m', 'now')
    ->timezone(Common\Constants::TIME_ZONE_BROWSER)
    ->withRow(new RowBuilder('Overview'))
    ->withPanel(
        (new Timeseries\PanelBuilder())
            ->title('Network received')
            ->unit('bps')
            ->min(0)
            ->withTarget(
                (new Prometheus\DataqueryBuilder())
                    ->expr('rate(node_network_receive_bytes_total{job="integrations/raspberrypi-node", device!="lo"}[$__rate_interval]) * 8')
                    ->legendFormat({{ `'{{ device }}'` }})
            )
    )
;

echo(json_encode($builder->build(), JSON_PRETTY_PRINT).PHP_EOL);
```

### Unmarshaling a dashboard

```php
<?php

use Grafana\Foundation\Dashboard\Dashboard;

require_once __DIR__.'/vendor/autoload.php';

$dashboardJSON = file_get_contents(__DIR__.'/dashboard.json');

$dashboard = Dashboard::fromArray(json_decode($dashboardJSON, true));

var_dump($dashboard);
```

### Defining a custom query type

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

### Defining a custom panel type

While the SDK ships with support for all core panels, it can be extended for
private/third-party plugins.

To do so, define a type for the custom panel's options:

```php
<?php

class CustomPanelOptions implements \JsonSerializable
{
    public bool $makeBeautiful;

    public function __construct(?bool $makeBeautiful = null)
    {
        $this->makeBeautiful = $makeBeautiful ?: false;
    }

    public function jsonSerialize(): array
    {
        return [
            "makeBeautiful" => $this->makeBeautiful,
        ];
    }

    /**
     * @param array{makeBeautiful?: bool} $data
     */
    public static function fromArray(array $data): self
    {
        return new self(
            makeBeautiful: $data["makeBeautiful"] ?? null,
        );
    }
}
```

Now, let's define a builder for that type:

```php
<?php

use Grafana\Foundation\Cog;
use Grafana\Foundation\Dashboard;

/**
 * @implements Cog\Builder<Dashboard\Panel>
 */
class CustomPanelBuilder extends Dashboard\PanelBuilder implements Cog\Builder
{
    public function __construct()
    {
        parent::__construct();
    }

    public function makeBeautiful(): static
    {
        if ($this->internal->options === null) {
            $this->internal->options = new CustomPanelOptions();
        }
        assert($this->internal->options instanceof CustomPanelOptions);
        $this->internal->options->makeBeautiful = true;
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

// This lets cog know about the newly created panel type and how to unmarshal it.
Cog\Runtime::get()->registerPanelcfgVariant(new Cog\PanelcfgConfig(
    identifier: 'custom-panel', // panel plugin ID
    optionsFromArray: [CustomPanelOptions::class, 'fromArray'],
));

$builder = (new DashboardBuilder(title: 'Custom panel type'))
    ->uid('test-custom-panel-type')
    ->refresh('1m')
    ->time('now-30m', 'now')
    ->withRow(new RowBuilder('Overview'))
    ->withPanel(
        (new CustomPanelBuilder())
            ->title('Sample panel')
            ->makeBeautiful()
    )
;

var_dump($builder->build());
```

## Maturity

The code in this repository should be considered as "public preview". While it is used by Grafana Labs in production, it still is under active development.

> [!NOTE]
> Bugs and issues are handled solely by Engineering teams. On-call support or SLAs are not available.

## License

[Apache 2.0 License](./LICENSE)
