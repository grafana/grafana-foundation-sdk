// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.azuremonitor;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;
import java.util.List;

public class BuilderQueryEditorColumnsExpression {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("columns")
    public List<String> columns;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("type")
    public BuilderQueryEditorExpressionType type;
    public BuilderQueryEditorColumnsExpression() {
        this.type = BuilderQueryEditorExpressionType.PROPERTY;
    }
    public BuilderQueryEditorColumnsExpression(List<String> columns,BuilderQueryEditorExpressionType type) {
        this.columns = columns;
        this.type = type;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
