# Code generated - EDITING IS FUTILE. DO NOT EDIT.

import typing
from ..cog import builder as cogbuilder
from ..models import azuremonitor
from ..models import dashboard


class AzureMonitorQuery(cogbuilder.Builder[azuremonitor.AzureMonitorQuery]):
    _internal: azuremonitor.AzureMonitorQuery

    def __init__(self):
        self._internal = azuremonitor.AzureMonitorQuery()

    def build(self) -> azuremonitor.AzureMonitorQuery:
        """
        Builds the object.
        """
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
    
    def subscription(self, subscription: str) -> typing.Self:    
        """
        Azure subscription containing the resource(s) to be queried.
        Also used for template variable queries
        """
            
        self._internal.subscription = subscription
    
        return self
    
    def subscriptions(self, subscriptions: list[str]) -> typing.Self:    
        """
        Subscriptions to be queried via Azure Resource Graph.
        """
            
        self._internal.subscriptions = subscriptions
    
        return self
    
    def azure_monitor(self, azure_monitor: cogbuilder.Builder[azuremonitor.AzureMetricQuery]) -> typing.Self:    
        """
        Azure Monitor Metrics sub-query properties.
        """
            
        azure_monitor_resource = azure_monitor.build()
        self._internal.azure_monitor = azure_monitor_resource
    
        return self
    
    def azure_log_analytics(self, azure_log_analytics: cogbuilder.Builder[azuremonitor.AzureLogsQuery]) -> typing.Self:    
        """
        Azure Monitor Logs sub-query properties.
        """
            
        azure_log_analytics_resource = azure_log_analytics.build()
        self._internal.azure_log_analytics = azure_log_analytics_resource
    
        return self
    
    def azure_resource_graph(self, azure_resource_graph: cogbuilder.Builder[azuremonitor.AzureResourceGraphQuery]) -> typing.Self:    
        """
        Azure Resource Graph sub-query properties.
        """
            
        azure_resource_graph_resource = azure_resource_graph.build()
        self._internal.azure_resource_graph = azure_resource_graph_resource
    
        return self
    
    def azure_traces(self, azure_traces: cogbuilder.Builder[azuremonitor.AzureTracesQuery]) -> typing.Self:    
        """
        Application Insights Traces sub-query properties.
        """
            
        azure_traces_resource = azure_traces.build()
        self._internal.azure_traces = azure_traces_resource
    
        return self
    
    def grafana_template_variable_fn(self, grafana_template_variable_fn: cogbuilder.Builder[azuremonitor.GrafanaTemplateVariableQuery]) -> typing.Self:    
        """
        @deprecated Legacy template variable support.
        """
            
        grafana_template_variable_fn_resource = grafana_template_variable_fn.build()
        self._internal.grafana_template_variable_fn = grafana_template_variable_fn_resource
    
        return self
    
    def resource_group(self, resource_group: str) -> typing.Self:    
        """
        Resource group used in template variable queries
        """
            
        self._internal.resource_group = resource_group
    
        return self
    
    def namespace(self, namespace: str) -> typing.Self:    
        """
        Namespace used in template variable queries
        """
            
        self._internal.namespace = namespace
    
        return self
    
    def resource(self, resource: str) -> typing.Self:    
        """
        Resource used in template variable queries
        """
            
        self._internal.resource = resource
    
        return self
    
    def region(self, region: str) -> typing.Self:    
        """
        Region used in template variable queries
        """
            
        self._internal.region = region
    
        return self
    
    def custom_namespace(self, custom_namespace: str) -> typing.Self:    
        """
        Custom namespace used in template variable queries
        """
            
        self._internal.custom_namespace = custom_namespace
    
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
    
    def query(self, query: str) -> typing.Self:    
        """
        Used only for exemplar queries from Prometheus
        """
            
        self._internal.query = query
    
        return self
    


