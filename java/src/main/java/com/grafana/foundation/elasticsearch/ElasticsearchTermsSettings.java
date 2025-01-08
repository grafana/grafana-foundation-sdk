// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class ElasticsearchTermsSettings {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("order")
    public TermsOrder order;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("size")
    public String size;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("min_doc_count")
    public String minDocCount;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("orderBy")
    public String orderBy;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("missing")
    public String missing;
    public ElasticsearchTermsSettings() {
    }
    
    public ElasticsearchTermsSettings(TermsOrder order,String size,String minDocCount,String orderBy,String missing) {
        this.order = order;
        this.size = size;
        this.minDocCount = minDocCount;
        this.orderBy = orderBy;
        this.missing = missing;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
