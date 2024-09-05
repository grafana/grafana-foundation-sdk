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
    # TODO: figure out how to define callbacks
    on_click_filter_label: typing.Optional[object]
    on_click_filter_out_label: typing.Optional[object]
    is_filter_label_active: typing.Optional[object]
    on_click_filter_string: typing.Optional[object]
    on_click_filter_out_string: typing.Optional[object]
    on_click_show_field: typing.Optional[object]
    on_click_hide_field: typing.Optional[object]
    displayed_fields: typing.Optional[list[str]]

    def __init__(self, show_labels: bool = False, show_common_labels: bool = False, show_time: bool = False, show_log_context_toggle: bool = False, wrap_log_message: bool = False, prettify_log_message: bool = False, enable_log_details: bool = False, sort_order: typing.Optional[common.LogsSortOrder] = None, dedup_strategy: typing.Optional[common.LogsDedupStrategy] = None, on_click_filter_label: typing.Optional[object] = None, on_click_filter_out_label: typing.Optional[object] = None, is_filter_label_active: typing.Optional[object] = None, on_click_filter_string: typing.Optional[object] = None, on_click_filter_out_string: typing.Optional[object] = None, on_click_show_field: typing.Optional[object] = None, on_click_hide_field: typing.Optional[object] = None, displayed_fields: typing.Optional[list[str]] = None):
        self.show_labels = show_labels
        self.show_common_labels = show_common_labels
        self.show_time = show_time
        self.show_log_context_toggle = show_log_context_toggle
        self.wrap_log_message = wrap_log_message
        self.prettify_log_message = prettify_log_message
        self.enable_log_details = enable_log_details
        self.sort_order = sort_order if sort_order is not None else common.LogsSortOrder.DESCENDING
        self.dedup_strategy = dedup_strategy if dedup_strategy is not None else common.LogsDedupStrategy.NONE
        self.on_click_filter_label = on_click_filter_label
        self.on_click_filter_out_label = on_click_filter_out_label
        self.is_filter_label_active = is_filter_label_active
        self.on_click_filter_string = on_click_filter_string
        self.on_click_filter_out_string = on_click_filter_out_string
        self.on_click_show_field = on_click_show_field
        self.on_click_hide_field = on_click_hide_field
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
        if "displayedFields" in data:
            args["displayed_fields"] = data["displayedFields"]        

        return cls(**args)


def variant_config():
    return cogruntime.PanelCfgConfig(
        identifier="logs",
        options_from_json_hook=Options.from_json,
        field_config_from_json_hook=None,
    )
