<?php

namespace {{ .Data.NamespaceRoot }}\Cog;

/**
 * @implements \ArrayAccess<string, mixed>
 */
final class UnknownDataquery implements \ArrayAccess, \JsonSerializable, Dataquery
{
    /**
     * @var array<string, mixed>
     */
    private $data = [];

    /**
     * @param array<string, mixed> $data
     */
	public function __construct(array $data)
	{
	    $this->data = $data;
	}

    /**
     * @return array<string, mixed>
     */
    public function toArray(): array
    {
        return $this->data;
    }

    public function dataqueryType(): string
    {
        return 'unknown';
    }

    /**
     * @param string $offset
     * @param mixed $value
     */
    public function offsetSet($offset, $value): void
    {
        $this->data[$offset] = $value;
    }

    /**
     * @param string $offset
     */
    public function offsetExists($offset): bool
    {
        return \array_key_exists($offset, $this->data);
    }

    /**
     * @param string $offset
     */
    public function offsetUnset($offset): void
    {
        unset($this->data[$offset]);
    }

    /**
     * @param string $offset
     */
    public function offsetGet($offset): mixed
    {
        if (!\array_key_exists($offset, $this->data)) {
            throw new \ValueError("offset '$offset' does not exist");
        }
        return $this->data[$offset] ?? null;
    }

    public function jsonSerialize(): mixed
    {
        return $this->data;
    }
}
