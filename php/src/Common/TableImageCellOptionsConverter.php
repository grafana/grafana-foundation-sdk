<?php

namespace Grafana\Foundation\Common;

final class TableImageCellOptionsConverter
{
    public static function convert(\Grafana\Foundation\Common\TableImageCellOptions $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Common\TableImageCellOptionsBuilder())',
        ];
            if ($input->alt !== null && $input->alt !== "") {
    
        
    $buffer = 'alt(';
        $arg0 =\var_export($input->alt, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->title !== null && $input->title !== "") {
    
        
    $buffer = 'title(';
        $arg0 =\var_export($input->title, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

