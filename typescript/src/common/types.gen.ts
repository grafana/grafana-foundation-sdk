// Code generated - EDITING IS FUTILE. DO NOT EDIT.

// A topic is attached to DataFrame metadata in query results.
// This specifies where the data should be used.
export enum DataTopic {
	Series = "series",
	Annotations = "annotations",
	AlertStates = "alertStates",
}

export const defaultDataTopic = (): DataTopic => (DataTopic.Series);

// TODO docs
export interface DataSourceJsonData {
	authType?: string;
	defaultRegion?: string;
	profile?: string;
	manageAlerts?: boolean;
	alertmanagerUid?: string;
}

export const defaultDataSourceJsonData = (): DataSourceJsonData => ({
});

// These are the common properties available to all queries in all datasources.
// Specific implementations will *extend* this interface, adding the required
// properties for the given context.
export interface DataQuery {
	// A unique identifier for the query within the list of targets.
	// In server side expressions, the refId is used as a variable name to identify results.
	// By default, the UI will assign A->Z; however setting meaningful names may be useful.
	refId: string;
	// If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.
	hide?: boolean;
	// Specify the query flavor
	// TODO make this required and give it a default
	queryType?: string;
	// For mixed data sources the selected datasource is on the query level.
	// For non mixed scenarios this is undefined.
	// TODO find a better way to do this ^ that's friendly to schema
	// TODO this shouldn't be unknown but DataSourceRef | null
	datasource?: any;
}

export const defaultDataQuery = (): DataQuery => ({
	refId: "",
});

export interface BaseDimensionConfig {
	// fixed: T -- will be added by each element
	field?: string;
}

export const defaultBaseDimensionConfig = (): BaseDimensionConfig => ({
});

export enum ScaleDimensionMode {
	Linear = "linear",
	Quad = "quad",
}

export const defaultScaleDimensionMode = (): ScaleDimensionMode => (ScaleDimensionMode.Linear);

export interface ScaleDimensionConfig {
	min: number;
	max: number;
	fixed?: number;
	// fixed: T -- will be added by each element
	field?: string;
	// | *"linear"
	mode?: ScaleDimensionMode;
}

export const defaultScaleDimensionConfig = (): ScaleDimensionConfig => ({
	min: 0,
	max: 0,
});

export interface ColorDimensionConfig {
	// color value
	fixed?: string;
	// fixed: T -- will be added by each element
	field?: string;
}

export const defaultColorDimensionConfig = (): ColorDimensionConfig => ({
});

export enum ScalarDimensionMode {
	Mod = "mod",
	Clamped = "clamped",
}

export const defaultScalarDimensionMode = (): ScalarDimensionMode => (ScalarDimensionMode.Mod);

export interface ScalarDimensionConfig {
	min: number;
	max: number;
	fixed?: number;
	// fixed: T -- will be added by each element
	field?: string;
	mode?: ScalarDimensionMode;
}

export const defaultScalarDimensionConfig = (): ScalarDimensionConfig => ({
	min: 0,
	max: 0,
});

export enum TextDimensionMode {
	Fixed = "fixed",
	Field = "field",
	Template = "template",
}

export const defaultTextDimensionMode = (): TextDimensionMode => (TextDimensionMode.Fixed);

export interface TextDimensionConfig {
	mode: TextDimensionMode;
	// fixed: T -- will be added by each element
	field?: string;
	fixed?: string;
}

export const defaultTextDimensionConfig = (): TextDimensionConfig => ({
	mode: TextDimensionMode.Fixed,
});

export enum ResourceDimensionMode {
	Fixed = "fixed",
	Field = "field",
	Mapping = "mapping",
}

export const defaultResourceDimensionMode = (): ResourceDimensionMode => (ResourceDimensionMode.Fixed);

export interface MapLayerOptions {
	type: string;
	// configured unique display name
	name: string;
	// Custom options depending on the type
	config?: any;
	// Common method to define geometry fields
	location?: FrameGeometrySource;
	// Defines a frame MatcherConfig that may filter data for the given layer
	filterData?: any;
	// Common properties:
	// https://openlayers.org/en/latest/apidoc/module-ol_layer_Base-BaseLayer.html
	// Layer opacity (0-1)
	opacity?: number;
	// Check tooltip (defaults to true)
	tooltip?: boolean;
}

export const defaultMapLayerOptions = (): MapLayerOptions => ({
	type: "",
	name: "",
});

