// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.azuremonitor;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class BuilderQueryEditorWhereExpressionItems {
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("property")
    public BuilderQueryEditorProperty property;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("operator")
    public BuilderQueryEditorOperator operator;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("type")
    public BuilderQueryEditorExpressionType type;
    public BuilderQueryEditorWhereExpressionItems() {
        this.property = new com.grafana.foundation.azuremonitor.BuilderQueryEditorPropertyBuilder().build();
        this.operator = new com.grafana.foundation.azuremonitor.BuilderQueryEditorOperatorBuilder().build();
        this.type = BuilderQueryEditorExpressionType.PROPERTY;
    }
    public BuilderQueryEditorWhereExpressionItems(BuilderQueryEditorProperty property,BuilderQueryEditorOperator operator,BuilderQueryEditorExpressionType type) {
        this.property = property;
        this.operator = operator;
        this.type = type;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
