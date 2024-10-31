# Code generated - EDITING IS FUTILE. DO NOT EDIT.

import typing
from ..cog import builder as cogbuilder
from ..models import cloudwatch
from ..models import dashboard


class MetricStat(cogbuilder.Builder[cloudwatch.MetricStat]):    
    _internal: cloudwatch.MetricStat

    def __init__(self):
        self._internal = cloudwatch.MetricStat()

    def build(self) -> cloudwatch.MetricStat:
        """
        Builds the object.
        """
        return self._internal    
    
    def region(self, region: str) -> typing.Self:    
        """
        AWS region to query for the metric
        """
            
        self._internal.region = region
    
        return self
    
    def namespace(self, namespace: str) -> typing.Self:    
        """
        A namespace is a container for CloudWatch metrics. Metrics in different namespaces are isolated from each other, so that metrics from different applications are not mistakenly aggregated into the same statistics. For example, Amazon EC2 uses the AWS/EC2 namespace.
        """
            
        self._internal.namespace = namespace
    
        return self
    
    def metric_name(self, metric_name: str) -> typing.Self:    
        """
        Name of the metric
        """
            
        self._internal.metric_name = metric_name
    
        return self
    
    def dimensions(self, dimensions: cloudwatch.Dimensions) -> typing.Self:    
        """
        The dimensions of the metric
        """
            
        self._internal.dimensions = dimensions
    
        return self
    
    def match_exact(self, match_exact: bool) -> typing.Self:    
        """
        Only show metrics that exactly match all defined dimension names.
        """
            
        self._internal.match_exact = match_exact
    
        return self
    
    def period(self, period: str) -> typing.Self:    
        """
        The length of time associated with a specific Amazon CloudWatch statistic. Can be specified by a number of seconds, 'auto', or as a duration string e.g. '15m' being 15 minutes
        """
            
        self._internal.period = period
    
        return self
    
    def account_id(self, account_id: str) -> typing.Self:    
        """
        The ID of the AWS account to query for the metric, specifying `all` will query all accounts that the monitoring account is permitted to query.
        """
            
        self._internal.account_id = account_id
    
        return self
    
    def statistic(self, statistic: str) -> typing.Self:    
        """
        Metric data aggregations over specified periods of time. For detailed definitions of the statistics supported by CloudWatch, see https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Statistics-definitions.html.
        """
            
        self._internal.statistic = statistic
    
        return self
    
    def statistics(self, statistics: list[str]) -> typing.Self:    
        """
        @deprecated use statistic
        """
            
        self._internal.statistics = statistics
    
        return self
    