export enum FrameGeometrySourceMode {
	Auto = "auto",
	Geohash = "geohash",
	Coords = "coords",
	Lookup = "lookup",
}

export const defaultFrameGeometrySourceMode = (): FrameGeometrySourceMode => (FrameGeometrySourceMode.Auto);

export enum HeatmapCalculationMode {
	Size = "size",
	Count = "count",
}

export const defaultHeatmapCalculationMode = (): HeatmapCalculationMode => (HeatmapCalculationMode.Size);

export enum HeatmapCellLayout {
	Le = "le",
	Ge = "ge",
	Unknown = "unknown",
	Auto = "auto",
}

export const defaultHeatmapCellLayout = (): HeatmapCellLayout => (HeatmapCellLayout.Le);

export interface HeatmapCalculationBucketConfig {
	// Sets the bucket calculation mode
	mode?: HeatmapCalculationMode;
	// The number of buckets to use for the axis in the heatmap
	value?: string;
	// Controls the scale of the buckets
	scale?: ScaleDistributionConfig;
}

export const defaultHeatmapCalculationBucketConfig = (): HeatmapCalculationBucketConfig => ({
});

export enum LogsSortOrder {
	Descending = "Descending",
	Ascending = "Ascending",
}

export const defaultLogsSortOrder = (): LogsSortOrder => (LogsSortOrder.Descending);

// TODO docs
export enum AxisPlacement {
	Auto = "auto",
	Top = "top",
	Right = "right",
	Bottom = "bottom",
	Left = "left",
	Hidden = "hidden",
}

export const defaultAxisPlacement = (): AxisPlacement => (AxisPlacement.Auto);

// TODO docs
export enum AxisColorMode {
	Text = "text",
	Series = "series",
}

export const defaultAxisColorMode = (): AxisColorMode => (AxisColorMode.Text);

// TODO docs
export enum VisibilityMode {
	Auto = "auto",
	Never = "never",
	Always = "always",
}

export const defaultVisibilityMode = (): VisibilityMode => (VisibilityMode.Auto);

// TODO docs
export enum GraphDrawStyle {
	Line = "line",
	Bars = "bars",
	Points = "points",
}

export const defaultGraphDrawStyle = (): GraphDrawStyle => (GraphDrawStyle.Line);

// TODO docs
export enum GraphTransform {
	Constant = "constant",
	NegativeY = "negative-Y",
}

export const defaultGraphTransform = (): GraphTransform => (GraphTransform.Constant);

// TODO docs
export enum LineInterpolation {
	Linear = "linear",
	Smooth = "smooth",
	StepBefore = "stepBefore",
	StepAfter = "stepAfter",
}

export const defaultLineInterpolation = (): LineInterpolation => (LineInterpolation.Linear);

// TODO docs
export enum ScaleDistribution {
	Linear = "linear",
	Log = "log",
	Ordinal = "ordinal",
	Symlog = "symlog",
}

export const defaultScaleDistribution = (): ScaleDistribution => (ScaleDistribution.Linear);

// TODO docs
export enum GraphGradientMode {
	None = "none",
	Opacity = "opacity",
	Hue = "hue",
	Scheme = "scheme",
}

export const defaultGraphGradientMode = (): GraphGradientMode => (GraphGradientMode.None);

// TODO docs
export enum StackingMode {
	None = "none",
	Normal = "normal",
	Percent = "percent",
}

export const defaultStackingMode = (): StackingMode => (StackingMode.None);

// TODO docs
export enum BarAlignment {
	Before = -1,
	Center = 0,
	After = 1,
}

export const defaultBarAlignment = (): BarAlignment => (BarAlignment.Before);

// TODO docs
export enum ScaleOrientation {
	Horizontal = 0,
	Vertical = 1,
}

export const defaultScaleOrientation = (): ScaleOrientation => (ScaleOrientation.Horizontal);

// TODO docs
export enum ScaleDirection {
	Up = 1,
	Right = 1,
	Down = -1,
	Left = -1,
}

export const defaultScaleDirection = (): ScaleDirection => (ScaleDirection.Up);

// TODO docs
export interface LineStyle {
	fill?: "solid" | "dash" | "dot" | "square";
	dash?: number[];
}

export const defaultLineStyle = (): LineStyle => ({
});

