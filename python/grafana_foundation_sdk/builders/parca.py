# Code generated - EDITING IS FUTILE. DO NOT EDIT.

import typing
from ..cog import builder as cogbuilder
from ..models import parca
from ..models import common
from ..models import dashboardv2beta1


class Dataquery(cogbuilder.Builder[parca.Dataquery]):
    _internal: parca.Dataquery

    def __init__(self) -> None:
        self._internal = parca.Dataquery()

    def build(self) -> parca.Dataquery:
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
    
    def profile_type_id(self, profile_type_id: str) -> typing.Self:    
        """
        Specifies the type of profile to query.
        """
            
        self._internal.profile_type_id = profile_type_id
    
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
    
    def query_type(self, query_type: str) -> typing.Self:    
        """
        Specify the query flavor
        TODO make this required and give it a default
        """
            
        self._internal.query_type = query_type
    
        return self
    
    def datasource(self, datasource: common.DataSourceRef) -> typing.Self:    
        """
        For mixed data sources the selected datasource is on the query level.
        For non mixed scenarios this is undefined.
        TODO find a better way to do this ^ that's friendly to schema
        TODO this shouldn't be unknown but DataSourceRef | null
        """
            
        self._internal.datasource = datasource
    
        return self
    


class Query(cogbuilder.Builder[dashboardv2beta1.DataQueryKind]):
    _internal: dashboardv2beta1.DataQueryKind

    def __init__(self) -> None:
        self._internal = dashboardv2beta1.DataQueryKind()        
        self._internal.kind = "DataQuery"        
        self._internal.group = "parca"

    def build(self) -> dashboardv2beta1.DataQueryKind:
        """
        Builds the object.
        """
        return self._internal    
    
    def version(self, version: str) -> typing.Self:    
        self._internal.version = version
    
        return self
    
    def datasource(self, ref: cogbuilder.Builder[dashboardv2beta1.Dashboardv2beta1DataQueryKindDatasource]) -> typing.Self:    
        """
        New type for datasource reference
        Not creating a new type until we figure out how to handle DS refs for group by, adhoc, and every place that uses DataSourceRef in TS.
        """
            
        ref_resource = ref.build()
        self._internal.datasource = ref_resource
    
        return self
    
    def label_selector(self, label_selector: str) -> typing.Self:    
        """
        Specifies the query label selectors.
        """
            
        if self._internal.spec is None:
            self._internal.spec = parca.Dataquery()
        assert isinstance(self._internal.spec, parca.Dataquery)
        self._internal.spec.label_selector = label_selector
    
        return self
    
    def profile_type_id(self, profile_type_id: str) -> typing.Self:    
        """
        Specifies the type of profile to query.
        """
            
        if self._internal.spec is None:
            self._internal.spec = parca.Dataquery()
        assert isinstance(self._internal.spec, parca.Dataquery)
        self._internal.spec.profile_type_id = profile_type_id
    
        return self
    
    def ref_id(self, ref_id: str) -> typing.Self:    
        """
        A unique identifier for the query within the list of targets.
        In server side expressions, the refId is used as a variable name to identify results.
        By default, the UI will assign A->Z; however setting meaningful names may be useful.
        """
            
        if self._internal.spec is None:
            self._internal.spec = parca.Dataquery()
        assert isinstance(self._internal.spec, parca.Dataquery)
        self._internal.spec.ref_id = ref_id
    
        return self
    
    def hide(self, hide: bool) -> typing.Self:    
        """
        If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.
        """
            
        if self._internal.spec is None:
            self._internal.spec = parca.Dataquery()
        assert isinstance(self._internal.spec, parca.Dataquery)
        self._internal.spec.hide = hide
    
        return self
    
    def query_type(self, query_type: str) -> typing.Self:    
        """
        Specify the query flavor
        TODO make this required and give it a default
        """
            
        if self._internal.spec is None:
            self._internal.spec = parca.Dataquery()
        assert isinstance(self._internal.spec, parca.Dataquery)
        self._internal.spec.query_type = query_type
    
        return self
    
    def old_datasource(self, datasource: common.DataSourceRef) -> typing.Self:    
        """
        For mixed data sources the selected datasource is on the query level.
        For non mixed scenarios this is undefined.
        TODO find a better way to do this ^ that's friendly to schema
        TODO this shouldn't be unknown but DataSourceRef | null
        """
            
        if self._internal.spec is None:
            self._internal.spec = parca.Dataquery()
        assert isinstance(self._internal.spec, parca.Dataquery)
        self._internal.spec.datasource = datasource
    
        return self
    
