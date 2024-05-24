// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package common

import (
	"encoding/json"
	"errors"
	"fmt"
)

func (resource BoolOrFloat64) MarshalJSON() ([]byte, error) {
	if resource.Bool != nil {
		return json.Marshal(resource.Bool)
	}

	if resource.Float64 != nil {
		return json.Marshal(resource.Float64)
	}

	return nil, fmt.Errorf("no value for disjunction of scalars")
}

func (resource *BoolOrFloat64) UnmarshalJSON(raw []byte) error {
	if raw == nil {
		return nil
	}

	var errList []error

	// Bool
	var Bool bool
	if err := json.Unmarshal(raw, &Bool); err != nil {
		errList = append(errList, err)
		resource.Bool = nil
	} else {
		resource.Bool = &Bool
		return nil
	}

	// Float64
	var Float64 float64
	if err := json.Unmarshal(raw, &Float64); err != nil {
		errList = append(errList, err)
		resource.Float64 = nil
	} else {
		resource.Float64 = &Float64
		return nil
	}

	return errors.Join(errList...)
}

func (resource BoolOrUint32) MarshalJSON() ([]byte, error) {
	if resource.Bool != nil {
		return json.Marshal(resource.Bool)
	}

	if resource.Uint32 != nil {
		return json.Marshal(resource.Uint32)
	}

	return nil, fmt.Errorf("no value for disjunction of scalars")
}

func (resource *BoolOrUint32) UnmarshalJSON(raw []byte) error {
	if raw == nil {
		return nil
	}

	var errList []error

	// Bool
	var Bool bool
	if err := json.Unmarshal(raw, &Bool); err != nil {
		errList = append(errList, err)
		resource.Bool = nil
	} else {
		resource.Bool = &Bool
		return nil
	}

	// Uint32
	var Uint32 uint32
	if err := json.Unmarshal(raw, &Uint32); err != nil {
		errList = append(errList, err)
		resource.Uint32 = nil
	} else {
		resource.Uint32 = &Uint32
		return nil
	}

	return errors.Join(errList...)
}

func (resource TableAutoCellOptionsOrTableSparklineCellOptionsOrTableBarGaugeCellOptionsOrTableColoredBackgroundCellOptionsOrTableColorTextCellOptionsOrTableImageCellOptionsOrTableDataLinksCellOptionsOrTableJsonViewCellOptions) MarshalJSON() ([]byte, error) {
	if resource.TableAutoCellOptions != nil {
		return json.Marshal(resource.TableAutoCellOptions)
	}
	if resource.TableSparklineCellOptions != nil {
		return json.Marshal(resource.TableSparklineCellOptions)
	}
	if resource.TableBarGaugeCellOptions != nil {
		return json.Marshal(resource.TableBarGaugeCellOptions)
	}
	if resource.TableColoredBackgroundCellOptions != nil {
		return json.Marshal(resource.TableColoredBackgroundCellOptions)
	}
	if resource.TableColorTextCellOptions != nil {
		return json.Marshal(resource.TableColorTextCellOptions)
	}
	if resource.TableImageCellOptions != nil {
		return json.Marshal(resource.TableImageCellOptions)
	}
	if resource.TableDataLinksCellOptions != nil {
		return json.Marshal(resource.TableDataLinksCellOptions)
	}
	if resource.TableJsonViewCellOptions != nil {
		return json.Marshal(resource.TableJsonViewCellOptions)
	}

	return nil, fmt.Errorf("no value for disjunction of refs")
}
func (resource *TableAutoCellOptionsOrTableSparklineCellOptionsOrTableBarGaugeCellOptionsOrTableColoredBackgroundCellOptionsOrTableColorTextCellOptionsOrTableImageCellOptionsOrTableDataLinksCellOptionsOrTableJsonViewCellOptions) UnmarshalJSON(raw []byte) error {
	if raw == nil {
		return nil
	}

	// FIXME: this is wasteful, we need to find a more efficient way to unmarshal this.
	parsedAsMap := make(map[string]any)
	if err := json.Unmarshal(raw, &parsedAsMap); err != nil {
		return err
	}

	discriminator, found := parsedAsMap["type"]
	if !found {
		return errors.New("discriminator field 'type' not found in payload")
	}

	switch discriminator {
	case "auto":
		var tableAutoCellOptions TableAutoCellOptions
		if err := json.Unmarshal(raw, &tableAutoCellOptions); err != nil {
			return err
		}

		resource.TableAutoCellOptions = &tableAutoCellOptions
		return nil
	case "color-background":
		var tableColoredBackgroundCellOptions TableColoredBackgroundCellOptions
		if err := json.Unmarshal(raw, &tableColoredBackgroundCellOptions); err != nil {
			return err
		}

		resource.TableColoredBackgroundCellOptions = &tableColoredBackgroundCellOptions
		return nil
	case "color-text":
		var tableColorTextCellOptions TableColorTextCellOptions
		if err := json.Unmarshal(raw, &tableColorTextCellOptions); err != nil {
			return err
		}

		resource.TableColorTextCellOptions = &tableColorTextCellOptions
		return nil
	case "data-links":
		var tableDataLinksCellOptions TableDataLinksCellOptions
		if err := json.Unmarshal(raw, &tableDataLinksCellOptions); err != nil {
			return err
		}

		resource.TableDataLinksCellOptions = &tableDataLinksCellOptions
		return nil
	case "gauge":
		var tableBarGaugeCellOptions TableBarGaugeCellOptions
		if err := json.Unmarshal(raw, &tableBarGaugeCellOptions); err != nil {
			return err
		}

		resource.TableBarGaugeCellOptions = &tableBarGaugeCellOptions
		return nil
	case "image":
		var tableImageCellOptions TableImageCellOptions
		if err := json.Unmarshal(raw, &tableImageCellOptions); err != nil {
			return err
		}

		resource.TableImageCellOptions = &tableImageCellOptions
		return nil
	case "json-view":
		var tableJsonViewCellOptions TableJsonViewCellOptions
		if err := json.Unmarshal(raw, &tableJsonViewCellOptions); err != nil {
			return err
		}

		resource.TableJsonViewCellOptions = &tableJsonViewCellOptions
		return nil
	case "sparkline":
		var tableSparklineCellOptions TableSparklineCellOptions
		if err := json.Unmarshal(raw, &tableSparklineCellOptions); err != nil {
			return err
		}

		resource.TableSparklineCellOptions = &tableSparklineCellOptions
		return nil
	}

	return fmt.Errorf("could not unmarshal resource with `type = %v`", discriminator)
}
