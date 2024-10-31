<?php

namespace Grafana\Foundation\Dashboard;

final class SpecialValueMapConverter
{
    public static function convert(\Grafana\Foundation\Dashboard\SpecialValueMap $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboard\SpecialValueMapBuilder())',
        ];
            
    
        {
    $buffer = 'options(';
        $arg0 = \Grafana\Foundation\Dashboard\DashboardSpecialValueMapOptionsConverter::convert($input->options);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    

        return \implode("\n\t->", $calls);
    }
}

