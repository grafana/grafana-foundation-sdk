<?php

use App\Red;
use Grafana\Foundation\Resource;

require_once __DIR__.'/vendor/autoload.php';

$redDashboard = Red\Dashboard::generate(
    title: 'RED method',
    serviceIDs: ['sample-service', 'payments', 'front-gateway'],
);

$manifest = new Resource\Manifest(
    apiVersion: 'dashboard.grafana.app/v1beta1',
    kind: Resource\Constants::DASHBOARD_KIND,
    metadata: new Resource\Metadata(name: $redDashboard->uid),
    spec: $redDashboard,
);

echo json_encode($manifest, JSON_PRETTY_PRINT).PHP_EOL;
