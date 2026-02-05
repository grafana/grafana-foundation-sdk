"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.VisualizationBuilder = void 0;
const tslib_1 = require("tslib");
const dashboardv2beta1 = tslib_1.__importStar(require("../dashboardv2beta1"));
const xychart = tslib_1.__importStar(require("../xychart"));
class VisualizationBuilder {
    constructor() {
        this.internal = dashboardv2beta1.defaultVizConfigKind();
        this.internal.kind = "VizConfig";
        this.internal.group = "xychart";
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    // The display value for this field.  This supports template variables blank is auto
    displayName(displayName) {
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
    displayNameFromDS(displayNameFromDS) {
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
    description(description) {
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
    path(path) {
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
    unit(unit) {
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
    decimals(decimals) {
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
    min(min) {
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
    max(max) {
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
    mappings(mappings) {
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
    thresholds(thresholds) {
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
    colorScheme(color) {
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
    dataLinks(links) {
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
    actions(actions) {
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
    noValue(noValue) {
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
    overrides(overrides) {
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
    override(override) {
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
    overrideByName(name, properties) {
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
    overrideByRegexp(regexp, properties) {
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
    overrideByFieldType(fieldType, properties) {
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
    overrideByQuery(queryRefId, properties) {
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
    show(show) {
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
            this.internal.spec.fieldConfig.defaults.custom = xychart.defaultFieldConfig();
        }
        this.internal.spec.fieldConfig.defaults.custom.show = show;
        return this;
    }
    pointSize(pointSize) {
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
            this.internal.spec.fieldConfig.defaults.custom = xychart.defaultFieldConfig();
        }
        this.internal.spec.fieldConfig.defaults.custom.pointSize = pointSize;
        return this;
    }
    pointShape(pointShape) {
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
            this.internal.spec.fieldConfig.defaults.custom = xychart.defaultFieldConfig();
        }
        this.internal.spec.fieldConfig.defaults.custom.pointShape = pointShape;
        return this;
    }
    pointStrokeWidth(pointStrokeWidth) {
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
            this.internal.spec.fieldConfig.defaults.custom = xychart.defaultFieldConfig();
        }
        if (!(pointStrokeWidth >= 0)) {
            throw new Error("pointStrokeWidth must be >= 0");
        }
        this.internal.spec.fieldConfig.defaults.custom.pointStrokeWidth = pointStrokeWidth;
        return this;
    }
    fillOpacity(fillOpacity) {
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
            this.internal.spec.fieldConfig.defaults.custom = xychart.defaultFieldConfig();
        }
        if (!(fillOpacity <= 100)) {
            throw new Error("fillOpacity must be <= 100");
        }
        this.internal.spec.fieldConfig.defaults.custom.fillOpacity = fillOpacity;
        return this;
    }
    lineWidth(lineWidth) {
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
            this.internal.spec.fieldConfig.defaults.custom = xychart.defaultFieldConfig();
        }
        if (!(lineWidth >= 0)) {
            throw new Error("lineWidth must be >= 0");
        }
        this.internal.spec.fieldConfig.defaults.custom.lineWidth = lineWidth;
        return this;
    }
    hideFrom(hideFrom) {
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
            this.internal.spec.fieldConfig.defaults.custom = xychart.defaultFieldConfig();
        }
        const hideFromResource = hideFrom.build();
        this.internal.spec.fieldConfig.defaults.custom.hideFrom = hideFromResource;
        return this;
    }
    axisPlacement(axisPlacement) {
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
            this.internal.spec.fieldConfig.defaults.custom = xychart.defaultFieldConfig();
        }
        this.internal.spec.fieldConfig.defaults.custom.axisPlacement = axisPlacement;
        return this;
    }
    axisColorMode(axisColorMode) {
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
            this.internal.spec.fieldConfig.defaults.custom = xychart.defaultFieldConfig();
        }
        this.internal.spec.fieldConfig.defaults.custom.axisColorMode = axisColorMode;
        return this;
    }
    axisLabel(axisLabel) {
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
            this.internal.spec.fieldConfig.defaults.custom = xychart.defaultFieldConfig();
        }
        this.internal.spec.fieldConfig.defaults.custom.axisLabel = axisLabel;
        return this;
    }
    axisWidth(axisWidth) {
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
            this.internal.spec.fieldConfig.defaults.custom = xychart.defaultFieldConfig();
        }
        this.internal.spec.fieldConfig.defaults.custom.axisWidth = axisWidth;
        return this;
    }
    axisSoftMin(axisSoftMin) {
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
            this.internal.spec.fieldConfig.defaults.custom = xychart.defaultFieldConfig();
        }
        this.internal.spec.fieldConfig.defaults.custom.axisSoftMin = axisSoftMin;
        return this;
    }
    axisSoftMax(axisSoftMax) {
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
            this.internal.spec.fieldConfig.defaults.custom = xychart.defaultFieldConfig();
        }
        this.internal.spec.fieldConfig.defaults.custom.axisSoftMax = axisSoftMax;
        return this;
    }
    axisGridShow(axisGridShow) {
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
            this.internal.spec.fieldConfig.defaults.custom = xychart.defaultFieldConfig();
        }
        this.internal.spec.fieldConfig.defaults.custom.axisGridShow = axisGridShow;
        return this;
    }
    scaleDistribution(scaleDistribution) {
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
            this.internal.spec.fieldConfig.defaults.custom = xychart.defaultFieldConfig();
        }
        const scaleDistributionResource = scaleDistribution.build();
        this.internal.spec.fieldConfig.defaults.custom.scaleDistribution = scaleDistributionResource;
        return this;
    }
    axisCenteredZero(axisCenteredZero) {
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
            this.internal.spec.fieldConfig.defaults.custom = xychart.defaultFieldConfig();
        }
        this.internal.spec.fieldConfig.defaults.custom.axisCenteredZero = axisCenteredZero;
        return this;
    }
    lineStyle(lineStyle) {
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
            this.internal.spec.fieldConfig.defaults.custom = xychart.defaultFieldConfig();
        }
        const lineStyleResource = lineStyle.build();
        this.internal.spec.fieldConfig.defaults.custom.lineStyle = lineStyleResource;
        return this;
    }
    axisBorderShow(axisBorderShow) {
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
            this.internal.spec.fieldConfig.defaults.custom = xychart.defaultFieldConfig();
        }
        this.internal.spec.fieldConfig.defaults.custom.axisBorderShow = axisBorderShow;
        return this;
    }
    mapping(mapping) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultVizConfigSpec();
        }
        if (!this.internal.spec.options) {
            this.internal.spec.options = xychart.defaultOptions();
        }
        this.internal.spec.options.mapping = mapping;
        return this;
    }
    legend(legend) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultVizConfigSpec();
        }
        if (!this.internal.spec.options) {
            this.internal.spec.options = xychart.defaultOptions();
        }
        const legendResource = legend.build();
        this.internal.spec.options.legend = legendResource;
        return this;
    }
    tooltip(tooltip) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultVizConfigSpec();
        }
        if (!this.internal.spec.options) {
            this.internal.spec.options = xychart.defaultOptions();
        }
        const tooltipResource = tooltip.build();
        this.internal.spec.options.tooltip = tooltipResource;
        return this;
    }
    series(series) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultVizConfigSpec();
        }
        if (!this.internal.spec.options) {
            this.internal.spec.options = xychart.defaultOptions();
        }
        const seriesResources = series.map(builder1 => builder1.build());
        this.internal.spec.options.series = seriesResources;
        return this;
    }
}
exports.VisualizationBuilder = VisualizationBuilder;
//# sourceMappingURL=visualizationBuilder.gen.js.map