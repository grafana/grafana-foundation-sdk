# Code generated - EDITING IS FUTILE. DO NOT EDIT.

import typing
from ..cog import builder as cogbuilder
from ..models import azuremonitor


class AzureMonitorQuery(cogbuilder.Builder[azuremonitor.AzureMonitorQuery]):    
    __internal: azuremonitor.AzureMonitorQuery

    def __init__(self):
        self.__internal = azuremonitor.AzureMonitorQuery()

    def build(self) -> azuremonitor.AzureMonitorQuery:
        return self.__internal    
    
    def ref_id(self, ref_id: str) -> typing.Self:    
        """
        A unique identifier for the query within the list of targets.
        In server side expressions, the refId is used as a variable name to identify results.
        By default, the UI will assign A->Z; however setting meaningful names may be useful.
        """
            
        self.__internal.ref_id = ref_id
    
        return self
    
    def hide(self, hide: bool) -> typing.Self:    
        """
        true if query is disabled (ie should not be returned to the dashboard)
        Note this does not always imply that the query should not be executed since
        the results from a hidden query may be used as the input to other queries (SSE etc)
        """
            
        self.__internal.hide = hide
    
        return self
    
    def query_type(self, query_type: str) -> typing.Self:    
        """
        Specify the query flavor
        TODO make this required and give it a default
        """
            
        self.__internal.query_type = query_type
    
        return self
    
    def subscription(self, subscription: str) -> typing.Self:    
        """
        Azure subscription containing the resource(s) to be queried.
        """
            
        self.__internal.subscription = subscription
    
        return self
    
    def subscriptions(self, subscriptions: list[str]) -> typing.Self:    
        """
        Subscriptions to be queried via Azure Resource Graph.
        """
            
        self.__internal.subscriptions = subscriptions
    
        return self
    
    def azure_monitor(self, azure_monitor: cogbuilder.Builder[azuremonitor.AzureMetricQuery]) -> typing.Self:    
        """
        Azure Monitor Metrics sub-query properties.
        """
            
        azure_monitor_resource = azure_monitor.build()
        self.__internal.azure_monitor = azure_monitor_resource
    
        return self
    
    def azure_log_analytics(self, azure_log_analytics: cogbuilder.Builder[azuremonitor.AzureLogsQuery]) -> typing.Self:    
        """
        Azure Monitor Logs sub-query properties.
        """
            
        azure_log_analytics_resource = azure_log_analytics.build()
        self.__internal.azure_log_analytics = azure_log_analytics_resource
    
        return self
    
    def azure_resource_graph(self, azure_resource_graph: cogbuilder.Builder[azuremonitor.AzureResourceGraphQuery]) -> typing.Self:    
        """
        Azure Resource Graph sub-query properties.
        """
            
        azure_resource_graph_resource = azure_resource_graph.build()
        self.__internal.azure_resource_graph = azure_resource_graph_resource
    
        return self
    
    def azure_traces(self, azure_traces: cogbuilder.Builder[azuremonitor.AzureTracesQuery]) -> typing.Self:    
        """
        Application Insights Traces sub-query properties.
        """
            
        azure_traces_resource = azure_traces.build()
        self.__internal.azure_traces = azure_traces_resource
    
        return self
    
    def grafana_template_variable_fn(self, grafana_template_variable_fn: azuremonitor.GrafanaTemplateVariableQuery) -> typing.Self:    
        """
        @deprecated Legacy template variable support.
        """
            
        self.__internal.grafana_template_variable_fn = grafana_template_variable_fn
    
        return self
    
    def resource_group(self, resource_group: str) -> typing.Self:    
        """
        Template variables params. These exist for backwards compatiblity with legacy template variables.
        """
            
        self.__internal.resource_group = resource_group
    
        return self
    
    def namespace(self, namespace: str) -> typing.Self:        
        self.__internal.namespace = namespace
    
        return self
    
    def resource(self, resource: str) -> typing.Self:        
        self.__internal.resource = resource
    
        return self
    
    def datasource(self, datasource: object) -> typing.Self:    
        """
        For mixed data sources the selected datasource is on the query level.
        For non mixed scenarios this is undefined.
        TODO find a better way to do this ^ that's friendly to schema
        TODO this shouldn't be unknown but DataSourceRef | null
        """
            
        self.__internal.datasource = datasource
    
        return self
    
    def region(self, region: str) -> typing.Self:    
        """
        Azure Monitor query type.
        queryType: #AzureQueryType
        """
            
        self.__internal.region = region
    
        return self
    

