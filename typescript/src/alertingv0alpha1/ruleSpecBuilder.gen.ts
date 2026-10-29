// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as alertingv0alpha1 from '../alertingv0alpha1';

export class RuleSpecBuilder implements cog.Builder<alertingv0alpha1.RuleSpec> {
    protected readonly internal: alertingv0alpha1.RuleSpec;

    constructor() {
        this.internal = alertingv0alpha1.defaultRuleSpec();
    }

    /**
     * Builds the object.
     */
    build(): alertingv0alpha1.RuleSpec {
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

    trigger(trigger: cog.Builder<alertingv0alpha1.IntervalTrigger>): this {
        const triggerResource = trigger.build();
        this.internal.trigger = triggerResource;
        return this;
    }

    labels(labels: Record<string, alertingv0alpha1.TemplateString>): this {
        this.internal.labels = labels;
        return this;
    }

    expressions(expressions: alertingv0alpha1.ExpressionMap): this {
        this.internal.expressions = expressions;
        return this;
    }
}

