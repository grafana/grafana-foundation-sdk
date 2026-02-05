"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.defaultPanelModel = exports.defaultLibraryElementDTOMetaUser = exports.defaultLibraryElementDTOMeta = exports.defaultLibraryPanel = void 0;
const defaultLibraryPanel = () => ({
    uid: "",
    name: "",
    type: "",
    version: 0,
    model: (0, exports.defaultPanelModel)(),
});
exports.defaultLibraryPanel = defaultLibraryPanel;
const defaultLibraryElementDTOMeta = () => ({
    folderName: "",
    folderUid: "",
    connectedDashboards: 0,
    created: "",
    updated: "",
    createdBy: (0, exports.defaultLibraryElementDTOMetaUser)(),
    updatedBy: (0, exports.defaultLibraryElementDTOMetaUser)(),
});
exports.defaultLibraryElementDTOMeta = defaultLibraryElementDTOMeta;
const defaultLibraryElementDTOMetaUser = () => ({
    id: 0,
    name: "",
    avatarUrl: "",
});
exports.defaultLibraryElementDTOMetaUser = defaultLibraryElementDTOMetaUser;
const defaultPanelModel = () => ({
    type: "",
    transparent: false,
    repeatDirection: "h",
});
exports.defaultPanelModel = defaultPanelModel;
//# sourceMappingURL=types.gen.js.map