class AzureMetricQuery(cogbuilder.Builder[azuremonitor.AzureMetricQuery]):
    _internal: azuremonitor.AzureMetricQuery

    def __init__(self):
        self._internal = azuremonitor.AzureMetricQuery()

    def build(self) -> azuremonitor.AzureMetricQuery:
        """
        Builds the object.
        """
        return self._internal    
    
    def resources(self, resources: list[cogbuilder.Builder[azuremonitor.AzureMonitorResource]]) -> typing.Self:    
        """
        Array of resource URIs to be queried.
        """
            
        resources_resources = [r1.build() for r1 in resources]
        self._internal.resources = resources_resources
    
        return self
    
    def metric_namespace(self, metric_namespace: str) -> typing.Self:    
        """
        metricNamespace is used as the resource type (or resource namespace).
        It's usually equal to the target metric namespace. e.g. microsoft.storage/storageaccounts
        Kept the name of the variable as metricNamespace to avoid backward incompatibility issues.
        """
            
        self._internal.metric_namespace = metric_namespace
    
        return self
    
    def custom_namespace(self, custom_namespace: str) -> typing.Self:    
        """
        Used as the value for the metricNamespace property when it's different from the resource namespace.
        """
            
        self._internal.custom_namespace = custom_namespace
    
        return self
    
    def metric_name(self, metric_name: str) -> typing.Self:    
        """
        The metric to query data for within the specified metricNamespace. e.g. UsedCapacity
        """
            
        self._internal.metric_name = metric_name
    
        return self
    
    def region(self, region: str) -> typing.Self:    
        """
        The Azure region containing the resource(s).
        """
            
        self._internal.region = region
    
        return self
    
    def time_grain(self, time_grain: str) -> typing.Self:    
        """
        The granularity of data points to be queried. Defaults to auto.
        """
            
        self._internal.time_grain = time_grain
    
        return self
    
    def aggregation(self, aggregation: str) -> typing.Self:    
        """
        The aggregation to be used within the query. Defaults to the primaryAggregationType defined by the metric.
        """
            
        self._internal.aggregation = aggregation
    
        return self
    
    def dimension_filters(self, dimension_filters: list[cogbuilder.Builder[azuremonitor.AzureMetricDimension]]) -> typing.Self:    
        """
        Filters to reduce the set of data returned. Dimensions that can be filtered on are defined by the metric.
        """
            
        dimension_filters_resources = [r1.build() for r1 in dimension_filters]
        self._internal.dimension_filters = dimension_filters_resources
    
        return self
    
    def top(self, top: str) -> typing.Self:    
        """
        Maximum number of records to return. Defaults to 10.
        """
            
        self._internal.top = top
    
        return self
    
    def allowed_time_grains_ms(self, allowed_time_grains_ms: list[int]) -> typing.Self:    
        """
        Time grains that are supported by the metric.
        """
            
        self._internal.allowed_time_grains_ms = allowed_time_grains_ms
    
        return self
    
    def alias(self, alias: str) -> typing.Self:    
        """
        Aliases can be set to modify the legend labels. e.g. {{ resourceGroup }}. See docs for more detail.
        """
            
        self._internal.alias = alias
    
        return self
    
    def time_grain_unit(self, time_grain_unit: str) -> typing.Self:    
        """
        @deprecated
        """
            
        self._internal.time_grain_unit = time_grain_unit
    
        return self
    
    def dimension(self, dimension: str) -> typing.Self:    
        """
        @deprecated This property was migrated to dimensionFilters and should only be accessed in the migration
        """
            
        self._internal.dimension = dimension
    
        return self
    
    def dimension_filter(self, dimension_filter: str) -> typing.Self:    
        """
        @deprecated This property was migrated to dimensionFilters and should only be accessed in the migration
        """
            
        self._internal.dimension_filter = dimension_filter
    
        return self
    
    def metric_definition(self, metric_definition: str) -> typing.Self:    
        """
        @deprecated Use metricNamespace instead
        """
            
        self._internal.metric_definition = metric_definition
    
        return self
    
    def resource_uri(self, resource_uri: str) -> typing.Self:    
        """
        @deprecated Use resourceGroup, resourceName and metricNamespace instead
        """
            
        self._internal.resource_uri = resource_uri
    
        return self
    
    def resource_group(self, resource_group: str) -> typing.Self:    
        """
        @deprecated Use resources instead
        """
            
        self._internal.resource_group = resource_group
    
        return self
    
    def resource_name(self, resource_name: str) -> typing.Self:    
        """
        @deprecated Use resources instead
        """
            
        self._internal.resource_name = resource_name
    
        return self
    


class AzureMonitorResource(cogbuilder.Builder[azuremonitor.AzureMonitorResource]):
    _internal: azuremonitor.AzureMonitorResource

    def __init__(self):
        self._internal = azuremonitor.AzureMonitorResource()

    def build(self) -> azuremonitor.AzureMonitorResource:
        """
        Builds the object.
        """
        return self._internal    
    
    def subscription(self, subscription: str) -> typing.Self:    
        self._internal.subscription = subscription
    
        return self
    
    def resource_group(self, resource_group: str) -> typing.Self:    
        self._internal.resource_group = resource_group
    
        return self
    
    def resource_name(self, resource_name: str) -> typing.Self:    
        self._internal.resource_name = resource_name
    
        return self
    
    def metric_namespace(self, metric_namespace: str) -> typing.Self:    
        self._internal.metric_namespace = metric_namespace
    
        return self
    
    def region(self, region: str) -> typing.Self:    
        self._internal.region = region
    
        return self
    


class AzureMetricDimension(cogbuilder.Builder[azuremonitor.AzureMetricDimension]):
    _internal: azuremonitor.AzureMetricDimension

    def __init__(self):
        self._internal = azuremonitor.AzureMetricDimension()

    def build(self) -> azuremonitor.AzureMetricDimension:
        """
        Builds the object.
        """
        return self._internal    
    
    def dimension(self, dimension: str) -> typing.Self:    
        """
        Name of Dimension to be filtered on.
        """
            
        self._internal.dimension = dimension
    
        return self
    
    def operator(self, operator: str) -> typing.Self:    
        """
        String denoting the filter operation. Supports 'eq' - equals,'ne' - not equals, 'sw' - starts with. Note that some dimensions may not support all operators.
        """
            
        self._internal.operator = operator
    
        return self
    
    def filters(self, filters: list[str]) -> typing.Self:    
        """
        Values to match with the filter.
        """
            
        self._internal.filters = filters
    
        return self
    
    def filter(self, filter_val: str) -> typing.Self:    
        """
        @deprecated filter is deprecated in favour of filters to support multiselect.
        """
            
        self._internal.filter_val = filter_val
    
        return self
    