class CloudWatchMetricsQuery(cogbuilder.Builder[cloudwatch.CloudWatchMetricsQuery]):    
    """
    Shape of a CloudWatch Metrics query
    """
    
    _internal: cloudwatch.CloudWatchMetricsQuery

    def __init__(self):
        self._internal = cloudwatch.CloudWatchMetricsQuery()

    def build(self) -> cloudwatch.CloudWatchMetricsQuery:
        """
        Builds the object.
        """
        return self._internal    
    
    def query_mode(self, query_mode: cloudwatch.CloudWatchQueryMode) -> typing.Self:    
        """
        Whether a query is a Metrics, Logs, or Annotations query
        """
            
        self._internal.query_mode = query_mode
    
        return self
    
    def metric_query_type(self, metric_query_type: cloudwatch.MetricQueryType) -> typing.Self:    
        """
        Whether to use a metric search or metric query. Metric query is referred to as "Metrics Insights" in the AWS console.
        """
            
        self._internal.metric_query_type = metric_query_type
    
        return self
    
    def metric_editor_mode(self, metric_editor_mode: cloudwatch.MetricEditorMode) -> typing.Self:    
        """
        Whether to use the query builder or code editor to create the query
        """
            
        self._internal.metric_editor_mode = metric_editor_mode
    
        return self
    
    def id_val(self, id_val: str) -> typing.Self:    
        """
        ID can be used to reference other queries in math expressions. The ID can include numbers, letters, and underscore, and must start with a lowercase letter.
        """
            
        self._internal.id_val = id_val
    
        return self
    
    def alias(self, alias: str) -> typing.Self:    
        """
        Deprecated: use label
        @deprecated use label
        """
            
        self._internal.alias = alias
    
        return self
    
    def label(self, label: str) -> typing.Self:    
        """
        Change the time series legend names using dynamic labels. See https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/graph-dynamic-labels.html for more details.
        """
            
        self._internal.label = label
    
        return self
    
    def expression(self, expression: str) -> typing.Self:    
        """
        Math expression query
        """
            
        self._internal.expression = expression
    
        return self
    
    def sql_expression(self, sql_expression: str) -> typing.Self:    
        """
        When the metric query type is `metricQueryType` is set to `Query`, this field is used to specify the query string.
        """
            
        self._internal.sql_expression = sql_expression
    
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
    
    def region(self, region: str) -> typing.Self:    
        """
        AWS region to query for the metric
        """
            
        self._internal.region = region
    
        return self
    
    def namespace(self, namespace: str) -> typing.Self:    
        """
        A namespace is a container for CloudWatch metrics. Metrics in different namespaces are isolated from each other, so that metrics from different applications are not mistakenly aggregated into the same statistics. For example, Amazon EC2 uses the AWS/EC2 namespace.
        """
            
        self._internal.namespace = namespace
    
        return self
    
    def metric_name(self, metric_name: str) -> typing.Self:    
        """
        Name of the metric
        """
            
        self._internal.metric_name = metric_name
    
        return self
    
    def dimensions(self, dimensions: cloudwatch.Dimensions) -> typing.Self:    
        """
        The dimensions of the metric
        """
            
        self._internal.dimensions = dimensions
    
        return self
    
    def match_exact(self, match_exact: bool) -> typing.Self:    
        """
        Only show metrics that exactly match all defined dimension names.
        """
            
        self._internal.match_exact = match_exact
    
        return self
    
    def period(self, period: str) -> typing.Self:    
        """
        The length of time associated with a specific Amazon CloudWatch statistic. Can be specified by a number of seconds, 'auto', or as a duration string e.g. '15m' being 15 minutes
        """
            
        self._internal.period = period
    
        return self
    
    def account_id(self, account_id: str) -> typing.Self:    
        """
        The ID of the AWS account to query for the metric, specifying `all` will query all accounts that the monitoring account is permitted to query.
        """
            
        self._internal.account_id = account_id
    
        return self
    
    def statistic(self, statistic: str) -> typing.Self:    
        """
        Metric data aggregations over specified periods of time. For detailed definitions of the statistics supported by CloudWatch, see https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Statistics-definitions.html.
        """
            
        self._internal.statistic = statistic
    
        return self
    
    def sql(self, sql: cogbuilder.Builder[cloudwatch.SQLExpression]) -> typing.Self:    
        """
        When the metric query type is `metricQueryType` is set to `Query` and the `metricEditorMode` is set to `Builder`, this field is used to build up an object representation of a SQL query.
        """
            
        sql_resource = sql.build()
        self._internal.sql = sql_resource
    
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
    
    def statistics(self, statistics: list[str]) -> typing.Self:    
        """
        @deprecated use statistic
        """
            
        self._internal.statistics = statistics
    
        return self
    

class SQLExpression(cogbuilder.Builder[cloudwatch.SQLExpression]):    
    _internal: cloudwatch.SQLExpression

    def __init__(self):
        self._internal = cloudwatch.SQLExpression()

    def build(self) -> cloudwatch.SQLExpression:
        """
        Builds the object.
        """
        return self._internal    
    
    def select(self, select: cogbuilder.Builder[cloudwatch.QueryEditorFunctionExpression]) -> typing.Self:    
        """
        SELECT part of the SQL expression
        """
            
        select_resource = select.build()
        self._internal.select = select_resource
    
        return self
    
    def from_val(self, from_val: typing.Union[cogbuilder.Builder[cloudwatch.QueryEditorPropertyExpression], cogbuilder.Builder[cloudwatch.QueryEditorFunctionExpression]]) -> typing.Self:    
        """
        FROM part of the SQL expression
        """
            
        from_val_resource = from_val.build()
        self._internal.from_val = from_val_resource
    
        return self
    
    def where(self, where: cogbuilder.Builder[cloudwatch.QueryEditorArrayExpression]) -> typing.Self:    
        """
        WHERE part of the SQL expression
        """
            
        where_resource = where.build()
        self._internal.where = where_resource
    
        return self
    
    def group_by(self, group_by: cogbuilder.Builder[cloudwatch.QueryEditorArrayExpression]) -> typing.Self:    
        """
        GROUP BY part of the SQL expression
        """
            
        group_by_resource = group_by.build()
        self._internal.group_by = group_by_resource
    
        return self
    
    def order_by(self, order_by: cogbuilder.Builder[cloudwatch.QueryEditorFunctionExpression]) -> typing.Self:    
        """
        ORDER BY part of the SQL expression
        """
            
        order_by_resource = order_by.build()
        self._internal.order_by = order_by_resource
    
        return self
    
    def order_by_direction(self, order_by_direction: str) -> typing.Self:    
        """
        The sort order of the SQL expression, `ASC` or `DESC`
        """
            
        self._internal.order_by_direction = order_by_direction
    
        return self
    
    def limit(self, limit: int) -> typing.Self:    
        """
        LIMIT part of the SQL expression
        """
            
        self._internal.limit = limit
    
        return self
    

