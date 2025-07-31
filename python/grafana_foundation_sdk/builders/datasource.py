# Code generated - EDITING IS FUTILE. DO NOT EDIT.

import typing
from ..cog import builder as cogbuilder
from ..models import datasource
from ..models import dashboard


class Dataquery(cogbuilder.Builder[datasource.Dataquery]):
    _internal: datasource.Dataquery

    def __init__(self):
        self._internal = datasource.Dataquery()

    def build(self) -> datasource.Dataquery:
        """
        Builds the object.
        """
        return self._internal    
    
    def panel_id(self, panel_id: int) -> typing.Self:    
        """
        Panel ID from wich the queries will be reused.
        """
            
        self._internal.panel_id = panel_id
    
        return self
    
    def ref_id(self, ref_id: str) -> typing.Self:    
        """
        A unique identifier for the query within the list of targets.
        In server side expressions, the refId is used as a variable name to identify results.
        By default, the UI will assign A->Z; however setting meaningful names may be useful.
        """
            
        self._internal.ref_id = ref_id
    
        return self
    
    def hide(self, hide: bool) -> typing.Self:    
        """
        If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.
        """
            
        self._internal.hide = hide
    
        return self
    
    def with_transforms(self, with_transforms: bool) -> typing.Self:    
        self._internal.with_transforms = with_transforms
    
        return self
    
    def datasource(self, datasource: dashboard.DataSourceRef) -> typing.Self:    
        """
        For mixed data sources the selected datasource is on the query level.
        For non mixed scenarios this is undefined.
        TODO find a better way to do this ^ that's friendly to schema
        TODO this shouldn't be unknown but DataSourceRef | null
        """
            
        self._internal.datasource = datasource
    
        return self
    
