# Code generated - EDITING IS FUTILE. DO NOT EDIT.

import typing
from ..cog import variants as cogvariants
import enum
from ..cog import runtime as cogruntime


class MetricStat:
    # AWS region to query for the metric
    region: str
    # A namespace is a container for CloudWatch metrics. Metrics in different namespaces are isolated from each other, so that metrics from different applications are not mistakenly aggregated into the same statistics. For example, Amazon EC2 uses the AWS/EC2 namespace.
    namespace: str
    # Name of the metric
    metric_name: typing.Optional[str]
    # The dimensions of the metric
    dimensions: typing.Optional['Dimensions']
    # Only show metrics that exactly match all defined dimension names.
    match_exact: typing.Optional[bool]
    # The length of time associated with a specific Amazon CloudWatch statistic. Can be specified by a number of seconds, 'auto', or as a duration string e.g. '15m' being 15 minutes
    period: typing.Optional[str]
    # The ID of the AWS account to query for the metric, specifying `all` will query all accounts that the monitoring account is permitted to query.
    account_id: typing.Optional[str]
    # Metric data aggregations over specified periods of time. For detailed definitions of the statistics supported by CloudWatch, see https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Statistics-definitions.html.
    statistic: typing.Optional[str]
    # @deprecated use statistic
    statistics: typing.Optional[list[str]]

    def __init__(self, region: str = "", namespace: str = "", metric_name: typing.Optional[str] = None, dimensions: typing.Optional['Dimensions'] = None, match_exact: typing.Optional[bool] = None, period: typing.Optional[str] = None, account_id: typing.Optional[str] = None, statistic: typing.Optional[str] = None, statistics: typing.Optional[list[str]] = None):
        self.region = region
        self.namespace = namespace
        self.metric_name = metric_name
        self.dimensions = dimensions
        self.match_exact = match_exact
        self.period = period
        self.account_id = account_id
        self.statistic = statistic
        self.statistics = statistics

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "region": self.region,
            "namespace": self.namespace,
        }
        if self.metric_name is not None:
            payload["metricName"] = self.metric_name
        if self.dimensions is not None:
            payload["dimensions"] = self.dimensions
        if self.match_exact is not None:
            payload["matchExact"] = self.match_exact
        if self.period is not None:
            payload["period"] = self.period
        if self.account_id is not None:
            payload["accountId"] = self.account_id
        if self.statistic is not None:
            payload["statistic"] = self.statistic
        if self.statistics is not None:
            payload["statistics"] = self.statistics
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "region" in data:
            args["region"] = data["region"]
        if "namespace" in data:
            args["namespace"] = data["namespace"]
        if "metricName" in data:
            args["metric_name"] = data["metricName"]
        if "dimensions" in data:
            args["dimensions"] = data["dimensions"]
        if "matchExact" in data:
            args["match_exact"] = data["matchExact"]
        if "period" in data:
            args["period"] = data["period"]
        if "accountId" in data:
            args["account_id"] = data["accountId"]
        if "statistic" in data:
            args["statistic"] = data["statistic"]
        if "statistics" in data:
            args["statistics"] = data["statistics"]        

        return cls(**args)


# A name/value pair that is part of the identity of a metric. For example, you can get statistics for a specific EC2 instance by specifying the InstanceId dimension when you search for metrics.
Dimensions: typing.TypeAlias = dict[str, typing.Union[str, list[str]]]


