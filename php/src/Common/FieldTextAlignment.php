<?php

namespace Grafana\Foundation\Common;

/**
 * TODO -- should not be table specific!
 * TODO docs
 */
final class FieldTextAlignment implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, FieldTextAlignment>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function auto(): self
    {
        if (!isset(self::$instances["Auto"])) {
            self::$instances["Auto"] = new self("auto");
        }

        return self::$instances["Auto"];
    }

    public static function left(): self
    {
        if (!isset(self::$instances["Left"])) {
            self::$instances["Left"] = new self("left");
        }

        return self::$instances["Left"];
    }

    public static function right(): self
    {
        if (!isset(self::$instances["Right"])) {
            self::$instances["Right"] = new self("right");
        }

        return self::$instances["Right"];
    }

    public static function center(): self
    {
        if (!isset(self::$instances["Center"])) {
            self::$instances["Center"] = new self("center");
        }

        return self::$instances["Center"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "auto") {
            return self::auto();
        }

        if ($value === "left") {
            return self::left();
        }

        if ($value === "right") {
            return self::right();
        }

        if ($value === "center") {
            return self::center();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum FieldTextAlignment");
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

