"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.ContactPointBuilder = void 0;
const tslib_1 = require("tslib");
const alerting = tslib_1.__importStar(require("../alerting"));
// EmbeddedContactPoint is the contact point type that is used
// by grafanas embedded alertmanager implementation.
class ContactPointBuilder {
    constructor() {
        this.internal = alerting.defaultContactPoint();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    disableResolveMessage(disableResolveMessage) {
        this.internal.disableResolveMessage = disableResolveMessage;
        return this;
    }
    // Name is used as grouping key in the UI. Contact points with the
    // same name will be grouped in the UI.
    name(name) {
        this.internal.name = name;
        return this;
    }
    provenance(provenance) {
        this.internal.provenance = provenance;
        return this;
    }
    settings(settings) {
        this.internal.settings = settings;
        return this;
    }
    type(type) {
        this.internal.type = type;
        return this;
    }
    // UID is the unique identifier of the contact point. The UID can be
    // set by the user.
    uid(uid) {
        if (!(uid.length >= 1)) {
            throw new Error("uid.length must be >= 1");
        }
        if (!(uid.length <= 40)) {
            throw new Error("uid.length must be <= 40");
        }
        this.internal.uid = uid;
        return this;
    }
}
exports.ContactPointBuilder = ContactPointBuilder;
//# sourceMappingURL=contactPointBuilder.gen.js.map