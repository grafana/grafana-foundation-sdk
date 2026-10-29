// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as alertingv0alpha1 from '../alertingv0alpha1';

export class RuleRefBuilder implements cog.Builder<alertingv0alpha1.RuleRef> {
    protected readonly internal: alertingv0alpha1.RuleRef;

    constructor() {
        this.internal = alertingv0alpha1.defaultRuleRef();
    }

    /**
     * Builds the object.
     */
    build(): alertingv0alpha1.RuleRef {
        return this.internal;
    }

    // name is the metadata.name of an AlertRule or RecordingRule resource.
    name(name: alertingv0alpha1.RuleUID): this {
        this.internal.name = name;
        return this;
    }
}

