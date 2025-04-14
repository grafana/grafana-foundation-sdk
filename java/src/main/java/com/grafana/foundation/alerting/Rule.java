// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.alerting;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;
import java.util.Map;
import com.fasterxml.jackson.annotation.JsonSetter;
import com.fasterxml.jackson.annotation.Nulls;
import java.util.List;

public class Rule {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("annotations")
    public Map<String, String> annotations;
    @JsonProperty("condition")
    public String condition;
    @JsonSetter(nulls = Nulls.AS_EMPTY)
    @JsonProperty("data")
    public List<Query> data;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("execErrState")
    public RuleExecErrState execErrState;
    @JsonProperty("folderUID")
    public String folderUID;
    // The amount of time, in seconds, for which the rule must be breached for the rule to be considered to be Firing.
    // Before this time has elapsed, the rule is only considered to be Pending.
    @JsonProperty("for")
    public String forArg;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("id")
    public Long id;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("isPaused")
    public Boolean isPaused;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("labels")
    public Map<String, String> labels;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("noDataState")
    public RuleNoDataState noDataState;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("notification_settings")
    public NotificationSettings notificationSettings;
    @JsonProperty("orgID")
    public Long orgID;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("provenance")
    public String provenance;
    @JsonProperty("ruleGroup")
    public String ruleGroup;
    @JsonProperty("title")
    public String title;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("uid")
    public String uid;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("updated")
    public String updated;
    public Rule() {
    }
    public Rule(Map<String, String> annotations,String condition,List<Query> data,RuleExecErrState execErrState,String folderUID,String forArg,Long id,Boolean isPaused,Map<String, String> labels,RuleNoDataState noDataState,NotificationSettings notificationSettings,Long orgID,String provenance,String ruleGroup,String title,String uid,String updated) {
        this.annotations = annotations;
        this.condition = condition;
        this.data = data;
        this.execErrState = execErrState;
        this.folderUID = folderUID;
        this.forArg = forArg;
        this.id = id;
        this.isPaused = isPaused;
        this.labels = labels;
        this.noDataState = noDataState;
        this.notificationSettings = notificationSettings;
        this.orgID = orgID;
        this.provenance = provenance;
        this.ruleGroup = ruleGroup;
        this.title = title;
        this.uid = uid;
        this.updated = updated;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
