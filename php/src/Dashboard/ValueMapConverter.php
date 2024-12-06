<?php

namespace Grafana\Foundation\Dashboard;

final class ValueMapConverter
{
    public static function convert(\Grafana\Foundation\Dashboard\ValueMap $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboard\ValueMapBuilder())',
        ];
            
    
        {
    $buffer = 'options(';
        $arg0 = "[";
        foreach ($input->options as $key => $arg1) {
            $tmpoptionsarg1 = \Grafana\Foundation\Dashboard\ValueMappingResultConverter::convert($arg1);
            $arg0 .= "\t".var_export($key, true)." => $tmpoptionsarg1,";
        }
        $arg0 .= "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    

        return \implode("\n\t->", $calls);
    }
}

