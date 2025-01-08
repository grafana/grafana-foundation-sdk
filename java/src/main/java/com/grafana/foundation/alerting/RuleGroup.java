// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.alerting;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;
import java.util.List;

public class RuleGroup {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("folderUid")
    public String folderUid;
    // The interval, in seconds, at which all rules in the group are evaluated.
    // If a group contains many rules, the rules are evaluated sequentially.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("interval")
    public Long interval;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("rules")
    public List<Rule> rules;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("title")
    public String title;
    public RuleGroup() {
    }
    
    public RuleGroup(String folderUid,Long interval,List<Rule> rules,String title) {
        this.folderUid = folderUid;
        this.interval = interval;
        this.rules = rules;
        this.title = title;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
