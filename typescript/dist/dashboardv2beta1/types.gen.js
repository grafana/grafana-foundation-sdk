"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.HttpRequestMethod = exports.defaultFetchOptions = exports.defaultActionType = exports.ActionType = exports.defaultAction = exports.defaultFieldColorSeriesByMode = exports.FieldColorSeriesByMode = exports.defaultFieldColorModeId = exports.FieldColorModeId = exports.defaultFieldColor = exports.defaultThreshold = exports.defaultThresholdsMode = exports.ThresholdsMode = exports.defaultThresholdsConfig = exports.defaultSpecialValueMatch = exports.SpecialValueMatch = exports.defaultSpecialValueMap = exports.defaultRegexMap = exports.defaultRangeMap = exports.defaultValueMappingResult = exports.defaultMappingType = exports.MappingType = exports.defaultValueMap = exports.defaultValueMapping = exports.defaultFieldConfig = exports.defaultFieldConfigSource = exports.defaultVizConfigSpec = exports.defaultVizConfigKind = exports.defaultQueryOptionsSpec = exports.defaultDataTopic = exports.DataTopic = exports.defaultMatcherConfig = exports.defaultDataTransformerConfig = exports.defaultTransformationKind = exports.defaultPanelQuerySpec = exports.defaultPanelQueryKind = exports.defaultQueryGroupSpec = exports.defaultQueryGroupKind = exports.defaultDataLink = exports.defaultPanelSpec = exports.defaultPanelKind = exports.defaultElement = exports.defaultDashboardCursorSync = exports.DashboardCursorSync = exports.AnnotationQueryPlacement = exports.defaultAnnotationPanelFilter = exports.defaultDataQueryKind = exports.defaultAnnotationQuerySpec = exports.defaultAnnotationQueryKind = exports.defaultDashboard = void 0;
exports.VariableHide = exports.defaultVariableOption = exports.defaultQueryVariableSpec = exports.defaultQueryVariableKind = exports.defaultVariableKind = exports.defaultTimeRangeOption = exports.defaultTimeSettingsSpec = exports.DashboardLinkPlacement = exports.defaultDashboardLinkType = exports.DashboardLinkType = exports.defaultDashboardLink = exports.defaultTabRepeatOptions = exports.defaultTabsLayoutTabSpec = exports.defaultTabsLayoutTabKind = exports.defaultTabsLayoutSpec = exports.defaultTabsLayoutKind = exports.defaultAutoGridRepeatOptions = exports.defaultAutoGridLayoutItemSpec = exports.defaultAutoGridLayoutItemKind = exports.defaultAutoGridLayoutSpec = exports.defaultAutoGridLayoutKind = exports.defaultRowRepeatOptions = exports.defaultConditionalRenderingTimeRangeSizeSpec = exports.defaultConditionalRenderingTimeRangeSizeKind = exports.defaultConditionalRenderingDataSpec = exports.defaultConditionalRenderingDataKind = exports.defaultConditionalRenderingVariableSpec = exports.defaultConditionalRenderingVariableKind = exports.defaultConditionalRenderingGroupSpec = exports.defaultConditionalRenderingGroupKind = exports.defaultRowsLayoutRowSpec = exports.defaultRowsLayoutRowKind = exports.defaultRowsLayoutSpec = exports.defaultRowsLayoutKind = exports.defaultRepeatMode = exports.RepeatMode = exports.defaultRepeatOptions = exports.defaultElementReference = exports.defaultGridLayoutItemSpec = exports.defaultGridLayoutItemKind = exports.defaultGridLayoutSpec = exports.defaultGridLayoutKind = exports.defaultLibraryPanelRef = exports.defaultLibraryPanelKindSpec = exports.defaultLibraryPanelKind = exports.defaultDynamicConfigValue = exports.ActionVariableType = exports.defaultActionVariable = exports.defaultInfinityOptions = exports.defaultHttpRequestMethod = void 0;
exports.defaultVariableValueOption = exports.defaultCustomFormatterVariable = exports.defaultVariableType = exports.VariableType = exports.defaultCustomVariableValue = exports.defaultVariableValueSingle = exports.defaultVariableValue = exports.defaultKind = exports.defaultSwitchVariableSpec = exports.defaultSwitchVariableKind = exports.defaultMetricFindValue = exports.FilterOrigin = exports.defaultAdHocFilterWithLabels = exports.defaultAdhocVariableSpec = exports.defaultAdhocVariableKind = exports.defaultGroupByVariableSpec = exports.defaultGroupByVariableKind = exports.defaultCustomVariableSpec = exports.defaultCustomVariableKind = exports.defaultIntervalVariableSpec = exports.defaultIntervalVariableKind = exports.defaultDatasourceVariableSpec = exports.defaultDatasourceVariableKind = exports.defaultConstantVariableSpec = exports.defaultConstantVariableKind = exports.defaultTextVariableSpec = exports.defaultTextVariableKind = exports.defaultVariableSort = exports.VariableSort = exports.defaultVariableRefresh = exports.VariableRefresh = exports.defaultVariableHide = void 0;
const defaultDashboard = () => ({
    annotations: [],
    cursorSync: DashboardCursorSync.Off,
    editable: true,
    elements: {},
    layout: (0, exports.defaultGridLayoutKind)(),
    links: [],
    preload: false,
    tags: [],
    timeSettings: (0, exports.defaultTimeSettingsSpec)(),
    title: "",
    variables: [],
});
exports.defaultDashboard = defaultDashboard;
const defaultAnnotationQueryKind = () => ({
    kind: "AnnotationQuery",
    spec: (0, exports.defaultAnnotationQuerySpec)(),
});
exports.defaultAnnotationQueryKind = defaultAnnotationQueryKind;
const defaultAnnotationQuerySpec = () => ({
    query: (0, exports.defaultDataQueryKind)(),
    enable: false,
    hide: false,
    iconColor: "",
    name: "",
    builtIn: false,
    placement: exports.AnnotationQueryPlacement,
});
exports.defaultAnnotationQuerySpec = defaultAnnotationQuerySpec;
const defaultDataQueryKind = () => ({
    kind: "DataQuery",
    group: "",
    version: "v0",
    spec: {},
});
exports.defaultDataQueryKind = defaultDataQueryKind;
const defaultAnnotationPanelFilter = () => ({
    exclude: false,
    ids: [],
});
exports.defaultAnnotationPanelFilter = defaultAnnotationPanelFilter;
// Annotation Query placement. Defines where the annotation query should be displayed.
// - "inControlsMenu" renders the annotation query in the dashboard controls dropdown menu
exports.AnnotationQueryPlacement = "inControlsMenu";
// "Off" for no shared crosshair or tooltip (default).
// "Crosshair" for shared crosshair.
// "Tooltip" for shared crosshair AND shared tooltip.
var DashboardCursorSync;
(function (DashboardCursorSync) {
    DashboardCursorSync["Crosshair"] = "Crosshair";
    DashboardCursorSync["Tooltip"] = "Tooltip";
    DashboardCursorSync["Off"] = "Off";
})(DashboardCursorSync || (exports.DashboardCursorSync = DashboardCursorSync = {}));
const defaultDashboardCursorSync = () => (DashboardCursorSync.Off);
exports.defaultDashboardCursorSync = defaultDashboardCursorSync;
const defaultElement = () => ((0, exports.defaultPanelKind)());
exports.defaultElement = defaultElement;
const defaultPanelKind = () => ({
    kind: "Panel",
    spec: (0, exports.defaultPanelSpec)(),
});
exports.defaultPanelKind = defaultPanelKind;
const defaultPanelSpec = () => ({
    id: 0,
    title: "",
    description: "",
    links: [],
    data: (0, exports.defaultQueryGroupKind)(),
    vizConfig: (0, exports.defaultVizConfigKind)(),
});
exports.defaultPanelSpec = defaultPanelSpec;
const defaultDataLink = () => ({
    title: "",
    url: "",
});
exports.defaultDataLink = defaultDataLink;
const defaultQueryGroupKind = () => ({
    kind: "QueryGroup",
    spec: (0, exports.defaultQueryGroupSpec)(),
});
exports.defaultQueryGroupKind = defaultQueryGroupKind;
const defaultQueryGroupSpec = () => ({
    queries: [],
    transformations: [],
    queryOptions: (0, exports.defaultQueryOptionsSpec)(),
});
exports.defaultQueryGroupSpec = defaultQueryGroupSpec;
const defaultPanelQueryKind = () => ({
    kind: "PanelQuery",
    spec: (0, exports.defaultPanelQuerySpec)(),
});
exports.defaultPanelQueryKind = defaultPanelQueryKind;
const defaultPanelQuerySpec = () => ({
    query: (0, exports.defaultDataQueryKind)(),
    refId: "",
    hidden: false,
});
exports.defaultPanelQuerySpec = defaultPanelQuerySpec;
const defaultTransformationKind = () => ({
    kind: "",
    spec: (0, exports.defaultDataTransformerConfig)(),
});
exports.defaultTransformationKind = defaultTransformationKind;
const defaultDataTransformerConfig = () => ({
    id: "",
    options: {},
});
exports.defaultDataTransformerConfig = defaultDataTransformerConfig;
const defaultMatcherConfig = () => ({
    id: "",
});
exports.defaultMatcherConfig = defaultMatcherConfig;
// A topic is attached to DataFrame metadata in query results.
// This specifies where the data should be used.
var DataTopic;
(function (DataTopic) {
    DataTopic["Series"] = "series";
    DataTopic["Annotations"] = "annotations";
    DataTopic["AlertStates"] = "alertStates";
})(DataTopic || (exports.DataTopic = DataTopic = {}));
const defaultDataTopic = () => (DataTopic.Series);
exports.defaultDataTopic = defaultDataTopic;
const defaultQueryOptionsSpec = () => ({});
exports.defaultQueryOptionsSpec = defaultQueryOptionsSpec;
const defaultVizConfigKind = () => ({
    kind: "VizConfig",
    group: "",
    version: "",
    spec: (0, exports.defaultVizConfigSpec)(),
});
exports.defaultVizConfigKind = defaultVizConfigKind;
const defaultVizConfigSpec = () => ({
    options: {},
    fieldConfig: (0, exports.defaultFieldConfigSource)(),
});
exports.defaultVizConfigSpec = defaultVizConfigSpec;
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
    type: MappingType.Value,
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
    MappingType["Value"] = "value";
    MappingType["Range"] = "range";
    MappingType["Regex"] = "regex";
    MappingType["Special"] = "special";
})(MappingType || (exports.MappingType = MappingType = {}));
const defaultMappingType = () => (MappingType.Value);
exports.defaultMappingType = defaultMappingType;
const defaultValueMappingResult = () => ({});
exports.defaultValueMappingResult = defaultValueMappingResult;
const defaultRangeMap = () => ({
    type: MappingType.Range,
    options: {
        from: 0,
        to: 0,
        result: (0, exports.defaultValueMappingResult)(),
    },
});
exports.defaultRangeMap = defaultRangeMap;
const defaultRegexMap = () => ({
    type: MappingType.Regex,
    options: {
        pattern: "",
        result: (0, exports.defaultValueMappingResult)(),
    },
});
exports.defaultRegexMap = defaultRegexMap;
const defaultSpecialValueMap = () => ({
    type: MappingType.Special,
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
    SpecialValueMatch["NullAndNaN"] = "null+nan";
    SpecialValueMatch["Empty"] = "empty";
})(SpecialValueMatch || (exports.SpecialValueMatch = SpecialValueMatch = {}));
const defaultSpecialValueMatch = () => (SpecialValueMatch.True);
exports.defaultSpecialValueMatch = defaultSpecialValueMatch;
const defaultThresholdsConfig = () => ({
    mode: ThresholdsMode.Absolute,
    steps: [],
});
exports.defaultThresholdsConfig = defaultThresholdsConfig;
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
const defaultAction = () => ({
    type: ActionType.Fetch,
    title: "",
});
exports.defaultAction = defaultAction;
var ActionType;
(function (ActionType) {
    ActionType["Fetch"] = "fetch";
    ActionType["Infinity"] = "infinity";
})(ActionType || (exports.ActionType = ActionType = {}));
const defaultActionType = () => (ActionType.Fetch);
exports.defaultActionType = defaultActionType;
const defaultFetchOptions = () => ({
    method: HttpRequestMethod.GET,
    url: "",
});
exports.defaultFetchOptions = defaultFetchOptions;
var HttpRequestMethod;
(function (HttpRequestMethod) {
    HttpRequestMethod["GET"] = "GET";
    HttpRequestMethod["PUT"] = "PUT";
    HttpRequestMethod["POST"] = "POST";
    HttpRequestMethod["DELETE"] = "DELETE";
    HttpRequestMethod["PATCH"] = "PATCH";
})(HttpRequestMethod || (exports.HttpRequestMethod = HttpRequestMethod = {}));
const defaultHttpRequestMethod = () => (HttpRequestMethod.GET);
exports.defaultHttpRequestMethod = defaultHttpRequestMethod;
const defaultInfinityOptions = () => ({
    method: HttpRequestMethod.GET,
    url: "",
    datasourceUid: "",
});
exports.defaultInfinityOptions = defaultInfinityOptions;
const defaultActionVariable = () => ({
    key: "",
    name: "",
    type: exports.ActionVariableType,
});
exports.defaultActionVariable = defaultActionVariable;
// Action variable type
exports.ActionVariableType = "string";
const defaultDynamicConfigValue = () => ({
    id: "",
});
exports.defaultDynamicConfigValue = defaultDynamicConfigValue;
const defaultLibraryPanelKind = () => ({
    kind: "LibraryPanel",
    spec: (0, exports.defaultLibraryPanelKindSpec)(),
});
exports.defaultLibraryPanelKind = defaultLibraryPanelKind;
const defaultLibraryPanelKindSpec = () => ({
    id: 0,
    title: "",
    libraryPanel: (0, exports.defaultLibraryPanelRef)(),
});
exports.defaultLibraryPanelKindSpec = defaultLibraryPanelKindSpec;
const defaultLibraryPanelRef = () => ({
    name: "",
    uid: "",
});
exports.defaultLibraryPanelRef = defaultLibraryPanelRef;
const defaultGridLayoutKind = () => ({
    kind: "GridLayout",
    spec: (0, exports.defaultGridLayoutSpec)(),
});
exports.defaultGridLayoutKind = defaultGridLayoutKind;
const defaultGridLayoutSpec = () => ({
    items: [],
});
exports.defaultGridLayoutSpec = defaultGridLayoutSpec;
const defaultGridLayoutItemKind = () => ({
    kind: "GridLayoutItem",
    spec: (0, exports.defaultGridLayoutItemSpec)(),
});
exports.defaultGridLayoutItemKind = defaultGridLayoutItemKind;
const defaultGridLayoutItemSpec = () => ({
    x: 0,
    y: 0,
    width: 0,
    height: 0,
    element: (0, exports.defaultElementReference)(),
});
exports.defaultGridLayoutItemSpec = defaultGridLayoutItemSpec;
const defaultElementReference = () => ({
    kind: "ElementReference",
    name: "",
});
exports.defaultElementReference = defaultElementReference;
const defaultRepeatOptions = () => ({
    mode: RepeatMode.Variable,
    value: "",
});
exports.defaultRepeatOptions = defaultRepeatOptions;
// other repeat modes will be added in the future: label, frame
var RepeatMode;
(function (RepeatMode) {
    RepeatMode["Variable"] = "variable";
})(RepeatMode || (exports.RepeatMode = RepeatMode = {}));
const defaultRepeatMode = () => (RepeatMode.Variable);
exports.defaultRepeatMode = defaultRepeatMode;
const defaultRowsLayoutKind = () => ({
    kind: "RowsLayout",
    spec: (0, exports.defaultRowsLayoutSpec)(),
});
exports.defaultRowsLayoutKind = defaultRowsLayoutKind;
const defaultRowsLayoutSpec = () => ({
    rows: [],
});
exports.defaultRowsLayoutSpec = defaultRowsLayoutSpec;
const defaultRowsLayoutRowKind = () => ({
    kind: "RowsLayoutRow",
    spec: (0, exports.defaultRowsLayoutRowSpec)(),
});
exports.defaultRowsLayoutRowKind = defaultRowsLayoutRowKind;
const defaultRowsLayoutRowSpec = () => ({
    layout: (0, exports.defaultGridLayoutKind)(),
});
exports.defaultRowsLayoutRowSpec = defaultRowsLayoutRowSpec;
const defaultConditionalRenderingGroupKind = () => ({
    kind: "ConditionalRenderingGroup",
    spec: (0, exports.defaultConditionalRenderingGroupSpec)(),
});
exports.defaultConditionalRenderingGroupKind = defaultConditionalRenderingGroupKind;
const defaultConditionalRenderingGroupSpec = () => ({
    visibility: "show",
    condition: "and",
    items: [],
});
exports.defaultConditionalRenderingGroupSpec = defaultConditionalRenderingGroupSpec;
const defaultConditionalRenderingVariableKind = () => ({
    kind: "ConditionalRenderingVariable",
    spec: (0, exports.defaultConditionalRenderingVariableSpec)(),
});
exports.defaultConditionalRenderingVariableKind = defaultConditionalRenderingVariableKind;
const defaultConditionalRenderingVariableSpec = () => ({
    variable: "",
    operator: "equals",
    value: "",
});
exports.defaultConditionalRenderingVariableSpec = defaultConditionalRenderingVariableSpec;
const defaultConditionalRenderingDataKind = () => ({
    kind: "ConditionalRenderingData",
    spec: (0, exports.defaultConditionalRenderingDataSpec)(),
});
exports.defaultConditionalRenderingDataKind = defaultConditionalRenderingDataKind;
const defaultConditionalRenderingDataSpec = () => ({
    value: false,
});
exports.defaultConditionalRenderingDataSpec = defaultConditionalRenderingDataSpec;
const defaultConditionalRenderingTimeRangeSizeKind = () => ({
    kind: "ConditionalRenderingTimeRangeSize",
    spec: (0, exports.defaultConditionalRenderingTimeRangeSizeSpec)(),
});
exports.defaultConditionalRenderingTimeRangeSizeKind = defaultConditionalRenderingTimeRangeSizeKind;
const defaultConditionalRenderingTimeRangeSizeSpec = () => ({
    value: "",
});
exports.defaultConditionalRenderingTimeRangeSizeSpec = defaultConditionalRenderingTimeRangeSizeSpec;
const defaultRowRepeatOptions = () => ({
    mode: RepeatMode.Variable,
    value: "",
});
exports.defaultRowRepeatOptions = defaultRowRepeatOptions;
const defaultAutoGridLayoutKind = () => ({
    kind: "AutoGridLayout",
    spec: (0, exports.defaultAutoGridLayoutSpec)(),
});
exports.defaultAutoGridLayoutKind = defaultAutoGridLayoutKind;
const defaultAutoGridLayoutSpec = () => ({
    maxColumnCount: 3,
    columnWidthMode: "standard",
    rowHeightMode: "standard",
    fillScreen: false,
    items: [],
});
exports.defaultAutoGridLayoutSpec = defaultAutoGridLayoutSpec;
const defaultAutoGridLayoutItemKind = () => ({
    kind: "AutoGridLayoutItem",
    spec: (0, exports.defaultAutoGridLayoutItemSpec)(),
});
exports.defaultAutoGridLayoutItemKind = defaultAutoGridLayoutItemKind;
const defaultAutoGridLayoutItemSpec = () => ({
    element: (0, exports.defaultElementReference)(),
});
exports.defaultAutoGridLayoutItemSpec = defaultAutoGridLayoutItemSpec;
const defaultAutoGridRepeatOptions = () => ({
    mode: RepeatMode.Variable,
    value: "",
});
exports.defaultAutoGridRepeatOptions = defaultAutoGridRepeatOptions;
const defaultTabsLayoutKind = () => ({
    kind: "TabsLayout",
    spec: (0, exports.defaultTabsLayoutSpec)(),
});
exports.defaultTabsLayoutKind = defaultTabsLayoutKind;
const defaultTabsLayoutSpec = () => ({
    tabs: [],
});
exports.defaultTabsLayoutSpec = defaultTabsLayoutSpec;
const defaultTabsLayoutTabKind = () => ({
    kind: "TabsLayoutTab",
    spec: (0, exports.defaultTabsLayoutTabSpec)(),
});
exports.defaultTabsLayoutTabKind = defaultTabsLayoutTabKind;
const defaultTabsLayoutTabSpec = () => ({
    layout: (0, exports.defaultGridLayoutKind)(),
});
exports.defaultTabsLayoutTabSpec = defaultTabsLayoutTabSpec;
const defaultTabRepeatOptions = () => ({
    mode: RepeatMode.Variable,
    value: "",
});
exports.defaultTabRepeatOptions = defaultTabRepeatOptions;
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
    placement: exports.DashboardLinkPlacement,
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
// Dashboard Link placement. Defines where the link should be displayed.
// - "inControlsMenu" renders the link in bottom part of the dashboard controls dropdown menu
exports.DashboardLinkPlacement = "inControlsMenu";
const defaultTimeSettingsSpec = () => ({
    timezone: "browser",
    from: "now-6h",
    to: "now",
    autoRefresh: "",
    autoRefreshIntervals: [
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
    hideTimepicker: false,
    fiscalYearStartMonth: 0,
});
exports.defaultTimeSettingsSpec = defaultTimeSettingsSpec;
const defaultTimeRangeOption = () => ({
    display: "Last 6 hours",
    from: "now-6h",
    to: "now",
});
exports.defaultTimeRangeOption = defaultTimeRangeOption;
const defaultVariableKind = () => ((0, exports.defaultQueryVariableKind)());
exports.defaultVariableKind = defaultVariableKind;
const defaultQueryVariableKind = () => ({
    kind: "QueryVariable",
    spec: (0, exports.defaultQueryVariableSpec)(),
});
exports.defaultQueryVariableKind = defaultQueryVariableKind;
const defaultQueryVariableSpec = () => ({
    name: "",
    current: { text: "", value: "", },
    hide: VariableHide.DontHide,
    refresh: VariableRefresh.Never,
    skipUrlSync: false,
    query: (0, exports.defaultDataQueryKind)(),
    regex: "",
    sort: VariableSort.Disabled,
    options: [],
    multi: false,
    includeAll: false,
    allowCustomValue: true,
});
exports.defaultQueryVariableSpec = defaultQueryVariableSpec;
const defaultVariableOption = () => ({
    text: "",
    value: "",
});
exports.defaultVariableOption = defaultVariableOption;
// Determine if the variable shows on dashboard
// Accepted values are `dontHide` (show label and value), `hideLabel` (show value only), `hideVariable` (show nothing), `inControlsMenu` (show in a drop-down menu).
var VariableHide;
(function (VariableHide) {
    VariableHide["DontHide"] = "dontHide";
    VariableHide["HideLabel"] = "hideLabel";
    VariableHide["HideVariable"] = "hideVariable";
    VariableHide["InControlsMenu"] = "inControlsMenu";
})(VariableHide || (exports.VariableHide = VariableHide = {}));
const defaultVariableHide = () => (VariableHide.DontHide);
exports.defaultVariableHide = defaultVariableHide;
// Options to config when to refresh a variable
// `never`: Never refresh the variable
// `onDashboardLoad`: Queries the data source every time the dashboard loads.
// `onTimeRangeChanged`: Queries the data source when the dashboard time range changes.
var VariableRefresh;
(function (VariableRefresh) {
    VariableRefresh["Never"] = "never";
    VariableRefresh["OnDashboardLoad"] = "onDashboardLoad";
    VariableRefresh["OnTimeRangeChanged"] = "onTimeRangeChanged";
})(VariableRefresh || (exports.VariableRefresh = VariableRefresh = {}));
const defaultVariableRefresh = () => (VariableRefresh.Never);
exports.defaultVariableRefresh = defaultVariableRefresh;
// Sort variable options
// Accepted values are:
// `disabled`: No sorting
// `alphabeticalAsc`: Alphabetical ASC
// `alphabeticalDesc`: Alphabetical DESC
// `numericalAsc`: Numerical ASC
// `numericalDesc`: Numerical DESC
// `alphabeticalCaseInsensitiveAsc`: Alphabetical Case Insensitive ASC
// `alphabeticalCaseInsensitiveDesc`: Alphabetical Case Insensitive DESC
// `naturalAsc`: Natural ASC
// `naturalDesc`: Natural DESC
// VariableSort enum with default value
var VariableSort;
(function (VariableSort) {
    VariableSort["Disabled"] = "disabled";
    VariableSort["AlphabeticalAsc"] = "alphabeticalAsc";
    VariableSort["AlphabeticalDesc"] = "alphabeticalDesc";
    VariableSort["NumericalAsc"] = "numericalAsc";
    VariableSort["NumericalDesc"] = "numericalDesc";
    VariableSort["AlphabeticalCaseInsensitiveAsc"] = "alphabeticalCaseInsensitiveAsc";
    VariableSort["AlphabeticalCaseInsensitiveDesc"] = "alphabeticalCaseInsensitiveDesc";
    VariableSort["NaturalAsc"] = "naturalAsc";
    VariableSort["NaturalDesc"] = "naturalDesc";
})(VariableSort || (exports.VariableSort = VariableSort = {}));
const defaultVariableSort = () => (VariableSort.Disabled);
exports.defaultVariableSort = defaultVariableSort;
const defaultTextVariableKind = () => ({
    kind: "TextVariable",
    spec: (0, exports.defaultTextVariableSpec)(),
});
exports.defaultTextVariableKind = defaultTextVariableKind;
const defaultTextVariableSpec = () => ({
    name: "",
    current: { text: "", value: "", },
    query: "",
    hide: VariableHide.DontHide,
    skipUrlSync: false,
});
exports.defaultTextVariableSpec = defaultTextVariableSpec;
const defaultConstantVariableKind = () => ({
    kind: "ConstantVariable",
    spec: (0, exports.defaultConstantVariableSpec)(),
});
exports.defaultConstantVariableKind = defaultConstantVariableKind;
const defaultConstantVariableSpec = () => ({
    name: "",
    query: "",
    current: { text: "", value: "", },
    hide: VariableHide.DontHide,
    skipUrlSync: false,
});
exports.defaultConstantVariableSpec = defaultConstantVariableSpec;
const defaultDatasourceVariableKind = () => ({
    kind: "DatasourceVariable",
    spec: (0, exports.defaultDatasourceVariableSpec)(),
});
exports.defaultDatasourceVariableKind = defaultDatasourceVariableKind;
const defaultDatasourceVariableSpec = () => ({
    name: "",
    pluginId: "",
    refresh: VariableRefresh.Never,
    regex: "",
    current: { text: "", value: "", },
    options: [],
    multi: false,
    includeAll: false,
    hide: VariableHide.DontHide,
    skipUrlSync: false,
    allowCustomValue: true,
});
exports.defaultDatasourceVariableSpec = defaultDatasourceVariableSpec;
const defaultIntervalVariableKind = () => ({
    kind: "IntervalVariable",
    spec: (0, exports.defaultIntervalVariableSpec)(),
});
exports.defaultIntervalVariableKind = defaultIntervalVariableKind;
const defaultIntervalVariableSpec = () => ({
    name: "",
    query: "",
    current: { text: "", value: "", },
    options: [],
    auto: false,
    auto_min: "",
    auto_count: 0,
    refresh: VariableRefresh.Never,
    hide: VariableHide.DontHide,
    skipUrlSync: false,
});
exports.defaultIntervalVariableSpec = defaultIntervalVariableSpec;
const defaultCustomVariableKind = () => ({
    kind: "CustomVariable",
    spec: (0, exports.defaultCustomVariableSpec)(),
});
exports.defaultCustomVariableKind = defaultCustomVariableKind;
const defaultCustomVariableSpec = () => ({
    name: "",
    query: "",
    current: (0, exports.defaultVariableOption)(),
    options: [],
    multi: false,
    includeAll: false,
    hide: VariableHide.DontHide,
    skipUrlSync: false,
    allowCustomValue: true,
});
exports.defaultCustomVariableSpec = defaultCustomVariableSpec;
const defaultGroupByVariableKind = () => ({
    kind: "GroupByVariable",
    group: "",
    spec: (0, exports.defaultGroupByVariableSpec)(),
});
exports.defaultGroupByVariableKind = defaultGroupByVariableKind;
const defaultGroupByVariableSpec = () => ({
    name: "",
    current: { text: "", value: "", },
    options: [],
    multi: false,
    hide: VariableHide.DontHide,
    skipUrlSync: false,
});
exports.defaultGroupByVariableSpec = defaultGroupByVariableSpec;
const defaultAdhocVariableKind = () => ({
    kind: "AdhocVariable",
    group: "",
    spec: (0, exports.defaultAdhocVariableSpec)(),
});
exports.defaultAdhocVariableKind = defaultAdhocVariableKind;
const defaultAdhocVariableSpec = () => ({
    name: "",
    baseFilters: [],
    filters: [],
    defaultKeys: [],
    hide: VariableHide.DontHide,
    skipUrlSync: false,
    allowCustomValue: true,
});
exports.defaultAdhocVariableSpec = defaultAdhocVariableSpec;
const defaultAdHocFilterWithLabels = () => ({
    key: "",
    operator: "",
    value: "",
    origin: exports.FilterOrigin,
});
exports.defaultAdHocFilterWithLabels = defaultAdHocFilterWithLabels;
// Determine the origin of the adhoc variable filter
exports.FilterOrigin = "dashboard";
const defaultMetricFindValue = () => ({
    text: "",
});
exports.defaultMetricFindValue = defaultMetricFindValue;
const defaultSwitchVariableKind = () => ({
    kind: "SwitchVariable",
    spec: (0, exports.defaultSwitchVariableSpec)(),
});
exports.defaultSwitchVariableKind = defaultSwitchVariableKind;
const defaultSwitchVariableSpec = () => ({
    name: "",
    current: "false",
    enabledValue: "true",
    disabledValue: "false",
    hide: VariableHide.DontHide,
    skipUrlSync: false,
});
exports.defaultSwitchVariableSpec = defaultSwitchVariableSpec;
const defaultKind = () => ({
    kind: "",
    spec: {},
});
exports.defaultKind = defaultKind;
const defaultVariableValue = () => ((0, exports.defaultVariableValueSingle)());
exports.defaultVariableValue = defaultVariableValue;
const defaultVariableValueSingle = () => ("");
exports.defaultVariableValueSingle = defaultVariableValueSingle;
const defaultCustomVariableValue = () => ({
    formatter: "",
});
exports.defaultCustomVariableValue = defaultCustomVariableValue;
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
    VariableType["Switch"] = "switch";
})(VariableType || (exports.VariableType = VariableType = {}));
const defaultVariableType = () => (VariableType.Query);
exports.defaultVariableType = defaultVariableType;
const defaultCustomFormatterVariable = () => ({
    name: "",
    type: VariableType.Query,
    multi: false,
    includeAll: false,
});
exports.defaultCustomFormatterVariable = defaultCustomFormatterVariable;
const defaultVariableValueOption = () => ({
    label: "",
    value: (0, exports.defaultVariableValueSingle)(),
});
exports.defaultVariableValueOption = defaultVariableValueOption;
//# sourceMappingURL=types.gen.js.map