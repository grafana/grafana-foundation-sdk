<?php

namespace Grafana\Foundation\Grafanapyroscope;

final class PyroscopeQueryType implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, PyroscopeQueryType>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function metrics(): self
    {
        if (!isset(self::$instances["Metrics"])) {
            self::$instances["Metrics"] = new self("metrics");
        }

        return self::$instances["Metrics"];
    }

    public static function profile(): self
    {
        if (!isset(self::$instances["Profile"])) {
            self::$instances["Profile"] = new self("profile");
        }

        return self::$instances["Profile"];
    }

    public static function both(): self
    {
        if (!isset(self::$instances["Both"])) {
            self::$instances["Both"] = new self("both");
        }

        return self::$instances["Both"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "metrics") {
            return self::metrics();
        }

        if ($value === "profile") {
            return self::profile();
        }

        if ($value === "both") {
            return self::both();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum PyroscopeQueryType");
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

