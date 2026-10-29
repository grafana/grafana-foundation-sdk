// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.alertingv0alpha1;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;
import java.util.Map;
import java.util.List;

/**
 * #RuleSearchHitFields is the per-kind field payload returned on each hit.
 * It carries the union of alert- and recording-rule search fields; only the
 * fields relevant to a hit's kind are populated. This maps to the kind's
 * declared searchFields.
 */
public class RuleSearchHitFields {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("title")
    public String title;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("folder")
    public String folder;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("type")
    public String type;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("interval")
    public String interval;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("paused")
    public Boolean paused;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("labels")
    public Map<String, String> labels;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("datasourceUIDs")
    public List<String> datasourceUIDs;
    // Alert-rule fields.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("annotations")
    public Map<String, String> annotations;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("for")
    public String forArg;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("keepFiringFor")
    public String keepFiringFor;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("dashboardUID")
    public String dashboardUID;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("panelID")
    public Long panelID;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("receiver")
    public String receiver;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("notificationType")
    public String notificationType;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("routingTree")
    public String routingTree;
    // Recording-rule fields.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("metric")
    public String metric;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("targetDatasourceUID")
    public String targetDatasourceUID;
    public RuleSearchHitFields() {
    }
    public RuleSearchHitFields(String title,String folder,String type,String interval,Boolean paused,Map<String, String> labels,List<String> datasourceUIDs,Map<String, String> annotations,String forArg,String keepFiringFor,String dashboardUID,Long panelID,String receiver,String notificationType,String routingTree,String metric,String targetDatasourceUID) {
        this.title = title;
        this.folder = folder;
        this.type = type;
        this.interval = interval;
        this.paused = paused;
        this.labels = labels;
        this.datasourceUIDs = datasourceUIDs;
        this.annotations = annotations;
        this.forArg = forArg;
        this.keepFiringFor = keepFiringFor;
        this.dashboardUID = dashboardUID;
        this.panelID = panelID;
        this.receiver = receiver;
        this.notificationType = notificationType;
        this.routingTree = routingTree;
        this.metric = metric;
        this.targetDatasourceUID = targetDatasourceUID;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
