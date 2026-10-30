// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as alertingv1beta1 from '../alertingv1beta1';

export class RuleSpecBuilder implements cog.Builder<alertingv1beta1.RuleSpec> {
    protected readonly internal: alertingv1beta1.RuleSpec;

    constructor() {
        this.internal = alertingv1beta1.defaultRuleSpec();
    }

    /**
     * Builds the object.
     */
    build(): alertingv1beta1.RuleSpec {
        return this.internal;
    }

    title(title: string): this {
        this.internal.title = title;
        return this;
    }

    paused(paused: boolean): this {
        this.internal.paused = paused;
        return this;
    }

    trigger(trigger: cog.Builder<alertingv1beta1.IntervalTrigger>): this {
        const triggerResource = trigger.build();
        this.internal.trigger = triggerResource;
        return this;
    }

    labels(labels: Record<string, alertingv1beta1.TemplateString>): this {
        this.internal.labels = labels;
        return this;
    }

    expressions(expressions: alertingv1beta1.ExpressionMap): this {
        this.internal.expressions = expressions;
        return this;
    }
}

