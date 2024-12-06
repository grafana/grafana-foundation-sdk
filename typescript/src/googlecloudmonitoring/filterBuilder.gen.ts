// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as googlecloudmonitoring from '../googlecloudmonitoring';

// Query filter representation.
export class FilterBuilder implements cog.Builder<googlecloudmonitoring.Filter> {
    protected readonly internal: googlecloudmonitoring.Filter;

    constructor() {
        this.internal = googlecloudmonitoring.defaultFilter();
    }

    /**
     * Builds the object.
     */
    build(): googlecloudmonitoring.Filter {
        return this.internal;
    }

    // Filter key.
    key(key: string): this {
        this.internal.key = key;
        return this;
    }

    // Filter operator.
    operator(operator: string): this {
        this.internal.operator = operator;
        return this;
    }

    // Filter value.
    value(value: string): this {
        this.internal.value = value;
        return this;
    }

    // Filter condition.
    condition(condition: string): this {
        this.internal.condition = condition;
        return this;
    }
}
