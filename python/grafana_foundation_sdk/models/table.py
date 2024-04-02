# Code generated - EDITING IS FUTILE. DO NOT EDIT.

import typing
from ..models import common
from ..cog import runtime as cogruntime


class Options:
    # Represents the index of the selected frame
    frame_index: float
    # Controls whether the panel should show the header
    show_header: bool
    # Controls whether the header should show icons for the column types
    show_type_icons: typing.Optional[bool]
    # Used to control row sorting
    sort_by: typing.Optional[list[common.TableSortByFieldState]]
    # Controls footer options
    footer: typing.Optional[common.TableFooterOptions]
    # Controls the height of the rows
    cell_height: typing.Optional[common.TableCellHeight]

    def __init__(self, frame_index: float = 0, show_header: bool = True, show_type_icons: typing.Optional[bool] = False, sort_by: typing.Optional[list[common.TableSortByFieldState]] = None, footer: typing.Optional[common.TableFooterOptions] = None, cell_height: typing.Optional[common.TableCellHeight] = None):
        self.frame_index = frame_index
        self.show_header = show_header
        self.show_type_icons = show_type_icons
        self.sort_by = sort_by
        self.footer = footer if footer is not None else common.TableFooterOptions(count_rows=False, reducer=None, show=False)
        self.cell_height = cell_height if cell_height is not None else common.TableCellHeight.SM

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "frameIndex": self.frame_index,
            "showHeader": self.show_header,
        }
        if self.show_type_icons is not None:
            payload["showTypeIcons"] = self.show_type_icons
        if self.sort_by is not None:
            payload["sortBy"] = self.sort_by
        if self.footer is not None:
            payload["footer"] = self.footer
        if self.cell_height is not None:
            payload["cellHeight"] = self.cell_height
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "frameIndex" in data:
            args["frame_index"] = data["frameIndex"]
        if "showHeader" in data:
            args["show_header"] = data["showHeader"]
        if "showTypeIcons" in data:
            args["show_type_icons"] = data["showTypeIcons"]
        if "sortBy" in data:
            args["sort_by"] = data["sortBy"]
        if "footer" in data:
            args["footer"] = common.TableFooterOptions.from_json(data["footer"])
        if "cellHeight" in data:
            args["cell_height"] = data["cellHeight"]        

        return cls(**args)


def variant_config():
    return cogruntime.PanelCfgConfig(
        identifier="table",
        options_from_json_hook=Options.from_json,
        field_config_from_json_hook=None,
    )
