// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.azuremonitor;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class BuilderQueryEditorFunctionParameterExpression {
    @JsonProperty("value")
    public String value;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("fieldType")
    public BuilderQueryEditorPropertyType fieldType;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("type")
    public BuilderQueryEditorExpressionType type;
    public BuilderQueryEditorFunctionParameterExpression() {
    }
    public BuilderQueryEditorFunctionParameterExpression(String value,BuilderQueryEditorPropertyType fieldType,BuilderQueryEditorExpressionType type) {
        this.value = value;
        this.fieldType = fieldType;
        this.type = type;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
