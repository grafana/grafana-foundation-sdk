"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.QueryVariableBuilder = void 0;
const tslib_1 = require("tslib");
const dashboard = tslib_1.__importStar(require("../dashboard"));
// A variable is a placeholder for a value. You can use variables in metric queries and in panel titles.
class QueryVariableBuilder {
    constructor(name) {
        this.internal = dashboard.defaultVariableModel();
        this.internal.name = name;
        this.internal.type = dashboard.VariableType.Query;
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    // Name of variable
    name(name) {
        this.internal.name = name;
        return this;
    }
    // Optional display name
    label(label) {
        this.internal.label = label;
        return this;
    }
    // Visibility configuration for the variable
    hide(hide) {
        this.internal.hide = hide;
        return this;
    }
    // Description of variable. It can be defined but `null`.
    description(description) {
        this.internal.description = description;
        return this;
    }
    // Query used to fetch values for a variable
    query(query) {
        this.internal.query = query;
        return this;
    }
    // Data source used to fetch values for a variable. It can be defined but `null`.
    datasource(datasource) {
        this.internal.datasource = datasource;
        return this;
    }
    // Shows current selected variable text/value on the dashboard
    current(current) {
        this.internal.current = current;
        return this;
    }
    // Whether multiple values can be selected or not from variable value list
    multi(multi) {
        this.internal.multi = multi;
        return this;
    }
    // Allow custom values to be entered in the variable
    allowCustomValue(allowCustomValue) {
        this.internal.allowCustomValue = allowCustomValue;
        return this;
    }
    // Options that can be selected for a variable.
    options(options) {
        this.internal.options = options;
        return this;
    }
    // Options to config when to refresh a variable
    refresh(refresh) {
        this.internal.refresh = refresh;
        return this;
    }
    // Options sort order
    sort(sort) {
        this.internal.sort = sort;
        return this;
    }
    // Whether all value option is available or not
    includeAll(includeAll) {
        this.internal.includeAll = includeAll;
        return this;
    }
    // Custom all value
    allValue(allValue) {
        this.internal.allValue = allValue;
        return this;
    }
    // Optional field, if you want to extract part of a series name or metric node segment.
    // Named capture groups can be used to separate the display text and value.
    regex(regex) {
        this.internal.regex = regex;
        return this;
    }
}
exports.QueryVariableBuilder = QueryVariableBuilder;
//# sourceMappingURL=queryVariableBuilder.gen.js.map