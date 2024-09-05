// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboard;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
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
    @JsonProperty("tags")
    public List<String> tags;
    // Only required/valid for the grafana datasource...
    // but code+tests is already depending on it so hard to change
    @JsonProperty("type")
    public String type;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<AnnotationTarget> {
        private final AnnotationTarget internal;
        
        public Builder() {
            this.internal = new AnnotationTarget();
        }
    public Builder limit(Long limit) {
    this.internal.limit = limit;
        return this;
    }
    
    public Builder matchAny(Boolean matchAny) {
    this.internal.matchAny = matchAny;
        return this;
    }
    
    public Builder tags(List<String> tags) {
    this.internal.tags = tags;
        return this;
    }
    
    public Builder type(String type) {
    this.internal.type = type;
        return this;
    }
    public AnnotationTarget build() {
            return this.internal;
        }
    }
}
