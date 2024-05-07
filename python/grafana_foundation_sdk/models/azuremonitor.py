# Code generated - EDITING IS FUTILE. DO NOT EDIT.

from ..cog import variants as cogvariants
import typing
from ..cog import runtime as cogruntime
import enum


class AzureMonitorQuery(cogvariants.Dataquery):
    # A unique identifier for the query within the list of targets.
    # In server side expressions, the refId is used as a variable name to identify results.
    # By default, the UI will assign A->Z; however setting meaningful names may be useful.
    ref_id: str
    # true if query is disabled (ie should not be returned to the dashboard)
    # Note this does not always imply that the query should not be executed since
    # the results from a hidden query may be used as the input to other queries (SSE etc)
    hide: typing.Optional[bool]
    # Specify the query flavor
    # TODO make this required and give it a default
    query_type: typing.Optional[str]
    # Azure subscription containing the resource(s) to be queried.
    subscription: typing.Optional[str]
    # Subscriptions to be queried via Azure Resource Graph.
    subscriptions: typing.Optional[list[str]]
    # Azure Monitor Metrics sub-query properties.
    azure_monitor: typing.Optional['AzureMetricQuery']
    # Azure Monitor Logs sub-query properties.
    azure_log_analytics: typing.Optional['AzureLogsQuery']
    # Azure Resource Graph sub-query properties.
    azure_resource_graph: typing.Optional['AzureResourceGraphQuery']
    # Application Insights Traces sub-query properties.
    azure_traces: typing.Optional['AzureTracesQuery']
    # @deprecated Legacy template variable support.
    grafana_template_variable_fn: typing.Optional['GrafanaTemplateVariableQuery']
    # Template variables params. These exist for backwards compatiblity with legacy template variables.
    resource_group: typing.Optional[str]
    namespace: typing.Optional[str]
    resource: typing.Optional[str]
    # For mixed data sources the selected datasource is on the query level.
    # For non mixed scenarios this is undefined.
    # TODO find a better way to do this ^ that's friendly to schema
    # TODO this shouldn't be unknown but DataSourceRef | null
    datasource: typing.Optional[object]
    # Azure Monitor query type.
    # queryType: #AzureQueryType
    region: typing.Optional[str]

    def __init__(self, ref_id: str = "", hide: typing.Optional[bool] = None, query_type: typing.Optional[str] = None, subscription: typing.Optional[str] = None, subscriptions: typing.Optional[list[str]] = None, azure_monitor: typing.Optional['AzureMetricQuery'] = None, azure_log_analytics: typing.Optional['AzureLogsQuery'] = None, azure_resource_graph: typing.Optional['AzureResourceGraphQuery'] = None, azure_traces: typing.Optional['AzureTracesQuery'] = None, grafana_template_variable_fn: typing.Optional['GrafanaTemplateVariableQuery'] = None, resource_group: typing.Optional[str] = None, namespace: typing.Optional[str] = None, resource: typing.Optional[str] = None, datasource: typing.Optional[object] = None, region: typing.Optional[str] = None):
        self.ref_id = ref_id
        self.hide = hide
        self.query_type = query_type
        self.subscription = subscription
        self.subscriptions = subscriptions
        self.azure_monitor = azure_monitor
        self.azure_log_analytics = azure_log_analytics
        self.azure_resource_graph = azure_resource_graph
        self.azure_traces = azure_traces
        self.grafana_template_variable_fn = grafana_template_variable_fn
        self.resource_group = resource_group
        self.namespace = namespace
        self.resource = resource
        self.datasource = datasource
        self.region = region

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "refId": self.ref_id,
        }
        if self.hide is not None:
            payload["hide"] = self.hide
        if self.query_type is not None:
            payload["queryType"] = self.query_type
        if self.subscription is not None:
            payload["subscription"] = self.subscription
        if self.subscriptions is not None:
            payload["subscriptions"] = self.subscriptions
        if self.azure_monitor is not None:
            payload["azureMonitor"] = self.azure_monitor
        if self.azure_log_analytics is not None:
            payload["azureLogAnalytics"] = self.azure_log_analytics
        if self.azure_resource_graph is not None:
            payload["azureResourceGraph"] = self.azure_resource_graph
        if self.azure_traces is not None:
            payload["azureTraces"] = self.azure_traces
        if self.grafana_template_variable_fn is not None:
            payload["grafanaTemplateVariableFn"] = self.grafana_template_variable_fn
        if self.resource_group is not None:
            payload["resourceGroup"] = self.resource_group
        if self.namespace is not None:
            payload["namespace"] = self.namespace
        if self.resource is not None:
            payload["resource"] = self.resource
        if self.datasource is not None:
            payload["datasource"] = self.datasource
        if self.region is not None:
            payload["region"] = self.region
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "refId" in data:
            args["ref_id"] = data["refId"]
        if "hide" in data:
            args["hide"] = data["hide"]
        if "queryType" in data:
            args["query_type"] = data["queryType"]
        if "subscription" in data:
            args["subscription"] = data["subscription"]
        if "subscriptions" in data:
            args["subscriptions"] = data["subscriptions"]
        if "azureMonitor" in data:
            args["azure_monitor"] = AzureMetricQuery.from_json(data["azureMonitor"])
        if "azureLogAnalytics" in data:
            args["azure_log_analytics"] = AzureLogsQuery.from_json(data["azureLogAnalytics"])
        if "azureResourceGraph" in data:
            args["azure_resource_graph"] = AzureResourceGraphQuery.from_json(data["azureResourceGraph"])
        if "azureTraces" in data:
            args["azure_traces"] = AzureTracesQuery.from_json(data["azureTraces"])
        if "grafanaTemplateVariableFn" in data:
            args["grafana_template_variable_fn"] = data["grafanaTemplateVariableFn"]
        if "resourceGroup" in data:
            args["resource_group"] = data["resourceGroup"]
        if "namespace" in data:
            args["namespace"] = data["namespace"]
        if "resource" in data:
            args["resource"] = data["resource"]
        if "datasource" in data:
            args["datasource"] = data["datasource"]
        if "region" in data:
            args["region"] = data["region"]        

        return cls(**args)


