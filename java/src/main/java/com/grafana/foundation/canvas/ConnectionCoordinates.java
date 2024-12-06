// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.canvas;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

public class ConnectionCoordinates {
    @JsonProperty("x")
    public Double x;
    @JsonProperty("y")
    public Double y;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<ConnectionCoordinates> {
        protected final ConnectionCoordinates internal;
        
        public Builder() {
            this.internal = new ConnectionCoordinates();
        }
    public Builder x(Double x) {
    this.internal.x = x;
        return this;
    }
    
    public Builder y(Double y) {
    this.internal.y = y;
        return this;
    }
    public ConnectionCoordinates build() {
            return this.internal;
        }
    }
}