class QueryEditorFunctionExpression(cogbuilder.Builder[cloudwatch.QueryEditorFunctionExpression]):    
    _internal: cloudwatch.QueryEditorFunctionExpression

    def __init__(self):
        self._internal = cloudwatch.QueryEditorFunctionExpression()        
        self._internal.type_val = "function"

    def build(self) -> cloudwatch.QueryEditorFunctionExpression:
        """
        Builds the object.
        """
        return self._internal    
    
    def name(self, name: str) -> typing.Self:        
        self._internal.name = name
    
        return self
    
    def parameters(self, parameters: list[cogbuilder.Builder[cloudwatch.QueryEditorFunctionParameterExpression]]) -> typing.Self:        
        parameters_resources = [r1.build() for r1 in parameters]
        self._internal.parameters = parameters_resources
    
        return self
    

class QueryEditorFunctionParameterExpression(cogbuilder.Builder[cloudwatch.QueryEditorFunctionParameterExpression]):    
    _internal: cloudwatch.QueryEditorFunctionParameterExpression

    def __init__(self):
        self._internal = cloudwatch.QueryEditorFunctionParameterExpression()        
        self._internal.type_val = "functionParameter"

    def build(self) -> cloudwatch.QueryEditorFunctionParameterExpression:
        """
        Builds the object.
        """
        return self._internal    
    
    def name(self, name: str) -> typing.Self:        
        self._internal.name = name
    
        return self
    

class QueryEditorPropertyExpression(cogbuilder.Builder[cloudwatch.QueryEditorPropertyExpression]):    
    _internal: cloudwatch.QueryEditorPropertyExpression

    def __init__(self):
        self._internal = cloudwatch.QueryEditorPropertyExpression()        
        self._internal.type_val = "property"

    def build(self) -> cloudwatch.QueryEditorPropertyExpression:
        """
        Builds the object.
        """
        return self._internal    
    
    def property_val(self, property_val: cogbuilder.Builder[cloudwatch.QueryEditorProperty]) -> typing.Self:        
        property_val_resource = property_val.build()
        self._internal.property_val = property_val_resource
    
        return self
    

class QueryEditorGroupByExpression(cogbuilder.Builder[cloudwatch.QueryEditorGroupByExpression]):    
    _internal: cloudwatch.QueryEditorGroupByExpression

    def __init__(self):
        self._internal = cloudwatch.QueryEditorGroupByExpression()        
        self._internal.type_val = "groupBy"

    def build(self) -> cloudwatch.QueryEditorGroupByExpression:
        """
        Builds the object.
        """
        return self._internal    
    
    def property_val(self, property_val: cogbuilder.Builder[cloudwatch.QueryEditorProperty]) -> typing.Self:        
        property_val_resource = property_val.build()
        self._internal.property_val = property_val_resource
    
        return self
    

class QueryEditorOperatorExpression(cogbuilder.Builder[cloudwatch.QueryEditorOperatorExpression]):    
    _internal: cloudwatch.QueryEditorOperatorExpression

    def __init__(self):
        self._internal = cloudwatch.QueryEditorOperatorExpression()        
        self._internal.type_val = "operator"

    def build(self) -> cloudwatch.QueryEditorOperatorExpression:
        """
        Builds the object.
        """
        return self._internal    
    
    def property_val(self, property_val: cogbuilder.Builder[cloudwatch.QueryEditorProperty]) -> typing.Self:        
        property_val_resource = property_val.build()
        self._internal.property_val = property_val_resource
    
        return self
    
    def operator(self, operator: cogbuilder.Builder[cloudwatch.QueryEditorOperator]) -> typing.Self:    
        """
        TS type is operator: QueryEditorOperator<QueryEditorOperatorValueType>, extended in veneer
        """
            
        operator_resource = operator.build()
        self._internal.operator = operator_resource
    
        return self
    

