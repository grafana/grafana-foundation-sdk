# Code generated - EDITING IS FUTILE. DO NOT EDIT.

import typing
import enum


class DataSourceJsonData:
    """
    TODO docs
    """

    auth_type: typing.Optional[str]
    default_region: typing.Optional[str]
    profile: typing.Optional[str]
    manage_alerts: typing.Optional[bool]
    alertmanager_uid: typing.Optional[str]

    def __init__(self, auth_type: typing.Optional[str] = None, default_region: typing.Optional[str] = None, profile: typing.Optional[str] = None, manage_alerts: typing.Optional[bool] = None, alertmanager_uid: typing.Optional[str] = None):
        self.auth_type = auth_type
        self.default_region = default_region
        self.profile = profile
        self.manage_alerts = manage_alerts
        self.alertmanager_uid = alertmanager_uid

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.auth_type is not None:
            payload["authType"] = self.auth_type
        if self.default_region is not None:
            payload["defaultRegion"] = self.default_region
        if self.profile is not None:
            payload["profile"] = self.profile
        if self.manage_alerts is not None:
            payload["manageAlerts"] = self.manage_alerts
        if self.alertmanager_uid is not None:
            payload["alertmanagerUid"] = self.alertmanager_uid
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "authType" in data:
            args["auth_type"] = data["authType"]
        if "defaultRegion" in data:
            args["default_region"] = data["defaultRegion"]
        if "profile" in data:
            args["profile"] = data["profile"]
        if "manageAlerts" in data:
            args["manage_alerts"] = data["manageAlerts"]
        if "alertmanagerUid" in data:
            args["alertmanager_uid"] = data["alertmanagerUid"]        

        return cls(**args)


class DataQuery:
    """
    These are the common properties available to all queries in all datasources.
    Specific implementations will *extend* this interface, adding the required
    properties for the given context.
    """

    # A unique identifier for the query within the list of targets.
    # In server side expressions, the refId is used as a variable name to identify results.
    # By default, the UI will assign A->Z; however setting meaningful names may be useful.
    ref_id: str
    # true if query is disabled (ie should not be returned to the dashboard)
    # Note this does not always imply that the query should not be executed since
    # the results from a hidden query may be used as the input to other queries (SSE etc)
    hide: typing.Optional[bool]
    # Specify the query flavor
    # TODO make this required and give it a default
    query_type: typing.Optional[str]
    # For mixed data sources the selected datasource is on the query level.
    # For non mixed scenarios this is undefined.
    # TODO find a better way to do this ^ that's friendly to schema
    # TODO this shouldn't be unknown but DataSourceRef | null
    datasource: typing.Optional[object]

    def __init__(self, ref_id: str = "", hide: typing.Optional[bool] = None, query_type: typing.Optional[str] = None, datasource: typing.Optional[object] = None):
        self.ref_id = ref_id
        self.hide = hide
        self.query_type = query_type
        self.datasource = datasource

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "refId": self.ref_id,
        }
        if self.hide is not None:
            payload["hide"] = self.hide
        if self.query_type is not None:
            payload["queryType"] = self.query_type
        if self.datasource is not None:
            payload["datasource"] = self.datasource
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "refId" in data:
            args["ref_id"] = data["refId"]
        if "hide" in data:
            args["hide"] = data["hide"]
        if "queryType" in data:
            args["query_type"] = data["queryType"]
        if "datasource" in data:
            args["datasource"] = data["datasource"]        

        return cls(**args)


class BaseDimensionConfig:
    # fixed: T -- will be added by each element
    field: typing.Optional[str]

    def __init__(self, field: typing.Optional[str] = None):
        self.field = field

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.field is not None:
            payload["field"] = self.field
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "field" in data:
            args["field"] = data["field"]        

        return cls(**args)


class ScaleDimensionMode(enum.StrEnum):
    LINEAR = "linear"
    QUAD = "quad"


class ScaleDimensionConfig:
    min_val: float
    max_val: float
    fixed: typing.Optional[float]
    # fixed: T -- will be added by each element
    field: typing.Optional[str]
    # | *"linear"
    mode: typing.Optional['ScaleDimensionMode']

    def __init__(self, min_val: float = 0, max_val: float = 0, fixed: typing.Optional[float] = None, field: typing.Optional[str] = None, mode: typing.Optional['ScaleDimensionMode'] = None):
        self.min_val = min_val
        self.max_val = max_val
        self.fixed = fixed
        self.field = field
        self.mode = mode

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "min": self.min_val,
            "max": self.max_val,
        }
        if self.fixed is not None:
            payload["fixed"] = self.fixed
        if self.field is not None:
            payload["field"] = self.field
        if self.mode is not None:
            payload["mode"] = self.mode
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "min" in data:
            args["min_val"] = data["min"]
        if "max" in data:
            args["max_val"] = data["max"]
        if "fixed" in data:
            args["fixed"] = data["fixed"]
        if "field" in data:
            args["field"] = data["field"]
        if "mode" in data:
            args["mode"] = data["mode"]        

        return cls(**args)


class ColorDimensionConfig:
    # color value
    fixed: typing.Optional[str]
    # fixed: T -- will be added by each element
    field: typing.Optional[str]

    def __init__(self, fixed: typing.Optional[str] = None, field: typing.Optional[str] = None):
        self.fixed = fixed
        self.field = field

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.fixed is not None:
            payload["fixed"] = self.fixed
        if self.field is not None:
            payload["field"] = self.field
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "fixed" in data:
            args["fixed"] = data["fixed"]
        if "field" in data:
            args["field"] = data["field"]        

        return cls(**args)


class ScalarDimensionMode(enum.StrEnum):
    MOD = "mod"
    CLAMPED = "clamped"


class ScalarDimensionConfig:
    min_val: float
    max_val: float
    fixed: typing.Optional[float]
    # fixed: T -- will be added by each element
    field: typing.Optional[str]
    mode: typing.Optional['ScalarDimensionMode']

    def __init__(self, min_val: float = 0, max_val: float = 0, fixed: typing.Optional[float] = None, field: typing.Optional[str] = None, mode: typing.Optional['ScalarDimensionMode'] = None):
        self.min_val = min_val
        self.max_val = max_val
        self.fixed = fixed
        self.field = field
        self.mode = mode

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "min": self.min_val,
            "max": self.max_val,
        }
        if self.fixed is not None:
            payload["fixed"] = self.fixed
        if self.field is not None:
            payload["field"] = self.field
        if self.mode is not None:
            payload["mode"] = self.mode
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "min" in data:
            args["min_val"] = data["min"]
        if "max" in data:
            args["max_val"] = data["max"]
        if "fixed" in data:
            args["fixed"] = data["fixed"]
        if "field" in data:
            args["field"] = data["field"]
        if "mode" in data:
            args["mode"] = data["mode"]        

        return cls(**args)


class TextDimensionMode(enum.StrEnum):
    FIXED = "fixed"
    FIELD = "field"
    TEMPLATE = "template"


class TextDimensionConfig:
    mode: 'TextDimensionMode'
    # fixed: T -- will be added by each element
    field: typing.Optional[str]
    fixed: typing.Optional[str]

    def __init__(self, mode: typing.Optional['TextDimensionMode'] = None, field: typing.Optional[str] = None, fixed: typing.Optional[str] = None):
        self.mode = mode if mode is not None else TextDimensionMode.FIXED
        self.field = field
        self.fixed = fixed

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "mode": self.mode,
        }
        if self.field is not None:
            payload["field"] = self.field
        if self.fixed is not None:
            payload["fixed"] = self.fixed
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "mode" in data:
            args["mode"] = data["mode"]
        if "field" in data:
            args["field"] = data["field"]
        if "fixed" in data:
            args["fixed"] = data["fixed"]        

        return cls(**args)


class ResourceDimensionMode(enum.StrEnum):
    FIXED = "fixed"
    FIELD = "field"
    MAPPING = "mapping"


class MapLayerOptions:
    type_val: str
    # configured unique display name
    name: str
    # Custom options depending on the type
    config: typing.Optional[object]
    # Common method to define geometry fields
    location: typing.Optional['FrameGeometrySource']
    # Defines a frame MatcherConfig that may filter data for the given layer
    filter_data: typing.Optional[object]
    # Common properties:
    # https://openlayers.org/en/latest/apidoc/module-ol_layer_Base-BaseLayer.html
    # Layer opacity (0-1)
    opacity: typing.Optional[int]
    # Check tooltip (defaults to true)
    tooltip: typing.Optional[bool]

    def __init__(self, type_val: str = "", name: str = "", config: typing.Optional[object] = None, location: typing.Optional['FrameGeometrySource'] = None, filter_data: typing.Optional[object] = None, opacity: typing.Optional[int] = None, tooltip: typing.Optional[bool] = None):
        self.type_val = type_val
        self.name = name
        self.config = config
        self.location = location
        self.filter_data = filter_data
        self.opacity = opacity
        self.tooltip = tooltip

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "type": self.type_val,
            "name": self.name,
        }
        if self.config is not None:
            payload["config"] = self.config
        if self.location is not None:
            payload["location"] = self.location
        if self.filter_data is not None:
            payload["filterData"] = self.filter_data
        if self.opacity is not None:
            payload["opacity"] = self.opacity
        if self.tooltip is not None:
            payload["tooltip"] = self.tooltip
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "type" in data:
            args["type_val"] = data["type"]
        if "name" in data:
            args["name"] = data["name"]
        if "config" in data:
            args["config"] = data["config"]
        if "location" in data:
            args["location"] = FrameGeometrySource.from_json(data["location"])
        if "filterData" in data:
            args["filter_data"] = data["filterData"]
        if "opacity" in data:
            args["opacity"] = data["opacity"]
        if "tooltip" in data:
            args["tooltip"] = data["tooltip"]        

        return cls(**args)


