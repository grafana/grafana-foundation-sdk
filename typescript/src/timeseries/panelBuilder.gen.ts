// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboard from '../dashboard';
import * as timeseries from '../timeseries';
import * as common from '../common';

// Dashboard panels are the basic visualization building blocks.
export class PanelBuilder implements cog.Builder<dashboard.Panel> {
    protected readonly internal: dashboard.Panel;

    constructor() {
        this.internal = dashboard.defaultPanel();
        this.internal.type = "timeseries";
    }

    build(): dashboard.Panel {
        return this.internal;
    }

    // Unique identifier of the panel. Generated by Grafana when creating a new panel. It must be unique within a dashboard, but not globally.
    id(id: number): this {
        this.internal.id = id;
        return this;
    }

    // Depends on the panel plugin. See the plugin documentation for details.
    withTarget(targets: cog.Builder<cog.Dataquery>): this {
        if (!this.internal.targets) {
            this.internal.targets = [];
        }
        const targetsResource = targets.build();
        this.internal.targets.push(targetsResource);
        return this;
    }

    // Panel title.
    title(title: string): this {
        this.internal.title = title;
        return this;
    }

    // Panel description.
    description(description: string): this {
        this.internal.description = description;
        return this;
    }

    // Whether to display the panel without a background.
    transparent(transparent: boolean): this {
        this.internal.transparent = transparent;
        return this;
    }

    // The datasource used in all targets.
    datasource(datasource: dashboard.DataSourceRef): this {
        this.internal.datasource = datasource;
        return this;
    }

    // Panel height. The height is the number of rows from the top edge of the panel.
    height(h: number): this {
        if (!this.internal.gridPos) {
            this.internal.gridPos = dashboard.defaultGridPos();
        }
        if (!(h > 0)) {
            throw new Error("h must be > 0");
        }
        this.internal.gridPos.h = h;
        return this;
    }

    // Panel width. The width is the number of columns from the left edge of the panel.
    span(w: number): this {
        if (!this.internal.gridPos) {
            this.internal.gridPos = dashboard.defaultGridPos();
        }
        if (!(w > 0)) {
            throw new Error("w must be > 0");
        }
        if (!(w <= 24)) {
            throw new Error("w must be <= 24");
        }
        this.internal.gridPos.w = w;
        return this;
    }

    // Panel links.
    links(links: cog.Builder<dashboard.DashboardLink>[]): this {
        const linksResources = links.map(builder1 => builder1.build());
        this.internal.links = linksResources;
        return this;
    }

    // Name of template variable to repeat for.
    repeat(repeat: string): this {
        this.internal.repeat = repeat;
        return this;
    }

    // Direction to repeat in if 'repeat' is set.
    // `h` for horizontal, `v` for vertical.
    repeatDirection(repeatDirection: "h" | "v"): this {
        this.internal.repeatDirection = repeatDirection;
        return this;
    }

    // Option for repeated panels that controls max items per row
    // Only relevant for horizontally repeated panels
    maxPerRow(maxPerRow: number): this {
        this.internal.maxPerRow = maxPerRow;
        return this;
    }

    // The maximum number of data points that the panel queries are retrieving.
    maxDataPoints(maxDataPoints: number): this {
        this.internal.maxDataPoints = maxDataPoints;
        return this;
    }

    // List of transformations that are applied to the panel data before rendering.
    // When there are multiple transformations, Grafana applies them in the order they are listed.
    // Each transformation creates a result set that then passes on to the next transformation in the processing pipeline.
    withTransformation(transformations: dashboard.DataTransformerConfig): this {
        if (!this.internal.transformations) {
            this.internal.transformations = [];
        }
        this.internal.transformations.push(transformations);
        return this;
    }

