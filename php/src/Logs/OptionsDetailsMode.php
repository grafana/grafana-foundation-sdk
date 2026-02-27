<?php

namespace Grafana\Foundation\Logs;

final class OptionsDetailsMode implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, OptionsDetailsMode>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function inline(): self
    {
        if (!isset(self::$instances["Inline"])) {
            self::$instances["Inline"] = new self("inline");
        }

        return self::$instances["Inline"];
    }

    public static function sidebar(): self
    {
        if (!isset(self::$instances["Sidebar"])) {
            self::$instances["Sidebar"] = new self("sidebar");
        }

        return self::$instances["Sidebar"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "inline") {
            return self::inline();
        }

        if ($value === "sidebar") {
            return self::sidebar();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum OptionsDetailsMode");
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

