# Code generated - EDITING IS FUTILE. DO NOT EDIT.

from ..models import common
import typing
from ..cog import runtime as cogruntime


class Options:
    timezone: typing.Optional[list[common.TimeZone]]
    legend: common.VizLegendOptions
    tooltip: common.VizTooltipOptions

    def __init__(self, timezone: typing.Optional[list[common.TimeZone]] = None, legend: typing.Optional[common.VizLegendOptions] = None, tooltip: typing.Optional[common.VizTooltipOptions] = None):
        self.timezone = timezone
        self.legend = legend if legend is not None else common.VizLegendOptions()
        self.tooltip = tooltip if tooltip is not None else common.VizTooltipOptions()

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "legend": self.legend,
            "tooltip": self.tooltip,
        }
        if self.timezone is not None:
            payload["timezone"] = self.timezone
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "timezone" in data:
            args["timezone"] = data["timezone"]
        if "legend" in data:
            args["legend"] = common.VizLegendOptions.from_json(data["legend"])
        if "tooltip" in data:
            args["tooltip"] = common.VizTooltipOptions.from_json(data["tooltip"])        

        return cls(**args)


FieldConfig: typing.TypeAlias = common.GraphFieldConfig





def variant_config():
    return cogruntime.PanelCfgConfig(
        identifier="timeseries",
        options_from_json_hook=Options.from_json,
        field_config_from_json_hook=FieldConfig.from_json,
    )
