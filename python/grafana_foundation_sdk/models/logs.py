# Code generated - EDITING IS FUTILE. DO NOT EDIT.

import typing
from ..models import common
from ..cog import runtime as cogruntime


class Options:
    show_labels: bool
    show_common_labels: bool
    show_time: bool
    show_log_context_toggle: bool
    show_controls: typing.Optional[bool]
    controls_storage_key: typing.Optional[str]
    wrap_log_message: bool
    prettify_log_message: bool
    enable_log_details: bool
    sort_order: common.LogsSortOrder
    dedup_strategy: common.LogsDedupStrategy
    enable_infinite_scrolling: typing.Optional[bool]
    # TODO: figure out how to define callbacks
    on_click_filter_label: typing.Optional[object]
    on_click_filter_out_label: typing.Optional[object]
    is_filter_label_active: typing.Optional[object]
    on_click_filter_string: typing.Optional[object]
    on_click_filter_out_string: typing.Optional[object]
    on_click_show_field: typing.Optional[object]
    on_click_hide_field: typing.Optional[object]
    on_log_options_change: typing.Optional[object]
    log_row_menu_icons_before: typing.Optional[object]
    log_row_menu_icons_after: typing.Optional[object]
    on_new_logs_received: typing.Optional[object]
    displayed_fields: typing.Optional[list[str]]

    def __init__(self, show_labels: bool = False, show_common_labels: bool = False, show_time: bool = False, show_log_context_toggle: bool = False, show_controls: typing.Optional[bool] = None, controls_storage_key: typing.Optional[str] = None, wrap_log_message: bool = False, prettify_log_message: bool = False, enable_log_details: bool = False, sort_order: typing.Optional[common.LogsSortOrder] = None, dedup_strategy: typing.Optional[common.LogsDedupStrategy] = None, enable_infinite_scrolling: typing.Optional[bool] = None, on_click_filter_label: typing.Optional[object] = None, on_click_filter_out_label: typing.Optional[object] = None, is_filter_label_active: typing.Optional[object] = None, on_click_filter_string: typing.Optional[object] = None, on_click_filter_out_string: typing.Optional[object] = None, on_click_show_field: typing.Optional[object] = None, on_click_hide_field: typing.Optional[object] = None, on_log_options_change: typing.Optional[object] = None, log_row_menu_icons_before: typing.Optional[object] = None, log_row_menu_icons_after: typing.Optional[object] = None, on_new_logs_received: typing.Optional[object] = None, displayed_fields: typing.Optional[list[str]] = None):
        self.show_labels = show_labels
        self.show_common_labels = show_common_labels
        self.show_time = show_time
        self.show_log_context_toggle = show_log_context_toggle
        self.show_controls = show_controls
        self.controls_storage_key = controls_storage_key
        self.wrap_log_message = wrap_log_message
        self.prettify_log_message = prettify_log_message
        self.enable_log_details = enable_log_details
        self.sort_order = sort_order if sort_order is not None else common.LogsSortOrder.DESCENDING
        self.dedup_strategy = dedup_strategy if dedup_strategy is not None else common.LogsDedupStrategy.NONE
        self.enable_infinite_scrolling = enable_infinite_scrolling
        self.on_click_filter_label = on_click_filter_label
        self.on_click_filter_out_label = on_click_filter_out_label
        self.is_filter_label_active = is_filter_label_active
        self.on_click_filter_string = on_click_filter_string
        self.on_click_filter_out_string = on_click_filter_out_string
        self.on_click_show_field = on_click_show_field
        self.on_click_hide_field = on_click_hide_field
        self.on_log_options_change = on_log_options_change
        self.log_row_menu_icons_before = log_row_menu_icons_before
        self.log_row_menu_icons_after = log_row_menu_icons_after
        self.on_new_logs_received = on_new_logs_received
        self.displayed_fields = displayed_fields

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
        if self.show_controls is not None:
            payload["showControls"] = self.show_controls
        if self.controls_storage_key is not None:
            payload["controlsStorageKey"] = self.controls_storage_key
        if self.enable_infinite_scrolling is not None:
            payload["enableInfiniteScrolling"] = self.enable_infinite_scrolling
        if self.on_click_filter_label is not None:
            payload["onClickFilterLabel"] = self.on_click_filter_label
        if self.on_click_filter_out_label is not None:
            payload["onClickFilterOutLabel"] = self.on_click_filter_out_label
        if self.is_filter_label_active is not None:
            payload["isFilterLabelActive"] = self.is_filter_label_active
        if self.on_click_filter_string is not None:
            payload["onClickFilterString"] = self.on_click_filter_string
        if self.on_click_filter_out_string is not None:
            payload["onClickFilterOutString"] = self.on_click_filter_out_string
        if self.on_click_show_field is not None:
            payload["onClickShowField"] = self.on_click_show_field
        if self.on_click_hide_field is not None:
            payload["onClickHideField"] = self.on_click_hide_field
        if self.on_log_options_change is not None:
            payload["onLogOptionsChange"] = self.on_log_options_change
        if self.log_row_menu_icons_before is not None:
            payload["logRowMenuIconsBefore"] = self.log_row_menu_icons_before
        if self.log_row_menu_icons_after is not None:
            payload["logRowMenuIconsAfter"] = self.log_row_menu_icons_after
        if self.on_new_logs_received is not None:
            payload["onNewLogsReceived"] = self.on_new_logs_received
        if self.displayed_fields is not None:
            payload["displayedFields"] = self.displayed_fields
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
        if "showControls" in data:
            args["show_controls"] = data["showControls"]
        if "controlsStorageKey" in data:
            args["controls_storage_key"] = data["controlsStorageKey"]
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
        if "onClickFilterLabel" in data:
            args["on_click_filter_label"] = data["onClickFilterLabel"]
        if "onClickFilterOutLabel" in data:
            args["on_click_filter_out_label"] = data["onClickFilterOutLabel"]
        if "isFilterLabelActive" in data:
            args["is_filter_label_active"] = data["isFilterLabelActive"]
        if "onClickFilterString" in data:
            args["on_click_filter_string"] = data["onClickFilterString"]
        if "onClickFilterOutString" in data:
            args["on_click_filter_out_string"] = data["onClickFilterOutString"]
        if "onClickShowField" in data:
            args["on_click_show_field"] = data["onClickShowField"]
        if "onClickHideField" in data:
            args["on_click_hide_field"] = data["onClickHideField"]
        if "onLogOptionsChange" in data:
            args["on_log_options_change"] = data["onLogOptionsChange"]
        if "logRowMenuIconsBefore" in data:
            args["log_row_menu_icons_before"] = data["logRowMenuIconsBefore"]
        if "logRowMenuIconsAfter" in data:
            args["log_row_menu_icons_after"] = data["logRowMenuIconsAfter"]
        if "onNewLogsReceived" in data:
            args["on_new_logs_received"] = data["onNewLogsReceived"]
        if "displayedFields" in data:
            args["displayed_fields"] = data["displayedFields"]        

        return cls(**args)


def variant_config() -> cogruntime.PanelCfgConfig:
    return cogruntime.PanelCfgConfig(
        identifier="logs",
        options_from_json_hook=Options.from_json,
        field_config_from_json_hook=None,
    )

