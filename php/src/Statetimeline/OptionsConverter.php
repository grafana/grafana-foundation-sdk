<?php

namespace Grafana\Foundation\Statetimeline;

final class OptionsConverter
{
    public static function convert(\Grafana\Foundation\Statetimeline\Options $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Statetimeline\OptionsBuilder())',
        ];
            
    
        {
    $buffer = 'showValue(';
        $arg0 ='\Grafana\Foundation\Common\VisibilityMode::fromValue("'.$input->showValue.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            if ($input->rowHeight !== (float) 0.9) {
    
        
    $buffer = 'rowHeight(';
        $arg0 =\var_export($input->rowHeight, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->mergeValues !== null && $input->mergeValues !== true) {
    
        
    $buffer = 'mergeValues(';
        $arg0 =\var_export($input->mergeValues, true);
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
            if ($input->alignValue !== null) {
    
        
    $buffer = 'alignValue(';
        $arg0 ='\Grafana\Foundation\Common\TimelineValueAlignment::fromValue("'.$input->alignValue.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

