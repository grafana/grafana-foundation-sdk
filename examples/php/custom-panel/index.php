<?php

use App\Custom;
use Grafana\Foundation\Cog;
use Grafana\Foundation\Dashboard\DashboardBuilder;
use Grafana\Foundation\Dashboard\RowBuilder;
use Grafana\Foundation\Resource;

require_once __DIR__.'/vendor/autoload.php';

// This lets cog know about the newly created panel type and how to unmarshal it.
Cog\Runtime::get()->registerPanelcfgVariant(new Cog\PanelcfgConfig(
    identifier: 'custom-panel', // panel plugin ID
    optionsFromArray: [Custom\Options::class, 'fromArray'],
));

$builder = (new DashboardBuilder(title: '[Example] Custom panel type'))
    ->uid('example-custom-panel')
    ->refresh('1m')
    ->time('now-30m', 'now')
    ->withRow(new RowBuilder('Overview'))
    ->withPanel(
        (new Custom\PanelBuilder())
            ->title('Sample panel')
            ->makeBeautiful()
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