    // The min time interval setting defines a lower limit for the $__interval and $__interval_ms variables.
    // This value must be formatted as a number followed by a valid time
    // identifier like: "40s", "3d", etc.
    // See: https://grafana.com/docs/grafana/latest/panels-visualizations/query-transform-data/#query-options
    interval(interval: string): this {
        this.internal.interval = interval;
        return this;
    }

    // Overrides the relative time range for individual panels,
    // which causes them to be different than what is selected in
    // the dashboard time picker in the top-right corner of the dashboard. You can use this to show metrics from different
    // time periods or days on the same dashboard.
    // The value is formatted as time operation like: `now-5m` (Last 5 minutes), `now/d` (the day so far),
    // `now-5d/d`(Last 5 days), `now/w` (This week so far), `now-2y/y` (Last 2 years).
    // Note: Panel time overrides have no effect when the dashboard’s time range is absolute.
    // See: https://grafana.com/docs/grafana/latest/panels-visualizations/query-transform-data/#query-options
    timeFrom(timeFrom: string): this {
        this.internal.timeFrom = timeFrom;
        return this;
    }

    // Overrides the time range for individual panels by shifting its start and end relative to the time picker.
    // For example, you can shift the time range for the panel to be two hours earlier than the dashboard time picker setting `2h`.
    // Note: Panel time overrides have no effect when the dashboard’s time range is absolute.
    // See: https://grafana.com/docs/grafana/latest/panels-visualizations/query-transform-data/#query-options
    timeShift(timeShift: string): this {
        this.internal.timeShift = timeShift;
        return this;
    }

    // Controls if the timeFrom or timeShift overrides are shown in the panel header
    hideTimeOverride(hideTimeOverride: boolean): this {
        this.internal.hideTimeOverride = hideTimeOverride;
        return this;
    }

    // Dynamically load the panel
    libraryPanel(libraryPanel: dashboard.LibraryPanelRef): this {
        this.internal.libraryPanel = libraryPanel;
        return this;
    }

    // The display value for this field.  This supports template variables blank is auto
    displayName(displayName: string): this {
        if (!this.internal.fieldConfig) {
            this.internal.fieldConfig = dashboard.defaultFieldConfigSource();
        }
        if (!this.internal.fieldConfig.defaults) {
            this.internal.fieldConfig.defaults = dashboard.defaultFieldConfig();
        }
        this.internal.fieldConfig.defaults.displayName = displayName;
        return this;
    }

    // Unit a field should use. The unit you select is applied to all fields except time.
    // You can use the units ID availables in Grafana or a custom unit.
    // Available units in Grafana: https://github.com/grafana/grafana/blob/main/packages/grafana-data/src/valueFormats/categories.ts
    // As custom unit, you can use the following formats:
    // `suffix:<suffix>` for custom unit that should go after value.
    // `prefix:<prefix>` for custom unit that should go before value.
    // `time:<format>` For custom date time formats type for example `time:YYYY-MM-DD`.
    // `si:<base scale><unit characters>` for custom SI units. For example: `si: mF`. This one is a bit more advanced as you can specify both a unit and the source data scale. So if your source data is represented as milli (thousands of) something prefix the unit with that SI scale character.
    // `count:<unit>` for a custom count unit.
    // `currency:<unit>` for custom a currency unit.
    unit(unit: string): this {
        if (!this.internal.fieldConfig) {
            this.internal.fieldConfig = dashboard.defaultFieldConfigSource();
        }
        if (!this.internal.fieldConfig.defaults) {
            this.internal.fieldConfig.defaults = dashboard.defaultFieldConfig();
        }
        this.internal.fieldConfig.defaults.unit = unit;
        return this;
    }

    // Specify the number of decimals Grafana includes in the rendered value.
    // If you leave this field blank, Grafana automatically truncates the number of decimals based on the value.
    // For example 1.1234 will display as 1.12 and 100.456 will display as 100.
    // To display all decimals, set the unit to `String`.
    decimals(decimals: number): this {
        if (!this.internal.fieldConfig) {
            this.internal.fieldConfig = dashboard.defaultFieldConfigSource();
        }
        if (!this.internal.fieldConfig.defaults) {
            this.internal.fieldConfig.defaults = dashboard.defaultFieldConfig();
        }
        this.internal.fieldConfig.defaults.decimals = decimals;
        return this;
    }