class CloudWatchMetricsQuery(cogvariants.Dataquery):
    """
    Shape of a CloudWatch Metrics query
    """

    # Whether a query is a Metrics, Logs, or Annotations query
    query_mode: 'CloudWatchQueryMode'
    # Whether to use a metric search or metric query. Metric query is referred to as "Metrics Insights" in the AWS console.
    metric_query_type: typing.Optional['MetricQueryType']
    # Whether to use the query builder or code editor to create the query
    metric_editor_mode: typing.Optional['MetricEditorMode']
    # ID can be used to reference other queries in math expressions. The ID can include numbers, letters, and underscore, and must start with a lowercase letter.
    id_val: str
    # Deprecated: use label
    # @deprecated use label
    alias: typing.Optional[str]
    # Change the time series legend names using dynamic labels. See https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/graph-dynamic-labels.html for more details.
    label: typing.Optional[str]
    # Math expression query
    expression: typing.Optional[str]
    # When the metric query type is `metricQueryType` is set to `Query`, this field is used to specify the query string.
    sql_expression: typing.Optional[str]
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
    # AWS region to query for the metric
    region: str
    # A namespace is a container for CloudWatch metrics. Metrics in different namespaces are isolated from each other, so that metrics from different applications are not mistakenly aggregated into the same statistics. For example, Amazon EC2 uses the AWS/EC2 namespace.
    namespace: str
    # Name of the metric
    metric_name: typing.Optional[str]
    # The dimensions of the metric
    dimensions: typing.Optional['Dimensions']
    # Only show metrics that exactly match all defined dimension names.
    match_exact: typing.Optional[bool]
    # The length of time associated with a specific Amazon CloudWatch statistic. Can be specified by a number of seconds, 'auto', or as a duration string e.g. '15m' being 15 minutes
    period: typing.Optional[str]
    # The ID of the AWS account to query for the metric, specifying `all` will query all accounts that the monitoring account is permitted to query.
    account_id: typing.Optional[str]
    # Metric data aggregations over specified periods of time. For detailed definitions of the statistics supported by CloudWatch, see https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Statistics-definitions.html.
    statistic: typing.Optional[str]
    # When the metric query type is `metricQueryType` is set to `Query` and the `metricEditorMode` is set to `Builder`, this field is used to build up an object representation of a SQL query.
    sql: typing.Optional['SQLExpression']
    # For mixed data sources the selected datasource is on the query level.
    # For non mixed scenarios this is undefined.
    # TODO find a better way to do this ^ that's friendly to schema
    # TODO this shouldn't be unknown but DataSourceRef | null
    datasource: typing.Optional[object]
    # @deprecated use statistic
    statistics: typing.Optional[list[str]]

    def __init__(self, query_mode: typing.Optional['CloudWatchQueryMode'] = None, metric_query_type: typing.Optional['MetricQueryType'] = None, metric_editor_mode: typing.Optional['MetricEditorMode'] = None, id_val: str = "", alias: typing.Optional[str] = None, label: typing.Optional[str] = None, expression: typing.Optional[str] = None, sql_expression: typing.Optional[str] = None, ref_id: str = "", hide: typing.Optional[bool] = None, query_type: typing.Optional[str] = None, region: str = "", namespace: str = "", metric_name: typing.Optional[str] = None, dimensions: typing.Optional['Dimensions'] = None, match_exact: typing.Optional[bool] = None, period: typing.Optional[str] = None, account_id: typing.Optional[str] = None, statistic: typing.Optional[str] = None, sql: typing.Optional['SQLExpression'] = None, datasource: typing.Optional[object] = None, statistics: typing.Optional[list[str]] = None):
        self.query_mode = query_mode if query_mode is not None else CloudWatchQueryMode.METRICS
        self.metric_query_type = metric_query_type
        self.metric_editor_mode = metric_editor_mode
        self.id_val = id_val
        self.alias = alias
        self.label = label
        self.expression = expression
        self.sql_expression = sql_expression
        self.ref_id = ref_id
        self.hide = hide
        self.query_type = query_type
        self.region = region
        self.namespace = namespace
        self.metric_name = metric_name
        self.dimensions = dimensions
        self.match_exact = match_exact
        self.period = period
        self.account_id = account_id
        self.statistic = statistic
        self.sql = sql
        self.datasource = datasource
        self.statistics = statistics

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "queryMode": self.query_mode,
            "id": self.id_val,
            "refId": self.ref_id,
            "region": self.region,
            "namespace": self.namespace,
        }
        if self.metric_query_type is not None:
            payload["metricQueryType"] = self.metric_query_type
        if self.metric_editor_mode is not None:
            payload["metricEditorMode"] = self.metric_editor_mode
        if self.alias is not None:
            payload["alias"] = self.alias
        if self.label is not None:
            payload["label"] = self.label
        if self.expression is not None:
            payload["expression"] = self.expression
        if self.sql_expression is not None:
            payload["sqlExpression"] = self.sql_expression
        if self.hide is not None:
            payload["hide"] = self.hide
        if self.query_type is not None:
            payload["queryType"] = self.query_type
        if self.metric_name is not None:
            payload["metricName"] = self.metric_name
        if self.dimensions is not None:
            payload["dimensions"] = self.dimensions
        if self.match_exact is not None:
            payload["matchExact"] = self.match_exact
        if self.period is not None:
            payload["period"] = self.period
        if self.account_id is not None:
            payload["accountId"] = self.account_id
        if self.statistic is not None:
            payload["statistic"] = self.statistic
        if self.sql is not None:
            payload["sql"] = self.sql
        if self.datasource is not None:
            payload["datasource"] = self.datasource
        if self.statistics is not None:
            payload["statistics"] = self.statistics
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "queryMode" in data:
            args["query_mode"] = data["queryMode"]
        if "metricQueryType" in data:
            args["metric_query_type"] = data["metricQueryType"]
        if "metricEditorMode" in data:
            args["metric_editor_mode"] = data["metricEditorMode"]
        if "id" in data:
            args["id_val"] = data["id"]
        if "alias" in data:
            args["alias"] = data["alias"]
        if "label" in data:
            args["label"] = data["label"]
        if "expression" in data:
            args["expression"] = data["expression"]
        if "sqlExpression" in data:
            args["sql_expression"] = data["sqlExpression"]
        if "refId" in data:
            args["ref_id"] = data["refId"]
        if "hide" in data:
            args["hide"] = data["hide"]
        if "queryType" in data:
            args["query_type"] = data["queryType"]
        if "region" in data:
            args["region"] = data["region"]
        if "namespace" in data:
            args["namespace"] = data["namespace"]
        if "metricName" in data:
            args["metric_name"] = data["metricName"]
        if "dimensions" in data:
            args["dimensions"] = data["dimensions"]
        if "matchExact" in data:
            args["match_exact"] = data["matchExact"]
        if "period" in data:
            args["period"] = data["period"]
        if "accountId" in data:
            args["account_id"] = data["accountId"]
        if "statistic" in data:
            args["statistic"] = data["statistic"]
        if "sql" in data:
            args["sql"] = SQLExpression.from_json(data["sql"])
        if "datasource" in data:
            args["datasource"] = data["datasource"]
        if "statistics" in data:
            args["statistics"] = data["statistics"]        

        return cls(**args)