class AzureMetricQuery(cogbuilder.Builder[azuremonitor.AzureMetricQuery]):    
    __internal: azuremonitor.AzureMetricQuery

    def __init__(self):
        self.__internal = azuremonitor.AzureMetricQuery()

    def build(self) -> azuremonitor.AzureMetricQuery:
        return self.__internal    
    
    def resources(self, resources: list[cogbuilder.Builder[azuremonitor.AzureMonitorResource]]) -> typing.Self:    
        """
        Array of resource URIs to be queried.
        """
            
        resources_resources = [r1.build() for r1 in resources]
        self.__internal.resources = resources_resources
    
        return self
    
    def metric_namespace(self, metric_namespace: str) -> typing.Self:    
        """
        metricNamespace is used as the resource type (or resource namespace).
        It's usually equal to the target metric namespace. e.g. microsoft.storage/storageaccounts
        Kept the name of the variable as metricNamespace to avoid backward incompatibility issues.
        """
            
        self.__internal.metric_namespace = metric_namespace
    
        return self
    
    def custom_namespace(self, custom_namespace: str) -> typing.Self:    
        """
        Used as the value for the metricNamespace property when it's different from the resource namespace.
        """
            
        self.__internal.custom_namespace = custom_namespace
    
        return self
    
    def metric_name(self, metric_name: str) -> typing.Self:    
        """
        The metric to query data for within the specified metricNamespace. e.g. UsedCapacity
        """
            
        self.__internal.metric_name = metric_name
    
        return self
    
    def region(self, region: str) -> typing.Self:    
        """
        The Azure region containing the resource(s).
        """
            
        self.__internal.region = region
    
        return self
    
    def time_grain(self, time_grain: str) -> typing.Self:    
        """
        The granularity of data points to be queried. Defaults to auto.
        """
            
        self.__internal.time_grain = time_grain
    
        return self
    
    def aggregation(self, aggregation: str) -> typing.Self:    
        """
        The aggregation to be used within the query. Defaults to the primaryAggregationType defined by the metric.
        """
            
        self.__internal.aggregation = aggregation
    
        return self
    
    def dimension_filters(self, dimension_filters: list[cogbuilder.Builder[azuremonitor.AzureMetricDimension]]) -> typing.Self:    
        """
        Filters to reduce the set of data returned. Dimensions that can be filtered on are defined by the metric.
        """
            
        dimension_filters_resources = [r1.build() for r1 in dimension_filters]
        self.__internal.dimension_filters = dimension_filters_resources
    
        return self
    
    def top(self, top: str) -> typing.Self:    
        """
        Maximum number of records to return. Defaults to 10.
        """
            
        self.__internal.top = top
    
        return self
    
    def allowed_time_grains_ms(self, allowed_time_grains_ms: list[int]) -> typing.Self:    
        """
        Time grains that are supported by the metric.
        """
            
        self.__internal.allowed_time_grains_ms = allowed_time_grains_ms
    
        return self
    
    def alias(self, alias: str) -> typing.Self:    
        """
        Aliases can be set to modify the legend labels. e.g. {{ resourceGroup }}. See docs for more detail.
        """
            
        self.__internal.alias = alias
    
        return self
    
    def time_grain_unit(self, time_grain_unit: str) -> typing.Self:    
        """
        @deprecated
        """
            
        self.__internal.time_grain_unit = time_grain_unit
    
        return self
    
    def dimension(self, dimension: str) -> typing.Self:    
        """
        @deprecated This property was migrated to dimensionFilters and should only be accessed in the migration
        """
            
        self.__internal.dimension = dimension
    
        return self
    
    def dimension_filter(self, dimension_filter: str) -> typing.Self:    
        """
        @deprecated This property was migrated to dimensionFilters and should only be accessed in the migration
        """
            
        self.__internal.dimension_filter = dimension_filter
    
        return self
    
    def metric_definition(self, metric_definition: str) -> typing.Self:    
        """
        @deprecated Use metricNamespace instead
        """
            
        self.__internal.metric_definition = metric_definition
    
        return self
    
    def resource_uri(self, resource_uri: str) -> typing.Self:    
        """
        @deprecated Use resourceGroup, resourceName and metricNamespace instead
        """
            
        self.__internal.resource_uri = resource_uri
    
        return self
    
    def resource_group(self, resource_group: str) -> typing.Self:    
        """
        @deprecated Use resources instead
        """
            
        self.__internal.resource_group = resource_group
    
        return self
    
    def resource_name(self, resource_name: str) -> typing.Self:    
        """
        @deprecated Use resources instead
        """
            
        self.__internal.resource_name = resource_name
    
        return self
    

