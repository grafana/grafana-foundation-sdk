<?php

namespace App\Red;

use Grafana\Foundation\Common;
use Grafana\Foundation\Common\GraphDrawStyle;
use Grafana\Foundation\Common\GraphGradientMode;
use Grafana\Foundation\Common\GraphThresholdsStyleConfigBuilder;
use Grafana\Foundation\Common\GraphThresholdsStyleMode;
use Grafana\Foundation\Common\LegendDisplayMode;
use Grafana\Foundation\Common\LegendPlacement;
use Grafana\Foundation\Common\LineInterpolation;
use Grafana\Foundation\Common\SortOrder;
use Grafana\Foundation\Common\TooltipDisplayMode;
use Grafana\Foundation\Common\VisibilityMode;
use Grafana\Foundation\Common\VizLegendOptionsBuilder;
use Grafana\Foundation\Common\VizTooltipOptionsBuilder;
use Grafana\Foundation\Dashboard as SDKDashboard;
use Grafana\Foundation\Dashboard\DashboardBuilder;
use Grafana\Foundation\Dashboard\DashboardCursorSync;
use Grafana\Foundation\Dashboard\DashboardLinkBuilder;
use Grafana\Foundation\Dashboard\DashboardLinkType;
use Grafana\Foundation\Dashboard\DataSourceRef;
use Grafana\Foundation\Dashboard\RowBuilder;
use Grafana\Foundation\Dashboard\TimePickerBuilder;
use Grafana\Foundation\Testdata;
use Grafana\Foundation\Timeseries;
use Grafana\Foundation\Units\Constants as Units;

class Dashboard
{
    /**
     * @param string[] $serviceIDs
     */
    public static function generate(string $title, array $serviceIDs): SDKDashboard\Dashboard
    {
        $builder = (new DashboardBuilder(title: '[Example] '.$title))
            ->uid('example-red-method')
            ->tags(['generated', 'red'])
            ->editable()
            ->tooltip(DashboardCursorSync::crosshair())
            ->refresh('30s')
            ->time('now-30m', 'now')
            ->timezone(Common\Constants::TIME_ZONE_BROWSER)
            ->timepicker(
                (new TimePickerBuilder())
                    ->refreshIntervals(['5s', '10s', '30s', '1m', '5m', '15m', '30m', '1h', '2h', '1d'])
            )
		    // More info about the RED method
            ->link(
                (new DashboardLinkBuilder('RED method'))
                    ->type(DashboardLinkType::link())
                    ->icon('question')
                    ->targetBlank(true)
                    ->url('https://grafana.com/blog/2018/08/02/the-red-method-how-to-instrument-your-services/#the-red-method')
            )
            ->withRow(new RowBuilder('Overview'))
        ;

        foreach ($serviceIDs as $serviceID) {
            $builder = $builder
                ->withRow(new RowBuilder($serviceID))
                ->withPanel(self::requestRateTimeseries()->span(8)->height(8))
                ->withPanel(self::errorRateTimeseries()->span(8)->height(8))
                ->withPanel(self::durationTimeseries()->span(8)->height(8))
            ;
        }

        return $builder->build();
    }

    private static function requestRateTimeseries(): Timeseries\PanelBuilder
    {
        return self::defaultTimeseries()
            ->title('Request rate')
            ->description('Number of requests handled by the service, per second.')
            ->unit(Units::REQUESTS_PER_SECOND)
            ->withTarget(self::randomData())
        ;
    }

    private static function errorRateTimeseries(): Timeseries\PanelBuilder
    {
        return self::defaultTimeseries()
            ->title('Error rate')
            ->description('Number of failed requests, per second.')
            ->unit(Units::REQUESTS_PER_SECOND)
            ->withTarget(self::randomData())
        ;
    }

    private static function durationTimeseries(): Timeseries\PanelBuilder
    {
        return self::defaultTimeseries()
            ->title('Duration')
            ->description('Time taken to process the requests.')
            ->unit(Units::SECONDS)
            ->withTarget(self::randomData())
        ;
    }

    private static function randomData(): Testdata\DataqueryBuilder
    {
        return (new Testdata\DataqueryBuilder())
            ->queryType('randomWalk')
            ->datasource(
                new DataSourceRef(type: 'grafana', uid: 'grafana')
            )
        ;
    }

    private static function defaultTimeseries(): Timeseries\PanelBuilder
    {
        return (new Timeseries\PanelBuilder())
            ->lineWidth(1)
            ->fillOpacity(30)
            ->showPoints(VisibilityMode::never())
            ->drawStyle(GraphDrawStyle::line())
            ->gradientMode(GraphGradientMode::opacity())
            ->spanNulls(false)
            ->axisBorderShow(false)
            ->lineInterpolation(LineInterpolation::smooth())
            ->legend(
                (new VizLegendOptionsBuilder())
                    ->showLegend(true)
                    ->placement(LegendPlacement::bottom())
                    ->displayMode(LegendDisplayMode::list())
            )
            ->tooltip(
                (new VizTooltipOptionsBuilder())
                    ->mode(TooltipDisplayMode::multi())
                    ->sort(SortOrder::descending())
            )
            ->thresholdsStyle(
                (new GraphThresholdsStyleConfigBuilder())
                    ->mode(GraphThresholdsStyleMode::off())
            )
        ;
    }
}
