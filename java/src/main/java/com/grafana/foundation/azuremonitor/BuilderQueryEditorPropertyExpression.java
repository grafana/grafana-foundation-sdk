// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.azuremonitor;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class BuilderQueryEditorPropertyExpression {
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("property")
    public BuilderQueryEditorProperty property;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("type")
    public BuilderQueryEditorExpressionType type;
    public BuilderQueryEditorPropertyExpression() {
        this.property = new com.grafana.foundation.azuremonitor.BuilderQueryEditorPropertyBuilder().build();
        this.type = BuilderQueryEditorExpressionType.PROPERTY;
    }
    public BuilderQueryEditorPropertyExpression(BuilderQueryEditorProperty property,BuilderQueryEditorExpressionType type) {
        this.property = property;
        this.type = type;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
