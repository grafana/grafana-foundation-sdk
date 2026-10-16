// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

/**
 * Dashboard v2 shape: no `id`, it is carried by the parent's `group`.
 */
public class V2TransformationSpec {
    // Disabled transformations are skipped
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("disabled")
    public Boolean disabled;
    // Optional frame matcher. When missing it will be applied to all results
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("filter")
    public MatcherConfig filter;
    // Where to pull DataFrames from as input to transformation
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("topic")
    public DataTopic topic;
    // Options to be passed to the transformer
    // Valid options depend on the transformer id
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("options")
    public Object options;
    public V2TransformationSpec() {
        this.options = new Object();
    }
    public V2TransformationSpec(Boolean disabled,MatcherConfig filter,DataTopic topic,Object options) {
        this.disabled = disabled;
        this.filter = filter;
        this.topic = topic;
        this.options = options;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
