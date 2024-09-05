// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.databind.annotation.JsonSerialize;

public class SerialDiff {
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("pipelineAgg")
    public String pipelineAgg;
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("field")
    public String field;
    @JsonProperty("type")
    public String type;
    @JsonProperty("id")
    public String id;
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("settings")
    public ElasticsearchSerialDiffSettings settings;
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("hide")
    public Boolean hide;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<SerialDiff> {
        private final SerialDiff internal;
        
        public Builder() {
            this.internal = new SerialDiff();
    this.internal.type = "serial_diff";
        }
    public Builder pipelineAgg(String pipelineAgg) {
    this.internal.pipelineAgg = pipelineAgg;
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
    
    public Builder settings(com.grafana.foundation.cog.Builder<ElasticsearchSerialDiffSettings> settings) {
    this.internal.settings = settings.build();
        return this;
    }
    
    public Builder hide(Boolean hide) {
    this.internal.hide = hide;
        return this;
    }
    public SerialDiff build() {
            return this.internal;
        }
    }
}
