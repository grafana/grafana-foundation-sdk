// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

public class Filters { 
    @JsonProperty("id")
    public String id; 
    @JsonProperty("type")
    public String type; 
    @JsonProperty("settings")
    public ElasticsearchFiltersSettings settings;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<Filters> {
        private final Filters internal;
        
        public Builder() {
            this.internal = new Filters();
    this.internal.type = "filters";
        }
    public Builder id(String id) {
    this.internal.id = id;
        return this;
    }
    
    public Builder settings(com.grafana.foundation.cog.Builder<ElasticsearchFiltersSettings> settings) {
    this.internal.settings = settings.build();
        return this;
    }
    public Filters build() {
            return this.internal;
        }
    }
}
