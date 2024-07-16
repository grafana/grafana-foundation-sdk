// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.azuremonitor;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

public class AppInsightsGroupByQuery { 
    @JsonProperty("rawQuery")
    public String rawQuery; 
    @JsonProperty("kind")
    public String kind; 
    @JsonProperty("metricName")
    public String metricName;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<AppInsightsGroupByQuery> {
        private final AppInsightsGroupByQuery internal;
        
        public Builder() {
            this.internal = new AppInsightsGroupByQuery();
    this.internal.kind = "AppInsightsGroupByQuery";
        }
    public Builder rawQuery(String rawQuery) {
    this.internal.rawQuery = rawQuery;
        return this;
    }
    
    public Builder metricName(String metricName) {
    this.internal.metricName = metricName;
        return this;
    }
    public AppInsightsGroupByQuery build() {
            return this.internal;
        }
    }
}