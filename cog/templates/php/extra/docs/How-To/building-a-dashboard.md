# Building a dashboard

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
