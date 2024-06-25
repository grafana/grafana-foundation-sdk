<?php

use App\Custom;
use Grafana\Foundation\Cog;
use Grafana\Foundation\Dashboard\DashboardBuilder;
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

var_dump($builder->build());