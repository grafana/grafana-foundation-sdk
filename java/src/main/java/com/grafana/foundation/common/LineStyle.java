// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;
import java.util.List;

// TODO docs
public class LineStyle {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("fill")
    public LineStyleFill fill;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("dash")
    public List<Double> dash;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<LineStyle> {
        protected final LineStyle internal;
        
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
