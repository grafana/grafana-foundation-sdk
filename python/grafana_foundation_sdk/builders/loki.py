# Code generated - EDITING IS FUTILE. DO NOT EDIT.

import typing
from ..cog import builder as cogbuilder
from ..models import loki
from ..models import common
from ..models import dashboardv2beta1


class Dataquery(cogbuilder.Builder[loki.Dataquery]):
    _internal: loki.Dataquery

    def __init__(self) -> None:
        self._internal = loki.Dataquery()

    def build(self) -> loki.Dataquery:
        """
        Builds the object.
        """
        return self._internal    
    
    def expr(self, expr: str) -> typing.Self:    
        """
        The LogQL query.
        """
            
        self._internal.expr = expr
    
        return self
    
    def legend_format(self, legend_format: str) -> typing.Self:    
        """
        Used to override the name of the series.
        """
            
        self._internal.legend_format = legend_format
    
        return self
    
    def max_lines(self, max_lines: int) -> typing.Self:    
        """
        Used to limit the number of log rows returned.
        """
            
        self._internal.max_lines = max_lines
    
        return self
    
    def resolution(self, resolution: int) -> typing.Self:    
        """
        @deprecated, now use step.
        """
            
        self._internal.resolution = resolution
    
        return self
    
    def editor_mode(self, editor_mode: loki.QueryEditorMode) -> typing.Self:    
        self._internal.editor_mode = editor_mode
    
        return self
    
    def range(self, range_val: bool) -> typing.Self:    
        """
        @deprecated, now use queryType.
        """
            
        self._internal.range_val = range_val
    
        return self
    
    def instant(self, instant: bool) -> typing.Self:    
        """
        @deprecated, now use queryType.
        """
            
        self._internal.instant = instant
    
        return self
    
    def step(self, step: str) -> typing.Self:    
        """
        Used to set step value for range queries.
        """
            
        self._internal.step = step
    
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
    
    def direction(self, direction: loki.LokiQueryDirection) -> typing.Self:    
        self._internal.direction = direction
    
        return self
    


class Query(cogbuilder.Builder[dashboardv2beta1.DataQueryKind]):
    _internal: dashboardv2beta1.DataQueryKind

    def __init__(self) -> None:
        self._internal = dashboardv2beta1.DataQueryKind()        
        self._internal.kind = "DataQuery"        
        self._internal.group = "loki"

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
    
    def expr(self, expr: str) -> typing.Self:    
        """
        The LogQL query.
        """
            
        if self._internal.spec is None:
            self._internal.spec = loki.Dataquery()
        assert isinstance(self._internal.spec, loki.Dataquery)
        self._internal.spec.expr = expr
    
        return self
    
    def legend_format(self, legend_format: str) -> typing.Self:    
        """
        Used to override the name of the series.
        """
            
        if self._internal.spec is None:
            self._internal.spec = loki.Dataquery()
        assert isinstance(self._internal.spec, loki.Dataquery)
        self._internal.spec.legend_format = legend_format
    
        return self
    
    def max_lines(self, max_lines: int) -> typing.Self:    
        """
        Used to limit the number of log rows returned.
        """
            
        if self._internal.spec is None:
            self._internal.spec = loki.Dataquery()
        assert isinstance(self._internal.spec, loki.Dataquery)
        self._internal.spec.max_lines = max_lines
    
        return self
    
    def resolution(self, resolution: int) -> typing.Self:    
        """
        @deprecated, now use step.
        """
            
        if self._internal.spec is None:
            self._internal.spec = loki.Dataquery()
        assert isinstance(self._internal.spec, loki.Dataquery)
        self._internal.spec.resolution = resolution
    
        return self
    
    def editor_mode(self, editor_mode: loki.QueryEditorMode) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = loki.Dataquery()
        assert isinstance(self._internal.spec, loki.Dataquery)
        self._internal.spec.editor_mode = editor_mode
    
        return self
    
    def range(self, range_val: bool) -> typing.Self:    
        """
        @deprecated, now use queryType.
        """
            
        if self._internal.spec is None:
            self._internal.spec = loki.Dataquery()
        assert isinstance(self._internal.spec, loki.Dataquery)
        self._internal.spec.range_val = range_val
    
        return self
    
    def instant(self, instant: bool) -> typing.Self:    
        """
        @deprecated, now use queryType.
        """
            
        if self._internal.spec is None:
            self._internal.spec = loki.Dataquery()
        assert isinstance(self._internal.spec, loki.Dataquery)
        self._internal.spec.instant = instant
    
        return self
    
    def step(self, step: str) -> typing.Self:    
        """
        Used to set step value for range queries.
        """
            
        if self._internal.spec is None:
            self._internal.spec = loki.Dataquery()
        assert isinstance(self._internal.spec, loki.Dataquery)
        self._internal.spec.step = step
    
        return self
    
    def ref_id(self, ref_id: str) -> typing.Self:    
        """
        A unique identifier for the query within the list of targets.
        In server side expressions, the refId is used as a variable name to identify results.
        By default, the UI will assign A->Z; however setting meaningful names may be useful.
        """
            
        if self._internal.spec is None:
            self._internal.spec = loki.Dataquery()
        assert isinstance(self._internal.spec, loki.Dataquery)
        self._internal.spec.ref_id = ref_id
    
        return self
    
    def hide(self, hide: bool) -> typing.Self:    
        """
        If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.
        """
            
        if self._internal.spec is None:
            self._internal.spec = loki.Dataquery()
        assert isinstance(self._internal.spec, loki.Dataquery)
        self._internal.spec.hide = hide
    
        return self
    
    def query_type(self, query_type: str) -> typing.Self:    
        """
        Specify the query flavor
        TODO make this required and give it a default
        """
            
        if self._internal.spec is None:
            self._internal.spec = loki.Dataquery()
        assert isinstance(self._internal.spec, loki.Dataquery)
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
            self._internal.spec = loki.Dataquery()
        assert isinstance(self._internal.spec, loki.Dataquery)
        self._internal.spec.datasource = datasource
    
        return self
    
    def direction(self, direction: loki.LokiQueryDirection) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = loki.Dataquery()
        assert isinstance(self._internal.spec, loki.Dataquery)
        self._internal.spec.direction = direction
    
        return self
    