class AzureLogsQuery(cogbuilder.Builder[azuremonitor.AzureLogsQuery]):    
    """
    Azure Monitor Logs sub-query properties
    """
    
    __internal: azuremonitor.AzureLogsQuery

    def __init__(self):
        self.__internal = azuremonitor.AzureLogsQuery()

    def build(self) -> azuremonitor.AzureLogsQuery:
        return self.__internal    
    
    def query(self, query: str) -> typing.Self:    
        """
        KQL query to be executed.
        """
            
        self.__internal.query = query
    
        return self
    
    def result_format(self, result_format: azuremonitor.ResultFormat) -> typing.Self:    
        """
        Specifies the format results should be returned as.
        """
            
        self.__internal.result_format = result_format
    
        return self
    
    def resources(self, resources: list[str]) -> typing.Self:    
        """
        Array of resource URIs to be queried.
        """
            
        self.__internal.resources = resources
    
        return self
    
    def dashboard_time(self, dashboard_time: bool) -> typing.Self:    
        """
        If set to true the dashboard time range will be used as a filter for the query. Otherwise the query time ranges will be used. Defaults to false.
        """
            
        self.__internal.dashboard_time = dashboard_time
    
        return self
    
    def time_column(self, time_column: str) -> typing.Self:    
        """
        If dashboardTime is set to true this value dictates which column the time filter will be applied to. Defaults to the first tables timeSpan column, the first datetime column found, or TimeGenerated
        """
            
        self.__internal.time_column = time_column
    
        return self
    
    def workspace(self, workspace: str) -> typing.Self:    
        """
        Workspace ID. This was removed in Grafana 8, but remains for backwards compat.
        """
            
        self.__internal.workspace = workspace
    
        return self
    
    def resource(self, resource: str) -> typing.Self:    
        """
        @deprecated Use resources instead
        """
            
        self.__internal.resource = resource
    
        return self
    
    def intersect_time(self, intersect_time: bool) -> typing.Self:    
        """
        @deprecated Use dashboardTime instead
        """
            
        self.__internal.intersect_time = intersect_time
    
        return self
    

class AzureTracesQuery(cogbuilder.Builder[azuremonitor.AzureTracesQuery]):    
    """
    Application Insights Traces sub-query properties
    """
    
    __internal: azuremonitor.AzureTracesQuery

    def __init__(self):
        self.__internal = azuremonitor.AzureTracesQuery()

    def build(self) -> azuremonitor.AzureTracesQuery:
        return self.__internal    
    
    def result_format(self, result_format: azuremonitor.ResultFormat) -> typing.Self:    
        """
        Specifies the format results should be returned as.
        """
            
        self.__internal.result_format = result_format
    
        return self
    
    def resources(self, resources: list[str]) -> typing.Self:    
        """
        Array of resource URIs to be queried.
        """
            
        self.__internal.resources = resources
    
        return self
    
    def operation_id(self, operation_id: str) -> typing.Self:    
        """
        Operation ID. Used only for Traces queries.
        """
            
        self.__internal.operation_id = operation_id
    
        return self
    
    def trace_types(self, trace_types: list[str]) -> typing.Self:    
        """
        Types of events to filter by.
        """
            
        self.__internal.trace_types = trace_types
    
        return self
    
    def filters(self, filters: list[cogbuilder.Builder[azuremonitor.AzureTracesFilter]]) -> typing.Self:    
        """
        Filters for property values.
        """
            
        filters_resources = [r1.build() for r1 in filters]
        self.__internal.filters = filters_resources
    
        return self
    
    def query(self, query: str) -> typing.Self:    
        """
        KQL query to be executed.
        """
            
        self.__internal.query = query
    
        return self
    

