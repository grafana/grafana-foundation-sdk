<?php

namespace Grafana\Foundation\Dashboardv2beta1;

final class NotebookSpecConverter
{
    public static function convert(\Grafana\Foundation\Dashboardv2beta1\NotebookSpec $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboardv2beta1\NotebookSpecBuilder())',
        ];
            if ($input->title !== "") {
    
        
    $buffer = 'title(';
        $arg0 =\var_export($input->title, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->description !== null && $input->description !== "") {
    
        
    $buffer = 'description(';
        $arg0 =\var_export($input->description, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if (count($input->tags) >= 1) {
    
        
    $buffer = 'tags(';
        $tmparg0 = [];
        foreach ($input->tags as $arg1) {
        $tmptagsarg1 =\var_export($arg1, true);
        $tmparg0[] = $tmptagsarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            
    
        {
    $buffer = 'timeSettings(';
        $arg0 = \Grafana\Foundation\Dashboardv2beta1\TimeSettingsConverter::convert($input->timeSettings);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            
    
        {
    $buffer = 'elements(';
        $arg0 = "[";
        foreach ($input->elements as $key => $arg1) {
            switch (true) {
            case $arg1 instanceof \Grafana\Foundation\Dashboardv2beta1\CellKind:
                $disjunctionarg1 = \Grafana\Foundation\Dashboardv2beta1\CellKindConverter::convert($arg1);
                $tmpelementsarg1 = $disjunctionarg1;
                break;
            case $arg1 instanceof \Grafana\Foundation\Dashboardv2beta1\V2PanelKind:
                $disjunctionarg1 = \Grafana\Foundation\Dashboardv2beta1\V2PanelKindConverter::convert($arg1);
                $tmpelementsarg1 = $disjunctionarg1;
                break;
            case $arg1 instanceof \Grafana\Foundation\Dashboardv2beta1\LibraryPanelKind:
                $disjunctionarg1 = \Grafana\Foundation\Dashboardv2beta1\LibraryPanelConverter::convert($arg1);
                $tmpelementsarg1 = $disjunctionarg1;
                break;
            default:
                throw new \ValueError('disjunction branch not handled');
        }
            $arg0 .= "\t".var_export($key, true)." => $tmpelementsarg1,";
        }
        $arg0 .= "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            
    
        {
    $buffer = 'layout(';
        $arg0 = \Grafana\Foundation\Dashboardv2beta1\NotebookLayoutKindConverter::convert($input->layout);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    

        return \implode("\n\t->", $calls);
    }
}

