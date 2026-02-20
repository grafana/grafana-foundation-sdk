<?php

use Grafana\Foundation\Dashboardv2beta1;

require_once __DIR__.'/vendor/autoload.php';

$manifest = Dashboardv2beta1\Dashboardv2beta1::manifest(
    name: 'example-dashboard-with-tabs-and-rows',
    dashboard: App\Dashboard::generate(),
);

echo json_encode($manifest->build(), JSON_PRETTY_PRINT).PHP_EOL;