    // The minimum value used in percentage threshold calculations. Leave blank for auto calculation based on all series and fields.
    min(min: number): this {
        if (!this.internal.fieldConfig) {
            this.internal.fieldConfig = dashboard.defaultFieldConfigSource();
        }
        if (!this.internal.fieldConfig.defaults) {
            this.internal.fieldConfig.defaults = dashboard.defaultFieldConfig();
        }
        this.internal.fieldConfig.defaults.min = min;
        return this;
    }

    // The maximum value used in percentage threshold calculations. Leave blank for auto calculation based on all series and fields.
    max(max: number): this {
        if (!this.internal.fieldConfig) {
            this.internal.fieldConfig = dashboard.defaultFieldConfigSource();
        }
        if (!this.internal.fieldConfig.defaults) {
            this.internal.fieldConfig.defaults = dashboard.defaultFieldConfig();
        }
        this.internal.fieldConfig.defaults.max = max;
        return this;
    }

    // Convert input values into a display string
    mappings(mappings: dashboard.ValueMapping[]): this {
        if (!this.internal.fieldConfig) {
            this.internal.fieldConfig = dashboard.defaultFieldConfigSource();
        }
        if (!this.internal.fieldConfig.defaults) {
            this.internal.fieldConfig.defaults = dashboard.defaultFieldConfig();
        }
        this.internal.fieldConfig.defaults.mappings = mappings;
        return this;
    }

    // Map numeric values to states
    thresholds(thresholds: cog.Builder<dashboard.ThresholdsConfig>): this {
        if (!this.internal.fieldConfig) {
            this.internal.fieldConfig = dashboard.defaultFieldConfigSource();
        }
        if (!this.internal.fieldConfig.defaults) {
            this.internal.fieldConfig.defaults = dashboard.defaultFieldConfig();
        }
        const thresholdsResource = thresholds.build();
        this.internal.fieldConfig.defaults.thresholds = thresholdsResource;
        return this;
    }

    // Alternative to empty string
    noValue(noValue: string): this {
        if (!this.internal.fieldConfig) {
            this.internal.fieldConfig = dashboard.defaultFieldConfigSource();
        }
        if (!this.internal.fieldConfig.defaults) {
            this.internal.fieldConfig.defaults = dashboard.defaultFieldConfig();
        }
        this.internal.fieldConfig.defaults.noValue = noValue;
        return this;
    }

    // Overrides are the options applied to specific fields overriding the defaults.
    withOverride(overrides: {
	matcher: dashboard.MatcherConfig;
	properties: dashboard.DynamicConfigValue[];
}): this {
        if (!this.internal.fieldConfig) {
            this.internal.fieldConfig = dashboard.defaultFieldConfigSource();
        }
        if (!this.internal.fieldConfig.overrides) {
            this.internal.fieldConfig.overrides = [];
        }
        this.internal.fieldConfig.overrides.push(overrides);
        return this;
    }

    timezone(timezone: common.TimeZone[]): this {
        if (!this.internal.options) {
            this.internal.options = timeseries.defaultOptions();
        }
        this.internal.options.timezone = timezone;
        return this;
    }

    legend(legend: cog.Builder<common.VizLegendOptions>): this {
        if (!this.internal.options) {
            this.internal.options = timeseries.defaultOptions();
        }
        const legendResource = legend.build();
        this.internal.options.legend = legendResource;
        return this;
    }

    tooltip(tooltip: cog.Builder<common.VizTooltipOptions>): this {
        if (!this.internal.options) {
            this.internal.options = timeseries.defaultOptions();
        }
        const tooltipResource = tooltip.build();
        this.internal.options.tooltip = tooltipResource;
        return this;
    }

