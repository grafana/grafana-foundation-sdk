// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.azuremonitor;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import java.util.LinkedList;
import com.fasterxml.jackson.annotation.JsonSetter;
import com.fasterxml.jackson.annotation.Nulls;
import java.util.List;
import com.fasterxml.jackson.annotation.JsonInclude;

public class BuilderQueryEditorWhereExpressionArray {
    @JsonSetter(nulls = Nulls.AS_EMPTY)
    @JsonProperty("expressions")
    public List<BuilderQueryEditorWhereExpression> expressions;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("type")
    public BuilderQueryEditorExpressionType type;
    public BuilderQueryEditorWhereExpressionArray() {
        this.expressions = new LinkedList<>();
        this.type = BuilderQueryEditorExpressionType.PROPERTY;
    }
    public BuilderQueryEditorWhereExpressionArray(List<BuilderQueryEditorWhereExpression> expressions,BuilderQueryEditorExpressionType type) {
        this.expressions = expressions;
        this.type = type;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
