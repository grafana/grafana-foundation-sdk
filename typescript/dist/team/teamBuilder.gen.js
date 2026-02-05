"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.TeamBuilder = void 0;
const tslib_1 = require("tslib");
const team = tslib_1.__importStar(require("../team"));
class TeamBuilder {
    constructor(name) {
        this.internal = team.defaultTeam();
        this.internal.name = name;
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    email(email) {
        this.internal.email = email;
        return this;
    }
    name(name) {
        this.internal.name = name;
        return this;
    }
}
exports.TeamBuilder = TeamBuilder;
//# sourceMappingURL=teamBuilder.gen.js.map