// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class InlineScript {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("String")
    public String string;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("ElasticsearchInlineScript")
    public ElasticsearchInlineScript elasticsearchInlineScript;
    public InlineScript() {
    }
    
    public InlineScript(String string,ElasticsearchInlineScript elasticsearchInlineScript) {
        this.string = string;
        this.elasticsearchInlineScript = elasticsearchInlineScript;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
