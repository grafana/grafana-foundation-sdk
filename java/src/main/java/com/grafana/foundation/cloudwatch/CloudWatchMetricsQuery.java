// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.cloudwatch;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;
import java.util.Map;
import com.grafana.foundation.dashboard.DataSourceRef;
import java.util.List;

// Shape of a CloudWatch Metrics query
public class CloudWatchMetricsQuery implements com.grafana.foundation.cog.variants.Dataquery {
    // Whether a query is a Metrics, Logs, or Annotations query
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("queryMode")
    public CloudWatchQueryMode queryMode;
    // Whether to use a metric search or metric insights query
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("metricQueryType")
    public MetricQueryType metricQueryType;
    // Whether to use the query builder or code editor to create the query
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("metricEditorMode")
    public MetricEditorMode metricEditorMode;
    // ID can be used to reference other queries in math expressions. The ID can include numbers, letters, and underscore, and must start with a lowercase letter.
    @JsonProperty("id")
    public String id;
    // Deprecated: use label
    // @deprecated use label
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("alias")
    public String alias;
    // Change the time series legend names using dynamic labels. See https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/graph-dynamic-labels.html for more details.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("label")
    public String label;
    // Math expression query
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("expression")
    public String expression;
    // When the metric query type is set to `Insights`, this field is used to specify the query string.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("sqlExpression")
    public String sqlExpression;
    // A unique identifier for the query within the list of targets.
    // In server side expressions, the refId is used as a variable name to identify results.
    // By default, the UI will assign A->Z; however setting meaningful names may be useful.
    @JsonProperty("refId")
    public String refId;
    // If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("hide")
    public Boolean hide;
    // Specify the query flavor
    // TODO make this required and give it a default
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("queryType")
    public String queryType;
    // AWS region to query for the metric
    @JsonProperty("region")
    public String region;
    // A namespace is a container for CloudWatch metrics. Metrics in different namespaces are isolated from each other, so that metrics from different applications are not mistakenly aggregated into the same statistics. For example, Amazon EC2 uses the AWS/EC2 namespace.
    @JsonProperty("namespace")
    public String namespace;
    // Name of the metric
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("metricName")
    public String metricName;
    // The dimensions of the metric
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("dimensions")
    public Map<String, StringOrArrayOfString> dimensions;
    // Only show metrics that exactly match all defined dimension names.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("matchExact")
    public Boolean matchExact;
    // The length of time associated with a specific Amazon CloudWatch statistic. Can be specified by a number of seconds, 'auto', or as a duration string e.g. '15m' being 15 minutes
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("period")
    public String period;
    // The ID of the AWS account to query for the metric, specifying `all` will query all accounts that the monitoring account is permitted to query.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("accountId")
    public String accountId;
    // Metric data aggregations over specified periods of time. For detailed definitions of the statistics supported by CloudWatch, see https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Statistics-definitions.html.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("statistic")
    public String statistic;
    // When the metric query type is set to `Insights` and the `metricEditorMode` is set to `Builder`, this field is used to build up an object representation of a SQL query.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("sql")
    public SQLExpression sql;
    // For mixed data sources the selected datasource is on the query level.
    // For non mixed scenarios this is undefined.
    // TODO find a better way to do this ^ that's friendly to schema
    // TODO this shouldn't be unknown but DataSourceRef | null
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("datasource")
    public DataSourceRef datasource;
    // @deprecated use statistic
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("statistics")
    public List<String> statistics;
    public CloudWatchMetricsQuery() {
        this.queryMode = CloudWatchQueryMode.METRICS;
    }
    public CloudWatchMetricsQuery(CloudWatchQueryMode queryMode,MetricQueryType metricQueryType,MetricEditorMode metricEditorMode,String id,String alias,String label,String expression,String sqlExpression,String refId,Boolean hide,String queryType,String region,String namespace,String metricName,Map<String, StringOrArrayOfString> dimensions,Boolean matchExact,String period,String accountId,String statistic,SQLExpression sql,DataSourceRef datasource,List<String> statistics) {
        this.queryMode = queryMode;
        this.metricQueryType = metricQueryType;
        this.metricEditorMode = metricEditorMode;
        this.id = id;
        this.alias = alias;
        this.label = label;
        this.expression = expression;
        this.sqlExpression = sqlExpression;
        this.refId = refId;
        this.hide = hide;
        this.queryType = queryType;
        this.region = region;
        this.namespace = namespace;
        this.metricName = metricName;
        this.dimensions = dimensions;
        this.matchExact = matchExact;
        this.period = period;
        this.accountId = accountId;
        this.statistic = statistic;
        this.sql = sql;
        this.datasource = datasource;
        this.statistics = statistics;
    }
    public String dataqueryName() {
        return "cloudwatch";
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
