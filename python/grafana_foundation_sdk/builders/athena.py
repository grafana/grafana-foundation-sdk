# Code generated - EDITING IS FUTILE. DO NOT EDIT.

import typing
from ..cog import builder as cogbuilder
from ..models import athena
from ..models import dashboard


class Dataquery(cogbuilder.Builder[athena.Dataquery]):    
    """
    Manually converted from https://github.com/grafana/athena-datasource/blob/57ad707147b7a11e9a521a836d6bf9799877e0e3/src/types.ts
    """
    
    _internal: athena.Dataquery

    def __init__(self):
        self._internal = athena.Dataquery()

    def build(self) -> athena.Dataquery:
        """
        Builds the object.
        """
        return self._internal    
    
    def format_val(self, format_val: athena.FormatOptions) -> typing.Self:        
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
    
    def raw_sql(self, raw_sql: str) -> typing.Self:        
        self._internal.raw_sql = raw_sql
    
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
    

class ConnectionArgs(cogbuilder.Builder[athena.ConnectionArgs]):    
    _internal: athena.ConnectionArgs

    def __init__(self):
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
    