class CloudWatchQueryMode(enum.StrEnum):
    METRICS = "Metrics"
    LOGS = "Logs"
    ANNOTATIONS = "Annotations"


class MetricQueryType(enum.IntEnum):
    SEARCH = 0
    QUERY = 1


class MetricEditorMode(enum.IntEnum):
    BUILDER = 0
    CODE = 1


class SQLExpression:
    # SELECT part of the SQL expression
    select: typing.Optional['QueryEditorFunctionExpression']
    # FROM part of the SQL expression
    from_val: typing.Optional[typing.Union['QueryEditorPropertyExpression', 'QueryEditorFunctionExpression']]
    # WHERE part of the SQL expression
    where: typing.Optional['QueryEditorArrayExpression']
    # GROUP BY part of the SQL expression
    group_by: typing.Optional['QueryEditorArrayExpression']
    # ORDER BY part of the SQL expression
    order_by: typing.Optional['QueryEditorFunctionExpression']
    # The sort order of the SQL expression, `ASC` or `DESC`
    order_by_direction: typing.Optional[str]
    # LIMIT part of the SQL expression
    limit: typing.Optional[int]

    def __init__(self, select: typing.Optional['QueryEditorFunctionExpression'] = None, from_val: typing.Optional[typing.Union['QueryEditorPropertyExpression', 'QueryEditorFunctionExpression']] = None, where: typing.Optional['QueryEditorArrayExpression'] = None, group_by: typing.Optional['QueryEditorArrayExpression'] = None, order_by: typing.Optional['QueryEditorFunctionExpression'] = None, order_by_direction: typing.Optional[str] = None, limit: typing.Optional[int] = None):
        self.select = select
        self.from_val = from_val
        self.where = where
        self.group_by = group_by
        self.order_by = order_by
        self.order_by_direction = order_by_direction
        self.limit = limit

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.select is not None:
            payload["select"] = self.select
        if self.from_val is not None:
            payload["from"] = self.from_val
        if self.where is not None:
            payload["where"] = self.where
        if self.group_by is not None:
            payload["groupBy"] = self.group_by
        if self.order_by is not None:
            payload["orderBy"] = self.order_by
        if self.order_by_direction is not None:
            payload["orderByDirection"] = self.order_by_direction
        if self.limit is not None:
            payload["limit"] = self.limit
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "select" in data:
            args["select"] = QueryEditorFunctionExpression.from_json(data["select"])
        if "from" in data:
            decoding_map: dict[str, typing.Union[typing.Type[QueryEditorPropertyExpression], typing.Type[QueryEditorFunctionExpression]]] = {"property": QueryEditorPropertyExpression, "function": QueryEditorFunctionExpression}
            args["from_val"] = decoding_map[data["from"]["type"]].from_json(data["from"])
        if "where" in data:
            args["where"] = QueryEditorArrayExpression.from_json(data["where"])
        if "groupBy" in data:
            args["group_by"] = QueryEditorArrayExpression.from_json(data["groupBy"])
        if "orderBy" in data:
            args["order_by"] = QueryEditorFunctionExpression.from_json(data["orderBy"])
        if "orderByDirection" in data:
            args["order_by_direction"] = data["orderByDirection"]
        if "limit" in data:
            args["limit"] = data["limit"]        

        return cls(**args)


