// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.azuremonitor;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

// @deprecated Use MetricNamespaceQuery instead
public class MetricDefinitionsQuery {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("rawQuery")
    public String rawQuery;
    @JsonProperty("kind")
    public String kind;
    @JsonProperty("subscription")
    public String subscription;
    @JsonProperty("resourceGroup")
    public String resourceGroup;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("metricNamespace")
    public String metricNamespace;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("resourceName")
    public String resourceName;
    public MetricDefinitionsQuery() {
    }
    
    public MetricDefinitionsQuery(String rawQuery,String kind,String subscription,String resourceGroup,String metricNamespace,String resourceName) {
        this.rawQuery = rawQuery;
        this.kind = kind;
        this.subscription = subscription;
        this.resourceGroup = resourceGroup;
        this.metricNamespace = metricNamespace;
        this.resourceName = resourceName;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
