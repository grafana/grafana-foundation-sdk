<?php

namespace Grafana\Foundation\Candlestick;

final class CandlestickColorsConverter
{
    public static function convert(\Grafana\Foundation\Candlestick\CandlestickColors $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Candlestick\CandlestickColorsBuilder())',
        ];
            if ($input->up !== "" && $input->up !== "green") {
    
        
    $buffer = 'up(';
        $arg0 =\var_export($input->up, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->down !== "" && $input->down !== "red") {
    
        
    $buffer = 'down(';
        $arg0 =\var_export($input->down, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->flat !== "" && $input->flat !== "gray") {
    
        
    $buffer = 'flat(';
        $arg0 =\var_export($input->flat, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

