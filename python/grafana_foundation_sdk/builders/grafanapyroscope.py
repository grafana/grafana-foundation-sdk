# Code generated - EDITING IS FUTILE. DO NOT EDIT.

import typing
from ..cog import builder as cogbuilder
from ..models import grafanapyroscope
from ..models import dashboard


class Dataquery(cogbuilder.Builder[grafanapyroscope.Dataquery]):    
    _internal: grafanapyroscope.Dataquery

    def __init__(self):
        self._internal = grafanapyroscope.Dataquery()

    def build(self) -> grafanapyroscope.Dataquery:
        """
        Builds the object.
        """
        return self._internal    
    
    def label_selector(self, label_selector: str) -> typing.Self:    
        """
        Specifies the query label selectors.
        """
            
        self._internal.label_selector = label_selector
    
        return self
    
    def span_selector(self, span_selector: list[str]) -> typing.Self:    
        """
        Specifies the query span selectors.
        """
            
        self._internal.span_selector = span_selector
    
        return self
    
    def profile_type_id(self, profile_type_id: str) -> typing.Self:    
        """
        Specifies the type of profile to query.
        """
            
        self._internal.profile_type_id = profile_type_id
    
        return self
    
    def group_by(self, group_by: list[str]) -> typing.Self:    
        """
        Allows to group the results.
        """
            
        self._internal.group_by = group_by
    
        return self
    
    def max_nodes(self, max_nodes: int) -> typing.Self:    
        """
        Sets the maximum number of nodes in the flamegraph.
        """
            
        self._internal.max_nodes = max_nodes
    
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
        true if query is disabled (ie should not be returned to the dashboard)
        Note this does not always imply that the query should not be executed since
        the results from a hidden query may be used as the input to other queries (SSE etc)
        """
            
        self._internal.hide = hide
    
        return self
    
    def query_type(self, query_type: str) -> typing.Self:    
        """
        Specify the query flavor
        TODO make this required and give it a default
        """
            
        self._internal.query_type = query_type
    
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
    