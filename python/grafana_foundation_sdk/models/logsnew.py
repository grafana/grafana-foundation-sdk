# Code generated - EDITING IS FUTILE. DO NOT EDIT.

from ..models import common
import typing
from ..cog import runtime as cogruntime


class Options:
    show_controls: bool
    show_time: bool
    wrap_log_message: bool
    enable_log_details: bool
    syntax_highlighting: bool
    sort_order: common.LogsSortOrder
    dedup_strategy: common.LogsDedupStrategy
    grammar: typing.Optional[object]
    enable_infinite_scrolling: typing.Optional[bool]
    on_log_options_change: typing.Optional[object]
    on_new_logs_received: typing.Optional[object]

    def __init__(self, show_controls: bool = False, show_time: bool = False, wrap_log_message: bool = False, enable_log_details: bool = False, syntax_highlighting: bool = False, sort_order: typing.Optional[common.LogsSortOrder] = None, dedup_strategy: typing.Optional[common.LogsDedupStrategy] = None, grammar: typing.Optional[object] = None, enable_infinite_scrolling: typing.Optional[bool] = None, on_log_options_change: typing.Optional[object] = None, on_new_logs_received: typing.Optional[object] = None):
        self.show_controls = show_controls
        self.show_time = show_time
        self.wrap_log_message = wrap_log_message
        self.enable_log_details = enable_log_details
        self.syntax_highlighting = syntax_highlighting
        self.sort_order = sort_order if sort_order is not None else common.LogsSortOrder.DESCENDING
        self.dedup_strategy = dedup_strategy if dedup_strategy is not None else common.LogsDedupStrategy.NONE
        self.grammar = grammar
        self.enable_infinite_scrolling = enable_infinite_scrolling
        self.on_log_options_change = on_log_options_change
        self.on_new_logs_received = on_new_logs_received

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "showControls": self.show_controls,
            "showTime": self.show_time,
            "wrapLogMessage": self.wrap_log_message,
            "enableLogDetails": self.enable_log_details,
            "syntaxHighlighting": self.syntax_highlighting,
            "sortOrder": self.sort_order,
            "dedupStrategy": self.dedup_strategy,
        }
        if self.grammar is not None:
            payload["grammar"] = self.grammar
        if self.enable_infinite_scrolling is not None:
            payload["enableInfiniteScrolling"] = self.enable_infinite_scrolling
        if self.on_log_options_change is not None:
            payload["onLogOptionsChange"] = self.on_log_options_change
        if self.on_new_logs_received is not None:
            payload["onNewLogsReceived"] = self.on_new_logs_received
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "showControls" in data:
            args["show_controls"] = data["showControls"]
        if "showTime" in data:
            args["show_time"] = data["showTime"]
        if "wrapLogMessage" in data:
            args["wrap_log_message"] = data["wrapLogMessage"]
        if "enableLogDetails" in data:
            args["enable_log_details"] = data["enableLogDetails"]
        if "syntaxHighlighting" in data:
            args["syntax_highlighting"] = data["syntaxHighlighting"]
        if "sortOrder" in data:
            args["sort_order"] = data["sortOrder"]
        if "dedupStrategy" in data:
            args["dedup_strategy"] = data["dedupStrategy"]
        if "grammar" in data:
            args["grammar"] = data["grammar"]
        if "enableInfiniteScrolling" in data:
            args["enable_infinite_scrolling"] = data["enableInfiniteScrolling"]
        if "onLogOptionsChange" in data:
            args["on_log_options_change"] = data["onLogOptionsChange"]
        if "onNewLogsReceived" in data:
            args["on_new_logs_received"] = data["onNewLogsReceived"]        

        return cls(**args)


def variant_config() -> cogruntime.PanelCfgConfig:
    return cogruntime.PanelCfgConfig(
        identifier="logsnew",
        options_from_json_hook=Options.from_json,
        field_config_from_json_hook=None,
    )

