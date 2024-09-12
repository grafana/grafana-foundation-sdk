// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package common

import (
	"encoding/json"
	"errors"
	"fmt"
)

// A topic is attached to DataFrame metadata in query results.
// This specifies where the data should be used.
type DataTopic string

const (
	DataTopicSeries      DataTopic = "series"
	DataTopicAnnotations DataTopic = "annotations"
	DataTopicAlertStates DataTopic = "alertStates"
)

// TODO docs
type DataSourceJsonData struct {
	AuthType        *string `json:"authType,omitempty"`
	DefaultRegion   *string `json:"defaultRegion,omitempty"`
	Profile         *string `json:"profile,omitempty"`
	ManageAlerts    *bool   `json:"manageAlerts,omitempty"`
	AlertmanagerUid *string `json:"alertmanagerUid,omitempty"`
}

// These are the common properties available to all queries in all datasources.
// Specific implementations will *extend* this interface, adding the required
// properties for the given context.
type DataQuery struct {
	// A unique identifier for the query within the list of targets.
	// In server side expressions, the refId is used as a variable name to identify results.
	// By default, the UI will assign A->Z; however setting meaningful names may be useful.
	RefId string `json:"refId"`
	// If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.
	Hide *bool `json:"hide,omitempty"`
	// Specify the query flavor
	// TODO make this required and give it a default
	QueryType *string `json:"queryType,omitempty"`
	// For mixed data sources the selected datasource is on the query level.
	// For non mixed scenarios this is undefined.
	// TODO find a better way to do this ^ that's friendly to schema
	// TODO this shouldn't be unknown but DataSourceRef | null
	Datasource any `json:"datasource,omitempty"`
}

type BaseDimensionConfig struct {
	// fixed: T -- will be added by each element
	Field *string `json:"field,omitempty"`
}

type ScaleDimensionMode string

const (
	ScaleDimensionModeLinear ScaleDimensionMode = "linear"
	ScaleDimensionModeQuad   ScaleDimensionMode = "quad"
)

type ScaleDimensionConfig struct {
	Min   float64  `json:"min"`
	Max   float64  `json:"max"`
	Fixed *float64 `json:"fixed,omitempty"`
	// fixed: T -- will be added by each element
	Field *string `json:"field,omitempty"`
	// | *"linear"
	Mode *ScaleDimensionMode `json:"mode,omitempty"`
}

type ColorDimensionConfig struct {
	// color value
	Fixed *string `json:"fixed,omitempty"`
	// fixed: T -- will be added by each element
	Field *string `json:"field,omitempty"`
}

type ScalarDimensionMode string

const (
	ScalarDimensionModeMod     ScalarDimensionMode = "mod"
	ScalarDimensionModeClamped ScalarDimensionMode = "clamped"
)

type ScalarDimensionConfig struct {
	Min   float64  `json:"min"`
	Max   float64  `json:"max"`
	Fixed *float64 `json:"fixed,omitempty"`
	// fixed: T -- will be added by each element
	Field *string              `json:"field,omitempty"`
	Mode  *ScalarDimensionMode `json:"mode,omitempty"`
}

type TextDimensionMode string

const (
	TextDimensionModeFixed    TextDimensionMode = "fixed"
	TextDimensionModeField    TextDimensionMode = "field"
	TextDimensionModeTemplate TextDimensionMode = "template"
)

type TextDimensionConfig struct {
	Mode TextDimensionMode `json:"mode"`
	// fixed: T -- will be added by each element
	Field *string `json:"field,omitempty"`
	Fixed *string `json:"fixed,omitempty"`
}

type ResourceDimensionMode string

const (
	ResourceDimensionModeFixed   ResourceDimensionMode = "fixed"
	ResourceDimensionModeField   ResourceDimensionMode = "field"
	ResourceDimensionModeMapping ResourceDimensionMode = "mapping"
)

type MapLayerOptions struct {
	Type string `json:"type"`
	// configured unique display name
	Name string `json:"name"`
	// Custom options depending on the type
	Config any `json:"config,omitempty"`
	// Common method to define geometry fields
	Location *FrameGeometrySource `json:"location,omitempty"`
	// Defines a frame MatcherConfig that may filter data for the given layer
	FilterData any `json:"filterData,omitempty"`
	// Common properties:
	// https://openlayers.org/en/latest/apidoc/module-ol_layer_Base-BaseLayer.html
	// Layer opacity (0-1)
	Opacity *int64 `json:"opacity,omitempty"`
	// Check tooltip (defaults to true)
	Tooltip *bool `json:"tooltip,omitempty"`
}

type FrameGeometrySourceMode string

const (
	FrameGeometrySourceModeAuto    FrameGeometrySourceMode = "auto"
	FrameGeometrySourceModeGeohash FrameGeometrySourceMode = "geohash"
	FrameGeometrySourceModeCoords  FrameGeometrySourceMode = "coords"
	FrameGeometrySourceModeLookup  FrameGeometrySourceMode = "lookup"
)

type HeatmapCalculationMode string

const (
	HeatmapCalculationModeSize  HeatmapCalculationMode = "size"
	HeatmapCalculationModeCount HeatmapCalculationMode = "count"
)

type HeatmapCellLayout string

