// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.canvas;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class Placement {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("top")
    public Double top;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("left")
    public Double left;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("right")
    public Double right;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("bottom")
    public Double bottom;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("width")
    public Double width;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("height")
    public Double height;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<Placement> {
        protected final Placement internal;
        
        public Builder() {
            this.internal = new Placement();
        }
    public Builder top(Double top) {
    this.internal.top = top;
        return this;
    }
    
    public Builder left(Double left) {
    this.internal.left = left;
        return this;
    }
    
    public Builder right(Double right) {
    this.internal.right = right;
        return this;
    }
    
    public Builder bottom(Double bottom) {
    this.internal.bottom = bottom;
        return this;
    }
    
    public Builder width(Double width) {
    this.internal.width = width;
        return this;
    }
    
    public Builder height(Double height) {
    this.internal.height = height;
        return this;
    }
    public Placement build() {
            return this.internal;
        }
    }
}
