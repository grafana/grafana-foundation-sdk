<?php

namespace Grafana\Foundation\Dashboard;

/**
 * Dashboard action
 */
class Action implements \JsonSerializable
{
    public \Grafana\Foundation\Dashboard\ActionType $type;

    public string $title;

    public ?\Grafana\Foundation\Dashboard\FetchOptions $fetch;

    public ?\Grafana\Foundation\Dashboard\InfinityOptions $infinity;

    public ?string $confirmation;

    public ?bool $oneClick;

    /**
     * @var array<\Grafana\Foundation\Dashboard\ActionVariable>|null
     */
    public ?array $variables;

    public ?\Grafana\Foundation\Dashboard\DashboardActionStyle $style;

    /**
     * @param \Grafana\Foundation\Dashboard\ActionType|null $type
     * @param string|null $title
     * @param \Grafana\Foundation\Dashboard\FetchOptions|null $fetch
     * @param \Grafana\Foundation\Dashboard\InfinityOptions|null $infinity
     * @param string|null $confirmation
     * @param bool|null $oneClick
     * @param array<\Grafana\Foundation\Dashboard\ActionVariable>|null $variables
     * @param \Grafana\Foundation\Dashboard\DashboardActionStyle|null $style
     */
    public function __construct(?\Grafana\Foundation\Dashboard\ActionType $type = null, ?string $title = null, ?\Grafana\Foundation\Dashboard\FetchOptions $fetch = null, ?\Grafana\Foundation\Dashboard\InfinityOptions $infinity = null, ?string $confirmation = null, ?bool $oneClick = null, ?array $variables = null, ?\Grafana\Foundation\Dashboard\DashboardActionStyle $style = null)
    {
        $this->type = $type ?: \Grafana\Foundation\Dashboard\ActionType::Fetch();
        $this->title = $title ?: "";
        $this->fetch = $fetch;
        $this->infinity = $infinity;
        $this->confirmation = $confirmation;
        $this->oneClick = $oneClick;
        $this->variables = $variables;
        $this->style = $style;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{type?: string, title?: string, fetch?: mixed, infinity?: mixed, confirmation?: string, oneClick?: bool, variables?: array<mixed>, style?: mixed} $inputData */
        $data = $inputData;
        return new self(
            type: isset($data["type"]) ? (function($input) { return \Grafana\Foundation\Dashboard\ActionType::fromValue($input); })($data["type"]) : null,
            title: $data["title"] ?? null,
            fetch: isset($data["fetch"]) ? (function($input) {
    	/** @var array{method?: string, url?: string, body?: string, queryParams?: array<array<string>>, headers?: array<array<string>>} */
    $val = $input;
    	return \Grafana\Foundation\Dashboard\FetchOptions::fromArray($val);
    })($data["fetch"]) : null,
            infinity: isset($data["infinity"]) ? (function($input) {
    	/** @var array{method?: string, url?: string, body?: string, queryParams?: array<array<string>>, headers?: array<array<string>>, datasourceUid?: string} */
    $val = $input;
    	return \Grafana\Foundation\Dashboard\InfinityOptions::fromArray($val);
    })($data["infinity"]) : null,
            confirmation: $data["confirmation"] ?? null,
            oneClick: $data["oneClick"] ?? null,
            variables: array_filter(array_map((function($input) {
    	/** @var array{key?: string, name?: string, type?: "string"} */
    $val = $input;
    	return \Grafana\Foundation\Dashboard\ActionVariable::fromArray($val);
    }), $data["variables"] ?? [])),
            style: isset($data["style"]) ? (function($input) {
    	/** @var array{backgroundColor?: string} */
    $val = $input;
    	return \Grafana\Foundation\Dashboard\DashboardActionStyle::fromArray($val);
    })($data["style"]) : null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->type = $this->type;
        $data->title = $this->title;
        if (isset($this->fetch)) {
            $data->fetch = $this->fetch;
        }
        if (isset($this->infinity)) {
            $data->infinity = $this->infinity;
        }
        if (isset($this->confirmation)) {
            $data->confirmation = $this->confirmation;
        }
        if (isset($this->oneClick)) {
            $data->oneClick = $this->oneClick;
        }
        if (isset($this->variables)) {
            $data->variables = $this->variables;
        }
        if (isset($this->style)) {
            $data->style = $this->style;
        }
        return $data;
    }
}
