// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.azuremonitor;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonSetter;
import com.fasterxml.jackson.annotation.Nulls;
import java.util.List;
import com.fasterxml.jackson.annotation.JsonInclude;

public class BuilderQueryEditorReduceExpressionArray {
    @JsonSetter(nulls = Nulls.AS_EMPTY)
    @JsonProperty("expressions")
    public List<BuilderQueryEditorReduceExpression> expressions;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("type")
    public BuilderQueryEditorExpressionType type;
    public BuilderQueryEditorReduceExpressionArray() {
    }
    public BuilderQueryEditorReduceExpressionArray(List<BuilderQueryEditorReduceExpression> expressions,BuilderQueryEditorExpressionType type) {
        this.expressions = expressions;
        this.type = type;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