const (
	HeatmapCellLayoutLe      HeatmapCellLayout = "le"
	HeatmapCellLayoutGe      HeatmapCellLayout = "ge"
	HeatmapCellLayoutUnknown HeatmapCellLayout = "unknown"
	HeatmapCellLayoutAuto    HeatmapCellLayout = "auto"
)

type HeatmapCalculationBucketConfig struct {
	// Sets the bucket calculation mode
	Mode *HeatmapCalculationMode `json:"mode,omitempty"`
	// The number of buckets to use for the axis in the heatmap
	Value *string `json:"value,omitempty"`
	// Controls the scale of the buckets
	Scale *ScaleDistributionConfig `json:"scale,omitempty"`
}

type LogsSortOrder string

const (
	LogsSortOrderDescending LogsSortOrder = "Descending"
	LogsSortOrderAscending  LogsSortOrder = "Ascending"
)

// TODO docs
type AxisPlacement string

const (
	AxisPlacementAuto   AxisPlacement = "auto"
	AxisPlacementTop    AxisPlacement = "top"
	AxisPlacementRight  AxisPlacement = "right"
	AxisPlacementBottom AxisPlacement = "bottom"
	AxisPlacementLeft   AxisPlacement = "left"
	AxisPlacementHidden AxisPlacement = "hidden"
)

// TODO docs
type AxisColorMode string

const (
	AxisColorModeText   AxisColorMode = "text"
	AxisColorModeSeries AxisColorMode = "series"
)

// TODO docs
type VisibilityMode string

const (
	VisibilityModeAuto   VisibilityMode = "auto"
	VisibilityModeNever  VisibilityMode = "never"
	VisibilityModeAlways VisibilityMode = "always"
)

// TODO docs
type GraphDrawStyle string

const (
	GraphDrawStyleLine   GraphDrawStyle = "line"
	GraphDrawStyleBars   GraphDrawStyle = "bars"
	GraphDrawStylePoints GraphDrawStyle = "points"
)

// TODO docs
type GraphTransform string

const (
	GraphTransformConstant  GraphTransform = "constant"
	GraphTransformNegativeY GraphTransform = "negative-Y"
)

// TODO docs
type LineInterpolation string

const (
	LineInterpolationLinear     LineInterpolation = "linear"
	LineInterpolationSmooth     LineInterpolation = "smooth"
	LineInterpolationStepBefore LineInterpolation = "stepBefore"
	LineInterpolationStepAfter  LineInterpolation = "stepAfter"
)

// TODO docs
type ScaleDistribution string

const (
	ScaleDistributionLinear  ScaleDistribution = "linear"
	ScaleDistributionLog     ScaleDistribution = "log"
	ScaleDistributionOrdinal ScaleDistribution = "ordinal"
	ScaleDistributionSymlog  ScaleDistribution = "symlog"
)

// TODO docs
type GraphGradientMode string

const (
	GraphGradientModeNone    GraphGradientMode = "none"
	GraphGradientModeOpacity GraphGradientMode = "opacity"
	GraphGradientModeHue     GraphGradientMode = "hue"
	GraphGradientModeScheme  GraphGradientMode = "scheme"
)

// TODO docs
type StackingMode string

const (
	StackingModeNone    StackingMode = "none"
	StackingModeNormal  StackingMode = "normal"
	StackingModePercent StackingMode = "percent"
)

// TODO docs
type BarAlignment int64

const (
	BarAlignmentBefore BarAlignment = -1
	BarAlignmentCenter BarAlignment = 0
	BarAlignmentAfter  BarAlignment = 1
)

// TODO docs
type ScaleOrientation int64

const (
	ScaleOrientationHorizontal ScaleOrientation = 0
	ScaleOrientationVertical   ScaleOrientation = 1
)

// TODO docs
type ScaleDirection int64

const (
	ScaleDirectionUp    ScaleDirection = 1
	ScaleDirectionRight ScaleDirection = 1
	ScaleDirectionDown  ScaleDirection = -1
	ScaleDirectionLeft  ScaleDirection = -1
)

// TODO docs
type LineStyle struct {
	Fill *LineStyleFill `json:"fill,omitempty"`
	Dash []float64      `json:"dash,omitempty"`
}

// TODO docs
type LineConfig struct {
	LineColor         *string            `json:"lineColor,omitempty"`
	LineWidth         *float64           `json:"lineWidth,omitempty"`
	LineInterpolation *LineInterpolation `json:"lineInterpolation,omitempty"`
	LineStyle         *LineStyle         `json:"lineStyle,omitempty"`
	// Indicate if null values should be treated as gaps or connected.
	// When the value is a number, it represents the maximum delta in the
	// X axis that should be considered connected.  For timeseries, this is milliseconds
	SpanNulls *BoolOrFloat64 `json:"spanNulls,omitempty"`
}

// TODO docs
type BarConfig struct {
	BarAlignment   *BarAlignment `json:"barAlignment,omitempty"`
	BarWidthFactor *float64      `json:"barWidthFactor,omitempty"`
	BarMaxWidth    *float64      `json:"barMaxWidth,omitempty"`
}

// TODO docs
type FillConfig struct {
	FillColor   *string  `json:"fillColor,omitempty"`
	FillOpacity *float64 `json:"fillOpacity,omitempty"`
	FillBelowTo *string  `json:"fillBelowTo,omitempty"`
}

