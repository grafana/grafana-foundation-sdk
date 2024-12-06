<?php

namespace Grafana\Foundation\Common;

final class TableAutoCellOptionsConverter
{
    public static function convert(\Grafana\Foundation\Common\TableAutoCellOptions $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Common\TableAutoCellOptionsBuilder())',
        ];
            if ($input->wrapText !== null) {
    
        
    $buffer = 'wrapText(';
        $arg0 =\var_export($input->wrapText, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

