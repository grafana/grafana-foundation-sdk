// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class ElasticsearchHistogramSettings {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("interval")
    public String interval;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("min_doc_count")
    public String minDocCount;
    public ElasticsearchHistogramSettings() {
    }
    public ElasticsearchHistogramSettings(String interval,String minDocCount) {
        this.interval = interval;
        this.minDocCount = minDocCount;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
