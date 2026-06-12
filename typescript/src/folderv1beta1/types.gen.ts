// Code generated - EDITING IS FUTILE. DO NOT EDIT.

/**
 * @deprecated Prefer using folder.Folder instead.
 */
export interface Folder {
	title: string;
	description?: string;
}

export const defaultFolder = (): Folder => ({
	title: "",
});

export const FolderApiVersion = "folder.grafana.app/v1beta1";

export const FolderKind = "Folder";

