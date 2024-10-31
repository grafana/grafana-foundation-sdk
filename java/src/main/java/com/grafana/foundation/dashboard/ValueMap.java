// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboard;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonSetter;
import com.fasterxml.jackson.annotation.Nulls;
import java.util.Map;

// Maps text values to a color or different display text and color.
// For example, you can configure a value mapping so that all instances of the value 10 appear as Perfection! rather than the number.
public class ValueMap {
    @JsonProperty("type")
    public String type;
    // Map with <value_to_match>: ValueMappingResult. For example: { "10": { text: "Perfection!", color: "green" } }
    @JsonSetter(nulls = Nulls.AS_EMPTY)
    @JsonProperty("options")
    public Map<String, ValueMappingResult> options;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<ValueMap> {
        protected final ValueMap internal;
        
        public Builder() {
            this.internal = new ValueMap();
    this.internal.type = "value";
        }
    public Builder options(Map<String, ValueMappingResult> options) {
    this.internal.options = options;
        return this;
    }
    public ValueMap build() {
            return this.internal;
        }
    }
}
