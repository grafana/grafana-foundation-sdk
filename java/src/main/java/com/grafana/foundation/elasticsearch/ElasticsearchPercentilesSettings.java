// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;
import java.util.List;

public class ElasticsearchPercentilesSettings {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("script")
    public InlineScript script;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("missing")
    public String missing;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("percents")
    public List<String> percents;
    public ElasticsearchPercentilesSettings() {
    }
    public ElasticsearchPercentilesSettings(InlineScript script,String missing,List<String> percents) {
        this.script = script;
        this.missing = missing;
        this.percents = percents;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
