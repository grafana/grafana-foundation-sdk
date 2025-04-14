// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.azuremonitor;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;
import java.util.List;

public class BuilderQueryEditorReduceExpression {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("property")
    public BuilderQueryEditorProperty property;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("reduce")
    public BuilderQueryEditorProperty reduce;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("parameters")
    public List<BuilderQueryEditorFunctionParameterExpression> parameters;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("focus")
    public Boolean focus;
    public BuilderQueryEditorReduceExpression() {
    }
    public BuilderQueryEditorReduceExpression(BuilderQueryEditorProperty property,BuilderQueryEditorProperty reduce,List<BuilderQueryEditorFunctionParameterExpression> parameters,Boolean focus) {
        this.property = property;
        this.reduce = reduce;
        this.parameters = parameters;
        this.focus = focus;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
