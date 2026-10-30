// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as alertingv1beta1 from '../alertingv1beta1';

/**
 * Count of alert instances per state. error also counts instances whose evaluation
 * errored but were mapped to another state via execErrState, so it can overlap.
 */
export class AlertRuleInstanceTotalsBuilder implements cog.Builder<alertingv1beta1.AlertRuleInstanceTotals> {
    protected readonly internal: alertingv1beta1.AlertRuleInstanceTotals;

    constructor() {
        this.internal = alertingv1beta1.defaultAlertRuleInstanceTotals();
    }

    /**
     * Builds the object.
     */
    build(): alertingv1beta1.AlertRuleInstanceTotals {
        return this.internal;
    }

    healthy(healthy: number): this {
        this.internal.healthy = healthy;
        return this;
    }

    firing(firing: number): this {
        this.internal.firing = firing;
        return this;
    }

    pending(pending: number): this {
        this.internal.pending = pending;
        return this;
    }

    recovering(recovering: number): this {
        this.internal.recovering = recovering;
        return this;
    }

    nodata(nodata: number): this {
        this.internal.nodata = nodata;
        return this;
    }

    error(error: number): this {
        this.internal.error = error;
        return this;
    }
}

