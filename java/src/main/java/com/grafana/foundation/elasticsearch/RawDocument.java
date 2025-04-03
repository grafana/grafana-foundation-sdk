// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class RawDocument {
    @JsonProperty("type")
    public MetricAggregationType type;
    @JsonProperty("id")
    public String id;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("settings")
    public ElasticsearchRawDocumentSettings settings;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("hide")
    public Boolean hide;
    public RawDocument() {
        this.type = MetricAggregationType.RAW_DOCUMENT;
    }
    public RawDocument(String id,ElasticsearchRawDocumentSettings settings,Boolean hide) {
        this.type = MetricAggregationType.RAW_DOCUMENT;
        this.id = id;
        this.settings = settings;
        this.hide = hide;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
