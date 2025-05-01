// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.azuremonitor;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class BuilderQueryEditorOrderByExpression {
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("property")
    public BuilderQueryEditorProperty property;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("order")
    public BuilderQueryEditorOrderByOptions order;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("type")
    public BuilderQueryEditorExpressionType type;
    public BuilderQueryEditorOrderByExpression() {
        this.property = new com.grafana.foundation.azuremonitor.BuilderQueryEditorPropertyBuilder().build();
        this.order = BuilderQueryEditorOrderByOptions.ASC;
        this.type = BuilderQueryEditorExpressionType.PROPERTY;
    }
    public BuilderQueryEditorOrderByExpression(BuilderQueryEditorProperty property,BuilderQueryEditorOrderByOptions order,BuilderQueryEditorExpressionType type) {
        this.property = property;
        this.order = order;
        this.type = type;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