def variant_config() -> cogruntime.DataqueryConfig:
    return cogruntime.DataqueryConfig(
        identifier="grafana-azure-monitor-datasource",
        from_json_hook=AzureMonitorQuery.from_json,
    )


class AzureQueryType(enum.StrEnum):
    """
    Defines the supported queryTypes. GrafanaTemplateVariableFn is deprecated
    """

    AZURE_MONITOR = "Azure Monitor"
    LOG_ANALYTICS = "Azure Log Analytics"
    AZURE_RESOURCE_GRAPH = "Azure Resource Graph"
    AZURE_TRACES = "Azure Traces"
    SUBSCRIPTIONS_QUERY = "Azure Subscriptions"
    RESOURCE_GROUPS_QUERY = "Azure Resource Groups"
    NAMESPACES_QUERY = "Azure Namespaces"
    RESOURCE_NAMES_QUERY = "Azure Resource Names"
    METRIC_NAMES_QUERY = "Azure Metric Names"
    WORKSPACES_QUERY = "Azure Workspaces"
    LOCATIONS_QUERY = "Azure Regions"
    GRAFANA_TEMPLATE_VARIABLE_FN = "Grafana Template Variable Function"


class AzureMetricQuery:
    # Array of resource URIs to be queried.
    resources: typing.Optional[list['AzureMonitorResource']]
    # metricNamespace is used as the resource type (or resource namespace).
    # It's usually equal to the target metric namespace. e.g. microsoft.storage/storageaccounts
    # Kept the name of the variable as metricNamespace to avoid backward incompatibility issues.
    metric_namespace: typing.Optional[str]
    # Used as the value for the metricNamespace property when it's different from the resource namespace.
    custom_namespace: typing.Optional[str]
    # The metric to query data for within the specified metricNamespace. e.g. UsedCapacity
    metric_name: typing.Optional[str]
    # The Azure region containing the resource(s).
    region: typing.Optional[str]
    # The granularity of data points to be queried. Defaults to auto.
    time_grain: typing.Optional[str]
    # The aggregation to be used within the query. Defaults to the primaryAggregationType defined by the metric.
    aggregation: typing.Optional[str]
    # Filters to reduce the set of data returned. Dimensions that can be filtered on are defined by the metric.
    dimension_filters: typing.Optional[list['AzureMetricDimension']]
    # Maximum number of records to return. Defaults to 10.
    top: typing.Optional[str]
    # Time grains that are supported by the metric.
    allowed_time_grains_ms: typing.Optional[list[int]]
    # Aliases can be set to modify the legend labels. e.g. {{ resourceGroup }}. See docs for more detail.
    alias: typing.Optional[str]
    # @deprecated
    time_grain_unit: typing.Optional[str]
    # @deprecated This property was migrated to dimensionFilters and should only be accessed in the migration
    dimension: typing.Optional[str]
    # @deprecated This property was migrated to dimensionFilters and should only be accessed in the migration
    dimension_filter: typing.Optional[str]
    # @deprecated Use metricNamespace instead
    metric_definition: typing.Optional[str]
    # @deprecated Use resourceGroup, resourceName and metricNamespace instead
    resource_uri: typing.Optional[str]
    # @deprecated Use resources instead
    resource_group: typing.Optional[str]
    # @deprecated Use resources instead
    resource_name: typing.Optional[str]

    def __init__(self, resources: typing.Optional[list['AzureMonitorResource']] = None, metric_namespace: typing.Optional[str] = None, custom_namespace: typing.Optional[str] = None, metric_name: typing.Optional[str] = None, region: typing.Optional[str] = None, time_grain: typing.Optional[str] = None, aggregation: typing.Optional[str] = None, dimension_filters: typing.Optional[list['AzureMetricDimension']] = None, top: typing.Optional[str] = None, allowed_time_grains_ms: typing.Optional[list[int]] = None, alias: typing.Optional[str] = None, time_grain_unit: typing.Optional[str] = None, dimension: typing.Optional[str] = None, dimension_filter: typing.Optional[str] = None, metric_definition: typing.Optional[str] = None, resource_uri: typing.Optional[str] = None, resource_group: typing.Optional[str] = None, resource_name: typing.Optional[str] = None):
        self.resources = resources
        self.metric_namespace = metric_namespace
        self.custom_namespace = custom_namespace
        self.metric_name = metric_name
        self.region = region
        self.time_grain = time_grain
        self.aggregation = aggregation
        self.dimension_filters = dimension_filters
        self.top = top
        self.allowed_time_grains_ms = allowed_time_grains_ms
        self.alias = alias
        self.time_grain_unit = time_grain_unit
        self.dimension = dimension
        self.dimension_filter = dimension_filter
        self.metric_definition = metric_definition
        self.resource_uri = resource_uri
        self.resource_group = resource_group
        self.resource_name = resource_name

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.resources is not None:
            payload["resources"] = self.resources
        if self.metric_namespace is not None:
            payload["metricNamespace"] = self.metric_namespace
        if self.custom_namespace is not None:
            payload["customNamespace"] = self.custom_namespace
        if self.metric_name is not None:
            payload["metricName"] = self.metric_name
        if self.region is not None:
            payload["region"] = self.region
        if self.time_grain is not None:
            payload["timeGrain"] = self.time_grain
        if self.aggregation is not None:
            payload["aggregation"] = self.aggregation
        if self.dimension_filters is not None:
            payload["dimensionFilters"] = self.dimension_filters
        if self.top is not None:
            payload["top"] = self.top
        if self.allowed_time_grains_ms is not None:
            payload["allowedTimeGrainsMs"] = self.allowed_time_grains_ms
        if self.alias is not None:
            payload["alias"] = self.alias
        if self.time_grain_unit is not None:
            payload["timeGrainUnit"] = self.time_grain_unit
        if self.dimension is not None:
            payload["dimension"] = self.dimension
        if self.dimension_filter is not None:
            payload["dimensionFilter"] = self.dimension_filter
        if self.metric_definition is not None:
            payload["metricDefinition"] = self.metric_definition
        if self.resource_uri is not None:
            payload["resourceUri"] = self.resource_uri
        if self.resource_group is not None:
            payload["resourceGroup"] = self.resource_group
        if self.resource_name is not None:
            payload["resourceName"] = self.resource_name
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "resources" in data:
            args["resources"] = data["resources"]
        if "metricNamespace" in data:
            args["metric_namespace"] = data["metricNamespace"]
        if "customNamespace" in data:
            args["custom_namespace"] = data["customNamespace"]
        if "metricName" in data:
            args["metric_name"] = data["metricName"]
        if "region" in data:
            args["region"] = data["region"]
        if "timeGrain" in data:
            args["time_grain"] = data["timeGrain"]
        if "aggregation" in data:
            args["aggregation"] = data["aggregation"]
        if "dimensionFilters" in data:
            args["dimension_filters"] = data["dimensionFilters"]
        if "top" in data:
            args["top"] = data["top"]
        if "allowedTimeGrainsMs" in data:
            args["allowed_time_grains_ms"] = data["allowedTimeGrainsMs"]
        if "alias" in data:
            args["alias"] = data["alias"]
        if "timeGrainUnit" in data:
            args["time_grain_unit"] = data["timeGrainUnit"]
        if "dimension" in data:
            args["dimension"] = data["dimension"]
        if "dimensionFilter" in data:
            args["dimension_filter"] = data["dimensionFilter"]
        if "metricDefinition" in data:
            args["metric_definition"] = data["metricDefinition"]
        if "resourceUri" in data:
            args["resource_uri"] = data["resourceUri"]
        if "resourceGroup" in data:
            args["resource_group"] = data["resourceGroup"]
        if "resourceName" in data:
            args["resource_name"] = data["resourceName"]        

        return cls(**args)


