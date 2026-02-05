// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';

// Interval variable kind
export class IntervalVariableBuilder implements cog.Builder<dashboardv2beta1.IntervalVariableKind> {
    protected readonly internal: dashboardv2beta1.IntervalVariableKind;

    constructor(name: string) {
        this.internal = dashboardv2beta1.defaultIntervalVariableKind();
        this.internal.kind = "IntervalVariable";
        this.internal.spec.name = name;
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.IntervalVariableKind {
        return this.internal;
    }

    spec(spec: dashboardv2beta1.IntervalVariableSpec): this {
        this.internal.spec = spec;
        return this;
    }

    name(name: string): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultIntervalVariableSpec();
        }
        this.internal.spec.name = name;
        return this;
    }

    query(query: string): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultIntervalVariableSpec();
        }
        this.internal.spec.query = query;
        return this;
    }

    current(current: dashboardv2beta1.VariableOption): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultIntervalVariableSpec();
        }
        this.internal.spec.current = current;
        return this;
    }

    options(options: dashboardv2beta1.VariableOption[]): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultIntervalVariableSpec();
        }
        this.internal.spec.options = options;
        return this;
    }

    auto(auto: boolean): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultIntervalVariableSpec();
        }
        this.internal.spec.auto = auto;
        return this;
    }

    autoMin(autoMin: string): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultIntervalVariableSpec();
        }
        this.internal.spec.auto_min = autoMin;
        return this;
    }

    autoCount(autoCount: number): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultIntervalVariableSpec();
        }
        this.internal.spec.auto_count = autoCount;
        return this;
    }

    refresh(refresh: dashboardv2beta1.VariableRefresh): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultIntervalVariableSpec();
        }
        this.internal.spec.refresh = refresh;
        return this;
    }

    label(label: string): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultIntervalVariableSpec();
        }
        this.internal.spec.label = label;
        return this;
    }

    hide(hide: dashboardv2beta1.VariableHide): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultIntervalVariableSpec();
        }
        this.internal.spec.hide = hide;
        return this;
    }

    skipUrlSync(skipUrlSync: boolean): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultIntervalVariableSpec();
        }
        this.internal.spec.skipUrlSync = skipUrlSync;
        return this;
    }

    description(description: string): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultIntervalVariableSpec();
        }
        this.internal.spec.description = description;
        return this;
    }
}