class FrameGeometrySourceMode(enum.StrEnum):
    AUTO = "auto"
    GEOHASH = "geohash"
    COORDS = "coords"
    LOOKUP = "lookup"


class HeatmapCalculationMode(enum.StrEnum):
    SIZE = "size"
    COUNT = "count"


class HeatmapCellLayout(enum.StrEnum):
    LE = "le"
    GE = "ge"
    UNKNOWN = "unknown"
    AUTO = "auto"


class HeatmapCalculationBucketConfig:
    # Sets the bucket calculation mode
    mode: typing.Optional['HeatmapCalculationMode']
    # The number of buckets to use for the axis in the heatmap
    value: typing.Optional[str]
    # Controls the scale of the buckets
    scale: typing.Optional['ScaleDistributionConfig']

    def __init__(self, mode: typing.Optional['HeatmapCalculationMode'] = None, value: typing.Optional[str] = None, scale: typing.Optional['ScaleDistributionConfig'] = None):
        self.mode = mode
        self.value = value
        self.scale = scale

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.mode is not None:
            payload["mode"] = self.mode
        if self.value is not None:
            payload["value"] = self.value
        if self.scale is not None:
            payload["scale"] = self.scale
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "mode" in data:
            args["mode"] = data["mode"]
        if "value" in data:
            args["value"] = data["value"]
        if "scale" in data:
            args["scale"] = ScaleDistributionConfig.from_json(data["scale"])        

        return cls(**args)


class LogsSortOrder(enum.StrEnum):
    DESCENDING = "Descending"
    ASCENDING = "Ascending"


class AxisPlacement(enum.StrEnum):
    """
    TODO docs
    """

    AUTO = "auto"
    TOP = "top"
    RIGHT = "right"
    BOTTOM = "bottom"
    LEFT = "left"
    HIDDEN = "hidden"


class AxisColorMode(enum.StrEnum):
    """
    TODO docs
    """

    TEXT = "text"
    SERIES = "series"


class VisibilityMode(enum.StrEnum):
    """
    TODO docs
    """

    AUTO = "auto"
    NEVER = "never"
    ALWAYS = "always"


class GraphDrawStyle(enum.StrEnum):
    """
    TODO docs
    """

    LINE = "line"
    BARS = "bars"
    POINTS = "points"


class GraphTransform(enum.StrEnum):
    """
    TODO docs
    """

    CONSTANT = "constant"
    NEGATIVE_Y = "negative-Y"


class LineInterpolation(enum.StrEnum):
    """
    TODO docs
    """

    LINEAR = "linear"
    SMOOTH = "smooth"
    STEP_BEFORE = "stepBefore"
    STEP_AFTER = "stepAfter"


class ScaleDistribution(enum.StrEnum):
    """
    TODO docs
    """

    LINEAR = "linear"
    LOG = "log"
    ORDINAL = "ordinal"
    SYMLOG = "symlog"


class GraphGradientMode(enum.StrEnum):
    """
    TODO docs
    """

    NONE = "none"
    OPACITY = "opacity"
    HUE = "hue"
    SCHEME = "scheme"


class StackingMode(enum.StrEnum):
    """
    TODO docs
    """

    NONE = "none"
    NORMAL = "normal"
    PERCENT = "percent"


class BarAlignment(enum.IntEnum):
    """
    TODO docs
    """

    BEFORE = -1
    CENTER = 0
    AFTER = 1


class ScaleOrientation(enum.IntEnum):
    """
    TODO docs
    """

    HORIZONTAL = 0
    VERTICAL = 1


class ScaleDirection(enum.IntEnum):
    """
    TODO docs
    """

    UP = 1
    RIGHT = 1
    DOWN = -1
    LEFT = -1


class LineStyle:
    """
    TODO docs
    """

    fill: typing.Optional[typing.Literal["solid", "dash", "dot", "square"]]
    dash: typing.Optional[list[float]]

    def __init__(self, fill: typing.Optional[typing.Literal["solid", "dash", "dot", "square"]] = None, dash: typing.Optional[list[float]] = None):
        self.fill = fill
        self.dash = dash

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.fill is not None:
            payload["fill"] = self.fill
        if self.dash is not None:
            payload["dash"] = self.dash
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "fill" in data:
            args["fill"] = data["fill"]
        if "dash" in data:
            args["dash"] = data["dash"]        

        return cls(**args)


class LineConfig:
    """
    TODO docs
    """

    line_color: typing.Optional[str]
    line_width: typing.Optional[float]
    line_interpolation: typing.Optional['LineInterpolation']
    line_style: typing.Optional['LineStyle']
    # Indicate if null values should be treated as gaps or connected.
    # When the value is a number, it represents the maximum delta in the
    # X axis that should be considered connected.  For timeseries, this is milliseconds
    span_nulls: typing.Optional[typing.Union[bool, float]]

    def __init__(self, line_color: typing.Optional[str] = None, line_width: typing.Optional[float] = None, line_interpolation: typing.Optional['LineInterpolation'] = None, line_style: typing.Optional['LineStyle'] = None, span_nulls: typing.Optional[typing.Union[bool, float]] = None):
        self.line_color = line_color
        self.line_width = line_width
        self.line_interpolation = line_interpolation
        self.line_style = line_style
        self.span_nulls = span_nulls

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.line_color is not None:
            payload["lineColor"] = self.line_color
        if self.line_width is not None:
            payload["lineWidth"] = self.line_width
        if self.line_interpolation is not None:
            payload["lineInterpolation"] = self.line_interpolation
        if self.line_style is not None:
            payload["lineStyle"] = self.line_style
        if self.span_nulls is not None:
            payload["spanNulls"] = self.span_nulls
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "lineColor" in data:
            args["line_color"] = data["lineColor"]
        if "lineWidth" in data:
            args["line_width"] = data["lineWidth"]
        if "lineInterpolation" in data:
            args["line_interpolation"] = data["lineInterpolation"]
        if "lineStyle" in data:
            args["line_style"] = LineStyle.from_json(data["lineStyle"])
        if "spanNulls" in data:
            args["span_nulls"] = data["spanNulls"]        

        return cls(**args)


class BarConfig:
    """
    TODO docs
    """

    bar_alignment: typing.Optional['BarAlignment']
    bar_width_factor: typing.Optional[float]
    bar_max_width: typing.Optional[float]

    def __init__(self, bar_alignment: typing.Optional['BarAlignment'] = None, bar_width_factor: typing.Optional[float] = None, bar_max_width: typing.Optional[float] = None):
        self.bar_alignment = bar_alignment
        self.bar_width_factor = bar_width_factor
        self.bar_max_width = bar_max_width

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.bar_alignment is not None:
            payload["barAlignment"] = self.bar_alignment
        if self.bar_width_factor is not None:
            payload["barWidthFactor"] = self.bar_width_factor
        if self.bar_max_width is not None:
            payload["barMaxWidth"] = self.bar_max_width
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "barAlignment" in data:
            args["bar_alignment"] = data["barAlignment"]
        if "barWidthFactor" in data:
            args["bar_width_factor"] = data["barWidthFactor"]
        if "barMaxWidth" in data:
            args["bar_max_width"] = data["barMaxWidth"]        

        return cls(**args)


class FillConfig:
    """
    TODO docs
    """

    fill_color: typing.Optional[str]
    fill_opacity: typing.Optional[float]
    fill_below_to: typing.Optional[str]

    def __init__(self, fill_color: typing.Optional[str] = None, fill_opacity: typing.Optional[float] = None, fill_below_to: typing.Optional[str] = None):
        self.fill_color = fill_color
        self.fill_opacity = fill_opacity
        self.fill_below_to = fill_below_to

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.fill_color is not None:
            payload["fillColor"] = self.fill_color
        if self.fill_opacity is not None:
            payload["fillOpacity"] = self.fill_opacity
        if self.fill_below_to is not None:
            payload["fillBelowTo"] = self.fill_below_to
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "fillColor" in data:
            args["fill_color"] = data["fillColor"]
        if "fillOpacity" in data:
            args["fill_opacity"] = data["fillOpacity"]
        if "fillBelowTo" in data:
            args["fill_below_to"] = data["fillBelowTo"]        

        return cls(**args)