class AzureTracesFilter(cogbuilder.Builder[azuremonitor.AzureTracesFilter]):    
    __internal: azuremonitor.AzureTracesFilter

    def __init__(self):
        self.__internal = azuremonitor.AzureTracesFilter()

    def build(self) -> azuremonitor.AzureTracesFilter:
        return self.__internal    
    
    def property_val(self, property_val: str) -> typing.Self:    
        """
        Property name, auto-populated based on available traces.
        """
            
        self.__internal.property_val = property_val
    
        return self
    
    def operation(self, operation: str) -> typing.Self:    
        """
        Comparison operator to use. Either equals or not equals.
        """
            
        self.__internal.operation = operation
    
        return self
    
    def filters(self, filters: list[str]) -> typing.Self:    
        """
        Values to filter by.
        """
            
        self.__internal.filters = filters
    
        return self
    

class AzureResourceGraphQuery(cogbuilder.Builder[azuremonitor.AzureResourceGraphQuery]):    
    __internal: azuremonitor.AzureResourceGraphQuery

    def __init__(self):
        self.__internal = azuremonitor.AzureResourceGraphQuery()

    def build(self) -> azuremonitor.AzureResourceGraphQuery:
        return self.__internal    
    
    def query(self, query: str) -> typing.Self:    
        """
        Azure Resource Graph KQL query to be executed.
        """
            
        self.__internal.query = query
    
        return self
    
    def result_format(self, result_format: str) -> typing.Self:    
        """
        Specifies the format results should be returned as. Defaults to table.
        """
            
        self.__internal.result_format = result_format
    
        return self
    

class AzureMonitorResource(cogbuilder.Builder[azuremonitor.AzureMonitorResource]):    
    __internal: azuremonitor.AzureMonitorResource

    def __init__(self):
        self.__internal = azuremonitor.AzureMonitorResource()

    def build(self) -> azuremonitor.AzureMonitorResource:
        return self.__internal    
    
    def subscription(self, subscription: str) -> typing.Self:        
        self.__internal.subscription = subscription
    
        return self
    
    def resource_group(self, resource_group: str) -> typing.Self:        
        self.__internal.resource_group = resource_group
    
        return self
    
    def resource_name(self, resource_name: str) -> typing.Self:        
        self.__internal.resource_name = resource_name
    
        return self
    
    def metric_namespace(self, metric_namespace: str) -> typing.Self:        
        self.__internal.metric_namespace = metric_namespace
    
        return self
    
    def region(self, region: str) -> typing.Self:        
        self.__internal.region = region
    
        return self
    

class AzureMetricDimension(cogbuilder.Builder[azuremonitor.AzureMetricDimension]):    
    __internal: azuremonitor.AzureMetricDimension

    def __init__(self):
        self.__internal = azuremonitor.AzureMetricDimension()

    def build(self) -> azuremonitor.AzureMetricDimension:
        return self.__internal    
    
    def dimension(self, dimension: str) -> typing.Self:    
        """
        Name of Dimension to be filtered on.
        """
            
        self.__internal.dimension = dimension
    
        return self
    
    def operator(self, operator: str) -> typing.Self:    
        """
        String denoting the filter operation. Supports 'eq' - equals,'ne' - not equals, 'sw' - starts with. Note that some dimensions may not support all operators.
        """
            
        self.__internal.operator = operator
    
        return self
    
    def filters(self, filters: list[str]) -> typing.Self:    
        """
        Values to match with the filter.
        """
            
        self.__internal.filters = filters
    
        return self
    
    def filter_val(self, filter_val: str) -> typing.Self:    
        """
        @deprecated filter is deprecated in favour of filters to support multiselect.
        """
            
        self.__internal.filter_val = filter_val
    
        return self
    

