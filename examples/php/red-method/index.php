<?php

use App\Red;

require_once __DIR__.'/vendor/autoload.php';

$redDashboard = Red\Dashboard::generate(
    title: 'RED method',
    serviceIDs: ['sample-service', 'payments', 'front-gateway'],
);

echo json_encode($redDashboard, JSON_PRETTY_PRINT | JSON_UNESCAPED_SLASHES).PHP_EOL;