class PointsConfig:
    """
    TODO docs
    """

    show_points: typing.Optional['VisibilityMode']
    point_size: typing.Optional[float]
    point_color: typing.Optional[str]
    point_symbol: typing.Optional[str]

    def __init__(self, show_points: typing.Optional['VisibilityMode'] = None, point_size: typing.Optional[float] = None, point_color: typing.Optional[str] = None, point_symbol: typing.Optional[str] = None):
        self.show_points = show_points
        self.point_size = point_size
        self.point_color = point_color
        self.point_symbol = point_symbol

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.show_points is not None:
            payload["showPoints"] = self.show_points
        if self.point_size is not None:
            payload["pointSize"] = self.point_size
        if self.point_color is not None:
            payload["pointColor"] = self.point_color
        if self.point_symbol is not None:
            payload["pointSymbol"] = self.point_symbol
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "showPoints" in data:
            args["show_points"] = data["showPoints"]
        if "pointSize" in data:
            args["point_size"] = data["pointSize"]
        if "pointColor" in data:
            args["point_color"] = data["pointColor"]
        if "pointSymbol" in data:
            args["point_symbol"] = data["pointSymbol"]        

        return cls(**args)


class ScaleDistributionConfig:
    """
    TODO docs
    """

    type_val: 'ScaleDistribution'
    log: typing.Optional[float]
    linear_threshold: typing.Optional[float]

    def __init__(self, type_val: typing.Optional['ScaleDistribution'] = None, log: typing.Optional[float] = None, linear_threshold: typing.Optional[float] = None):
        self.type_val = type_val if type_val is not None else ScaleDistribution.LINEAR
        self.log = log
        self.linear_threshold = linear_threshold

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "type": self.type_val,
        }
        if self.log is not None:
            payload["log"] = self.log
        if self.linear_threshold is not None:
            payload["linearThreshold"] = self.linear_threshold
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "type" in data:
            args["type_val"] = data["type"]
        if "log" in data:
            args["log"] = data["log"]
        if "linearThreshold" in data:
            args["linear_threshold"] = data["linearThreshold"]        

        return cls(**args)


class AxisConfig:
    """
    TODO docs
    """

    axis_placement: typing.Optional['AxisPlacement']
    axis_color_mode: typing.Optional['AxisColorMode']
    axis_label: typing.Optional[str]
    axis_width: typing.Optional[float]
    axis_soft_min: typing.Optional[float]
    axis_soft_max: typing.Optional[float]
    axis_grid_show: typing.Optional[bool]
    scale_distribution: typing.Optional['ScaleDistributionConfig']
    axis_centered_zero: typing.Optional[bool]
    axis_border_show: typing.Optional[bool]

    def __init__(self, axis_placement: typing.Optional['AxisPlacement'] = None, axis_color_mode: typing.Optional['AxisColorMode'] = None, axis_label: typing.Optional[str] = None, axis_width: typing.Optional[float] = None, axis_soft_min: typing.Optional[float] = None, axis_soft_max: typing.Optional[float] = None, axis_grid_show: typing.Optional[bool] = None, scale_distribution: typing.Optional['ScaleDistributionConfig'] = None, axis_centered_zero: typing.Optional[bool] = None, axis_border_show: typing.Optional[bool] = None):
        self.axis_placement = axis_placement
        self.axis_color_mode = axis_color_mode
        self.axis_label = axis_label
        self.axis_width = axis_width
        self.axis_soft_min = axis_soft_min
        self.axis_soft_max = axis_soft_max
        self.axis_grid_show = axis_grid_show
        self.scale_distribution = scale_distribution
        self.axis_centered_zero = axis_centered_zero
        self.axis_border_show = axis_border_show

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.axis_placement is not None:
            payload["axisPlacement"] = self.axis_placement
        if self.axis_color_mode is not None:
            payload["axisColorMode"] = self.axis_color_mode
        if self.axis_label is not None:
            payload["axisLabel"] = self.axis_label
        if self.axis_width is not None:
            payload["axisWidth"] = self.axis_width
        if self.axis_soft_min is not None:
            payload["axisSoftMin"] = self.axis_soft_min
        if self.axis_soft_max is not None:
            payload["axisSoftMax"] = self.axis_soft_max
        if self.axis_grid_show is not None:
            payload["axisGridShow"] = self.axis_grid_show
        if self.scale_distribution is not None:
            payload["scaleDistribution"] = self.scale_distribution
        if self.axis_centered_zero is not None:
            payload["axisCenteredZero"] = self.axis_centered_zero
        if self.axis_border_show is not None:
            payload["axisBorderShow"] = self.axis_border_show
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "axisPlacement" in data:
            args["axis_placement"] = data["axisPlacement"]
        if "axisColorMode" in data:
            args["axis_color_mode"] = data["axisColorMode"]
        if "axisLabel" in data:
            args["axis_label"] = data["axisLabel"]
        if "axisWidth" in data:
            args["axis_width"] = data["axisWidth"]
        if "axisSoftMin" in data:
            args["axis_soft_min"] = data["axisSoftMin"]
        if "axisSoftMax" in data:
            args["axis_soft_max"] = data["axisSoftMax"]
        if "axisGridShow" in data:
            args["axis_grid_show"] = data["axisGridShow"]
        if "scaleDistribution" in data:
            args["scale_distribution"] = ScaleDistributionConfig.from_json(data["scaleDistribution"])
        if "axisCenteredZero" in data:
            args["axis_centered_zero"] = data["axisCenteredZero"]
        if "axisBorderShow" in data:
            args["axis_border_show"] = data["axisBorderShow"]        

        return cls(**args)


class HideSeriesConfig:
    """
    TODO docs
    """

    tooltip: bool
    legend: bool
    viz: bool

    def __init__(self, tooltip: bool = False, legend: bool = False, viz: bool = False):
        self.tooltip = tooltip
        self.legend = legend
        self.viz = viz

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "tooltip": self.tooltip,
            "legend": self.legend,
            "viz": self.viz,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "tooltip" in data:
            args["tooltip"] = data["tooltip"]
        if "legend" in data:
            args["legend"] = data["legend"]
        if "viz" in data:
            args["viz"] = data["viz"]        

        return cls(**args)


class StackingConfig:
    """
    TODO docs
    """

    mode: typing.Optional['StackingMode']
    group: typing.Optional[str]

    def __init__(self, mode: typing.Optional['StackingMode'] = None, group: typing.Optional[str] = None):
        self.mode = mode
        self.group = group

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.mode is not None:
            payload["mode"] = self.mode
        if self.group is not None:
            payload["group"] = self.group
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "mode" in data:
            args["mode"] = data["mode"]
        if "group" in data:
            args["group"] = data["group"]        

        return cls(**args)


class StackableFieldConfig:
    """
    TODO docs
    """

    stacking: typing.Optional['StackingConfig']

    def __init__(self, stacking: typing.Optional['StackingConfig'] = None):
        self.stacking = stacking

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.stacking is not None:
            payload["stacking"] = self.stacking
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "stacking" in data:
            args["stacking"] = StackingConfig.from_json(data["stacking"])        

        return cls(**args)


class HideableFieldConfig:
    """
    TODO docs
    """

    hide_from: typing.Optional['HideSeriesConfig']

    def __init__(self, hide_from: typing.Optional['HideSeriesConfig'] = None):
        self.hide_from = hide_from

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.hide_from is not None:
            payload["hideFrom"] = self.hide_from
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "hideFrom" in data:
            args["hide_from"] = HideSeriesConfig.from_json(data["hideFrom"])        

        return cls(**args)


class GraphTresholdsStyleMode(enum.StrEnum):
    """
    TODO docs
    """

    OFF = "off"
    LINE = "line"
    DASHED = "dashed"
    AREA = "area"
    LINE_AND_AREA = "line+area"
    DASHED_AND_AREA = "dashed+area"
    SERIES = "series"


class GraphThresholdsStyleConfig:
    """
    TODO docs
    """

    mode: 'GraphTresholdsStyleMode'

    def __init__(self, mode: typing.Optional['GraphTresholdsStyleMode'] = None):
        self.mode = mode if mode is not None else GraphTresholdsStyleMode.OFF

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "mode": self.mode,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "mode" in data:
            args["mode"] = data["mode"]        

        return cls(**args)


class LegendPlacement(enum.StrEnum):
    """
    TODO docs
    """

    BOTTOM = "bottom"
    RIGHT = "right"


class LegendDisplayMode(enum.StrEnum):
    """
    TODO docs
    Note: "hidden" needs to remain as an option for plugins compatibility
    """

    LIST = "list"
    TABLE = "table"
    HIDDEN = "hidden"


