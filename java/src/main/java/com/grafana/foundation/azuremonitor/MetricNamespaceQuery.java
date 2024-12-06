// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.azuremonitor;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class MetricNamespaceQuery {
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
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<MetricNamespaceQuery> {
        protected final MetricNamespaceQuery internal;
        
        public Builder() {
            this.internal = new MetricNamespaceQuery();
    this.internal.kind = "MetricNamespaceQuery";
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
    public MetricNamespaceQuery build() {
            return this.internal;
        }
    }
}
