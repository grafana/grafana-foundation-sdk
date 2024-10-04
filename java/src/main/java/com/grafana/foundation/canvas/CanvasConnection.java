// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.canvas;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.grafana.foundation.common.ColorDimensionConfig;
import com.grafana.foundation.common.ScaleDimensionConfig;

public class CanvasConnection {
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("source")
    public ConnectionCoordinates source;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("target")
    public ConnectionCoordinates target;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("targetName")
    public String targetName;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("path")
    public ConnectionPath path;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("color")
    public ColorDimensionConfig color;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("size")
    public ScaleDimensionConfig size;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<CanvasConnection> {
        protected final CanvasConnection internal;
        
        public Builder() {
            this.internal = new CanvasConnection();
        }
    public Builder source(com.grafana.foundation.cog.Builder<ConnectionCoordinates> source) {
    this.internal.source = source.build();
        return this;
    }
    
    public Builder target(com.grafana.foundation.cog.Builder<ConnectionCoordinates> target) {
    this.internal.target = target.build();
        return this;
    }
    
    public Builder targetName(String targetName) {
    this.internal.targetName = targetName;
        return this;
    }
    
    public Builder path(ConnectionPath path) {
    this.internal.path = path;
        return this;
    }
    
    public Builder color(com.grafana.foundation.cog.Builder<ColorDimensionConfig> color) {
    this.internal.color = color.build();
        return this;
    }
    
    public Builder size(com.grafana.foundation.cog.Builder<ScaleDimensionConfig> size) {
    this.internal.size = size.build();
        return this;
    }
    public CanvasConnection build() {
            return this.internal;
        }
    }
}
