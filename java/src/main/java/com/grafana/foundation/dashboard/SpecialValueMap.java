// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboard;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

// Maps special values like Null, NaN (not a number), and boolean values like true and false to a display text and color.
// See SpecialValueMatch to see the list of special values.
// For example, you can configure a special value mapping so that null values appear as N/A.
public class SpecialValueMap {
    @JsonProperty("type")
    public MappingType type;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("options")
    public DashboardSpecialValueMapOptions options;
    public SpecialValueMap() {
        this.type = MappingType.SPECIAL_VALUE;
    }
    public SpecialValueMap(DashboardSpecialValueMapOptions options) {
        this.type = MappingType.SPECIAL_VALUE;
        this.options = options;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
