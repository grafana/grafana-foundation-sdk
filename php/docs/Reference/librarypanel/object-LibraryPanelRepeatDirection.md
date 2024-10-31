---
title: <span class="badge object-type-enum"></span> LibraryPanelRepeatDirection
---
# <span class="badge object-type-enum"></span> LibraryPanelRepeatDirection

## Definition

```php
final class LibraryPanelRepeatDirection implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, LibraryPanelRepeatDirection>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function h(): self
    {
        if (!isset(self::$instances["H"])) {
            self::$instances["H"] = new self("h");
        }

        return self::$instances["H"];
    }

    public static function v(): self
    {
        if (!isset(self::$instances["V"])) {
            self::$instances["V"] = new self("v");
        }

        return self::$instances["V"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "h") {
            return self::h();
        }

        if ($value === "v") {
            return self::v();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum LibraryPanelRepeatDirection");
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
