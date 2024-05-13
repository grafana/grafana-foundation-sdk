// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package alerting

import (
	"encoding/json"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

func (resource *Query) UnmarshalJSON(raw []byte) error {
	if raw == nil {
		return nil
	}
	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}

	if fields["datasourceUid"] != nil {
		if err := json.Unmarshal(fields["datasourceUid"], &resource.DatasourceUid); err != nil {
			return err
		}
	}

	if fields["queryType"] != nil {
		if err := json.Unmarshal(fields["queryType"], &resource.QueryType); err != nil {
			return err
		}
	}

	if fields["refId"] != nil {
		if err := json.Unmarshal(fields["refId"], &resource.RefId); err != nil {
			return err
		}
	}

	if fields["relativeTimeRange"] != nil {
		if err := json.Unmarshal(fields["relativeTimeRange"], &resource.RelativeTimeRange); err != nil {
			return err
		}
	}

	dataqueryTypeHint := ""

	if fields["model"] != nil {
		model, err := cog.UnmarshalDataquery(fields["model"], dataqueryTypeHint)
		if err != nil {
			return err
		}
		resource.Model = model
	}

	return nil
}
