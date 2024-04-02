# Code generated - EDITING IS FUTILE. DO NOT EDIT.

from ..models import common
import typing
from ..cog import runtime as cogruntime


class Options:
    show_labels: bool
    show_common_labels: bool
    show_time: bool
    wrap_log_message: bool
    prettify_log_message: bool
    enable_log_details: bool
    sort_order: common.LogsSortOrder
    dedup_strategy: common.LogsDedupStrategy

    def __init__(self, show_labels: bool = False, show_common_labels: bool = False, show_time: bool = False, wrap_log_message: bool = False, prettify_log_message: bool = False, enable_log_details: bool = False, sort_order: typing.Optional[common.LogsSortOrder] = None, dedup_strategy: typing.Optional[common.LogsDedupStrategy] = None):
        self.show_labels = show_labels
        self.show_common_labels = show_common_labels
        self.show_time = show_time
        self.wrap_log_message = wrap_log_message
        self.prettify_log_message = prettify_log_message
        self.enable_log_details = enable_log_details
        self.sort_order = sort_order if sort_order is not None else common.LogsSortOrder.DESCENDING
        self.dedup_strategy = dedup_strategy if dedup_strategy is not None else common.LogsDedupStrategy.NONE

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "showLabels": self.show_labels,
            "showCommonLabels": self.show_common_labels,
            "showTime": self.show_time,
            "wrapLogMessage": self.wrap_log_message,
            "prettifyLogMessage": self.prettify_log_message,
            "enableLogDetails": self.enable_log_details,
            "sortOrder": self.sort_order,
            "dedupStrategy": self.dedup_strategy,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "showLabels" in data:
            args["show_labels"] = data["showLabels"]
        if "showCommonLabels" in data:
            args["show_common_labels"] = data["showCommonLabels"]
        if "showTime" in data:
            args["show_time"] = data["showTime"]
        if "wrapLogMessage" in data:
            args["wrap_log_message"] = data["wrapLogMessage"]
        if "prettifyLogMessage" in data:
            args["prettify_log_message"] = data["prettifyLogMessage"]
        if "enableLogDetails" in data:
            args["enable_log_details"] = data["enableLogDetails"]
        if "sortOrder" in data:
            args["sort_order"] = data["sortOrder"]
        if "dedupStrategy" in data:
            args["dedup_strategy"] = data["dedupStrategy"]        

        return cls(**args)


def variant_config():
    return cogruntime.PanelCfgConfig(
        identifier="logs",
        options_from_json_hook=Options.from_json,
        field_config_from_json_hook=None,
    )
