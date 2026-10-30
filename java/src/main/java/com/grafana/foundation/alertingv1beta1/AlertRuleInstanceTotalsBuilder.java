// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.alertingv1beta1;


/**
 * Count of alert instances per state. error also counts instances whose evaluation
 * errored but were mapped to another state via execErrState, so it can overlap.
 */
public class AlertRuleInstanceTotalsBuilder implements com.grafana.foundation.cog.Builder<AlertRuleInstanceTotals> {
    protected final AlertRuleInstanceTotals internal;
    
    public AlertRuleInstanceTotalsBuilder() {
        this.internal = new AlertRuleInstanceTotals();
    }
    public AlertRuleInstanceTotalsBuilder healthy(Long healthy) {
        this.internal.healthy = healthy;
        return this;
    }
    
    public AlertRuleInstanceTotalsBuilder firing(Long firing) {
        this.internal.firing = firing;
        return this;
    }
    
    public AlertRuleInstanceTotalsBuilder pending(Long pending) {
        this.internal.pending = pending;
        return this;
    }
    
    public AlertRuleInstanceTotalsBuilder recovering(Long recovering) {
        this.internal.recovering = recovering;
        return this;
    }
    
    public AlertRuleInstanceTotalsBuilder nodata(Long nodata) {
        this.internal.nodata = nodata;
        return this;
    }
    
    public AlertRuleInstanceTotalsBuilder error(Long error) {
        this.internal.error = error;
        return this;
    }
    public AlertRuleInstanceTotals build() {
        return this.internal;
    }
}
