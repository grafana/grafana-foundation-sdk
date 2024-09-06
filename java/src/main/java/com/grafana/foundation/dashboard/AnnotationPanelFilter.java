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
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<AnnotationPanelFilter> {
        private final AnnotationPanelFilter internal;
        
        public Builder() {
            this.internal = new AnnotationPanelFilter();
        this.exclude(false);
        }
    public Builder exclude(Boolean exclude) {
    this.internal.exclude = exclude;
        return this;
    }
    
    public Builder ids(List<Integer> ids) {
    this.internal.ids = ids;
        return this;
    }
    public AnnotationPanelFilter build() {
            return this.internal;
        }
    }
}
