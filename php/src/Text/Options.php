<?php

namespace Grafana\Foundation\Text;

class Options implements \JsonSerializable
{
    public \Grafana\Foundation\Text\TextMode $mode;

    public ?\Grafana\Foundation\Text\CodeOptions $code;

    public string $content;

    /**
     * @param \Grafana\Foundation\Text\TextMode|null $mode
     * @param \Grafana\Foundation\Text\CodeOptions|null $code
     * @param string|null $content
     */
    public function __construct(?\Grafana\Foundation\Text\TextMode $mode = null, ?\Grafana\Foundation\Text\CodeOptions $code = null, ?string $content = null)
    {
        $this->mode = $mode ?: \Grafana\Foundation\Text\TextMode::markdown();
        $this->code = $code;
        $this->content = $content ?: "# Title\n\nFor markdown syntax help: [commonmark.org/help](https://commonmark.org/help/)";
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{mode?: string, code?: mixed, content?: string} $inputData */
        $data = $inputData;
        return new self(
            mode: isset($data["mode"]) ? (function($input) { return \Grafana\Foundation\Text\TextMode::fromValue($input); })($data["mode"]) : null,
            code: isset($data["code"]) ? (function($input) {
    	/** @var array{language?: string, showLineNumbers?: bool, showMiniMap?: bool} */
    $val = $input;
    	return \Grafana\Foundation\Text\CodeOptions::fromArray($val);
    })($data["code"]) : null,
            content: $data["content"] ?? null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->mode = $this->mode;
        $data->content = $this->content;
        if (isset($this->code)) {
            $data->code = $this->code;
        }
        return $data;
    }
}