// TODO docs
export interface LineConfig {
	lineColor?: string;
	lineWidth?: number;
	lineInterpolation?: LineInterpolation;
	lineStyle?: LineStyle;
	// Indicate if null values should be treated as gaps or connected.
	// When the value is a number, it represents the maximum delta in the
	// X axis that should be considered connected.  For timeseries, this is milliseconds
	spanNulls?: boolean | number;
}

export const defaultLineConfig = (): LineConfig => ({
});

// TODO docs
export interface BarConfig {
	barAlignment?: BarAlignment;
	barWidthFactor?: number;
	barMaxWidth?: number;
}

export const defaultBarConfig = (): BarConfig => ({
});

// TODO docs
export interface FillConfig {
	fillColor?: string;
	fillOpacity?: number;
	fillBelowTo?: string;
}

export const defaultFillConfig = (): FillConfig => ({
});

// TODO docs
export interface PointsConfig {
	showPoints?: VisibilityMode;
	pointSize?: number;
	pointColor?: string;
	pointSymbol?: string;
}

export const defaultPointsConfig = (): PointsConfig => ({
});

// TODO docs
export interface ScaleDistributionConfig {
	type: ScaleDistribution;
	log?: number;
	linearThreshold?: number;
}

export const defaultScaleDistributionConfig = (): ScaleDistributionConfig => ({
	type: ScaleDistribution.Linear,
});

// TODO docs
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

export const defaultAxisConfig = (): AxisConfig => ({
});

// TODO docs
export interface HideSeriesConfig {
	tooltip: boolean;
	legend: boolean;
	viz: boolean;
}

export const defaultHideSeriesConfig = (): HideSeriesConfig => ({
	tooltip: false,
	legend: false,
	viz: false,
});

// TODO docs
export interface StackingConfig {
	mode?: StackingMode;
	group?: string;
}

export const defaultStackingConfig = (): StackingConfig => ({
});

// TODO docs
export interface StackableFieldConfig {
	stacking?: StackingConfig;
}

export const defaultStackableFieldConfig = (): StackableFieldConfig => ({
});

// TODO docs
export interface HideableFieldConfig {
	hideFrom?: HideSeriesConfig;
}

export const defaultHideableFieldConfig = (): HideableFieldConfig => ({
});

// TODO docs
export enum GraphThresholdsStyleMode {
	Off = "off",
	Line = "line",
	Dashed = "dashed",
	Area = "area",
	LineAndArea = "line+area",
	DashedAndArea = "dashed+area",
	Series = "series",
}

export const defaultGraphThresholdsStyleMode = (): GraphThresholdsStyleMode => (GraphThresholdsStyleMode.Off);

// TODO docs
export interface GraphThresholdsStyleConfig {
	mode: GraphThresholdsStyleMode;
}

export const defaultGraphThresholdsStyleConfig = (): GraphThresholdsStyleConfig => ({
	mode: GraphThresholdsStyleMode.Off,
});

// TODO docs
export enum LegendPlacement {
	Bottom = "bottom",
	Right = "right",
}

export const defaultLegendPlacement = (): LegendPlacement => (LegendPlacement.Bottom);

// TODO docs
// Note: "hidden" needs to remain as an option for plugins compatibility
export enum LegendDisplayMode {
	List = "list",
	Table = "table",
	Hidden = "hidden",
}

export const defaultLegendDisplayMode = (): LegendDisplayMode => (LegendDisplayMode.List);

// TODO docs
export interface SingleStatBaseOptions {
	reduceOptions: ReduceDataOptions;
	text?: VizTextDisplayOptions;
	orientation: VizOrientation;
}

export const defaultSingleStatBaseOptions = (): SingleStatBaseOptions => ({
	reduceOptions: defaultReduceDataOptions(),
	orientation: VizOrientation.Auto,
});

// TODO docs
export interface ReduceDataOptions {
	// If true show each row value
	values?: boolean;
	// if showing all values limit
	limit?: number;
	// When !values, pick one value for the whole field
	calcs: string[];
	// Which fields to show.  By default this is only numeric fields
	fields?: string;
}

export const defaultReduceDataOptions = (): ReduceDataOptions => ({
	calcs: [],
});

// TODO docs
export enum VizOrientation {
	Auto = "auto",
	Vertical = "vertical",
	Horizontal = "horizontal",
}

export const defaultVizOrientation = (): VizOrientation => (VizOrientation.Auto);

