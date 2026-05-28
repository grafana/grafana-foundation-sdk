// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2 from '../dashboardv2';

// Interval variable kind
export class IntervalVariableBuilder implements cog.Builder<dashboardv2.IntervalVariableKind> {
    protected readonly internal: dashboardv2.IntervalVariableKind;

    constructor(name: string) {
        this.internal = dashboardv2.defaultIntervalVariableKind();
        this.internal.kind = "IntervalVariable";
        this.internal.spec.refresh = "onTimeRangeChanged";
        this.internal.spec.name = name;
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2.IntervalVariableKind {
        return this.internal;
    }

    name(name: string): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultIntervalVariableSpec();
        }
        this.internal.spec.name = name;
        return this;
    }

    query(query: string): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultIntervalVariableSpec();
        }
        this.internal.spec.query = query;
        return this;
    }

    current(current: dashboardv2.VariableOption): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultIntervalVariableSpec();
        }
        this.internal.spec.current = current;
        return this;
    }

    options(options: dashboardv2.VariableOption[]): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultIntervalVariableSpec();
        }
        this.internal.spec.options = options;
        return this;
    }

    auto(auto: boolean): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultIntervalVariableSpec();
        }
        this.internal.spec.auto = auto;
        return this;
    }

    autoMin(autoMin: string): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultIntervalVariableSpec();
        }
        this.internal.spec.auto_min = autoMin;
        return this;
    }

    autoCount(autoCount: number): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultIntervalVariableSpec();
        }
        this.internal.spec.auto_count = autoCount;
        return this;
    }

    label(label: string): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultIntervalVariableSpec();
        }
        this.internal.spec.label = label;
        return this;
    }

    hide(hide: dashboardv2.VariableHide): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultIntervalVariableSpec();
        }
        this.internal.spec.hide = hide;
        return this;
    }

    skipUrlSync(skipUrlSync: boolean): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultIntervalVariableSpec();
        }
        this.internal.spec.skipUrlSync = skipUrlSync;
        return this;
    }

    description(description: string): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultIntervalVariableSpec();
        }
        this.internal.spec.description = description;
        return this;
    }

    origin(origin: cog.Builder<dashboardv2.ControlSourceRef>): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultIntervalVariableSpec();
        }
        const originResource = origin.build();
        this.internal.spec.origin = originResource;
        return this;
    }
}