    drawStyle(drawStyle: common.GraphDrawStyle): this {
        if (!this.internal.fieldConfig) {
            this.internal.fieldConfig = dashboard.defaultFieldConfigSource();
        }
        if (!this.internal.fieldConfig.defaults) {
            this.internal.fieldConfig.defaults = dashboard.defaultFieldConfig();
        }
        if (!this.internal.fieldConfig.defaults.custom) {
            this.internal.fieldConfig.defaults.custom = timeseries.defaultFieldConfig();
        }
        this.internal.fieldConfig.defaults.custom.drawStyle = drawStyle;
        return this;
    }

    gradientMode(gradientMode: common.GraphGradientMode): this {
        if (!this.internal.fieldConfig) {
            this.internal.fieldConfig = dashboard.defaultFieldConfigSource();
        }
        if (!this.internal.fieldConfig.defaults) {
            this.internal.fieldConfig.defaults = dashboard.defaultFieldConfig();
        }
        if (!this.internal.fieldConfig.defaults.custom) {
            this.internal.fieldConfig.defaults.custom = timeseries.defaultFieldConfig();
        }
        this.internal.fieldConfig.defaults.custom.gradientMode = gradientMode;
        return this;
    }

    thresholdsStyle(thresholdsStyle: cog.Builder<common.GraphThresholdsStyleConfig>): this {
        if (!this.internal.fieldConfig) {
            this.internal.fieldConfig = dashboard.defaultFieldConfigSource();
        }
        if (!this.internal.fieldConfig.defaults) {
            this.internal.fieldConfig.defaults = dashboard.defaultFieldConfig();
        }
        if (!this.internal.fieldConfig.defaults.custom) {
            this.internal.fieldConfig.defaults.custom = timeseries.defaultFieldConfig();
        }
        const thresholdsStyleResource = thresholdsStyle.build();
        this.internal.fieldConfig.defaults.custom.thresholdsStyle = thresholdsStyleResource;
        return this;
    }

    lineColor(lineColor: string): this {
        if (!this.internal.fieldConfig) {
            this.internal.fieldConfig = dashboard.defaultFieldConfigSource();
        }
        if (!this.internal.fieldConfig.defaults) {
            this.internal.fieldConfig.defaults = dashboard.defaultFieldConfig();
        }
        if (!this.internal.fieldConfig.defaults.custom) {
            this.internal.fieldConfig.defaults.custom = timeseries.defaultFieldConfig();
        }
        this.internal.fieldConfig.defaults.custom.lineColor = lineColor;
        return this;
    }

    lineWidth(lineWidth: number): this {
        if (!this.internal.fieldConfig) {
            this.internal.fieldConfig = dashboard.defaultFieldConfigSource();
        }
        if (!this.internal.fieldConfig.defaults) {
            this.internal.fieldConfig.defaults = dashboard.defaultFieldConfig();
        }
        if (!this.internal.fieldConfig.defaults.custom) {
            this.internal.fieldConfig.defaults.custom = timeseries.defaultFieldConfig();
        }
        this.internal.fieldConfig.defaults.custom.lineWidth = lineWidth;
        return this;
    }

    lineInterpolation(lineInterpolation: common.LineInterpolation): this {
        if (!this.internal.fieldConfig) {
            this.internal.fieldConfig = dashboard.defaultFieldConfigSource();
        }
        if (!this.internal.fieldConfig.defaults) {
            this.internal.fieldConfig.defaults = dashboard.defaultFieldConfig();
        }
        if (!this.internal.fieldConfig.defaults.custom) {
            this.internal.fieldConfig.defaults.custom = timeseries.defaultFieldConfig();
        }
        this.internal.fieldConfig.defaults.custom.lineInterpolation = lineInterpolation;
        return this;
    }

