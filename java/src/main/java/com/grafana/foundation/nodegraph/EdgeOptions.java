// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.nodegraph;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class EdgeOptions {
    // Unit for the main stat to override what ever is set in the data frame.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("mainStatUnit")
    public String mainStatUnit;
    // Unit for the secondary stat to override what ever is set in the data frame.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("secondaryStatUnit")
    public String secondaryStatUnit;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<EdgeOptions> {
        protected final EdgeOptions internal;
        
        public Builder() {
            this.internal = new EdgeOptions();
        }
    public Builder mainStatUnit(String mainStatUnit) {
    this.internal.mainStatUnit = mainStatUnit;
        return this;
    }
    
    public Builder secondaryStatUnit(String secondaryStatUnit) {
    this.internal.secondaryStatUnit = secondaryStatUnit;
        return this;
    }
    public EdgeOptions build() {
            return this.internal;
        }
    }
}
