// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;
import java.util.List;

public class BucketScript {
    @JsonProperty("type")
    public MetricAggregationType type;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("pipelineVariables")
    public List<PipelineVariable> pipelineVariables;
    @JsonProperty("id")
    public String id;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("settings")
    public ElasticsearchBucketScriptSettings settings;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("hide")
    public Boolean hide;
    public BucketScript() {
        this.type = MetricAggregationType.BUCKET_SCRIPT;
    }
    public BucketScript(List<PipelineVariable> pipelineVariables,String id,ElasticsearchBucketScriptSettings settings,Boolean hide) {
        this.type = MetricAggregationType.BUCKET_SCRIPT;
        this.pipelineVariables = pipelineVariables;
        this.id = id;
        this.settings = settings;
        this.hide = hide;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