class AzureLogsQuery:
    """
    Azure Monitor Logs sub-query properties
    """

    # KQL query to be executed.
    query: typing.Optional[str]
    # Specifies the format results should be returned as.
    result_format: typing.Optional['ResultFormat']
    # Array of resource URIs to be queried.
    resources: typing.Optional[list[str]]
    # If set to true the dashboard time range will be used as a filter for the query. Otherwise the query time ranges will be used. Defaults to false.
    dashboard_time: typing.Optional[bool]
    # If dashboardTime is set to true this value dictates which column the time filter will be applied to. Defaults to the first tables timeSpan column, the first datetime column found, or TimeGenerated
    time_column: typing.Optional[str]
    # Workspace ID. This was removed in Grafana 8, but remains for backwards compat.
    workspace: typing.Optional[str]
    # @deprecated Use resources instead
    resource: typing.Optional[str]
    # @deprecated Use dashboardTime instead
    intersect_time: typing.Optional[bool]

    def __init__(self, query: typing.Optional[str] = None, result_format: typing.Optional['ResultFormat'] = None, resources: typing.Optional[list[str]] = None, dashboard_time: typing.Optional[bool] = None, time_column: typing.Optional[str] = None, workspace: typing.Optional[str] = None, resource: typing.Optional[str] = None, intersect_time: typing.Optional[bool] = None):
        self.query = query
        self.result_format = result_format
        self.resources = resources
        self.dashboard_time = dashboard_time
        self.time_column = time_column
        self.workspace = workspace
        self.resource = resource
        self.intersect_time = intersect_time

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.query is not None:
            payload["query"] = self.query
        if self.result_format is not None:
            payload["resultFormat"] = self.result_format
        if self.resources is not None:
            payload["resources"] = self.resources
        if self.dashboard_time is not None:
            payload["dashboardTime"] = self.dashboard_time
        if self.time_column is not None:
            payload["timeColumn"] = self.time_column
        if self.workspace is not None:
            payload["workspace"] = self.workspace
        if self.resource is not None:
            payload["resource"] = self.resource
        if self.intersect_time is not None:
            payload["intersectTime"] = self.intersect_time
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "query" in data:
            args["query"] = data["query"]
        if "resultFormat" in data:
            args["result_format"] = data["resultFormat"]
        if "resources" in data:
            args["resources"] = data["resources"]
        if "dashboardTime" in data:
            args["dashboard_time"] = data["dashboardTime"]
        if "timeColumn" in data:
            args["time_column"] = data["timeColumn"]
        if "workspace" in data:
            args["workspace"] = data["workspace"]
        if "resource" in data:
            args["resource"] = data["resource"]
        if "intersectTime" in data:
            args["intersect_time"] = data["intersectTime"]        

        return cls(**args)


