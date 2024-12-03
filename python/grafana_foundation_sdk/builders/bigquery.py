# Code generated - EDITING IS FUTILE. DO NOT EDIT.

import typing
from ..cog import builder as cogbuilder
from ..models import bigquery
from ..models import dashboard


class SQLExpression(cogbuilder.Builder[bigquery.SQLExpression]):    
    _internal: bigquery.SQLExpression

    def __init__(self):
        self._internal = bigquery.SQLExpression()

    def build(self) -> bigquery.SQLExpression:
        """
        Builds the object.
        """
        return self._internal    
    
    def columns(self, columns: list[cogbuilder.Builder[bigquery.QueryEditorFunctionExpression]]) -> typing.Self:        
        columns_resources = [r1.build() for r1 in columns]
        self._internal.columns = columns_resources
    
        return self
    
    def from_val(self, from_val: str) -> typing.Self:        
        self._internal.from_val = from_val
    
        return self
    
    def where_string(self, where_string: str) -> typing.Self:    
        """
        whereJsonTree?: _
        """
            
        self._internal.where_string = where_string
    
        return self
    
    def group_by(self, group_by: list[cogbuilder.Builder[bigquery.QueryEditorGroupByExpression]]) -> typing.Self:        
        group_by_resources = [r1.build() for r1 in group_by]
        self._internal.group_by = group_by_resources
    
        return self
    
    def order_by(self, order_by: cogbuilder.Builder[bigquery.QueryEditorPropertyExpression]) -> typing.Self:        
        order_by_resource = order_by.build()
        self._internal.order_by = order_by_resource
    
        return self
    
    def order_by_direction(self, order_by_direction: bigquery.OrderByDirection) -> typing.Self:        
        self._internal.order_by_direction = order_by_direction
    
        return self
    
    def limit(self, limit: int) -> typing.Self:        
        self._internal.limit = limit
    
        return self
    
    def offset(self, offset: int) -> typing.Self:        
        self._internal.offset = offset
    
        return self
    

class QueryEditorFunctionExpression(cogbuilder.Builder[bigquery.QueryEditorFunctionExpression]):    
    _internal: bigquery.QueryEditorFunctionExpression

    def __init__(self):
        self._internal = bigquery.QueryEditorFunctionExpression()        
        self._internal.type_val = "function"

    def build(self) -> bigquery.QueryEditorFunctionExpression:
        """
        Builds the object.
        """
        return self._internal    
    
    def name(self, name: str) -> typing.Self:        
        self._internal.name = name
    
        return self
    
    def parameters(self, parameters: list[cogbuilder.Builder[bigquery.QueryEditorFunctionParameterExpression]]) -> typing.Self:        
        parameters_resources = [r1.build() for r1 in parameters]
        self._internal.parameters = parameters_resources
    
        return self
    

class QueryEditorFunctionParameterExpression(cogbuilder.Builder[bigquery.QueryEditorFunctionParameterExpression]):    
    _internal: bigquery.QueryEditorFunctionParameterExpression

    def __init__(self):
        self._internal = bigquery.QueryEditorFunctionParameterExpression()        
        self._internal.type_val = "functionParameter"

    def build(self) -> bigquery.QueryEditorFunctionParameterExpression:
        """
        Builds the object.
        """
        return self._internal    
    
    def name(self, name: str) -> typing.Self:        
        self._internal.name = name
    
        return self
    

class QueryEditorGroupByExpression(cogbuilder.Builder[bigquery.QueryEditorGroupByExpression]):    
    _internal: bigquery.QueryEditorGroupByExpression

    def __init__(self):
        self._internal = bigquery.QueryEditorGroupByExpression()        
        self._internal.type_val = "groupBy"

    def build(self) -> bigquery.QueryEditorGroupByExpression:
        """
        Builds the object.
        """
        return self._internal    
    
    def property_val(self, property_val: cogbuilder.Builder[bigquery.QueryEditorProperty]) -> typing.Self:        
        property_val_resource = property_val.build()
        self._internal.property_val = property_val_resource
    
        return self
    

