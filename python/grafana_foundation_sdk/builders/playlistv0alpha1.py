# Code generated - EDITING IS FUTILE. DO NOT EDIT.

import typing
from ..cog import builder as cogbuilder
from ..models import playlistv0alpha1
from ..builders import resource as resource_builder
from ..models import resource


class Item(cogbuilder.Builder[playlistv0alpha1.Item]):
    _internal: playlistv0alpha1.Item

    def __init__(self) -> None:
        self._internal = playlistv0alpha1.Item()

    def build(self) -> playlistv0alpha1.Item:
        """
        Builds the object.
        """
        return self._internal    
    
    def type(self, type_val: typing.Literal["dashboard_by_tag", "dashboard_by_uid", "dashboard_by_id"]) -> typing.Self:    
        """
        type of the item.
        """
            
        self._internal.type_val = type_val
    
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
            
        self._internal.value = value
    
        return self
    


class Playlist(cogbuilder.Builder[playlistv0alpha1.Playlist]):
    _internal: playlistv0alpha1.Playlist

    def __init__(self, title: str) -> None:
        self._internal = playlistv0alpha1.Playlist()        
        self._internal.title = title

    def build(self) -> playlistv0alpha1.Playlist:
        """
        Builds the object.
        """
        return self._internal    
    
    def title(self, title: str) -> typing.Self:    
        self._internal.title = title
    
        return self
    
    def interval(self, interval: str) -> typing.Self:    
        self._internal.interval = interval
    
        return self
    
    def items(self, items: list[cogbuilder.Builder[playlistv0alpha1.Item]]) -> typing.Self:    
        items_resources = [r1.build() for r1 in items]
        self._internal.items = items_resources
    
        return self
    
    def item(self, item: cogbuilder.Builder[playlistv0alpha1.Item]) -> typing.Self:    
        if self._internal.items is None:
            self._internal.items = []
        
        item_resource = item.build()
        self._internal.items.append(item_resource)
    
        return self
    

"""
Creates a resource manifest from a Playlist.
"""
def manifest(name: str, playlist: cogbuilder.Builder[playlistv0alpha1.Playlist]) -> cogbuilder.Builder[resource.Manifest]:
    builder = resource_builder.Manifest()
    builder.api_version("playlist.grafana.app/playlistv0alpha1")
    builder.kind("Playlist")
    builder.metadata(resource_builder.named(name))
    builder.spec(playlist.build())

    return builder
