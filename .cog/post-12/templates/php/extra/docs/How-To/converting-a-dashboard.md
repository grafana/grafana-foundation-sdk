# Converting a dashboard

```php
<?php

use Grafana\Foundation\Dashboard;

require_once __DIR__.'/vendor/autoload.php';

$dashboardJSON = file_get_contents('__DIR__.'/dashboard.json'');

$dashboard = Dashboard\Dashboard::fromArray(json_decode($dashboardJSON, true));

$converted = Dashboard\DashboardConverter::convert($dashboard);

var_dump($converted);
```
