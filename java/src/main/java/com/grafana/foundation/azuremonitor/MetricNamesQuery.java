// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.azuremonitor;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class MetricNamesQuery {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("rawQuery")
    public String rawQuery;
    @JsonProperty("kind")
    public String kind;
    @JsonProperty("subscription")
    public String subscription;
    @JsonProperty("resourceGroup")
    public String resourceGroup;
    @JsonProperty("resourceName")
    public String resourceName;
    @JsonProperty("metricNamespace")
    public String metricNamespace;
    public MetricNamesQuery() {
    }
    public MetricNamesQuery(String rawQuery,String kind,String subscription,String resourceGroup,String resourceName,String metricNamespace) {
        this.rawQuery = rawQuery;
        this.kind = kind;
        this.subscription = subscription;
        this.resourceGroup = resourceGroup;
        this.resourceName = resourceName;
        this.metricNamespace = metricNamespace;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
