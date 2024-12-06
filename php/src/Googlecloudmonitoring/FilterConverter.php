<?php

namespace Grafana\Foundation\Googlecloudmonitoring;

final class FilterConverter
{
    public static function convert(\Grafana\Foundation\Googlecloudmonitoring\Filter $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Googlecloudmonitoring\FilterBuilder())',
        ];
            if ($input->key !== "") {
    
        
    $buffer = 'key(';
        $arg0 =\var_export($input->key, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->operator !== "") {
    
        
    $buffer = 'operator(';
        $arg0 =\var_export($input->operator, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->value !== "") {
    
        
    $buffer = 'value(';
        $arg0 =\var_export($input->value, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->condition !== null && $input->condition !== "") {
    
        
    $buffer = 'condition(';
        $arg0 =\var_export($input->condition, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

