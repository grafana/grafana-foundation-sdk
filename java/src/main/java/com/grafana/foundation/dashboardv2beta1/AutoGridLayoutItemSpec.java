// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class AutoGridLayoutItemSpec {
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("element")
    public ElementReference element;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("repeat")
    public AutoGridRepeatOptions repeat;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("conditionalRendering")
    public ConditionalRenderingGroupKind conditionalRendering;
    public AutoGridLayoutItemSpec() {
        this.element = new com.grafana.foundation.dashboardv2beta1.ElementReference();
    }
    public AutoGridLayoutItemSpec(ElementReference element,AutoGridRepeatOptions repeat,ConditionalRenderingGroupKind conditionalRendering) {
        this.element = element;
        this.repeat = repeat;
        this.conditionalRendering = conditionalRendering;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
