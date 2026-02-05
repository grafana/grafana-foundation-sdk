// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

// TODO docs
public class HideableFieldConfig {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("hideFrom")
    public HideSeriesConfig hideFrom;
    public HideableFieldConfig() {
    }
    public HideableFieldConfig(HideSeriesConfig hideFrom) {
        this.hideFrom = hideFrom;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