class QueryEditorOperator(cogbuilder.Builder[cloudwatch.QueryEditorOperator]):    
    """
    TS type is QueryEditorOperator<T extends QueryEditorOperatorValueType>, extended in veneer
    """
    
    _internal: cloudwatch.QueryEditorOperator

    def __init__(self):
        self._internal = cloudwatch.QueryEditorOperator()

    def build(self) -> cloudwatch.QueryEditorOperator:
        """
        Builds the object.
        """
        return self._internal    
    
    def name(self, name: str) -> typing.Self:        
        self._internal.name = name
    
        return self
    
    def value(self, value: typing.Union[str, bool, int, list[cloudwatch.QueryEditorOperatorType]]) -> typing.Self:        
        self._internal.value = value
    
        return self
    

class QueryEditorProperty(cogbuilder.Builder[cloudwatch.QueryEditorProperty]):    
    _internal: cloudwatch.QueryEditorProperty

    def __init__(self):
        self._internal = cloudwatch.QueryEditorProperty()

    def build(self) -> cloudwatch.QueryEditorProperty:
        """
        Builds the object.
        """
        return self._internal    
    
    def type_val(self, type_val: cloudwatch.QueryEditorPropertyType) -> typing.Self:        
        self._internal.type_val = type_val
    
        return self
    
    def name(self, name: str) -> typing.Self:        
        self._internal.name = name
    
        return self
    

class QueryEditorArrayExpression(cogbuilder.Builder[cloudwatch.QueryEditorArrayExpression]):    
    _internal: cloudwatch.QueryEditorArrayExpression

    def __init__(self):
        self._internal = cloudwatch.QueryEditorArrayExpression()

    def build(self) -> cloudwatch.QueryEditorArrayExpression:
        """
        Builds the object.
        """
        return self._internal    
    
    def type_val(self, type_val: typing.Literal["and", "or"]) -> typing.Self:        
        self._internal.type_val = type_val
    
        return self
    
    def expressions(self, expressions: list[cloudwatch.QueryEditorExpression]) -> typing.Self:        
        self._internal.expressions = expressions
    
        return self
    

class CloudWatchLogsQuery(cogbuilder.Builder[cloudwatch.CloudWatchLogsQuery]):    
    """
    Shape of a CloudWatch Logs query
    """
    
    _internal: cloudwatch.CloudWatchLogsQuery

    def __init__(self):
        self._internal = cloudwatch.CloudWatchLogsQuery()

    def build(self) -> cloudwatch.CloudWatchLogsQuery:
        """
        Builds the object.
        """
        return self._internal    
    
    def query_mode(self, query_mode: cloudwatch.CloudWatchQueryMode) -> typing.Self:    
        """
        Whether a query is a Metrics, Logs, or Annotations query
        """
            
        self._internal.query_mode = query_mode
    
        return self
    
    def id_val(self, id_val: str) -> typing.Self:        
        self._internal.id_val = id_val
    
        return self
    
    def region(self, region: str) -> typing.Self:    
        """
        AWS region to query for the logs
        """
            
        self._internal.region = region
    
        return self
    
    def expression(self, expression: str) -> typing.Self:    
        """
        The CloudWatch Logs Insights query to execute
        """
            
        self._internal.expression = expression
    
        return self
    
    def stats_groups(self, stats_groups: list[str]) -> typing.Self:    
        """
        Fields to group the results by, this field is automatically populated whenever the query is updated
        """
            
        self._internal.stats_groups = stats_groups
    
        return self
    
    def log_groups(self, log_groups: list[cogbuilder.Builder[cloudwatch.LogGroup]]) -> typing.Self:    
        """
        Log groups to query
        """
            
        log_groups_resources = [r1.build() for r1 in log_groups]
        self._internal.log_groups = log_groups_resources
    
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
    
    def log_group_names(self, log_group_names: list[str]) -> typing.Self:    
        """
        @deprecated use logGroups
        """
            
        self._internal.log_group_names = log_group_names
    
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
    

class LogGroup(cogbuilder.Builder[cloudwatch.LogGroup]):    
    _internal: cloudwatch.LogGroup

    def __init__(self):
        self._internal = cloudwatch.LogGroup()

    def build(self) -> cloudwatch.LogGroup:
        """
        Builds the object.
        """
        return self._internal    
    
    def arn(self, arn: str) -> typing.Self:    
        """
        ARN of the log group
        """
            
        self._internal.arn = arn
    
        return self
    
    def name(self, name: str) -> typing.Self:    
        """
        Name of the log group
        """
            
        self._internal.name = name
    
        return self
    
    def account_id(self, account_id: str) -> typing.Self:    
        """
        AccountId of the log group
        """
            
        self._internal.account_id = account_id
    
        return self
    
    def account_label(self, account_label: str) -> typing.Self:    
        """
        Label of the log group
        """
            
        self._internal.account_label = account_label
    
        return self
    

