// Code generated - EDITING IS FUTILE. DO NOT EDIT.

// TODO:
// common metadata will soon support setting the parent folder in the metadata
export interface Folder {
	// Unique folder id. (will be k8s name)
	uid: string;
	// Folder title
	title: string;
	// Description of the folder.
	description?: string;
}

export const defaultFolder = (): Folder => ({
	uid: "",
	title: "",
});

