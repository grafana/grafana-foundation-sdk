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
    // Whether to use a metric search or metric query. Metric query is referred to as "Metrics Insights" in the AWS console.
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
    // When the metric query type is `metricQueryType` is set to `Query`, this field is used to specify the query string.
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
    // When the metric query type is `metricQueryType` is set to `Query` and the `metricEditorMode` is set to `Builder`, this field is used to build up an object representation of a SQL query.
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
    public String dataqueryName() {
        return "cloudwatch";
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<CloudWatchMetricsQuery> {
        protected final CloudWatchMetricsQuery internal;
        
        public Builder() {
            this.internal = new CloudWatchMetricsQuery();
        this.queryMode(CloudWatchQueryMode.METRICS);
        }
    public Builder queryMode(CloudWatchQueryMode queryMode) {
    this.internal.queryMode = queryMode;
        return this;
    }
    
    public Builder metricQueryType(MetricQueryType metricQueryType) {
    this.internal.metricQueryType = metricQueryType;
        return this;
    }
    
    public Builder metricEditorMode(MetricEditorMode metricEditorMode) {
    this.internal.metricEditorMode = metricEditorMode;
        return this;
    }
    
    public Builder id(String id) {
    this.internal.id = id;
        return this;
    }
    
    public Builder alias(String alias) {
    this.internal.alias = alias;
        return this;
    }
    
    public Builder label(String label) {
    this.internal.label = label;
        return this;
    }
    
    public Builder expression(String expression) {
    this.internal.expression = expression;
        return this;
    }
    
    public Builder sqlExpression(String sqlExpression) {
    this.internal.sqlExpression = sqlExpression;
        return this;
    }
    
    public Builder refId(String refId) {
    this.internal.refId = refId;
        return this;
    }
    
    public Builder hide(Boolean hide) {
    this.internal.hide = hide;
        return this;
    }
    
    public Builder queryType(String queryType) {
    this.internal.queryType = queryType;
        return this;
    }
    
    public Builder region(String region) {
    this.internal.region = region;
        return this;
    }
    
    public Builder namespace(String namespace) {
    this.internal.namespace = namespace;
        return this;
    }
    
    public Builder metricName(String metricName) {
    this.internal.metricName = metricName;
        return this;
    }
    
    public Builder dimensions(Map<String, StringOrArrayOfString> dimensions) {
    this.internal.dimensions = dimensions;
        return this;
    }
    
    public Builder matchExact(Boolean matchExact) {
    this.internal.matchExact = matchExact;
        return this;
    }
    
    public Builder period(String period) {
    this.internal.period = period;
        return this;
    }
    
    public Builder accountId(String accountId) {
    this.internal.accountId = accountId;
        return this;
    }
    
    public Builder statistic(String statistic) {
    this.internal.statistic = statistic;
        return this;
    }
    
    public Builder sql(com.grafana.foundation.cog.Builder<SQLExpression> sql) {
    this.internal.sql = sql.build();
        return this;
    }
    
    public Builder datasource(DataSourceRef datasource) {
    this.internal.datasource = datasource;
        return this;
    }
    
    public Builder statistics(List<String> statistics) {
    this.internal.statistics = statistics;
        return this;
    }
    public CloudWatchMetricsQuery build() {
            return this.internal;
        }
    }
}