class BaseGrafanaTemplateVariableQuery(cogbuilder.Builder[azuremonitor.BaseGrafanaTemplateVariableQuery]):    
    __internal: azuremonitor.BaseGrafanaTemplateVariableQuery

    def __init__(self):
        self.__internal = azuremonitor.BaseGrafanaTemplateVariableQuery()

    def build(self) -> azuremonitor.BaseGrafanaTemplateVariableQuery:
        return self.__internal    
    
    def raw_query(self, raw_query: str) -> typing.Self:        
        self.__internal.raw_query = raw_query
    
        return self
    

class UnknownQuery(cogbuilder.Builder[azuremonitor.UnknownQuery]):    
    __internal: azuremonitor.UnknownQuery

    def __init__(self):
        self.__internal = azuremonitor.UnknownQuery()        
        self.__internal.kind = "UnknownQuery"

    def build(self) -> azuremonitor.UnknownQuery:
        return self.__internal    
    
    def raw_query(self, raw_query: str) -> typing.Self:        
        self.__internal.raw_query = raw_query
    
        return self
    

class AppInsightsMetricNameQuery(cogbuilder.Builder[azuremonitor.AppInsightsMetricNameQuery]):    
    __internal: azuremonitor.AppInsightsMetricNameQuery

    def __init__(self):
        self.__internal = azuremonitor.AppInsightsMetricNameQuery()        
        self.__internal.kind = "AppInsightsMetricNameQuery"

    def build(self) -> azuremonitor.AppInsightsMetricNameQuery:
        return self.__internal    
    
    def raw_query(self, raw_query: str) -> typing.Self:        
        self.__internal.raw_query = raw_query
    
        return self
    

class AppInsightsGroupByQuery(cogbuilder.Builder[azuremonitor.AppInsightsGroupByQuery]):    
    __internal: azuremonitor.AppInsightsGroupByQuery

    def __init__(self):
        self.__internal = azuremonitor.AppInsightsGroupByQuery()        
        self.__internal.kind = "AppInsightsGroupByQuery"

    def build(self) -> azuremonitor.AppInsightsGroupByQuery:
        return self.__internal    
    
    def raw_query(self, raw_query: str) -> typing.Self:        
        self.__internal.raw_query = raw_query
    
        return self
    
    def metric_name(self, metric_name: str) -> typing.Self:        
        self.__internal.metric_name = metric_name
    
        return self
    

class SubscriptionsQuery(cogbuilder.Builder[azuremonitor.SubscriptionsQuery]):    
    __internal: azuremonitor.SubscriptionsQuery

    def __init__(self):
        self.__internal = azuremonitor.SubscriptionsQuery()        
        self.__internal.kind = "SubscriptionsQuery"

    def build(self) -> azuremonitor.SubscriptionsQuery:
        return self.__internal    
    
    def raw_query(self, raw_query: str) -> typing.Self:        
        self.__internal.raw_query = raw_query
    
        return self
    

class ResourceGroupsQuery(cogbuilder.Builder[azuremonitor.ResourceGroupsQuery]):    
    __internal: azuremonitor.ResourceGroupsQuery

    def __init__(self):
        self.__internal = azuremonitor.ResourceGroupsQuery()        
        self.__internal.kind = "ResourceGroupsQuery"

    def build(self) -> azuremonitor.ResourceGroupsQuery:
        return self.__internal    
    
    def raw_query(self, raw_query: str) -> typing.Self:        
        self.__internal.raw_query = raw_query
    
        return self
    
    def subscription(self, subscription: str) -> typing.Self:        
        self.__internal.subscription = subscription
    
        return self
    

class ResourceNamesQuery(cogbuilder.Builder[azuremonitor.ResourceNamesQuery]):    
    __internal: azuremonitor.ResourceNamesQuery

    def __init__(self):
        self.__internal = azuremonitor.ResourceNamesQuery()        
        self.__internal.kind = "ResourceNamesQuery"

    def build(self) -> azuremonitor.ResourceNamesQuery:
        return self.__internal    
    
    def raw_query(self, raw_query: str) -> typing.Self:        
        self.__internal.raw_query = raw_query
    
        return self
    
    def subscription(self, subscription: str) -> typing.Self:        
        self.__internal.subscription = subscription
    
        return self
    
    def resource_group(self, resource_group: str) -> typing.Self:        
        self.__internal.resource_group = resource_group
    
        return self
    
    def metric_namespace(self, metric_namespace: str) -> typing.Self:        
        self.__internal.metric_namespace = metric_namespace
    
        return self
    