class AzureTracesQuery:
    """
    Application Insights Traces sub-query properties
    """

    # Specifies the format results should be returned as.
    result_format: typing.Optional['ResultFormat']
    # Array of resource URIs to be queried.
    resources: typing.Optional[list[str]]
    # Operation ID. Used only for Traces queries.
    operation_id: typing.Optional[str]
    # Types of events to filter by.
    trace_types: typing.Optional[list[str]]
    # Filters for property values.
    filters: typing.Optional[list['AzureTracesFilter']]
    # KQL query to be executed.
    query: typing.Optional[str]

    def __init__(self, result_format: typing.Optional['ResultFormat'] = None, resources: typing.Optional[list[str]] = None, operation_id: typing.Optional[str] = None, trace_types: typing.Optional[list[str]] = None, filters: typing.Optional[list['AzureTracesFilter']] = None, query: typing.Optional[str] = None):
        self.result_format = result_format
        self.resources = resources
        self.operation_id = operation_id
        self.trace_types = trace_types
        self.filters = filters
        self.query = query

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.result_format is not None:
            payload["resultFormat"] = self.result_format
        if self.resources is not None:
            payload["resources"] = self.resources
        if self.operation_id is not None:
            payload["operationId"] = self.operation_id
        if self.trace_types is not None:
            payload["traceTypes"] = self.trace_types
        if self.filters is not None:
            payload["filters"] = self.filters
        if self.query is not None:
            payload["query"] = self.query
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "resultFormat" in data:
            args["result_format"] = data["resultFormat"]
        if "resources" in data:
            args["resources"] = data["resources"]
        if "operationId" in data:
            args["operation_id"] = data["operationId"]
        if "traceTypes" in data:
            args["trace_types"] = data["traceTypes"]
        if "filters" in data:
            args["filters"] = data["filters"]
        if "query" in data:
            args["query"] = data["query"]        

        return cls(**args)