// TODO docs
type PointsConfig struct {
	ShowPoints  *VisibilityMode `json:"showPoints,omitempty"`
	PointSize   *float64        `json:"pointSize,omitempty"`
	PointColor  *string         `json:"pointColor,omitempty"`
	PointSymbol *string         `json:"pointSymbol,omitempty"`
}

// TODO docs
type ScaleDistributionConfig struct {
	Type            ScaleDistribution `json:"type"`
	Log             *float64          `json:"log,omitempty"`
	LinearThreshold *float64          `json:"linearThreshold,omitempty"`
}

// TODO docs
type AxisConfig struct {
	AxisPlacement     *AxisPlacement           `json:"axisPlacement,omitempty"`
	AxisColorMode     *AxisColorMode           `json:"axisColorMode,omitempty"`
	AxisLabel         *string                  `json:"axisLabel,omitempty"`
	AxisWidth         *float64                 `json:"axisWidth,omitempty"`
	AxisSoftMin       *float64                 `json:"axisSoftMin,omitempty"`
	AxisSoftMax       *float64                 `json:"axisSoftMax,omitempty"`
	AxisGridShow      *bool                    `json:"axisGridShow,omitempty"`
	ScaleDistribution *ScaleDistributionConfig `json:"scaleDistribution,omitempty"`
	AxisCenteredZero  *bool                    `json:"axisCenteredZero,omitempty"`
	AxisBorderShow    *bool                    `json:"axisBorderShow,omitempty"`
}

// TODO docs
type HideSeriesConfig struct {
	Tooltip bool `json:"tooltip"`
	Legend  bool `json:"legend"`
	Viz     bool `json:"viz"`
}

// TODO docs
type StackingConfig struct {
	Mode  *StackingMode `json:"mode,omitempty"`
	Group *string       `json:"group,omitempty"`
}

// TODO docs
type StackableFieldConfig struct {
	Stacking *StackingConfig `json:"stacking,omitempty"`
}

// TODO docs
type HideableFieldConfig struct {
	HideFrom *HideSeriesConfig `json:"hideFrom,omitempty"`
}

// TODO docs
type GraphThresholdsStyleMode string

const (
	GraphThresholdsStyleModeOff           GraphThresholdsStyleMode = "off"
	GraphThresholdsStyleModeLine          GraphThresholdsStyleMode = "line"
	GraphThresholdsStyleModeDashed        GraphThresholdsStyleMode = "dashed"
	GraphThresholdsStyleModeArea          GraphThresholdsStyleMode = "area"
	GraphThresholdsStyleModeLineAndArea   GraphThresholdsStyleMode = "line+area"
	GraphThresholdsStyleModeDashedAndArea GraphThresholdsStyleMode = "dashed+area"
	GraphThresholdsStyleModeSeries        GraphThresholdsStyleMode = "series"
)

// TODO docs
type GraphThresholdsStyleConfig struct {
	Mode GraphThresholdsStyleMode `json:"mode"`
}

// TODO docs
type LegendPlacement string

const (
	LegendPlacementBottom LegendPlacement = "bottom"
	LegendPlacementRight  LegendPlacement = "right"
)

// TODO docs
// Note: "hidden" needs to remain as an option for plugins compatibility
type LegendDisplayMode string

const (
	LegendDisplayModeList   LegendDisplayMode = "list"
	LegendDisplayModeTable  LegendDisplayMode = "table"
	LegendDisplayModeHidden LegendDisplayMode = "hidden"
)

// TODO docs
type SingleStatBaseOptions struct {
	ReduceOptions ReduceDataOptions      `json:"reduceOptions"`
	Text          *VizTextDisplayOptions `json:"text,omitempty"`
	Orientation   VizOrientation         `json:"orientation"`
}

// TODO docs
type ReduceDataOptions struct {
	// If true show each row value
	Values *bool `json:"values,omitempty"`
	// if showing all values limit
	Limit *float64 `json:"limit,omitempty"`
	// When !values, pick one value for the whole field
	Calcs []string `json:"calcs"`
	// Which fields to show.  By default this is only numeric fields
	Fields *string `json:"fields,omitempty"`
}

// TODO docs
type VizOrientation string

const (
	VizOrientationAuto       VizOrientation = "auto"
	VizOrientationVertical   VizOrientation = "vertical"
	VizOrientationHorizontal VizOrientation = "horizontal"
)

// TODO docs
type OptionsWithTooltip struct {
	Tooltip VizTooltipOptions `json:"tooltip"`
}

// TODO docs
type OptionsWithLegend struct {
	Legend VizLegendOptions `json:"legend"`
}

// TODO docs
type OptionsWithTimezones struct {
	Timezone []TimeZone `json:"timezone,omitempty"`
}

// TODO docs
type OptionsWithTextFormatting struct {
	Text *VizTextDisplayOptions `json:"text,omitempty"`
}

// TODO docs
type BigValueColorMode string

const (
	BigValueColorModeValue           BigValueColorMode = "value"
	BigValueColorModeBackground      BigValueColorMode = "background"
	BigValueColorModeBackgroundSolid BigValueColorMode = "background_solid"
	BigValueColorModeNone            BigValueColorMode = "none"
)

// TODO docs
type BigValueGraphMode string

const (
	BigValueGraphModeNone BigValueGraphMode = "none"
	BigValueGraphModeLine BigValueGraphMode = "line"
	BigValueGraphModeArea BigValueGraphMode = "area"
)

// TODO docs
type BigValueJustifyMode string