class SingleStatBaseOptions:
    """
    TODO docs
    """

    reduce_options: 'ReduceDataOptions'
    text: typing.Optional['VizTextDisplayOptions']
    orientation: 'VizOrientation'

    def __init__(self, reduce_options: typing.Optional['ReduceDataOptions'] = None, text: typing.Optional['VizTextDisplayOptions'] = None, orientation: typing.Optional['VizOrientation'] = None):
        self.reduce_options = reduce_options if reduce_options is not None else ReduceDataOptions()
        self.text = text
        self.orientation = orientation if orientation is not None else VizOrientation.AUTO

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "reduceOptions": self.reduce_options,
            "orientation": self.orientation,
        }
        if self.text is not None:
            payload["text"] = self.text
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "reduceOptions" in data:
            args["reduce_options"] = ReduceDataOptions.from_json(data["reduceOptions"])
        if "text" in data:
            args["text"] = VizTextDisplayOptions.from_json(data["text"])
        if "orientation" in data:
            args["orientation"] = data["orientation"]        

        return cls(**args)


class ReduceDataOptions:
    """
    TODO docs
    """

    # If true show each row value
    values: typing.Optional[bool]
    # if showing all values limit
    limit: typing.Optional[float]
    # When !values, pick one value for the whole field
    calcs: list[str]
    # Which fields to show.  By default this is only numeric fields
    fields: typing.Optional[str]

    def __init__(self, values: typing.Optional[bool] = None, limit: typing.Optional[float] = None, calcs: typing.Optional[list[str]] = None, fields: typing.Optional[str] = None):
        self.values = values
        self.limit = limit
        self.calcs = calcs if calcs is not None else []
        self.fields = fields

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "calcs": self.calcs,
        }
        if self.values is not None:
            payload["values"] = self.values
        if self.limit is not None:
            payload["limit"] = self.limit
        if self.fields is not None:
            payload["fields"] = self.fields
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "values" in data:
            args["values"] = data["values"]
        if "limit" in data:
            args["limit"] = data["limit"]
        if "calcs" in data:
            args["calcs"] = data["calcs"]
        if "fields" in data:
            args["fields"] = data["fields"]        

        return cls(**args)


class VizOrientation(enum.StrEnum):
    """
    TODO docs
    """

    AUTO = "auto"
    VERTICAL = "vertical"
    HORIZONTAL = "horizontal"


class OptionsWithTooltip:
    """
    TODO docs
    """

    tooltip: 'VizTooltipOptions'

    def __init__(self, tooltip: typing.Optional['VizTooltipOptions'] = None):
        self.tooltip = tooltip if tooltip is not None else VizTooltipOptions()

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "tooltip": self.tooltip,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "tooltip" in data:
            args["tooltip"] = VizTooltipOptions.from_json(data["tooltip"])        

        return cls(**args)


class OptionsWithLegend:
    """
    TODO docs
    """

    legend: 'VizLegendOptions'

    def __init__(self, legend: typing.Optional['VizLegendOptions'] = None):
        self.legend = legend if legend is not None else VizLegendOptions()

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "legend": self.legend,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "legend" in data:
            args["legend"] = VizLegendOptions.from_json(data["legend"])        

        return cls(**args)


class OptionsWithTimezones:
    """
    TODO docs
    """

    timezone: typing.Optional[list['TimeZone']]

    def __init__(self, timezone: typing.Optional[list['TimeZone']] = None):
        self.timezone = timezone

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.timezone is not None:
            payload["timezone"] = self.timezone
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "timezone" in data:
            args["timezone"] = data["timezone"]        

        return cls(**args)


class OptionsWithTextFormatting:
    """
    TODO docs
    """

    text: typing.Optional['VizTextDisplayOptions']

    def __init__(self, text: typing.Optional['VizTextDisplayOptions'] = None):
        self.text = text

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.text is not None:
            payload["text"] = self.text
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "text" in data:
            args["text"] = VizTextDisplayOptions.from_json(data["text"])        

        return cls(**args)


class BigValueColorMode(enum.StrEnum):
    """
    TODO docs
    """

    VALUE = "value"
    BACKGROUND = "background"
    BACKGROUND_SOLID = "background_solid"
    NONE = "none"


class BigValueGraphMode(enum.StrEnum):
    """
    TODO docs
    """

    NONE = "none"
    LINE = "line"
    AREA = "area"


class BigValueJustifyMode(enum.StrEnum):
    """
    TODO docs
    """

    AUTO = "auto"
    CENTER = "center"


class BigValueTextMode(enum.StrEnum):
    """
    TODO docs
    """

    AUTO = "auto"
    VALUE = "value"
    VALUE_AND_NAME = "value_and_name"
    NAME = "name"
    NONE = "none"


class FieldTextAlignment(enum.StrEnum):
    """
    TODO -- should not be table specific!
    TODO docs
    """

    AUTO = "auto"
    LEFT = "left"
    RIGHT = "right"
    CENTER = "center"


class TimelineValueAlignment(enum.StrEnum):
    """
    Controls the value alignment in the TimelineChart component
    """

    CENTER = "center"
    LEFT = "left"
    RIGHT = "right"


class VizTextDisplayOptions:
    """
    TODO docs
    """

    # Explicit title text size
    title_size: typing.Optional[float]
    # Explicit value text size
    value_size: typing.Optional[float]

    def __init__(self, title_size: typing.Optional[float] = None, value_size: typing.Optional[float] = None):
        self.title_size = title_size
        self.value_size = value_size

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.title_size is not None:
            payload["titleSize"] = self.title_size
        if self.value_size is not None:
            payload["valueSize"] = self.value_size
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "titleSize" in data:
            args["title_size"] = data["titleSize"]
        if "valueSize" in data:
            args["value_size"] = data["valueSize"]        

        return cls(**args)


class TooltipDisplayMode(enum.StrEnum):
    """
    TODO docs
    """

    SINGLE = "single"
    MULTI = "multi"
    NONE = "none"


class SortOrder(enum.StrEnum):
    """
    TODO docs
    """

    ASCENDING = "asc"
    DESCENDING = "desc"
    NONE = "none"


