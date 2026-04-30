// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2;

import com.fasterxml.jackson.annotation.JsonUnwrapped;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.databind.annotation.JsonDeserialize;
import com.fasterxml.jackson.annotation.JsonInclude;

@JsonDeserialize(using = AutoGridLayoutKindOrGridLayoutKindDeserializer.class)
public class AutoGridLayoutKindOrGridLayoutKind {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonUnwrapped
    protected AutoGridLayoutKind autoGridLayoutKind;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonUnwrapped
    protected GridLayoutKind gridLayoutKind;
    protected AutoGridLayoutKindOrGridLayoutKind() {}
    public static AutoGridLayoutKindOrGridLayoutKind createAutoGridLayoutKind(AutoGridLayoutKind autoGridLayoutKind) {
        AutoGridLayoutKindOrGridLayoutKind autoGridLayoutKindOrGridLayoutKind = new AutoGridLayoutKindOrGridLayoutKind();
        autoGridLayoutKindOrGridLayoutKind.autoGridLayoutKind = autoGridLayoutKind;
        return autoGridLayoutKindOrGridLayoutKind;
    }
    public static AutoGridLayoutKindOrGridLayoutKind createGridLayoutKind(GridLayoutKind gridLayoutKind) {
        AutoGridLayoutKindOrGridLayoutKind autoGridLayoutKindOrGridLayoutKind = new AutoGridLayoutKindOrGridLayoutKind();
        autoGridLayoutKindOrGridLayoutKind.gridLayoutKind = gridLayoutKind;
        return autoGridLayoutKindOrGridLayoutKind;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
