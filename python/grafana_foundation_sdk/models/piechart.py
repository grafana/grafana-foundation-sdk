# Code generated - EDITING IS FUTILE. DO NOT EDIT.

import enum
from ..models import common
import typing
from ..cog import runtime as cogruntime


class PieChartType(enum.StrEnum):
    """
    Select the pie chart display style.
    """

    PIE = "pie"
    DONUT = "donut"


class PieChartLabels(enum.StrEnum):
    """
    Select labels to display on the pie chart.
     - Name - The series or field name.
     - Percent - The percentage of the whole.
     - Value - The raw numerical value.
    """

    NAME = "name"
    VALUE = "value"
    PERCENT = "percent"


class PieChartLegendValues(enum.StrEnum):
    """
    Select values to display in the legend.
     - Percent: The percentage of the whole.
     - Value: The raw numerical value.
    """

    VALUE = "value"
    PERCENT = "percent"


class PieChartLegendOptions:
    values: list['PieChartLegendValues']
    display_mode: common.LegendDisplayMode
    placement: common.LegendPlacement
    show_legend: bool
    as_table: typing.Optional[bool]
    is_visible: typing.Optional[bool]
    sort_by: typing.Optional[str]
    sort_desc: typing.Optional[bool]
    width: typing.Optional[float]
    calcs: list[str]

    def __init__(self, values: typing.Optional[list['PieChartLegendValues']] = None, display_mode: typing.Optional[common.LegendDisplayMode] = None, placement: typing.Optional[common.LegendPlacement] = None, show_legend: bool = False, as_table: typing.Optional[bool] = None, is_visible: typing.Optional[bool] = None, sort_by: typing.Optional[str] = None, sort_desc: typing.Optional[bool] = None, width: typing.Optional[float] = None, calcs: typing.Optional[list[str]] = None):
        self.values = values if values is not None else []
        self.display_mode = display_mode if display_mode is not None else common.LegendDisplayMode.LIST
        self.placement = placement if placement is not None else common.LegendPlacement.BOTTOM
        self.show_legend = show_legend
        self.as_table = as_table
        self.is_visible = is_visible
        self.sort_by = sort_by
        self.sort_desc = sort_desc
        self.width = width
        self.calcs = calcs if calcs is not None else []

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "values": self.values,
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
        
        if "values" in data:
            args["values"] = data["values"]
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


class Options:
    pie_type: 'PieChartType'
    display_labels: typing.Optional[list['PieChartLabels']]
    tooltip: common.VizTooltipOptions
    reduce_options: common.ReduceDataOptions
    text: typing.Optional[common.VizTextDisplayOptions]
    legend: 'PieChartLegendOptions'
    orientation: common.VizOrientation

    def __init__(self, pie_type: typing.Optional['PieChartType'] = None, display_labels: typing.Optional[list['PieChartLabels']] = None, tooltip: typing.Optional[common.VizTooltipOptions] = None, reduce_options: typing.Optional[common.ReduceDataOptions] = None, text: typing.Optional[common.VizTextDisplayOptions] = None, legend: typing.Optional['PieChartLegendOptions'] = None, orientation: typing.Optional[common.VizOrientation] = None):
        self.pie_type = pie_type if pie_type is not None else PieChartType.PIE
        self.display_labels = display_labels
        self.tooltip = tooltip if tooltip is not None else common.VizTooltipOptions()
        self.reduce_options = reduce_options if reduce_options is not None else common.ReduceDataOptions()
        self.text = text
        self.legend = legend if legend is not None else PieChartLegendOptions()
        self.orientation = orientation if orientation is not None else common.VizOrientation.AUTO

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "pieType": self.pie_type,
            "tooltip": self.tooltip,
            "reduceOptions": self.reduce_options,
            "legend": self.legend,
            "orientation": self.orientation,
        }
        if self.display_labels is not None:
            payload["displayLabels"] = self.display_labels
        if self.text is not None:
            payload["text"] = self.text
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "pieType" in data:
            args["pie_type"] = data["pieType"]
        if "displayLabels" in data:
            args["display_labels"] = data["displayLabels"]
        if "tooltip" in data:
            args["tooltip"] = common.VizTooltipOptions.from_json(data["tooltip"])
        if "reduceOptions" in data:
            args["reduce_options"] = common.ReduceDataOptions.from_json(data["reduceOptions"])
        if "text" in data:
            args["text"] = common.VizTextDisplayOptions.from_json(data["text"])
        if "legend" in data:
            args["legend"] = PieChartLegendOptions.from_json(data["legend"])
        if "orientation" in data:
            args["orientation"] = data["orientation"]        

        return cls(**args)


FieldConfig: typing.TypeAlias = common.HideableFieldConfig





def variant_config():
    return cogruntime.PanelCfgConfig(
        identifier="piechart",
        options_from_json_hook=Options.from_json,
        field_config_from_json_hook=FieldConfig.from_json,
    )
