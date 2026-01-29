// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import com.fasterxml.jackson.core.JsonGenerator;
import com.fasterxml.jackson.databind.JsonSerializer;
import com.fasterxml.jackson.databind.SerializerProvider;

import java.io.IOException;

public class ConditionalRenderingVariableKindOrConditionalRenderingDataKindOrConditionalRenderingTimeRangeSizeKindSerializer extends JsonSerializer<ConditionalRenderingVariableKindOrConditionalRenderingDataKindOrConditionalRenderingTimeRangeSizeKind> {

    @Override
    public void serialize(ConditionalRenderingVariableKindOrConditionalRenderingDataKindOrConditionalRenderingTimeRangeSizeKind value, JsonGenerator gen, SerializerProvider serializerProvider) throws IOException {
        if (value.conditionalRenderingVariableKind != null) {
            gen.writeObject(value.conditionalRenderingVariableKind);
        }
        else if (value.conditionalRenderingDataKind != null) {
            gen.writeObject(value.conditionalRenderingDataKind);
        }
        else if (value.conditionalRenderingTimeRangeSizeKind != null) {
            gen.writeObject(value.conditionalRenderingTimeRangeSizeKind);
        }
    }
}
