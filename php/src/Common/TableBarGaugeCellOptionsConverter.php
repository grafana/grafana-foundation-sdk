<?php

namespace Grafana\Foundation\Common;

final class TableBarGaugeCellOptionsConverter
{
    public static function convert(\Grafana\Foundation\Common\TableBarGaugeCellOptions $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Common\TableBarGaugeCellOptionsBuilder())',
        ];
            if ($input->mode !== null) {
    
        
    $buffer = 'mode(';
        $arg0 ='\Grafana\Foundation\Common\BarGaugeDisplayMode::fromValue("'.$input->mode.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->valueDisplayMode !== null) {
    
        
    $buffer = 'valueDisplayMode(';
        $arg0 ='\Grafana\Foundation\Common\BarGaugeValueMode::fromValue("'.$input->valueDisplayMode.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