class AzureTracesFilter:
    # Property name, auto-populated based on available traces.
    property_val: str
    # Comparison operator to use. Either equals or not equals.
    operation: str
    # Values to filter by.
    filters: list[str]

    def __init__(self, property_val: str = "", operation: str = "", filters: typing.Optional[list[str]] = None):
        self.property_val = property_val
        self.operation = operation
        self.filters = filters if filters is not None else []

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "property": self.property_val,
            "operation": self.operation,
            "filters": self.filters,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "property" in data:
            args["property_val"] = data["property"]
        if "operation" in data:
            args["operation"] = data["operation"]
        if "filters" in data:
            args["filters"] = data["filters"]        

        return cls(**args)


class ResultFormat(enum.StrEnum):
    TABLE = "table"
    TIME_SERIES = "time_series"
    TRACE = "trace"
    LOGS = "logs"


class AzureResourceGraphQuery:
    # Azure Resource Graph KQL query to be executed.
    query: typing.Optional[str]
    # Specifies the format results should be returned as. Defaults to table.
    result_format: typing.Optional[str]

    def __init__(self, query: typing.Optional[str] = None, result_format: typing.Optional[str] = None):
        self.query = query
        self.result_format = result_format

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.query is not None:
            payload["query"] = self.query
        if self.result_format is not None:
            payload["resultFormat"] = self.result_format
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "query" in data:
            args["query"] = data["query"]
        if "resultFormat" in data:
            args["result_format"] = data["resultFormat"]        

        return cls(**args)


class AzureMonitorResource:
    subscription: typing.Optional[str]
    resource_group: typing.Optional[str]
    resource_name: typing.Optional[str]
    metric_namespace: typing.Optional[str]
    region: typing.Optional[str]

    def __init__(self, subscription: typing.Optional[str] = None, resource_group: typing.Optional[str] = None, resource_name: typing.Optional[str] = None, metric_namespace: typing.Optional[str] = None, region: typing.Optional[str] = None):
        self.subscription = subscription
        self.resource_group = resource_group
        self.resource_name = resource_name
        self.metric_namespace = metric_namespace
        self.region = region

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.subscription is not None:
            payload["subscription"] = self.subscription
        if self.resource_group is not None:
            payload["resourceGroup"] = self.resource_group
        if self.resource_name is not None:
            payload["resourceName"] = self.resource_name
        if self.metric_namespace is not None:
            payload["metricNamespace"] = self.metric_namespace
        if self.region is not None:
            payload["region"] = self.region
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "subscription" in data:
            args["subscription"] = data["subscription"]
        if "resourceGroup" in data:
            args["resource_group"] = data["resourceGroup"]
        if "resourceName" in data:
            args["resource_name"] = data["resourceName"]
        if "metricNamespace" in data:
            args["metric_namespace"] = data["metricNamespace"]
        if "region" in data:
            args["region"] = data["region"]        

        return cls(**args)