const (
	BigValueJustifyModeAuto   BigValueJustifyMode = "auto"
	BigValueJustifyModeCenter BigValueJustifyMode = "center"
)

// TODO docs
type BigValueTextMode string

const (
	BigValueTextModeAuto         BigValueTextMode = "auto"
	BigValueTextModeValue        BigValueTextMode = "value"
	BigValueTextModeValueAndName BigValueTextMode = "value_and_name"
	BigValueTextModeName         BigValueTextMode = "name"
	BigValueTextModeNone         BigValueTextMode = "none"
)

// TODO docs
type PercentChangeColorMode string

const (
	PercentChangeColorModeStandard    PercentChangeColorMode = "standard"
	PercentChangeColorModeInverted    PercentChangeColorMode = "inverted"
	PercentChangeColorModeSameAsValue PercentChangeColorMode = "same_as_value"
)

// TODO -- should not be table specific!
// TODO docs
type FieldTextAlignment string

const (
	FieldTextAlignmentAuto   FieldTextAlignment = "auto"
	FieldTextAlignmentLeft   FieldTextAlignment = "left"
	FieldTextAlignmentRight  FieldTextAlignment = "right"
	FieldTextAlignmentCenter FieldTextAlignment = "center"
)

// Controls the value alignment in the TimelineChart component
type TimelineValueAlignment string

const (
	TimelineValueAlignmentCenter TimelineValueAlignment = "center"
	TimelineValueAlignmentLeft   TimelineValueAlignment = "left"
	TimelineValueAlignmentRight  TimelineValueAlignment = "right"
)

// TODO docs
type VizTextDisplayOptions struct {
	// Explicit title text size
	TitleSize *float64 `json:"titleSize,omitempty"`
	// Explicit value text size
	ValueSize *float64 `json:"valueSize,omitempty"`
}

// TODO docs
type TooltipDisplayMode string

const (
	TooltipDisplayModeSingle TooltipDisplayMode = "single"
	TooltipDisplayModeMulti  TooltipDisplayMode = "multi"
	TooltipDisplayModeNone   TooltipDisplayMode = "none"
)

// TODO docs
type SortOrder string

const (
	SortOrderAscending  SortOrder = "asc"
	SortOrderDescending SortOrder = "desc"
	SortOrderNone       SortOrder = "none"
)

// TODO docs
type GraphFieldConfig struct {
	DrawStyle         *GraphDrawStyle             `json:"drawStyle,omitempty"`
	GradientMode      *GraphGradientMode          `json:"gradientMode,omitempty"`
	ThresholdsStyle   *GraphThresholdsStyleConfig `json:"thresholdsStyle,omitempty"`
	Transform         *GraphTransform             `json:"transform,omitempty"`
	LineColor         *string                     `json:"lineColor,omitempty"`
	LineWidth         *float64                    `json:"lineWidth,omitempty"`
	LineInterpolation *LineInterpolation          `json:"lineInterpolation,omitempty"`
	LineStyle         *LineStyle                  `json:"lineStyle,omitempty"`
	FillColor         *string                     `json:"fillColor,omitempty"`
	FillOpacity       *float64                    `json:"fillOpacity,omitempty"`
	ShowPoints        *VisibilityMode             `json:"showPoints,omitempty"`
	PointSize         *float64                    `json:"pointSize,omitempty"`
	PointColor        *string                     `json:"pointColor,omitempty"`
	AxisPlacement     *AxisPlacement              `json:"axisPlacement,omitempty"`
	AxisColorMode     *AxisColorMode              `json:"axisColorMode,omitempty"`
	AxisLabel         *string                     `json:"axisLabel,omitempty"`
	AxisWidth         *float64                    `json:"axisWidth,omitempty"`
	AxisSoftMin       *float64                    `json:"axisSoftMin,omitempty"`
	AxisSoftMax       *float64                    `json:"axisSoftMax,omitempty"`
	AxisGridShow      *bool                       `json:"axisGridShow,omitempty"`
	ScaleDistribution *ScaleDistributionConfig    `json:"scaleDistribution,omitempty"`
	AxisCenteredZero  *bool                       `json:"axisCenteredZero,omitempty"`
	BarAlignment      *BarAlignment               `json:"barAlignment,omitempty"`
	BarWidthFactor    *float64                    `json:"barWidthFactor,omitempty"`
	Stacking          *StackingConfig             `json:"stacking,omitempty"`
	HideFrom          *HideSeriesConfig           `json:"hideFrom,omitempty"`
	InsertNulls       *BoolOrFloat64              `json:"insertNulls,omitempty"`
	// Indicate if null values should be treated as gaps or connected.
	// When the value is a number, it represents the maximum delta in the
	// X axis that should be considered connected.  For timeseries, this is milliseconds
	SpanNulls      *BoolOrFloat64 `json:"spanNulls,omitempty"`
	FillBelowTo    *string        `json:"fillBelowTo,omitempty"`
	PointSymbol    *string        `json:"pointSymbol,omitempty"`
	AxisBorderShow *bool          `json:"axisBorderShow,omitempty"`
	BarMaxWidth    *float64       `json:"barMaxWidth,omitempty"`
}

