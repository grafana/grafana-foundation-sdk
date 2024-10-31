# Defining a custom panel type

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
        $this->internal->type = "custom-panel";
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
