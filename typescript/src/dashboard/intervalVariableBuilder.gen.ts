// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboard from '../dashboard';

// A variable is a placeholder for a value. You can use variables in metric queries and in panel titles.
export class IntervalVariableBuilder implements cog.Builder<dashboard.VariableModel> {
    protected readonly internal: dashboard.VariableModel;

    constructor(name: string) {
        this.internal = dashboard.defaultVariableModel();
        this.internal.name = name;
        this.internal.type = dashboard.VariableType.Interval;
    }

    build(): dashboard.VariableModel {
        return this.internal;
    }

    // Unique numeric identifier for the variable.
    id(id: string): this {
        this.internal.id = id;
        return this;
    }

    // Name of variable
    name(name: string): this {
        this.internal.name = name;
        return this;
    }

    // Optional display name
    label(label: string): this {
        this.internal.label = label;
        return this;
    }

    // Visibility configuration for the variable
    hide(hide: dashboard.VariableHide): this {
        this.internal.hide = hide;
        return this;
    }

    // Description of variable. It can be defined but `null`.
    description(description: string): this {
        this.internal.description = description;
        return this;
    }

    // Query used to fetch values for a variable
    values(query: string | Record<string, any>): this {
        this.internal.query = query;
        return this;
    }

    // Format to use while fetching all values from data source, eg: wildcard, glob, regex, pipe, etc.
    allFormat(allFormat: string): this {
        this.internal.allFormat = allFormat;
        return this;
    }

    // Shows current selected variable text/value on the dashboard
    current(current: dashboard.VariableOption): this {
        this.internal.current = current;
        return this;
    }

    // Options that can be selected for a variable.
    options(options: dashboard.VariableOption[]): this {
        this.internal.options = options;
        return this;
    }
}
