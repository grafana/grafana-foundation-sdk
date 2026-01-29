# Code generated - EDITING IS FUTILE. DO NOT EDIT.

import typing
from ..cog import builder as cogbuilder
from ..models import prometheus
from ..models import common
from ..models import dashboardv2beta1


class Dataquery(cogbuilder.Builder[prometheus.Dataquery]):
    _internal: prometheus.Dataquery

    def __init__(self) -> None:
        self._internal = prometheus.Dataquery()

    def build(self) -> prometheus.Dataquery:
        """
        Builds the object.
        """
        return self._internal    
    
    def expr(self, expr: str) -> typing.Self:    
        """
        The actual expression/query that will be evaluated by Prometheus
        """
            
        self._internal.expr = expr
    
        return self
    
    def instant(self) -> typing.Self:    
        """
        Returns only the latest value that Prometheus has scraped for the requested time series
        """
            
        self._internal.instant = True    
        self._internal.range_val = False
    
        return self
    
    def range(self) -> typing.Self:    
        """
        Returns a Range vector, comprised of a set of time series containing a range of data points over time for each time series
        """
            
        self._internal.range_val = True    
        self._internal.instant = False
    
        return self
    
    def exemplar(self, exemplar: bool) -> typing.Self:    
        """
        Execute an additional query to identify interesting raw samples relevant for the given expr
        """
            
        self._internal.exemplar = exemplar
    
        return self
    
    def editor_mode(self, editor_mode: prometheus.QueryEditorMode) -> typing.Self:    
        """
        Specifies which editor is being used to prepare the query. It can be "code" or "builder"
        """
            
        self._internal.editor_mode = editor_mode
    
        return self
    
    def format(self, format_val: prometheus.PromQueryFormat) -> typing.Self:    
        """
        Query format to determine how to display data points in panel. It can be "time_series", "table", "heatmap"
        """
            
        self._internal.format_val = format_val
    
        return self
    
    def legend_format(self, legend_format: str) -> typing.Self:    
        """
        Series name override or template. Ex. {{hostname}} will be replaced with label value for hostname
        """
            
        self._internal.legend_format = legend_format
    
        return self
    
    def interval_factor(self, interval_factor: float) -> typing.Self:    
        """
        @deprecated Used to specify how many times to divide max data points by. We use max data points under query options
        See https://github.com/grafana/grafana/issues/48081
        """
            
        self._internal.interval_factor = interval_factor
    
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
    
    def interval(self, interval: str) -> typing.Self:    
        """
        An additional lower limit for the step parameter of the Prometheus query and for the
        `$__interval` and `$__rate_interval` variables.
        """
            
        self._internal.interval = interval
    
        return self
    
    def range_and_instant(self) -> typing.Self:    
        self._internal.range_val = True    
        self._internal.instant = True
    
        return self
    


