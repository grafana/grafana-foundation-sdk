# Code generated - EDITING IS FUTILE. DO NOT EDIT.

from ..models import common
import typing
from ..cog import runtime as cogruntime


class Options:
    show_labels: bool
    show_common_labels: bool
    show_time: bool
    show_log_context_toggle: bool
    wrap_log_message: bool
    prettify_log_message: bool
    enable_log_details: bool
    sort_order: common.LogsSortOrder
    dedup_strategy: common.LogsDedupStrategy
    enable_infinite_scrolling: typing.Optional[bool]
    # Display controls to jump to the last or first log line, and filters by log level.
    show_controls: typing.Optional[bool]
    # Show a component to manage the displayed fields from the logs.
    show_field_selector: typing.Optional[bool]
    # Use a predefined coloring scheme to highlight relevant parts of the log lines.
    syntax_highlighting: typing.Optional[bool]
    font_size: typing.Optional[typing.Literal["default", "small"]]
    details_mode: typing.Optional[typing.Literal["inline", "sidebar"]]

    def __init__(self, show_labels: bool = False, show_common_labels: bool = False, show_time: bool = False, show_log_context_toggle: bool = False, wrap_log_message: bool = False, prettify_log_message: bool = False, enable_log_details: bool = False, sort_order: typing.Optional[common.LogsSortOrder] = None, dedup_strategy: typing.Optional[common.LogsDedupStrategy] = None, enable_infinite_scrolling: typing.Optional[bool] = None, show_controls: typing.Optional[bool] = None, show_field_selector: typing.Optional[bool] = None, syntax_highlighting: typing.Optional[bool] = None, font_size: typing.Optional[typing.Literal["default", "small"]] = None, details_mode: typing.Optional[typing.Literal["inline", "sidebar"]] = None) -> None:
        self.show_labels = show_labels
        self.show_common_labels = show_common_labels
        self.show_time = show_time
        self.show_log_context_toggle = show_log_context_toggle
        self.wrap_log_message = wrap_log_message
        self.prettify_log_message = prettify_log_message
        self.enable_log_details = enable_log_details
        self.sort_order = sort_order if sort_order is not None else common.LogsSortOrder.DESCENDING
        self.dedup_strategy = dedup_strategy if dedup_strategy is not None else common.LogsDedupStrategy.NONE
        self.enable_infinite_scrolling = enable_infinite_scrolling
        self.show_controls = show_controls
        self.show_field_selector = show_field_selector
        self.syntax_highlighting = syntax_highlighting
        self.font_size = font_size
        self.details_mode = details_mode

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "showLabels": self.show_labels,
            "showCommonLabels": self.show_common_labels,
            "showTime": self.show_time,
            "showLogContextToggle": self.show_log_context_toggle,
            "wrapLogMessage": self.wrap_log_message,
            "prettifyLogMessage": self.prettify_log_message,
            "enableLogDetails": self.enable_log_details,
            "sortOrder": self.sort_order,
            "dedupStrategy": self.dedup_strategy,
        }
        if self.enable_infinite_scrolling is not None:
            payload["enableInfiniteScrolling"] = self.enable_infinite_scrolling
        if self.show_controls is not None:
            payload["showControls"] = self.show_controls
        if self.show_field_selector is not None:
            payload["showFieldSelector"] = self.show_field_selector
        if self.syntax_highlighting is not None:
            payload["syntaxHighlighting"] = self.syntax_highlighting
        if self.font_size is not None:
            payload["fontSize"] = self.font_size
        if self.details_mode is not None:
            payload["detailsMode"] = self.details_mode
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
        if "showLogContextToggle" in data:
            args["show_log_context_toggle"] = data["showLogContextToggle"]
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
        if "enableInfiniteScrolling" in data:
            args["enable_infinite_scrolling"] = data["enableInfiniteScrolling"]
        if "showControls" in data:
            args["show_controls"] = data["showControls"]
        if "showFieldSelector" in data:
            args["show_field_selector"] = data["showFieldSelector"]
        if "syntaxHighlighting" in data:
            args["syntax_highlighting"] = data["syntaxHighlighting"]
        if "fontSize" in data:
            args["font_size"] = data["fontSize"]
        if "detailsMode" in data:
            args["details_mode"] = data["detailsMode"]        

        return cls(**args)


def variant_config() -> cogruntime.PanelCfgConfig:
    return cogruntime.PanelCfgConfig(
        identifier="logs",
        options_from_json_hook=Options.from_json,
        field_config_from_json_hook=None,
    )

