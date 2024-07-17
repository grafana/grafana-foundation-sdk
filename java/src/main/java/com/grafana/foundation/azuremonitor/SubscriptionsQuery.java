// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.azuremonitor;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

public class SubscriptionsQuery { 
    @JsonProperty("rawQuery")
    public String rawQuery; 
    @JsonProperty("kind")
    public String kind;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<SubscriptionsQuery> {
        private final SubscriptionsQuery internal;
        
        public Builder() {
            this.internal = new SubscriptionsQuery();
    this.internal.kind = "SubscriptionsQuery";
        }
    public Builder rawQuery(String rawQuery) {
    this.internal.rawQuery = rawQuery;
        return this;
    }
    public SubscriptionsQuery build() {
            return this.internal;
        }
    }
}
