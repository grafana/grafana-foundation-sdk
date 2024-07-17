// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.cloudwatch;

import com.fasterxml.jackson.annotation.JsonUnwrapped;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.databind.annotation.JsonDeserialize;

@JsonDeserialize(using = CloudWatchQueryDeserializer.class)
public class CloudWatchQuery implements com.grafana.foundation.cog.variants.Dataquery { 
    @JsonUnwrapped
    public CloudWatchMetricsQuery cloudWatchMetricsQuery; 
    @JsonUnwrapped
    public CloudWatchLogsQuery cloudWatchLogsQuery; 
    @JsonUnwrapped
    public CloudWatchAnnotationQuery cloudWatchAnnotationQuery;
    
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
