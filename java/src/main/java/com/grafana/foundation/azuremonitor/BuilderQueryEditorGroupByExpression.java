// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.azuremonitor;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class BuilderQueryEditorGroupByExpression {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("property")
    public BuilderQueryEditorProperty property;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("interval")
    public BuilderQueryEditorProperty interval;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("focus")
    public Boolean focus;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("type")
    public BuilderQueryEditorExpressionType type;
    public BuilderQueryEditorGroupByExpression() {
    }
    public BuilderQueryEditorGroupByExpression(BuilderQueryEditorProperty property,BuilderQueryEditorProperty interval,Boolean focus,BuilderQueryEditorExpressionType type) {
        this.property = property;
        this.interval = interval;
        this.focus = focus;
        this.type = type;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
