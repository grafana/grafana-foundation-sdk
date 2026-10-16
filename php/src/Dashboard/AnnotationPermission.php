<?php

namespace Grafana\Foundation\Dashboard;

/**
 * +k8s:deepcopy-gen=true
 */
class AnnotationPermission implements \JsonSerializable
{
    /**
     * +k8s:deepcopy-gen=true
     */
    public ?\Grafana\Foundation\Dashboard\AnnotationActions $dashboard;

    /**
     * @param \Grafana\Foundation\Dashboard\AnnotationActions|null $dashboard
     */
    public function __construct(?\Grafana\Foundation\Dashboard\AnnotationActions $dashboard = null)
    {
        $this->dashboard = $dashboard;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{dashboard?: mixed} $inputData */
        $data = $inputData;
        return new self(
            dashboard: isset($data["dashboard"]) ? (function($input) {
    	/** @var array{canAdd?: bool, canDelete?: bool, canEdit?: bool} */
    $val = $input;
    	return \Grafana\Foundation\Dashboard\AnnotationActions::fromArray($val);
    })($data["dashboard"]) : null,
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
        return $data;
    }
}
