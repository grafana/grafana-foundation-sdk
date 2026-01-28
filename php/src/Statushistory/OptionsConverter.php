<?php

namespace Grafana\Foundation\Statushistory;

final class OptionsConverter
{
    public static function convert(\Grafana\Foundation\Statushistory\Options $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Statushistory\OptionsBuilder())',
        ];
            if ($input->rowHeight !== (float) 0.9) {
    
        
    $buffer = 'rowHeight(';
        $arg0 =\var_export($input->rowHeight, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            
    
        {
    $buffer = 'showValue(';
        $arg0 ='\Grafana\Foundation\Common\VisibilityMode::fromValue("'.$input->showValue.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            if ($input->colWidth !== null && $input->colWidth !== (float) 0.9) {
    
        
    $buffer = 'colWidth(';
        $arg0 =\var_export($input->colWidth, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            
    
        {
    $buffer = 'legend(';
        $arg0 = \Grafana\Foundation\Common\VizLegendOptionsConverter::convert($input->legend);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            
    
        {
    $buffer = 'tooltip(';
        $arg0 = \Grafana\Foundation\Common\VizTooltipOptionsConverter::convert($input->tooltip);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            if ($input->timezone !== null && count($input->timezone) >= 1) {
    
        
    $buffer = 'timezone(';
        $tmparg0 = [];
        foreach ($input->timezone as $arg1) {
        $tmptimezonearg1 =\var_export($arg1, true);
        $tmparg0[] = $tmptimezonearg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->perPage !== null && $input->perPage !== (float) 20) {
    
        
    $buffer = 'perPage(';
        $arg0 =\var_export($input->perPage, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