class QueryEditorFunctionExpression:
    type_val: typing.Literal["function"]
    name: typing.Optional[str]
    parameters: typing.Optional[list['QueryEditorFunctionParameterExpression']]

    def __init__(self, name: typing.Optional[str] = None, parameters: typing.Optional[list['QueryEditorFunctionParameterExpression']] = None):
        self.type_val = "function"
        self.name = name
        self.parameters = parameters

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "type": self.type_val,
        }
        if self.name is not None:
            payload["name"] = self.name
        if self.parameters is not None:
            payload["parameters"] = self.parameters
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "name" in data:
            args["name"] = data["name"]
        if "parameters" in data:
            args["parameters"] = data["parameters"]        

        return cls(**args)


class QueryEditorExpressionType(enum.StrEnum):
    PROPERTY = "property"
    OPERATOR = "operator"
    OR = "or"
    AND = "and"
    GROUP_BY = "groupBy"
    FUNCTION = "function"
    FUNCTION_PARAMETER = "functionParameter"


class QueryEditorFunctionParameterExpression:
    type_val: typing.Literal["functionParameter"]
    name: typing.Optional[str]

    def __init__(self, name: typing.Optional[str] = None):
        self.type_val = "functionParameter"
        self.name = name

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "type": self.type_val,
        }
        if self.name is not None:
            payload["name"] = self.name
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "name" in data:
            args["name"] = data["name"]        

        return cls(**args)


class QueryEditorPropertyExpression:
    type_val: typing.Literal["property"]
    property_val: 'QueryEditorProperty'

    def __init__(self, property_val: typing.Optional['QueryEditorProperty'] = None):
        self.type_val = "property"
        self.property_val = property_val if property_val is not None else QueryEditorProperty()

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "type": self.type_val,
            "property": self.property_val,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "property" in data:
            args["property_val"] = QueryEditorProperty.from_json(data["property"])        

        return cls(**args)


