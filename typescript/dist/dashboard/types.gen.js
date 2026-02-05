"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.defaultAnnotationPanelFilter = exports.defaultAnnotationQuery = exports.defaultAnnotationContainer = exports.defaultVariableSort = exports.VariableSort = exports.defaultVariableRefresh = exports.VariableRefresh = exports.defaultVariableOption = exports.defaultVariableHide = exports.VariableHide = exports.defaultVariableType = exports.VariableType = exports.defaultVariableModel = exports.defaultRowPanel = exports.defaultDynamicConfigValue = exports.defaultFieldColorSeriesByMode = exports.FieldColorSeriesByMode = exports.defaultFieldColorModeId = exports.FieldColorModeId = exports.defaultFieldColor = exports.defaultThreshold = exports.defaultThresholdsMode = exports.ThresholdsMode = exports.defaultThresholdsConfig = exports.defaultSpecialValueMatch = exports.SpecialValueMatch = exports.defaultSpecialValueMap = exports.defaultRegexMap = exports.defaultRangeMap = exports.defaultValueMappingResult = exports.defaultMappingType = exports.MappingType = exports.defaultValueMap = exports.defaultValueMapping = exports.defaultFieldConfig = exports.defaultFieldConfigSource = exports.defaultLibraryPanelRef = exports.defaultMatcherConfig = exports.defaultDataTransformerConfig = exports.defaultDashboardLinkType = exports.DashboardLinkType = exports.defaultDashboardLink = exports.defaultGridPos = exports.defaultDataSourceRef = exports.defaultPanel = exports.defaultTimeOption = exports.defaultTimePickerConfig = exports.defaultDashboardCursorSync = exports.DashboardCursorSync = exports.defaultDashboard = void 0;
exports.defaultDashboardMeta = exports.defaultAnnotationPermission = exports.defaultAnnotationActions = exports.defaultSnapshot = exports.defaultAnnotationTarget = void 0;
const tslib_1 = require("tslib");
const common = tslib_1.__importStar(require("../common"));
const defaultDashboard = () => ({
    timezone: "browser",
    editable: true,
    graphTooltip: DashboardCursorSync.Off,
    fiscalYearStartMonth: 0,
    schemaVersion: 41,
    templating: {},
    annotations: (0, exports.defaultAnnotationContainer)(),
});
exports.defaultDashboard = defaultDashboard;
// 0 for no shared crosshair or tooltip (default).
// 1 for shared crosshair.
// 2 for shared crosshair AND shared tooltip.
var DashboardCursorSync;
(function (DashboardCursorSync) {
    DashboardCursorSync[DashboardCursorSync["Off"] = 0] = "Off";
    DashboardCursorSync[DashboardCursorSync["Crosshair"] = 1] = "Crosshair";
    DashboardCursorSync[DashboardCursorSync["Tooltip"] = 2] = "Tooltip";
})(DashboardCursorSync || (exports.DashboardCursorSync = DashboardCursorSync = {}));
const defaultDashboardCursorSync = () => (DashboardCursorSync.Off);
exports.defaultDashboardCursorSync = defaultDashboardCursorSync;
const defaultTimePickerConfig = () => ({
    hidden: false,
    refresh_intervals: [
        "5s",
        "10s",
        "30s",
        "1m",
        "5m",
        "15m",
        "30m",
        "1h",
        "2h",
        "1d",
    ],
});
exports.defaultTimePickerConfig = defaultTimePickerConfig;
const defaultTimeOption = () => ({
    display: "",
    from: "",
    to: "",
});
exports.defaultTimeOption = defaultTimeOption;
const defaultPanel = () => ({
    type: "",
    transparent: false,
    repeatDirection: "h",
});
exports.defaultPanel = defaultPanel;
const defaultDataSourceRef = () => (common.defaultDataSourceRef());
exports.defaultDataSourceRef = defaultDataSourceRef;
const defaultGridPos = () => ({
    h: 9,
    w: 12,
    x: 0,
    y: 0,
});
exports.defaultGridPos = defaultGridPos;
const defaultDashboardLink = () => ({
    title: "",
    type: DashboardLinkType.Link,
    icon: "",
    tooltip: "",
    tags: [],
    asDropdown: false,
    targetBlank: false,
    includeVars: false,
    keepTime: false,
});
exports.defaultDashboardLink = defaultDashboardLink;
// Dashboard Link type. Accepted values are dashboards (to refer to another dashboard) and link (to refer to an external resource)
var DashboardLinkType;
(function (DashboardLinkType) {
    DashboardLinkType["Link"] = "link";
    DashboardLinkType["Dashboards"] = "dashboards";
})(DashboardLinkType || (exports.DashboardLinkType = DashboardLinkType = {}));
const defaultDashboardLinkType = () => (DashboardLinkType.Link);
exports.defaultDashboardLinkType = defaultDashboardLinkType;
const defaultDataTransformerConfig = () => ({
    id: "",
    options: {},
});
exports.defaultDataTransformerConfig = defaultDataTransformerConfig;
const defaultMatcherConfig = () => ({
    id: "",
});
exports.defaultMatcherConfig = defaultMatcherConfig;
const defaultLibraryPanelRef = () => ({
    name: "",
    uid: "",
});
exports.defaultLibraryPanelRef = defaultLibraryPanelRef;
const defaultFieldConfigSource = () => ({
    defaults: (0, exports.defaultFieldConfig)(),
    overrides: [],
});
exports.defaultFieldConfigSource = defaultFieldConfigSource;
const defaultFieldConfig = () => ({});
exports.defaultFieldConfig = defaultFieldConfig;
const defaultValueMapping = () => ((0, exports.defaultValueMap)());
exports.defaultValueMapping = defaultValueMapping;
const defaultValueMap = () => ({
    type: MappingType.ValueToText,
    options: {},
});
exports.defaultValueMap = defaultValueMap;
// Supported value mapping types
// `value`: Maps text values to a color or different display text and color. For example, you can configure a value mapping so that all instances of the value 10 appear as Perfection! rather than the number.
// `range`: Maps numerical ranges to a display text and color. For example, if a value is within a certain range, you can configure a range value mapping to display Low or High rather than the number.
// `regex`: Maps regular expressions to replacement text and a color. For example, if a value is www.example.com, you can configure a regex value mapping so that Grafana displays www and truncates the domain.
// `special`: Maps special values like Null, NaN (not a number), and boolean values like true and false to a display text and color. See SpecialValueMatch to see the list of special values. For example, you can configure a special value mapping so that null values appear as N/A.
var MappingType;
(function (MappingType) {
    MappingType["ValueToText"] = "value";
    MappingType["RangeToText"] = "range";
    MappingType["RegexToText"] = "regex";
    MappingType["SpecialValue"] = "special";
})(MappingType || (exports.MappingType = MappingType = {}));
const defaultMappingType = () => (MappingType.ValueToText);
exports.defaultMappingType = defaultMappingType;
const defaultValueMappingResult = () => ({});
exports.defaultValueMappingResult = defaultValueMappingResult;
const defaultRangeMap = () => ({
    type: MappingType.RangeToText,
    options: {
        from: 0,
        to: 0,
        result: (0, exports.defaultValueMappingResult)(),
    },
});
exports.defaultRangeMap = defaultRangeMap;
const defaultRegexMap = () => ({
    type: MappingType.RegexToText,
    options: {
        pattern: "",
        result: (0, exports.defaultValueMappingResult)(),
    },
});
exports.defaultRegexMap = defaultRegexMap;
const defaultSpecialValueMap = () => ({
    type: MappingType.SpecialValue,
    options: {
        match: SpecialValueMatch.True,
        result: (0, exports.defaultValueMappingResult)(),
    },
});
exports.defaultSpecialValueMap = defaultSpecialValueMap;
// Special value types supported by the `SpecialValueMap`
var SpecialValueMatch;
(function (SpecialValueMatch) {
    SpecialValueMatch["True"] = "true";
    SpecialValueMatch["False"] = "false";
    SpecialValueMatch["Null"] = "null";
    SpecialValueMatch["NotANumber"] = "nan";
    SpecialValueMatch["NullAndNan"] = "null+nan";
    SpecialValueMatch["Empty"] = "empty";
})(SpecialValueMatch || (exports.SpecialValueMatch = SpecialValueMatch = {}));
const defaultSpecialValueMatch = () => (SpecialValueMatch.True);
exports.defaultSpecialValueMatch = defaultSpecialValueMatch;
const defaultThresholdsConfig = () => ({
    mode: ThresholdsMode.Absolute,
    steps: [],
});
exports.defaultThresholdsConfig = defaultThresholdsConfig;
// Thresholds can either be `absolute` (specific number) or `percentage` (relative to min or max, it will be values between 0 and 1).
var ThresholdsMode;
(function (ThresholdsMode) {
    ThresholdsMode["Absolute"] = "absolute";
    ThresholdsMode["Percentage"] = "percentage";
})(ThresholdsMode || (exports.ThresholdsMode = ThresholdsMode = {}));
const defaultThresholdsMode = () => (ThresholdsMode.Absolute);
exports.defaultThresholdsMode = defaultThresholdsMode;
const defaultThreshold = () => ({
    value: 0,
    color: "",
});
exports.defaultThreshold = defaultThreshold;
const defaultFieldColor = () => ({
    mode: FieldColorModeId.Thresholds,
});
exports.defaultFieldColor = defaultFieldColor;
// Color mode for a field. You can specify a single color, or select a continuous (gradient) color schemes, based on a value.
// Continuous color interpolates a color using the percentage of a value relative to min and max.
// Accepted values are:
// `thresholds`: From thresholds. Informs Grafana to take the color from the matching threshold
// `palette-classic`: Classic palette. Grafana will assign color by looking up a color in a palette by series index. Useful for Graphs and pie charts and other categorical data visualizations
// `palette-classic-by-name`: Classic palette (by name). Grafana will assign color by looking up a color in a palette by series name. Useful for Graphs and pie charts and other categorical data visualizations
// `continuous-GrYlRd`: ontinuous Green-Yellow-Red palette mode
// `continuous-RdYlGr`: Continuous Red-Yellow-Green palette mode
// `continuous-BlYlRd`: Continuous Blue-Yellow-Red palette mode
// `continuous-YlRd`: Continuous Yellow-Red palette mode
// `continuous-BlPu`: Continuous Blue-Purple palette mode
// `continuous-YlBl`: Continuous Yellow-Blue palette mode
// `continuous-blues`: Continuous Blue palette mode
// `continuous-reds`: Continuous Red palette mode
// `continuous-greens`: Continuous Green palette mode
// `continuous-purples`: Continuous Purple palette mode
// `shades`: Shades of a single color. Specify a single color, useful in an override rule.
// `fixed`: Fixed color mode. Specify a single color, useful in an override rule.
var FieldColorModeId;
(function (FieldColorModeId) {
    FieldColorModeId["Thresholds"] = "thresholds";
    FieldColorModeId["PaletteClassic"] = "palette-classic";
    FieldColorModeId["PaletteClassicByName"] = "palette-classic-by-name";
    FieldColorModeId["ContinuousGrYlRd"] = "continuous-GrYlRd";
    FieldColorModeId["ContinuousRdYlGr"] = "continuous-RdYlGr";
    FieldColorModeId["ContinuousBlYlRd"] = "continuous-BlYlRd";
    FieldColorModeId["ContinuousYlRd"] = "continuous-YlRd";
    FieldColorModeId["ContinuousBlPu"] = "continuous-BlPu";
    FieldColorModeId["ContinuousYlBl"] = "continuous-YlBl";
    FieldColorModeId["ContinuousBlues"] = "continuous-blues";
    FieldColorModeId["ContinuousReds"] = "continuous-reds";
    FieldColorModeId["ContinuousGreens"] = "continuous-greens";
    FieldColorModeId["ContinuousPurples"] = "continuous-purples";
    FieldColorModeId["Fixed"] = "fixed";
    FieldColorModeId["Shades"] = "shades";
})(FieldColorModeId || (exports.FieldColorModeId = FieldColorModeId = {}));
const defaultFieldColorModeId = () => (FieldColorModeId.Thresholds);
exports.defaultFieldColorModeId = defaultFieldColorModeId;
// Defines how to assign a series color from "by value" color schemes. For example for an aggregated data points like a timeseries, the color can be assigned by the min, max or last value.
var FieldColorSeriesByMode;
(function (FieldColorSeriesByMode) {
    FieldColorSeriesByMode["Min"] = "min";
    FieldColorSeriesByMode["Max"] = "max";
    FieldColorSeriesByMode["Last"] = "last";
})(FieldColorSeriesByMode || (exports.FieldColorSeriesByMode = FieldColorSeriesByMode = {}));
const defaultFieldColorSeriesByMode = () => (FieldColorSeriesByMode.Min);
exports.defaultFieldColorSeriesByMode = defaultFieldColorSeriesByMode;
const defaultDynamicConfigValue = () => ({
    id: "",
});
exports.defaultDynamicConfigValue = defaultDynamicConfigValue;
const defaultRowPanel = () => ({
    type: "row",
    collapsed: false,
    id: 0,
    panels: [],
});
exports.defaultRowPanel = defaultRowPanel;
const defaultVariableModel = () => ({
    type: VariableType.Query,
    name: "",
    skipUrlSync: false,
    multi: false,
    allowCustomValue: true,
    includeAll: false,
    auto: false,
    auto_min: "10s",
    auto_count: 30,
});
exports.defaultVariableModel = defaultVariableModel;
// Dashboard variable type
// `query`: Query-generated list of values such as metric names, server names, sensor IDs, data centers, and so on.
// `adhoc`: Key/value filters that are automatically added to all metric queries for a data source (Prometheus, Loki, InfluxDB, and Elasticsearch only).
// `constant`: 	Define a hidden constant.
// `datasource`: Quickly change the data source for an entire dashboard.
// `interval`: Interval variables represent time spans.
// `textbox`: Display a free text input field with an optional default value.
// `custom`: Define the variable options manually using a comma-separated list.
// `system`: Variables defined by Grafana. See: https://grafana.com/docs/grafana/latest/dashboards/variables/add-template-variables/#global-variables
var VariableType;
(function (VariableType) {
    VariableType["Query"] = "query";
    VariableType["Adhoc"] = "adhoc";
    VariableType["Groupby"] = "groupby";
    VariableType["Constant"] = "constant";
    VariableType["Datasource"] = "datasource";
    VariableType["Interval"] = "interval";
    VariableType["Textbox"] = "textbox";
    VariableType["Custom"] = "custom";
    VariableType["System"] = "system";
    VariableType["Snapshot"] = "snapshot";
})(VariableType || (exports.VariableType = VariableType = {}));
const defaultVariableType = () => (VariableType.Query);
exports.defaultVariableType = defaultVariableType;
// Determine if the variable shows on dashboard
// Accepted values are 0 (show label and value), 1 (show value only), 2 (show nothing).
var VariableHide;
(function (VariableHide) {
    VariableHide[VariableHide["DontHide"] = 0] = "DontHide";
    VariableHide[VariableHide["HideLabel"] = 1] = "HideLabel";
    VariableHide[VariableHide["HideVariable"] = 2] = "HideVariable";
})(VariableHide || (exports.VariableHide = VariableHide = {}));
const defaultVariableHide = () => (VariableHide.DontHide);
exports.defaultVariableHide = defaultVariableHide;
const defaultVariableOption = () => ({
    text: "",
    value: "",
});
exports.defaultVariableOption = defaultVariableOption;
// Options to config when to refresh a variable
// `0`: Never refresh the variable
// `1`: Queries the data source every time the dashboard loads.
// `2`: Queries the data source when the dashboard time range changes.
var VariableRefresh;
(function (VariableRefresh) {
    VariableRefresh[VariableRefresh["Never"] = 0] = "Never";
    VariableRefresh[VariableRefresh["OnDashboardLoad"] = 1] = "OnDashboardLoad";
    VariableRefresh[VariableRefresh["OnTimeRangeChanged"] = 2] = "OnTimeRangeChanged";
})(VariableRefresh || (exports.VariableRefresh = VariableRefresh = {}));
const defaultVariableRefresh = () => (VariableRefresh.Never);
exports.defaultVariableRefresh = defaultVariableRefresh;
// Sort variable options
// Accepted values are:
// `0`: No sorting
// `1`: Alphabetical ASC
// `2`: Alphabetical DESC
// `3`: Numerical ASC
// `4`: Numerical DESC
// `5`: Alphabetical Case Insensitive ASC
// `6`: Alphabetical Case Insensitive DESC
// `7`: Natural ASC
// `8`: Natural DESC
var VariableSort;
(function (VariableSort) {
    VariableSort[VariableSort["Disabled"] = 0] = "Disabled";
    VariableSort[VariableSort["AlphabeticalAsc"] = 1] = "AlphabeticalAsc";
    VariableSort[VariableSort["AlphabeticalDesc"] = 2] = "AlphabeticalDesc";
    VariableSort[VariableSort["NumericalAsc"] = 3] = "NumericalAsc";
    VariableSort[VariableSort["NumericalDesc"] = 4] = "NumericalDesc";
    VariableSort[VariableSort["AlphabeticalCaseInsensitiveAsc"] = 5] = "AlphabeticalCaseInsensitiveAsc";
    VariableSort[VariableSort["AlphabeticalCaseInsensitiveDesc"] = 6] = "AlphabeticalCaseInsensitiveDesc";
    VariableSort[VariableSort["NaturalAsc"] = 7] = "NaturalAsc";
    VariableSort[VariableSort["NaturalDesc"] = 8] = "NaturalDesc";
})(VariableSort || (exports.VariableSort = VariableSort = {}));
const defaultVariableSort = () => (VariableSort.Disabled);
exports.defaultVariableSort = defaultVariableSort;
const defaultAnnotationContainer = () => ({});
exports.defaultAnnotationContainer = defaultAnnotationContainer;
const defaultAnnotationQuery = () => ({
    name: "",
    datasource: common.defaultDataSourceRef(),
    enable: true,
    hide: false,
    iconColor: "",
    builtIn: 0,
});
exports.defaultAnnotationQuery = defaultAnnotationQuery;
const defaultAnnotationPanelFilter = () => ({
    exclude: false,
    ids: [],
});
exports.defaultAnnotationPanelFilter = defaultAnnotationPanelFilter;
const defaultAnnotationTarget = () => ({
    limit: 0,
    matchAny: false,
    tags: [],
    type: "",
});
exports.defaultAnnotationTarget = defaultAnnotationTarget;
const defaultSnapshot = () => ({
    created: "",
    expires: "",
    external: false,
    externalUrl: "",
    originalUrl: "",
    id: 0,
    key: "",
    name: "",
    orgId: 0,
    updated: "",
    userId: 0,
});
exports.defaultSnapshot = defaultSnapshot;
const defaultAnnotationActions = () => ({});
exports.defaultAnnotationActions = defaultAnnotationActions;
const defaultAnnotationPermission = () => ({});
exports.defaultAnnotationPermission = defaultAnnotationPermission;
const defaultDashboardMeta = () => ({});
exports.defaultDashboardMeta = defaultDashboardMeta;
//# sourceMappingURL=types.gen.js.map