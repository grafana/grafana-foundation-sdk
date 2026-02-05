// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import com.fasterxml.jackson.annotation.JsonUnwrapped;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.databind.annotation.JsonDeserialize;
import com.fasterxml.jackson.annotation.JsonInclude;

@JsonDeserialize(using = ConditionalRenderingVariableKindOrConditionalRenderingDataKindOrConditionalRenderingTimeRangeSizeKindDeserializer.class)
public class ConditionalRenderingVariableKindOrConditionalRenderingDataKindOrConditionalRenderingTimeRangeSizeKind {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonUnwrapped
    protected ConditionalRenderingVariableKind conditionalRenderingVariableKind;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonUnwrapped
    protected ConditionalRenderingDataKind conditionalRenderingDataKind;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonUnwrapped
    protected ConditionalRenderingTimeRangeSizeKind conditionalRenderingTimeRangeSizeKind;
    protected ConditionalRenderingVariableKindOrConditionalRenderingDataKindOrConditionalRenderingTimeRangeSizeKind() {}
    public static ConditionalRenderingVariableKindOrConditionalRenderingDataKindOrConditionalRenderingTimeRangeSizeKind createConditionalRenderingVariableKind(ConditionalRenderingVariableKind conditionalRenderingVariableKind) {
        ConditionalRenderingVariableKindOrConditionalRenderingDataKindOrConditionalRenderingTimeRangeSizeKind conditionalRenderingVariableKindOrConditionalRenderingDataKindOrConditionalRenderingTimeRangeSizeKind = new ConditionalRenderingVariableKindOrConditionalRenderingDataKindOrConditionalRenderingTimeRangeSizeKind();
        conditionalRenderingVariableKindOrConditionalRenderingDataKindOrConditionalRenderingTimeRangeSizeKind.conditionalRenderingVariableKind = conditionalRenderingVariableKind;
        return conditionalRenderingVariableKindOrConditionalRenderingDataKindOrConditionalRenderingTimeRangeSizeKind;
    }
    public static ConditionalRenderingVariableKindOrConditionalRenderingDataKindOrConditionalRenderingTimeRangeSizeKind createConditionalRenderingDataKind(ConditionalRenderingDataKind conditionalRenderingDataKind) {
        ConditionalRenderingVariableKindOrConditionalRenderingDataKindOrConditionalRenderingTimeRangeSizeKind conditionalRenderingVariableKindOrConditionalRenderingDataKindOrConditionalRenderingTimeRangeSizeKind = new ConditionalRenderingVariableKindOrConditionalRenderingDataKindOrConditionalRenderingTimeRangeSizeKind();
        conditionalRenderingVariableKindOrConditionalRenderingDataKindOrConditionalRenderingTimeRangeSizeKind.conditionalRenderingDataKind = conditionalRenderingDataKind;
        return conditionalRenderingVariableKindOrConditionalRenderingDataKindOrConditionalRenderingTimeRangeSizeKind;
    }
    public static ConditionalRenderingVariableKindOrConditionalRenderingDataKindOrConditionalRenderingTimeRangeSizeKind createConditionalRenderingTimeRangeSizeKind(ConditionalRenderingTimeRangeSizeKind conditionalRenderingTimeRangeSizeKind) {
        ConditionalRenderingVariableKindOrConditionalRenderingDataKindOrConditionalRenderingTimeRangeSizeKind conditionalRenderingVariableKindOrConditionalRenderingDataKindOrConditionalRenderingTimeRangeSizeKind = new ConditionalRenderingVariableKindOrConditionalRenderingDataKindOrConditionalRenderingTimeRangeSizeKind();
        conditionalRenderingVariableKindOrConditionalRenderingDataKindOrConditionalRenderingTimeRangeSizeKind.conditionalRenderingTimeRangeSizeKind = conditionalRenderingTimeRangeSizeKind;
        return conditionalRenderingVariableKindOrConditionalRenderingDataKindOrConditionalRenderingTimeRangeSizeKind;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