class GraphFieldConfig:
    """
    TODO docs
    """

    draw_style: typing.Optional['GraphDrawStyle']
    gradient_mode: typing.Optional['GraphGradientMode']
    thresholds_style: typing.Optional['GraphThresholdsStyleConfig']
    line_color: typing.Optional[str]
    line_width: typing.Optional[float]
    line_interpolation: typing.Optional['LineInterpolation']
    line_style: typing.Optional['LineStyle']
    fill_color: typing.Optional[str]
    fill_opacity: typing.Optional[float]
    show_points: typing.Optional['VisibilityMode']
    point_size: typing.Optional[float]
    point_color: typing.Optional[str]
    axis_placement: typing.Optional['AxisPlacement']
    axis_color_mode: typing.Optional['AxisColorMode']
    axis_label: typing.Optional[str]
    axis_width: typing.Optional[float]
    axis_soft_min: typing.Optional[float]
    axis_soft_max: typing.Optional[float]
    axis_grid_show: typing.Optional[bool]
    scale_distribution: typing.Optional['ScaleDistributionConfig']
    axis_centered_zero: typing.Optional[bool]
    bar_alignment: typing.Optional['BarAlignment']
    bar_width_factor: typing.Optional[float]
    stacking: typing.Optional['StackingConfig']
    hide_from: typing.Optional['HideSeriesConfig']
    transform: typing.Optional['GraphTransform']
    # Indicate if null values should be treated as gaps or connected.
    # When the value is a number, it represents the maximum delta in the
    # X axis that should be considered connected.  For timeseries, this is milliseconds
    span_nulls: typing.Optional[typing.Union[bool, float]]
    fill_below_to: typing.Optional[str]
    point_symbol: typing.Optional[str]
    axis_border_show: typing.Optional[bool]
    bar_max_width: typing.Optional[float]
    insert_nulls: typing.Optional[typing.Union[bool, int]]

    def __init__(self, draw_style: typing.Optional['GraphDrawStyle'] = None, gradient_mode: typing.Optional['GraphGradientMode'] = None, thresholds_style: typing.Optional['GraphThresholdsStyleConfig'] = None, line_color: typing.Optional[str] = None, line_width: typing.Optional[float] = None, line_interpolation: typing.Optional['LineInterpolation'] = None, line_style: typing.Optional['LineStyle'] = None, fill_color: typing.Optional[str] = None, fill_opacity: typing.Optional[float] = None, show_points: typing.Optional['VisibilityMode'] = None, point_size: typing.Optional[float] = None, point_color: typing.Optional[str] = None, axis_placement: typing.Optional['AxisPlacement'] = None, axis_color_mode: typing.Optional['AxisColorMode'] = None, axis_label: typing.Optional[str] = None, axis_width: typing.Optional[float] = None, axis_soft_min: typing.Optional[float] = None, axis_soft_max: typing.Optional[float] = None, axis_grid_show: typing.Optional[bool] = None, scale_distribution: typing.Optional['ScaleDistributionConfig'] = None, axis_centered_zero: typing.Optional[bool] = None, bar_alignment: typing.Optional['BarAlignment'] = None, bar_width_factor: typing.Optional[float] = None, stacking: typing.Optional['StackingConfig'] = None, hide_from: typing.Optional['HideSeriesConfig'] = None, transform: typing.Optional['GraphTransform'] = None, span_nulls: typing.Optional[typing.Union[bool, float]] = None, fill_below_to: typing.Optional[str] = None, point_symbol: typing.Optional[str] = None, axis_border_show: typing.Optional[bool] = None, bar_max_width: typing.Optional[float] = None, insert_nulls: typing.Optional[typing.Union[bool, int]] = None):
        self.draw_style = draw_style
        self.gradient_mode = gradient_mode
        self.thresholds_style = thresholds_style
        self.line_color = line_color
        self.line_width = line_width
        self.line_interpolation = line_interpolation
        self.line_style = line_style
        self.fill_color = fill_color
        self.fill_opacity = fill_opacity
        self.show_points = show_points
        self.point_size = point_size
        self.point_color = point_color
        self.axis_placement = axis_placement
        self.axis_color_mode = axis_color_mode
        self.axis_label = axis_label
        self.axis_width = axis_width
        self.axis_soft_min = axis_soft_min
        self.axis_soft_max = axis_soft_max
        self.axis_grid_show = axis_grid_show
        self.scale_distribution = scale_distribution
        self.axis_centered_zero = axis_centered_zero
        self.bar_alignment = bar_alignment
        self.bar_width_factor = bar_width_factor
        self.stacking = stacking
        self.hide_from = hide_from
        self.transform = transform
        self.span_nulls = span_nulls
        self.fill_below_to = fill_below_to
        self.point_symbol = point_symbol
        self.axis_border_show = axis_border_show
        self.bar_max_width = bar_max_width
        self.insert_nulls = insert_nulls

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.draw_style is not None:
            payload["drawStyle"] = self.draw_style
        if self.gradient_mode is not None:
            payload["gradientMode"] = self.gradient_mode
        if self.thresholds_style is not None:
            payload["thresholdsStyle"] = self.thresholds_style
        if self.line_color is not None:
            payload["lineColor"] = self.line_color
        if self.line_width is not None:
            payload["lineWidth"] = self.line_width
        if self.line_interpolation is not None:
            payload["lineInterpolation"] = self.line_interpolation
        if self.line_style is not None:
            payload["lineStyle"] = self.line_style
        if self.fill_color is not None:
            payload["fillColor"] = self.fill_color
        if self.fill_opacity is not None:
            payload["fillOpacity"] = self.fill_opacity
        if self.show_points is not None:
            payload["showPoints"] = self.show_points
        if self.point_size is not None:
            payload["pointSize"] = self.point_size
        if self.point_color is not None:
            payload["pointColor"] = self.point_color
        if self.axis_placement is not None:
            payload["axisPlacement"] = self.axis_placement
        if self.axis_color_mode is not None:
            payload["axisColorMode"] = self.axis_color_mode
        if self.axis_label is not None:
            payload["axisLabel"] = self.axis_label
        if self.axis_width is not None:
            payload["axisWidth"] = self.axis_width
        if self.axis_soft_min is not None:
            payload["axisSoftMin"] = self.axis_soft_min
        if self.axis_soft_max is not None:
            payload["axisSoftMax"] = self.axis_soft_max
        if self.axis_grid_show is not None:
            payload["axisGridShow"] = self.axis_grid_show
        if self.scale_distribution is not None:
            payload["scaleDistribution"] = self.scale_distribution
        if self.axis_centered_zero is not None:
            payload["axisCenteredZero"] = self.axis_centered_zero
        if self.bar_alignment is not None:
            payload["barAlignment"] = self.bar_alignment
        if self.bar_width_factor is not None:
            payload["barWidthFactor"] = self.bar_width_factor
        if self.stacking is not None:
            payload["stacking"] = self.stacking
        if self.hide_from is not None:
            payload["hideFrom"] = self.hide_from
        if self.transform is not None:
            payload["transform"] = self.transform
        if self.span_nulls is not None:
            payload["spanNulls"] = self.span_nulls
        if self.fill_below_to is not None:
            payload["fillBelowTo"] = self.fill_below_to
        if self.point_symbol is not None:
            payload["pointSymbol"] = self.point_symbol
        if self.axis_border_show is not None:
            payload["axisBorderShow"] = self.axis_border_show
        if self.bar_max_width is not None:
            payload["barMaxWidth"] = self.bar_max_width
        if self.insert_nulls is not None:
            payload["insertNulls"] = self.insert_nulls
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "drawStyle" in data:
            args["draw_style"] = data["drawStyle"]
        if "gradientMode" in data:
            args["gradient_mode"] = data["gradientMode"]
        if "thresholdsStyle" in data:
            args["thresholds_style"] = GraphThresholdsStyleConfig.from_json(data["thresholdsStyle"])
        if "lineColor" in data:
            args["line_color"] = data["lineColor"]
        if "lineWidth" in data:
            args["line_width"] = data["lineWidth"]
        if "lineInterpolation" in data:
            args["line_interpolation"] = data["lineInterpolation"]
        if "lineStyle" in data:
            args["line_style"] = LineStyle.from_json(data["lineStyle"])
        if "fillColor" in data:
            args["fill_color"] = data["fillColor"]
        if "fillOpacity" in data:
            args["fill_opacity"] = data["fillOpacity"]
        if "showPoints" in data:
            args["show_points"] = data["showPoints"]
        if "pointSize" in data:
            args["point_size"] = data["pointSize"]
        if "pointColor" in data:
            args["point_color"] = data["pointColor"]
        if "axisPlacement" in data:
            args["axis_placement"] = data["axisPlacement"]
        if "axisColorMode" in data:
            args["axis_color_mode"] = data["axisColorMode"]
        if "axisLabel" in data:
            args["axis_label"] = data["axisLabel"]
        if "axisWidth" in data:
            args["axis_width"] = data["axisWidth"]
        if "axisSoftMin" in data:
            args["axis_soft_min"] = data["axisSoftMin"]
        if "axisSoftMax" in data:
            args["axis_soft_max"] = data["axisSoftMax"]
        if "axisGridShow" in data:
            args["axis_grid_show"] = data["axisGridShow"]
        if "scaleDistribution" in data:
            args["scale_distribution"] = ScaleDistributionConfig.from_json(data["scaleDistribution"])
        if "axisCenteredZero" in data:
            args["axis_centered_zero"] = data["axisCenteredZero"]
        if "barAlignment" in data:
            args["bar_alignment"] = data["barAlignment"]
        if "barWidthFactor" in data:
            args["bar_width_factor"] = data["barWidthFactor"]
        if "stacking" in data:
            args["stacking"] = StackingConfig.from_json(data["stacking"])
        if "hideFrom" in data:
            args["hide_from"] = HideSeriesConfig.from_json(data["hideFrom"])
        if "transform" in data:
            args["transform"] = data["transform"]
        if "spanNulls" in data:
            args["span_nulls"] = data["spanNulls"]
        if "fillBelowTo" in data:
            args["fill_below_to"] = data["fillBelowTo"]
        if "pointSymbol" in data:
            args["point_symbol"] = data["pointSymbol"]
        if "axisBorderShow" in data:
            args["axis_border_show"] = data["axisBorderShow"]
        if "barMaxWidth" in data:
            args["bar_max_width"] = data["barMaxWidth"]
        if "insertNulls" in data:
            args["insert_nulls"] = data["insertNulls"]        

        return cls(**args)


class VizLegendOptions:
    """
    TODO docs
    """

    display_mode: 'LegendDisplayMode'
    placement: 'LegendPlacement'
    show_legend: bool
    as_table: typing.Optional[bool]
    is_visible: typing.Optional[bool]
    sort_by: typing.Optional[str]
    sort_desc: typing.Optional[bool]
    width: typing.Optional[float]
    calcs: list[str]

    def __init__(self, display_mode: typing.Optional['LegendDisplayMode'] = None, placement: typing.Optional['LegendPlacement'] = None, show_legend: bool = False, as_table: typing.Optional[bool] = None, is_visible: typing.Optional[bool] = None, sort_by: typing.Optional[str] = None, sort_desc: typing.Optional[bool] = None, width: typing.Optional[float] = None, calcs: typing.Optional[list[str]] = None):
        self.display_mode = display_mode if display_mode is not None else LegendDisplayMode.LIST
        self.placement = placement if placement is not None else LegendPlacement.BOTTOM
        self.show_legend = show_legend
        self.as_table = as_table
        self.is_visible = is_visible
        self.sort_by = sort_by
        self.sort_desc = sort_desc
        self.width = width
        self.calcs = calcs if calcs is not None else []

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "displayMode": self.display_mode,
            "placement": self.placement,
            "showLegend": self.show_legend,
            "calcs": self.calcs,
        }
        if self.as_table is not None:
            payload["asTable"] = self.as_table
        if self.is_visible is not None:
            payload["isVisible"] = self.is_visible
        if self.sort_by is not None:
            payload["sortBy"] = self.sort_by
        if self.sort_desc is not None:
            payload["sortDesc"] = self.sort_desc
        if self.width is not None:
            payload["width"] = self.width
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "displayMode" in data:
            args["display_mode"] = data["displayMode"]
        if "placement" in data:
            args["placement"] = data["placement"]
        if "showLegend" in data:
            args["show_legend"] = data["showLegend"]
        if "asTable" in data:
            args["as_table"] = data["asTable"]
        if "isVisible" in data:
            args["is_visible"] = data["isVisible"]
        if "sortBy" in data:
            args["sort_by"] = data["sortBy"]
        if "sortDesc" in data:
            args["sort_desc"] = data["sortDesc"]
        if "width" in data:
            args["width"] = data["width"]
        if "calcs" in data:
            args["calcs"] = data["calcs"]        

        return cls(**args)