class AzureLogsQuery(cogbuilder.Builder[azuremonitor.AzureLogsQuery]):    
    """
    Azure Monitor Logs sub-query properties
    """
    
    _internal: azuremonitor.AzureLogsQuery

    def __init__(self):
        self._internal = azuremonitor.AzureLogsQuery()

    def build(self) -> azuremonitor.AzureLogsQuery:
        """
        Builds the object.
        """
        return self._internal    
    
    def query(self, query: str) -> typing.Self:    
        """
        KQL query to be executed.
        """
            
        self._internal.query = query
    
        return self
    
    def result_format(self, result_format: azuremonitor.ResultFormat) -> typing.Self:    
        """
        Specifies the format results should be returned as.
        """
            
        self._internal.result_format = result_format
    
        return self
    
    def resources(self, resources: list[str]) -> typing.Self:    
        """
        Array of resource URIs to be queried.
        """
            
        self._internal.resources = resources
    
        return self
    
    def dashboard_time(self, dashboard_time: bool) -> typing.Self:    
        """
        If set to true the dashboard time range will be used as a filter for the query. Otherwise the query time ranges will be used. Defaults to false.
        """
            
        self._internal.dashboard_time = dashboard_time
    
        return self
    
    def time_column(self, time_column: str) -> typing.Self:    
        """
        If dashboardTime is set to true this value dictates which column the time filter will be applied to. Defaults to the first tables timeSpan column, the first datetime column found, or TimeGenerated
        """
            
        self._internal.time_column = time_column
    
        return self
    
    def basic_logs_query(self, basic_logs_query: bool) -> typing.Self:    
        """
        If set to true the query will be run as a basic logs query
        """
            
        self._internal.basic_logs_query = basic_logs_query
    
        return self
    
    def workspace(self, workspace: str) -> typing.Self:    
        """
        Workspace ID. This was removed in Grafana 8, but remains for backwards compat.
        """
            
        self._internal.workspace = workspace
    
        return self
    
    def mode(self, mode: azuremonitor.LogsEditorMode) -> typing.Self:    
        """
        Denotes if logs query editor is in builder mode
        """
            
        self._internal.mode = mode
    
        return self
    
    def builder_query(self, builder_query: cogbuilder.Builder[azuremonitor.BuilderQueryExpression]) -> typing.Self:    
        """
        Builder query to be executed.
        """
            
        builder_query_resource = builder_query.build()
        self._internal.builder_query = builder_query_resource
    
        return self
    
    def resource(self, resource: str) -> typing.Self:    
        """
        @deprecated Use resources instead
        """
            
        self._internal.resource = resource
    
        return self
    
    def intersect_time(self, intersect_time: bool) -> typing.Self:    
        """
        @deprecated Use dashboardTime instead
        """
            
        self._internal.intersect_time = intersect_time
    
        return self
    


class BuilderQueryExpression(cogbuilder.Builder[azuremonitor.BuilderQueryExpression]):
    _internal: azuremonitor.BuilderQueryExpression

    def __init__(self):
        self._internal = azuremonitor.BuilderQueryExpression()

    def build(self) -> azuremonitor.BuilderQueryExpression:
        """
        Builds the object.
        """
        return self._internal    
    
    def from_val(self, from_val: cogbuilder.Builder[azuremonitor.BuilderQueryEditorPropertyExpression]) -> typing.Self:    
        from_val_resource = from_val.build()
        self._internal.from_val = from_val_resource
    
        return self
    
    def columns(self, columns: cogbuilder.Builder[azuremonitor.BuilderQueryEditorColumnsExpression]) -> typing.Self:    
        columns_resource = columns.build()
        self._internal.columns = columns_resource
    
        return self
    
    def where(self, where: cogbuilder.Builder[azuremonitor.BuilderQueryEditorWhereExpressionArray]) -> typing.Self:    
        where_resource = where.build()
        self._internal.where = where_resource
    
        return self
    
    def reduce(self, reduce: cogbuilder.Builder[azuremonitor.BuilderQueryEditorReduceExpressionArray]) -> typing.Self:    
        reduce_resource = reduce.build()
        self._internal.reduce = reduce_resource
    
        return self
    
    def group_by(self, group_by: cogbuilder.Builder[azuremonitor.BuilderQueryEditorGroupByExpressionArray]) -> typing.Self:    
        group_by_resource = group_by.build()
        self._internal.group_by = group_by_resource
    
        return self
    
    def limit(self, limit: int) -> typing.Self:    
        self._internal.limit = limit
    
        return self
    
    def order_by(self, order_by: cogbuilder.Builder[azuremonitor.BuilderQueryEditorOrderByExpressionArray]) -> typing.Self:    
        order_by_resource = order_by.build()
        self._internal.order_by = order_by_resource
    
        return self
    
    def fuzzy_search(self, fuzzy_search: cogbuilder.Builder[azuremonitor.BuilderQueryEditorWhereExpressionArray]) -> typing.Self:    
        fuzzy_search_resource = fuzzy_search.build()
        self._internal.fuzzy_search = fuzzy_search_resource
    
        return self
    
    def time_filter(self, time_filter: cogbuilder.Builder[azuremonitor.BuilderQueryEditorWhereExpressionArray]) -> typing.Self:    
        time_filter_resource = time_filter.build()
        self._internal.time_filter = time_filter_resource
    
        return self
    


