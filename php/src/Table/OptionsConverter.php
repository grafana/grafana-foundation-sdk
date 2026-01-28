<?php

namespace Grafana\Foundation\Table;

final class OptionsConverter
{
    public static function convert(\Grafana\Foundation\Table\Options $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Table\OptionsBuilder())',
        ];
            if ($input->frameIndex !== (float) 0) {
    
        
    $buffer = 'frameIndex(';
        $arg0 =\var_export($input->frameIndex, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->showHeader !== true) {
    
        
    $buffer = 'showHeader(';
        $arg0 =\var_export($input->showHeader, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->showTypeIcons !== null && $input->showTypeIcons !== false) {
    
        
    $buffer = 'showTypeIcons(';
        $arg0 =\var_export($input->showTypeIcons, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->sortBy !== null && count($input->sortBy) >= 1) {
    
        
    $buffer = 'sortBy(';
        $tmparg0 = [];
        foreach ($input->sortBy as $arg1) {
        $tmpsortByarg1 = \Grafana\Foundation\Common\TableSortByFieldStateConverter::convert($arg1);
        $tmparg0[] = $tmpsortByarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->footer !== null) {
    
        
    $buffer = 'footer(';
        $arg0 = \Grafana\Foundation\Common\TableFooterOptionsConverter::convert($input->footer);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->cellHeight !== null) {
    
        
    $buffer = 'cellHeight(';
        $arg0 ='\Grafana\Foundation\Common\TableCellHeight::fromValue("'.$input->cellHeight.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

