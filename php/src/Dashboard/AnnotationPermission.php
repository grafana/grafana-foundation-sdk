<?php

namespace Grafana\Foundation\Dashboard;

class AnnotationPermission implements \JsonSerializable
{
    public ?\Grafana\Foundation\Dashboard\AnnotationActions $dashboard;

    public ?\Grafana\Foundation\Dashboard\AnnotationActions $organization;

    /**
     * @param \Grafana\Foundation\Dashboard\AnnotationActions|null $dashboard
     * @param \Grafana\Foundation\Dashboard\AnnotationActions|null $organization
     */
    public function __construct(?\Grafana\Foundation\Dashboard\AnnotationActions $dashboard = null, ?\Grafana\Foundation\Dashboard\AnnotationActions $organization = null)
    {
        $this->dashboard = $dashboard;
        $this->organization = $organization;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{dashboard?: mixed, organization?: mixed} $inputData */
        $data = $inputData;
        return new self(
            dashboard: isset($data["dashboard"]) ? (function($input) {
    	/** @var array{canAdd?: bool, canDelete?: bool, canEdit?: bool} */
    $val = $input;
    	return \Grafana\Foundation\Dashboard\AnnotationActions::fromArray($val);
    })($data["dashboard"]) : null,
            organization: isset($data["organization"]) ? (function($input) {
    	/** @var array{canAdd?: bool, canDelete?: bool, canEdit?: bool} */
    $val = $input;
    	return \Grafana\Foundation\Dashboard\AnnotationActions::fromArray($val);
    })($data["organization"]) : null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        if (isset($this->dashboard)) {
            $data->dashboard = $this->dashboard;
        }
        if (isset($this->organization)) {
            $data->organization = $this->organization;
        }
        return $data;
    }
}
