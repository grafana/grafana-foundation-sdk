// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.azuremonitor;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

// @deprecated Use MetricNamespaceQuery instead
public class MetricDefinitionsQuery { 
    @JsonProperty("rawQuery")
    public String rawQuery; 
    @JsonProperty("kind")
    public String kind; 
    @JsonProperty("subscription")
    public String subscription; 
    @JsonProperty("resourceGroup")
    public String resourceGroup; 
    @JsonProperty("metricNamespace")
    public String metricNamespace; 
    @JsonProperty("resourceName")
    public String resourceName;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<MetricDefinitionsQuery> {
        private final MetricDefinitionsQuery internal;
        
        public Builder() {
            this.internal = new MetricDefinitionsQuery();
    this.internal.kind = "MetricDefinitionsQuery";
        }
    public Builder rawQuery(String rawQuery) {
    this.internal.rawQuery = rawQuery;
        return this;
    }
    
    public Builder subscription(String subscription) {
    this.internal.subscription = subscription;
        return this;
    }
    
    public Builder resourceGroup(String resourceGroup) {
    this.internal.resourceGroup = resourceGroup;
        return this;
    }
    
    public Builder metricNamespace(String metricNamespace) {
    this.internal.metricNamespace = metricNamespace;
        return this;
    }
    
    public Builder resourceName(String resourceName) {
    this.internal.resourceName = resourceName;
        return this;
    }
    public MetricDefinitionsQuery build() {
            return this.internal;
        }
    }
}
