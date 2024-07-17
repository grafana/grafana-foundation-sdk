// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

public class ExtendedStats { 
    @JsonProperty("type")
    public String type; 
    @JsonProperty("settings")
    public ElasticsearchExtendedStatsSettings settings; 
    @JsonProperty("field")
    public String field; 
    @JsonProperty("id")
    public String id; 
    @JsonProperty("meta")
    public Object meta; 
    @JsonProperty("hide")
    public Boolean hide;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<ExtendedStats> {
        private final ExtendedStats internal;
        
        public Builder() {
            this.internal = new ExtendedStats();
    this.internal.type = "extended_stats";
        }
    public Builder settings(com.grafana.foundation.cog.Builder<ElasticsearchExtendedStatsSettings> settings) {
    this.internal.settings = settings.build();
        return this;
    }
    
    public Builder field(String field) {
    this.internal.field = field;
        return this;
    }
    
    public Builder id(String id) {
    this.internal.id = id;
        return this;
    }
    
    public Builder meta(Object meta) {
    this.internal.meta = meta;
        return this;
    }
    
    public Builder hide(Boolean hide) {
    this.internal.hide = hide;
        return this;
    }
    public ExtendedStats build() {
            return this.internal;
        }
    }
}