class BarGaugeDisplayMode(enum.StrEnum):
    """
    Enum expressing the possible display modes
    for the bar gauge component of Grafana UI
    """

    BASIC = "basic"
    LCD = "lcd"
    GRADIENT = "gradient"


class BarGaugeValueMode(enum.StrEnum):
    """
    Allows for the table cell gauge display type to set the gauge mode.
    """

    COLOR = "color"
    TEXT = "text"
    HIDDEN = "hidden"


class BarGaugeNamePlacement(enum.StrEnum):
    """
    Allows for the bar gauge name to be placed explicitly
    """

    AUTO = "auto"
    TOP = "top"
    LEFT = "left"


class BarGaugeSizing(enum.StrEnum):
    """
    Allows for the bar gauge size to be set explicitly
    """

    AUTO = "auto"
    MANUAL = "manual"


class VizTooltipOptions:
    """
    TODO docs
    """

    mode: 'TooltipDisplayMode'
    sort: 'SortOrder'

    def __init__(self, mode: typing.Optional['TooltipDisplayMode'] = None, sort: typing.Optional['SortOrder'] = None):
        self.mode = mode if mode is not None else TooltipDisplayMode.SINGLE
        self.sort = sort if sort is not None else SortOrder.ASCENDING

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "mode": self.mode,
            "sort": self.sort,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "mode" in data:
            args["mode"] = data["mode"]
        if "sort" in data:
            args["sort"] = data["sort"]        

        return cls(**args)


Labels: typing.TypeAlias = dict[str, str]


class TableCellDisplayMode(enum.StrEnum):
    """
    Internally, this is the "type" of cell that's being displayed
    in the table such as colored text, JSON, gauge, etc.
    The color-background-solid, gradient-gauge, and lcd-gauge
    modes are deprecated in favor of new cell subOptions
    """

    AUTO = "auto"
    COLOR_TEXT = "color-text"
    COLOR_BACKGROUND = "color-background"
    COLOR_BACKGROUND_SOLID = "color-background-solid"
    GRADIENT_GAUGE = "gradient-gauge"
    LCD_GAUGE = "lcd-gauge"
    JSON_VIEW = "json-view"
    BASIC_GAUGE = "basic"
    IMAGE = "image"
    GAUGE = "gauge"
    SPARKLINE = "sparkline"
    CUSTOM = "custom"


class TableCellBackgroundDisplayMode(enum.StrEnum):
    """
    Display mode to the "Colored Background" display
    mode for table cells. Either displays a solid color (basic mode)
    or a gradient.
    """

    BASIC = "basic"
    GRADIENT = "gradient"


class TableSortByFieldState:
    """
    Sort by field state
    """

    # Sets the display name of the field to sort by
    display_name: str
    # Flag used to indicate descending sort order
    desc: typing.Optional[bool]

    def __init__(self, display_name: str = "", desc: typing.Optional[bool] = None):
        self.display_name = display_name
        self.desc = desc

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "displayName": self.display_name,
        }
        if self.desc is not None:
            payload["desc"] = self.desc
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "displayName" in data:
            args["display_name"] = data["displayName"]
        if "desc" in data:
            args["desc"] = data["desc"]        

        return cls(**args)


class TableFooterOptions:
    """
    Footer options
    """

    show: bool
    # actually 1 value
    reducer: list[str]
    fields: typing.Optional[list[str]]
    enable_pagination: typing.Optional[bool]
    count_rows: typing.Optional[bool]

    def __init__(self, show: bool = False, reducer: typing.Optional[list[str]] = None, fields: typing.Optional[list[str]] = None, enable_pagination: typing.Optional[bool] = None, count_rows: typing.Optional[bool] = None):
        self.show = show
        self.reducer = reducer if reducer is not None else []
        self.fields = fields
        self.enable_pagination = enable_pagination
        self.count_rows = count_rows

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "show": self.show,
            "reducer": self.reducer,
        }
        if self.fields is not None:
            payload["fields"] = self.fields
        if self.enable_pagination is not None:
            payload["enablePagination"] = self.enable_pagination
        if self.count_rows is not None:
            payload["countRows"] = self.count_rows
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "show" in data:
            args["show"] = data["show"]
        if "reducer" in data:
            args["reducer"] = data["reducer"]
        if "fields" in data:
            args["fields"] = data["fields"]
        if "enablePagination" in data:
            args["enable_pagination"] = data["enablePagination"]
        if "countRows" in data:
            args["count_rows"] = data["countRows"]        

        return cls(**args)


class TableAutoCellOptions:
    """
    Auto mode table cell options
    """

    type_val: typing.Literal["auto"]

    def __init__(self, ):
        self.type_val = "auto"

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "type": self.type_val,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        return cls(**args)


class TableColorTextCellOptions:
    """
    Colored text cell options
    """

    type_val: typing.Literal["color-text"]

    def __init__(self, ):
        self.type_val = "color-text"

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "type": self.type_val,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        return cls(**args)


class TableJsonViewCellOptions:
    """
    Json view cell options
    """

    type_val: typing.Literal["json-view"]

    def __init__(self, ):
        self.type_val = "json-view"

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "type": self.type_val,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        return cls(**args)


class TableImageCellOptions:
    """
    Json view cell options
    """

    type_val: typing.Literal["image"]

    def __init__(self, ):
        self.type_val = "image"

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "type": self.type_val,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        return cls(**args)


class TableBarGaugeCellOptions:
    """
    Gauge cell options
    """

    type_val: typing.Literal["gauge"]
    mode: typing.Optional['BarGaugeDisplayMode']
    value_display_mode: typing.Optional['BarGaugeValueMode']

    def __init__(self, mode: typing.Optional['BarGaugeDisplayMode'] = None, value_display_mode: typing.Optional['BarGaugeValueMode'] = None):
        self.type_val = "gauge"
        self.mode = mode
        self.value_display_mode = value_display_mode

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "type": self.type_val,
        }
        if self.mode is not None:
            payload["mode"] = self.mode
        if self.value_display_mode is not None:
            payload["valueDisplayMode"] = self.value_display_mode
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "mode" in data:
            args["mode"] = data["mode"]
        if "valueDisplayMode" in data:
            args["value_display_mode"] = data["valueDisplayMode"]        

        return cls(**args)


