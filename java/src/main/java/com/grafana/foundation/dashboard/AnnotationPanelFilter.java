// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboard;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonSetter;
import com.fasterxml.jackson.annotation.Nulls;
import java.util.List;

public class AnnotationPanelFilter {
    // Should the specified panels be included or excluded
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("exclude")
    public Boolean exclude;
    // Panel IDs that should be included or excluded
    @JsonSetter(nulls = Nulls.AS_EMPTY)
    @JsonProperty("ids")
    public List<Integer> ids;
    public AnnotationPanelFilter() {
        this.exclude = false;
    }
    public AnnotationPanelFilter(Boolean exclude,List<Integer> ids) {
        this.exclude = exclude;
        this.ids = ids;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
