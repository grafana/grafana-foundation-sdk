// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboard from '../dashboard';

// A variable is a placeholder for a value. You can use variables in metric queries and in panel titles.
export class DatasourceVariableBuilder implements cog.Builder<dashboard.VariableModel> {
    protected readonly internal: dashboard.VariableModel;

    constructor(name: string) {
        this.internal = dashboard.defaultVariableModel();
        this.internal.name = name;
        this.internal.type = dashboard.VariableType.Datasource;
    }

    /**
     * Builds the object.
     */
    build(): dashboard.VariableModel {
        return this.internal;
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
    type(query: string | Record<string, any>): this {
        this.internal.query = query;
        return this;
    }

    // Shows current selected variable text/value on the dashboard
    current(current: dashboard.VariableOption): this {
        this.internal.current = current;
        return this;
    }

    // Whether multiple values can be selected or not from variable value list
    multi(multi: boolean): this {
        this.internal.multi = multi;
        return this;
    }

    // Whether all value option is available or not
    includeAll(includeAll: boolean): this {
        this.internal.includeAll = includeAll;
        return this;
    }

    // Custom all value
    allValue(allValue: string): this {
        this.internal.allValue = allValue;
        return this;
    }

    // Optional field, if you want to extract part of a series name or metric node segment.
    // Named capture groups can be used to separate the display text and value.
    regex(regex: string): this {
        this.internal.regex = regex;
        return this;
    }
}
