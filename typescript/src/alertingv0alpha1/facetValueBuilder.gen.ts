// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as alertingv0alpha1 from '../alertingv0alpha1';

/**
 * #FacetValue is a single value/count pair in a facet breakdown.
 */
export class FacetValueBuilder implements cog.Builder<alertingv0alpha1.FacetValue> {
    protected readonly internal: alertingv0alpha1.FacetValue;

    constructor() {
        this.internal = alertingv0alpha1.defaultFacetValue();
    }

    /**
     * Builds the object.
     */
    build(): alertingv0alpha1.FacetValue {
        return this.internal;
    }

    value(value: string): this {
        this.internal.value = value;
        return this;
    }

    count(count: number): this {
        this.internal.count = count;
        return this;
    }
}

