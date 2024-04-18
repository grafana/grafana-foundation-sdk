# Code generated - EDITING IS FUTILE. DO NOT EDIT.

import typing
from ..cog import builder as cogbuilder
from ..models import tempo


class TempoQuery(cogbuilder.Builder[tempo.TempoQuery]):    
    _internal: tempo.TempoQuery

    def __init__(self):
        self._internal = tempo.TempoQuery()

    def build(self) -> tempo.TempoQuery:
        return self._internal    
    
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
    
    def query(self, query: str) -> typing.Self:    
        """
        TraceQL query or trace ID
        """
            
        self._internal.query = query
    
        return self
    
    def search(self, search: str) -> typing.Self:    
        """
        @deprecated Logfmt query to filter traces by their tags. Example: http.status_code=200 error=true
        """
            
        self._internal.search = search
    
        return self
    
    def service_name(self, service_name: str) -> typing.Self:    
        """
        @deprecated Query traces by service name
        """
            
        self._internal.service_name = service_name
    
        return self
    
    def span_name(self, span_name: str) -> typing.Self:    
        """
        @deprecated Query traces by span name
        """
            
        self._internal.span_name = span_name
    
        return self
    
    def min_duration(self, min_duration: str) -> typing.Self:    
        """
        @deprecated Define the minimum duration to select traces. Use duration format, for example: 1.2s, 100ms
        """
            
        self._internal.min_duration = min_duration
    
        return self
    
    def max_duration(self, max_duration: str) -> typing.Self:    
        """
        @deprecated Define the maximum duration to select traces. Use duration format, for example: 1.2s, 100ms
        """
            
        self._internal.max_duration = max_duration
    
        return self
    
    def service_map_query(self, service_map_query: str) -> typing.Self:    
        """
        Filters to be included in a PromQL query to select data for the service graph. Example: {client="app",service="app"}
        """
            
        self._internal.service_map_query = service_map_query
    
        return self
    
    def service_map_include_namespace(self, service_map_include_namespace: bool) -> typing.Self:    
        """
        Use service.namespace in addition to service.name to uniquely identify a service.
        """
            
        self._internal.service_map_include_namespace = service_map_include_namespace
    
        return self
    
    def limit(self, limit: int) -> typing.Self:    
        """
        Defines the maximum number of traces that are returned from Tempo
        """
            
        self._internal.limit = limit
    
        return self
    
    def spss(self, spss: int) -> typing.Self:    
        """
        Defines the maximum number of spans per spanset that are returned from Tempo
        """
            
        self._internal.spss = spss
    
        return self
    
    def filters(self, filters: list[cogbuilder.Builder[tempo.TraceqlFilter]]) -> typing.Self:        
        filters_resources = [r1.build() for r1 in filters]
        self._internal.filters = filters_resources
    
        return self
    
    def group_by(self, group_by: list[cogbuilder.Builder[tempo.TraceqlFilter]]) -> typing.Self:    
        """
        Filters that are used to query the metrics summary
        """
            
        group_by_resources = [r1.build() for r1 in group_by]
        self._internal.group_by = group_by_resources
    
        return self
    
    def datasource(self, datasource: object) -> typing.Self:    
        """
        For mixed data sources the selected datasource is on the query level.
        For non mixed scenarios this is undefined.
        TODO find a better way to do this ^ that's friendly to schema
        TODO this shouldn't be unknown but DataSourceRef | null
        """
            
        self._internal.datasource = datasource
    
        return self
    
    def table_type(self, table_type: tempo.SearchTableType) -> typing.Self:    
        """
        The type of the table that is used to display the search results
        """
            
        self._internal.table_type = table_type
    
        return self
    

class TraceqlFilter(cogbuilder.Builder[tempo.TraceqlFilter]):    
    _internal: tempo.TraceqlFilter

    def __init__(self):
        self._internal = tempo.TraceqlFilter()

    def build(self) -> tempo.TraceqlFilter:
        return self._internal    
    
    def id_val(self, id_val: str) -> typing.Self:    
        """
        Uniquely identify the filter, will not be used in the query generation
        """
            
        self._internal.id_val = id_val
    
        return self
    
    def tag(self, tag: str) -> typing.Self:    
        """
        The tag for the search filter, for example: .http.status_code, .service.name, status
        """
            
        self._internal.tag = tag
    
        return self
    
    def operator(self, operator: str) -> typing.Self:    
        """
        The operator that connects the tag to the value, for example: =, >, !=, =~
        """
            
        self._internal.operator = operator
    
        return self
    
    def value(self, value: typing.Union[str, list[str]]) -> typing.Self:    
        """
        The value for the search filter
        """
            
        self._internal.value = value
    
        return self
    
    def value_type(self, value_type: str) -> typing.Self:    
        """
        The type of the value, used for example to check whether we need to wrap the value in quotes when generating the query
        """
            
        self._internal.value_type = value_type
    
        return self
    
    def scope(self, scope: tempo.TraceqlSearchScope) -> typing.Self:    
        """
        The scope of the filter, can either be unscoped/all scopes, resource or span
        """
            
        self._internal.scope = scope
    
        return self
    