class QueryEditorGroupByExpression:
    type_val: typing.Literal["groupBy"]
    property_val: 'QueryEditorProperty'

    def __init__(self, property_val: typing.Optional['QueryEditorProperty'] = None):
        self.type_val = "groupBy"
        self.property_val = property_val if property_val is not None else QueryEditorProperty()

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "type": self.type_val,
            "property": self.property_val,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "property" in data:
            args["property_val"] = QueryEditorProperty.from_json(data["property"])        

        return cls(**args)


class QueryEditorOperatorExpression:
    type_val: typing.Literal["operator"]
    property_val: 'QueryEditorProperty'
    # TS type is operator: QueryEditorOperator<QueryEditorOperatorValueType>, extended in veneer
    operator: 'QueryEditorOperator'

    def __init__(self, property_val: typing.Optional['QueryEditorProperty'] = None, operator: typing.Optional['QueryEditorOperator'] = None):
        self.type_val = "operator"
        self.property_val = property_val if property_val is not None else QueryEditorProperty()
        self.operator = operator if operator is not None else QueryEditorOperator()

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "type": self.type_val,
            "property": self.property_val,
            "operator": self.operator,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "property" in data:
            args["property_val"] = QueryEditorProperty.from_json(data["property"])
        if "operator" in data:
            args["operator"] = QueryEditorOperator.from_json(data["operator"])        

        return cls(**args)


class QueryEditorOperator:
    """
    TS type is QueryEditorOperator<T extends QueryEditorOperatorValueType>, extended in veneer
    """

    name: typing.Optional[str]
    value: typing.Optional[typing.Union[str, bool, int, list['QueryEditorOperatorType']]]

    def __init__(self, name: typing.Optional[str] = None, value: typing.Optional[typing.Union[str, bool, int, list['QueryEditorOperatorType']]] = None):
        self.name = name
        self.value = value

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.name is not None:
            payload["name"] = self.name
        if self.value is not None:
            payload["value"] = self.value
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "name" in data:
            args["name"] = data["name"]
        if "value" in data:
            args["value"] = data["value"]        

        return cls(**args)


QueryEditorOperatorValueType: typing.TypeAlias = typing.Union[str, bool, int, list['QueryEditorOperatorType']]


QueryEditorOperatorType: typing.TypeAlias = typing.Union[str, bool, int]


class QueryEditorProperty:
    type_val: 'QueryEditorPropertyType'
    name: typing.Optional[str]

    def __init__(self, type_val: typing.Optional['QueryEditorPropertyType'] = None, name: typing.Optional[str] = None):
        self.type_val = type_val if type_val is not None else QueryEditorPropertyType.STRING
        self.name = name

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "type": self.type_val,
        }
        if self.name is not None:
            payload["name"] = self.name
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "type" in data:
            args["type_val"] = data["type"]
        if "name" in data:
            args["name"] = data["name"]        

        return cls(**args)


class QueryEditorPropertyType(enum.StrEnum):
    STRING = "string"


class QueryEditorArrayExpression:
    type_val: typing.Literal["and", "or"]
    expressions: list['QueryEditorExpression']

    def __init__(self, type_val: typing.Optional[typing.Literal["and", "or"]] = None, expressions: typing.Optional[list['QueryEditorExpression']] = None):
        self.type_val = type_val if type_val is not None else "and"
        self.expressions = expressions if expressions is not None else []

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "type": self.type_val,
            "expressions": self.expressions,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "type" in data:
            args["type_val"] = data["type"]
        if "expressions" in data:
            args["expressions"] = data["expressions"]        

        return cls(**args)


