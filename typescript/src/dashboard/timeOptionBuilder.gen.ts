// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboard from '../dashboard';

// Counterpart for TypeScript's TimeOption type.
export class TimeOptionBuilder implements cog.Builder<dashboard.TimeOption> {
    protected readonly internal: dashboard.TimeOption;

    constructor() {
        this.internal = dashboard.defaultTimeOption();
    }

    /**
     * Builds the object.
     */
    build(): dashboard.TimeOption {
        return this.internal;
    }

    display(display: string): this {
        this.internal.display = display;
        return this;
    }

    from(from: string): this {
        this.internal.from = from;
        return this;
    }

    to(to: string): this {
        this.internal.to = to;
        return this;
    }
}

