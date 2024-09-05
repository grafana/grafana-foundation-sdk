// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.testdata;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.databind.annotation.JsonSerialize;

// TODO: Should this live here given it's not used in the dataquery?
public class Scenario {
    @JsonProperty("id")
    public String id;
    @JsonProperty("name")
    public String name;
    @JsonProperty("stringInput")
    public String stringInput;
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("description")
    public String description;
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("hideAliasField")
    public Boolean hideAliasField;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<Scenario> {
        private final Scenario internal;
        
        public Builder() {
            this.internal = new Scenario();
        }
    public Builder id(String id) {
    this.internal.id = id;
        return this;
    }
    
    public Builder name(String name) {
    this.internal.name = name;
        return this;
    }
    
    public Builder stringInput(String stringInput) {
    this.internal.stringInput = stringInput;
        return this;
    }
    
    public Builder description(String description) {
    this.internal.description = description;
        return this;
    }
    
    public Builder hideAliasField(Boolean hideAliasField) {
    this.internal.hideAliasField = hideAliasField;
        return this;
    }
    public Scenario build() {
            return this.internal;
        }
    }
}
