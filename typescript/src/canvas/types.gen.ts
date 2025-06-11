// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as common from '../common';


export enum HorizontalConstraint {
	Left = "left",
	Right = "right",
	LeftRight = "leftright",
	Center = "center",
	Scale = "scale",
}

export const defaultHorizontalConstraint = (): HorizontalConstraint => (HorizontalConstraint.Left);

export enum VerticalConstraint {
	Top = "top",
	Bottom = "bottom",
	TopBottom = "topbottom",
	Center = "center",
	Scale = "scale",
}

export const defaultVerticalConstraint = (): VerticalConstraint => (VerticalConstraint.Top);

export interface Constraint {
	horizontal?: HorizontalConstraint;
	vertical?: VerticalConstraint;
}

export const defaultConstraint = (): Constraint => ({
});

export interface Placement {
	top?: number;
	left?: number;
	right?: number;
	bottom?: number;
	width?: number;
	height?: number;
}

export const defaultPlacement = (): Placement => ({
});

export enum BackgroundImageSize {
	Original = "original",
	Contain = "contain",
	Cover = "cover",
	Fill = "fill",
	Tile = "tile",
}

export const defaultBackgroundImageSize = (): BackgroundImageSize => (BackgroundImageSize.Original);

export interface BackgroundConfig {
	color?: common.ColorDimensionConfig;
	image?: common.ResourceDimensionConfig;
	size?: BackgroundImageSize;
}

export const defaultBackgroundConfig = (): BackgroundConfig => ({
});

export interface LineConfig {
	color?: common.ColorDimensionConfig;
	width?: number;
}

export const defaultLineConfig = (): LineConfig => ({
});

export interface ConnectionCoordinates {
	x: number;
	y: number;
}

export const defaultConnectionCoordinates = (): ConnectionCoordinates => ({
	x: 0,
	y: 0,
});

export enum ConnectionPath {
	Straight = "straight",
}

export const defaultConnectionPath = (): ConnectionPath => (ConnectionPath.Straight);

export interface CanvasConnection {
	source: ConnectionCoordinates;
	target: ConnectionCoordinates;
	targetName?: string;
	path: ConnectionPath.Straight;
	color?: common.ColorDimensionConfig;
	size?: common.ScaleDimensionConfig;
}

export const defaultCanvasConnection = (): CanvasConnection => ({
	source: defaultConnectionCoordinates(),
	target: defaultConnectionCoordinates(),
	path: ConnectionPath.Straight,
});

export interface CanvasElementOptions {
	name: string;
	type: string;
	// TODO: figure out how to define this (element config(s))
	config?: any;
	constraint?: Constraint;
	placement?: Placement;
	background?: BackgroundConfig;
	border?: LineConfig;
	connections?: CanvasConnection[];
}

export const defaultCanvasElementOptions = (): CanvasElementOptions => ({
	name: "",
	type: "",
});

export interface Options {
	// Enable inline editing
	inlineEditing: boolean;
	// Show all available element types
	showAdvancedTypes: boolean;
	// The root element of canvas (frame), where all canvas elements are nested
	// TODO: Figure out how to define a default value for this
	root: {
		// Name of the root element
		name: string;
		// Type of root element (frame)
		type: "frame";
		// The list of canvas elements attached to the root element
		elements: CanvasElementOptions[];
	};
}

export const defaultOptions = (): Options => ({
	inlineEditing: true,
	showAdvancedTypes: true,
	root: {
	name: "",
	type: "frame",
	elements: [],
},
});

