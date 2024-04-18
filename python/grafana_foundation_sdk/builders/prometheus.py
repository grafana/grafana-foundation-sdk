# Code generated - EDITING IS FUTILE. DO NOT EDIT.

import typing
from ..cog import builder as cogbuilder
from ..models import prometheus


class Dataquery(cogbuilder.Builder[prometheus.Dataquery]):    
    _internal: prometheus.Dataquery

    def __init__(self):
        self._internal = prometheus.Dataquery()

    def build(self) -> prometheus.Dataquery:
        return self._internal    
    
    def expr(self, expr: str) -> typing.Self:        
        self._internal.expr = expr
    
        return self
    
    def instant(self, instant: bool) -> typing.Self:        
        self._internal.instant = instant
    
        return self
    
    def range_val(self, range_val: bool) -> typing.Self:        
        self._internal.range_val = range_val
    
        return self
    
    def exemplar(self, exemplar: bool) -> typing.Self:        
        self._internal.exemplar = exemplar
    
        return self
    
    def editor_mode(self, editor_mode: prometheus.QueryEditorMode) -> typing.Self:        
        self._internal.editor_mode = editor_mode
    
        return self
    
    def format_val(self, format_val: prometheus.PromQueryFormat) -> typing.Self:        
        self._internal.format_val = format_val
    
        return self
    
    def legend_format(self, legend_format: str) -> typing.Self:        
        self._internal.legend_format = legend_format
    
        return self
    
    def interval_factor(self, interval_factor: float) -> typing.Self:        
        self._internal.interval_factor = interval_factor
    
        return self
    
    def scope(self, scope: cogbuilder.Builder[prometheus.PrometheusDataqueryScope]) -> typing.Self:        
        scope_resource = scope.build()
        self._internal.scope = scope_resource
    
        return self
    
    def ref_id(self, ref_id: str) -> typing.Self:        
        self._internal.ref_id = ref_id
    
        return self
    
    def hide(self, hide: bool) -> typing.Self:        
        self._internal.hide = hide
    
        return self
    
    def query_type(self, query_type: str) -> typing.Self:        
        self._internal.query_type = query_type
    
        return self
    
    def datasource(self, datasource: object) -> typing.Self:        
        self._internal.datasource = datasource
    
        return self
    
    def interval(self, interval: str) -> typing.Self:    
        """
        An additional lower limit for the step parameter of the Prometheus query and for the
        `$__interval` and `$__rate_interval` variables.
        """
            
        self._internal.interval = interval
    
        return self
    

class PrometheusDataqueryScope(cogbuilder.Builder[prometheus.PrometheusDataqueryScope]):    
    _internal: prometheus.PrometheusDataqueryScope

    def __init__(self):
        self._internal = prometheus.PrometheusDataqueryScope()

    def build(self) -> prometheus.PrometheusDataqueryScope:
        return self._internal    
    
    def matchers(self, matchers: str) -> typing.Self:        
        self._internal.matchers = matchers
    
        return self
    