    lineStyle(lineStyle: cog.Builder<common.LineStyle>): this {
        if (!this.internal.fieldConfig) {
            this.internal.fieldConfig = dashboard.defaultFieldConfigSource();
        }
        if (!this.internal.fieldConfig.defaults) {
            this.internal.fieldConfig.defaults = dashboard.defaultFieldConfig();
        }
        if (!this.internal.fieldConfig.defaults.custom) {
            this.internal.fieldConfig.defaults.custom = timeseries.defaultFieldConfig();
        }
        const lineStyleResource = lineStyle.build();
        this.internal.fieldConfig.defaults.custom.lineStyle = lineStyleResource;
        return this;
    }

    fillColor(fillColor: string): this {
        if (!this.internal.fieldConfig) {
            this.internal.fieldConfig = dashboard.defaultFieldConfigSource();
        }
        if (!this.internal.fieldConfig.defaults) {
            this.internal.fieldConfig.defaults = dashboard.defaultFieldConfig();
        }
        if (!this.internal.fieldConfig.defaults.custom) {
            this.internal.fieldConfig.defaults.custom = timeseries.defaultFieldConfig();
        }
        this.internal.fieldConfig.defaults.custom.fillColor = fillColor;
        return this;
    }

    fillOpacity(fillOpacity: number): this {
        if (!this.internal.fieldConfig) {
            this.internal.fieldConfig = dashboard.defaultFieldConfigSource();
        }
        if (!this.internal.fieldConfig.defaults) {
            this.internal.fieldConfig.defaults = dashboard.defaultFieldConfig();
        }
        if (!this.internal.fieldConfig.defaults.custom) {
            this.internal.fieldConfig.defaults.custom = timeseries.defaultFieldConfig();
        }
        this.internal.fieldConfig.defaults.custom.fillOpacity = fillOpacity;
        return this;
    }

    showPoints(showPoints: common.VisibilityMode): this {
        if (!this.internal.fieldConfig) {
            this.internal.fieldConfig = dashboard.defaultFieldConfigSource();
        }
        if (!this.internal.fieldConfig.defaults) {
            this.internal.fieldConfig.defaults = dashboard.defaultFieldConfig();
        }
        if (!this.internal.fieldConfig.defaults.custom) {
            this.internal.fieldConfig.defaults.custom = timeseries.defaultFieldConfig();
        }
        this.internal.fieldConfig.defaults.custom.showPoints = showPoints;
        return this;
    }

    pointSize(pointSize: number): this {
        if (!this.internal.fieldConfig) {
            this.internal.fieldConfig = dashboard.defaultFieldConfigSource();
        }
        if (!this.internal.fieldConfig.defaults) {
            this.internal.fieldConfig.defaults = dashboard.defaultFieldConfig();
        }
        if (!this.internal.fieldConfig.defaults.custom) {
            this.internal.fieldConfig.defaults.custom = timeseries.defaultFieldConfig();
        }
        this.internal.fieldConfig.defaults.custom.pointSize = pointSize;
        return this;
    }

    pointColor(pointColor: string): this {
        if (!this.internal.fieldConfig) {
            this.internal.fieldConfig = dashboard.defaultFieldConfigSource();
        }
        if (!this.internal.fieldConfig.defaults) {
            this.internal.fieldConfig.defaults = dashboard.defaultFieldConfig();
        }
        if (!this.internal.fieldConfig.defaults.custom) {
            this.internal.fieldConfig.defaults.custom = timeseries.defaultFieldConfig();
        }
        this.internal.fieldConfig.defaults.custom.pointColor = pointColor;
        return this;
    }

