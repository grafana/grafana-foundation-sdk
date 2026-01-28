<?php

namespace Grafana\Foundation\Debug;

final class OptionsConverter
{
    public static function convert(\Grafana\Foundation\Debug\Options $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Debug\OptionsBuilder())',
        ];
            
    
        {
    $buffer = 'mode(';
        $arg0 ='\Grafana\Foundation\Debug\DebugMode::fromValue("'.$input->mode.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            if ($input->counters !== null) {
    
        
    $buffer = 'counters(';
        $arg0 = \Grafana\Foundation\Debug\UpdateConfigConverter::convert($input->counters);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

