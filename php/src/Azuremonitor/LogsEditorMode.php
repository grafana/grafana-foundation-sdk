<?php

namespace Grafana\Foundation\Azuremonitor;

final class LogsEditorMode implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, LogsEditorMode>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function builder(): self
    {
        if (!isset(self::$instances["Builder"])) {
            self::$instances["Builder"] = new self("builder");
        }

        return self::$instances["Builder"];
    }

    public static function raw(): self
    {
        if (!isset(self::$instances["Raw"])) {
            self::$instances["Raw"] = new self("raw");
        }

        return self::$instances["Raw"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "builder") {
            return self::builder();
        }

        if ($value === "raw") {
            return self::raw();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum LogsEditorMode");
    }

    public function jsonSerialize(): string
    {
        return $this->value;
    }

    public function __toString(): string
    {
        return $this->value;
    }
}