QueryEditorExpression: typing.TypeAlias = typing.Union['QueryEditorArrayExpression', 'QueryEditorPropertyExpression', 'QueryEditorGroupByExpression', 'QueryEditorFunctionExpression', 'QueryEditorFunctionParameterExpression', 'QueryEditorOperatorExpression']


class CloudWatchLogsQuery(cogvariants.Dataquery):
    """
    Shape of a CloudWatch Logs query
    """

    # Whether a query is a Metrics, Logs, or Annotations query
    query_mode: 'CloudWatchQueryMode'
    id_val: str
    # AWS region to query for the logs
    region: str
    # The CloudWatch Logs Insights query to execute
    expression: typing.Optional[str]
    # Fields to group the results by, this field is automatically populated whenever the query is updated
    stats_groups: typing.Optional[list[str]]
    # Log groups to query
    log_groups: typing.Optional[list['LogGroup']]
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
    # @deprecated use logGroups
    log_group_names: typing.Optional[list[str]]
    # For mixed data sources the selected datasource is on the query level.
    # For non mixed scenarios this is undefined.
    # TODO find a better way to do this ^ that's friendly to schema
    # TODO this shouldn't be unknown but DataSourceRef | null
    datasource: typing.Optional[object]

    def __init__(self, query_mode: typing.Optional['CloudWatchQueryMode'] = None, id_val: str = "", region: str = "", expression: typing.Optional[str] = None, stats_groups: typing.Optional[list[str]] = None, log_groups: typing.Optional[list['LogGroup']] = None, ref_id: str = "", hide: typing.Optional[bool] = None, query_type: typing.Optional[str] = None, log_group_names: typing.Optional[list[str]] = None, datasource: typing.Optional[object] = None):
        self.query_mode = query_mode if query_mode is not None else CloudWatchQueryMode.LOGS
        self.id_val = id_val
        self.region = region
        self.expression = expression
        self.stats_groups = stats_groups
        self.log_groups = log_groups
        self.ref_id = ref_id
        self.hide = hide
        self.query_type = query_type
        self.log_group_names = log_group_names
        self.datasource = datasource

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "queryMode": self.query_mode,
            "id": self.id_val,
            "region": self.region,
            "refId": self.ref_id,
        }
        if self.expression is not None:
            payload["expression"] = self.expression
        if self.stats_groups is not None:
            payload["statsGroups"] = self.stats_groups
        if self.log_groups is not None:
            payload["logGroups"] = self.log_groups
        if self.hide is not None:
            payload["hide"] = self.hide
        if self.query_type is not None:
            payload["queryType"] = self.query_type
        if self.log_group_names is not None:
            payload["logGroupNames"] = self.log_group_names
        if self.datasource is not None:
            payload["datasource"] = self.datasource
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "queryMode" in data:
            args["query_mode"] = data["queryMode"]
        if "id" in data:
            args["id_val"] = data["id"]
        if "region" in data:
            args["region"] = data["region"]
        if "expression" in data:
            args["expression"] = data["expression"]
        if "statsGroups" in data:
            args["stats_groups"] = data["statsGroups"]
        if "logGroups" in data:
            args["log_groups"] = data["logGroups"]
        if "refId" in data:
            args["ref_id"] = data["refId"]
        if "hide" in data:
            args["hide"] = data["hide"]
        if "queryType" in data:
            args["query_type"] = data["queryType"]
        if "logGroupNames" in data:
            args["log_group_names"] = data["logGroupNames"]
        if "datasource" in data:
            args["datasource"] = data["datasource"]        

        return cls(**args)


class LogGroup:
    # ARN of the log group
    arn: str
    # Name of the log group
    name: str
    # AccountId of the log group
    account_id: typing.Optional[str]
    # Label of the log group
    account_label: typing.Optional[str]

    def __init__(self, arn: str = "", name: str = "", account_id: typing.Optional[str] = None, account_label: typing.Optional[str] = None):
        self.arn = arn
        self.name = name
        self.account_id = account_id
        self.account_label = account_label

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "arn": self.arn,
            "name": self.name,
        }
        if self.account_id is not None:
            payload["accountId"] = self.account_id
        if self.account_label is not None:
            payload["accountLabel"] = self.account_label
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "arn" in data:
            args["arn"] = data["arn"]
        if "name" in data:
            args["name"] = data["name"]
        if "accountId" in data:
            args["account_id"] = data["accountId"]
        if "accountLabel" in data:
            args["account_label"] = data["accountLabel"]        

        return cls(**args)


