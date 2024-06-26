<?php

require_once __DIR__.'/vendor/autoload.php';

$dashboard = \App\Linux\Dashboard::generate();

echo json_encode($dashboard, JSON_PRETTY_PRINT).PHP_EOL;