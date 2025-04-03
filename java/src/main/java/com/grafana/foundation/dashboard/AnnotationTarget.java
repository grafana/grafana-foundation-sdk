// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboard;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonSetter;
import com.fasterxml.jackson.annotation.Nulls;
import java.util.List;

// TODO: this should be a regular DataQuery that depends on the selected dashboard
// these match the properties of the "grafana" datasouce that is default in most dashboards
public class AnnotationTarget {
    // Only required/valid for the grafana datasource...
    // but code+tests is already depending on it so hard to change
    @JsonProperty("limit")
    public Long limit;
    // Only required/valid for the grafana datasource...
    // but code+tests is already depending on it so hard to change
    @JsonProperty("matchAny")
    public Boolean matchAny;
    // Only required/valid for the grafana datasource...
    // but code+tests is already depending on it so hard to change
    @JsonSetter(nulls = Nulls.AS_EMPTY)
    @JsonProperty("tags")
    public List<String> tags;
    // Only required/valid for the grafana datasource...
    // but code+tests is already depending on it so hard to change
    @JsonProperty("type")
    public String type;
    public AnnotationTarget() {
    }
    public AnnotationTarget(Long limit,Boolean matchAny,List<String> tags,String type) {
        this.limit = limit;
        this.matchAny = matchAny;
        this.tags = tags;
        this.type = type;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
