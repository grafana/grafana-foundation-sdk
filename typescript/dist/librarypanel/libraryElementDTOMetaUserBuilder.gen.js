"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.LibraryElementDTOMetaUserBuilder = void 0;
const tslib_1 = require("tslib");
const librarypanel = tslib_1.__importStar(require("../librarypanel"));
class LibraryElementDTOMetaUserBuilder {
    constructor() {
        this.internal = librarypanel.defaultLibraryElementDTOMetaUser();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    id(id) {
        this.internal.id = id;
        return this;
    }
    name(name) {
        this.internal.name = name;
        return this;
    }
    avatarUrl(avatarUrl) {
        this.internal.avatarUrl = avatarUrl;
        return this;
    }
}
exports.LibraryElementDTOMetaUserBuilder = LibraryElementDTOMetaUserBuilder;
//# sourceMappingURL=libraryElementDTOMetaUserBuilder.gen.js.map