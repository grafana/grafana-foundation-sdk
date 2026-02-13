<?php

use Grafana\Foundation\Resource;

require_once __DIR__.'/vendor/autoload.php';

$dashboard = \App\Linux\Dashboard::generate();

$manifest = new Resource\Manifest(
    apiVersion: 'dashboard.grafana.app/v1beta1',
    kind: Resource\Constants::DASHBOARD_KIND,
    metadata: new Resource\Metadata(name: $dashboard->uid),
    spec: $dashboard,
);

echo json_encode($manifest, JSON_PRETTY_PRINT).PHP_EOL;