class AzureMetricDimension:
    # Name of Dimension to be filtered on.
    dimension: typing.Optional[str]
    # String denoting the filter operation. Supports 'eq' - equals,'ne' - not equals, 'sw' - starts with. Note that some dimensions may not support all operators.
    operator: typing.Optional[str]
    # Values to match with the filter.
    filters: typing.Optional[list[str]]
    # @deprecated filter is deprecated in favour of filters to support multiselect.
    filter_val: typing.Optional[str]

    def __init__(self, dimension: typing.Optional[str] = None, operator: typing.Optional[str] = None, filters: typing.Optional[list[str]] = None, filter_val: typing.Optional[str] = None):
        self.dimension = dimension
        self.operator = operator
        self.filters = filters
        self.filter_val = filter_val

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.dimension is not None:
            payload["dimension"] = self.dimension
        if self.operator is not None:
            payload["operator"] = self.operator
        if self.filters is not None:
            payload["filters"] = self.filters
        if self.filter_val is not None:
            payload["filter"] = self.filter_val
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "dimension" in data:
            args["dimension"] = data["dimension"]
        if "operator" in data:
            args["operator"] = data["operator"]
        if "filters" in data:
            args["filters"] = data["filters"]
        if "filter" in data:
            args["filter_val"] = data["filter"]        

        return cls(**args)


class GrafanaTemplateVariableQueryType(enum.StrEnum):
    APP_INSIGHTS_METRIC_NAME_QUERY = "AppInsightsMetricNameQuery"
    APP_INSIGHTS_GROUP_BY_QUERY = "AppInsightsGroupByQuery"
    SUBSCRIPTIONS_QUERY = "SubscriptionsQuery"
    RESOURCE_GROUPS_QUERY = "ResourceGroupsQuery"
    RESOURCE_NAMES_QUERY = "ResourceNamesQuery"
    METRIC_NAMESPACE_QUERY = "MetricNamespaceQuery"
    METRIC_NAMES_QUERY = "MetricNamesQuery"
    WORKSPACES_QUERY = "WorkspacesQuery"
    UNKNOWN_QUERY = "UnknownQuery"


class BaseGrafanaTemplateVariableQuery:
    raw_query: typing.Optional[str]

    def __init__(self, raw_query: typing.Optional[str] = None):
        self.raw_query = raw_query

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.raw_query is not None:
            payload["rawQuery"] = self.raw_query
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "rawQuery" in data:
            args["raw_query"] = data["rawQuery"]        

        return cls(**args)


class UnknownQuery:
    raw_query: typing.Optional[str]
    kind: typing.Literal["UnknownQuery"]

    def __init__(self, raw_query: typing.Optional[str] = None):
        self.raw_query = raw_query
        self.kind = "UnknownQuery"

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "kind": self.kind,
        }
        if self.raw_query is not None:
            payload["rawQuery"] = self.raw_query
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "rawQuery" in data:
            args["raw_query"] = data["rawQuery"]        

        return cls(**args)


class AppInsightsMetricNameQuery:
    raw_query: typing.Optional[str]
    kind: typing.Literal["AppInsightsMetricNameQuery"]

    def __init__(self, raw_query: typing.Optional[str] = None):
        self.raw_query = raw_query
        self.kind = "AppInsightsMetricNameQuery"

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "kind": self.kind,
        }
        if self.raw_query is not None:
            payload["rawQuery"] = self.raw_query
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "rawQuery" in data:
            args["raw_query"] = data["rawQuery"]        

        return cls(**args)


