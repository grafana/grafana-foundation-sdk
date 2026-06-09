// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.cloudwatch;

import com.fasterxml.jackson.annotation.JsonUnwrapped;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.databind.annotation.JsonDeserialize;
import com.fasterxml.jackson.annotation.JsonInclude;

@JsonDeserialize(using = RequestDeserializer.class)
public class Request implements com.grafana.foundation.cog.variants.Dataquery {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonUnwrapped
    protected MetricsQuery metricsQuery;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonUnwrapped
    protected LogsQuery logsQuery;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonUnwrapped
    protected AnnotationQuery annotationQuery;
    protected Request() {}
    public static Request createMetricsQuery(MetricsQuery metricsQuery) {
        Request request = new Request();
        request.metricsQuery = metricsQuery;
        return request;
    }
    public static Request createLogsQuery(LogsQuery logsQuery) {
        Request request = new Request();
        request.logsQuery = logsQuery;
        return request;
    }
    public static Request createAnnotationQuery(AnnotationQuery annotationQuery) {
        Request request = new Request();
        request.annotationQuery = annotationQuery;
        return request;
    }
    public String dataqueryName() {
        return "cloudwatch";
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
