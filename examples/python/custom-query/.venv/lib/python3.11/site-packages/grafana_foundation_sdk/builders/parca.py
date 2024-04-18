# Code generated - EDITING IS FUTILE. DO NOT EDIT.

import typing
from ..cog import builder as cogbuilder
from ..models import parca


class Dataquery(cogbuilder.Builder[parca.Dataquery]):    
    __internal: parca.Dataquery

    def __init__(self):
        self.__internal = parca.Dataquery()

    def build(self) -> parca.Dataquery:
        return self.__internal    
    
    def label_selector(self, label_selector: str) -> typing.Self:        
        self.__internal.label_selector = label_selector
    
        return self
    
    def profile_type_id(self, profile_type_id: str) -> typing.Self:        
        self.__internal.profile_type_id = profile_type_id
    
        return self
    
    def ref_id(self, ref_id: str) -> typing.Self:        
        self.__internal.ref_id = ref_id
    
        return self
    
    def hide(self, hide: bool) -> typing.Self:        
        self.__internal.hide = hide
    
        return self
    
    def query_type(self, query_type: str) -> typing.Self:        
        self.__internal.query_type = query_type
    
        return self
    
    def datasource(self, datasource: object) -> typing.Self:        
        self.__internal.datasource = datasource
    
        return self
    