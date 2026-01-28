// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import java.util.LinkedList;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonSetter;
import com.fasterxml.jackson.annotation.Nulls;
import java.util.List;

public class ConditionalRenderingGroupSpec {
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("visibility")
    public ConditionalRenderingGroupSpecVisibility visibility;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("condition")
    public ConditionalRenderingGroupSpecCondition condition;
    @JsonSetter(nulls = Nulls.AS_EMPTY)
    @JsonProperty("items")
    public List<ConditionalRenderingVariableKindOrConditionalRenderingDataKindOrConditionalRenderingTimeRangeSizeKind> items;
    public ConditionalRenderingGroupSpec() {
        this.visibility = ConditionalRenderingGroupSpecVisibility.SHOW;
        this.condition = ConditionalRenderingGroupSpecCondition.AND;
        this.items = new LinkedList<>();
    }
    public ConditionalRenderingGroupSpec(ConditionalRenderingGroupSpecVisibility visibility,ConditionalRenderingGroupSpecCondition condition,List<ConditionalRenderingVariableKindOrConditionalRenderingDataKindOrConditionalRenderingTimeRangeSizeKind> items) {
        this.visibility = visibility;
        this.condition = condition;
        this.items = items;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
