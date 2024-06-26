<?php

require_once __DIR__.'/vendor/autoload.php';

$dashboard = \App\Agent\Dashboard::generate();

echo json_encode($dashboard, JSON_PRETTY_PRINT).PHP_EOL;