// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as alerting from '../alerting';

export class MatcherBuilder implements cog.Builder<alerting.Matcher> {
    protected readonly internal: alerting.Matcher;

    constructor() {
        this.internal = alerting.defaultMatcher();
    }

    /**
     * Builds the object.
     */
    build(): alerting.Matcher {
        return this.internal;
    }

    name(name: string): this {
        this.internal.Name = name;
        return this;
    }

    type(type: alerting.MatchType): this {
        this.internal.Type = type;
        return this;
    }

    value(value: string): this {
        this.internal.Value = value;
        return this;
    }
}
