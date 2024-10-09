<?php

namespace Grafana\Foundation\Common;

final class HideableFieldConfigConverter
{
    public static function convert(\Grafana\Foundation\Common\HideableFieldConfig $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Common\HideableFieldConfigBuilder())',
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

