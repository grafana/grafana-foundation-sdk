<?php

namespace App;

use Grafana\Foundation\Dashboardv2beta1 as SDKDashboard;
use Grafana\Foundation\Dashboardv2beta1\Dashboardv2beta1 as Dash;
use Grafana\Foundation\Timeseries;
use Grafana\Foundation\Testdata;

class Dashboard
{
    public static function generate(): SDKDashboard\DashboardBuilder
    {
        return (new SDKDashboard\DashboardBuilder(title: '[Example] Dashboard with tabs and rows'))
            ->description('Dashboard with tabs and rows generated with grafana-foundation-sdk')
            // Available panels
            ->element("panel-1", self::randomWalkPanel(1, 4))
            ->element("panel-2", self::randomWalkPanel(2, 5))
            ->element("panel-3", self::randomWalkPanel(3, 1))
            ->element("panel-4", self::randomWalkPanel(4, 1))
		    // Layout building
            ->layout(
                Dash::tabs()
                    ->tab(
                        Dash::tab('Tab without rows')
                            ->layout(
                                Dash::grid()
                                    ->item(Dash::gridItem("panel-1")->width(12)->height(8))
                            )
                    )
                    ->tab(
                        Dash::tab('Tab with rows')
                            ->layout(
                                Dash::rows()
                                    ->row(
                                        Dash::row('Row without tabs')
                                            ->collapse(false)
                                            ->layout(
                                                Dash::autoGrid()
                                                    ->item(Dash::autoGridItem('panel-3'))
                                            )
                                    )
                                    ->row(
                                        Dash::row('Row with tabs')
                                            ->collapse(true)
                                            ->layout(
                                                Dash::tabs()
                                                    ->tab(
                                                        Dash::tab('First tab')
                                                            ->layout(
                                                                Dash::autoGrid()
                                                                    ->item(Dash::autoGridItem('panel-3'))
                                                            )
                                                    )
                                                    ->tab(
                                                        Dash::tab('Second tab')
                                                            ->layout(
                                                                Dash::autoGrid()
                                                                    ->item(Dash::autoGridItem('panel-4'))
                                                            )
                                                    )
                                            )
                                    )
                                    ->row(
                                        Dash::row('Empty row')
                                            ->layout(Dash::autoGrid())
                                    )
                                    ->row(
                                        Dash::row('Hide header row')
                                            ->hideHeader(true)
                                            ->layout(Dash::autoGrid())
                                    )
                            )
                    )
            )
        ;
    }

    private static function randomWalkPanel(int $panelId, int $seriesCount): SDKDashboard\PanelBuilder
    {
        return (new SDKDashboard\PanelBuilder())
            ->title("New panel {$panelId}")
            ->visualization(new Timeseries\VisualizationBuilder())
            ->data(new SDKDashboard\QueryGroupBuilder()
                ->target(new SDKDashboard\TargetBuilder()
                    ->query(
                        new Testdata\QueryBuilder()
                            ->scenarioId(Testdata\DataqueryScenarioId::randomWalk())
                            ->seriesCount($seriesCount)
                    )
                )
            )
        ;
    }
}
