// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.azuremonitor;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.databind.annotation.JsonSerialize;

public class AzureMonitorResource {
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("subscription")
    public String subscription;
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("resourceGroup")
    public String resourceGroup;
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("resourceName")
    public String resourceName;
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("metricNamespace")
    public String metricNamespace;
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("region")
    public String region;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<AzureMonitorResource> {
        private final AzureMonitorResource internal;
        
        public Builder() {
            this.internal = new AzureMonitorResource();
        }
    public Builder subscription(String subscription) {
    this.internal.subscription = subscription;
        return this;
    }
    
    public Builder resourceGroup(String resourceGroup) {
    this.internal.resourceGroup = resourceGroup;
        return this;
    }
    
    public Builder resourceName(String resourceName) {
    this.internal.resourceName = resourceName;
        return this;
    }
    
    public Builder metricNamespace(String metricNamespace) {
    this.internal.metricNamespace = metricNamespace;
        return this;
    }
    
    public Builder region(String region) {
    this.internal.region = region;
        return this;
    }
    public AzureMonitorResource build() {
            return this.internal;
        }
    }
}