class CloudWatchAnnotationQuery(cogvariants.Dataquery):
    """
    Shape of a CloudWatch Annotation query
    TS type is CloudWatchDefaultQuery = Omit<CloudWatchLogsQuery, 'queryMode'> & CloudWatchMetricsQuery, declared in veneer
    #CloudWatchDefaultQuery: #CloudWatchLogsQuery & #CloudWatchMetricsQuery @cuetsy(kind="type")
    """

    # Whether a query is a Metrics, Logs, or Annotations query
    query_mode: 'CloudWatchQueryMode'
    # Enable matching on the prefix of the action name or alarm name, specify the prefixes with actionPrefix and/or alarmNamePrefix
    prefix_matching: typing.Optional[bool]
    # Use this parameter to filter the results of the operation to only those alarms
    # that use a certain alarm action. For example, you could specify the ARN of
    # an SNS topic to find all alarms that send notifications to that topic.
    # e.g. `arn:aws:sns:us-east-1:123456789012:my-app-` would match `arn:aws:sns:us-east-1:123456789012:my-app-action`
    # but not match `arn:aws:sns:us-east-1:123456789012:your-app-action`
    action_prefix: typing.Optional[str]
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
    # AWS region to query for the metric
    region: str
    # A namespace is a container for CloudWatch metrics. Metrics in different namespaces are isolated from each other, so that metrics from different applications are not mistakenly aggregated into the same statistics. For example, Amazon EC2 uses the AWS/EC2 namespace.
    namespace: str
    # Name of the metric
    metric_name: typing.Optional[str]
    # The dimensions of the metric
    dimensions: typing.Optional['Dimensions']
    # Only show metrics that exactly match all defined dimension names.
    match_exact: typing.Optional[bool]
    # The length of time associated with a specific Amazon CloudWatch statistic. Can be specified by a number of seconds, 'auto', or as a duration string e.g. '15m' being 15 minutes
    period: typing.Optional[str]
    # The ID of the AWS account to query for the metric, specifying `all` will query all accounts that the monitoring account is permitted to query.
    account_id: typing.Optional[str]
    # Metric data aggregations over specified periods of time. For detailed definitions of the statistics supported by CloudWatch, see https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Statistics-definitions.html.
    statistic: typing.Optional[str]
    # An alarm name prefix. If you specify this parameter, you receive information
    # about all alarms that have names that start with this prefix.
    # e.g. `my-team-service-` would match `my-team-service-high-cpu` but not match `your-team-service-high-cpu`
    alarm_name_prefix: typing.Optional[str]
    # For mixed data sources the selected datasource is on the query level.
    # For non mixed scenarios this is undefined.
    # TODO find a better way to do this ^ that's friendly to schema
    # TODO this shouldn't be unknown but DataSourceRef | null
    datasource: typing.Optional[object]
    # @deprecated use statistic
    statistics: typing.Optional[list[str]]

    def __init__(self, query_mode: typing.Optional['CloudWatchQueryMode'] = None, prefix_matching: typing.Optional[bool] = None, action_prefix: typing.Optional[str] = None, ref_id: str = "", hide: typing.Optional[bool] = None, query_type: typing.Optional[str] = None, region: str = "", namespace: str = "", metric_name: typing.Optional[str] = None, dimensions: typing.Optional['Dimensions'] = None, match_exact: typing.Optional[bool] = None, period: typing.Optional[str] = None, account_id: typing.Optional[str] = None, statistic: typing.Optional[str] = None, alarm_name_prefix: typing.Optional[str] = None, datasource: typing.Optional[object] = None, statistics: typing.Optional[list[str]] = None):
        self.query_mode = query_mode if query_mode is not None else CloudWatchQueryMode.ANNOTATIONS
        self.prefix_matching = prefix_matching
        self.action_prefix = action_prefix
        self.ref_id = ref_id
        self.hide = hide
        self.query_type = query_type
        self.region = region
        self.namespace = namespace
        self.metric_name = metric_name
        self.dimensions = dimensions
        self.match_exact = match_exact
        self.period = period
        self.account_id = account_id
        self.statistic = statistic
        self.alarm_name_prefix = alarm_name_prefix
        self.datasource = datasource
        self.statistics = statistics

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "queryMode": self.query_mode,
            "refId": self.ref_id,
            "region": self.region,
            "namespace": self.namespace,
        }
        if self.prefix_matching is not None:
            payload["prefixMatching"] = self.prefix_matching
        if self.action_prefix is not None:
            payload["actionPrefix"] = self.action_prefix
        if self.hide is not None:
            payload["hide"] = self.hide
        if self.query_type is not None:
            payload["queryType"] = self.query_type
        if self.metric_name is not None:
            payload["metricName"] = self.metric_name
        if self.dimensions is not None:
            payload["dimensions"] = self.dimensions
        if self.match_exact is not None:
            payload["matchExact"] = self.match_exact
        if self.period is not None:
            payload["period"] = self.period
        if self.account_id is not None:
            payload["accountId"] = self.account_id
        if self.statistic is not None:
            payload["statistic"] = self.statistic
        if self.alarm_name_prefix is not None:
            payload["alarmNamePrefix"] = self.alarm_name_prefix
        if self.datasource is not None:
            payload["datasource"] = self.datasource
        if self.statistics is not None:
            payload["statistics"] = self.statistics
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "queryMode" in data:
            args["query_mode"] = data["queryMode"]
        if "prefixMatching" in data:
            args["prefix_matching"] = data["prefixMatching"]
        if "actionPrefix" in data:
            args["action_prefix"] = data["actionPrefix"]
        if "refId" in data:
            args["ref_id"] = data["refId"]
        if "hide" in data:
            args["hide"] = data["hide"]
        if "queryType" in data:
            args["query_type"] = data["queryType"]
        if "region" in data:
            args["region"] = data["region"]
        if "namespace" in data:
            args["namespace"] = data["namespace"]
        if "metricName" in data:
            args["metric_name"] = data["metricName"]
        if "dimensions" in data:
            args["dimensions"] = data["dimensions"]
        if "matchExact" in data:
            args["match_exact"] = data["matchExact"]
        if "period" in data:
            args["period"] = data["period"]
        if "accountId" in data:
            args["account_id"] = data["accountId"]
        if "statistic" in data:
            args["statistic"] = data["statistic"]
        if "alarmNamePrefix" in data:
            args["alarm_name_prefix"] = data["alarmNamePrefix"]
        if "datasource" in data:
            args["datasource"] = data["datasource"]
        if "statistics" in data:
            args["statistics"] = data["statistics"]        

        return cls(**args)


CloudWatchQuery: typing.TypeAlias = typing.Union['CloudWatchMetricsQuery', 'CloudWatchLogsQuery', 'CloudWatchAnnotationQuery']


def variant_config() -> cogruntime.DataqueryConfig:
    decoding_map: dict[str, typing.Union[typing.Type[CloudWatchMetricsQuery], typing.Type[CloudWatchLogsQuery], typing.Type[CloudWatchAnnotationQuery]]] = {"Metrics": CloudWatchMetricsQuery, "Logs": CloudWatchLogsQuery, "Annotations": CloudWatchAnnotationQuery}
    return cogruntime.DataqueryConfig(
        identifier="cloudwatch",
        from_json_hook=lambda data: decoding_map[data["queryMode"]].from_json(data),
    )



