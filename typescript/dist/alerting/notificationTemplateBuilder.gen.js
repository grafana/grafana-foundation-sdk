"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.NotificationTemplateBuilder = void 0;
const tslib_1 = require("tslib");
const alerting = tslib_1.__importStar(require("../alerting"));
class NotificationTemplateBuilder {
    constructor() {
        this.internal = alerting.defaultNotificationTemplate();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    name(name) {
        this.internal.name = name;
        return this;
    }
    provenance(provenance) {
        this.internal.provenance = provenance;
        return this;
    }
    template(template) {
        this.internal.template = template;
        return this;
    }
    version(version) {
        this.internal.version = version;
        return this;
    }
}
exports.NotificationTemplateBuilder = NotificationTemplateBuilder;
//# sourceMappingURL=notificationTemplateBuilder.gen.js.map