    axisPlacement(axisPlacement: common.AxisPlacement): this {
        if (!this.internal.fieldConfig) {
            this.internal.fieldConfig = dashboard.defaultFieldConfigSource();
        }
        if (!this.internal.fieldConfig.defaults) {
            this.internal.fieldConfig.defaults = dashboard.defaultFieldConfig();
        }
        if (!this.internal.fieldConfig.defaults.custom) {
            this.internal.fieldConfig.defaults.custom = timeseries.defaultFieldConfig();
        }
        this.internal.fieldConfig.defaults.custom.axisPlacement = axisPlacement;
        return this;
    }

    axisColorMode(axisColorMode: common.AxisColorMode): this {
        if (!this.internal.fieldConfig) {
            this.internal.fieldConfig = dashboard.defaultFieldConfigSource();
        }
        if (!this.internal.fieldConfig.defaults) {
            this.internal.fieldConfig.defaults = dashboard.defaultFieldConfig();
        }
        if (!this.internal.fieldConfig.defaults.custom) {
            this.internal.fieldConfig.defaults.custom = timeseries.defaultFieldConfig();
        }
        this.internal.fieldConfig.defaults.custom.axisColorMode = axisColorMode;
        return this;
    }

    axisLabel(axisLabel: string): this {
        if (!this.internal.fieldConfig) {
            this.internal.fieldConfig = dashboard.defaultFieldConfigSource();
        }
        if (!this.internal.fieldConfig.defaults) {
            this.internal.fieldConfig.defaults = dashboard.defaultFieldConfig();
        }
        if (!this.internal.fieldConfig.defaults.custom) {
            this.internal.fieldConfig.defaults.custom = timeseries.defaultFieldConfig();
        }
        this.internal.fieldConfig.defaults.custom.axisLabel = axisLabel;
        return this;
    }

    axisWidth(axisWidth: number): this {
        if (!this.internal.fieldConfig) {
            this.internal.fieldConfig = dashboard.defaultFieldConfigSource();
        }
        if (!this.internal.fieldConfig.defaults) {
            this.internal.fieldConfig.defaults = dashboard.defaultFieldConfig();
        }
        if (!this.internal.fieldConfig.defaults.custom) {
            this.internal.fieldConfig.defaults.custom = timeseries.defaultFieldConfig();
        }
        this.internal.fieldConfig.defaults.custom.axisWidth = axisWidth;
        return this;
    }

    axisSoftMin(axisSoftMin: number): this {
        if (!this.internal.fieldConfig) {
            this.internal.fieldConfig = dashboard.defaultFieldConfigSource();
        }
        if (!this.internal.fieldConfig.defaults) {
            this.internal.fieldConfig.defaults = dashboard.defaultFieldConfig();
        }
        if (!this.internal.fieldConfig.defaults.custom) {
            this.internal.fieldConfig.defaults.custom = timeseries.defaultFieldConfig();
        }
        this.internal.fieldConfig.defaults.custom.axisSoftMin = axisSoftMin;
        return this;
    }

    axisSoftMax(axisSoftMax: number): this {
        if (!this.internal.fieldConfig) {
            this.internal.fieldConfig = dashboard.defaultFieldConfigSource();
        }
        if (!this.internal.fieldConfig.defaults) {
            this.internal.fieldConfig.defaults = dashboard.defaultFieldConfig();
        }
        if (!this.internal.fieldConfig.defaults.custom) {
            this.internal.fieldConfig.defaults.custom = timeseries.defaultFieldConfig();
        }
        this.internal.fieldConfig.defaults.custom.axisSoftMax = axisSoftMax;
        return this;
    }

    axisGridShow(axisGridShow: boolean): this {
        if (!this.internal.fieldConfig) {
            this.internal.fieldConfig = dashboard.defaultFieldConfigSource();
        }
        if (!this.internal.fieldConfig.defaults) {
            this.internal.fieldConfig.defaults = dashboard.defaultFieldConfig();
        }
        if (!this.internal.fieldConfig.defaults.custom) {
            this.internal.fieldConfig.defaults.custom = timeseries.defaultFieldConfig();
        }
        this.internal.fieldConfig.defaults.custom.axisGridShow = axisGridShow;
        return this;
    }

