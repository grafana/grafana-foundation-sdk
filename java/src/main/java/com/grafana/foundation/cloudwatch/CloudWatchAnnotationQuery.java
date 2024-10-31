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

// Shape of a CloudWatch Annotation query
// TS type is CloudWatchDefaultQuery = Omit<CloudWatchLogsQuery, 'queryMode'> & CloudWatchMetricsQuery, declared in veneer
// #CloudWatchDefaultQuery: #CloudWatchLogsQuery & #CloudWatchMetricsQuery @cuetsy(kind="type")
public class CloudWatchAnnotationQuery implements com.grafana.foundation.cog.variants.Dataquery {
    // Whether a query is a Metrics, Logs, or Annotations query
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("queryMode")
    public CloudWatchQueryMode queryMode;
    // Enable matching on the prefix of the action name or alarm name, specify the prefixes with actionPrefix and/or alarmNamePrefix
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("prefixMatching")
    public Boolean prefixMatching;
    // Use this parameter to filter the results of the operation to only those alarms
    // that use a certain alarm action. For example, you could specify the ARN of
    // an SNS topic to find all alarms that send notifications to that topic.
    // e.g. `arn:aws:sns:us-east-1:123456789012:my-app-` would match `arn:aws:sns:us-east-1:123456789012:my-app-action`
    // but not match `arn:aws:sns:us-east-1:123456789012:your-app-action`
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("actionPrefix")
    public String actionPrefix;
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
    // An alarm name prefix. If you specify this parameter, you receive information
    // about all alarms that have names that start with this prefix.
    // e.g. `my-team-service-` would match `my-team-service-high-cpu` but not match `your-team-service-high-cpu`
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("alarmNamePrefix")
    public String alarmNamePrefix;
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

    
    public static class Builder implements com.grafana.foundation.cog.Builder<CloudWatchAnnotationQuery> {
        protected final CloudWatchAnnotationQuery internal;
        
        public Builder() {
            this.internal = new CloudWatchAnnotationQuery();
        this.queryMode(CloudWatchQueryMode.ANNOTATIONS);
        }
    public Builder queryMode(CloudWatchQueryMode queryMode) {
    this.internal.queryMode = queryMode;
        return this;
    }
    
    public Builder prefixMatching(Boolean prefixMatching) {
    this.internal.prefixMatching = prefixMatching;
        return this;
    }
    
    public Builder actionPrefix(String actionPrefix) {
    this.internal.actionPrefix = actionPrefix;
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
    
    public Builder alarmNamePrefix(String alarmNamePrefix) {
    this.internal.alarmNamePrefix = alarmNamePrefix;
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
    public CloudWatchAnnotationQuery build() {
            return this.internal;
        }
    }
}
