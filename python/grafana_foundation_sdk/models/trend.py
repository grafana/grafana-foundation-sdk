# Code generated - EDITING IS FUTILE. DO NOT EDIT.

from ..models import common
import typing
from ..cog import runtime as cogruntime


class Options:
    """
    Identical to timeseries... except it does not have timezone settings
    """

    legend: common.VizLegendOptions
    tooltip: common.VizTooltipOptions
    # Name of the x field to use (defaults to first number)
    x_field: typing.Optional[str]

    def __init__(self, legend: typing.Optional[common.VizLegendOptions] = None, tooltip: typing.Optional[common.VizTooltipOptions] = None, x_field: typing.Optional[str] = None):
        self.legend = legend if legend is not None else common.VizLegendOptions()
        self.tooltip = tooltip if tooltip is not None else common.VizTooltipOptions()
        self.x_field = x_field

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "legend": self.legend,
            "tooltip": self.tooltip,
        }
        if self.x_field is not None:
            payload["xField"] = self.x_field
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "legend" in data:
            args["legend"] = common.VizLegendOptions.from_json(data["legend"])
        if "tooltip" in data:
            args["tooltip"] = common.VizTooltipOptions.from_json(data["tooltip"])
        if "xField" in data:
            args["x_field"] = data["xField"]        

        return cls(**args)


FieldConfig: typing.TypeAlias = common.GraphFieldConfig





def variant_config():
    return cogruntime.PanelCfgConfig(
        identifier="trend",
        options_from_json_hook=Options.from_json,
        field_config_from_json_hook=FieldConfig.from_json,
    )
