// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as alertingv1beta1 from '../alertingv1beta1';

export class RuleRefBuilder implements cog.Builder<alertingv1beta1.RuleRef> {
    protected readonly internal: alertingv1beta1.RuleRef;

    constructor() {
        this.internal = alertingv1beta1.defaultRuleRef();
    }

    /**
     * Builds the object.
     */
    build(): alertingv1beta1.RuleRef {
        return this.internal;
    }

    // name is the metadata.name of an AlertRule or RecordingRule resource.
    name(name: alertingv1beta1.RuleUID): this {
        this.internal.name = name;
        return this;
    }
}

