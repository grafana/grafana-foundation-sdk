<?php

namespace Grafana\Foundation\Dashboardv2beta1;

final class AutoGridLayoutSpecConverter
{
    public static function convert(\Grafana\Foundation\Dashboardv2beta1\AutoGridLayoutSpec $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboardv2beta1\AutoGridLayoutSpecBuilder())',
        ];
            if ($input->maxColumnCount !== null && $input->maxColumnCount !== (float) 3) {
    
        
    $buffer = 'maxColumnCount(';
        $arg0 =\var_export($input->maxColumnCount, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            
    
        {
    $buffer = 'columnWidthMode(';
        $arg0 ='\Grafana\Foundation\Dashboardv2beta1\AutoGridLayoutSpecColumnWidthMode::fromValue("'.$input->columnWidthMode.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            if ($input->columnWidth !== null) {
    
        
    $buffer = 'columnWidth(';
        $arg0 =\var_export($input->columnWidth, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            
    
        {
    $buffer = 'rowHeightMode(';
        $arg0 ='\Grafana\Foundation\Dashboardv2beta1\AutoGridLayoutSpecRowHeightMode::fromValue("'.$input->rowHeightMode.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            if ($input->rowHeight !== null) {
    
        
    $buffer = 'rowHeight(';
        $arg0 =\var_export($input->rowHeight, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->fillScreen !== null && $input->fillScreen !== false) {
    
        
    $buffer = 'fillScreen(';
        $arg0 =\var_export($input->fillScreen, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if (count($input->items) >= 1) {
    
        
    $buffer = 'items(';
        $tmparg0 = [];
        foreach ($input->items as $arg1) {
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

