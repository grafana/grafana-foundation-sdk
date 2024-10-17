<?php

namespace App\Linux;

use Grafana\Foundation\Common as SDKCommon;
use Grafana\Foundation\Dashboard as SDKDashboard;
use Grafana\Foundation\Prometheus;
use Grafana\Foundation\Stat;
use Grafana\Foundation\Timeseries;

class Common
{
    public static function queryVariable(string $name, string $label, string $query): SDKDashboard\QueryVariableBuilder
    {
        return (new SDKDashboard\QueryVariableBuilder($name))
            ->label($label)
            ->query($query)
            ->datasource(new SDKDashboard\DataSourceRef(uid: '$datasource'))
            ->current(
                new SDKDashboard\VariableOption(
                    selected: true,
                    text: 'All',
                    value: '$__all',
                )
            )
            ->refresh(SDKDashboard\VariableRefresh::onTimeRangeChanged())
            ->sort(SDKDashboard\VariableSort::alphabeticalAsc())
            ->multi(true)
            ->includeAll(true)
        ;
    }

    public static function defaultStat(): Stat\PanelBuilder
    {
        return (new Stat\PanelBuilder())
            ->datasource(new SDKDashboard\DataSourceRef(uid: '$datasource'))
            ->height(2)
            ->span(6)
            ->decimals(1)
            ->reduceOptions((new SDKCommon\ReduceDataOptionsBuilder())->calcs(['lastNotNull']))
            ->colorMode(SDKCommon\BigValueColorMode::value())
            ->graphMode(SDKCommon\BigValueGraphMode::none())
        ;
    }

    public static function defaultTimeseries(): Timeseries\PanelBuilder
    {
        return (new Timeseries\PanelBuilder())
            ->datasource(new SDKDashboard\DataSourceRef(uid: '$datasource'))
            ->lineWidth(2)
            ->fillOpacity(30)
            ->lineInterpolation(SDKCommon\LineInterpolation::smooth())
            ->showPoints(SDKCommon\VisibilityMode::never())
            ->drawStyle(SDKCommon\GraphDrawStyle::line())
            ->gradientMode(SDKCommon\GraphGradientMode::scheme())
            ->legend(
                (new SDKCommon\VizLegendOptionsBuilder())
                    ->showLegend(true)
                    ->placement(SDKCommon\LegendPlacement::bottom())
                    ->displayMode(SDKCommon\LegendDisplayMode::list())
            )
            ->tooltip(
                (new SDKCommon\VizTooltipOptionsBuilder())
                    ->mode(SDKCommon\TooltipDisplayMode::multi())
                    ->sort(SDKCommon\SortOrder::descending())
            )
            ->thresholdsStyle(
                (new SDKCommon\GraphThresholdsStyleConfigBuilder())
                    ->mode(SDKCommon\GraphThresholdsStyleMode::off())
            )
            ->spanNulls(false)
            ->axisBorderShow(false)
        ;
    }

    public static function unameStat(string $panelTitle, string $panelDescription, string $field): Stat\PanelBuilder
    {
        return self::defaultStat()
            ->title($panelTitle)
            ->description($panelDescription)
            ->withTarget(
                self::tablePrometheusQuery('node_uname_info{job=~"integrations/(node_exporter|unix)",cluster=~"$cluster",job=~"$job",instance=~"$instance"}', 'A')
            )
            ->reduceOptions(
                (new SDKCommon\ReduceDataOptionsBuilder())
                    ->calcs(['lastNotNull'])
                    ->fields('/^'.$field.'$/')
            )
            ->colorMode(SDKCommon\BigValueColorMode::none())
        ;
    }

    public static function prometheusQuery(string $query, string $legend): Prometheus\DataqueryBuilder
    {
        return (new Prometheus\DataqueryBuilder())
            ->expr($query)
            ->legendFormat($legend)
        ;
    }

    public static function tablePrometheusQuery(string $query, string $refId): Prometheus\DataqueryBuilder
    {
        return (new Prometheus\DataqueryBuilder())
            ->expr($query)
            ->format(Prometheus\PromQueryFormat::table())
            ->refId($refId)
        ;
    }
}