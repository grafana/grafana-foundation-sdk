// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.canvas;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;
import java.util.List;

public class CanvasElementOptions {
    @JsonProperty("name")
    public String name;
    @JsonProperty("type")
    public String type;
    // TODO: figure out how to define this (element config(s))
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("config")
    public Object config;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("constraint")
    public Constraint constraint;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("placement")
    public Placement placement;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("background")
    public BackgroundConfig background;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("border")
    public LineConfig border;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("connections")
    public List<CanvasConnection> connections;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<CanvasElementOptions> {
        protected final CanvasElementOptions internal;
        
        public Builder() {
            this.internal = new CanvasElementOptions();
        }
    public Builder name(String name) {
    this.internal.name = name;
        return this;
    }
    
    public Builder type(String type) {
    this.internal.type = type;
        return this;
    }
    
    public Builder config(Object config) {
    this.internal.config = config;
        return this;
    }
    
    public Builder constraint(com.grafana.foundation.cog.Builder<Constraint> constraint) {
    this.internal.constraint = constraint.build();
        return this;
    }
    
    public Builder placement(com.grafana.foundation.cog.Builder<Placement> placement) {
    this.internal.placement = placement.build();
        return this;
    }
    
    public Builder background(com.grafana.foundation.cog.Builder<BackgroundConfig> background) {
    this.internal.background = background.build();
        return this;
    }
    
    public Builder border(com.grafana.foundation.cog.Builder<LineConfig> border) {
    this.internal.border = border.build();
        return this;
    }
    
    public Builder connections(com.grafana.foundation.cog.Builder<List<CanvasConnection>> connections) {
    this.internal.connections = connections.build();
        return this;
    }
    public CanvasElementOptions build() {
            return this.internal;
        }
    }
}