class BuilderQueryEditorPropertyExpression(cogbuilder.Builder[azuremonitor.BuilderQueryEditorPropertyExpression]):
    _internal: azuremonitor.BuilderQueryEditorPropertyExpression

    def __init__(self):
        self._internal = azuremonitor.BuilderQueryEditorPropertyExpression()

    def build(self) -> azuremonitor.BuilderQueryEditorPropertyExpression:
        """
        Builds the object.
        """
        return self._internal    
    
    def property(self, property_val: cogbuilder.Builder[azuremonitor.BuilderQueryEditorProperty]) -> typing.Self:    
        property_val_resource = property_val.build()
        self._internal.property_val = property_val_resource
    
        return self
    
    def type(self, type_val: azuremonitor.BuilderQueryEditorExpressionType) -> typing.Self:    
        self._internal.type_val = type_val
    
        return self
    


class BuilderQueryEditorProperty(cogbuilder.Builder[azuremonitor.BuilderQueryEditorProperty]):
    _internal: azuremonitor.BuilderQueryEditorProperty

    def __init__(self):
        self._internal = azuremonitor.BuilderQueryEditorProperty()

    def build(self) -> azuremonitor.BuilderQueryEditorProperty:
        """
        Builds the object.
        """
        return self._internal    
    
    def type(self, type_val: azuremonitor.BuilderQueryEditorPropertyType) -> typing.Self:    
        self._internal.type_val = type_val
    
        return self
    
    def name(self, name: str) -> typing.Self:    
        self._internal.name = name
    
        return self
    


class BuilderQueryEditorColumnsExpression(cogbuilder.Builder[azuremonitor.BuilderQueryEditorColumnsExpression]):
    _internal: azuremonitor.BuilderQueryEditorColumnsExpression

    def __init__(self):
        self._internal = azuremonitor.BuilderQueryEditorColumnsExpression()

    def build(self) -> azuremonitor.BuilderQueryEditorColumnsExpression:
        """
        Builds the object.
        """
        return self._internal    
    
    def columns(self, columns: list[str]) -> typing.Self:    
        self._internal.columns = columns
    
        return self
    
    def type(self, type_val: azuremonitor.BuilderQueryEditorExpressionType) -> typing.Self:    
        self._internal.type_val = type_val
    
        return self
    


class BuilderQueryEditorWhereExpressionArray(cogbuilder.Builder[azuremonitor.BuilderQueryEditorWhereExpressionArray]):
    _internal: azuremonitor.BuilderQueryEditorWhereExpressionArray

    def __init__(self):
        self._internal = azuremonitor.BuilderQueryEditorWhereExpressionArray()

    def build(self) -> azuremonitor.BuilderQueryEditorWhereExpressionArray:
        """
        Builds the object.
        """
        return self._internal    
    
    def expressions(self, expressions: list[cogbuilder.Builder[azuremonitor.BuilderQueryEditorWhereExpression]]) -> typing.Self:    
        expressions_resources = [r1.build() for r1 in expressions]
        self._internal.expressions = expressions_resources
    
        return self
    
    def type(self, type_val: azuremonitor.BuilderQueryEditorExpressionType) -> typing.Self:    
        self._internal.type_val = type_val
    
        return self
    


class BuilderQueryEditorWhereExpression(cogbuilder.Builder[azuremonitor.BuilderQueryEditorWhereExpression]):
    _internal: azuremonitor.BuilderQueryEditorWhereExpression

    def __init__(self):
        self._internal = azuremonitor.BuilderQueryEditorWhereExpression()

    def build(self) -> azuremonitor.BuilderQueryEditorWhereExpression:
        """
        Builds the object.
        """
        return self._internal    
    
    def type(self, type_val: azuremonitor.BuilderQueryEditorExpressionType) -> typing.Self:    
        self._internal.type_val = type_val
    
        return self
    
    def expressions(self, expressions: list[cogbuilder.Builder[azuremonitor.BuilderQueryEditorWhereExpressionItems]]) -> typing.Self:    
        expressions_resources = [r1.build() for r1 in expressions]
        self._internal.expressions = expressions_resources
    
        return self
    


class BuilderQueryEditorWhereExpressionItems(cogbuilder.Builder[azuremonitor.BuilderQueryEditorWhereExpressionItems]):
    _internal: azuremonitor.BuilderQueryEditorWhereExpressionItems

    def __init__(self):
        self._internal = azuremonitor.BuilderQueryEditorWhereExpressionItems()

    def build(self) -> azuremonitor.BuilderQueryEditorWhereExpressionItems:
        """
        Builds the object.
        """
        return self._internal    
    
    def property(self, property_val: cogbuilder.Builder[azuremonitor.BuilderQueryEditorProperty]) -> typing.Self:    
        property_val_resource = property_val.build()
        self._internal.property_val = property_val_resource
    
        return self
    
    def operator(self, operator: cogbuilder.Builder[azuremonitor.BuilderQueryEditorOperator]) -> typing.Self:    
        operator_resource = operator.build()
        self._internal.operator = operator_resource
    
        return self
    
    def type(self, type_val: azuremonitor.BuilderQueryEditorExpressionType) -> typing.Self:    
        self._internal.type_val = type_val
    
        return self
    


