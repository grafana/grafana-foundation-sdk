// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as common from '../common';

// TODO docs
export class OptionsWithTimezonesBuilder implements cog.Builder<common.OptionsWithTimezones> {
    protected readonly internal: common.OptionsWithTimezones;

    constructor() {
        this.internal = common.defaultOptionsWithTimezones();
    }

    /**
     * Builds the object.
     */
    build(): common.OptionsWithTimezones {
        return this.internal;
    }

    timezone(timezone: common.TimeZone[]): this {
        this.internal.timezone = timezone;
        return this;
    }
}
