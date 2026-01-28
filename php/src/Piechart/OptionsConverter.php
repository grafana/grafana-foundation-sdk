<?php

namespace Grafana\Foundation\Piechart;

final class OptionsConverter
{
    public static function convert(\Grafana\Foundation\Piechart\Options $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Piechart\OptionsBuilder())',
        ];
            
    
        {
    $buffer = 'pieType(';
        $arg0 ='\Grafana\Foundation\Piechart\PieChartType::fromValue("'.$input->pieType.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            if ($input->displayLabels !== null && count($input->displayLabels) >= 1) {
    
        
    $buffer = 'displayLabels(';
        $tmparg0 = [];
        foreach ($input->displayLabels as $arg1) {
        $tmpdisplayLabelsarg1 ='\Grafana\Foundation\Piechart\PieChartLabels::fromValue("'.$arg1.'")';
        $tmparg0[] = $tmpdisplayLabelsarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
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
    
    
            
    
        {
    $buffer = 'reduceOptions(';
        $arg0 = \Grafana\Foundation\Common\ReduceDataOptionsConverter::convert($input->reduceOptions);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            if ($input->text !== null) {
    
        
    $buffer = 'text(';
        $arg0 = \Grafana\Foundation\Common\VizTextDisplayOptionsConverter::convert($input->text);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            
    
        {
    $buffer = 'legend(';
        $arg0 = \Grafana\Foundation\Piechart\PieChartLegendOptionsConverter::convert($input->legend);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            
    
        {
    $buffer = 'orientation(';
        $arg0 ='\Grafana\Foundation\Common\VizOrientation::fromValue("'.$input->orientation.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    

        return \implode("\n\t->", $calls);
    }
}