class BuilderQueryEditorOperator(cogbuilder.Builder[azuremonitor.BuilderQueryEditorOperator]):
    _internal: azuremonitor.BuilderQueryEditorOperator

    def __init__(self):
        self._internal = azuremonitor.BuilderQueryEditorOperator()

    def build(self) -> azuremonitor.BuilderQueryEditorOperator:
        """
        Builds the object.
        """
        return self._internal    
    
    def name(self, name: str) -> typing.Self:    
        self._internal.name = name
    
        return self
    
    def value(self, value: str) -> typing.Self:    
        self._internal.value = value
    
        return self
    
    def label_value(self, label_value: str) -> typing.Self:    
        self._internal.label_value = label_value
    
        return self
    


class BuilderQueryEditorReduceExpressionArray(cogbuilder.Builder[azuremonitor.BuilderQueryEditorReduceExpressionArray]):
    _internal: azuremonitor.BuilderQueryEditorReduceExpressionArray

    def __init__(self):
        self._internal = azuremonitor.BuilderQueryEditorReduceExpressionArray()

    def build(self) -> azuremonitor.BuilderQueryEditorReduceExpressionArray:
        """
        Builds the object.
        """
        return self._internal    
    
    def expressions(self, expressions: list[cogbuilder.Builder[azuremonitor.BuilderQueryEditorReduceExpression]]) -> typing.Self:    
        expressions_resources = [r1.build() for r1 in expressions]
        self._internal.expressions = expressions_resources
    
        return self
    
    def type(self, type_val: azuremonitor.BuilderQueryEditorExpressionType) -> typing.Self:    
        self._internal.type_val = type_val
    
        return self
    


class BuilderQueryEditorReduceExpression(cogbuilder.Builder[azuremonitor.BuilderQueryEditorReduceExpression]):
    _internal: azuremonitor.BuilderQueryEditorReduceExpression

    def __init__(self):
        self._internal = azuremonitor.BuilderQueryEditorReduceExpression()

    def build(self) -> azuremonitor.BuilderQueryEditorReduceExpression:
        """
        Builds the object.
        """
        return self._internal    
    
    def property(self, property_val: cogbuilder.Builder[azuremonitor.BuilderQueryEditorProperty]) -> typing.Self:    
        property_val_resource = property_val.build()
        self._internal.property_val = property_val_resource
    
        return self
    
    def reduce(self, reduce: cogbuilder.Builder[azuremonitor.BuilderQueryEditorProperty]) -> typing.Self:    
        reduce_resource = reduce.build()
        self._internal.reduce = reduce_resource
    
        return self
    
    def parameters(self, parameters: list[cogbuilder.Builder[azuremonitor.BuilderQueryEditorFunctionParameterExpression]]) -> typing.Self:    
        parameters_resources = [r1.build() for r1 in parameters]
        self._internal.parameters = parameters_resources
    
        return self
    
    def focus(self, focus: bool) -> typing.Self:    
        self._internal.focus = focus
    
        return self
    


class BuilderQueryEditorFunctionParameterExpression(cogbuilder.Builder[azuremonitor.BuilderQueryEditorFunctionParameterExpression]):
    _internal: azuremonitor.BuilderQueryEditorFunctionParameterExpression

    def __init__(self):
        self._internal = azuremonitor.BuilderQueryEditorFunctionParameterExpression()

    def build(self) -> azuremonitor.BuilderQueryEditorFunctionParameterExpression:
        """
        Builds the object.
        """
        return self._internal    
    
    def value(self, value: str) -> typing.Self:    
        self._internal.value = value
    
        return self
    
    def field_type(self, field_type: azuremonitor.BuilderQueryEditorPropertyType) -> typing.Self:    
        self._internal.field_type = field_type
    
        return self
    
    def type(self, type_val: azuremonitor.BuilderQueryEditorExpressionType) -> typing.Self:    
        self._internal.type_val = type_val
    
        return self
    


class BuilderQueryEditorGroupByExpressionArray(cogbuilder.Builder[azuremonitor.BuilderQueryEditorGroupByExpressionArray]):
    _internal: azuremonitor.BuilderQueryEditorGroupByExpressionArray

    def __init__(self):
        self._internal = azuremonitor.BuilderQueryEditorGroupByExpressionArray()

    def build(self) -> azuremonitor.BuilderQueryEditorGroupByExpressionArray:
        """
        Builds the object.
        """
        return self._internal    
    
    def expressions(self, expressions: list[cogbuilder.Builder[azuremonitor.BuilderQueryEditorGroupByExpression]]) -> typing.Self:    
        expressions_resources = [r1.build() for r1 in expressions]
        self._internal.expressions = expressions_resources
    
        return self
    
    def type(self, type_val: azuremonitor.BuilderQueryEditorExpressionType) -> typing.Self:    
        self._internal.type_val = type_val
    
        return self
    


