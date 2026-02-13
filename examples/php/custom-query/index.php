<?php

use App\Custom;
use Grafana\Foundation\Cog;
use Grafana\Foundation\Dashboard\DashboardBuilder;
use Grafana\Foundation\Resource;
use Grafana\Foundation\Timeseries;

require_once __DIR__.'/vendor/autoload.php';

// This lets cog know about the newly created query type and how to unmarshal it.
Cog\Runtime::get()->registerDataqueryVariant(new Cog\DataqueryConfig(
    identifier: 'custom', // datasource plugin ID
    fromArray: [Custom\Query::class, 'fromArray'],
));

$builder = (new DashboardBuilder(title: '[Example] Custom query'))
    ->uid('example-custom-query')
    ->withPanel(
        (new Timeseries\PanelBuilder())
            ->title('Sample panel')
            ->withTarget(new Custom\QueryBuilder('query here'))
    )
;

$dashboard = $builder->build();

$manifest = new Resource\Manifest(
    apiVersion: 'dashboard.grafana.app/v1beta1',
    kind: Resource\Constants::DASHBOARD_KIND,
    metadata: new Resource\Metadata(name: $dashboard->uid),
    spec: $dashboard,
);

echo json_encode($manifest, JSON_PRETTY_PRINT).PHP_EOL;