    scaleDistribution(scaleDistribution: cog.Builder<common.ScaleDistributionConfig>): this {
        if (!this.internal.fieldConfig) {
            this.internal.fieldConfig = dashboard.defaultFieldConfigSource();
        }
        if (!this.internal.fieldConfig.defaults) {
            this.internal.fieldConfig.defaults = dashboard.defaultFieldConfig();
        }
        if (!this.internal.fieldConfig.defaults.custom) {
            this.internal.fieldConfig.defaults.custom = timeseries.defaultFieldConfig();
        }
        const scaleDistributionResource = scaleDistribution.build();
        this.internal.fieldConfig.defaults.custom.scaleDistribution = scaleDistributionResource;
        return this;
    }

    axisCenteredZero(axisCenteredZero: boolean): this {
        if (!this.internal.fieldConfig) {
            this.internal.fieldConfig = dashboard.defaultFieldConfigSource();
        }
        if (!this.internal.fieldConfig.defaults) {
            this.internal.fieldConfig.defaults = dashboard.defaultFieldConfig();
        }
        if (!this.internal.fieldConfig.defaults.custom) {
            this.internal.fieldConfig.defaults.custom = timeseries.defaultFieldConfig();
        }
        this.internal.fieldConfig.defaults.custom.axisCenteredZero = axisCenteredZero;
        return this;
    }

    barAlignment(barAlignment: common.BarAlignment): this {
        if (!this.internal.fieldConfig) {
            this.internal.fieldConfig = dashboard.defaultFieldConfigSource();
        }
        if (!this.internal.fieldConfig.defaults) {
            this.internal.fieldConfig.defaults = dashboard.defaultFieldConfig();
        }
        if (!this.internal.fieldConfig.defaults.custom) {
            this.internal.fieldConfig.defaults.custom = timeseries.defaultFieldConfig();
        }
        this.internal.fieldConfig.defaults.custom.barAlignment = barAlignment;
        return this;
    }

    barWidthFactor(barWidthFactor: number): this {
        if (!this.internal.fieldConfig) {
            this.internal.fieldConfig = dashboard.defaultFieldConfigSource();
        }
        if (!this.internal.fieldConfig.defaults) {
            this.internal.fieldConfig.defaults = dashboard.defaultFieldConfig();
        }
        if (!this.internal.fieldConfig.defaults.custom) {
            this.internal.fieldConfig.defaults.custom = timeseries.defaultFieldConfig();
        }
        this.internal.fieldConfig.defaults.custom.barWidthFactor = barWidthFactor;
        return this;
    }

    stacking(stacking: cog.Builder<common.StackingConfig>): this {
        if (!this.internal.fieldConfig) {
            this.internal.fieldConfig = dashboard.defaultFieldConfigSource();
        }
        if (!this.internal.fieldConfig.defaults) {
            this.internal.fieldConfig.defaults = dashboard.defaultFieldConfig();
        }
        if (!this.internal.fieldConfig.defaults.custom) {
            this.internal.fieldConfig.defaults.custom = timeseries.defaultFieldConfig();
        }
        const stackingResource = stacking.build();
        this.internal.fieldConfig.defaults.custom.stacking = stackingResource;
        return this;
    }

    hideFrom(hideFrom: cog.Builder<common.HideSeriesConfig>): this {
        if (!this.internal.fieldConfig) {
            this.internal.fieldConfig = dashboard.defaultFieldConfigSource();
        }
        if (!this.internal.fieldConfig.defaults) {
            this.internal.fieldConfig.defaults = dashboard.defaultFieldConfig();
        }
        if (!this.internal.fieldConfig.defaults.custom) {
            this.internal.fieldConfig.defaults.custom = timeseries.defaultFieldConfig();
        }
        const hideFromResource = hideFrom.build();
        this.internal.fieldConfig.defaults.custom.hideFrom = hideFromResource;
        return this;
    }

