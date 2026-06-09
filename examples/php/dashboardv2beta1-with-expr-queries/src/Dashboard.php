<?php

namespace App;

use Grafana\Foundation\Dashboardv2beta1 as SDKDashboard;
use Grafana\Foundation\Dashboardv2beta1\Dashboardv2beta1 as Dash;
use Grafana\Foundation\Expr;
use Grafana\Foundation\Timeseries;
use Grafana\Foundation\Stat;

class Dashboard
{
    public static function generate(): SDKDashboard\DashboardBuilder
    {
        return (new SDKDashboard\DashboardBuilder(title: '[Example] Dashboard with expr queries'))
            ->editable(true)
            ->element("expr-sql", self::timeseries(
                1,
                (new Expr\QueryBuilder())->sql(
                    (new Expr\TypeSqlBuilder())->expression('select * from A'),
                ),
            ))
            ->element("expr-math", self::stat(
                2,
                (new Expr\QueryBuilder())->math(
                    (new Expr\TypeMathBuilder())->expression('2 + 4'),
                ),
            ))
		    // Layout building
            ->layout(
                Dash::autoGrid()
                    ->item(Dash::autoGridItem("expr-sql"))
                    ->item(Dash::autoGridItem("expr-math"))
            )
        ;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\DataQueryKind> $query
     */
    private static function timeseries(int $panelId, \Grafana\Foundation\Cog\Builder $query): SDKDashboard\PanelBuilder
    {
        return (new SDKDashboard\PanelBuilder())
            ->id((float)$panelId)
            ->title("Panel {$panelId}")
            ->visualization(new Timeseries\VisualizationBuilder())
            ->data((new SDKDashboard\QueryGroupBuilder())
                ->target((new SDKDashboard\TargetBuilder())
                    ->refId("target-$panelId")
                    ->query($query)
                )
            )
        ;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\DataQueryKind> $query
     */
    private static function stat(int $panelId, \Grafana\Foundation\Cog\Builder $query): SDKDashboard\PanelBuilder
    {
        return (new SDKDashboard\PanelBuilder())
            ->id((float)$panelId)
            ->title("Panel {$panelId}")
            ->visualization(new Stat\VisualizationBuilder())
            ->data((new SDKDashboard\QueryGroupBuilder())
                ->target((new SDKDashboard\TargetBuilder())
                    ->refId("target-$panelId")
                    ->query($query)
                )
            )
        ;
    }
}
