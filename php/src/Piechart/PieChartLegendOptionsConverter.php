<?php

namespace Grafana\Foundation\Piechart;

final class PieChartLegendOptionsConverter
{
    public static function convert(\Grafana\Foundation\Piechart\PieChartLegendOptions $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Piechart\PieChartLegendOptionsBuilder())',
        ];
            if (count($input->values) >= 1) {
    
        
    $buffer = 'values(';
        $tmparg0 = [];
        foreach ($input->values as $arg1) {
        $tmpvaluesarg1 ='\Grafana\Foundation\Piechart\PieChartLegendValues::fromValue("'.$arg1.'")';
        $tmparg0[] = $tmpvaluesarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            
    
        {
    $buffer = 'displayMode(';
        $arg0 ='\Grafana\Foundation\Common\LegendDisplayMode::fromValue("'.$input->displayMode.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            
    
        {
    $buffer = 'placement(';
        $arg0 ='\Grafana\Foundation\Common\LegendPlacement::fromValue("'.$input->placement.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            
    
        {
    $buffer = 'showLegend(';
        $arg0 =\var_export($input->showLegend, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            if ($input->asTable !== null) {
    
        
    $buffer = 'asTable(';
        $arg0 =\var_export($input->asTable, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->isVisible !== null) {
    
        
    $buffer = 'isVisible(';
        $arg0 =\var_export($input->isVisible, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->sortBy !== null && $input->sortBy !== "") {
    
        
    $buffer = 'sortBy(';
        $arg0 =\var_export($input->sortBy, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->sortDesc !== null) {
    
        
    $buffer = 'sortDesc(';
        $arg0 =\var_export($input->sortDesc, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->width !== null) {
    
        
    $buffer = 'width(';
        $arg0 =\var_export($input->width, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if (count($input->calcs) >= 1) {
    
        
    $buffer = 'calcs(';
        $tmparg0 = [];
        foreach ($input->calcs as $arg1) {
        $tmpcalcsarg1 =\var_export($arg1, true);
        $tmparg0[] = $tmpcalcsarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

