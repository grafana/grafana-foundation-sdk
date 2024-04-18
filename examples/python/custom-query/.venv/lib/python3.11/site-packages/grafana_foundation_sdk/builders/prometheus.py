# Code generated - EDITING IS FUTILE. DO NOT EDIT.

import typing
from ..cog import builder as cogbuilder
from ..models import prometheus


class Dataquery(cogbuilder.Builder[prometheus.Dataquery]):    
    __internal: prometheus.Dataquery

    def __init__(self):
        self.__internal = prometheus.Dataquery()

    def build(self) -> prometheus.Dataquery:
        return self.__internal    
    
    def expr(self, expr: str) -> typing.Self:        
        self.__internal.expr = expr
    
        return self
    
    def instant(self, instant: bool) -> typing.Self:        
        self.__internal.instant = instant
    
        return self
    
    def range_val(self, range_val: bool) -> typing.Self:        
        self.__internal.range_val = range_val
    
        return self
    
    def exemplar(self, exemplar: bool) -> typing.Self:        
        self.__internal.exemplar = exemplar
    
        return self
    
    def editor_mode(self, editor_mode: prometheus.QueryEditorMode) -> typing.Self:        
        self.__internal.editor_mode = editor_mode
    
        return self
    
    def format_val(self, format_val: prometheus.PromQueryFormat) -> typing.Self:        
        self.__internal.format_val = format_val
    
        return self
    
    def legend_format(self, legend_format: str) -> typing.Self:        
        self.__internal.legend_format = legend_format
    
        return self
    
    def interval_factor(self, interval_factor: float) -> typing.Self:        
        self.__internal.interval_factor = interval_factor
    
        return self
    
    def scope(self, scope: cogbuilder.Builder[prometheus.PrometheusDataqueryScope]) -> typing.Self:        
        scope_resource = scope.build()
        self.__internal.scope = scope_resource
    
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
    
    def interval(self, interval: str) -> typing.Self:    
        """
        An additional lower limit for the step parameter of the Prometheus query and for the
        `$__interval` and `$__rate_interval` variables.
        """
            
        self.__internal.interval = interval
    
        return self
    

class PrometheusDataqueryScope(cogbuilder.Builder[prometheus.PrometheusDataqueryScope]):    
    __internal: prometheus.PrometheusDataqueryScope

    def __init__(self):
        self.__internal = prometheus.PrometheusDataqueryScope()

    def build(self) -> prometheus.PrometheusDataqueryScope:
        return self.__internal    
    
    def matchers(self, matchers: str) -> typing.Self:        
        self.__internal.matchers = matchers
    
        return self
    