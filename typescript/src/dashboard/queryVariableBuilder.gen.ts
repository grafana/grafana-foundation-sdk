// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboard from '../dashboard';

// A variable is a placeholder for a value. You can use variables in metric queries and in panel titles.
export class QueryVariableBuilder implements cog.Builder<dashboard.VariableModel> {
    protected readonly internal: dashboard.VariableModel;

    constructor(name: string) {
        this.internal = dashboard.defaultVariableModel();
        this.internal.name = name;
        this.internal.type = dashboard.VariableType.Query;
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
    query(query: string | Record<string, any>): this {
        this.internal.query = query;
        return this;
    }

    // Data source used to fetch values for a variable. It can be defined but `null`.
    datasource(datasource: dashboard.DataSourceRef): this {
        this.internal.datasource = datasource;
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

    // Options that can be selected for a variable.
    options(options: dashboard.VariableOption[]): this {
        this.internal.options = options;
        return this;
    }

    refresh(refresh: dashboard.VariableRefresh): this {
        this.internal.refresh = refresh;
        return this;
    }

    // Options sort order
    sort(sort: dashboard.VariableSort): this {
        this.internal.sort = sort;
        return this;
    }
}