class BuilderQueryEditorGroupByExpression(cogbuilder.Builder[azuremonitor.BuilderQueryEditorGroupByExpression]):
    _internal: azuremonitor.BuilderQueryEditorGroupByExpression

    def __init__(self):
        self._internal = azuremonitor.BuilderQueryEditorGroupByExpression()

    def build(self) -> azuremonitor.BuilderQueryEditorGroupByExpression:
        """
        Builds the object.
        """
        return self._internal    
    
    def property(self, property_val: cogbuilder.Builder[azuremonitor.BuilderQueryEditorProperty]) -> typing.Self:    
        property_val_resource = property_val.build()
        self._internal.property_val = property_val_resource
    
        return self
    
    def interval(self, interval: cogbuilder.Builder[azuremonitor.BuilderQueryEditorProperty]) -> typing.Self:    
        interval_resource = interval.build()
        self._internal.interval = interval_resource
    
        return self
    
    def focus(self, focus: bool) -> typing.Self:    
        self._internal.focus = focus
    
        return self
    
    def type(self, type_val: azuremonitor.BuilderQueryEditorExpressionType) -> typing.Self:    
        self._internal.type_val = type_val
    
        return self
    


class BuilderQueryEditorOrderByExpressionArray(cogbuilder.Builder[azuremonitor.BuilderQueryEditorOrderByExpressionArray]):
    _internal: azuremonitor.BuilderQueryEditorOrderByExpressionArray

    def __init__(self):
        self._internal = azuremonitor.BuilderQueryEditorOrderByExpressionArray()

    def build(self) -> azuremonitor.BuilderQueryEditorOrderByExpressionArray:
        """
        Builds the object.
        """
        return self._internal    
    
    def expressions(self, expressions: list[cogbuilder.Builder[azuremonitor.BuilderQueryEditorOrderByExpression]]) -> typing.Self:    
        expressions_resources = [r1.build() for r1 in expressions]
        self._internal.expressions = expressions_resources
    
        return self
    
    def type(self, type_val: azuremonitor.BuilderQueryEditorExpressionType) -> typing.Self:    
        self._internal.type_val = type_val
    
        return self
    


class BuilderQueryEditorOrderByExpression(cogbuilder.Builder[azuremonitor.BuilderQueryEditorOrderByExpression]):
    _internal: azuremonitor.BuilderQueryEditorOrderByExpression

    def __init__(self):
        self._internal = azuremonitor.BuilderQueryEditorOrderByExpression()

    def build(self) -> azuremonitor.BuilderQueryEditorOrderByExpression:
        """
        Builds the object.
        """
        return self._internal    
    
    def property(self, property_val: cogbuilder.Builder[azuremonitor.BuilderQueryEditorProperty]) -> typing.Self:    
        property_val_resource = property_val.build()
        self._internal.property_val = property_val_resource
    
        return self
    
    def order(self, order: azuremonitor.BuilderQueryEditorOrderByOptions) -> typing.Self:    
        self._internal.order = order
    
        return self
    
    def type(self, type_val: azuremonitor.BuilderQueryEditorExpressionType) -> typing.Self:    
        self._internal.type_val = type_val
    
        return self
    


class AzureResourceGraphQuery(cogbuilder.Builder[azuremonitor.AzureResourceGraphQuery]):
    _internal: azuremonitor.AzureResourceGraphQuery

    def __init__(self):
        self._internal = azuremonitor.AzureResourceGraphQuery()

    def build(self) -> azuremonitor.AzureResourceGraphQuery:
        """
        Builds the object.
        """
        return self._internal    
    
    def query(self, query: str) -> typing.Self:    
        """
        Azure Resource Graph KQL query to be executed.
        """
            
        self._internal.query = query
    
        return self
    
    def result_format(self, result_format: str) -> typing.Self:    
        """
        Specifies the format results should be returned as. Defaults to table.
        """
            
        self._internal.result_format = result_format
    
        return self
    


class AzureTracesQuery(cogbuilder.Builder[azuremonitor.AzureTracesQuery]):    
    """
    Application Insights Traces sub-query properties
    """
    
    _internal: azuremonitor.AzureTracesQuery

    def __init__(self):
        self._internal = azuremonitor.AzureTracesQuery()

    def build(self) -> azuremonitor.AzureTracesQuery:
        """
        Builds the object.
        """
        return self._internal    
    
    def result_format(self, result_format: azuremonitor.ResultFormat) -> typing.Self:    
        """
        Specifies the format results should be returned as.
        """
            
        self._internal.result_format = result_format
    
        return self
    
    def resources(self, resources: list[str]) -> typing.Self:    
        """
        Array of resource URIs to be queried.
        """
            
        self._internal.resources = resources
    
        return self
    
    def operation_id(self, operation_id: str) -> typing.Self:    
        """
        Operation ID. Used only for Traces queries.
        """
            
        self._internal.operation_id = operation_id
    
        return self
    
    def trace_types(self, trace_types: list[str]) -> typing.Self:    
        """
        Types of events to filter by.
        """
            
        self._internal.trace_types = trace_types
    
        return self
    
    def filters(self, filters: list[cogbuilder.Builder[azuremonitor.AzureTracesFilter]]) -> typing.Self:    
        """
        Filters for property values.
        """
            
        filters_resources = [r1.build() for r1 in filters]
        self._internal.filters = filters_resources
    
        return self
    
    def query(self, query: str) -> typing.Self:    
        """
        KQL query to be executed.
        """
            
        self._internal.query = query
    
        return self
    


