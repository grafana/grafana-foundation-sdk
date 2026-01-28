<?php

namespace Grafana\Foundation\Piechart;

final class FieldConfigConverter
{
    public static function convert(\Grafana\Foundation\Piechart\FieldConfig $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Piechart\FieldConfigBuilder())',
        ];
            if ($input->hideFrom !== null) {
    
        
    $buffer = 'hideFrom(';
        $arg0 = \Grafana\Foundation\Common\HideSeriesConfigConverter::convert($input->hideFrom);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

