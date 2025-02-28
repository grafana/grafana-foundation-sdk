# Code generated - EDITING IS FUTILE. DO NOT EDIT.

from ..models import common
import typing
from ..cog import runtime as cogruntime


class Options:
    show_time: bool
    wrap_log_message: bool
    enable_log_details: bool
    sort_order: common.LogsSortOrder
    dedup_strategy: common.LogsDedupStrategy
    enable_infinite_scrolling: typing.Optional[bool]
    on_new_logs_received: typing.Optional[object]

    def __init__(self, show_time: bool = False, wrap_log_message: bool = False, enable_log_details: bool = False, sort_order: typing.Optional[common.LogsSortOrder] = None, dedup_strategy: typing.Optional[common.LogsDedupStrategy] = None, enable_infinite_scrolling: typing.Optional[bool] = None, on_new_logs_received: typing.Optional[object] = None):
        self.show_time = show_time
        self.wrap_log_message = wrap_log_message
        self.enable_log_details = enable_log_details
        self.sort_order = sort_order if sort_order is not None else common.LogsSortOrder.DESCENDING
        self.dedup_strategy = dedup_strategy if dedup_strategy is not None else common.LogsDedupStrategy.NONE
        self.enable_infinite_scrolling = enable_infinite_scrolling
        self.on_new_logs_received = on_new_logs_received

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "showTime": self.show_time,
            "wrapLogMessage": self.wrap_log_message,
            "enableLogDetails": self.enable_log_details,
            "sortOrder": self.sort_order,
            "dedupStrategy": self.dedup_strategy,
        }
        if self.enable_infinite_scrolling is not None:
            payload["enableInfiniteScrolling"] = self.enable_infinite_scrolling
        if self.on_new_logs_received is not None:
            payload["onNewLogsReceived"] = self.on_new_logs_received
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "showTime" in data:
            args["show_time"] = data["showTime"]
        if "wrapLogMessage" in data:
            args["wrap_log_message"] = data["wrapLogMessage"]
        if "enableLogDetails" in data:
            args["enable_log_details"] = data["enableLogDetails"]
        if "sortOrder" in data:
            args["sort_order"] = data["sortOrder"]
        if "dedupStrategy" in data:
            args["dedup_strategy"] = data["dedupStrategy"]
        if "enableInfiniteScrolling" in data:
            args["enable_infinite_scrolling"] = data["enableInfiniteScrolling"]
        if "onNewLogsReceived" in data:
            args["on_new_logs_received"] = data["onNewLogsReceived"]        

        return cls(**args)


def variant_config() -> cogruntime.PanelCfgConfig:
    return cogruntime.PanelCfgConfig(
        identifier="logsnew",
        options_from_json_hook=Options.from_json,
        field_config_from_json_hook=None,
    )

