# Code generated - EDITING IS FUTILE. DO NOT EDIT.

import typing
from ..cog import builder as cogbuilder
from ..models import athena
from ..models import common
from ..models import dashboardv2beta1


class Dataquery(cogbuilder.Builder[athena.Dataquery]):    
    """
    Manually converted from https://github.com/grafana/athena-datasource/blob/57ad707147b7a11e9a521a836d6bf9799877e0e3/src/types.ts
    """
    
    _internal: athena.Dataquery

    def __init__(self) -> None:
        self._internal = athena.Dataquery()

    def build(self) -> athena.Dataquery:
        """
        Builds the object.
        """
        return self._internal    
    
    def format(self, format_val: athena.FormatOptions) -> typing.Self:    
        self._internal.format_val = format_val
    
        return self
    
    def connection_args(self, connection_args: cogbuilder.Builder[athena.ConnectionArgs]) -> typing.Self:    
        connection_args_resource = connection_args.build()
        self._internal.connection_args = connection_args_resource
    
        return self
    
    def table(self, table: str) -> typing.Self:    
        self._internal.table = table
    
        return self
    
    def column(self, column: str) -> typing.Self:    
        self._internal.column = column
    
        return self
    
    def query_id(self, query_id: str) -> typing.Self:    
        self._internal.query_id = query_id
    
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
    
    def raw_sql(self, raw_sql: str) -> typing.Self:    
        self._internal.raw_sql = raw_sql
    
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
    


class ConnectionArgs(cogbuilder.Builder[athena.ConnectionArgs]):
    _internal: athena.ConnectionArgs

    def __init__(self) -> None:
        self._internal = athena.ConnectionArgs()

    def build(self) -> athena.ConnectionArgs:
        """
        Builds the object.
        """
        return self._internal    
    
    def region(self, region: str) -> typing.Self:    
        self._internal.region = region
    
        return self
    
    def catalog(self, catalog: str) -> typing.Self:    
        self._internal.catalog = catalog
    
        return self
    
    def database(self, database: str) -> typing.Self:    
        self._internal.database = database
    
        return self
    
    def result_reuse_enabled(self, result_reuse_enabled: bool) -> typing.Self:    
        self._internal.result_reuse_enabled = result_reuse_enabled
    
        return self
    
    def result_reuse_max_age_in_minutes(self, result_reuse_max_age_in_minutes: float) -> typing.Self:    
        self._internal.result_reuse_max_age_in_minutes = result_reuse_max_age_in_minutes
    
        return self
    


class Query(cogbuilder.Builder[dashboardv2beta1.DataQueryKind]):
    _internal: dashboardv2beta1.DataQueryKind

    def __init__(self) -> None:
        self._internal = dashboardv2beta1.DataQueryKind()        
        self._internal.kind = "DataQuery"        
        self._internal.group = "grafana-athena-datasource"

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
    
    def format(self, format_val: athena.FormatOptions) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = athena.Dataquery()
        assert isinstance(self._internal.spec, athena.Dataquery)
        self._internal.spec.format_val = format_val
    
        return self
    
    def connection_args(self, connection_args: cogbuilder.Builder[athena.ConnectionArgs]) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = athena.Dataquery()
        assert isinstance(self._internal.spec, athena.Dataquery)
        connection_args_resource = connection_args.build()
        self._internal.spec.connection_args = connection_args_resource
    
        return self
    
    def table(self, table: str) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = athena.Dataquery()
        assert isinstance(self._internal.spec, athena.Dataquery)
        self._internal.spec.table = table
    
        return self
    
    def column(self, column: str) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = athena.Dataquery()
        assert isinstance(self._internal.spec, athena.Dataquery)
        self._internal.spec.column = column
    
        return self
    
    def query_id(self, query_id: str) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = athena.Dataquery()
        assert isinstance(self._internal.spec, athena.Dataquery)
        self._internal.spec.query_id = query_id
    
        return self
    
    def ref_id(self, ref_id: str) -> typing.Self:    
        """
        A unique identifier for the query within the list of targets.
        In server side expressions, the refId is used as a variable name to identify results.
        By default, the UI will assign A->Z; however setting meaningful names may be useful.
        """
            
        if self._internal.spec is None:
            self._internal.spec = athena.Dataquery()
        assert isinstance(self._internal.spec, athena.Dataquery)
        self._internal.spec.ref_id = ref_id
    
        return self
    
    def hide(self, hide: bool) -> typing.Self:    
        """
        If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.
        """
            
        if self._internal.spec is None:
            self._internal.spec = athena.Dataquery()
        assert isinstance(self._internal.spec, athena.Dataquery)
        self._internal.spec.hide = hide
    
        return self
    
    def query_type(self, query_type: str) -> typing.Self:    
        """
        Specify the query flavor
        TODO make this required and give it a default
        """
            
        if self._internal.spec is None:
            self._internal.spec = athena.Dataquery()
        assert isinstance(self._internal.spec, athena.Dataquery)
        self._internal.spec.query_type = query_type
    
        return self
    
    def raw_sql(self, raw_sql: str) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = athena.Dataquery()
        assert isinstance(self._internal.spec, athena.Dataquery)
        self._internal.spec.raw_sql = raw_sql
    
        return self
    
    def old_datasource(self, datasource: common.DataSourceRef) -> typing.Self:    
        """
        For mixed data sources the selected datasource is on the query level.
        For non mixed scenarios this is undefined.
        TODO find a better way to do this ^ that's friendly to schema
        TODO this shouldn't be unknown but DataSourceRef | null
        """
            
        if self._internal.spec is None:
            self._internal.spec = athena.Dataquery()
        assert isinstance(self._internal.spec, athena.Dataquery)
        self._internal.spec.datasource = datasource
    
        return self
    
