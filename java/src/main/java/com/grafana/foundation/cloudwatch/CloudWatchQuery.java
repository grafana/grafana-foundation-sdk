// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.cloudwatch;

import com.fasterxml.jackson.annotation.JsonUnwrapped;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.databind.annotation.JsonDeserialize;
import com.fasterxml.jackson.annotation.JsonInclude;

@JsonDeserialize(using = CloudWatchQueryDeserializer.class)
public class CloudWatchQuery implements com.grafana.foundation.cog.variants.Dataquery {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonUnwrapped
    protected CloudWatchMetricsQuery cloudWatchMetricsQuery;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonUnwrapped
    protected CloudWatchLogsQuery cloudWatchLogsQuery;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonUnwrapped
    protected CloudWatchAnnotationQuery cloudWatchAnnotationQuery;
    protected CloudWatchQuery() {}
    public static CloudWatchQuery createCloudWatchMetricsQuery(CloudWatchMetricsQuery cloudWatchMetricsQuery) {
        CloudWatchQuery cloudWatchQuery = new CloudWatchQuery();
        cloudWatchQuery.cloudWatchMetricsQuery = cloudWatchMetricsQuery;
        return cloudWatchQuery;
    }
    public static CloudWatchQuery createCloudWatchLogsQuery(CloudWatchLogsQuery cloudWatchLogsQuery) {
        CloudWatchQuery cloudWatchQuery = new CloudWatchQuery();
        cloudWatchQuery.cloudWatchLogsQuery = cloudWatchLogsQuery;
        return cloudWatchQuery;
    }
    public static CloudWatchQuery createCloudWatchAnnotationQuery(CloudWatchAnnotationQuery cloudWatchAnnotationQuery) {
        CloudWatchQuery cloudWatchQuery = new CloudWatchQuery();
        cloudWatchQuery.cloudWatchAnnotationQuery = cloudWatchAnnotationQuery;
        return cloudWatchQuery;
    }
    public String dataqueryName() {
        return "cloudwatch";
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
