<?php

namespace Grafana\Foundation\Common;

final class TableColoredBackgroundCellOptionsConverter
{
    public static function convert(\Grafana\Foundation\Common\TableColoredBackgroundCellOptions $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Common\TableColoredBackgroundCellOptionsBuilder())',
        ];
            if ($input->mode !== null) {
    
        
    $buffer = 'mode(';
        $arg0 ='\Grafana\Foundation\Common\TableCellBackgroundDisplayMode::fromValue("'.$input->mode.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->applyToRow !== null) {
    
        
    $buffer = 'applyToRow(';
        $arg0 =\var_export($input->applyToRow, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

