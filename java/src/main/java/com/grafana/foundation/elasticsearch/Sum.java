// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.databind.annotation.JsonSerialize;

public class Sum {
    @JsonProperty("type")
    public String type;
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("field")
    public String field;
    @JsonProperty("id")
    public String id;
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("settings")
    public ElasticsearchSumSettings settings;
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("hide")
    public Boolean hide;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<Sum> {
        private final Sum internal;
        
        public Builder() {
            this.internal = new Sum();
    this.internal.type = "sum";
        }
    public Builder field(String field) {
    this.internal.field = field;
        return this;
    }
    
    public Builder id(String id) {
    this.internal.id = id;
        return this;
    }
    
    public Builder settings(com.grafana.foundation.cog.Builder<ElasticsearchSumSettings> settings) {
    this.internal.settings = settings.build();
        return this;
    }
    
    public Builder hide(Boolean hide) {
    this.internal.hide = hide;
        return this;
    }
    public Sum build() {
            return this.internal;
        }
    }
}
