# Code generated - EDITING IS FUTILE. DO NOT EDIT.

import typing
from ..cog import builder as cogbuilder
from ..models import loki


class Dataquery(cogbuilder.Builder[loki.Dataquery]):    
    _internal: loki.Dataquery

    def __init__(self):
        self._internal = loki.Dataquery()

    def build(self) -> loki.Dataquery:
        return self._internal    
    
    def expr(self, expr: str) -> typing.Self:        
        self._internal.expr = expr
    
        return self
    
    def legend_format(self, legend_format: str) -> typing.Self:        
        self._internal.legend_format = legend_format
    
        return self
    
    def max_lines(self, max_lines: int) -> typing.Self:        
        self._internal.max_lines = max_lines
    
        return self
    
    def resolution(self, resolution: int) -> typing.Self:        
        self._internal.resolution = resolution
    
        return self
    
    def editor_mode(self, editor_mode: loki.QueryEditorMode) -> typing.Self:        
        self._internal.editor_mode = editor_mode
    
        return self
    
    def range_val(self, range_val: bool) -> typing.Self:        
        self._internal.range_val = range_val
    
        return self
    
    def instant(self, instant: bool) -> typing.Self:        
        self._internal.instant = instant
    
        return self
    
    def step(self, step: str) -> typing.Self:        
        self._internal.step = step
    
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
    