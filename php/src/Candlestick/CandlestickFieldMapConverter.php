<?php

namespace Grafana\Foundation\Candlestick;

final class CandlestickFieldMapConverter
{
    public static function convert(\Grafana\Foundation\Candlestick\CandlestickFieldMap $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Candlestick\CandlestickFieldMapBuilder())',
        ];
            if ($input->open !== null && $input->open !== "") {
    
        
    $buffer = 'open(';
        $arg0 =\var_export($input->open, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->high !== null && $input->high !== "") {
    
        
    $buffer = 'high(';
        $arg0 =\var_export($input->high, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->low !== null && $input->low !== "") {
    
        
    $buffer = 'low(';
        $arg0 =\var_export($input->low, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->close !== null && $input->close !== "") {
    
        
    $buffer = 'close(';
        $arg0 =\var_export($input->close, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->volume !== null && $input->volume !== "") {
    
        
    $buffer = 'volume(';
        $arg0 =\var_export($input->volume, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

