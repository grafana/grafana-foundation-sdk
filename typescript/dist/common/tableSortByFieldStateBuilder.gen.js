"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.TableSortByFieldStateBuilder = void 0;
const tslib_1 = require("tslib");
const common = tslib_1.__importStar(require("../common"));
// Sort by field state
class TableSortByFieldStateBuilder {
    constructor() {
        this.internal = common.defaultTableSortByFieldState();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    // Sets the display name of the field to sort by
    displayName(displayName) {
        this.internal.displayName = displayName;
        return this;
    }
    // Flag used to indicate descending sort order
    desc(desc) {
        this.internal.desc = desc;
        return this;
    }
}
exports.TableSortByFieldStateBuilder = TableSortByFieldStateBuilder;
//# sourceMappingURL=tableSortByFieldStateBuilder.gen.js.map