class TableSparklineCellOptions:
    """
    Sparkline cell options
    """

    type_val: typing.Literal["sparkline"]
    draw_style: typing.Optional['GraphDrawStyle']
    gradient_mode: typing.Optional['GraphGradientMode']
    thresholds_style: typing.Optional['GraphThresholdsStyleConfig']
    line_color: typing.Optional[str]
    line_width: typing.Optional[float]
    line_interpolation: typing.Optional['LineInterpolation']
    line_style: typing.Optional['LineStyle']
    fill_color: typing.Optional[str]
    fill_opacity: typing.Optional[float]
    show_points: typing.Optional['VisibilityMode']
    point_size: typing.Optional[float]
    point_color: typing.Optional[str]
    axis_placement: typing.Optional['AxisPlacement']
    axis_color_mode: typing.Optional['AxisColorMode']
    axis_label: typing.Optional[str]
    axis_width: typing.Optional[float]
    axis_soft_min: typing.Optional[float]
    axis_soft_max: typing.Optional[float]
    axis_grid_show: typing.Optional[bool]
    scale_distribution: typing.Optional['ScaleDistributionConfig']
    axis_centered_zero: typing.Optional[bool]
    bar_alignment: typing.Optional['BarAlignment']
    bar_width_factor: typing.Optional[float]
    stacking: typing.Optional['StackingConfig']
    hide_from: typing.Optional['HideSeriesConfig']
    hide_value: typing.Optional[bool]
    transform: typing.Optional['GraphTransform']
    # Indicate if null values should be treated as gaps or connected.
    # When the value is a number, it represents the maximum delta in the
    # X axis that should be considered connected.  For timeseries, this is milliseconds
    span_nulls: typing.Optional[typing.Union[bool, float]]
    fill_below_to: typing.Optional[str]
    point_symbol: typing.Optional[str]
    axis_border_show: typing.Optional[bool]
    bar_max_width: typing.Optional[float]

    def __init__(self, draw_style: typing.Optional['GraphDrawStyle'] = None, gradient_mode: typing.Optional['GraphGradientMode'] = None, thresholds_style: typing.Optional['GraphThresholdsStyleConfig'] = None, line_color: typing.Optional[str] = None, line_width: typing.Optional[float] = None, line_interpolation: typing.Optional['LineInterpolation'] = None, line_style: typing.Optional['LineStyle'] = None, fill_color: typing.Optional[str] = None, fill_opacity: typing.Optional[float] = None, show_points: typing.Optional['VisibilityMode'] = None, point_size: typing.Optional[float] = None, point_color: typing.Optional[str] = None, axis_placement: typing.Optional['AxisPlacement'] = None, axis_color_mode: typing.Optional['AxisColorMode'] = None, axis_label: typing.Optional[str] = None, axis_width: typing.Optional[float] = None, axis_soft_min: typing.Optional[float] = None, axis_soft_max: typing.Optional[float] = None, axis_grid_show: typing.Optional[bool] = None, scale_distribution: typing.Optional['ScaleDistributionConfig'] = None, axis_centered_zero: typing.Optional[bool] = None, bar_alignment: typing.Optional['BarAlignment'] = None, bar_width_factor: typing.Optional[float] = None, stacking: typing.Optional['StackingConfig'] = None, hide_from: typing.Optional['HideSeriesConfig'] = None, hide_value: typing.Optional[bool] = None, transform: typing.Optional['GraphTransform'] = None, span_nulls: typing.Optional[typing.Union[bool, float]] = None, fill_below_to: typing.Optional[str] = None, point_symbol: typing.Optional[str] = None, axis_border_show: typing.Optional[bool] = None, bar_max_width: typing.Optional[float] = None):
        self.type_val = "sparkline"
        self.draw_style = draw_style
        self.gradient_mode = gradient_mode
        self.thresholds_style = thresholds_style
        self.line_color = line_color
        self.line_width = line_width
        self.line_interpolation = line_interpolation
        self.line_style = line_style
        self.fill_color = fill_color
        self.fill_opacity = fill_opacity
        self.show_points = show_points
        self.point_size = point_size
        self.point_color = point_color
        self.axis_placement = axis_placement
        self.axis_color_mode = axis_color_mode
        self.axis_label = axis_label
        self.axis_width = axis_width
        self.axis_soft_min = axis_soft_min
        self.axis_soft_max = axis_soft_max
        self.axis_grid_show = axis_grid_show
        self.scale_distribution = scale_distribution
        self.axis_centered_zero = axis_centered_zero
        self.bar_alignment = bar_alignment
        self.bar_width_factor = bar_width_factor
        self.stacking = stacking
        self.hide_from = hide_from
        self.hide_value = hide_value
        self.transform = transform
        self.span_nulls = span_nulls
        self.fill_below_to = fill_below_to
        self.point_symbol = point_symbol
        self.axis_border_show = axis_border_show
        self.bar_max_width = bar_max_width

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "type": self.type_val,
        }
        if self.draw_style is not None:
            payload["drawStyle"] = self.draw_style
        if self.gradient_mode is not None:
            payload["gradientMode"] = self.gradient_mode
        if self.thresholds_style is not None:
            payload["thresholdsStyle"] = self.thresholds_style
        if self.line_color is not None:
            payload["lineColor"] = self.line_color
        if self.line_width is not None:
            payload["lineWidth"] = self.line_width
        if self.line_interpolation is not None:
            payload["lineInterpolation"] = self.line_interpolation
        if self.line_style is not None:
            payload["lineStyle"] = self.line_style
        if self.fill_color is not None:
            payload["fillColor"] = self.fill_color
        if self.fill_opacity is not None:
            payload["fillOpacity"] = self.fill_opacity
        if self.show_points is not None:
            payload["showPoints"] = self.show_points
        if self.point_size is not None:
            payload["pointSize"] = self.point_size
        if self.point_color is not None:
            payload["pointColor"] = self.point_color
        if self.axis_placement is not None:
            payload["axisPlacement"] = self.axis_placement
        if self.axis_color_mode is not None:
            payload["axisColorMode"] = self.axis_color_mode
        if self.axis_label is not None:
            payload["axisLabel"] = self.axis_label
        if self.axis_width is not None:
            payload["axisWidth"] = self.axis_width
        if self.axis_soft_min is not None:
            payload["axisSoftMin"] = self.axis_soft_min
        if self.axis_soft_max is not None:
            payload["axisSoftMax"] = self.axis_soft_max
        if self.axis_grid_show is not None:
            payload["axisGridShow"] = self.axis_grid_show
        if self.scale_distribution is not None:
            payload["scaleDistribution"] = self.scale_distribution
        if self.axis_centered_zero is not None:
            payload["axisCenteredZero"] = self.axis_centered_zero
        if self.bar_alignment is not None:
            payload["barAlignment"] = self.bar_alignment
        if self.bar_width_factor is not None:
            payload["barWidthFactor"] = self.bar_width_factor
        if self.stacking is not None:
            payload["stacking"] = self.stacking
        if self.hide_from is not None:
            payload["hideFrom"] = self.hide_from
        if self.hide_value is not None:
            payload["hideValue"] = self.hide_value
        if self.transform is not None:
            payload["transform"] = self.transform
        if self.span_nulls is not None:
            payload["spanNulls"] = self.span_nulls
        if self.fill_below_to is not None:
            payload["fillBelowTo"] = self.fill_below_to
        if self.point_symbol is not None:
            payload["pointSymbol"] = self.point_symbol
        if self.axis_border_show is not None:
            payload["axisBorderShow"] = self.axis_border_show
        if self.bar_max_width is not None:
            payload["barMaxWidth"] = self.bar_max_width
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "drawStyle" in data:
            args["draw_style"] = data["drawStyle"]
        if "gradientMode" in data:
            args["gradient_mode"] = data["gradientMode"]
        if "thresholdsStyle" in data:
            args["thresholds_style"] = GraphThresholdsStyleConfig.from_json(data["thresholdsStyle"])
        if "lineColor" in data:
            args["line_color"] = data["lineColor"]
        if "lineWidth" in data:
            args["line_width"] = data["lineWidth"]
        if "lineInterpolation" in data:
            args["line_interpolation"] = data["lineInterpolation"]
        if "lineStyle" in data:
            args["line_style"] = LineStyle.from_json(data["lineStyle"])
        if "fillColor" in data:
            args["fill_color"] = data["fillColor"]
        if "fillOpacity" in data:
            args["fill_opacity"] = data["fillOpacity"]
        if "showPoints" in data:
            args["show_points"] = data["showPoints"]
        if "pointSize" in data:
            args["point_size"] = data["pointSize"]
        if "pointColor" in data:
            args["point_color"] = data["pointColor"]
        if "axisPlacement" in data:
            args["axis_placement"] = data["axisPlacement"]
        if "axisColorMode" in data:
            args["axis_color_mode"] = data["axisColorMode"]
        if "axisLabel" in data:
            args["axis_label"] = data["axisLabel"]
        if "axisWidth" in data:
            args["axis_width"] = data["axisWidth"]
        if "axisSoftMin" in data:
            args["axis_soft_min"] = data["axisSoftMin"]
        if "axisSoftMax" in data:
            args["axis_soft_max"] = data["axisSoftMax"]
        if "axisGridShow" in data:
            args["axis_grid_show"] = data["axisGridShow"]
        if "scaleDistribution" in data:
            args["scale_distribution"] = ScaleDistributionConfig.from_json(data["scaleDistribution"])
        if "axisCenteredZero" in data:
            args["axis_centered_zero"] = data["axisCenteredZero"]
        if "barAlignment" in data:
            args["bar_alignment"] = data["barAlignment"]
        if "barWidthFactor" in data:
            args["bar_width_factor"] = data["barWidthFactor"]
        if "stacking" in data:
            args["stacking"] = StackingConfig.from_json(data["stacking"])
        if "hideFrom" in data:
            args["hide_from"] = HideSeriesConfig.from_json(data["hideFrom"])
        if "hideValue" in data:
            args["hide_value"] = data["hideValue"]
        if "transform" in data:
            args["transform"] = data["transform"]
        if "spanNulls" in data:
            args["span_nulls"] = data["spanNulls"]
        if "fillBelowTo" in data:
            args["fill_below_to"] = data["fillBelowTo"]
        if "pointSymbol" in data:
            args["point_symbol"] = data["pointSymbol"]
        if "axisBorderShow" in data:
            args["axis_border_show"] = data["axisBorderShow"]
        if "barMaxWidth" in data:
            args["bar_max_width"] = data["barMaxWidth"]        

        return cls(**args)


class TableColoredBackgroundCellOptions:
    """
    Colored background cell options
    """

    type_val: typing.Literal["color-background"]
    mode: typing.Optional['TableCellBackgroundDisplayMode']

    def __init__(self, mode: typing.Optional['TableCellBackgroundDisplayMode'] = None):
        self.type_val = "color-background"
        self.mode = mode

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "type": self.type_val,
        }
        if self.mode is not None:
            payload["mode"] = self.mode
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "mode" in data:
            args["mode"] = data["mode"]        

        return cls(**args)


class TableCellHeight(enum.StrEnum):
    """
    Height of a table cell
    """

    SM = "sm"
    MD = "md"
    LG = "lg"