// TODO docs
type VizLegendOptions struct {
	DisplayMode LegendDisplayMode `json:"displayMode"`
	Placement   LegendPlacement   `json:"placement"`
	ShowLegend  bool              `json:"showLegend"`
	AsTable     *bool             `json:"asTable,omitempty"`
	IsVisible   *bool             `json:"isVisible,omitempty"`
	SortBy      *string           `json:"sortBy,omitempty"`
	SortDesc    *bool             `json:"sortDesc,omitempty"`
	Width       *float64          `json:"width,omitempty"`
	Calcs       []string          `json:"calcs"`
}

// Enum expressing the possible display modes
// for the bar gauge component of Grafana UI
type BarGaugeDisplayMode string

const (
	BarGaugeDisplayModeBasic    BarGaugeDisplayMode = "basic"
	BarGaugeDisplayModeLcd      BarGaugeDisplayMode = "lcd"
	BarGaugeDisplayModeGradient BarGaugeDisplayMode = "gradient"
)

// Allows for the table cell gauge display type to set the gauge mode.
type BarGaugeValueMode string

const (
	BarGaugeValueModeColor  BarGaugeValueMode = "color"
	BarGaugeValueModeText   BarGaugeValueMode = "text"
	BarGaugeValueModeHidden BarGaugeValueMode = "hidden"
)

// Allows for the bar gauge name to be placed explicitly
type BarGaugeNamePlacement string

const (
	BarGaugeNamePlacementAuto BarGaugeNamePlacement = "auto"
	BarGaugeNamePlacementTop  BarGaugeNamePlacement = "top"
	BarGaugeNamePlacementLeft BarGaugeNamePlacement = "left"
)

// Allows for the bar gauge size to be set explicitly
type BarGaugeSizing string

const (
	BarGaugeSizingAuto   BarGaugeSizing = "auto"
	BarGaugeSizingManual BarGaugeSizing = "manual"
)

// TODO docs
type VizTooltipOptions struct {
	Mode      TooltipDisplayMode `json:"mode"`
	Sort      SortOrder          `json:"sort"`
	MaxWidth  *float64           `json:"maxWidth,omitempty"`
	MaxHeight *float64           `json:"maxHeight,omitempty"`
}

type Labels map[string]string

// Internally, this is the "type" of cell that's being displayed
// in the table such as colored text, JSON, gauge, etc.
// The color-background-solid, gradient-gauge, and lcd-gauge
// modes are deprecated in favor of new cell subOptions
type TableCellDisplayMode string

const (
	TableCellDisplayModeAuto                 TableCellDisplayMode = "auto"
	TableCellDisplayModeColorText            TableCellDisplayMode = "color-text"
	TableCellDisplayModeColorBackground      TableCellDisplayMode = "color-background"
	TableCellDisplayModeColorBackgroundSolid TableCellDisplayMode = "color-background-solid"
	TableCellDisplayModeGradientGauge        TableCellDisplayMode = "gradient-gauge"
	TableCellDisplayModeLcdGauge             TableCellDisplayMode = "lcd-gauge"
	TableCellDisplayModeJSONView             TableCellDisplayMode = "json-view"
	TableCellDisplayModeBasicGauge           TableCellDisplayMode = "basic"
	TableCellDisplayModeImage                TableCellDisplayMode = "image"
	TableCellDisplayModeGauge                TableCellDisplayMode = "gauge"
	TableCellDisplayModeSparkline            TableCellDisplayMode = "sparkline"
	TableCellDisplayModeDataLinks            TableCellDisplayMode = "data-links"
	TableCellDisplayModeCustom               TableCellDisplayMode = "custom"
)

// Display mode to the "Colored Background" display
// mode for table cells. Either displays a solid color (basic mode)
// or a gradient.
type TableCellBackgroundDisplayMode string

const (
	TableCellBackgroundDisplayModeBasic    TableCellBackgroundDisplayMode = "basic"
	TableCellBackgroundDisplayModeGradient TableCellBackgroundDisplayMode = "gradient"
)

// Sort by field state
type TableSortByFieldState struct {
	// Sets the display name of the field to sort by
	DisplayName string `json:"displayName"`
	// Flag used to indicate descending sort order
	Desc *bool `json:"desc,omitempty"`
}

// Footer options
type TableFooterOptions struct {
	Show bool `json:"show"`
	// actually 1 value
	Reducer          []string `json:"reducer"`
	Fields           []string `json:"fields,omitempty"`
	EnablePagination *bool    `json:"enablePagination,omitempty"`
	CountRows        *bool    `json:"countRows,omitempty"`
}

// Auto mode table cell options
type TableAutoCellOptions struct {
	Type     string `json:"type"`
	WrapText *bool  `json:"wrapText,omitempty"`
}

// Colored text cell options
type TableColorTextCellOptions struct {
	Type     string `json:"type"`
	WrapText *bool  `json:"wrapText,omitempty"`
}

// Json view cell options
type TableJsonViewCellOptions struct {
	Type string `json:"type"`
}

// Json view cell options
type TableImageCellOptions struct {
	Type  string  `json:"type"`
	Alt   *string `json:"alt,omitempty"`
	Title *string `json:"title,omitempty"`
}

// Show data links in the cell
type TableDataLinksCellOptions struct {
	Type string `json:"type"`
}

// Gauge cell options
type TableBarGaugeCellOptions struct {
	Type             string               `json:"type"`
	Mode             *BarGaugeDisplayMode `json:"mode,omitempty"`
	ValueDisplayMode *BarGaugeValueMode   `json:"valueDisplayMode,omitempty"`
}

