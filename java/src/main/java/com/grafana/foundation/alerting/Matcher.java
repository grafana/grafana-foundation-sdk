// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.alerting;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.databind.annotation.JsonSerialize;

public class Matcher {
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("Name")
    public String name;
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("Type")
    public MatchType type;
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("Value")
    public String value;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<Matcher> {
        private final Matcher internal;
        
        public Builder() {
            this.internal = new Matcher();
        }
    public Builder name(String name) {
    this.internal.name = name;
        return this;
    }
    
    public Builder type(MatchType type) {
    this.internal.type = type;
        return this;
    }
    
    public Builder value(String value) {
    this.internal.value = value;
        return this;
    }
    public Matcher build() {
            return this.internal;
        }
    }
}
