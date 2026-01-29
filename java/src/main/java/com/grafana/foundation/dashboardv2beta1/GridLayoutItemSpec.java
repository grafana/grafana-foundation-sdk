// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class GridLayoutItemSpec {
    @JsonProperty("x")
    public Long x;
    @JsonProperty("y")
    public Long y;
    @JsonProperty("width")
    public Long width;
    @JsonProperty("height")
    public Long height;
    // reference to a PanelKind from dashboard.spec.elements Expressed as JSON Schema reference
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("element")
    public ElementReference element;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("repeat")
    public RepeatOptions repeat;
    public GridLayoutItemSpec() {
        this.x = 0L;
        this.y = 0L;
        this.width = 0L;
        this.height = 0L;
        this.element = new com.grafana.foundation.dashboardv2beta1.ElementReference();
    }
    public GridLayoutItemSpec(Long x,Long y,Long width,Long height,ElementReference element,RepeatOptions repeat) {
        this.x = x;
        this.y = y;
        this.width = width;
        this.height = height;
        this.element = element;
        this.repeat = repeat;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