// Sparkline cell options
type TableSparklineCellOptions struct {
	Type              string                      `json:"type"`
	DrawStyle         *GraphDrawStyle             `json:"drawStyle,omitempty"`
	GradientMode      *GraphGradientMode          `json:"gradientMode,omitempty"`
	ThresholdsStyle   *GraphThresholdsStyleConfig `json:"thresholdsStyle,omitempty"`
	Transform         *GraphTransform             `json:"transform,omitempty"`
	LineColor         *string                     `json:"lineColor,omitempty"`
	LineWidth         *float64                    `json:"lineWidth,omitempty"`
	LineInterpolation *LineInterpolation          `json:"lineInterpolation,omitempty"`
	LineStyle         *LineStyle                  `json:"lineStyle,omitempty"`
	FillColor         *string                     `json:"fillColor,omitempty"`
	FillOpacity       *float64                    `json:"fillOpacity,omitempty"`
	ShowPoints        *VisibilityMode             `json:"showPoints,omitempty"`
	PointSize         *float64                    `json:"pointSize,omitempty"`
	PointColor        *string                     `json:"pointColor,omitempty"`
	AxisPlacement     *AxisPlacement              `json:"axisPlacement,omitempty"`
	AxisColorMode     *AxisColorMode              `json:"axisColorMode,omitempty"`
	AxisLabel         *string                     `json:"axisLabel,omitempty"`
	AxisWidth         *float64                    `json:"axisWidth,omitempty"`
	AxisSoftMin       *float64                    `json:"axisSoftMin,omitempty"`
	AxisSoftMax       *float64                    `json:"axisSoftMax,omitempty"`
	AxisGridShow      *bool                       `json:"axisGridShow,omitempty"`
	ScaleDistribution *ScaleDistributionConfig    `json:"scaleDistribution,omitempty"`
	AxisCenteredZero  *bool                       `json:"axisCenteredZero,omitempty"`
	BarAlignment      *BarAlignment               `json:"barAlignment,omitempty"`
	BarWidthFactor    *float64                    `json:"barWidthFactor,omitempty"`
	Stacking          *StackingConfig             `json:"stacking,omitempty"`
	HideFrom          *HideSeriesConfig           `json:"hideFrom,omitempty"`
	HideValue         *bool                       `json:"hideValue,omitempty"`
	InsertNulls       *BoolOrFloat64              `json:"insertNulls,omitempty"`
	// Indicate if null values should be treated as gaps or connected.
	// When the value is a number, it represents the maximum delta in the
	// X axis that should be considered connected.  For timeseries, this is milliseconds
	SpanNulls      *BoolOrFloat64 `json:"spanNulls,omitempty"`
	FillBelowTo    *string        `json:"fillBelowTo,omitempty"`
	PointSymbol    *string        `json:"pointSymbol,omitempty"`
	AxisBorderShow *bool          `json:"axisBorderShow,omitempty"`
	BarMaxWidth    *float64       `json:"barMaxWidth,omitempty"`
}

// Colored background cell options
type TableColoredBackgroundCellOptions struct {
	Type       string                          `json:"type"`
	Mode       *TableCellBackgroundDisplayMode `json:"mode,omitempty"`
	ApplyToRow *bool                           `json:"applyToRow,omitempty"`
	WrapText   *bool                           `json:"wrapText,omitempty"`
}

// Height of a table cell
type TableCellHeight string

const (
	TableCellHeightSm   TableCellHeight = "sm"
	TableCellHeightMd   TableCellHeight = "md"
	TableCellHeightLg   TableCellHeight = "lg"
	TableCellHeightAuto TableCellHeight = "auto"
)

// Table cell options. Each cell has a display mode
// and other potential options for that display.
type TableCellOptions = TableAutoCellOptionsOrTableSparklineCellOptionsOrTableBarGaugeCellOptionsOrTableColoredBackgroundCellOptionsOrTableColorTextCellOptionsOrTableImageCellOptionsOrTableDataLinksCellOptionsOrTableJsonViewCellOptions

// Use UTC/GMT timezone
const TimeZoneUtc = "utc"

// Use the timezone defined by end user web browser
const TimeZoneBrowser = "browser"

// Optional formats for the template variable replace functions
// See also https://grafana.com/docs/grafana/latest/dashboards/variables/variable-syntax/#advanced-variable-format-options
type VariableFormatID string

const (
	VariableFormatIDLucene        VariableFormatID = "lucene"
	VariableFormatIDRaw           VariableFormatID = "raw"
	VariableFormatIDRegex         VariableFormatID = "regex"
	VariableFormatIDPipe          VariableFormatID = "pipe"
	VariableFormatIDDistributed   VariableFormatID = "distributed"
	VariableFormatIDCSV           VariableFormatID = "csv"
	VariableFormatIDHTML          VariableFormatID = "html"
	VariableFormatIDJSON          VariableFormatID = "json"
	VariableFormatIDPercentEncode VariableFormatID = "percentencode"
	VariableFormatIDUriEncode     VariableFormatID = "uriencode"
	VariableFormatIDSingleQuote   VariableFormatID = "singlequote"
	VariableFormatIDDoubleQuote   VariableFormatID = "doublequote"
	VariableFormatIDSQLString     VariableFormatID = "sqlstring"
	VariableFormatIDDate          VariableFormatID = "date"
	VariableFormatIDGlob          VariableFormatID = "glob"
	VariableFormatIDText          VariableFormatID = "text"
	VariableFormatIDQueryParam    VariableFormatID = "queryparam"
)

