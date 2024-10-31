// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

// TODO docs
public class OptionsWithTextFormatting {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("text")
    public VizTextDisplayOptions text;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<OptionsWithTextFormatting> {
        protected final OptionsWithTextFormatting internal;
        
        public Builder() {
            this.internal = new OptionsWithTextFormatting();
        }
    public Builder text(com.grafana.foundation.cog.Builder<VizTextDisplayOptions> text) {
    this.internal.text = text.build();
        return this;
    }
    public OptionsWithTextFormatting build() {
            return this.internal;
        }
    }
}
