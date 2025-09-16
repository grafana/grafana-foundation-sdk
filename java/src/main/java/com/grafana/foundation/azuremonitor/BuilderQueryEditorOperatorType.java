// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.azuremonitor;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.databind.annotation.JsonDeserialize;
import com.fasterxml.jackson.annotation.JsonInclude;

@JsonDeserialize(using = BuilderQueryEditorOperatorTypeDeserializer.class)
public class BuilderQueryEditorOperatorType {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("String")
    protected String string;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("Bool")
    protected Boolean bool;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("Float64")
    protected Double float64;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("SelectableValue")
    protected SelectableValue selectableValue;
    protected BuilderQueryEditorOperatorType() {}
    public static BuilderQueryEditorOperatorType createString(String string) {
        BuilderQueryEditorOperatorType builderQueryEditorOperatorType = new BuilderQueryEditorOperatorType();
        builderQueryEditorOperatorType.string = string;
        return builderQueryEditorOperatorType;
    }
    public static BuilderQueryEditorOperatorType createBool(Boolean bool) {
        BuilderQueryEditorOperatorType builderQueryEditorOperatorType = new BuilderQueryEditorOperatorType();
        builderQueryEditorOperatorType.bool = bool;
        return builderQueryEditorOperatorType;
    }
    public static BuilderQueryEditorOperatorType createFloat64(Double float64) {
        BuilderQueryEditorOperatorType builderQueryEditorOperatorType = new BuilderQueryEditorOperatorType();
        builderQueryEditorOperatorType.float64 = float64;
        return builderQueryEditorOperatorType;
    }
    public static BuilderQueryEditorOperatorType createSelectableValue(SelectableValue selectableValue) {
        BuilderQueryEditorOperatorType builderQueryEditorOperatorType = new BuilderQueryEditorOperatorType();
        builderQueryEditorOperatorType.selectableValue = selectableValue;
        return builderQueryEditorOperatorType;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