// TODO docs
export interface OptionsWithTooltip {
	tooltip: VizTooltipOptions;
}

export const defaultOptionsWithTooltip = (): OptionsWithTooltip => ({
	tooltip: defaultVizTooltipOptions(),
});

// TODO docs
export interface OptionsWithLegend {
	legend: VizLegendOptions;
}

export const defaultOptionsWithLegend = (): OptionsWithLegend => ({
	legend: defaultVizLegendOptions(),
});

// TODO docs
export interface OptionsWithTimezones {
	timezone?: TimeZone[];
}

export const defaultOptionsWithTimezones = (): OptionsWithTimezones => ({
});

// TODO docs
export interface OptionsWithTextFormatting {
	text?: VizTextDisplayOptions;
}

export const defaultOptionsWithTextFormatting = (): OptionsWithTextFormatting => ({
});

// TODO docs
export enum BigValueColorMode {
	Value = "value",
	Background = "background",
	BackgroundSolid = "background_solid",
	None = "none",
}

export const defaultBigValueColorMode = (): BigValueColorMode => (BigValueColorMode.Value);

// TODO docs
export enum BigValueGraphMode {
	None = "none",
	Line = "line",
	Area = "area",
}

export const defaultBigValueGraphMode = (): BigValueGraphMode => (BigValueGraphMode.None);

// TODO docs
export enum BigValueJustifyMode {
	Auto = "auto",
	Center = "center",
}

export const defaultBigValueJustifyMode = (): BigValueJustifyMode => (BigValueJustifyMode.Auto);

// TODO docs
export enum BigValueTextMode {
	Auto = "auto",
	Value = "value",
	ValueAndName = "value_and_name",
	Name = "name",
	None = "none",
}

export const defaultBigValueTextMode = (): BigValueTextMode => (BigValueTextMode.Auto);

// TODO docs
export enum PercentChangeColorMode {
	Standard = "standard",
	Inverted = "inverted",
	SameAsValue = "same_as_value",
}

export const defaultPercentChangeColorMode = (): PercentChangeColorMode => (PercentChangeColorMode.Standard);

// TODO -- should not be table specific!
// TODO docs
export enum FieldTextAlignment {
	Auto = "auto",
	Left = "left",
	Right = "right",
	Center = "center",
}

export const defaultFieldTextAlignment = (): FieldTextAlignment => (FieldTextAlignment.Auto);

// Controls the value alignment in the TimelineChart component
export enum TimelineValueAlignment {
	Center = "center",
	Left = "left",
	Right = "right",
}

export const defaultTimelineValueAlignment = (): TimelineValueAlignment => (TimelineValueAlignment.Center);

// TODO docs
export interface VizTextDisplayOptions {
	// Explicit title text size
	titleSize?: number;
	// Explicit value text size
	valueSize?: number;
}

export const defaultVizTextDisplayOptions = (): VizTextDisplayOptions => ({
});

// TODO docs
export enum TooltipDisplayMode {
	Single = "single",
	Multi = "multi",
	None = "none",
}

export const defaultTooltipDisplayMode = (): TooltipDisplayMode => (TooltipDisplayMode.Single);

// TODO docs
export enum SortOrder {
	Ascending = "asc",
	Descending = "desc",
	None = "none",
}

export const defaultSortOrder = (): SortOrder => (SortOrder.Ascending);

// TODO docs
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
	// Indicate if null values should be treated as gaps or connected.
	// When the value is a number, it represents the maximum delta in the
	// X axis that should be considered connected.  For timeseries, this is milliseconds
	spanNulls?: boolean | number;
	fillBelowTo?: string;
	pointSymbol?: string;
	axisBorderShow?: boolean;
	barMaxWidth?: number;
}

export const defaultGraphFieldConfig = (): GraphFieldConfig => ({
});

// TODO docs
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

export const defaultVizLegendOptions = (): VizLegendOptions => ({
	displayMode: LegendDisplayMode.List,
	placement: LegendPlacement.Bottom,
	showLegend: false,
	calcs: [
],
});

// Enum expressing the possible display modes
// for the bar gauge component of Grafana UI
export enum BarGaugeDisplayMode {
	Basic = "basic",
	Lcd = "lcd",
	Gradient = "gradient",
}

export const defaultBarGaugeDisplayMode = (): BarGaugeDisplayMode => (BarGaugeDisplayMode.Basic);

