// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.preferences;

import java.util.List;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

public class NavbarPreference { 
    @JsonProperty("savedItemIds")
    public List<String> savedItemIds;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<NavbarPreference> {
        private final NavbarPreference internal;
        
        public Builder() {
            this.internal = new NavbarPreference();
        }
    public Builder savedItemIds(List<String> savedItemIds) {
    this.internal.savedItemIds = savedItemIds;
        return this;
    }
    public NavbarPreference build() {
            return this.internal;
        }
    }
}
