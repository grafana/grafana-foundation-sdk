// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.testdata;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.databind.annotation.JsonSerialize;

public class Key {
    @JsonProperty("tick")
    public Double tick;
    @JsonProperty("type")
    public String type;
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("uid")
    public String uid;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<Key> {
        private final Key internal;
        
        public Builder() {
            this.internal = new Key();
        }
    public Builder tick(Double tick) {
    this.internal.tick = tick;
        return this;
    }
    
    public Builder type(String type) {
    this.internal.type = type;
        return this;
    }
    
    public Builder uid(String uid) {
    this.internal.uid = uid;
        return this;
    }
    public Key build() {
            return this.internal;
        }
    }
}
