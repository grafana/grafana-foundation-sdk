<?php

namespace Grafana\Foundation\Dashboard;

final class RegexMapConverter
{
    public static function convert(\Grafana\Foundation\Dashboard\RegexMap $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboard\RegexMapBuilder())',
        ];
            
    
        {
    $buffer = 'options(';
        $arg0 = \Grafana\Foundation\Dashboard\DashboardRegexMapOptionsConverter::convert($input->options);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    

        return \implode("\n\t->", $calls);
    }
}

