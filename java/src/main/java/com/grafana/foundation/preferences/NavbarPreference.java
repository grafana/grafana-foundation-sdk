// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.preferences;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonSetter;
import com.fasterxml.jackson.annotation.Nulls;
import java.util.List;

public class NavbarPreference {
    @JsonSetter(nulls = Nulls.AS_EMPTY)
    @JsonProperty("bookmarkUrls")
    public List<String> bookmarkUrls;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<NavbarPreference> {
        private final NavbarPreference internal;
        
        public Builder() {
            this.internal = new NavbarPreference();
        }
    public Builder bookmarkUrls(List<String> bookmarkUrls) {
    this.internal.bookmarkUrls = bookmarkUrls;
        return this;
    }
    public NavbarPreference build() {
            return this.internal;
        }
    }
}