class MetricNamespaceQuery(cogbuilder.Builder[azuremonitor.MetricNamespaceQuery]):    
    __internal: azuremonitor.MetricNamespaceQuery

    def __init__(self):
        self.__internal = azuremonitor.MetricNamespaceQuery()        
        self.__internal.kind = "MetricNamespaceQuery"

    def build(self) -> azuremonitor.MetricNamespaceQuery:
        return self.__internal    
    
    def raw_query(self, raw_query: str) -> typing.Self:        
        self.__internal.raw_query = raw_query
    
        return self
    
    def subscription(self, subscription: str) -> typing.Self:        
        self.__internal.subscription = subscription
    
        return self
    
    def resource_group(self, resource_group: str) -> typing.Self:        
        self.__internal.resource_group = resource_group
    
        return self
    
    def metric_namespace(self, metric_namespace: str) -> typing.Self:        
        self.__internal.metric_namespace = metric_namespace
    
        return self
    
    def resource_name(self, resource_name: str) -> typing.Self:        
        self.__internal.resource_name = resource_name
    
        return self
    

class MetricDefinitionsQuery(cogbuilder.Builder[azuremonitor.MetricDefinitionsQuery]):    
    """
    @deprecated Use MetricNamespaceQuery instead
    """
    
    __internal: azuremonitor.MetricDefinitionsQuery

    def __init__(self):
        self.__internal = azuremonitor.MetricDefinitionsQuery()        
        self.__internal.kind = "MetricDefinitionsQuery"

    def build(self) -> azuremonitor.MetricDefinitionsQuery:
        return self.__internal    
    
    def raw_query(self, raw_query: str) -> typing.Self:        
        self.__internal.raw_query = raw_query
    
        return self
    
    def subscription(self, subscription: str) -> typing.Self:        
        self.__internal.subscription = subscription
    
        return self
    
    def resource_group(self, resource_group: str) -> typing.Self:        
        self.__internal.resource_group = resource_group
    
        return self
    
    def metric_namespace(self, metric_namespace: str) -> typing.Self:        
        self.__internal.metric_namespace = metric_namespace
    
        return self
    
    def resource_name(self, resource_name: str) -> typing.Self:        
        self.__internal.resource_name = resource_name
    
        return self
    

class MetricNamesQuery(cogbuilder.Builder[azuremonitor.MetricNamesQuery]):    
    __internal: azuremonitor.MetricNamesQuery

    def __init__(self):
        self.__internal = azuremonitor.MetricNamesQuery()        
        self.__internal.kind = "MetricNamesQuery"

    def build(self) -> azuremonitor.MetricNamesQuery:
        return self.__internal    
    
    def raw_query(self, raw_query: str) -> typing.Self:        
        self.__internal.raw_query = raw_query
    
        return self
    
    def subscription(self, subscription: str) -> typing.Self:        
        self.__internal.subscription = subscription
    
        return self
    
    def resource_group(self, resource_group: str) -> typing.Self:        
        self.__internal.resource_group = resource_group
    
        return self
    
    def resource_name(self, resource_name: str) -> typing.Self:        
        self.__internal.resource_name = resource_name
    
        return self
    
    def metric_namespace(self, metric_namespace: str) -> typing.Self:        
        self.__internal.metric_namespace = metric_namespace
    
        return self
    

class WorkspacesQuery(cogbuilder.Builder[azuremonitor.WorkspacesQuery]):    
    __internal: azuremonitor.WorkspacesQuery

    def __init__(self):
        self.__internal = azuremonitor.WorkspacesQuery()        
        self.__internal.kind = "WorkspacesQuery"

    def build(self) -> azuremonitor.WorkspacesQuery:
        return self.__internal    
    
    def raw_query(self, raw_query: str) -> typing.Self:        
        self.__internal.raw_query = raw_query
    
        return self
    
    def subscription(self, subscription: str) -> typing.Self:        
        self.__internal.subscription = subscription
    
        return self
    