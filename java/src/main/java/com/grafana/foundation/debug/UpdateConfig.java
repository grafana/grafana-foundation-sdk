// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.debug;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

public class UpdateConfig {
    @JsonProperty("render")
    public Boolean render;
    @JsonProperty("dataChanged")
    public Boolean dataChanged;
    @JsonProperty("schemaChanged")
    public Boolean schemaChanged;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<UpdateConfig> {
        protected final UpdateConfig internal;
        
        public Builder() {
            this.internal = new UpdateConfig();
        }
    public Builder render(Boolean render) {
    this.internal.render = render;
        return this;
    }
    
    public Builder dataChanged(Boolean dataChanged) {
    this.internal.dataChanged = dataChanged;
        return this;
    }
    
    public Builder schemaChanged(Boolean schemaChanged) {
    this.internal.schemaChanged = schemaChanged;
        return this;
    }
    public UpdateConfig build() {
            return this.internal;
        }
    }
}