type DataSourceRef struct {
	// The plugin type-id
	Type *string `json:"type,omitempty"`
	// Specific datasource instance
	Uid *string `json:"uid,omitempty"`
	// Datasource API version
	ApiVersion *string `json:"apiVersion,omitempty"`
}

// Links to a resource (image/svg path)
type ResourceDimensionConfig struct {
	Mode ResourceDimensionMode `json:"mode"`
	// fixed: T -- will be added by each element
	Field *string `json:"field,omitempty"`
	Fixed *string `json:"fixed,omitempty"`
}

type FrameGeometrySource struct {
	Mode FrameGeometrySourceMode `json:"mode"`
	// Field mappings
	Geohash   *string `json:"geohash,omitempty"`
	Latitude  *string `json:"latitude,omitempty"`
	Longitude *string `json:"longitude,omitempty"`
	Wkt       *string `json:"wkt,omitempty"`
	Lookup    *string `json:"lookup,omitempty"`
	// Path to Gazetteer
	Gazetteer *string `json:"gazetteer,omitempty"`
}

type HeatmapCalculationOptions struct {
	// The number of buckets to use for the xAxis in the heatmap
	XBuckets *HeatmapCalculationBucketConfig `json:"xBuckets,omitempty"`
	// The number of buckets to use for the yAxis in the heatmap
	YBuckets *HeatmapCalculationBucketConfig `json:"yBuckets,omitempty"`
}

type LogsDedupStrategy string

const (
	LogsDedupStrategyNone      LogsDedupStrategy = "none"
	LogsDedupStrategyExact     LogsDedupStrategy = "exact"
	LogsDedupStrategyNumbers   LogsDedupStrategy = "numbers"
	LogsDedupStrategySignature LogsDedupStrategy = "signature"
)

// Compare two values
type ComparisonOperation string

const (
	ComparisonOperationEQ  ComparisonOperation = "eq"
	ComparisonOperationNEQ ComparisonOperation = "neq"
	ComparisonOperationLT  ComparisonOperation = "lt"
	ComparisonOperationLTE ComparisonOperation = "lte"
	ComparisonOperationGT  ComparisonOperation = "gt"
	ComparisonOperationGTE ComparisonOperation = "gte"
)

// Field options for each field within a table (e.g 10, "The String", 64.20, etc.)
// Generally defines alignment, filtering capabilties, display options, etc.
type TableFieldOptions struct {
	Width    *float64           `json:"width,omitempty"`
	MinWidth *float64           `json:"minWidth,omitempty"`
	Align    FieldTextAlignment `json:"align"`
	// This field is deprecated in favor of using cellOptions
	DisplayMode *TableCellDisplayMode `json:"displayMode,omitempty"`
	CellOptions *TableCellOptions     `json:"cellOptions,omitempty"`
	// ?? default is missing or false ??
	Hidden     *bool `json:"hidden,omitempty"`
	Inspect    bool  `json:"inspect"`
	Filterable *bool `json:"filterable,omitempty"`
	// Hides any header for a column, useful for columns that show some static content or buttons.
	HideHeader *bool `json:"hideHeader,omitempty"`
}

// A specific timezone from https://en.wikipedia.org/wiki/Tz_database
type TimeZone string

type LineStyleFill string

const (
	LineStyleFillSolid  LineStyleFill = "solid"
	LineStyleFillDash   LineStyleFill = "dash"
	LineStyleFillDot    LineStyleFill = "dot"
	LineStyleFillSquare LineStyleFill = "square"
)

type BoolOrFloat64 struct {
	Bool    *bool    `json:"Bool,omitempty"`
	Float64 *float64 `json:"Float64,omitempty"`
}

func (resource BoolOrFloat64) MarshalJSON() ([]byte, error) {
	if resource.Bool != nil {
		return json.Marshal(resource.Bool)
	}

	if resource.Float64 != nil {
		return json.Marshal(resource.Float64)
	}

	return nil, fmt.Errorf("no value for disjunction of scalars")
}

func (resource *BoolOrFloat64) UnmarshalJSON(raw []byte) error {
	if raw == nil {
		return nil
	}

	var errList []error

	// Bool
	var Bool bool
	if err := json.Unmarshal(raw, &Bool); err != nil {
		errList = append(errList, err)
		resource.Bool = nil
	} else {
		resource.Bool = &Bool
		return nil
	}

	// Float64
	var Float64 float64
	if err := json.Unmarshal(raw, &Float64); err != nil {
		errList = append(errList, err)
		resource.Float64 = nil
	} else {
		resource.Float64 = &Float64
		return nil
	}

	return errors.Join(errList...)
}