class CloudWatchAnnotationQuery(cogbuilder.Builder[cloudwatch.CloudWatchAnnotationQuery]):    
    """
    Shape of a CloudWatch Annotation query
    TS type is CloudWatchDefaultQuery = Omit<CloudWatchLogsQuery, 'queryMode'> & CloudWatchMetricsQuery, declared in veneer
    #CloudWatchDefaultQuery: #CloudWatchLogsQuery & #CloudWatchMetricsQuery @cuetsy(kind="type")
    """
    
    _internal: cloudwatch.CloudWatchAnnotationQuery

    def __init__(self):
        self._internal = cloudwatch.CloudWatchAnnotationQuery()

    def build(self) -> cloudwatch.CloudWatchAnnotationQuery:
        """
        Builds the object.
        """
        return self._internal    
    
    def query_mode(self, query_mode: cloudwatch.CloudWatchQueryMode) -> typing.Self:    
        """
        Whether a query is a Metrics, Logs, or Annotations query
        """
            
        self._internal.query_mode = query_mode
    
        return self
    
    def prefix_matching(self, prefix_matching: bool) -> typing.Self:    
        """
        Enable matching on the prefix of the action name or alarm name, specify the prefixes with actionPrefix and/or alarmNamePrefix
        """
            
        self._internal.prefix_matching = prefix_matching
    
        return self
    
    def action_prefix(self, action_prefix: str) -> typing.Self:    
        """
        Use this parameter to filter the results of the operation to only those alarms
        that use a certain alarm action. For example, you could specify the ARN of
        an SNS topic to find all alarms that send notifications to that topic.
        e.g. `arn:aws:sns:us-east-1:123456789012:my-app-` would match `arn:aws:sns:us-east-1:123456789012:my-app-action`
        but not match `arn:aws:sns:us-east-1:123456789012:your-app-action`
        """
            
        self._internal.action_prefix = action_prefix
    
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
    
    def region(self, region: str) -> typing.Self:    
        """
        AWS region to query for the metric
        """
            
        self._internal.region = region
    
        return self
    
    def namespace(self, namespace: str) -> typing.Self:    
        """
        A namespace is a container for CloudWatch metrics. Metrics in different namespaces are isolated from each other, so that metrics from different applications are not mistakenly aggregated into the same statistics. For example, Amazon EC2 uses the AWS/EC2 namespace.
        """
            
        self._internal.namespace = namespace
    
        return self
    
    def metric_name(self, metric_name: str) -> typing.Self:    
        """
        Name of the metric
        """
            
        self._internal.metric_name = metric_name
    
        return self
    
    def dimensions(self, dimensions: cloudwatch.Dimensions) -> typing.Self:    
        """
        The dimensions of the metric
        """
            
        self._internal.dimensions = dimensions
    
        return self
    
    def match_exact(self, match_exact: bool) -> typing.Self:    
        """
        Only show metrics that exactly match all defined dimension names.
        """
            
        self._internal.match_exact = match_exact
    
        return self
    
    def period(self, period: str) -> typing.Self:    
        """
        The length of time associated with a specific Amazon CloudWatch statistic. Can be specified by a number of seconds, 'auto', or as a duration string e.g. '15m' being 15 minutes
        """
            
        self._internal.period = period
    
        return self
    
    def account_id(self, account_id: str) -> typing.Self:    
        """
        The ID of the AWS account to query for the metric, specifying `all` will query all accounts that the monitoring account is permitted to query.
        """
            
        self._internal.account_id = account_id
    
        return self
    
    def statistic(self, statistic: str) -> typing.Self:    
        """
        Metric data aggregations over specified periods of time. For detailed definitions of the statistics supported by CloudWatch, see https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Statistics-definitions.html.
        """
            
        self._internal.statistic = statistic
    
        return self
    
    def alarm_name_prefix(self, alarm_name_prefix: str) -> typing.Self:    
        """
        An alarm name prefix. If you specify this parameter, you receive information
        about all alarms that have names that start with this prefix.
        e.g. `my-team-service-` would match `my-team-service-high-cpu` but not match `your-team-service-high-cpu`
        """
            
        self._internal.alarm_name_prefix = alarm_name_prefix
    
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
    
    def statistics(self, statistics: list[str]) -> typing.Self:    
        """
        @deprecated use statistic
        """
            
        self._internal.statistics = statistics
    
        return self
    