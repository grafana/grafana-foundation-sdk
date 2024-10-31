// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

// TODO docs
public class SingleStatBaseOptions {
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("reduceOptions")
    public ReduceDataOptions reduceOptions;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("text")
    public VizTextDisplayOptions text;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("orientation")
    public VizOrientation orientation;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<SingleStatBaseOptions> {
        protected final SingleStatBaseOptions internal;
        
        public Builder() {
            this.internal = new SingleStatBaseOptions();
        }
    public Builder reduceOptions(com.grafana.foundation.cog.Builder<ReduceDataOptions> reduceOptions) {
    this.internal.reduceOptions = reduceOptions.build();
        return this;
    }
    
    public Builder text(com.grafana.foundation.cog.Builder<VizTextDisplayOptions> text) {
    this.internal.text = text.build();
        return this;
    }
    
    public Builder orientation(VizOrientation orientation) {
    this.internal.orientation = orientation;
        return this;
    }
    public SingleStatBaseOptions build() {
            return this.internal;
        }
    }
}