class AppInsightsGroupByQuery:
    raw_query: typing.Optional[str]
    kind: typing.Literal["AppInsightsGroupByQuery"]
    metric_name: str

    def __init__(self, raw_query: typing.Optional[str] = None, metric_name: str = ""):
        self.raw_query = raw_query
        self.kind = "AppInsightsGroupByQuery"
        self.metric_name = metric_name

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "kind": self.kind,
            "metricName": self.metric_name,
        }
        if self.raw_query is not None:
            payload["rawQuery"] = self.raw_query
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "rawQuery" in data:
            args["raw_query"] = data["rawQuery"]
        if "metricName" in data:
            args["metric_name"] = data["metricName"]        

        return cls(**args)


class SubscriptionsQuery:
    raw_query: typing.Optional[str]
    kind: typing.Literal["SubscriptionsQuery"]

    def __init__(self, raw_query: typing.Optional[str] = None):
        self.raw_query = raw_query
        self.kind = "SubscriptionsQuery"

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "kind": self.kind,
        }
        if self.raw_query is not None:
            payload["rawQuery"] = self.raw_query
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "rawQuery" in data:
            args["raw_query"] = data["rawQuery"]        

        return cls(**args)


class ResourceGroupsQuery:
    raw_query: typing.Optional[str]
    kind: typing.Literal["ResourceGroupsQuery"]
    subscription: str

    def __init__(self, raw_query: typing.Optional[str] = None, subscription: str = ""):
        self.raw_query = raw_query
        self.kind = "ResourceGroupsQuery"
        self.subscription = subscription

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "kind": self.kind,
            "subscription": self.subscription,
        }
        if self.raw_query is not None:
            payload["rawQuery"] = self.raw_query
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "rawQuery" in data:
            args["raw_query"] = data["rawQuery"]
        if "subscription" in data:
            args["subscription"] = data["subscription"]        

        return cls(**args)


class ResourceNamesQuery:
    raw_query: typing.Optional[str]
    kind: typing.Literal["ResourceNamesQuery"]
    subscription: str
    resource_group: str
    metric_namespace: str

    def __init__(self, raw_query: typing.Optional[str] = None, subscription: str = "", resource_group: str = "", metric_namespace: str = ""):
        self.raw_query = raw_query
        self.kind = "ResourceNamesQuery"
        self.subscription = subscription
        self.resource_group = resource_group
        self.metric_namespace = metric_namespace

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "kind": self.kind,
            "subscription": self.subscription,
            "resourceGroup": self.resource_group,
            "metricNamespace": self.metric_namespace,
        }
        if self.raw_query is not None:
            payload["rawQuery"] = self.raw_query
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "rawQuery" in data:
            args["raw_query"] = data["rawQuery"]
        if "subscription" in data:
            args["subscription"] = data["subscription"]
        if "resourceGroup" in data:
            args["resource_group"] = data["resourceGroup"]
        if "metricNamespace" in data:
            args["metric_namespace"] = data["metricNamespace"]        

        return cls(**args)


class MetricNamespaceQuery:
    raw_query: typing.Optional[str]
    kind: typing.Literal["MetricNamespaceQuery"]
    subscription: str
    resource_group: str
    metric_namespace: typing.Optional[str]
    resource_name: typing.Optional[str]

    def __init__(self, raw_query: typing.Optional[str] = None, subscription: str = "", resource_group: str = "", metric_namespace: typing.Optional[str] = None, resource_name: typing.Optional[str] = None):
        self.raw_query = raw_query
        self.kind = "MetricNamespaceQuery"
        self.subscription = subscription
        self.resource_group = resource_group
        self.metric_namespace = metric_namespace
        self.resource_name = resource_name

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "kind": self.kind,
            "subscription": self.subscription,
            "resourceGroup": self.resource_group,
        }
        if self.raw_query is not None:
            payload["rawQuery"] = self.raw_query
        if self.metric_namespace is not None:
            payload["metricNamespace"] = self.metric_namespace
        if self.resource_name is not None:
            payload["resourceName"] = self.resource_name
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "rawQuery" in data:
            args["raw_query"] = data["rawQuery"]
        if "subscription" in data:
            args["subscription"] = data["subscription"]
        if "resourceGroup" in data:
            args["resource_group"] = data["resourceGroup"]
        if "metricNamespace" in data:
            args["metric_namespace"] = data["metricNamespace"]
        if "resourceName" in data:
            args["resource_name"] = data["resourceName"]        

        return cls(**args)