type TableAutoCellOptionsOrTableSparklineCellOptionsOrTableBarGaugeCellOptionsOrTableColoredBackgroundCellOptionsOrTableColorTextCellOptionsOrTableImageCellOptionsOrTableDataLinksCellOptionsOrTableJsonViewCellOptions struct {
	TableAutoCellOptions              *TableAutoCellOptions              `json:"TableAutoCellOptions,omitempty"`
	TableSparklineCellOptions         *TableSparklineCellOptions         `json:"TableSparklineCellOptions,omitempty"`
	TableBarGaugeCellOptions          *TableBarGaugeCellOptions          `json:"TableBarGaugeCellOptions,omitempty"`
	TableColoredBackgroundCellOptions *TableColoredBackgroundCellOptions `json:"TableColoredBackgroundCellOptions,omitempty"`
	TableColorTextCellOptions         *TableColorTextCellOptions         `json:"TableColorTextCellOptions,omitempty"`
	TableImageCellOptions             *TableImageCellOptions             `json:"TableImageCellOptions,omitempty"`
	TableDataLinksCellOptions         *TableDataLinksCellOptions         `json:"TableDataLinksCellOptions,omitempty"`
	TableJsonViewCellOptions          *TableJsonViewCellOptions          `json:"TableJsonViewCellOptions,omitempty"`
}

func (resource TableAutoCellOptionsOrTableSparklineCellOptionsOrTableBarGaugeCellOptionsOrTableColoredBackgroundCellOptionsOrTableColorTextCellOptionsOrTableImageCellOptionsOrTableDataLinksCellOptionsOrTableJsonViewCellOptions) MarshalJSON() ([]byte, error) {
	if resource.TableAutoCellOptions != nil {
		return json.Marshal(resource.TableAutoCellOptions)
	}
	if resource.TableSparklineCellOptions != nil {
		return json.Marshal(resource.TableSparklineCellOptions)
	}
	if resource.TableBarGaugeCellOptions != nil {
		return json.Marshal(resource.TableBarGaugeCellOptions)
	}
	if resource.TableColoredBackgroundCellOptions != nil {
		return json.Marshal(resource.TableColoredBackgroundCellOptions)
	}
	if resource.TableColorTextCellOptions != nil {
		return json.Marshal(resource.TableColorTextCellOptions)
	}
	if resource.TableImageCellOptions != nil {
		return json.Marshal(resource.TableImageCellOptions)
	}
	if resource.TableDataLinksCellOptions != nil {
		return json.Marshal(resource.TableDataLinksCellOptions)
	}
	if resource.TableJsonViewCellOptions != nil {
		return json.Marshal(resource.TableJsonViewCellOptions)
	}

	return nil, fmt.Errorf("no value for disjunction of refs")
}

func (resource *TableAutoCellOptionsOrTableSparklineCellOptionsOrTableBarGaugeCellOptionsOrTableColoredBackgroundCellOptionsOrTableColorTextCellOptionsOrTableImageCellOptionsOrTableDataLinksCellOptionsOrTableJsonViewCellOptions) UnmarshalJSON(raw []byte) error {
	if raw == nil {
		return nil
	}

	// FIXME: this is wasteful, we need to find a more efficient way to unmarshal this.
	parsedAsMap := make(map[string]any)
	if err := json.Unmarshal(raw, &parsedAsMap); err != nil {
		return err
	}

	discriminator, found := parsedAsMap["type"]
	if !found {
		return errors.New("discriminator field 'type' not found in payload")
	}

	switch discriminator {
	case "auto":
		var tableAutoCellOptions TableAutoCellOptions
		if err := json.Unmarshal(raw, &tableAutoCellOptions); err != nil {
			return err
		}

		resource.TableAutoCellOptions = &tableAutoCellOptions
		return nil
	case "color-background":
		var tableColoredBackgroundCellOptions TableColoredBackgroundCellOptions
		if err := json.Unmarshal(raw, &tableColoredBackgroundCellOptions); err != nil {
			return err
		}

		resource.TableColoredBackgroundCellOptions = &tableColoredBackgroundCellOptions
		return nil
	case "color-text":
		var tableColorTextCellOptions TableColorTextCellOptions
		if err := json.Unmarshal(raw, &tableColorTextCellOptions); err != nil {
			return err
		}

		resource.TableColorTextCellOptions = &tableColorTextCellOptions
		return nil
	case "data-links":
		var tableDataLinksCellOptions TableDataLinksCellOptions
		if err := json.Unmarshal(raw, &tableDataLinksCellOptions); err != nil {
			return err
		}

		resource.TableDataLinksCellOptions = &tableDataLinksCellOptions
		return nil
	case "gauge":
		var tableBarGaugeCellOptions TableBarGaugeCellOptions
		if err := json.Unmarshal(raw, &tableBarGaugeCellOptions); err != nil {
			return err
		}

		resource.TableBarGaugeCellOptions = &tableBarGaugeCellOptions
		return nil
	case "image":
		var tableImageCellOptions TableImageCellOptions
		if err := json.Unmarshal(raw, &tableImageCellOptions); err != nil {
			return err
		}

		resource.TableImageCellOptions = &tableImageCellOptions
		return nil
	case "json-view":
		var tableJsonViewCellOptions TableJsonViewCellOptions
		if err := json.Unmarshal(raw, &tableJsonViewCellOptions); err != nil {
			return err
		}

		resource.TableJsonViewCellOptions = &tableJsonViewCellOptions
		return nil
	case "sparkline":
		var tableSparklineCellOptions TableSparklineCellOptions
		if err := json.Unmarshal(raw, &tableSparklineCellOptions); err != nil {
			return err
		}

		resource.TableSparklineCellOptions = &tableSparklineCellOptions
		return nil
	}

	return fmt.Errorf("could not unmarshal resource with `type = %v`", discriminator)
}
