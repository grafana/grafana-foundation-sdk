---
title: <span class="badge object-type-enum"></span> TypeReduceMode
---
# <span class="badge object-type-enum"></span> TypeReduceMode

## Definition

```php
final class TypeReduceMode implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, TypeReduceMode>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function dropNN(): self
    {
        if (!isset(self::$instances["DropNN"])) {
            self::$instances["DropNN"] = new self("dropNN");
        }

        return self::$instances["DropNN"];
    }

    public static function replaceNN(): self
    {
        if (!isset(self::$instances["ReplaceNN"])) {
            self::$instances["ReplaceNN"] = new self("replaceNN");
        }

        return self::$instances["ReplaceNN"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "dropNN") {
            return self::dropNN();
        }

        if ($value === "replaceNN") {
            return self::replaceNN();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum TypeReduceMode");
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

```