class MetricDefinitionsQuery:
    """
    @deprecated Use MetricNamespaceQuery instead
    """

    raw_query: typing.Optional[str]
    kind: typing.Literal["MetricDefinitionsQuery"]
    subscription: str
    resource_group: str
    metric_namespace: typing.Optional[str]
    resource_name: typing.Optional[str]

    def __init__(self, raw_query: typing.Optional[str] = None, subscription: str = "", resource_group: str = "", metric_namespace: typing.Optional[str] = None, resource_name: typing.Optional[str] = None):
        self.raw_query = raw_query
        self.kind = "MetricDefinitionsQuery"
        self.subscription = subscription
        self.resource_group = resource_group
        self.metric_namespace = metric_namespace
        self.resource_name = resource_name

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "kind": self.kind,
            "subscription": self.subscription,
            "resourceGroup": self.resource_group,
        }
        if self.raw_query is not None:
            payload["rawQuery"] = self.raw_query
        if self.metric_namespace is not None:
            payload["metricNamespace"] = self.metric_namespace
        if self.resource_name is not None:
            payload["resourceName"] = self.resource_name
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "rawQuery" in data:
            args["raw_query"] = data["rawQuery"]
        if "subscription" in data:
            args["subscription"] = data["subscription"]
        if "resourceGroup" in data:
            args["resource_group"] = data["resourceGroup"]
        if "metricNamespace" in data:
            args["metric_namespace"] = data["metricNamespace"]
        if "resourceName" in data:
            args["resource_name"] = data["resourceName"]        

        return cls(**args)


class MetricNamesQuery:
    raw_query: typing.Optional[str]
    kind: typing.Literal["MetricNamesQuery"]
    subscription: str
    resource_group: str
    resource_name: str
    metric_namespace: str

    def __init__(self, raw_query: typing.Optional[str] = None, subscription: str = "", resource_group: str = "", resource_name: str = "", metric_namespace: str = ""):
        self.raw_query = raw_query
        self.kind = "MetricNamesQuery"
        self.subscription = subscription
        self.resource_group = resource_group
        self.resource_name = resource_name
        self.metric_namespace = metric_namespace

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "kind": self.kind,
            "subscription": self.subscription,
            "resourceGroup": self.resource_group,
            "resourceName": self.resource_name,
            "metricNamespace": self.metric_namespace,
        }
        if self.raw_query is not None:
            payload["rawQuery"] = self.raw_query
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "rawQuery" in data:
            args["raw_query"] = data["rawQuery"]
        if "subscription" in data:
            args["subscription"] = data["subscription"]
        if "resourceGroup" in data:
            args["resource_group"] = data["resourceGroup"]
        if "resourceName" in data:
            args["resource_name"] = data["resourceName"]
        if "metricNamespace" in data:
            args["metric_namespace"] = data["metricNamespace"]        

        return cls(**args)


class WorkspacesQuery:
    raw_query: typing.Optional[str]
    kind: typing.Literal["WorkspacesQuery"]
    subscription: str

    def __init__(self, raw_query: typing.Optional[str] = None, subscription: str = ""):
        self.raw_query = raw_query
        self.kind = "WorkspacesQuery"
        self.subscription = subscription

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "kind": self.kind,
            "subscription": self.subscription,
        }
        if self.raw_query is not None:
            payload["rawQuery"] = self.raw_query
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "rawQuery" in data:
            args["raw_query"] = data["rawQuery"]
        if "subscription" in data:
            args["subscription"] = data["subscription"]        

        return cls(**args)


GrafanaTemplateVariableQuery: typing.TypeAlias = typing.Union['AppInsightsMetricNameQuery', 'AppInsightsGroupByQuery', 'SubscriptionsQuery', 'ResourceGroupsQuery', 'ResourceNamesQuery', 'MetricNamespaceQuery', 'MetricDefinitionsQuery', 'MetricNamesQuery', 'WorkspacesQuery', 'UnknownQuery']