class Query(cogbuilder.Builder[dashboardv2beta1.DataQueryKind]):
    _internal: dashboardv2beta1.DataQueryKind

    def __init__(self) -> None:
        self._internal = dashboardv2beta1.DataQueryKind()        
        self._internal.kind = "DataQuery"        
        self._internal.group = "prometheus"

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
        The actual expression/query that will be evaluated by Prometheus
        """
            
        if self._internal.spec is None:
            self._internal.spec = prometheus.Dataquery()
        assert isinstance(self._internal.spec, prometheus.Dataquery)
        self._internal.spec.expr = expr
    
        return self
    
    def instant(self) -> typing.Self:    
        """
        Returns only the latest value that Prometheus has scraped for the requested time series
        """
            
        if self._internal.spec is None:
            self._internal.spec = prometheus.Dataquery()
        assert isinstance(self._internal.spec, prometheus.Dataquery)
        self._internal.spec.instant = True
    
        return self
    
    def range(self) -> typing.Self:    
        """
        Returns a Range vector, comprised of a set of time series containing a range of data points over time for each time series
        """
            
        if self._internal.spec is None:
            self._internal.spec = prometheus.Dataquery()
        assert isinstance(self._internal.spec, prometheus.Dataquery)
        self._internal.spec.range_val = True
    
        return self
    
    def exemplar(self, exemplar: bool) -> typing.Self:    
        """
        Execute an additional query to identify interesting raw samples relevant for the given expr
        """
            
        if self._internal.spec is None:
            self._internal.spec = prometheus.Dataquery()
        assert isinstance(self._internal.spec, prometheus.Dataquery)
        self._internal.spec.exemplar = exemplar
    
        return self
    
    def editor_mode(self, editor_mode: prometheus.QueryEditorMode) -> typing.Self:    
        """
        Specifies which editor is being used to prepare the query. It can be "code" or "builder"
        """
            
        if self._internal.spec is None:
            self._internal.spec = prometheus.Dataquery()
        assert isinstance(self._internal.spec, prometheus.Dataquery)
        self._internal.spec.editor_mode = editor_mode
    
        return self
    
    def format(self, format_val: prometheus.PromQueryFormat) -> typing.Self:    
        """
        Query format to determine how to display data points in panel. It can be "time_series", "table", "heatmap"
        """
            
        if self._internal.spec is None:
            self._internal.spec = prometheus.Dataquery()
        assert isinstance(self._internal.spec, prometheus.Dataquery)
        self._internal.spec.format_val = format_val
    
        return self
    
    def legend_format(self, legend_format: str) -> typing.Self:    
        """
        Series name override or template. Ex. {{hostname}} will be replaced with label value for hostname
        """
            
        if self._internal.spec is None:
            self._internal.spec = prometheus.Dataquery()
        assert isinstance(self._internal.spec, prometheus.Dataquery)
        self._internal.spec.legend_format = legend_format
    
        return self
    
    def interval_factor(self, interval_factor: float) -> typing.Self:    
        """
        @deprecated Used to specify how many times to divide max data points by. We use max data points under query options
        See https://github.com/grafana/grafana/issues/48081
        """
            
        if self._internal.spec is None:
            self._internal.spec = prometheus.Dataquery()
        assert isinstance(self._internal.spec, prometheus.Dataquery)
        self._internal.spec.interval_factor = interval_factor
    
        return self
    
    def ref_id(self, ref_id: str) -> typing.Self:    
        """
        A unique identifier for the query within the list of targets.
        In server side expressions, the refId is used as a variable name to identify results.
        By default, the UI will assign A->Z; however setting meaningful names may be useful.
        """
            
        if self._internal.spec is None:
            self._internal.spec = prometheus.Dataquery()
        assert isinstance(self._internal.spec, prometheus.Dataquery)
        self._internal.spec.ref_id = ref_id
    
        return self
    
    def hide(self, hide: bool) -> typing.Self:    
        """
        If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.
        """
            
        if self._internal.spec is None:
            self._internal.spec = prometheus.Dataquery()
        assert isinstance(self._internal.spec, prometheus.Dataquery)
        self._internal.spec.hide = hide
    
        return self
    
    def query_type(self, query_type: str) -> typing.Self:    
        """
        Specify the query flavor
        TODO make this required and give it a default
        """
            
        if self._internal.spec is None:
            self._internal.spec = prometheus.Dataquery()
        assert isinstance(self._internal.spec, prometheus.Dataquery)
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
            self._internal.spec = prometheus.Dataquery()
        assert isinstance(self._internal.spec, prometheus.Dataquery)
        self._internal.spec.datasource = datasource
    
        return self
    
    def interval(self, interval: str) -> typing.Self:    
        """
        An additional lower limit for the step parameter of the Prometheus query and for the
        `$__interval` and `$__rate_interval` variables.
        """
            
        if self._internal.spec is None:
            self._internal.spec = prometheus.Dataquery()
        assert isinstance(self._internal.spec, prometheus.Dataquery)
        self._internal.spec.interval = interval
    
        return self
    
    def range_and_instant(self) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = prometheus.Dataquery()
        assert isinstance(self._internal.spec, prometheus.Dataquery)
        self._internal.spec.range_val = True    
        self._internal.spec.instant = True
    
        return self
    
