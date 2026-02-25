// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';
import * as candlestick from '../candlestick';
import * as common from '../common';

export class VisualizationBuilder implements cog.Builder<dashboardv2beta1.VizConfigKind> {
    protected readonly internal: dashboardv2beta1.VizConfigKind;

    constructor() {
        this.internal = dashboardv2beta1.defaultVizConfigKind();
        this.internal.kind = "VizConfig";
        this.internal.group = "candlestick";
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.VizConfigKind {
        return this.internal;
    }

    // The display value for this field.  This supports template variables blank is auto
    displayName(displayName: string): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultVizConfigSpec();
        }
        if (!this.internal.spec.fieldConfig) {
            this.internal.spec.fieldConfig = dashboardv2beta1.defaultFieldConfigSource();
        }
        if (!this.internal.spec.fieldConfig.defaults) {
            this.internal.spec.fieldConfig.defaults = dashboardv2beta1.defaultFieldConfig();
        }
        this.internal.spec.fieldConfig.defaults.displayName = displayName;
        return this;
    }

    // This can be used by data sources that return and explicit naming structure for values and labels
    // When this property is configured, this value is used rather than the default naming strategy.
    displayNameFromDS(displayNameFromDS: string): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultVizConfigSpec();
        }
        if (!this.internal.spec.fieldConfig) {
            this.internal.spec.fieldConfig = dashboardv2beta1.defaultFieldConfigSource();
        }
        if (!this.internal.spec.fieldConfig.defaults) {
            this.internal.spec.fieldConfig.defaults = dashboardv2beta1.defaultFieldConfig();
        }
        this.internal.spec.fieldConfig.defaults.displayNameFromDS = displayNameFromDS;
        return this;
    }

    // Human readable field metadata
    description(description: string): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultVizConfigSpec();
        }
        if (!this.internal.spec.fieldConfig) {
            this.internal.spec.fieldConfig = dashboardv2beta1.defaultFieldConfigSource();
        }
        if (!this.internal.spec.fieldConfig.defaults) {
            this.internal.spec.fieldConfig.defaults = dashboardv2beta1.defaultFieldConfig();
        }
        this.internal.spec.fieldConfig.defaults.description = description;
        return this;
    }

    // An explicit path to the field in the datasource.  When the frame meta includes a path,
    // This will default to `${frame.meta.path}/${field.name}
    // 
    // When defined, this value can be used as an identifier within the datasource scope, and
    // may be used to update the results
    path(path: string): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultVizConfigSpec();
        }
        if (!this.internal.spec.fieldConfig) {
            this.internal.spec.fieldConfig = dashboardv2beta1.defaultFieldConfigSource();
        }
        if (!this.internal.spec.fieldConfig.defaults) {
            this.internal.spec.fieldConfig.defaults = dashboardv2beta1.defaultFieldConfig();
        }
        this.internal.spec.fieldConfig.defaults.path = path;
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
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultVizConfigSpec();
        }
        if (!this.internal.spec.fieldConfig) {
            this.internal.spec.fieldConfig = dashboardv2beta1.defaultFieldConfigSource();
        }
        if (!this.internal.spec.fieldConfig.defaults) {
            this.internal.spec.fieldConfig.defaults = dashboardv2beta1.defaultFieldConfig();
        }
        this.internal.spec.fieldConfig.defaults.unit = unit;
        return this;
    }

    // Specify the number of decimals Grafana includes in the rendered value.
    // If you leave this field blank, Grafana automatically truncates the number of decimals based on the value.
    // For example 1.1234 will display as 1.12 and 100.456 will display as 100.
    // To display all decimals, set the unit to `String`.
    decimals(decimals: number): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultVizConfigSpec();
        }
        if (!this.internal.spec.fieldConfig) {
            this.internal.spec.fieldConfig = dashboardv2beta1.defaultFieldConfigSource();
        }
        if (!this.internal.spec.fieldConfig.defaults) {
            this.internal.spec.fieldConfig.defaults = dashboardv2beta1.defaultFieldConfig();
        }
        this.internal.spec.fieldConfig.defaults.decimals = decimals;
        return this;
    }

    // The minimum value used in percentage threshold calculations. Leave blank for auto calculation based on all series and fields.
    min(min: number): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultVizConfigSpec();
        }
        if (!this.internal.spec.fieldConfig) {
            this.internal.spec.fieldConfig = dashboardv2beta1.defaultFieldConfigSource();
        }
        if (!this.internal.spec.fieldConfig.defaults) {
            this.internal.spec.fieldConfig.defaults = dashboardv2beta1.defaultFieldConfig();
        }
        this.internal.spec.fieldConfig.defaults.min = min;
        return this;
    }

    // The maximum value used in percentage threshold calculations. Leave blank for auto calculation based on all series and fields.
    max(max: number): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultVizConfigSpec();
        }
        if (!this.internal.spec.fieldConfig) {
            this.internal.spec.fieldConfig = dashboardv2beta1.defaultFieldConfigSource();
        }
        if (!this.internal.spec.fieldConfig.defaults) {
            this.internal.spec.fieldConfig.defaults = dashboardv2beta1.defaultFieldConfig();
        }
        this.internal.spec.fieldConfig.defaults.max = max;
        return this;
    }

    // Convert input values into a display string
    mappings(mappings: dashboardv2beta1.ValueMapping[]): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultVizConfigSpec();
        }
        if (!this.internal.spec.fieldConfig) {
            this.internal.spec.fieldConfig = dashboardv2beta1.defaultFieldConfigSource();
        }
        if (!this.internal.spec.fieldConfig.defaults) {
            this.internal.spec.fieldConfig.defaults = dashboardv2beta1.defaultFieldConfig();
        }
        this.internal.spec.fieldConfig.defaults.mappings = mappings;
        return this;
    }

    // Map numeric values to states
    thresholds(thresholds: cog.Builder<dashboardv2beta1.ThresholdsConfig>): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultVizConfigSpec();
        }
        if (!this.internal.spec.fieldConfig) {
            this.internal.spec.fieldConfig = dashboardv2beta1.defaultFieldConfigSource();
        }
        if (!this.internal.spec.fieldConfig.defaults) {
            this.internal.spec.fieldConfig.defaults = dashboardv2beta1.defaultFieldConfig();
        }
        const thresholdsResource = thresholds.build();
        this.internal.spec.fieldConfig.defaults.thresholds = thresholdsResource;
        return this;
    }

    // Panel color configuration
    colorScheme(color: cog.Builder<dashboardv2beta1.FieldColor>): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultVizConfigSpec();
        }
        if (!this.internal.spec.fieldConfig) {
            this.internal.spec.fieldConfig = dashboardv2beta1.defaultFieldConfigSource();
        }
        if (!this.internal.spec.fieldConfig.defaults) {
            this.internal.spec.fieldConfig.defaults = dashboardv2beta1.defaultFieldConfig();
        }
        const colorResource = color.build();
        this.internal.spec.fieldConfig.defaults.color = colorResource;
        return this;
    }

    // The behavior when clicking on a result
    dataLinks(links: any[]): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultVizConfigSpec();
        }
        if (!this.internal.spec.fieldConfig) {
            this.internal.spec.fieldConfig = dashboardv2beta1.defaultFieldConfigSource();
        }
        if (!this.internal.spec.fieldConfig.defaults) {
            this.internal.spec.fieldConfig.defaults = dashboardv2beta1.defaultFieldConfig();
        }
        this.internal.spec.fieldConfig.defaults.links = links;
        return this;
    }

    // Define interactive HTTP requests that can be triggered from data visualizations.
    actions(actions: cog.Builder<dashboardv2beta1.Action>[]): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultVizConfigSpec();
        }
        if (!this.internal.spec.fieldConfig) {
            this.internal.spec.fieldConfig = dashboardv2beta1.defaultFieldConfigSource();
        }
        if (!this.internal.spec.fieldConfig.defaults) {
            this.internal.spec.fieldConfig.defaults = dashboardv2beta1.defaultFieldConfig();
        }
        const actionsResources = actions.map(builder1 => builder1.build());
        this.internal.spec.fieldConfig.defaults.actions = actionsResources;
        return this;
    }

    // Alternative to empty string
    noValue(noValue: string): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultVizConfigSpec();
        }
        if (!this.internal.spec.fieldConfig) {
            this.internal.spec.fieldConfig = dashboardv2beta1.defaultFieldConfigSource();
        }
        if (!this.internal.spec.fieldConfig.defaults) {
            this.internal.spec.fieldConfig.defaults = dashboardv2beta1.defaultFieldConfig();
        }
        this.internal.spec.fieldConfig.defaults.noValue = noValue;
        return this;
    }

    // Overrides are the options applied to specific fields overriding the defaults.
    overrides(overrides: {
	matcher: dashboardv2beta1.MatcherConfig;
	properties: dashboardv2beta1.DynamicConfigValue[];
}[]): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultVizConfigSpec();
        }
        if (!this.internal.spec.fieldConfig) {
            this.internal.spec.fieldConfig = dashboardv2beta1.defaultFieldConfigSource();
        }
        this.internal.spec.fieldConfig.overrides = overrides;
        return this;
    }

    // Overrides are the options applied to specific fields overriding the defaults.
    override(override: {
	matcher: dashboardv2beta1.MatcherConfig;
	properties: dashboardv2beta1.DynamicConfigValue[];
}): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultVizConfigSpec();
        }
        if (!this.internal.spec.fieldConfig) {
            this.internal.spec.fieldConfig = dashboardv2beta1.defaultFieldConfigSource();
        }
        if (!this.internal.spec.fieldConfig.overrides) {
            this.internal.spec.fieldConfig.overrides = [];
        }
        this.internal.spec.fieldConfig.overrides.push(override);
        return this;
    }

    // Sets which dimensions are used for the visualization
    mode(mode: candlestick.VizDisplayMode): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultVizConfigSpec();
        }
        if (!this.internal.spec.options) {
            this.internal.spec.options = candlestick.defaultOptions();
        }
        this.internal.spec.options.mode = mode;
        return this;
    }

    // Sets the style of the candlesticks
    candleStyle(candleStyle: candlestick.CandleStyle): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultVizConfigSpec();
        }
        if (!this.internal.spec.options) {
            this.internal.spec.options = candlestick.defaultOptions();
        }
        this.internal.spec.options.candleStyle = candleStyle;
        return this;
    }

    // Sets the color strategy for the candlesticks
    colorStrategy(colorStrategy: candlestick.ColorStrategy): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultVizConfigSpec();
        }
        if (!this.internal.spec.options) {
            this.internal.spec.options = candlestick.defaultOptions();
        }
        this.internal.spec.options.colorStrategy = colorStrategy;
        return this;
    }

    // Map fields to appropriate dimension
    fields(fields: cog.Builder<candlestick.CandlestickFieldMap>): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultVizConfigSpec();
        }
        if (!this.internal.spec.options) {
            this.internal.spec.options = candlestick.defaultOptions();
        }
        const fieldsResource = fields.build();
        this.internal.spec.options.fields = fieldsResource;
        return this;
    }

    // Set which colors are used when the price movement is up or down
    colors(colors: cog.Builder<candlestick.CandlestickColors>): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultVizConfigSpec();
        }
        if (!this.internal.spec.options) {
            this.internal.spec.options = candlestick.defaultOptions();
        }
        const colorsResource = colors.build();
        this.internal.spec.options.colors = colorsResource;
        return this;
    }

    legend(legend: cog.Builder<common.VizLegendOptions>): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultVizConfigSpec();
        }
        if (!this.internal.spec.options) {
            this.internal.spec.options = candlestick.defaultOptions();
        }
        const legendResource = legend.build();
        this.internal.spec.options.legend = legendResource;
        return this;
    }

    tooltip(tooltip: cog.Builder<common.VizTooltipOptions>): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultVizConfigSpec();
        }
        if (!this.internal.spec.options) {
            this.internal.spec.options = candlestick.defaultOptions();
        }
        const tooltipResource = tooltip.build();
        this.internal.spec.options.tooltip = tooltipResource;
        return this;
    }

    // When enabled, all fields will be sent to the graph
    includeAllFields(includeAllFields: boolean): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultVizConfigSpec();
        }
        if (!this.internal.spec.options) {
            this.internal.spec.options = candlestick.defaultOptions();
        }
        this.internal.spec.options.includeAllFields = includeAllFields;
        return this;
    }

    drawStyle(drawStyle: common.GraphDrawStyle): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultVizConfigSpec();
        }
        if (!this.internal.spec.fieldConfig) {
            this.internal.spec.fieldConfig = dashboardv2beta1.defaultFieldConfigSource();
        }
        if (!this.internal.spec.fieldConfig.defaults) {
            this.internal.spec.fieldConfig.defaults = dashboardv2beta1.defaultFieldConfig();
        }
        if (!this.internal.spec.fieldConfig.defaults.custom) {
            this.internal.spec.fieldConfig.defaults.custom = candlestick.defaultFieldConfig();
        }
        this.internal.spec.fieldConfig.defaults.custom.drawStyle = drawStyle;
        return this;
    }

    gradientMode(gradientMode: common.GraphGradientMode): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultVizConfigSpec();
        }
        if (!this.internal.spec.fieldConfig) {
            this.internal.spec.fieldConfig = dashboardv2beta1.defaultFieldConfigSource();
        }
        if (!this.internal.spec.fieldConfig.defaults) {
            this.internal.spec.fieldConfig.defaults = dashboardv2beta1.defaultFieldConfig();
        }
        if (!this.internal.spec.fieldConfig.defaults.custom) {
            this.internal.spec.fieldConfig.defaults.custom = candlestick.defaultFieldConfig();
        }
        this.internal.spec.fieldConfig.defaults.custom.gradientMode = gradientMode;
        return this;
    }

    thresholdsStyle(thresholdsStyle: cog.Builder<common.GraphThresholdsStyleConfig>): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultVizConfigSpec();
        }
        if (!this.internal.spec.fieldConfig) {
            this.internal.spec.fieldConfig = dashboardv2beta1.defaultFieldConfigSource();
        }
        if (!this.internal.spec.fieldConfig.defaults) {
            this.internal.spec.fieldConfig.defaults = dashboardv2beta1.defaultFieldConfig();
        }
        if (!this.internal.spec.fieldConfig.defaults.custom) {
            this.internal.spec.fieldConfig.defaults.custom = candlestick.defaultFieldConfig();
        }
        const thresholdsStyleResource = thresholdsStyle.build();
        this.internal.spec.fieldConfig.defaults.custom.thresholdsStyle = thresholdsStyleResource;
        return this;
    }

    transform(transform: common.GraphTransform): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultVizConfigSpec();
        }
        if (!this.internal.spec.fieldConfig) {
            this.internal.spec.fieldConfig = dashboardv2beta1.defaultFieldConfigSource();
        }
        if (!this.internal.spec.fieldConfig.defaults) {
            this.internal.spec.fieldConfig.defaults = dashboardv2beta1.defaultFieldConfig();
        }
        if (!this.internal.spec.fieldConfig.defaults.custom) {
            this.internal.spec.fieldConfig.defaults.custom = candlestick.defaultFieldConfig();
        }
        this.internal.spec.fieldConfig.defaults.custom.transform = transform;
        return this;
    }

    lineColor(lineColor: string): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultVizConfigSpec();
        }
        if (!this.internal.spec.fieldConfig) {
            this.internal.spec.fieldConfig = dashboardv2beta1.defaultFieldConfigSource();
        }
        if (!this.internal.spec.fieldConfig.defaults) {
            this.internal.spec.fieldConfig.defaults = dashboardv2beta1.defaultFieldConfig();
        }
        if (!this.internal.spec.fieldConfig.defaults.custom) {
            this.internal.spec.fieldConfig.defaults.custom = candlestick.defaultFieldConfig();
        }
        this.internal.spec.fieldConfig.defaults.custom.lineColor = lineColor;
        return this;
    }

    lineWidth(lineWidth: number): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultVizConfigSpec();
        }
        if (!this.internal.spec.fieldConfig) {
            this.internal.spec.fieldConfig = dashboardv2beta1.defaultFieldConfigSource();
        }
        if (!this.internal.spec.fieldConfig.defaults) {
            this.internal.spec.fieldConfig.defaults = dashboardv2beta1.defaultFieldConfig();
        }
        if (!this.internal.spec.fieldConfig.defaults.custom) {
            this.internal.spec.fieldConfig.defaults.custom = candlestick.defaultFieldConfig();
        }
        this.internal.spec.fieldConfig.defaults.custom.lineWidth = lineWidth;
        return this;
    }

    lineInterpolation(lineInterpolation: common.LineInterpolation): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultVizConfigSpec();
        }
        if (!this.internal.spec.fieldConfig) {
            this.internal.spec.fieldConfig = dashboardv2beta1.defaultFieldConfigSource();
        }
        if (!this.internal.spec.fieldConfig.defaults) {
            this.internal.spec.fieldConfig.defaults = dashboardv2beta1.defaultFieldConfig();
        }
        if (!this.internal.spec.fieldConfig.defaults.custom) {
            this.internal.spec.fieldConfig.defaults.custom = candlestick.defaultFieldConfig();
        }
        this.internal.spec.fieldConfig.defaults.custom.lineInterpolation = lineInterpolation;
        return this;
    }

    lineStyle(lineStyle: cog.Builder<common.LineStyle>): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultVizConfigSpec();
        }
        if (!this.internal.spec.fieldConfig) {
            this.internal.spec.fieldConfig = dashboardv2beta1.defaultFieldConfigSource();
        }
        if (!this.internal.spec.fieldConfig.defaults) {
            this.internal.spec.fieldConfig.defaults = dashboardv2beta1.defaultFieldConfig();
        }
        if (!this.internal.spec.fieldConfig.defaults.custom) {
            this.internal.spec.fieldConfig.defaults.custom = candlestick.defaultFieldConfig();
        }
        const lineStyleResource = lineStyle.build();
        this.internal.spec.fieldConfig.defaults.custom.lineStyle = lineStyleResource;
        return this;
    }

    fillColor(fillColor: string): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultVizConfigSpec();
        }
        if (!this.internal.spec.fieldConfig) {
            this.internal.spec.fieldConfig = dashboardv2beta1.defaultFieldConfigSource();
        }
        if (!this.internal.spec.fieldConfig.defaults) {
            this.internal.spec.fieldConfig.defaults = dashboardv2beta1.defaultFieldConfig();
        }
        if (!this.internal.spec.fieldConfig.defaults.custom) {
            this.internal.spec.fieldConfig.defaults.custom = candlestick.defaultFieldConfig();
        }
        this.internal.spec.fieldConfig.defaults.custom.fillColor = fillColor;
        return this;
    }

    fillOpacity(fillOpacity: number): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultVizConfigSpec();
        }
        if (!this.internal.spec.fieldConfig) {
            this.internal.spec.fieldConfig = dashboardv2beta1.defaultFieldConfigSource();
        }
        if (!this.internal.spec.fieldConfig.defaults) {
            this.internal.spec.fieldConfig.defaults = dashboardv2beta1.defaultFieldConfig();
        }
        if (!this.internal.spec.fieldConfig.defaults.custom) {
            this.internal.spec.fieldConfig.defaults.custom = candlestick.defaultFieldConfig();
        }
        this.internal.spec.fieldConfig.defaults.custom.fillOpacity = fillOpacity;
        return this;
    }

    showPoints(showPoints: common.VisibilityMode): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultVizConfigSpec();
        }
        if (!this.internal.spec.fieldConfig) {
            this.internal.spec.fieldConfig = dashboardv2beta1.defaultFieldConfigSource();
        }
        if (!this.internal.spec.fieldConfig.defaults) {
            this.internal.spec.fieldConfig.defaults = dashboardv2beta1.defaultFieldConfig();
        }
        if (!this.internal.spec.fieldConfig.defaults.custom) {
            this.internal.spec.fieldConfig.defaults.custom = candlestick.defaultFieldConfig();
        }
        this.internal.spec.fieldConfig.defaults.custom.showPoints = showPoints;
        return this;
    }

    pointSize(pointSize: number): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultVizConfigSpec();
        }
        if (!this.internal.spec.fieldConfig) {
            this.internal.spec.fieldConfig = dashboardv2beta1.defaultFieldConfigSource();
        }
        if (!this.internal.spec.fieldConfig.defaults) {
            this.internal.spec.fieldConfig.defaults = dashboardv2beta1.defaultFieldConfig();
        }
        if (!this.internal.spec.fieldConfig.defaults.custom) {
            this.internal.spec.fieldConfig.defaults.custom = candlestick.defaultFieldConfig();
        }
        this.internal.spec.fieldConfig.defaults.custom.pointSize = pointSize;
        return this;
    }

    pointColor(pointColor: string): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultVizConfigSpec();
        }
        if (!this.internal.spec.fieldConfig) {
            this.internal.spec.fieldConfig = dashboardv2beta1.defaultFieldConfigSource();
        }
        if (!this.internal.spec.fieldConfig.defaults) {
            this.internal.spec.fieldConfig.defaults = dashboardv2beta1.defaultFieldConfig();
        }
        if (!this.internal.spec.fieldConfig.defaults.custom) {
            this.internal.spec.fieldConfig.defaults.custom = candlestick.defaultFieldConfig();
        }
        this.internal.spec.fieldConfig.defaults.custom.pointColor = pointColor;
        return this;
    }

    axisPlacement(axisPlacement: common.AxisPlacement): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultVizConfigSpec();
        }
        if (!this.internal.spec.fieldConfig) {
            this.internal.spec.fieldConfig = dashboardv2beta1.defaultFieldConfigSource();
        }
        if (!this.internal.spec.fieldConfig.defaults) {
            this.internal.spec.fieldConfig.defaults = dashboardv2beta1.defaultFieldConfig();
        }
        if (!this.internal.spec.fieldConfig.defaults.custom) {
            this.internal.spec.fieldConfig.defaults.custom = candlestick.defaultFieldConfig();
        }
        this.internal.spec.fieldConfig.defaults.custom.axisPlacement = axisPlacement;
        return this;
    }

    axisColorMode(axisColorMode: common.AxisColorMode): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultVizConfigSpec();
        }
        if (!this.internal.spec.fieldConfig) {
            this.internal.spec.fieldConfig = dashboardv2beta1.defaultFieldConfigSource();
        }
        if (!this.internal.spec.fieldConfig.defaults) {
            this.internal.spec.fieldConfig.defaults = dashboardv2beta1.defaultFieldConfig();
        }
        if (!this.internal.spec.fieldConfig.defaults.custom) {
            this.internal.spec.fieldConfig.defaults.custom = candlestick.defaultFieldConfig();
        }
        this.internal.spec.fieldConfig.defaults.custom.axisColorMode = axisColorMode;
        return this;
    }

    axisLabel(axisLabel: string): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultVizConfigSpec();
        }
        if (!this.internal.spec.fieldConfig) {
            this.internal.spec.fieldConfig = dashboardv2beta1.defaultFieldConfigSource();
        }
        if (!this.internal.spec.fieldConfig.defaults) {
            this.internal.spec.fieldConfig.defaults = dashboardv2beta1.defaultFieldConfig();
        }
        if (!this.internal.spec.fieldConfig.defaults.custom) {
            this.internal.spec.fieldConfig.defaults.custom = candlestick.defaultFieldConfig();
        }
        this.internal.spec.fieldConfig.defaults.custom.axisLabel = axisLabel;
        return this;
    }

    axisWidth(axisWidth: number): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultVizConfigSpec();
        }
        if (!this.internal.spec.fieldConfig) {
            this.internal.spec.fieldConfig = dashboardv2beta1.defaultFieldConfigSource();
        }
        if (!this.internal.spec.fieldConfig.defaults) {
            this.internal.spec.fieldConfig.defaults = dashboardv2beta1.defaultFieldConfig();
        }
        if (!this.internal.spec.fieldConfig.defaults.custom) {
            this.internal.spec.fieldConfig.defaults.custom = candlestick.defaultFieldConfig();
        }
        this.internal.spec.fieldConfig.defaults.custom.axisWidth = axisWidth;
        return this;
    }

    axisSoftMin(axisSoftMin: number): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultVizConfigSpec();
        }
        if (!this.internal.spec.fieldConfig) {
            this.internal.spec.fieldConfig = dashboardv2beta1.defaultFieldConfigSource();
        }
        if (!this.internal.spec.fieldConfig.defaults) {
            this.internal.spec.fieldConfig.defaults = dashboardv2beta1.defaultFieldConfig();
        }
        if (!this.internal.spec.fieldConfig.defaults.custom) {
            this.internal.spec.fieldConfig.defaults.custom = candlestick.defaultFieldConfig();
        }
        this.internal.spec.fieldConfig.defaults.custom.axisSoftMin = axisSoftMin;
        return this;
    }

    axisSoftMax(axisSoftMax: number): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultVizConfigSpec();
        }
        if (!this.internal.spec.fieldConfig) {
            this.internal.spec.fieldConfig = dashboardv2beta1.defaultFieldConfigSource();
        }
        if (!this.internal.spec.fieldConfig.defaults) {
            this.internal.spec.fieldConfig.defaults = dashboardv2beta1.defaultFieldConfig();
        }
        if (!this.internal.spec.fieldConfig.defaults.custom) {
            this.internal.spec.fieldConfig.defaults.custom = candlestick.defaultFieldConfig();
        }
        this.internal.spec.fieldConfig.defaults.custom.axisSoftMax = axisSoftMax;
        return this;
    }

    axisGridShow(axisGridShow: boolean): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultVizConfigSpec();
        }
        if (!this.internal.spec.fieldConfig) {
            this.internal.spec.fieldConfig = dashboardv2beta1.defaultFieldConfigSource();
        }
        if (!this.internal.spec.fieldConfig.defaults) {
            this.internal.spec.fieldConfig.defaults = dashboardv2beta1.defaultFieldConfig();
        }
        if (!this.internal.spec.fieldConfig.defaults.custom) {
            this.internal.spec.fieldConfig.defaults.custom = candlestick.defaultFieldConfig();
        }
        this.internal.spec.fieldConfig.defaults.custom.axisGridShow = axisGridShow;
        return this;
    }

    scaleDistribution(scaleDistribution: cog.Builder<common.ScaleDistributionConfig>): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultVizConfigSpec();
        }
        if (!this.internal.spec.fieldConfig) {
            this.internal.spec.fieldConfig = dashboardv2beta1.defaultFieldConfigSource();
        }
        if (!this.internal.spec.fieldConfig.defaults) {
            this.internal.spec.fieldConfig.defaults = dashboardv2beta1.defaultFieldConfig();
        }
        if (!this.internal.spec.fieldConfig.defaults.custom) {
            this.internal.spec.fieldConfig.defaults.custom = candlestick.defaultFieldConfig();
        }
        const scaleDistributionResource = scaleDistribution.build();
        this.internal.spec.fieldConfig.defaults.custom.scaleDistribution = scaleDistributionResource;
        return this;
    }

    axisCenteredZero(axisCenteredZero: boolean): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultVizConfigSpec();
        }
        if (!this.internal.spec.fieldConfig) {
            this.internal.spec.fieldConfig = dashboardv2beta1.defaultFieldConfigSource();
        }
        if (!this.internal.spec.fieldConfig.defaults) {
            this.internal.spec.fieldConfig.defaults = dashboardv2beta1.defaultFieldConfig();
        }
        if (!this.internal.spec.fieldConfig.defaults.custom) {
            this.internal.spec.fieldConfig.defaults.custom = candlestick.defaultFieldConfig();
        }
        this.internal.spec.fieldConfig.defaults.custom.axisCenteredZero = axisCenteredZero;
        return this;
    }

    barAlignment(barAlignment: common.BarAlignment): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultVizConfigSpec();
        }
        if (!this.internal.spec.fieldConfig) {
            this.internal.spec.fieldConfig = dashboardv2beta1.defaultFieldConfigSource();
        }
        if (!this.internal.spec.fieldConfig.defaults) {
            this.internal.spec.fieldConfig.defaults = dashboardv2beta1.defaultFieldConfig();
        }
        if (!this.internal.spec.fieldConfig.defaults.custom) {
            this.internal.spec.fieldConfig.defaults.custom = candlestick.defaultFieldConfig();
        }
        this.internal.spec.fieldConfig.defaults.custom.barAlignment = barAlignment;
        return this;
    }

    barWidthFactor(barWidthFactor: number): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultVizConfigSpec();
        }
        if (!this.internal.spec.fieldConfig) {
            this.internal.spec.fieldConfig = dashboardv2beta1.defaultFieldConfigSource();
        }
        if (!this.internal.spec.fieldConfig.defaults) {
            this.internal.spec.fieldConfig.defaults = dashboardv2beta1.defaultFieldConfig();
        }
        if (!this.internal.spec.fieldConfig.defaults.custom) {
            this.internal.spec.fieldConfig.defaults.custom = candlestick.defaultFieldConfig();
        }
        this.internal.spec.fieldConfig.defaults.custom.barWidthFactor = barWidthFactor;
        return this;
    }

    stacking(stacking: cog.Builder<common.StackingConfig>): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultVizConfigSpec();
        }
        if (!this.internal.spec.fieldConfig) {
            this.internal.spec.fieldConfig = dashboardv2beta1.defaultFieldConfigSource();
        }
        if (!this.internal.spec.fieldConfig.defaults) {
            this.internal.spec.fieldConfig.defaults = dashboardv2beta1.defaultFieldConfig();
        }
        if (!this.internal.spec.fieldConfig.defaults.custom) {
            this.internal.spec.fieldConfig.defaults.custom = candlestick.defaultFieldConfig();
        }
        const stackingResource = stacking.build();
        this.internal.spec.fieldConfig.defaults.custom.stacking = stackingResource;
        return this;
    }

    hideFrom(hideFrom: cog.Builder<common.HideSeriesConfig>): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultVizConfigSpec();
        }
        if (!this.internal.spec.fieldConfig) {
            this.internal.spec.fieldConfig = dashboardv2beta1.defaultFieldConfigSource();
        }
        if (!this.internal.spec.fieldConfig.defaults) {
            this.internal.spec.fieldConfig.defaults = dashboardv2beta1.defaultFieldConfig();
        }
        if (!this.internal.spec.fieldConfig.defaults.custom) {
            this.internal.spec.fieldConfig.defaults.custom = candlestick.defaultFieldConfig();
        }
        const hideFromResource = hideFrom.build();
        this.internal.spec.fieldConfig.defaults.custom.hideFrom = hideFromResource;
        return this;
    }

    insertNulls(insertNulls: boolean | number): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultVizConfigSpec();
        }
        if (!this.internal.spec.fieldConfig) {
            this.internal.spec.fieldConfig = dashboardv2beta1.defaultFieldConfigSource();
        }
        if (!this.internal.spec.fieldConfig.defaults) {
            this.internal.spec.fieldConfig.defaults = dashboardv2beta1.defaultFieldConfig();
        }
        if (!this.internal.spec.fieldConfig.defaults.custom) {
            this.internal.spec.fieldConfig.defaults.custom = candlestick.defaultFieldConfig();
        }
        this.internal.spec.fieldConfig.defaults.custom.insertNulls = insertNulls;
        return this;
    }

    // Indicate if null values should be treated as gaps or connected.
    // When the value is a number, it represents the maximum delta in the
    // X axis that should be considered connected.  For timeseries, this is milliseconds
    spanNulls(spanNulls: boolean | number): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultVizConfigSpec();
        }
        if (!this.internal.spec.fieldConfig) {
            this.internal.spec.fieldConfig = dashboardv2beta1.defaultFieldConfigSource();
        }
        if (!this.internal.spec.fieldConfig.defaults) {
            this.internal.spec.fieldConfig.defaults = dashboardv2beta1.defaultFieldConfig();
        }
        if (!this.internal.spec.fieldConfig.defaults.custom) {
            this.internal.spec.fieldConfig.defaults.custom = candlestick.defaultFieldConfig();
        }
        this.internal.spec.fieldConfig.defaults.custom.spanNulls = spanNulls;
        return this;
    }

    fillBelowTo(fillBelowTo: string): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultVizConfigSpec();
        }
        if (!this.internal.spec.fieldConfig) {
            this.internal.spec.fieldConfig = dashboardv2beta1.defaultFieldConfigSource();
        }
        if (!this.internal.spec.fieldConfig.defaults) {
            this.internal.spec.fieldConfig.defaults = dashboardv2beta1.defaultFieldConfig();
        }
        if (!this.internal.spec.fieldConfig.defaults.custom) {
            this.internal.spec.fieldConfig.defaults.custom = candlestick.defaultFieldConfig();
        }
        this.internal.spec.fieldConfig.defaults.custom.fillBelowTo = fillBelowTo;
        return this;
    }

    pointSymbol(pointSymbol: string): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultVizConfigSpec();
        }
        if (!this.internal.spec.fieldConfig) {
            this.internal.spec.fieldConfig = dashboardv2beta1.defaultFieldConfigSource();
        }
        if (!this.internal.spec.fieldConfig.defaults) {
            this.internal.spec.fieldConfig.defaults = dashboardv2beta1.defaultFieldConfig();
        }
        if (!this.internal.spec.fieldConfig.defaults.custom) {
            this.internal.spec.fieldConfig.defaults.custom = candlestick.defaultFieldConfig();
        }
        this.internal.spec.fieldConfig.defaults.custom.pointSymbol = pointSymbol;
        return this;
    }

    axisBorderShow(axisBorderShow: boolean): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultVizConfigSpec();
        }
        if (!this.internal.spec.fieldConfig) {
            this.internal.spec.fieldConfig = dashboardv2beta1.defaultFieldConfigSource();
        }
        if (!this.internal.spec.fieldConfig.defaults) {
            this.internal.spec.fieldConfig.defaults = dashboardv2beta1.defaultFieldConfig();
        }
        if (!this.internal.spec.fieldConfig.defaults.custom) {
            this.internal.spec.fieldConfig.defaults.custom = candlestick.defaultFieldConfig();
        }
        this.internal.spec.fieldConfig.defaults.custom.axisBorderShow = axisBorderShow;
        return this;
    }

    barMaxWidth(barMaxWidth: number): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultVizConfigSpec();
        }
        if (!this.internal.spec.fieldConfig) {
            this.internal.spec.fieldConfig = dashboardv2beta1.defaultFieldConfigSource();
        }
        if (!this.internal.spec.fieldConfig.defaults) {
            this.internal.spec.fieldConfig.defaults = dashboardv2beta1.defaultFieldConfig();
        }
        if (!this.internal.spec.fieldConfig.defaults.custom) {
            this.internal.spec.fieldConfig.defaults.custom = candlestick.defaultFieldConfig();
        }
        this.internal.spec.fieldConfig.defaults.custom.barMaxWidth = barMaxWidth;
        return this;
    }

    // Adds override rules for a specific field, referred to by its name.
    overrideByName(name: string,properties: dashboardv2beta1.DynamicConfigValue[]): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultVizConfigSpec();
        }
        if (!this.internal.spec.fieldConfig) {
            this.internal.spec.fieldConfig = dashboardv2beta1.defaultFieldConfigSource();
        }
        if (!this.internal.spec.fieldConfig.overrides) {
            this.internal.spec.fieldConfig.overrides = [];
        }
        this.internal.spec.fieldConfig.overrides.push({
        matcher: {
        id: "byName",
        options: name,
    },
        properties: properties,
    });
        return this;
    }

    // Adds override rules for the fields whose name match the given regexp.
    overrideByRegexp(regexp: string,properties: dashboardv2beta1.DynamicConfigValue[]): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultVizConfigSpec();
        }
        if (!this.internal.spec.fieldConfig) {
            this.internal.spec.fieldConfig = dashboardv2beta1.defaultFieldConfigSource();
        }
        if (!this.internal.spec.fieldConfig.overrides) {
            this.internal.spec.fieldConfig.overrides = [];
        }
        this.internal.spec.fieldConfig.overrides.push({
        matcher: {
        id: "byRegexp",
        options: regexp,
    },
        properties: properties,
    });
        return this;
    }

    // Adds override rules for all the fields of the given type.
    overrideByFieldType(fieldType: string,properties: dashboardv2beta1.DynamicConfigValue[]): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultVizConfigSpec();
        }
        if (!this.internal.spec.fieldConfig) {
            this.internal.spec.fieldConfig = dashboardv2beta1.defaultFieldConfigSource();
        }
        if (!this.internal.spec.fieldConfig.overrides) {
            this.internal.spec.fieldConfig.overrides = [];
        }
        this.internal.spec.fieldConfig.overrides.push({
        matcher: {
        id: "byType",
        options: fieldType,
    },
        properties: properties,
    });
        return this;
    }

    overrideByQuery(queryRefId: string,properties: dashboardv2beta1.DynamicConfigValue[]): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultVizConfigSpec();
        }
        if (!this.internal.spec.fieldConfig) {
            this.internal.spec.fieldConfig = dashboardv2beta1.defaultFieldConfigSource();
        }
        if (!this.internal.spec.fieldConfig.overrides) {
            this.internal.spec.fieldConfig.overrides = [];
        }
        this.internal.spec.fieldConfig.overrides.push({
        matcher: {
        id: "byFrameRefID",
        options: queryRefId,
    },
        properties: properties,
    });
        return this;
    }
}

