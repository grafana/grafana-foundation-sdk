<?php

namespace Grafana\Foundation\Common;

final class StackableFieldConfigConverter
{
    public static function convert(\Grafana\Foundation\Common\StackableFieldConfig $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Common\StackableFieldConfigBuilder())',
        ];
            if ($input->stacking !== null) {
    
        
    $buffer = 'stacking(';
        $arg0 = \Grafana\Foundation\Common\StackingConfigConverter::convert($input->stacking);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