class QueryEditorProperty(cogbuilder.Builder[bigquery.QueryEditorProperty]):    
    _internal: bigquery.QueryEditorProperty

    def __init__(self):
        self._internal = bigquery.QueryEditorProperty()

    def build(self) -> bigquery.QueryEditorProperty:
        """
        Builds the object.
        """
        return self._internal    
    
    def type_val(self, type_val: bigquery.QueryEditorPropertyType) -> typing.Self:        
        self._internal.type_val = type_val
    
        return self
    
    def name(self, name: str) -> typing.Self:        
        self._internal.name = name
    
        return self
    

class QueryEditorPropertyExpression(cogbuilder.Builder[bigquery.QueryEditorPropertyExpression]):    
    _internal: bigquery.QueryEditorPropertyExpression

    def __init__(self):
        self._internal = bigquery.QueryEditorPropertyExpression()        
        self._internal.type_val = "property"

    def build(self) -> bigquery.QueryEditorPropertyExpression:
        """
        Builds the object.
        """
        return self._internal    
    
    def property_val(self, property_val: cogbuilder.Builder[bigquery.QueryEditorProperty]) -> typing.Self:        
        property_val_resource = property_val.build()
        self._internal.property_val = property_val_resource
    
        return self
    

class Dataquery(cogbuilder.Builder[bigquery.Dataquery]):    
    """
    Manually converted from https://github.com/grafana/google-bigquery-datasource/blob/18680e42ba557791d109c7c540c2c3f2647592f0/src/types.ts#L75
    """
    
    _internal: bigquery.Dataquery

    def __init__(self):
        self._internal = bigquery.Dataquery()

    def build(self) -> bigquery.Dataquery:
        """
        Builds the object.
        """
        return self._internal    
    
    def dataset(self, dataset: str) -> typing.Self:        
        self._internal.dataset = dataset
    
        return self
    
    def table(self, table: str) -> typing.Self:        
        self._internal.table = table
    
        return self
    
    def project(self, project: str) -> typing.Self:        
        self._internal.project = project
    
        return self
    
    def format_val(self, format_val: bigquery.QueryFormat) -> typing.Self:        
        self._internal.format_val = format_val
    
        return self
    
    def raw_query(self, raw_query: bool) -> typing.Self:        
        self._internal.raw_query = raw_query
    
        return self
    
    def raw_sql(self, raw_sql: str) -> typing.Self:        
        self._internal.raw_sql = raw_sql
    
        return self
    
    def location(self, location: str) -> typing.Self:        
        self._internal.location = location
    
        return self
    
    def partitioned(self, partitioned: bool) -> typing.Self:        
        self._internal.partitioned = partitioned
    
        return self
    
    def partitioned_field(self, partitioned_field: str) -> typing.Self:        
        self._internal.partitioned_field = partitioned_field
    
        return self
    
    def convert_to_utc(self, convert_to_utc: bool) -> typing.Self:        
        self._internal.convert_to_utc = convert_to_utc
    
        return self
    
    def sharded(self, sharded: bool) -> typing.Self:        
        self._internal.sharded = sharded
    
        return self
    
    def query_priority(self, query_priority: bigquery.QueryPriority) -> typing.Self:        
        self._internal.query_priority = query_priority
    
        return self
    
    def time_shift(self, time_shift: str) -> typing.Self:        
        self._internal.time_shift = time_shift
    
        return self
    
    def editor_mode(self, editor_mode: bigquery.EditorMode) -> typing.Self:        
        self._internal.editor_mode = editor_mode
    
        return self
    
    def sql(self, sql: cogbuilder.Builder[bigquery.SQLExpression]) -> typing.Self:        
        sql_resource = sql.build()
        self._internal.sql = sql_resource
    
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
    