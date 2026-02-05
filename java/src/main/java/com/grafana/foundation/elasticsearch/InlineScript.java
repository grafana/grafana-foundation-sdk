// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.databind.annotation.JsonDeserialize;
import com.fasterxml.jackson.annotation.JsonInclude;

@JsonDeserialize(using = InlineScriptDeserializer.class)
public class InlineScript {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("String")
    protected String string;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("ElasticsearchInlineScript")
    protected ElasticsearchInlineScript elasticsearchInlineScript;
    protected InlineScript() {}
    public static InlineScript createString(String string) {
        InlineScript inlineScript = new InlineScript();
        inlineScript.string = string;
        return inlineScript;
    }
    public static InlineScript createElasticsearchInlineScript(ElasticsearchInlineScript elasticsearchInlineScript) {
        InlineScript inlineScript = new InlineScript();
        inlineScript.elasticsearchInlineScript = elasticsearchInlineScript;
        return inlineScript;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
