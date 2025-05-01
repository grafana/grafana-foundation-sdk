// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.cloudwatch;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonSetter;
import com.fasterxml.jackson.annotation.Nulls;
import java.util.Map;
import java.util.List;

public class MetricStat {
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
    @JsonSetter(nulls = Nulls.AS_EMPTY)
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
    // @deprecated use statistic
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("statistics")
    public List<String> statistics;
    public MetricStat() {
        this.region = "";
        this.namespace = "";
    }
    public MetricStat(String region,String namespace,String metricName,Map<String, StringOrArrayOfString> dimensions,Boolean matchExact,String period,String accountId,String statistic,List<String> statistics) {
        this.region = region;
        this.namespace = namespace;
        this.metricName = metricName;
        this.dimensions = dimensions;
        this.matchExact = matchExact;
        this.period = period;
        this.accountId = accountId;
        this.statistic = statistic;
        this.statistics = statistics;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
