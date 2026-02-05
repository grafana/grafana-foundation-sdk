"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.ScaleOrientation = exports.defaultBarAlignment = exports.BarAlignment = exports.defaultStackingMode = exports.StackingMode = exports.defaultGraphGradientMode = exports.GraphGradientMode = exports.defaultLineInterpolation = exports.LineInterpolation = exports.defaultGraphTransform = exports.GraphTransform = exports.defaultGraphDrawStyle = exports.GraphDrawStyle = exports.defaultVisibilityMode = exports.VisibilityMode = exports.defaultAxisColorMode = exports.AxisColorMode = exports.defaultAxisPlacement = exports.AxisPlacement = exports.defaultLogsSortOrder = exports.LogsSortOrder = exports.defaultScaleDistribution = exports.ScaleDistribution = exports.defaultScaleDistributionConfig = exports.defaultHeatmapCalculationBucketConfig = exports.defaultHeatmapCellLayout = exports.HeatmapCellLayout = exports.defaultHeatmapCalculationMode = exports.HeatmapCalculationMode = exports.defaultFrameGeometrySourceMode = exports.FrameGeometrySourceMode = exports.defaultFrameGeometrySource = exports.defaultMapLayerOptions = exports.defaultResourceDimensionMode = exports.ResourceDimensionMode = exports.defaultTextDimensionConfig = exports.defaultTextDimensionMode = exports.TextDimensionMode = exports.defaultScalarDimensionConfig = exports.defaultScalarDimensionMode = exports.ScalarDimensionMode = exports.defaultColorDimensionConfig = exports.defaultScaleDimensionConfig = exports.defaultScaleDimensionMode = exports.ScaleDimensionMode = exports.defaultBaseDimensionConfig = exports.defaultDataQuery = exports.defaultDataSourceJsonData = exports.defaultDataTopic = exports.DataTopic = void 0;
exports.defaultTimelineValueAlignment = exports.TimelineValueAlignment = exports.defaultFieldTextAlignment = exports.FieldTextAlignment = exports.defaultPercentChangeColorMode = exports.PercentChangeColorMode = exports.defaultBigValueTextMode = exports.BigValueTextMode = exports.defaultBigValueJustifyMode = exports.BigValueJustifyMode = exports.defaultBigValueGraphMode = exports.BigValueGraphMode = exports.defaultBigValueColorMode = exports.BigValueColorMode = exports.defaultOptionsWithTextFormatting = exports.defaultTimeZone = exports.defaultOptionsWithTimezones = exports.defaultVizLegendOptions = exports.defaultOptionsWithLegend = exports.defaultSortOrder = exports.SortOrder = exports.defaultTooltipDisplayMode = exports.TooltipDisplayMode = exports.defaultVizTooltipOptions = exports.defaultOptionsWithTooltip = exports.defaultVizOrientation = exports.VizOrientation = exports.defaultVizTextDisplayOptions = exports.defaultReduceDataOptions = exports.defaultSingleStatBaseOptions = exports.defaultLegendDisplayMode = exports.LegendDisplayMode = exports.defaultLegendPlacement = exports.LegendPlacement = exports.defaultGraphThresholdsStyleConfig = exports.defaultGraphThresholdsStyleMode = exports.GraphThresholdsStyleMode = exports.defaultHideableFieldConfig = exports.defaultStackableFieldConfig = exports.defaultStackingConfig = exports.defaultHideSeriesConfig = exports.defaultAxisConfig = exports.defaultPointsConfig = exports.defaultFillConfig = exports.defaultBarConfig = exports.defaultLineConfig = exports.defaultLineStyle = exports.defaultScaleDirection = exports.ScaleDirection = exports.defaultScaleOrientation = void 0;
exports.defaultTableFieldOptions = exports.defaultComparisonOperation = exports.ComparisonOperation = exports.defaultLogsDedupStrategy = exports.LogsDedupStrategy = exports.defaultHeatmapCalculationOptions = exports.defaultResourceDimensionConfig = exports.defaultDataSourceRef = exports.defaultVariableFormatID = exports.VariableFormatID = exports.defaultTableCellOptions = exports.defaultTableCellHeight = exports.TableCellHeight = exports.defaultTableColoredBackgroundCellOptions = exports.defaultTableSparklineCellOptions = exports.defaultTableBarGaugeCellOptions = exports.defaultTableActionsCellOptions = exports.defaultTableDataLinksCellOptions = exports.defaultTableImageCellOptions = exports.defaultTableJsonViewCellOptions = exports.defaultTableColorTextCellOptions = exports.defaultTableAutoCellOptions = exports.defaultTableFooterOptions = exports.defaultTableSortByFieldState = exports.defaultTableCellBackgroundDisplayMode = exports.TableCellBackgroundDisplayMode = exports.defaultTableCellDisplayMode = exports.TableCellDisplayMode = exports.defaultLabels = exports.defaultBarGaugeSizing = exports.BarGaugeSizing = exports.defaultBarGaugeNamePlacement = exports.BarGaugeNamePlacement = exports.defaultBarGaugeValueMode = exports.BarGaugeValueMode = exports.defaultBarGaugeDisplayMode = exports.BarGaugeDisplayMode = exports.defaultGraphFieldConfig = void 0;
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
const defaultDataSourceJsonData = () => ({});
exports.defaultDataSourceJsonData = defaultDataSourceJsonData;
const defaultDataQuery = () => ({
    refId: "",
});
exports.defaultDataQuery = defaultDataQuery;
const defaultBaseDimensionConfig = () => ({});
exports.defaultBaseDimensionConfig = defaultBaseDimensionConfig;
var ScaleDimensionMode;
(function (ScaleDimensionMode) {
    ScaleDimensionMode["Linear"] = "linear";
    ScaleDimensionMode["Quad"] = "quad";
})(ScaleDimensionMode || (exports.ScaleDimensionMode = ScaleDimensionMode = {}));
const defaultScaleDimensionMode = () => (ScaleDimensionMode.Linear);
exports.defaultScaleDimensionMode = defaultScaleDimensionMode;
const defaultScaleDimensionConfig = () => ({
    min: 0,
    max: 0,
});
exports.defaultScaleDimensionConfig = defaultScaleDimensionConfig;
const defaultColorDimensionConfig = () => ({});
exports.defaultColorDimensionConfig = defaultColorDimensionConfig;
var ScalarDimensionMode;
(function (ScalarDimensionMode) {
    ScalarDimensionMode["Mod"] = "mod";
    ScalarDimensionMode["Clamped"] = "clamped";
})(ScalarDimensionMode || (exports.ScalarDimensionMode = ScalarDimensionMode = {}));
const defaultScalarDimensionMode = () => (ScalarDimensionMode.Mod);
exports.defaultScalarDimensionMode = defaultScalarDimensionMode;
const defaultScalarDimensionConfig = () => ({
    min: 0,
    max: 0,
});
exports.defaultScalarDimensionConfig = defaultScalarDimensionConfig;
var TextDimensionMode;
(function (TextDimensionMode) {
    TextDimensionMode["Fixed"] = "fixed";
    TextDimensionMode["Field"] = "field";
    TextDimensionMode["Template"] = "template";
})(TextDimensionMode || (exports.TextDimensionMode = TextDimensionMode = {}));
const defaultTextDimensionMode = () => (TextDimensionMode.Fixed);
exports.defaultTextDimensionMode = defaultTextDimensionMode;
const defaultTextDimensionConfig = () => ({
    mode: TextDimensionMode.Fixed,
});
exports.defaultTextDimensionConfig = defaultTextDimensionConfig;
var ResourceDimensionMode;
(function (ResourceDimensionMode) {
    ResourceDimensionMode["Fixed"] = "fixed";
    ResourceDimensionMode["Field"] = "field";
    ResourceDimensionMode["Mapping"] = "mapping";
})(ResourceDimensionMode || (exports.ResourceDimensionMode = ResourceDimensionMode = {}));
const defaultResourceDimensionMode = () => (ResourceDimensionMode.Fixed);
exports.defaultResourceDimensionMode = defaultResourceDimensionMode;
const defaultMapLayerOptions = () => ({
    type: "",
    name: "",
});
exports.defaultMapLayerOptions = defaultMapLayerOptions;
const defaultFrameGeometrySource = () => ({
    mode: FrameGeometrySourceMode.Auto,
});
exports.defaultFrameGeometrySource = defaultFrameGeometrySource;
var FrameGeometrySourceMode;
(function (FrameGeometrySourceMode) {
    FrameGeometrySourceMode["Auto"] = "auto";
    FrameGeometrySourceMode["Geohash"] = "geohash";
    FrameGeometrySourceMode["Coords"] = "coords";
    FrameGeometrySourceMode["Lookup"] = "lookup";
})(FrameGeometrySourceMode || (exports.FrameGeometrySourceMode = FrameGeometrySourceMode = {}));
const defaultFrameGeometrySourceMode = () => (FrameGeometrySourceMode.Auto);
exports.defaultFrameGeometrySourceMode = defaultFrameGeometrySourceMode;
var HeatmapCalculationMode;
(function (HeatmapCalculationMode) {
    HeatmapCalculationMode["Size"] = "size";
    HeatmapCalculationMode["Count"] = "count";
})(HeatmapCalculationMode || (exports.HeatmapCalculationMode = HeatmapCalculationMode = {}));
const defaultHeatmapCalculationMode = () => (HeatmapCalculationMode.Size);
exports.defaultHeatmapCalculationMode = defaultHeatmapCalculationMode;
var HeatmapCellLayout;
(function (HeatmapCellLayout) {
    HeatmapCellLayout["Le"] = "le";
    HeatmapCellLayout["Ge"] = "ge";
    HeatmapCellLayout["Unknown"] = "unknown";
    HeatmapCellLayout["Auto"] = "auto";
})(HeatmapCellLayout || (exports.HeatmapCellLayout = HeatmapCellLayout = {}));
const defaultHeatmapCellLayout = () => (HeatmapCellLayout.Le);
exports.defaultHeatmapCellLayout = defaultHeatmapCellLayout;
const defaultHeatmapCalculationBucketConfig = () => ({});
exports.defaultHeatmapCalculationBucketConfig = defaultHeatmapCalculationBucketConfig;
const defaultScaleDistributionConfig = () => ({
    type: ScaleDistribution.Linear,
});
exports.defaultScaleDistributionConfig = defaultScaleDistributionConfig;
// TODO docs
var ScaleDistribution;
(function (ScaleDistribution) {
    ScaleDistribution["Linear"] = "linear";
    ScaleDistribution["Log"] = "log";
    ScaleDistribution["Ordinal"] = "ordinal";
    ScaleDistribution["Symlog"] = "symlog";
})(ScaleDistribution || (exports.ScaleDistribution = ScaleDistribution = {}));
const defaultScaleDistribution = () => (ScaleDistribution.Linear);
exports.defaultScaleDistribution = defaultScaleDistribution;
var LogsSortOrder;
(function (LogsSortOrder) {
    LogsSortOrder["Descending"] = "Descending";
    LogsSortOrder["Ascending"] = "Ascending";
})(LogsSortOrder || (exports.LogsSortOrder = LogsSortOrder = {}));
const defaultLogsSortOrder = () => (LogsSortOrder.Descending);
exports.defaultLogsSortOrder = defaultLogsSortOrder;
// TODO docs
var AxisPlacement;
(function (AxisPlacement) {
    AxisPlacement["Auto"] = "auto";
    AxisPlacement["Top"] = "top";
    AxisPlacement["Right"] = "right";
    AxisPlacement["Bottom"] = "bottom";
    AxisPlacement["Left"] = "left";
    AxisPlacement["Hidden"] = "hidden";
})(AxisPlacement || (exports.AxisPlacement = AxisPlacement = {}));
const defaultAxisPlacement = () => (AxisPlacement.Auto);
exports.defaultAxisPlacement = defaultAxisPlacement;
// TODO docs
var AxisColorMode;
(function (AxisColorMode) {
    AxisColorMode["Text"] = "text";
    AxisColorMode["Series"] = "series";
})(AxisColorMode || (exports.AxisColorMode = AxisColorMode = {}));
const defaultAxisColorMode = () => (AxisColorMode.Text);
exports.defaultAxisColorMode = defaultAxisColorMode;
// TODO docs
var VisibilityMode;
(function (VisibilityMode) {
    VisibilityMode["Auto"] = "auto";
    VisibilityMode["Never"] = "never";
    VisibilityMode["Always"] = "always";
})(VisibilityMode || (exports.VisibilityMode = VisibilityMode = {}));
const defaultVisibilityMode = () => (VisibilityMode.Auto);
exports.defaultVisibilityMode = defaultVisibilityMode;
// TODO docs
var GraphDrawStyle;
(function (GraphDrawStyle) {
    GraphDrawStyle["Line"] = "line";
    GraphDrawStyle["Bars"] = "bars";
    GraphDrawStyle["Points"] = "points";
})(GraphDrawStyle || (exports.GraphDrawStyle = GraphDrawStyle = {}));
const defaultGraphDrawStyle = () => (GraphDrawStyle.Line);
exports.defaultGraphDrawStyle = defaultGraphDrawStyle;
// TODO docs
var GraphTransform;
(function (GraphTransform) {
    GraphTransform["Constant"] = "constant";
    GraphTransform["NegativeY"] = "negative-Y";
})(GraphTransform || (exports.GraphTransform = GraphTransform = {}));
const defaultGraphTransform = () => (GraphTransform.Constant);
exports.defaultGraphTransform = defaultGraphTransform;
// TODO docs
var LineInterpolation;
(function (LineInterpolation) {
    LineInterpolation["Linear"] = "linear";
    LineInterpolation["Smooth"] = "smooth";
    LineInterpolation["StepBefore"] = "stepBefore";
    LineInterpolation["StepAfter"] = "stepAfter";
})(LineInterpolation || (exports.LineInterpolation = LineInterpolation = {}));
const defaultLineInterpolation = () => (LineInterpolation.Linear);
exports.defaultLineInterpolation = defaultLineInterpolation;
// TODO docs
var GraphGradientMode;
(function (GraphGradientMode) {
    GraphGradientMode["None"] = "none";
    GraphGradientMode["Opacity"] = "opacity";
    GraphGradientMode["Hue"] = "hue";
    GraphGradientMode["Scheme"] = "scheme";
})(GraphGradientMode || (exports.GraphGradientMode = GraphGradientMode = {}));
const defaultGraphGradientMode = () => (GraphGradientMode.None);
exports.defaultGraphGradientMode = defaultGraphGradientMode;
// TODO docs
var StackingMode;
(function (StackingMode) {
    StackingMode["None"] = "none";
    StackingMode["Normal"] = "normal";
    StackingMode["Percent"] = "percent";
})(StackingMode || (exports.StackingMode = StackingMode = {}));
const defaultStackingMode = () => (StackingMode.None);
exports.defaultStackingMode = defaultStackingMode;
// TODO docs
var BarAlignment;
(function (BarAlignment) {
    BarAlignment[BarAlignment["Before"] = -1] = "Before";
    BarAlignment[BarAlignment["Center"] = 0] = "Center";
    BarAlignment[BarAlignment["After"] = 1] = "After";
})(BarAlignment || (exports.BarAlignment = BarAlignment = {}));
const defaultBarAlignment = () => (BarAlignment.Before);
exports.defaultBarAlignment = defaultBarAlignment;
// TODO docs
var ScaleOrientation;
(function (ScaleOrientation) {
    ScaleOrientation[ScaleOrientation["Horizontal"] = 0] = "Horizontal";
    ScaleOrientation[ScaleOrientation["Vertical"] = 1] = "Vertical";
})(ScaleOrientation || (exports.ScaleOrientation = ScaleOrientation = {}));
const defaultScaleOrientation = () => (ScaleOrientation.Horizontal);
exports.defaultScaleOrientation = defaultScaleOrientation;
// TODO docs
var ScaleDirection;
(function (ScaleDirection) {
    ScaleDirection[ScaleDirection["Up"] = 1] = "Up";
    ScaleDirection[ScaleDirection["Right"] = 1] = "Right";
    ScaleDirection[ScaleDirection["Down"] = -1] = "Down";
    ScaleDirection[ScaleDirection["Left"] = -1] = "Left";
})(ScaleDirection || (exports.ScaleDirection = ScaleDirection = {}));
const defaultScaleDirection = () => (ScaleDirection.Up);
exports.defaultScaleDirection = defaultScaleDirection;
const defaultLineStyle = () => ({});
exports.defaultLineStyle = defaultLineStyle;
const defaultLineConfig = () => ({});
exports.defaultLineConfig = defaultLineConfig;
const defaultBarConfig = () => ({});
exports.defaultBarConfig = defaultBarConfig;
const defaultFillConfig = () => ({});
exports.defaultFillConfig = defaultFillConfig;
const defaultPointsConfig = () => ({});
exports.defaultPointsConfig = defaultPointsConfig;
const defaultAxisConfig = () => ({});
exports.defaultAxisConfig = defaultAxisConfig;
const defaultHideSeriesConfig = () => ({
    tooltip: false,
    legend: false,
    viz: false,
});
exports.defaultHideSeriesConfig = defaultHideSeriesConfig;
const defaultStackingConfig = () => ({});
exports.defaultStackingConfig = defaultStackingConfig;
const defaultStackableFieldConfig = () => ({});
exports.defaultStackableFieldConfig = defaultStackableFieldConfig;
const defaultHideableFieldConfig = () => ({});
exports.defaultHideableFieldConfig = defaultHideableFieldConfig;
// TODO docs
var GraphThresholdsStyleMode;
(function (GraphThresholdsStyleMode) {
    GraphThresholdsStyleMode["Off"] = "off";
    GraphThresholdsStyleMode["Line"] = "line";
    GraphThresholdsStyleMode["Dashed"] = "dashed";
    GraphThresholdsStyleMode["Area"] = "area";
    GraphThresholdsStyleMode["LineAndArea"] = "line+area";
    GraphThresholdsStyleMode["DashedAndArea"] = "dashed+area";
    GraphThresholdsStyleMode["Series"] = "series";
})(GraphThresholdsStyleMode || (exports.GraphThresholdsStyleMode = GraphThresholdsStyleMode = {}));
const defaultGraphThresholdsStyleMode = () => (GraphThresholdsStyleMode.Off);
exports.defaultGraphThresholdsStyleMode = defaultGraphThresholdsStyleMode;
const defaultGraphThresholdsStyleConfig = () => ({
    mode: GraphThresholdsStyleMode.Off,
});
exports.defaultGraphThresholdsStyleConfig = defaultGraphThresholdsStyleConfig;
// TODO docs
var LegendPlacement;
(function (LegendPlacement) {
    LegendPlacement["Bottom"] = "bottom";
    LegendPlacement["Right"] = "right";
})(LegendPlacement || (exports.LegendPlacement = LegendPlacement = {}));
const defaultLegendPlacement = () => (LegendPlacement.Bottom);
exports.defaultLegendPlacement = defaultLegendPlacement;
// TODO docs
// Note: "hidden" needs to remain as an option for plugins compatibility
var LegendDisplayMode;
(function (LegendDisplayMode) {
    LegendDisplayMode["List"] = "list";
    LegendDisplayMode["Table"] = "table";
    LegendDisplayMode["Hidden"] = "hidden";
})(LegendDisplayMode || (exports.LegendDisplayMode = LegendDisplayMode = {}));
const defaultLegendDisplayMode = () => (LegendDisplayMode.List);
exports.defaultLegendDisplayMode = defaultLegendDisplayMode;
const defaultSingleStatBaseOptions = () => ({
    reduceOptions: (0, exports.defaultReduceDataOptions)(),
    orientation: VizOrientation.Auto,
});
exports.defaultSingleStatBaseOptions = defaultSingleStatBaseOptions;
const defaultReduceDataOptions = () => ({
    calcs: [],
});
exports.defaultReduceDataOptions = defaultReduceDataOptions;
const defaultVizTextDisplayOptions = () => ({});
exports.defaultVizTextDisplayOptions = defaultVizTextDisplayOptions;
// TODO docs
var VizOrientation;
(function (VizOrientation) {
    VizOrientation["Auto"] = "auto";
    VizOrientation["Vertical"] = "vertical";
    VizOrientation["Horizontal"] = "horizontal";
})(VizOrientation || (exports.VizOrientation = VizOrientation = {}));
const defaultVizOrientation = () => (VizOrientation.Auto);
exports.defaultVizOrientation = defaultVizOrientation;
const defaultOptionsWithTooltip = () => ({
    tooltip: (0, exports.defaultVizTooltipOptions)(),
});
exports.defaultOptionsWithTooltip = defaultOptionsWithTooltip;
const defaultVizTooltipOptions = () => ({
    mode: TooltipDisplayMode.Single,
    sort: SortOrder.Ascending,
});
exports.defaultVizTooltipOptions = defaultVizTooltipOptions;
// TODO docs
var TooltipDisplayMode;
(function (TooltipDisplayMode) {
    TooltipDisplayMode["Single"] = "single";
    TooltipDisplayMode["Multi"] = "multi";
    TooltipDisplayMode["None"] = "none";
})(TooltipDisplayMode || (exports.TooltipDisplayMode = TooltipDisplayMode = {}));
const defaultTooltipDisplayMode = () => (TooltipDisplayMode.Single);
exports.defaultTooltipDisplayMode = defaultTooltipDisplayMode;
// TODO docs
var SortOrder;
(function (SortOrder) {
    SortOrder["Ascending"] = "asc";
    SortOrder["Descending"] = "desc";
    SortOrder["None"] = "none";
})(SortOrder || (exports.SortOrder = SortOrder = {}));
const defaultSortOrder = () => (SortOrder.Ascending);
exports.defaultSortOrder = defaultSortOrder;
const defaultOptionsWithLegend = () => ({
    legend: (0, exports.defaultVizLegendOptions)(),
});
exports.defaultOptionsWithLegend = defaultOptionsWithLegend;
const defaultVizLegendOptions = () => ({
    displayMode: LegendDisplayMode.List,
    placement: LegendPlacement.Bottom,
    showLegend: false,
    calcs: [],
});
exports.defaultVizLegendOptions = defaultVizLegendOptions;
const defaultOptionsWithTimezones = () => ({});
exports.defaultOptionsWithTimezones = defaultOptionsWithTimezones;
const defaultTimeZone = () => ("browser");
exports.defaultTimeZone = defaultTimeZone;
const defaultOptionsWithTextFormatting = () => ({});
exports.defaultOptionsWithTextFormatting = defaultOptionsWithTextFormatting;
// TODO docs
var BigValueColorMode;
(function (BigValueColorMode) {
    BigValueColorMode["Value"] = "value";
    BigValueColorMode["Background"] = "background";
    BigValueColorMode["BackgroundSolid"] = "background_solid";
    BigValueColorMode["None"] = "none";
})(BigValueColorMode || (exports.BigValueColorMode = BigValueColorMode = {}));
const defaultBigValueColorMode = () => (BigValueColorMode.Value);
exports.defaultBigValueColorMode = defaultBigValueColorMode;
// TODO docs
var BigValueGraphMode;
(function (BigValueGraphMode) {
    BigValueGraphMode["None"] = "none";
    BigValueGraphMode["Line"] = "line";
    BigValueGraphMode["Area"] = "area";
})(BigValueGraphMode || (exports.BigValueGraphMode = BigValueGraphMode = {}));
const defaultBigValueGraphMode = () => (BigValueGraphMode.None);
exports.defaultBigValueGraphMode = defaultBigValueGraphMode;
// TODO docs
var BigValueJustifyMode;
(function (BigValueJustifyMode) {
    BigValueJustifyMode["Auto"] = "auto";
    BigValueJustifyMode["Center"] = "center";
})(BigValueJustifyMode || (exports.BigValueJustifyMode = BigValueJustifyMode = {}));
const defaultBigValueJustifyMode = () => (BigValueJustifyMode.Auto);
exports.defaultBigValueJustifyMode = defaultBigValueJustifyMode;
// TODO docs
var BigValueTextMode;
(function (BigValueTextMode) {
    BigValueTextMode["Auto"] = "auto";
    BigValueTextMode["Value"] = "value";
    BigValueTextMode["ValueAndName"] = "value_and_name";
    BigValueTextMode["Name"] = "name";
    BigValueTextMode["None"] = "none";
})(BigValueTextMode || (exports.BigValueTextMode = BigValueTextMode = {}));
const defaultBigValueTextMode = () => (BigValueTextMode.Auto);
exports.defaultBigValueTextMode = defaultBigValueTextMode;
// TODO docs
var PercentChangeColorMode;
(function (PercentChangeColorMode) {
    PercentChangeColorMode["Standard"] = "standard";
    PercentChangeColorMode["Inverted"] = "inverted";
    PercentChangeColorMode["SameAsValue"] = "same_as_value";
})(PercentChangeColorMode || (exports.PercentChangeColorMode = PercentChangeColorMode = {}));
const defaultPercentChangeColorMode = () => (PercentChangeColorMode.Standard);
exports.defaultPercentChangeColorMode = defaultPercentChangeColorMode;
// TODO -- should not be table specific!
// TODO docs
var FieldTextAlignment;
(function (FieldTextAlignment) {
    FieldTextAlignment["Auto"] = "auto";
    FieldTextAlignment["Left"] = "left";
    FieldTextAlignment["Right"] = "right";
    FieldTextAlignment["Center"] = "center";
})(FieldTextAlignment || (exports.FieldTextAlignment = FieldTextAlignment = {}));
const defaultFieldTextAlignment = () => (FieldTextAlignment.Auto);
exports.defaultFieldTextAlignment = defaultFieldTextAlignment;
// Controls the value alignment in the TimelineChart component
var TimelineValueAlignment;
(function (TimelineValueAlignment) {
    TimelineValueAlignment["Center"] = "center";
    TimelineValueAlignment["Left"] = "left";
    TimelineValueAlignment["Right"] = "right";
})(TimelineValueAlignment || (exports.TimelineValueAlignment = TimelineValueAlignment = {}));
const defaultTimelineValueAlignment = () => (TimelineValueAlignment.Center);
exports.defaultTimelineValueAlignment = defaultTimelineValueAlignment;
const defaultGraphFieldConfig = () => ({});
exports.defaultGraphFieldConfig = defaultGraphFieldConfig;
// Enum expressing the possible display modes
// for the bar gauge component of Grafana UI
var BarGaugeDisplayMode;
(function (BarGaugeDisplayMode) {
    BarGaugeDisplayMode["Basic"] = "basic";
    BarGaugeDisplayMode["Lcd"] = "lcd";
    BarGaugeDisplayMode["Gradient"] = "gradient";
})(BarGaugeDisplayMode || (exports.BarGaugeDisplayMode = BarGaugeDisplayMode = {}));
const defaultBarGaugeDisplayMode = () => (BarGaugeDisplayMode.Basic);
exports.defaultBarGaugeDisplayMode = defaultBarGaugeDisplayMode;
// Allows for the table cell gauge display type to set the gauge mode.
var BarGaugeValueMode;
(function (BarGaugeValueMode) {
    BarGaugeValueMode["Color"] = "color";
    BarGaugeValueMode["Text"] = "text";
    BarGaugeValueMode["Hidden"] = "hidden";
})(BarGaugeValueMode || (exports.BarGaugeValueMode = BarGaugeValueMode = {}));
const defaultBarGaugeValueMode = () => (BarGaugeValueMode.Color);
exports.defaultBarGaugeValueMode = defaultBarGaugeValueMode;
// Allows for the bar gauge name to be placed explicitly
var BarGaugeNamePlacement;
(function (BarGaugeNamePlacement) {
    BarGaugeNamePlacement["Auto"] = "auto";
    BarGaugeNamePlacement["Top"] = "top";
    BarGaugeNamePlacement["Left"] = "left";
    BarGaugeNamePlacement["Hidden"] = "hidden";
})(BarGaugeNamePlacement || (exports.BarGaugeNamePlacement = BarGaugeNamePlacement = {}));
const defaultBarGaugeNamePlacement = () => (BarGaugeNamePlacement.Auto);
exports.defaultBarGaugeNamePlacement = defaultBarGaugeNamePlacement;
// Allows for the bar gauge size to be set explicitly
var BarGaugeSizing;
(function (BarGaugeSizing) {
    BarGaugeSizing["Auto"] = "auto";
    BarGaugeSizing["Manual"] = "manual";
})(BarGaugeSizing || (exports.BarGaugeSizing = BarGaugeSizing = {}));
const defaultBarGaugeSizing = () => (BarGaugeSizing.Auto);
exports.defaultBarGaugeSizing = defaultBarGaugeSizing;
const defaultLabels = () => ({});
exports.defaultLabels = defaultLabels;
// Internally, this is the "type" of cell that's being displayed
// in the table such as colored text, JSON, gauge, etc.
// The color-background-solid, gradient-gauge, and lcd-gauge
// modes are deprecated in favor of new cell subOptions
var TableCellDisplayMode;
(function (TableCellDisplayMode) {
    TableCellDisplayMode["Auto"] = "auto";
    TableCellDisplayMode["ColorText"] = "color-text";
    TableCellDisplayMode["ColorBackground"] = "color-background";
    TableCellDisplayMode["ColorBackgroundSolid"] = "color-background-solid";
    TableCellDisplayMode["GradientGauge"] = "gradient-gauge";
    TableCellDisplayMode["LcdGauge"] = "lcd-gauge";
    TableCellDisplayMode["JSONView"] = "json-view";
    TableCellDisplayMode["BasicGauge"] = "basic";
    TableCellDisplayMode["Image"] = "image";
    TableCellDisplayMode["Gauge"] = "gauge";
    TableCellDisplayMode["Sparkline"] = "sparkline";
    TableCellDisplayMode["DataLinks"] = "data-links";
    TableCellDisplayMode["Custom"] = "custom";
    TableCellDisplayMode["Actions"] = "actions";
})(TableCellDisplayMode || (exports.TableCellDisplayMode = TableCellDisplayMode = {}));
const defaultTableCellDisplayMode = () => (TableCellDisplayMode.Auto);
exports.defaultTableCellDisplayMode = defaultTableCellDisplayMode;
// Display mode to the "Colored Background" display
// mode for table cells. Either displays a solid color (basic mode)
// or a gradient.
var TableCellBackgroundDisplayMode;
(function (TableCellBackgroundDisplayMode) {
    TableCellBackgroundDisplayMode["Basic"] = "basic";
    TableCellBackgroundDisplayMode["Gradient"] = "gradient";
})(TableCellBackgroundDisplayMode || (exports.TableCellBackgroundDisplayMode = TableCellBackgroundDisplayMode = {}));
const defaultTableCellBackgroundDisplayMode = () => (TableCellBackgroundDisplayMode.Basic);
exports.defaultTableCellBackgroundDisplayMode = defaultTableCellBackgroundDisplayMode;
const defaultTableSortByFieldState = () => ({
    displayName: "",
});
exports.defaultTableSortByFieldState = defaultTableSortByFieldState;
const defaultTableFooterOptions = () => ({
    show: false,
    reducer: [],
});
exports.defaultTableFooterOptions = defaultTableFooterOptions;
const defaultTableAutoCellOptions = () => ({
    type: TableCellDisplayMode.Auto,
});
exports.defaultTableAutoCellOptions = defaultTableAutoCellOptions;
const defaultTableColorTextCellOptions = () => ({
    type: TableCellDisplayMode.ColorText,
});
exports.defaultTableColorTextCellOptions = defaultTableColorTextCellOptions;
const defaultTableJsonViewCellOptions = () => ({
    type: TableCellDisplayMode.JSONView,
});
exports.defaultTableJsonViewCellOptions = defaultTableJsonViewCellOptions;
const defaultTableImageCellOptions = () => ({
    type: TableCellDisplayMode.Image,
});
exports.defaultTableImageCellOptions = defaultTableImageCellOptions;
const defaultTableDataLinksCellOptions = () => ({
    type: TableCellDisplayMode.DataLinks,
});
exports.defaultTableDataLinksCellOptions = defaultTableDataLinksCellOptions;
const defaultTableActionsCellOptions = () => ({
    type: TableCellDisplayMode.Actions,
});
exports.defaultTableActionsCellOptions = defaultTableActionsCellOptions;
const defaultTableBarGaugeCellOptions = () => ({
    type: TableCellDisplayMode.Gauge,
});
exports.defaultTableBarGaugeCellOptions = defaultTableBarGaugeCellOptions;
const defaultTableSparklineCellOptions = () => ({
    type: TableCellDisplayMode.Sparkline,
});
exports.defaultTableSparklineCellOptions = defaultTableSparklineCellOptions;
const defaultTableColoredBackgroundCellOptions = () => ({
    type: TableCellDisplayMode.ColorBackground,
});
exports.defaultTableColoredBackgroundCellOptions = defaultTableColoredBackgroundCellOptions;
// Height of a table cell
var TableCellHeight;
(function (TableCellHeight) {
    TableCellHeight["Sm"] = "sm";
    TableCellHeight["Md"] = "md";
    TableCellHeight["Lg"] = "lg";
    TableCellHeight["Auto"] = "auto";
})(TableCellHeight || (exports.TableCellHeight = TableCellHeight = {}));
const defaultTableCellHeight = () => (TableCellHeight.Sm);
exports.defaultTableCellHeight = defaultTableCellHeight;
const defaultTableCellOptions = () => ((0, exports.defaultTableAutoCellOptions)());
exports.defaultTableCellOptions = defaultTableCellOptions;
// Optional formats for the template variable replace functions
// See also https://grafana.com/docs/grafana/latest/dashboards/variables/variable-syntax/#advanced-variable-format-options
var VariableFormatID;
(function (VariableFormatID) {
    VariableFormatID["Lucene"] = "lucene";
    VariableFormatID["Raw"] = "raw";
    VariableFormatID["Regex"] = "regex";
    VariableFormatID["Pipe"] = "pipe";
    VariableFormatID["Distributed"] = "distributed";
    VariableFormatID["CSV"] = "csv";
    VariableFormatID["HTML"] = "html";
    VariableFormatID["JSON"] = "json";
    VariableFormatID["PercentEncode"] = "percentencode";
    VariableFormatID["UriEncode"] = "uriencode";
    VariableFormatID["SingleQuote"] = "singlequote";
    VariableFormatID["DoubleQuote"] = "doublequote";
    VariableFormatID["SQLString"] = "sqlstring";
    VariableFormatID["Date"] = "date";
    VariableFormatID["Glob"] = "glob";
    VariableFormatID["Text"] = "text";
    VariableFormatID["QueryParam"] = "queryparam";
})(VariableFormatID || (exports.VariableFormatID = VariableFormatID = {}));
const defaultVariableFormatID = () => (VariableFormatID.Lucene);
exports.defaultVariableFormatID = defaultVariableFormatID;
const defaultDataSourceRef = () => ({});
exports.defaultDataSourceRef = defaultDataSourceRef;
const defaultResourceDimensionConfig = () => ({
    mode: ResourceDimensionMode.Fixed,
});
exports.defaultResourceDimensionConfig = defaultResourceDimensionConfig;
const defaultHeatmapCalculationOptions = () => ({});
exports.defaultHeatmapCalculationOptions = defaultHeatmapCalculationOptions;
var LogsDedupStrategy;
(function (LogsDedupStrategy) {
    LogsDedupStrategy["None"] = "none";
    LogsDedupStrategy["Exact"] = "exact";
    LogsDedupStrategy["Numbers"] = "numbers";
    LogsDedupStrategy["Signature"] = "signature";
})(LogsDedupStrategy || (exports.LogsDedupStrategy = LogsDedupStrategy = {}));
const defaultLogsDedupStrategy = () => (LogsDedupStrategy.None);
exports.defaultLogsDedupStrategy = defaultLogsDedupStrategy;
// Compare two values
var ComparisonOperation;
(function (ComparisonOperation) {
    ComparisonOperation["EQ"] = "eq";
    ComparisonOperation["NEQ"] = "neq";
    ComparisonOperation["LT"] = "lt";
    ComparisonOperation["LTE"] = "lte";
    ComparisonOperation["GT"] = "gt";
    ComparisonOperation["GTE"] = "gte";
})(ComparisonOperation || (exports.ComparisonOperation = ComparisonOperation = {}));
const defaultComparisonOperation = () => (ComparisonOperation.EQ);
exports.defaultComparisonOperation = defaultComparisonOperation;
const defaultTableFieldOptions = () => ({
    align: FieldTextAlignment.Auto,
    inspect: false,
});
exports.defaultTableFieldOptions = defaultTableFieldOptions;
//# sourceMappingURL=types.gen.js.map