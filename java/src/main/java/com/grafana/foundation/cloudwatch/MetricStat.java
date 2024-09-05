// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.cloudwatch;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.databind.annotation.JsonSerialize;
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
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("metricName")
    public String metricName;
    // The dimensions of the metric
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("dimensions")
    public Map<String, StringOrArrayOfString> dimensions;
    // Only show metrics that exactly match all defined dimension names.
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("matchExact")
    public Boolean matchExact;
    // The length of time associated with a specific Amazon CloudWatch statistic. Can be specified by a number of seconds, 'auto', or as a duration string e.g. '15m' being 15 minutes
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("period")
    public String period;
    // The ID of the AWS account to query for the metric, specifying `all` will query all accounts that the monitoring account is permitted to query.
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("accountId")
    public String accountId;
    // Metric data aggregations over specified periods of time. For detailed definitions of the statistics supported by CloudWatch, see https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Statistics-definitions.html.
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("statistic")
    public String statistic;
    // @deprecated use statistic
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("statistics")
    public List<String> statistics;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<MetricStat> {
        private final MetricStat internal;
        
        public Builder() {
            this.internal = new MetricStat();
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
    
    public Builder statistics(List<String> statistics) {
    this.internal.statistics = statistics;
        return this;
    }
    public MetricStat build() {
            return this.internal;
        }
    }
}
