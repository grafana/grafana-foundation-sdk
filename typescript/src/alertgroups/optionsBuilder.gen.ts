// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as alertgroups from '../alertgroups';

export class OptionsBuilder implements cog.Builder<alertgroups.Options> {
    protected readonly internal: alertgroups.Options;

    constructor() {
        this.internal = alertgroups.defaultOptions();
    }

    /**
     * Builds the object.
     */
    build(): alertgroups.Options {
        return this.internal;
    }

    // Comma-separated list of values used to filter alert results
    labels(labels: string): this {
        this.internal.labels = labels;
        return this;
    }

    // Name of the alertmanager used as a source for alerts
    alertmanager(alertmanager: string): this {
        this.internal.alertmanager = alertmanager;
        return this;
    }

    // Expand all alert groups by default
    expandAll(expandAll: boolean): this {
        this.internal.expandAll = expandAll;
        return this;
    }
}