// Allows for the table cell gauge display type to set the gauge mode.
export enum BarGaugeValueMode {
	Color = "color",
	Text = "text",
	Hidden = "hidden",
}

export const defaultBarGaugeValueMode = (): BarGaugeValueMode => (BarGaugeValueMode.Color);

// Allows for the bar gauge name to be placed explicitly
export enum BarGaugeNamePlacement {
	Auto = "auto",
	Top = "top",
	Left = "left",
}

export const defaultBarGaugeNamePlacement = (): BarGaugeNamePlacement => (BarGaugeNamePlacement.Auto);

// Allows for the bar gauge size to be set explicitly
export enum BarGaugeSizing {
	Auto = "auto",
	Manual = "manual",
}

export const defaultBarGaugeSizing = (): BarGaugeSizing => (BarGaugeSizing.Auto);

// TODO docs
export interface VizTooltipOptions {
	mode: TooltipDisplayMode;
	sort: SortOrder;
	maxWidth?: number;
	maxHeight?: number;
}

export const defaultVizTooltipOptions = (): VizTooltipOptions => ({
	mode: TooltipDisplayMode.Single,
	sort: SortOrder.Ascending,
});

export type Labels = Record<string, string>;

export const defaultLabels = (): Labels => ({});

// Internally, this is the "type" of cell that's being displayed
// in the table such as colored text, JSON, gauge, etc.
// The color-background-solid, gradient-gauge, and lcd-gauge
// modes are deprecated in favor of new cell subOptions
export enum TableCellDisplayMode {
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
}

export const defaultTableCellDisplayMode = (): TableCellDisplayMode => (TableCellDisplayMode.Auto);

// Display mode to the "Colored Background" display
// mode for table cells. Either displays a solid color (basic mode)
// or a gradient.
export enum TableCellBackgroundDisplayMode {
	Basic = "basic",
	Gradient = "gradient",
}

export const defaultTableCellBackgroundDisplayMode = (): TableCellBackgroundDisplayMode => (TableCellBackgroundDisplayMode.Basic);

// Sort by field state
export interface TableSortByFieldState {
	// Sets the display name of the field to sort by
	displayName: string;
	// Flag used to indicate descending sort order
	desc?: boolean;
}

export const defaultTableSortByFieldState = (): TableSortByFieldState => ({
	displayName: "",
});

// Footer options
export interface TableFooterOptions {
	show: boolean;
	// actually 1 value
	reducer: string[];
	fields?: string[];
	enablePagination?: boolean;
	countRows?: boolean;
}

export const defaultTableFooterOptions = (): TableFooterOptions => ({
	show: false,
	reducer: [],
});

// Auto mode table cell options
export interface TableAutoCellOptions {
	type: "auto";
	wrapText?: boolean;
}

export const defaultTableAutoCellOptions = (): TableAutoCellOptions => ({
	type: "auto",
});

// Colored text cell options
export interface TableColorTextCellOptions {
	type: "color-text";
	wrapText?: boolean;
}

export const defaultTableColorTextCellOptions = (): TableColorTextCellOptions => ({
	type: "color-text",
});

// Json view cell options
export interface TableJsonViewCellOptions {
	type: "json-view";
}

export const defaultTableJsonViewCellOptions = (): TableJsonViewCellOptions => ({
	type: "json-view",
});

// Json view cell options
export interface TableImageCellOptions {
	type: "image";
}

export const defaultTableImageCellOptions = (): TableImageCellOptions => ({
	type: "image",
});

// Show data links in the cell
export interface TableDataLinksCellOptions {
	type: "data-links";
}

export const defaultTableDataLinksCellOptions = (): TableDataLinksCellOptions => ({
	type: "data-links",
});

// Gauge cell options
export interface TableBarGaugeCellOptions {
	type: "gauge";
	mode?: BarGaugeDisplayMode;
	valueDisplayMode?: BarGaugeValueMode;
}

export const defaultTableBarGaugeCellOptions = (): TableBarGaugeCellOptions => ({
	type: "gauge",
});

// Sparkline cell options
export interface TableSparklineCellOptions {
	type: "sparkline";
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
	// Indicate if null values should be treated as gaps or connected.
	// When the value is a number, it represents the maximum delta in the
	// X axis that should be considered connected.  For timeseries, this is milliseconds
	spanNulls?: boolean | number;
	fillBelowTo?: string;
	pointSymbol?: string;
	axisBorderShow?: boolean;
	barMaxWidth?: number;
}

