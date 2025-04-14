// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class ElasticsearchMovingAverageHoltWintersModelSettingsSettings {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("alpha")
    public String alpha;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("beta")
    public String beta;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("gamma")
    public String gamma;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("period")
    public String period;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("pad")
    public Boolean pad;
    public ElasticsearchMovingAverageHoltWintersModelSettingsSettings() {
    }
    public ElasticsearchMovingAverageHoltWintersModelSettingsSettings(String alpha,String beta,String gamma,String period,Boolean pad) {
        this.alpha = alpha;
        this.beta = beta;
        this.gamma = gamma;
        this.period = period;
        this.pad = pad;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
