// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as tempo from '../tempo';

export class TraceqlFilterBuilder implements cog.Builder<tempo.TraceqlFilter> {
    protected readonly internal: tempo.TraceqlFilter;

    constructor() {
        this.internal = tempo.defaultTraceqlFilter();
    }

    /**
     * Builds the object.
     */
    build(): tempo.TraceqlFilter {
        return this.internal;
    }

    // Uniquely identify the filter, will not be used in the query generation
    id(id: string): this {
        this.internal.id = id;
        return this;
    }

    // The tag for the search filter, for example: .http.status_code, .service.name, status
    tag(tag: string): this {
        this.internal.tag = tag;
        return this;
    }

    // The operator that connects the tag to the value, for example: =, >, !=, =~
    operator(operator: string): this {
        this.internal.operator = operator;
        return this;
    }

    // The value for the search filter
    value(value: string | string[]): this {
        this.internal.value = value;
        return this;
    }

    // The type of the value, used for example to check whether we need to wrap the value in quotes when generating the query
    valueType(valueType: string): this {
        this.internal.valueType = valueType;
        return this;
    }

    // The scope of the filter, can either be unscoped/all scopes, resource or span
    scope(scope: tempo.TraceqlSearchScope): this {
        this.internal.scope = scope;
        return this;
    }
}
