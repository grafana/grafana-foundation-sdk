<?php

namespace Grafana\Foundation\Dashboardv2beta1;

final class AutoGridLayoutConverter
{
    public static function convert(\Grafana\Foundation\Dashboardv2beta1\AutoGridLayoutKind $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboardv2beta1\AutoGridLayoutBuilder())',
        ];
            if ($input->spec->maxColumnCount !== null && $input->spec->maxColumnCount !== (float) 3) {
    
        
    $buffer = 'maxColumnCount(';
        $arg0 =\var_export($input->spec->maxColumnCount, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            
    
        {
    $buffer = 'columnWidthMode(';
        $arg0 ='\Grafana\Foundation\Dashboardv2beta1\AutoGridLayoutSpecColumnWidthMode::fromValue("'.$input->spec->columnWidthMode.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            if ($input->spec->columnWidth !== null) {
    
        
    $buffer = 'columnWidth(';
        $arg0 =\var_export($input->spec->columnWidth, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            
    
        {
    $buffer = 'rowHeightMode(';
        $arg0 ='\Grafana\Foundation\Dashboardv2beta1\AutoGridLayoutSpecRowHeightMode::fromValue("'.$input->spec->rowHeightMode.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            if ($input->spec->rowHeight !== null) {
    
        
    $buffer = 'rowHeight(';
        $arg0 =\var_export($input->spec->rowHeight, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec->fillScreen !== null && $input->spec->fillScreen !== false) {
    
        
    $buffer = 'fillScreen(';
        $arg0 =\var_export($input->spec->fillScreen, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if (count($input->spec->items) >= 1) {
    
        
    $buffer = 'items(';
        $tmparg0 = [];
        foreach ($input->spec->items as $arg1) {
        $tmpitemsarg1 = \Grafana\Foundation\Dashboardv2beta1\AutoGridLayoutItemConverter::convert($arg1);
        $tmparg0[] = $tmpitemsarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

