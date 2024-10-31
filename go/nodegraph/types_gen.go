// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package nodegraph

import (
	"encoding/json"
	"fmt"
	"strconv"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	variants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
	dashboard "github.com/grafana/grafana-foundation-sdk/go/dashboard"
)

type ArcOption struct {
	// Field from which to get the value. Values should be less than 1, representing fraction of a circle.
	Field *string `json:"field,omitempty"`
	// The color of the arc.
	Color *string `json:"color,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ArcOption` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *ArcOption) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "field"
	if fields["field"] != nil {
		if string(fields["field"]) != "null" {
			if err := json.Unmarshal(fields["field"], &resource.Field); err != nil {
				errs = append(errs, cog.MakeBuildErrors("field", err)...)
			}

		}
		delete(fields, "field")

	}
	// Field "color"
	if fields["color"] != nil {
		if string(fields["color"]) != "null" {
			if err := json.Unmarshal(fields["color"], &resource.Color); err != nil {
				errs = append(errs, cog.MakeBuildErrors("color", err)...)
			}

		}
		delete(fields, "color")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("ArcOption", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `ArcOption` objects.
func (resource ArcOption) Equals(other ArcOption) bool {
	if resource.Field == nil && other.Field != nil || resource.Field != nil && other.Field == nil {
		return false
	}

	if resource.Field != nil {
		if *resource.Field != *other.Field {
			return false
		}
	}
	if resource.Color == nil && other.Color != nil || resource.Color != nil && other.Color == nil {
		return false
	}

	if resource.Color != nil {
		if *resource.Color != *other.Color {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `ArcOption` fields for violations and returns them.
func (resource ArcOption) Validate() error {
	return nil
}

type NodeOptions struct {
	// Unit for the main stat to override what ever is set in the data frame.
	MainStatUnit *string `json:"mainStatUnit,omitempty"`
	// Unit for the secondary stat to override what ever is set in the data frame.
	SecondaryStatUnit *string `json:"secondaryStatUnit,omitempty"`
	// Define which fields are shown as part of the node arc (colored circle around the node).
	Arcs []ArcOption `json:"arcs,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `NodeOptions` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *NodeOptions) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "mainStatUnit"
	if fields["mainStatUnit"] != nil {
		if string(fields["mainStatUnit"]) != "null" {
			if err := json.Unmarshal(fields["mainStatUnit"], &resource.MainStatUnit); err != nil {
				errs = append(errs, cog.MakeBuildErrors("mainStatUnit", err)...)
			}

		}
		delete(fields, "mainStatUnit")

	}
	// Field "secondaryStatUnit"
	if fields["secondaryStatUnit"] != nil {
		if string(fields["secondaryStatUnit"]) != "null" {
			if err := json.Unmarshal(fields["secondaryStatUnit"], &resource.SecondaryStatUnit); err != nil {
				errs = append(errs, cog.MakeBuildErrors("secondaryStatUnit", err)...)
			}

		}
		delete(fields, "secondaryStatUnit")

	}
	// Field "arcs"
	if fields["arcs"] != nil {
		if string(fields["arcs"]) != "null" {

			partialArray := []json.RawMessage{}
			if err := json.Unmarshal(fields["arcs"], &partialArray); err != nil {
				return err
			}

			for i1 := range partialArray {
				var result1 ArcOption

				result1 = ArcOption{}
				if err := result1.UnmarshalJSONStrict(partialArray[i1]); err != nil {
					errs = append(errs, cog.MakeBuildErrors("arcs["+strconv.Itoa(i1)+"]", err)...)
				}
				resource.Arcs = append(resource.Arcs, result1)
			}

		}
		delete(fields, "arcs")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("NodeOptions", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `NodeOptions` objects.
func (resource NodeOptions) Equals(other NodeOptions) bool {
	if resource.MainStatUnit == nil && other.MainStatUnit != nil || resource.MainStatUnit != nil && other.MainStatUnit == nil {
		return false
	}

	if resource.MainStatUnit != nil {
		if *resource.MainStatUnit != *other.MainStatUnit {
			return false
		}
	}
	if resource.SecondaryStatUnit == nil && other.SecondaryStatUnit != nil || resource.SecondaryStatUnit != nil && other.SecondaryStatUnit == nil {
		return false
	}

	if resource.SecondaryStatUnit != nil {
		if *resource.SecondaryStatUnit != *other.SecondaryStatUnit {
			return false
		}
	}

	if len(resource.Arcs) != len(other.Arcs) {
		return false
	}

	for i1 := range resource.Arcs {
		if !resource.Arcs[i1].Equals(other.Arcs[i1]) {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `NodeOptions` fields for violations and returns them.
func (resource NodeOptions) Validate() error {
	var errs cog.BuildErrors

	for i1 := range resource.Arcs {
		if err := resource.Arcs[i1].Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("arcs["+strconv.Itoa(i1)+"]", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type EdgeOptions struct {
	// Unit for the main stat to override what ever is set in the data frame.
	MainStatUnit *string `json:"mainStatUnit,omitempty"`
	// Unit for the secondary stat to override what ever is set in the data frame.
	SecondaryStatUnit *string `json:"secondaryStatUnit,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `EdgeOptions` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *EdgeOptions) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "mainStatUnit"
	if fields["mainStatUnit"] != nil {
		if string(fields["mainStatUnit"]) != "null" {
			if err := json.Unmarshal(fields["mainStatUnit"], &resource.MainStatUnit); err != nil {
				errs = append(errs, cog.MakeBuildErrors("mainStatUnit", err)...)
			}

		}
		delete(fields, "mainStatUnit")

	}
	// Field "secondaryStatUnit"
	if fields["secondaryStatUnit"] != nil {
		if string(fields["secondaryStatUnit"]) != "null" {
			if err := json.Unmarshal(fields["secondaryStatUnit"], &resource.SecondaryStatUnit); err != nil {
				errs = append(errs, cog.MakeBuildErrors("secondaryStatUnit", err)...)
			}

		}
		delete(fields, "secondaryStatUnit")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("EdgeOptions", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `EdgeOptions` objects.
func (resource EdgeOptions) Equals(other EdgeOptions) bool {
	if resource.MainStatUnit == nil && other.MainStatUnit != nil || resource.MainStatUnit != nil && other.MainStatUnit == nil {
		return false
	}

	if resource.MainStatUnit != nil {
		if *resource.MainStatUnit != *other.MainStatUnit {
			return false
		}
	}
	if resource.SecondaryStatUnit == nil && other.SecondaryStatUnit != nil || resource.SecondaryStatUnit != nil && other.SecondaryStatUnit == nil {
		return false
	}

	if resource.SecondaryStatUnit != nil {
		if *resource.SecondaryStatUnit != *other.SecondaryStatUnit {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `EdgeOptions` fields for violations and returns them.
func (resource EdgeOptions) Validate() error {
	return nil
}

type Options struct {
	Nodes *NodeOptions `json:"nodes,omitempty"`
	Edges *EdgeOptions `json:"edges,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `Options` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *Options) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "nodes"
	if fields["nodes"] != nil {
		if string(fields["nodes"]) != "null" {

			resource.Nodes = &NodeOptions{}
			if err := resource.Nodes.UnmarshalJSONStrict(fields["nodes"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("nodes", err)...)
			}

		}
		delete(fields, "nodes")

	}
	// Field "edges"
	if fields["edges"] != nil {
		if string(fields["edges"]) != "null" {

			resource.Edges = &EdgeOptions{}
			if err := resource.Edges.UnmarshalJSONStrict(fields["edges"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("edges", err)...)
			}

		}
		delete(fields, "edges")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("Options", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `Options` objects.
func (resource Options) Equals(other Options) bool {
	if resource.Nodes == nil && other.Nodes != nil || resource.Nodes != nil && other.Nodes == nil {
		return false
	}

	if resource.Nodes != nil {
		if !resource.Nodes.Equals(*other.Nodes) {
			return false
		}
	}
	if resource.Edges == nil && other.Edges != nil || resource.Edges != nil && other.Edges == nil {
		return false
	}

	if resource.Edges != nil {
		if !resource.Edges.Equals(*other.Edges) {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `Options` fields for violations and returns them.
func (resource Options) Validate() error {
	var errs cog.BuildErrors
	if resource.Nodes != nil {
		if err := resource.Nodes.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("nodes", err)...)
		}
	}
	if resource.Edges != nil {
		if err := resource.Edges.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("edges", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// VariantConfig returns the configuration related to nodegraph panels.
// This configuration describes how to unmarshal it, convert it to code, …
func VariantConfig() variants.PanelcfgConfig {
	return variants.PanelcfgConfig{
		Identifier: "nodegraph",
		OptionsUnmarshaler: func(raw []byte) (any, error) {
			options := &Options{}

			if err := json.Unmarshal(raw, options); err != nil {
				return nil, err
			}

			return options, nil
		},
		StrictOptionsUnmarshaler: func(raw []byte) (any, error) {
			options := &Options{}

			if err := options.UnmarshalJSONStrict(raw); err != nil {
				return nil, err
			}

			return options, nil
		},
		GoConverter: func(inputPanel any) string {
			if panel, ok := inputPanel.(*dashboard.Panel); ok {
				return PanelConverter(*panel)
			}

			return PanelConverter(inputPanel.(dashboard.Panel))
		},
	}
}
