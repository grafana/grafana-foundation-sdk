<?php

namespace App;

use Grafana\Foundation\Dashboardv2beta1 as SDKDashboard;
use Grafana\Foundation\Timeseries;
use Grafana\Foundation\Testdata;

class Dashboard
{
    public static function generate(): SDKDashboard\DashboardBuilder
    {
        return (new SDKDashboard\DashboardBuilder(title: '[Example] Dashboard with tabs and rows'))
            ->description('Dashboard with tabs and rows generated with grafana-foundation-sdk')
            // Available panels
            ->panel("panel-1", self::randomWalkPanel(1, 4))
            ->panel("panel-2", self::randomWalkPanel(2, 5))
            ->panel("panel-3", self::randomWalkPanel(3, 1))
            ->panel("panel-4", self::randomWalkPanel(4, 1))
		    // Layout building
            ->tabsLayout(
                (new SDKDashboard\TabsLayoutBuilder())
                    ->tab(
                        new SDKDashboard\TabsLayoutTabBuilder('Tab without rows')
                            ->gridLayout(
                                new SDKDashboard\GridLayoutBuilder()
                                    ->item(new SDKDashboard\GridLayoutItemBuilder("panel-1")->width(12)->height(8))
                            )
                    )
                    ->tab(
                        new SDKDashboard\TabsLayoutTabBuilder('Tab with rows')
                            ->rowsLayout(
                                (new SDKDashboard\RowsLayoutBuilder())
                                    ->row(
                                        new SDKDashboard\RowsLayoutRowBuilder('Row without tabs')
                                            ->collapse(false)
                                            ->autoGridLayout(
                                                new SDKDashboard\AutoGridLayoutBuilder()
                                                    ->item(new SDKDashboard\AutoGridLayoutItemBuilder('panel-3'))
                                            )
                                    )
                                    ->row(
                                        new SDKDashboard\RowsLayoutRowBuilder('Row with tabs')
                                            ->collapse(true)
                                            ->tabsLayout(
                                                new SDKDashboard\TabsLayoutBuilder()
                                                    ->tab(
                                                        new SDKDashboard\TabsLayoutTabBuilder('First tab')
                                                            ->autoGridLayout(
                                                                new SDKDashboard\AutoGridLayoutBuilder()
                                                                    ->item(new SDKDashboard\AutoGridLayoutItemBuilder('panel-3'))
                                                            )
                                                    )
                                                    ->tab(
                                                        new SDKDashboard\TabsLayoutTabBuilder('Second tab')
                                                            ->autoGridLayout(
                                                                new SDKDashboard\AutoGridLayoutBuilder()
                                                                    ->item(new SDKDashboard\AutoGridLayoutItemBuilder('panel-4'))
                                                            )
                                                    )
                                            )
                                    )
                                    ->row(
                                        new SDKDashboard\RowsLayoutRowBuilder('Empty row')
                                            ->autoGridLayout(new SDKDashboard\AutoGridLayoutBuilder())
                                    )
                                    ->row(
                                        new SDKDashboard\RowsLayoutRowBuilder('Hide header row')
                                            ->hideHeader(true)
                                            ->autoGridLayout(new SDKDashboard\AutoGridLayoutBuilder())
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
