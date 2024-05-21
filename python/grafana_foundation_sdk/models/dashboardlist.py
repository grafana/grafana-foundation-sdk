# Code generated - EDITING IS FUTILE. DO NOT EDIT.

import typing
from ..cog import runtime as cogruntime


class Options:
    keep_time: bool
    include_vars: bool
    show_starred: bool
    show_recently_viewed: bool
    show_search: bool
    show_headings: bool
    show_folder_names: bool
    max_items: int
    query: str
    tags: list[str]
    # folderId is deprecated, and migrated to folderUid on panel init
    folder_id: typing.Optional[int]
    folder_uid: typing.Optional[str]

    def __init__(self, keep_time: bool = False, include_vars: bool = False, show_starred: bool = True, show_recently_viewed: bool = False, show_search: bool = False, show_headings: bool = True, show_folder_names: bool = True, max_items: int = 10, query: str = "", tags: typing.Optional[list[str]] = None, folder_id: typing.Optional[int] = None, folder_uid: typing.Optional[str] = None):
        self.keep_time = keep_time
        self.include_vars = include_vars
        self.show_starred = show_starred
        self.show_recently_viewed = show_recently_viewed
        self.show_search = show_search
        self.show_headings = show_headings
        self.show_folder_names = show_folder_names
        self.max_items = max_items
        self.query = query
        self.tags = tags if tags is not None else []
        self.folder_id = folder_id
        self.folder_uid = folder_uid

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "keepTime": self.keep_time,
            "includeVars": self.include_vars,
            "showStarred": self.show_starred,
            "showRecentlyViewed": self.show_recently_viewed,
            "showSearch": self.show_search,
            "showHeadings": self.show_headings,
            "showFolderNames": self.show_folder_names,
            "maxItems": self.max_items,
            "query": self.query,
            "tags": self.tags,
        }
        if self.folder_id is not None:
            payload["folderId"] = self.folder_id
        if self.folder_uid is not None:
            payload["folderUID"] = self.folder_uid
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "keepTime" in data:
            args["keep_time"] = data["keepTime"]
        if "includeVars" in data:
            args["include_vars"] = data["includeVars"]
        if "showStarred" in data:
            args["show_starred"] = data["showStarred"]
        if "showRecentlyViewed" in data:
            args["show_recently_viewed"] = data["showRecentlyViewed"]
        if "showSearch" in data:
            args["show_search"] = data["showSearch"]
        if "showHeadings" in data:
            args["show_headings"] = data["showHeadings"]
        if "showFolderNames" in data:
            args["show_folder_names"] = data["showFolderNames"]
        if "maxItems" in data:
            args["max_items"] = data["maxItems"]
        if "query" in data:
            args["query"] = data["query"]
        if "tags" in data:
            args["tags"] = data["tags"]
        if "folderId" in data:
            args["folder_id"] = data["folderId"]
        if "folderUID" in data:
            args["folder_uid"] = data["folderUID"]        

        return cls(**args)


def variant_config():
    return cogruntime.PanelCfgConfig(
        identifier="dashlist",
        options_from_json_hook=Options.from_json,
        field_config_from_json_hook=None,
    )