# Table cell options. Each cell has a display mode
# and other potential options for that display.
TableCellOptions: typing.TypeAlias = typing.Union['TableAutoCellOptions', 'TableSparklineCellOptions', 'TableBarGaugeCellOptions', 'TableColoredBackgroundCellOptions', 'TableColorTextCellOptions', 'TableImageCellOptions', 'TableJsonViewCellOptions']


# Use UTC/GMT timezone
TimeZoneUtc: typing.Literal["utc"] = "utc"


# Use the timezone defined by end user web browser
TimeZoneBrowser: typing.Literal["browser"] = "browser"


class VariableFormatID(enum.StrEnum):
    """
    Optional formats for the template variable replace functions
    See also https://grafana.com/docs/grafana/latest/dashboards/variables/variable-syntax/#advanced-variable-format-options
    """

    LUCENE = "lucene"
    RAW = "raw"
    REGEX = "regex"
    PIPE = "pipe"
    DISTRIBUTED = "distributed"
    CSV = "csv"
    HTML = "html"
    JSON = "json"
    PERCENT_ENCODE = "percentencode"
    URI_ENCODE = "uriencode"
    SINGLE_QUOTE = "singlequote"
    DOUBLE_QUOTE = "doublequote"
    SQL_STRING = "sqlstring"
    DATE = "date"
    GLOB = "glob"
    TEXT = "text"
    QUERY_PARAM = "queryparam"


class DataSourceRef:
    # The plugin type-id
    type_val: typing.Optional[str]
    # Specific datasource instance
    uid: typing.Optional[str]

    def __init__(self, type_val: typing.Optional[str] = None, uid: typing.Optional[str] = None):
        self.type_val = type_val
        self.uid = uid

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.type_val is not None:
            payload["type"] = self.type_val
        if self.uid is not None:
            payload["uid"] = self.uid
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "type" in data:
            args["type_val"] = data["type"]
        if "uid" in data:
            args["uid"] = data["uid"]        

        return cls(**args)


class ResourceDimensionConfig:
    """
    Links to a resource (image/svg path)
    """

    mode: 'ResourceDimensionMode'
    # fixed: T -- will be added by each element
    field: typing.Optional[str]
    fixed: typing.Optional[str]

    def __init__(self, mode: typing.Optional['ResourceDimensionMode'] = None, field: typing.Optional[str] = None, fixed: typing.Optional[str] = None):
        self.mode = mode if mode is not None else ResourceDimensionMode.FIXED
        self.field = field
        self.fixed = fixed

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "mode": self.mode,
        }
        if self.field is not None:
            payload["field"] = self.field
        if self.fixed is not None:
            payload["fixed"] = self.fixed
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "mode" in data:
            args["mode"] = data["mode"]
        if "field" in data:
            args["field"] = data["field"]
        if "fixed" in data:
            args["fixed"] = data["fixed"]        

        return cls(**args)


class FrameGeometrySource:
    mode: 'FrameGeometrySourceMode'
    # Field mappings
    geohash: typing.Optional[str]
    latitude: typing.Optional[str]
    longitude: typing.Optional[str]
    wkt: typing.Optional[str]
    lookup: typing.Optional[str]
    # Path to Gazetteer
    gazetteer: typing.Optional[str]

    def __init__(self, mode: typing.Optional['FrameGeometrySourceMode'] = None, geohash: typing.Optional[str] = None, latitude: typing.Optional[str] = None, longitude: typing.Optional[str] = None, wkt: typing.Optional[str] = None, lookup: typing.Optional[str] = None, gazetteer: typing.Optional[str] = None):
        self.mode = mode if mode is not None else FrameGeometrySourceMode.AUTO
        self.geohash = geohash
        self.latitude = latitude
        self.longitude = longitude
        self.wkt = wkt
        self.lookup = lookup
        self.gazetteer = gazetteer

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "mode": self.mode,
        }
        if self.geohash is not None:
            payload["geohash"] = self.geohash
        if self.latitude is not None:
            payload["latitude"] = self.latitude
        if self.longitude is not None:
            payload["longitude"] = self.longitude
        if self.wkt is not None:
            payload["wkt"] = self.wkt
        if self.lookup is not None:
            payload["lookup"] = self.lookup
        if self.gazetteer is not None:
            payload["gazetteer"] = self.gazetteer
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "mode" in data:
            args["mode"] = data["mode"]
        if "geohash" in data:
            args["geohash"] = data["geohash"]
        if "latitude" in data:
            args["latitude"] = data["latitude"]
        if "longitude" in data:
            args["longitude"] = data["longitude"]
        if "wkt" in data:
            args["wkt"] = data["wkt"]
        if "lookup" in data:
            args["lookup"] = data["lookup"]
        if "gazetteer" in data:
            args["gazetteer"] = data["gazetteer"]        

        return cls(**args)


class HeatmapCalculationOptions:
    # The number of buckets to use for the xAxis in the heatmap
    x_buckets: typing.Optional['HeatmapCalculationBucketConfig']
    # The number of buckets to use for the yAxis in the heatmap
    y_buckets: typing.Optional['HeatmapCalculationBucketConfig']

    def __init__(self, x_buckets: typing.Optional['HeatmapCalculationBucketConfig'] = None, y_buckets: typing.Optional['HeatmapCalculationBucketConfig'] = None):
        self.x_buckets = x_buckets
        self.y_buckets = y_buckets

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.x_buckets is not None:
            payload["xBuckets"] = self.x_buckets
        if self.y_buckets is not None:
            payload["yBuckets"] = self.y_buckets
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "xBuckets" in data:
            args["x_buckets"] = HeatmapCalculationBucketConfig.from_json(data["xBuckets"])
        if "yBuckets" in data:
            args["y_buckets"] = HeatmapCalculationBucketConfig.from_json(data["yBuckets"])        

        return cls(**args)


class LogsDedupStrategy(enum.StrEnum):
    NONE = "none"
    EXACT = "exact"
    NUMBERS = "numbers"
    SIGNATURE = "signature"


class ComparisonOperation(enum.StrEnum):
    """
    Compare two values
    """

    EQ = "eq"
    NEQ = "neq"
    LT = "lt"
    LTE = "lte"
    GT = "gt"
    GTE = "gte"


class TableFieldOptions:
    """
    Field options for each field within a table (e.g 10, "The String", 64.20, etc.)
    Generally defines alignment, filtering capabilties, display options, etc.
    """

    width: typing.Optional[float]
    min_width: typing.Optional[float]
    align: 'FieldTextAlignment'
    # This field is deprecated in favor of using cellOptions
    display_mode: typing.Optional['TableCellDisplayMode']
    cell_options: typing.Optional['TableCellOptions']
    # ?? default is missing or false ??
    hidden: typing.Optional[bool]
    inspect: bool
    filterable: typing.Optional[bool]
    # Hides any header for a column, useful for columns that show some static content or buttons.
    hide_header: typing.Optional[bool]

    def __init__(self, width: typing.Optional[float] = None, min_width: typing.Optional[float] = None, align: typing.Optional['FieldTextAlignment'] = None, display_mode: typing.Optional['TableCellDisplayMode'] = None, cell_options: typing.Optional['TableCellOptions'] = None, hidden: typing.Optional[bool] = None, inspect: bool = False, filterable: typing.Optional[bool] = None, hide_header: typing.Optional[bool] = None):
        self.width = width
        self.min_width = min_width
        self.align = align if align is not None else FieldTextAlignment.AUTO
        self.display_mode = display_mode
        self.cell_options = cell_options
        self.hidden = hidden
        self.inspect = inspect
        self.filterable = filterable
        self.hide_header = hide_header

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "align": self.align,
            "inspect": self.inspect,
        }
        if self.width is not None:
            payload["width"] = self.width
        if self.min_width is not None:
            payload["minWidth"] = self.min_width
        if self.display_mode is not None:
            payload["displayMode"] = self.display_mode
        if self.cell_options is not None:
            payload["cellOptions"] = self.cell_options
        if self.hidden is not None:
            payload["hidden"] = self.hidden
        if self.filterable is not None:
            payload["filterable"] = self.filterable
        if self.hide_header is not None:
            payload["hideHeader"] = self.hide_header
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "width" in data:
            args["width"] = data["width"]
        if "minWidth" in data:
            args["min_width"] = data["minWidth"]
        if "align" in data:
            args["align"] = data["align"]
        if "displayMode" in data:
            args["display_mode"] = data["displayMode"]
        if "cellOptions" in data:
            args["cell_options"] = data["cellOptions"]
        if "hidden" in data:
            args["hidden"] = data["hidden"]
        if "inspect" in data:
            args["inspect"] = data["inspect"]
        if "filterable" in data:
            args["filterable"] = data["filterable"]
        if "hideHeader" in data:
            args["hide_header"] = data["hideHeader"]        

        return cls(**args)


# A specific timezone from https://en.wikipedia.org/wiki/Tz_database
TimeZone: typing.TypeAlias = typing.Union[typing.Literal["utc"], typing.Literal["browser"], str]



