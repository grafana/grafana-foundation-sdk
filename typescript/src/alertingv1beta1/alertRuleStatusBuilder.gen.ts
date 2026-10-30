// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as alertingv1beta1 from '../alertingv1beta1';

export class AlertRuleStatusBuilder implements cog.Builder<alertingv1beta1.AlertRuleStatus> {
    protected readonly internal: alertingv1beta1.AlertRuleStatus;

    constructor() {
        this.internal = alertingv1beta1.defaultAlertRuleStatus();
    }

    /**
     * Builds the object.
     */
    build(): alertingv1beta1.AlertRuleStatus {
        return this.internal;
    }

    health(health: alertingv1beta1.AlertRuleHealth): this {
        this.internal.health = health;
        return this;
    }

    state(state: alertingv1beta1.AlertRuleState): this {
        this.internal.state = state;
        return this;
    }

    stateReason(stateReason: alertingv1beta1.AlertRuleStateReason): this {
        this.internal.stateReason = stateReason;
        return this;
    }

    lastEvaluationTime(lastEvaluationTime: string): this {
        this.internal.lastEvaluationTime = lastEvaluationTime;
        return this;
    }

    // duration of the last evaluation in seconds
    evaluationDuration(evaluationDuration: number): this {
        this.internal.evaluationDuration = evaluationDuration;
        return this;
    }

    lastError(lastError: string): this {
        this.internal.lastError = lastError;
        return this;
    }

    totals(totals: cog.Builder<alertingv1beta1.AlertRuleInstanceTotals>): this {
        const totalsResource = totals.build();
        this.internal.totals = totalsResource;
        return this;
    }
}

