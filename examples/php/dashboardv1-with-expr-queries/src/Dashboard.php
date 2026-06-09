<?php

namespace App;

use Grafana\Foundation\Dashboard as SDKDashboard;
use Grafana\Foundation\Dashboard\Dashboard as Dash;
use Grafana\Foundation\Expr;
use Grafana\Foundation\Timeseries;
use Grafana\Foundation\Stat;

class Dashboard
{
    public static function generate(): SDKDashboard\DashboardBuilder
    {
        return (new SDKDashboard\DashboardBuilder(title: '[Example] Dashboard with expr queries'))
            ->editable()
            ->uid('example-dashboard-with-expr-queries')
            ->withPanel(self::timeseries(
                1,
                (new Expr\TypeSqlBuilder())->expression('select * from A'),
            ))
            ->withPanel(self::stat(
                2,
                (new Expr\TypeMathBuilder())->expression('2 + 4'),
            ))
        ;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Cog\Dataquery> $query
     */
    private static function timeseries(int $panelId, \Grafana\Foundation\Cog\Builder $query): Timeseries\PanelBuilder
    {
        return (new Timeseries\PanelBuilder())
            ->id((float)$panelId)
            ->title("Panel {$panelId}")
            ->withTarget($query)
        ;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Cog\Dataquery> $query
     */
    private static function stat(int $panelId, \Grafana\Foundation\Cog\Builder $query): Stat\PanelBuilder
    {
        return (new Stat\PanelBuilder())
            ->id((float)$panelId)
            ->title("Panel {$panelId}")
            ->withTarget($query)
        ;
    }
}
