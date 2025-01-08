// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.xychart;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

// NOTE: (copied from dashboard_kind.cue, since not exported)
// Matcher is a predicate configuration. Based on the config a set of field(s) or values is filtered in order to apply override / transformation.
// It comes with in id ( to resolve implementation from registry) and a configuration that’s specific to a particular matcher type.
public class MatcherConfig {
    // The matcher id. This is used to find the matcher implementation from registry.
    @JsonProperty("id")
    public String id;
    // The matcher options. This is specific to the matcher implementation.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("options")
    public Object options;
    public MatcherConfig() {
        this.id = "";
    }
    
    public MatcherConfig(String id,Object options) {
        this.id = id;
        this.options = options;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
