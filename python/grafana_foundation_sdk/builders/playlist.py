# Code generated - EDITING IS FUTILE. DO NOT EDIT.

import typing
from ..cog import builder as cogbuilder
from ..models import playlist


class Playlist(cogbuilder.Builder[playlist.Playlist]):    
    __internal: playlist.Playlist

    def __init__(self):
        self.__internal = playlist.Playlist()

    def build(self) -> playlist.Playlist:
        return self.__internal    
    
    def uid(self, uid: str) -> typing.Self:    
        """
        Unique playlist identifier. Generated on creation, either by the
        creator of the playlist of by the application.
        """
            
        self.__internal.uid = uid
    
        return self
    
    def name(self, name: str) -> typing.Self:    
        """
        Name of the playlist.
        """
            
        self.__internal.name = name
    
        return self
    
    def interval(self, interval: str) -> typing.Self:    
        """
        Interval sets the time between switching views in a playlist.
        FIXME: Is this based on a standardized format or what options are available? Can datemath be used?
        """
            
        self.__internal.interval = interval
    
        return self
    
    def items(self, items: list[cogbuilder.Builder[playlist.PlaylistItem]]) -> typing.Self:    
        """
        The ordered list of items that the playlist will iterate over.
        FIXME! This should not be optional, but changing it makes the godegen awkward
        """
            
        items_resources = [r1.build() for r1 in items]
        self.__internal.items = items_resources
    
        return self
    

class PlaylistItem(cogbuilder.Builder[playlist.PlaylistItem]):    
    __internal: playlist.PlaylistItem

    def __init__(self):
        self.__internal = playlist.PlaylistItem()

    def build(self) -> playlist.PlaylistItem:
        return self.__internal    
    
    def type_val(self, type_val: typing.Literal["dashboard_by_uid", "dashboard_by_id", "dashboard_by_tag"]) -> typing.Self:    
        """
        Type of the item.
        """
            
        self.__internal.type_val = type_val
    
        return self
    
    def value(self, value: str) -> typing.Self:    
        """
        Value depends on type and describes the playlist item.
        
         - dashboard_by_id: The value is an internal numerical identifier set by Grafana. This
         is not portable as the numerical identifier is non-deterministic between different instances.
         Will be replaced by dashboard_by_uid in the future. (deprecated)
         - dashboard_by_tag: The value is a tag which is set on any number of dashboards. All
         dashboards behind the tag will be added to the playlist.
         - dashboard_by_uid: The value is the dashboard UID
        """
            
        self.__internal.value = value
    
        return self
    
    def title(self, title: str) -> typing.Self:    
        """
        Title is an unused property -- it will be removed in the future
        """
            
        self.__internal.title = title
    
        return self
    