export const defaultTableSparklineCellOptions = (): TableSparklineCellOptions => ({
	type: "sparkline",
});

// Colored background cell options
export interface TableColoredBackgroundCellOptions {
	type: "color-background";
	mode?: TableCellBackgroundDisplayMode;
	applyToRow?: boolean;
	wrapText?: boolean;
}

export const defaultTableColoredBackgroundCellOptions = (): TableColoredBackgroundCellOptions => ({
	type: "color-background",
});

// Height of a table cell
export enum TableCellHeight {
	Sm = "sm",
	Md = "md",
	Lg = "lg",
	Auto = "auto",
}

export const defaultTableCellHeight = (): TableCellHeight => (TableCellHeight.Sm);

// Table cell options. Each cell has a display mode
// and other potential options for that display.
export type TableCellOptions = TableAutoCellOptions | TableSparklineCellOptions | TableBarGaugeCellOptions | TableColoredBackgroundCellOptions | TableColorTextCellOptions | TableImageCellOptions | TableDataLinksCellOptions | TableJsonViewCellOptions;

export const defaultTableCellOptions = (): TableCellOptions => (defaultTableAutoCellOptions());

// Use UTC/GMT timezone
export type TimeZoneUtc = "utc";

// Use the timezone defined by end user web browser
export type TimeZoneBrowser = "browser";

// Optional formats for the template variable replace functions
// See also https://grafana.com/docs/grafana/latest/dashboards/variables/variable-syntax/#advanced-variable-format-options
export enum VariableFormatID {
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
	QueryParam = "queryparam",
}

export const defaultVariableFormatID = (): VariableFormatID => (VariableFormatID.Lucene);

// Links to a resource (image/svg path)
export interface ResourceDimensionConfig {
	mode: ResourceDimensionMode;
	// fixed: T -- will be added by each element
	field?: string;
	fixed?: string;
}

export const defaultResourceDimensionConfig = (): ResourceDimensionConfig => ({
	mode: ResourceDimensionMode.Fixed,
});

export interface FrameGeometrySource {
	mode: FrameGeometrySourceMode;
	// Field mappings
	geohash?: string;
	latitude?: string;
	longitude?: string;
	wkt?: string;
	lookup?: string;
	// Path to Gazetteer
	gazetteer?: string;
}

export const defaultFrameGeometrySource = (): FrameGeometrySource => ({
	mode: FrameGeometrySourceMode.Auto,
});

export interface HeatmapCalculationOptions {
	// The number of buckets to use for the xAxis in the heatmap
	xBuckets?: HeatmapCalculationBucketConfig;
	// The number of buckets to use for the yAxis in the heatmap
	yBuckets?: HeatmapCalculationBucketConfig;
}

export const defaultHeatmapCalculationOptions = (): HeatmapCalculationOptions => ({
});

export enum LogsDedupStrategy {
	None = "none",
	Exact = "exact",
	Numbers = "numbers",
	Signature = "signature",
}

export const defaultLogsDedupStrategy = (): LogsDedupStrategy => (LogsDedupStrategy.None);

// Compare two values
export enum ComparisonOperation {
	EQ = "eq",
	NEQ = "neq",
	LT = "lt",
	LTE = "lte",
	GT = "gt",
	GTE = "gte",
}

export const defaultComparisonOperation = (): ComparisonOperation => (ComparisonOperation.EQ);

// Field options for each field within a table (e.g 10, "The String", 64.20, etc.)
// Generally defines alignment, filtering capabilties, display options, etc.
export interface TableFieldOptions {
	width?: number;
	minWidth?: number;
	align: FieldTextAlignment;
	// This field is deprecated in favor of using cellOptions
	displayMode?: TableCellDisplayMode;
	cellOptions?: TableCellOptions;
	// ?? default is missing or false ??
	hidden?: boolean;
	inspect: boolean;
	filterable?: boolean;
	// Hides any header for a column, useful for columns that show some static content or buttons.
	hideHeader?: boolean;
}

export const defaultTableFieldOptions = (): TableFieldOptions => ({
	align: FieldTextAlignment.Auto,
	inspect: false,
});

// A specific timezone from https://en.wikipedia.org/wiki/Tz_database
export type TimeZone = string;

export const defaultTimeZone = (): TimeZone => ("browser");

