// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';
import * as heatmap from '../heatmap';
import * as common from '../common';

export class VisualizationBuilder implements cog.Builder<dashboardv2beta1.VizConfigKind> {
    protected readonly internal: dashboardv2beta1.VizConfigKind;

    constructor() {
        this.internal = dashboardv2beta1.defaultVizConfigKind();
        this.internal.kind = "VizConfig";
        this.internal.group = "heatmap";
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

    // Calculate min max per field
    fieldMinMax(fieldMinMax: boolean): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultVizConfigSpec();
        }
        if (!this.internal.spec.fieldConfig) {
            this.internal.spec.fieldConfig = dashboardv2beta1.defaultFieldConfigSource();
        }
        if (!this.internal.spec.fieldConfig.defaults) {
            this.internal.spec.fieldConfig.defaults = dashboardv2beta1.defaultFieldConfig();
        }
        this.internal.spec.fieldConfig.defaults.fieldMinMax = fieldMinMax;
        return this;
    }

    // How null values should be handled when calculating field stats
    // "null" - Include null values, "connected" - Ignore nulls, "null as zero" - Treat nulls as zero
    nullValueMode(nullValueMode: dashboardv2beta1.NullValueMode): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultVizConfigSpec();
        }
        if (!this.internal.spec.fieldConfig) {
            this.internal.spec.fieldConfig = dashboardv2beta1.defaultFieldConfigSource();
        }
        if (!this.internal.spec.fieldConfig.defaults) {
            this.internal.spec.fieldConfig.defaults = dashboardv2beta1.defaultFieldConfig();
        }
        this.internal.spec.fieldConfig.defaults.nullValueMode = nullValueMode;
        return this;
    }

    // Overrides are the options applied to specific fields overriding the defaults.
    overrides(overrides: {
	// Describes config override rules created when interacting with Grafana.
	__systemRef?: string;
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
	// Describes config override rules created when interacting with Grafana.
	__systemRef?: string;
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

    // Controls if the heatmap should be calculated from data
    calculate(calculate: boolean): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultVizConfigSpec();
        }
        if (!this.internal.spec.options) {
            this.internal.spec.options = heatmap.defaultOptions();
        }
        this.internal.spec.options.calculate = calculate;
        return this;
    }

    // Calculation options for the heatmap
    calculation(calculation: cog.Builder<common.HeatmapCalculationOptions>): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultVizConfigSpec();
        }
        if (!this.internal.spec.options) {
            this.internal.spec.options = heatmap.defaultOptions();
        }
        const calculationResource = calculation.build();
        this.internal.spec.options.calculation = calculationResource;
        return this;
    }

    // Controls the color options
    color(color: cog.Builder<heatmap.HeatmapColorOptions>): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultVizConfigSpec();
        }
        if (!this.internal.spec.options) {
            this.internal.spec.options = heatmap.defaultOptions();
        }
        const colorResource = color.build();
        this.internal.spec.options.color = colorResource;
        return this;
    }

    // Filters values between a given range
    filterValues(filterValues: cog.Builder<heatmap.FilterValueRange>): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultVizConfigSpec();
        }
        if (!this.internal.spec.options) {
            this.internal.spec.options = heatmap.defaultOptions();
        }
        const filterValuesResource = filterValues.build();
        this.internal.spec.options.filterValues = filterValuesResource;
        return this;
    }

    // Controls tick alignment and value name when not calculating from data
    rowsFrame(rowsFrame: cog.Builder<heatmap.RowsHeatmapOptions>): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultVizConfigSpec();
        }
        if (!this.internal.spec.options) {
            this.internal.spec.options = heatmap.defaultOptions();
        }
        const rowsFrameResource = rowsFrame.build();
        this.internal.spec.options.rowsFrame = rowsFrameResource;
        return this;
    }

    // | *{
    // 	layout: ui.HeatmapCellLayout & "auto" // TODO: fix after remove when https://github.com/grafana/cuetsy/issues/74 is fixed
    // }
    // Controls the display of the value in the cell
    showValue(showValue: common.VisibilityMode): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultVizConfigSpec();
        }
        if (!this.internal.spec.options) {
            this.internal.spec.options = heatmap.defaultOptions();
        }
        this.internal.spec.options.showValue = showValue;
        return this;
    }

    // Controls gap between cells
    cellGap(cellGap: number): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultVizConfigSpec();
        }
        if (!this.internal.spec.options) {
            this.internal.spec.options = heatmap.defaultOptions();
        }
        if (!(cellGap <= 25)) {
            throw new Error("cellGap must be <= 25");
        }
        this.internal.spec.options.cellGap = cellGap;
        return this;
    }

    // Controls cell radius
    cellRadius(cellRadius: number): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultVizConfigSpec();
        }
        if (!this.internal.spec.options) {
            this.internal.spec.options = heatmap.defaultOptions();
        }
        this.internal.spec.options.cellRadius = cellRadius;
        return this;
    }

    // Controls cell value unit
    cellValues(cellValues: cog.Builder<heatmap.CellValues>): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultVizConfigSpec();
        }
        if (!this.internal.spec.options) {
            this.internal.spec.options = heatmap.defaultOptions();
        }
        const cellValuesResource = cellValues.build();
        this.internal.spec.options.cellValues = cellValuesResource;
        return this;
    }

    // Controls yAxis placement
    yAxis(yAxis: cog.Builder<heatmap.YAxisConfig>): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultVizConfigSpec();
        }
        if (!this.internal.spec.options) {
            this.internal.spec.options = heatmap.defaultOptions();
        }
        const yAxisResource = yAxis.build();
        this.internal.spec.options.yAxis = yAxisResource;
        return this;
    }

    // | *{
    // 	axisPlacement: ui.AxisPlacement & "left" // TODO: fix after remove when https://github.com/grafana/cuetsy/issues/74 is fixed
    // }
    // Controls legend options
    legend(legend: cog.Builder<heatmap.HeatmapLegend>): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultVizConfigSpec();
        }
        if (!this.internal.spec.options) {
            this.internal.spec.options = heatmap.defaultOptions();
        }
        const legendResource = legend.build();
        this.internal.spec.options.legend = legendResource;
        return this;
    }

    // Controls tooltip options
    tooltip(tooltip: cog.Builder<heatmap.HeatmapTooltip>): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultVizConfigSpec();
        }
        if (!this.internal.spec.options) {
            this.internal.spec.options = heatmap.defaultOptions();
        }
        const tooltipResource = tooltip.build();
        this.internal.spec.options.tooltip = tooltipResource;
        return this;
    }

    // Controls exemplar options
    exemplars(exemplars: cog.Builder<heatmap.ExemplarConfig>): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultVizConfigSpec();
        }
        if (!this.internal.spec.options) {
            this.internal.spec.options = heatmap.defaultOptions();
        }
        const exemplarsResource = exemplars.build();
        this.internal.spec.options.exemplars = exemplarsResource;
        return this;
    }

    // Controls which axis to allow selection on
    selectionMode(selectionMode: heatmap.HeatmapSelectionMode): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultVizConfigSpec();
        }
        if (!this.internal.spec.options) {
            this.internal.spec.options = heatmap.defaultOptions();
        }
        this.internal.spec.options.selectionMode = selectionMode;
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
            this.internal.spec.fieldConfig.defaults.custom = heatmap.defaultFieldConfig();
        }
        const scaleDistributionResource = scaleDistribution.build();
        this.internal.spec.fieldConfig.defaults.custom.scaleDistribution = scaleDistributionResource;
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
            this.internal.spec.fieldConfig.defaults.custom = heatmap.defaultFieldConfig();
        }
        const hideFromResource = hideFrom.build();
        this.internal.spec.fieldConfig.defaults.custom.hideFrom = hideFromResource;
        return this;
    }
}

