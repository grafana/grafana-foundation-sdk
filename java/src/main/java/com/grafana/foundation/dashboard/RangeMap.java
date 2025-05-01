// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboard;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

// Maps numerical ranges to a display text and color.
// For example, if a value is within a certain range, you can configure a range value mapping to display Low or High rather than the number.
public class RangeMap {
    @JsonProperty("type")
    public MappingType type;
    // Range to match against and the result to apply when the value is within the range
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("options")
    public DashboardRangeMapOptions options;
    public RangeMap() {
        this.type = MappingType.RANGE_TO_TEXT;
        this.options = new com.grafana.foundation.dashboard.DashboardRangeMapOptionsBuilder().build();
    }
    public RangeMap(DashboardRangeMapOptions options) {
        this.type = MappingType.RANGE_TO_TEXT;
        this.options = options;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
