// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as alertingv1beta1 from '../alertingv1beta1';

export class RecordingRuleStatusBuilder implements cog.Builder<alertingv1beta1.RecordingRuleStatus> {
    protected readonly internal: alertingv1beta1.RecordingRuleStatus;

    constructor() {
        this.internal = alertingv1beta1.defaultRecordingRuleStatus();
    }

    /**
     * Builds the object.
     */
    build(): alertingv1beta1.RecordingRuleStatus {
        return this.internal;
    }

    health(health: alertingv1beta1.RecordingRuleHealth): this {
        this.internal.health = health;
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
}

