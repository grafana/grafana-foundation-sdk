export declare enum DataTopic {
    Series = "series",
    Annotations = "annotations",
    AlertStates = "alertStates"
}
export declare const defaultDataTopic: () => DataTopic;
export interface DataSourceJsonData {
    authType?: string;
    defaultRegion?: string;
    profile?: string;
    manageAlerts?: boolean;
    alertmanagerUid?: string;
}
export declare const defaultDataSourceJsonData: () => DataSourceJsonData;
export interface DataQuery {
    refId: string;
    hide?: boolean;
    queryType?: string;
    datasource?: any;
}
export declare const defaultDataQuery: () => DataQuery;
export interface BaseDimensionConfig {
    field?: string;
}
export declare const defaultBaseDimensionConfig: () => BaseDimensionConfig;
export declare enum ScaleDimensionMode {
    Linear = "linear",
    Quad = "quad"
}
export declare const defaultScaleDimensionMode: () => ScaleDimensionMode;
export interface ScaleDimensionConfig {
    min: number;
    max: number;
    fixed?: number;
    field?: string;
    mode?: ScaleDimensionMode;
}
export declare const defaultScaleDimensionConfig: () => ScaleDimensionConfig;
export interface ColorDimensionConfig {
    fixed?: string;
    field?: string;
}
export declare const defaultColorDimensionConfig: () => ColorDimensionConfig;
export declare enum ScalarDimensionMode {
    Mod = "mod",
    Clamped = "clamped"
}
export declare const defaultScalarDimensionMode: () => ScalarDimensionMode;
export interface ScalarDimensionConfig {
    min: number;
    max: number;
    fixed?: number;
    field?: string;
    mode?: ScalarDimensionMode;
}
export declare const defaultScalarDimensionConfig: () => ScalarDimensionConfig;
export declare enum TextDimensionMode {
    Fixed = "fixed",
    Field = "field",
    Template = "template"
}
export declare const defaultTextDimensionMode: () => TextDimensionMode;
export interface TextDimensionConfig {
    mode: TextDimensionMode;
    field?: string;
    fixed?: string;
}
export declare const defaultTextDimensionConfig: () => TextDimensionConfig;
export declare enum ResourceDimensionMode {
    Fixed = "fixed",
    Field = "field",
    Mapping = "mapping"
}
export declare const defaultResourceDimensionMode: () => ResourceDimensionMode;
export interface MapLayerOptions {
    type: string;
    name: string;
    config?: any;
    location?: FrameGeometrySource;
    filterData?: any;
    opacity?: number;
    tooltip?: boolean;
}
export declare const defaultMapLayerOptions: () => MapLayerOptions;
export interface FrameGeometrySource {
    mode: FrameGeometrySourceMode;
    geohash?: string;
    latitude?: string;
    longitude?: string;
    wkt?: string;
    lookup?: string;
    gazetteer?: string;
}
export declare const defaultFrameGeometrySource: () => FrameGeometrySource;
export declare enum FrameGeometrySourceMode {
    Auto = "auto",
    Geohash = "geohash",
    Coords = "coords",
    Lookup = "lookup"
}
export declare const defaultFrameGeometrySourceMode: () => FrameGeometrySourceMode;
export declare enum HeatmapCalculationMode {
    Size = "size",
    Count = "count"
}
export declare const defaultHeatmapCalculationMode: () => HeatmapCalculationMode;
export declare enum HeatmapCellLayout {
    Le = "le",
    Ge = "ge",
    Unknown = "unknown",
    Auto = "auto"
}
export declare const defaultHeatmapCellLayout: () => HeatmapCellLayout;
export interface HeatmapCalculationBucketConfig {
    mode?: HeatmapCalculationMode;
    value?: string;
    scale?: ScaleDistributionConfig;
}
export declare const defaultHeatmapCalculationBucketConfig: () => HeatmapCalculationBucketConfig;
export interface ScaleDistributionConfig {
    type: ScaleDistribution;
    log?: number;
    linearThreshold?: number;
}
export declare const defaultScaleDistributionConfig: () => ScaleDistributionConfig;
export declare enum ScaleDistribution {
    Linear = "linear",
    Log = "log",
    Ordinal = "ordinal",
    Symlog = "symlog"
}
export declare const defaultScaleDistribution: () => ScaleDistribution;
export declare enum LogsSortOrder {
    Descending = "Descending",
    Ascending = "Ascending"
}
export declare const defaultLogsSortOrder: () => LogsSortOrder;
export declare enum AxisPlacement {
    Auto = "auto",
    Top = "top",
    Right = "right",
    Bottom = "bottom",
    Left = "left",
    Hidden = "hidden"
}
export declare const defaultAxisPlacement: () => AxisPlacement;
export declare enum AxisColorMode {
    Text = "text",
    Series = "series"
}
export declare const defaultAxisColorMode: () => AxisColorMode;
export declare enum VisibilityMode {
    Auto = "auto",
    Never = "never",
    Always = "always"
}
export declare const defaultVisibilityMode: () => VisibilityMode;
export declare enum GraphDrawStyle {
    Line = "line",
    Bars = "bars",
    Points = "points"
}
export declare const defaultGraphDrawStyle: () => GraphDrawStyle;
export declare enum GraphTransform {
    Constant = "constant",
    NegativeY = "negative-Y"
}
export declare const defaultGraphTransform: () => GraphTransform;
export declare enum LineInterpolation {
    Linear = "linear",
    Smooth = "smooth",
    StepBefore = "stepBefore",
    StepAfter = "stepAfter"
}
export declare const defaultLineInterpolation: () => LineInterpolation;
export declare enum GraphGradientMode {
    None = "none",
    Opacity = "opacity",
    Hue = "hue",
    Scheme = "scheme"
}
export declare const defaultGraphGradientMode: () => GraphGradientMode;
export declare enum StackingMode {
    None = "none",
    Normal = "normal",
    Percent = "percent"
}
export declare const defaultStackingMode: () => StackingMode;
export declare enum BarAlignment {
    Before = -1,
    Center = 0,
    After = 1
}
export declare const defaultBarAlignment: () => BarAlignment;
export declare enum ScaleOrientation {
    Horizontal = 0,
    Vertical = 1
}
export declare const defaultScaleOrientation: () => ScaleOrientation;
export declare enum ScaleDirection {
    Up = 1,
    Right = 1,
    Down = -1,
    Left = -1
}
export declare const defaultScaleDirection: () => ScaleDirection;
export interface LineStyle {
    fill?: "solid" | "dash" | "dot" | "square";
    dash?: number[];
}
export declare const defaultLineStyle: () => LineStyle;
export interface LineConfig {
    lineColor?: string;
    lineWidth?: number;
    lineInterpolation?: LineInterpolation;
    lineStyle?: LineStyle;
    spanNulls?: boolean | number;
}
export declare const defaultLineConfig: () => LineConfig;
export interface BarConfig {
    barAlignment?: BarAlignment;
    barWidthFactor?: number;
    barMaxWidth?: number;
}
export declare const defaultBarConfig: () => BarConfig;
export interface FillConfig {
    fillColor?: string;
    fillOpacity?: number;
    fillBelowTo?: string;
}
export declare const defaultFillConfig: () => FillConfig;
export interface PointsConfig {
    showPoints?: VisibilityMode;
    pointSize?: number;
    pointColor?: string;
    pointSymbol?: string;
}
export declare const defaultPointsConfig: () => PointsConfig;
export interface AxisConfig {
    axisPlacement?: AxisPlacement;
    axisColorMode?: AxisColorMode;
    axisLabel?: string;
    axisWidth?: number;
    axisSoftMin?: number;
    axisSoftMax?: number;
    axisGridShow?: boolean;
    scaleDistribution?: ScaleDistributionConfig;
    axisCenteredZero?: boolean;
    axisBorderShow?: boolean;
}
export declare const defaultAxisConfig: () => AxisConfig;
export interface HideSeriesConfig {
    tooltip: boolean;
    legend: boolean;
    viz: boolean;
}
export declare const defaultHideSeriesConfig: () => HideSeriesConfig;
export interface StackingConfig {
    mode?: StackingMode;
    group?: string;
}
export declare const defaultStackingConfig: () => StackingConfig;
export interface StackableFieldConfig {
    stacking?: StackingConfig;
}
export declare const defaultStackableFieldConfig: () => StackableFieldConfig;
export interface HideableFieldConfig {
    hideFrom?: HideSeriesConfig;
}
export declare const defaultHideableFieldConfig: () => HideableFieldConfig;
export declare enum GraphThresholdsStyleMode {
    Off = "off",
    Line = "line",
    Dashed = "dashed",
    Area = "area",
    LineAndArea = "line+area",
    DashedAndArea = "dashed+area",
    Series = "series"
}
export declare const defaultGraphThresholdsStyleMode: () => GraphThresholdsStyleMode;
export interface GraphThresholdsStyleConfig {
    mode: GraphThresholdsStyleMode;
}
export declare const defaultGraphThresholdsStyleConfig: () => GraphThresholdsStyleConfig;
export declare enum LegendPlacement {
    Bottom = "bottom",
    Right = "right"
}
export declare const defaultLegendPlacement: () => LegendPlacement;
export declare enum LegendDisplayMode {
    List = "list",
    Table = "table",
    Hidden = "hidden"
}
export declare const defaultLegendDisplayMode: () => LegendDisplayMode;
export interface SingleStatBaseOptions {
    reduceOptions: ReduceDataOptions;
    text?: VizTextDisplayOptions;
    orientation: VizOrientation;
}
export declare const defaultSingleStatBaseOptions: () => SingleStatBaseOptions;
export interface ReduceDataOptions {
    values?: boolean;
    limit?: number;
    calcs: string[];
    fields?: string;
}
export declare const defaultReduceDataOptions: () => ReduceDataOptions;
export interface VizTextDisplayOptions {
    titleSize?: number;
    valueSize?: number;
    percentSize?: number;
}
export declare const defaultVizTextDisplayOptions: () => VizTextDisplayOptions;
export declare enum VizOrientation {
    Auto = "auto",
    Vertical = "vertical",
    Horizontal = "horizontal"
}
export declare const defaultVizOrientation: () => VizOrientation;
export interface OptionsWithTooltip {
    tooltip: VizTooltipOptions;
}
export declare const defaultOptionsWithTooltip: () => OptionsWithTooltip;
export interface VizTooltipOptions {
    mode: TooltipDisplayMode;
    sort: SortOrder;
    maxWidth?: number;
    maxHeight?: number;
    hideZeros?: boolean;
}
export declare const defaultVizTooltipOptions: () => VizTooltipOptions;
export declare enum TooltipDisplayMode {
    Single = "single",
    Multi = "multi",
    None = "none"
}
export declare const defaultTooltipDisplayMode: () => TooltipDisplayMode;
export declare enum SortOrder {
    Ascending = "asc",
    Descending = "desc",
    None = "none"
}
export declare const defaultSortOrder: () => SortOrder;
export interface OptionsWithLegend {
    legend: VizLegendOptions;
}
export declare const defaultOptionsWithLegend: () => OptionsWithLegend;
export interface VizLegendOptions {
    displayMode: LegendDisplayMode;
    placement: LegendPlacement;
    showLegend: boolean;
    asTable?: boolean;
    isVisible?: boolean;
    sortBy?: string;
    sortDesc?: boolean;
    width?: number;
    calcs: string[];
}
export declare const defaultVizLegendOptions: () => VizLegendOptions;
export interface OptionsWithTimezones {
    timezone?: TimeZone[];
}
export declare const defaultOptionsWithTimezones: () => OptionsWithTimezones;
export type TimeZone = string;
export declare const defaultTimeZone: () => TimeZone;
export type TimeZoneUtc = "utc";
export type TimeZoneBrowser = "browser";
export interface OptionsWithTextFormatting {
    text?: VizTextDisplayOptions;
}
export declare const defaultOptionsWithTextFormatting: () => OptionsWithTextFormatting;
export declare enum BigValueColorMode {
    Value = "value",
    Background = "background",
    BackgroundSolid = "background_solid",
    None = "none"
}
export declare const defaultBigValueColorMode: () => BigValueColorMode;
export declare enum BigValueGraphMode {
    None = "none",
    Line = "line",
    Area = "area"
}
export declare const defaultBigValueGraphMode: () => BigValueGraphMode;
export declare enum BigValueJustifyMode {
    Auto = "auto",
    Center = "center"
}
export declare const defaultBigValueJustifyMode: () => BigValueJustifyMode;
export declare enum BigValueTextMode {
    Auto = "auto",
    Value = "value",
    ValueAndName = "value_and_name",
    Name = "name",
    None = "none"
}
export declare const defaultBigValueTextMode: () => BigValueTextMode;
export declare enum PercentChangeColorMode {
    Standard = "standard",
    Inverted = "inverted",
    SameAsValue = "same_as_value"
}
export declare const defaultPercentChangeColorMode: () => PercentChangeColorMode;
export declare enum FieldTextAlignment {
    Auto = "auto",
    Left = "left",
    Right = "right",
    Center = "center"
}
export declare const defaultFieldTextAlignment: () => FieldTextAlignment;
export declare enum TimelineValueAlignment {
    Center = "center",
    Left = "left",
    Right = "right"
}
export declare const defaultTimelineValueAlignment: () => TimelineValueAlignment;
export interface GraphFieldConfig {
    drawStyle?: GraphDrawStyle;
    gradientMode?: GraphGradientMode;
    thresholdsStyle?: GraphThresholdsStyleConfig;
    transform?: GraphTransform;
    lineColor?: string;
    lineWidth?: number;
    lineInterpolation?: LineInterpolation;
    lineStyle?: LineStyle;
    fillColor?: string;
    fillOpacity?: number;
    showPoints?: VisibilityMode;
    pointSize?: number;
    pointColor?: string;
    axisPlacement?: AxisPlacement;
    axisColorMode?: AxisColorMode;
    axisLabel?: string;
    axisWidth?: number;
    axisSoftMin?: number;
    axisSoftMax?: number;
    axisGridShow?: boolean;
    scaleDistribution?: ScaleDistributionConfig;
    axisCenteredZero?: boolean;
    barAlignment?: BarAlignment;
    barWidthFactor?: number;
    stacking?: StackingConfig;
    hideFrom?: HideSeriesConfig;
    insertNulls?: boolean | number;
    spanNulls?: boolean | number;
    fillBelowTo?: string;
    pointSymbol?: string;
    axisBorderShow?: boolean;
    barMaxWidth?: number;
}
export declare const defaultGraphFieldConfig: () => GraphFieldConfig;
export declare enum BarGaugeDisplayMode {
    Basic = "basic",
    Lcd = "lcd",
    Gradient = "gradient"
}
export declare const defaultBarGaugeDisplayMode: () => BarGaugeDisplayMode;
export declare enum BarGaugeValueMode {
    Color = "color",
    Text = "text",
    Hidden = "hidden"
}
export declare const defaultBarGaugeValueMode: () => BarGaugeValueMode;
export declare enum BarGaugeNamePlacement {
    Auto = "auto",
    Top = "top",
    Left = "left",
    Hidden = "hidden"
}
export declare const defaultBarGaugeNamePlacement: () => BarGaugeNamePlacement;
export declare enum BarGaugeSizing {
    Auto = "auto",
    Manual = "manual"
}
export declare const defaultBarGaugeSizing: () => BarGaugeSizing;
export type Labels = Record<string, string>;
export declare const defaultLabels: () => Labels;
export declare enum TableCellDisplayMode {
    Auto = "auto",
    ColorText = "color-text",
    ColorBackground = "color-background",
    ColorBackgroundSolid = "color-background-solid",
    GradientGauge = "gradient-gauge",
    LcdGauge = "lcd-gauge",
    JSONView = "json-view",
    BasicGauge = "basic",
    Image = "image",
    Gauge = "gauge",
    Sparkline = "sparkline",
    DataLinks = "data-links",
    Custom = "custom",
    Actions = "actions"
}
export declare const defaultTableCellDisplayMode: () => TableCellDisplayMode;
export declare enum TableCellBackgroundDisplayMode {
    Basic = "basic",
    Gradient = "gradient"
}
export declare const defaultTableCellBackgroundDisplayMode: () => TableCellBackgroundDisplayMode;
export interface TableSortByFieldState {
    displayName: string;
    desc?: boolean;
}
export declare const defaultTableSortByFieldState: () => TableSortByFieldState;
export interface TableFooterOptions {
    show: boolean;
    reducer: string[];
    fields?: string[];
    enablePagination?: boolean;
    countRows?: boolean;
}
export declare const defaultTableFooterOptions: () => TableFooterOptions;
export interface TableAutoCellOptions {
    type: TableCellDisplayMode.Auto;
    wrapText?: boolean;
}
export declare const defaultTableAutoCellOptions: () => TableAutoCellOptions;
export interface TableColorTextCellOptions {
    type: TableCellDisplayMode.ColorText;
    wrapText?: boolean;
}
export declare const defaultTableColorTextCellOptions: () => TableColorTextCellOptions;
export interface TableJsonViewCellOptions {
    type: TableCellDisplayMode.JSONView;
}
export declare const defaultTableJsonViewCellOptions: () => TableJsonViewCellOptions;
export interface TableImageCellOptions {
    type: TableCellDisplayMode.Image;
    alt?: string;
    title?: string;
}
export declare const defaultTableImageCellOptions: () => TableImageCellOptions;
export interface TableDataLinksCellOptions {
    type: TableCellDisplayMode.DataLinks;
}
export declare const defaultTableDataLinksCellOptions: () => TableDataLinksCellOptions;
export interface TableActionsCellOptions {
    type: TableCellDisplayMode.Actions;
}
export declare const defaultTableActionsCellOptions: () => TableActionsCellOptions;
export interface TableBarGaugeCellOptions {
    type: TableCellDisplayMode.Gauge;
    mode?: BarGaugeDisplayMode;
    valueDisplayMode?: BarGaugeValueMode;
}
export declare const defaultTableBarGaugeCellOptions: () => TableBarGaugeCellOptions;
export interface TableSparklineCellOptions {
    type: TableCellDisplayMode.Sparkline;
    drawStyle?: GraphDrawStyle;
    gradientMode?: GraphGradientMode;
    thresholdsStyle?: GraphThresholdsStyleConfig;
    transform?: GraphTransform;
    lineColor?: string;
    lineWidth?: number;
    lineInterpolation?: LineInterpolation;
    lineStyle?: LineStyle;
    fillColor?: string;
    fillOpacity?: number;
    showPoints?: VisibilityMode;
    pointSize?: number;
    pointColor?: string;
    axisPlacement?: AxisPlacement;
    axisColorMode?: AxisColorMode;
    axisLabel?: string;
    axisWidth?: number;
    axisSoftMin?: number;
    axisSoftMax?: number;
    axisGridShow?: boolean;
    scaleDistribution?: ScaleDistributionConfig;
    axisCenteredZero?: boolean;
    barAlignment?: BarAlignment;
    barWidthFactor?: number;
    stacking?: StackingConfig;
    hideFrom?: HideSeriesConfig;
    hideValue?: boolean;
    insertNulls?: boolean | number;
    spanNulls?: boolean | number;
    fillBelowTo?: string;
    pointSymbol?: string;
    axisBorderShow?: boolean;
    barMaxWidth?: number;
}
export declare const defaultTableSparklineCellOptions: () => TableSparklineCellOptions;
export interface TableColoredBackgroundCellOptions {
    type: TableCellDisplayMode.ColorBackground;
    mode?: TableCellBackgroundDisplayMode;
    applyToRow?: boolean;
    wrapText?: boolean;
}
export declare const defaultTableColoredBackgroundCellOptions: () => TableColoredBackgroundCellOptions;
export declare enum TableCellHeight {
    Sm = "sm",
    Md = "md",
    Lg = "lg",
    Auto = "auto"
}
export declare const defaultTableCellHeight: () => TableCellHeight;
export type TableCellOptions = TableAutoCellOptions | TableSparklineCellOptions | TableBarGaugeCellOptions | TableColoredBackgroundCellOptions | TableColorTextCellOptions | TableImageCellOptions | TableDataLinksCellOptions | TableActionsCellOptions | TableJsonViewCellOptions;
export declare const defaultTableCellOptions: () => TableCellOptions;
export declare enum VariableFormatID {
    Lucene = "lucene",
    Raw = "raw",
    Regex = "regex",
    Pipe = "pipe",
    Distributed = "distributed",
    CSV = "csv",
    HTML = "html",
    JSON = "json",
    PercentEncode = "percentencode",
    UriEncode = "uriencode",
    SingleQuote = "singlequote",
    DoubleQuote = "doublequote",
    SQLString = "sqlstring",
    Date = "date",
    Glob = "glob",
    Text = "text",
    QueryParam = "queryparam"
}
export declare const defaultVariableFormatID: () => VariableFormatID;
export interface DataSourceRef {
    type?: string;
    uid?: string;
}
export declare const defaultDataSourceRef: () => DataSourceRef;
export interface ResourceDimensionConfig {
    mode: ResourceDimensionMode;
    field?: string;
    fixed?: string;
}
export declare const defaultResourceDimensionConfig: () => ResourceDimensionConfig;
export interface HeatmapCalculationOptions {
    xBuckets?: HeatmapCalculationBucketConfig;
    yBuckets?: HeatmapCalculationBucketConfig;
}
export declare const defaultHeatmapCalculationOptions: () => HeatmapCalculationOptions;
export declare enum LogsDedupStrategy {
    None = "none",
    Exact = "exact",
    Numbers = "numbers",
    Signature = "signature"
}
export declare const defaultLogsDedupStrategy: () => LogsDedupStrategy;
export declare enum ComparisonOperation {
    EQ = "eq",
    NEQ = "neq",
    LT = "lt",
    LTE = "lte",
    GT = "gt",
    GTE = "gte"
}
export declare const defaultComparisonOperation: () => ComparisonOperation;
export interface TableFieldOptions {
    width?: number;
    minWidth?: number;
    align: FieldTextAlignment;
    displayMode?: TableCellDisplayMode;
    cellOptions?: TableCellOptions;
    hidden?: boolean;
    inspect: boolean;
    filterable?: boolean;
    hideHeader?: boolean;
}
export declare const defaultTableFieldOptions: () => TableFieldOptions;
