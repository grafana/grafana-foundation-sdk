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
    public static CloudWatchQuery createCloudWatchMetricsQuery(com.grafana.foundation.cog.Builder<CloudWatchMetricsQuery> cloudWatchMetricsQuery) {
        CloudWatchQuery cloudWatchQuery = new CloudWatchQuery();
        cloudWatchQuery.cloudWatchMetricsQuery = cloudWatchMetricsQuery.build();
        return cloudWatchQuery;
    }
    public static CloudWatchQuery createCloudWatchLogsQuery(com.grafana.foundation.cog.Builder<CloudWatchLogsQuery> cloudWatchLogsQuery) {
        CloudWatchQuery cloudWatchQuery = new CloudWatchQuery();
        cloudWatchQuery.cloudWatchLogsQuery = cloudWatchLogsQuery.build();
        return cloudWatchQuery;
    }
    public static CloudWatchQuery createCloudWatchAnnotationQuery(com.grafana.foundation.cog.Builder<CloudWatchAnnotationQuery> cloudWatchAnnotationQuery) {
        CloudWatchQuery cloudWatchQuery = new CloudWatchQuery();
        cloudWatchQuery.cloudWatchAnnotationQuery = cloudWatchAnnotationQuery.build();
        return cloudWatchQuery;
    }
    public String dataqueryName() {
        return "cloudwatch";
    }
    
    public String toJSON() throws JsonProcessingException {
        if (cloudWatchMetricsQuery != null) {
            ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
            return ow.writeValueAsString(cloudWatchMetricsQuery);
        }
        if (cloudWatchLogsQuery != null) {
            ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
            return ow.writeValueAsString(cloudWatchLogsQuery);
        }
        if (cloudWatchAnnotationQuery != null) {
            ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
            return ow.writeValueAsString(cloudWatchAnnotationQuery);
        }
        
        return null;
    }

}
