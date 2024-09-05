// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.testdata;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.databind.annotation.JsonSerialize;
import java.util.List;

public class USAQuery {
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("fields")
    public List<String> fields;
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("mode")
    public String mode;
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("period")
    public String period;
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("states")
    public List<String> states;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<USAQuery> {
        private final USAQuery internal;
        
        public Builder() {
            this.internal = new USAQuery();
        }
    public Builder fields(List<String> fields) {
    this.internal.fields = fields;
        return this;
    }
    
    public Builder mode(String mode) {
    this.internal.mode = mode;
        return this;
    }
    
    public Builder period(String period) {
    this.internal.period = period;
        return this;
    }
    
    public Builder states(List<String> states) {
    this.internal.states = states;
        return this;
    }
    public USAQuery build() {
            return this.internal;
        }
    }
}
