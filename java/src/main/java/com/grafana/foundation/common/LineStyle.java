// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;

import java.util.List;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

// TODO docs
public class LineStyle { 
    @JsonProperty("fill")
    public LineStyleFill fill; 
    @JsonProperty("dash")
    public List<Double> dash;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<LineStyle> {
        private final LineStyle internal;
        
        public Builder() {
            this.internal = new LineStyle();
        }
    public Builder fill(LineStyleFill fill) {
    this.internal.fill = fill;
        return this;
    }
    
    public Builder dash(List<Double> dash) {
    this.internal.dash = dash;
        return this;
    }
    public LineStyle build() {
            return this.internal;
        }
    }
}