class AzureTracesFilter(cogbuilder.Builder[azuremonitor.AzureTracesFilter]):
    _internal: azuremonitor.AzureTracesFilter

    def __init__(self):
        self._internal = azuremonitor.AzureTracesFilter()

    def build(self) -> azuremonitor.AzureTracesFilter:
        """
        Builds the object.
        """
        return self._internal    
    
    def property(self, property_val: str) -> typing.Self:    
        """
        Property name, auto-populated based on available traces.
        """
            
        self._internal.property_val = property_val
    
        return self
    
    def operation(self, operation: str) -> typing.Self:    
        """
        Comparison operator to use. Either equals or not equals.
        """
            
        self._internal.operation = operation
    
        return self
    
    def filters(self, filters: list[str]) -> typing.Self:    
        """
        Values to filter by.
        """
            
        self._internal.filters = filters
    
        return self
    


class AppInsightsMetricNameQuery(cogbuilder.Builder[azuremonitor.AppInsightsMetricNameQuery]):
    _internal: azuremonitor.AppInsightsMetricNameQuery

    def __init__(self):
        self._internal = azuremonitor.AppInsightsMetricNameQuery()        
        self._internal.kind = "AppInsightsMetricNameQuery"

    def build(self) -> azuremonitor.AppInsightsMetricNameQuery:
        """
        Builds the object.
        """
        return self._internal    
    
    def raw_query(self, raw_query: str) -> typing.Self:    
        self._internal.raw_query = raw_query
    
        return self
    


class AppInsightsGroupByQuery(cogbuilder.Builder[azuremonitor.AppInsightsGroupByQuery]):
    _internal: azuremonitor.AppInsightsGroupByQuery

    def __init__(self):
        self._internal = azuremonitor.AppInsightsGroupByQuery()        
        self._internal.kind = "AppInsightsGroupByQuery"

    def build(self) -> azuremonitor.AppInsightsGroupByQuery:
        """
        Builds the object.
        """
        return self._internal    
    
    def raw_query(self, raw_query: str) -> typing.Self:    
        self._internal.raw_query = raw_query
    
        return self
    
    def metric_name(self, metric_name: str) -> typing.Self:    
        self._internal.metric_name = metric_name
    
        return self
    


class SubscriptionsQuery(cogbuilder.Builder[azuremonitor.SubscriptionsQuery]):
    _internal: azuremonitor.SubscriptionsQuery

    def __init__(self):
        self._internal = azuremonitor.SubscriptionsQuery()        
        self._internal.kind = "SubscriptionsQuery"

    def build(self) -> azuremonitor.SubscriptionsQuery:
        """
        Builds the object.
        """
        return self._internal    
    
    def raw_query(self, raw_query: str) -> typing.Self:    
        self._internal.raw_query = raw_query
    
        return self
    


class ResourceGroupsQuery(cogbuilder.Builder[azuremonitor.ResourceGroupsQuery]):
    _internal: azuremonitor.ResourceGroupsQuery

    def __init__(self):
        self._internal = azuremonitor.ResourceGroupsQuery()        
        self._internal.kind = "ResourceGroupsQuery"

    def build(self) -> azuremonitor.ResourceGroupsQuery:
        """
        Builds the object.
        """
        return self._internal    
    
    def raw_query(self, raw_query: str) -> typing.Self:    
        self._internal.raw_query = raw_query
    
        return self
    
    def subscription(self, subscription: str) -> typing.Self:    
        self._internal.subscription = subscription
    
        return self
    


class ResourceNamesQuery(cogbuilder.Builder[azuremonitor.ResourceNamesQuery]):
    _internal: azuremonitor.ResourceNamesQuery

    def __init__(self):
        self._internal = azuremonitor.ResourceNamesQuery()        
        self._internal.kind = "ResourceNamesQuery"

    def build(self) -> azuremonitor.ResourceNamesQuery:
        """
        Builds the object.
        """
        return self._internal    
    
    def raw_query(self, raw_query: str) -> typing.Self:    
        self._internal.raw_query = raw_query
    
        return self
    
    def subscription(self, subscription: str) -> typing.Self:    
        self._internal.subscription = subscription
    
        return self
    
    def resource_group(self, resource_group: str) -> typing.Self:    
        self._internal.resource_group = resource_group
    
        return self
    
    def metric_namespace(self, metric_namespace: str) -> typing.Self:    
        self._internal.metric_namespace = metric_namespace
    
        return self
    