    transform(transform: common.GraphTransform): this {
        if (!this.internal.fieldConfig) {
            this.internal.fieldConfig = dashboard.defaultFieldConfigSource();
        }
        if (!this.internal.fieldConfig.defaults) {
            this.internal.fieldConfig.defaults = dashboard.defaultFieldConfig();
        }
        if (!this.internal.fieldConfig.defaults.custom) {
            this.internal.fieldConfig.defaults.custom = timeseries.defaultFieldConfig();
        }
        this.internal.fieldConfig.defaults.custom.transform = transform;
        return this;
    }

    // Indicate if null values should be treated as gaps or connected.
    // When the value is a number, it represents the maximum delta in the
    // X axis that should be considered connected.  For timeseries, this is milliseconds
    spanNulls(spanNulls: boolean | number): this {
        if (!this.internal.fieldConfig) {
            this.internal.fieldConfig = dashboard.defaultFieldConfigSource();
        }
        if (!this.internal.fieldConfig.defaults) {
            this.internal.fieldConfig.defaults = dashboard.defaultFieldConfig();
        }
        if (!this.internal.fieldConfig.defaults.custom) {
            this.internal.fieldConfig.defaults.custom = timeseries.defaultFieldConfig();
        }
        this.internal.fieldConfig.defaults.custom.spanNulls = spanNulls;
        return this;
    }

    fillBelowTo(fillBelowTo: string): this {
        if (!this.internal.fieldConfig) {
            this.internal.fieldConfig = dashboard.defaultFieldConfigSource();
        }
        if (!this.internal.fieldConfig.defaults) {
            this.internal.fieldConfig.defaults = dashboard.defaultFieldConfig();
        }
        if (!this.internal.fieldConfig.defaults.custom) {
            this.internal.fieldConfig.defaults.custom = timeseries.defaultFieldConfig();
        }
        this.internal.fieldConfig.defaults.custom.fillBelowTo = fillBelowTo;
        return this;
    }

    pointSymbol(pointSymbol: string): this {
        if (!this.internal.fieldConfig) {
            this.internal.fieldConfig = dashboard.defaultFieldConfigSource();
        }
        if (!this.internal.fieldConfig.defaults) {
            this.internal.fieldConfig.defaults = dashboard.defaultFieldConfig();
        }
        if (!this.internal.fieldConfig.defaults.custom) {
            this.internal.fieldConfig.defaults.custom = timeseries.defaultFieldConfig();
        }
        this.internal.fieldConfig.defaults.custom.pointSymbol = pointSymbol;
        return this;
    }

    axisBorderShow(axisBorderShow: boolean): this {
        if (!this.internal.fieldConfig) {
            this.internal.fieldConfig = dashboard.defaultFieldConfigSource();
        }
        if (!this.internal.fieldConfig.defaults) {
            this.internal.fieldConfig.defaults = dashboard.defaultFieldConfig();
        }
        if (!this.internal.fieldConfig.defaults.custom) {
            this.internal.fieldConfig.defaults.custom = timeseries.defaultFieldConfig();
        }
        this.internal.fieldConfig.defaults.custom.axisBorderShow = axisBorderShow;
        return this;
    }

    barMaxWidth(barMaxWidth: number): this {
        if (!this.internal.fieldConfig) {
            this.internal.fieldConfig = dashboard.defaultFieldConfigSource();
        }
        if (!this.internal.fieldConfig.defaults) {
            this.internal.fieldConfig.defaults = dashboard.defaultFieldConfig();
        }
        if (!this.internal.fieldConfig.defaults.custom) {
            this.internal.fieldConfig.defaults.custom = timeseries.defaultFieldConfig();
        }
        this.internal.fieldConfig.defaults.custom.barMaxWidth = barMaxWidth;
        return this;
    }
}