class MetricNamespaceQuery(cogbuilder.Builder[azuremonitor.MetricNamespaceQuery]):
    _internal: azuremonitor.MetricNamespaceQuery

    def __init__(self):
        self._internal = azuremonitor.MetricNamespaceQuery()        
        self._internal.kind = "MetricNamespaceQuery"

    def build(self) -> azuremonitor.MetricNamespaceQuery:
        """
        Builds the object.
        """
        return self._internal    
    
    def raw_query(self, raw_query: str) -> typing.Self:    
        self._internal.raw_query = raw_query
    
        return self
    
    def subscription(self, subscription: str) -> typing.Self:    
        self._internal.subscription = subscription
    
        return self
    
    def resource_group(self, resource_group: str) -> typing.Self:    
        self._internal.resource_group = resource_group
    
        return self
    
    def metric_namespace(self, metric_namespace: str) -> typing.Self:    
        self._internal.metric_namespace = metric_namespace
    
        return self
    
    def resource_name(self, resource_name: str) -> typing.Self:    
        self._internal.resource_name = resource_name
    
        return self
    


class MetricDefinitionsQuery(cogbuilder.Builder[azuremonitor.MetricDefinitionsQuery]):    
    """
    @deprecated Use MetricNamespaceQuery instead
    """
    
    _internal: azuremonitor.MetricDefinitionsQuery

    def __init__(self):
        self._internal = azuremonitor.MetricDefinitionsQuery()        
        self._internal.kind = "MetricDefinitionsQuery"

    def build(self) -> azuremonitor.MetricDefinitionsQuery:
        """
        Builds the object.
        """
        return self._internal    
    
    def raw_query(self, raw_query: str) -> typing.Self:    
        self._internal.raw_query = raw_query
    
        return self
    
    def subscription(self, subscription: str) -> typing.Self:    
        self._internal.subscription = subscription
    
        return self
    
    def resource_group(self, resource_group: str) -> typing.Self:    
        self._internal.resource_group = resource_group
    
        return self
    
    def metric_namespace(self, metric_namespace: str) -> typing.Self:    
        self._internal.metric_namespace = metric_namespace
    
        return self
    
    def resource_name(self, resource_name: str) -> typing.Self:    
        self._internal.resource_name = resource_name
    
        return self
    


class MetricNamesQuery(cogbuilder.Builder[azuremonitor.MetricNamesQuery]):
    _internal: azuremonitor.MetricNamesQuery

    def __init__(self):
        self._internal = azuremonitor.MetricNamesQuery()        
        self._internal.kind = "MetricNamesQuery"

    def build(self) -> azuremonitor.MetricNamesQuery:
        """
        Builds the object.
        """
        return self._internal    
    
    def raw_query(self, raw_query: str) -> typing.Self:    
        self._internal.raw_query = raw_query
    
        return self
    
    def subscription(self, subscription: str) -> typing.Self:    
        self._internal.subscription = subscription
    
        return self
    
    def resource_group(self, resource_group: str) -> typing.Self:    
        self._internal.resource_group = resource_group
    
        return self
    
    def resource_name(self, resource_name: str) -> typing.Self:    
        self._internal.resource_name = resource_name
    
        return self
    
    def metric_namespace(self, metric_namespace: str) -> typing.Self:    
        self._internal.metric_namespace = metric_namespace
    
        return self
    


class WorkspacesQuery(cogbuilder.Builder[azuremonitor.WorkspacesQuery]):
    _internal: azuremonitor.WorkspacesQuery

    def __init__(self):
        self._internal = azuremonitor.WorkspacesQuery()        
        self._internal.kind = "WorkspacesQuery"

    def build(self) -> azuremonitor.WorkspacesQuery:
        """
        Builds the object.
        """
        return self._internal    
    
    def raw_query(self, raw_query: str) -> typing.Self:    
        self._internal.raw_query = raw_query
    
        return self
    
    def subscription(self, subscription: str) -> typing.Self:    
        self._internal.subscription = subscription
    
        return self
    


class UnknownQuery(cogbuilder.Builder[azuremonitor.UnknownQuery]):
    _internal: azuremonitor.UnknownQuery

    def __init__(self):
        self._internal = azuremonitor.UnknownQuery()        
        self._internal.kind = "UnknownQuery"

    def build(self) -> azuremonitor.UnknownQuery:
        """
        Builds the object.
        """
        return self._internal    
    
    def raw_query(self, raw_query: str) -> typing.Self:    
        self._internal.raw_query = raw_query
    
        return self
    


class SelectableValue(cogbuilder.Builder[azuremonitor.SelectableValue]):
    _internal: azuremonitor.SelectableValue

    def __init__(self):
        self._internal = azuremonitor.SelectableValue()

    def build(self) -> azuremonitor.SelectableValue:
        """
        Builds the object.
        """
        return self._internal    
    
    def label(self, label: str) -> typing.Self:    
        self._internal.label = label
    
        return self
    
    def value(self, value: str) -> typing.Self:    
        self._internal.value = value
    
        return self
    


class BaseGrafanaTemplateVariableQuery(cogbuilder.Builder[azuremonitor.BaseGrafanaTemplateVariableQuery]):
    _internal: azuremonitor.BaseGrafanaTemplateVariableQuery

    def __init__(self):
        self._internal = azuremonitor.BaseGrafanaTemplateVariableQuery()

    def build(self) -> azuremonitor.BaseGrafanaTemplateVariableQuery:
        """
        Builds the object.
        """
        return self._internal    
    
    def raw_query(self, raw_query: str) -> typing.Self:    
        self._internal.raw_query = raw_query
    
        return self
    
