<?php

namespace Grafana\Foundation\Dashboardv2beta1;

final class DashboardConverter
{
    public static function convert(\Grafana\Foundation\Dashboardv2beta1\Dashboard $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboardv2beta1\DashboardBuilder('.\var_export($input->title, true).'))',
        ];
            if (count($input->annotations) >= 1) {
    
        
    $buffer = 'annotations(';
        $tmparg0 = [];
        foreach ($input->annotations as $arg1) {
        $tmpannotationsarg1 = \Grafana\Foundation\Dashboardv2beta1\AnnotationQueryConverter::convert($arg1);
        $tmparg0[] = $tmpannotationsarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            
    
        {
    $buffer = 'cursorSync(';
        $arg0 ='\Grafana\Foundation\Dashboardv2beta1\DashboardCursorSync::fromValue("'.$input->cursorSync.'")';
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
            if ($input->editable !== null && $input->editable !== true) {
    
        
    $buffer = 'editable(';
        $arg0 =\var_export($input->editable, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            
    
        {
    $buffer = 'elements(';
        $arg0 = "[";
        foreach ($input->elements as $key => $arg1) {
            switch (true) {
            case $arg1 instanceof \Grafana\Foundation\Dashboardv2beta1\PanelKind:
                $disjunctionarg1 = \Grafana\Foundation\Dashboardv2beta1\PanelConverter::convert($arg1);
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
    
    
            
    foreach ($input->elements as $key => $value) {
        {
    $buffer = 'element(';
        $arg0 =\var_export($key, true);
        $buffer .= $arg0;
        $buffer .= ', ';
        switch (true) {
            case $value instanceof \Grafana\Foundation\Dashboardv2beta1\PanelKind:
                $disjunctionvalue = \Grafana\Foundation\Dashboardv2beta1\PanelConverter::convert($value);
                $arg1 = $disjunctionvalue;
                break;
            case $value instanceof \Grafana\Foundation\Dashboardv2beta1\LibraryPanelKind:
                $disjunctionvalue = \Grafana\Foundation\Dashboardv2beta1\LibraryPanelConverter::convert($value);
                $arg1 = $disjunctionvalue;
                break;
            default:
                throw new \ValueError('disjunction branch not handled');
        }
        $buffer .= $arg1;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    }
    
            
    
        {
    $buffer = 'layout(';
        switch (true) {
            case $input->layout instanceof \Grafana\Foundation\Dashboardv2beta1\GridLayoutKind:
                $disjunctionlayout = \Grafana\Foundation\Dashboardv2beta1\GridConverter::convert($input->layout);
                $arg0 = $disjunctionlayout;
                break;
            case $input->layout instanceof \Grafana\Foundation\Dashboardv2beta1\RowsLayoutKind:
                $disjunctionlayout = \Grafana\Foundation\Dashboardv2beta1\RowsConverter::convert($input->layout);
                $arg0 = $disjunctionlayout;
                break;
            case $input->layout instanceof \Grafana\Foundation\Dashboardv2beta1\AutoGridLayoutKind:
                $disjunctionlayout = \Grafana\Foundation\Dashboardv2beta1\AutoGridConverter::convert($input->layout);
                $arg0 = $disjunctionlayout;
                break;
            case $input->layout instanceof \Grafana\Foundation\Dashboardv2beta1\TabsLayoutKind:
                $disjunctionlayout = \Grafana\Foundation\Dashboardv2beta1\TabsConverter::convert($input->layout);
                $arg0 = $disjunctionlayout;
                break;
            default:
                throw new \ValueError('disjunction branch not handled');
        }
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            if (count($input->links) >= 1) {
    
        
    $buffer = 'links(';
        $tmparg0 = [];
        foreach ($input->links as $arg1) {
        $tmplinksarg1 = \Grafana\Foundation\Dashboardv2beta1\DashboardLinkConverter::convert($arg1);
        $tmparg0[] = $tmplinksarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->liveNow !== null) {
    
        
    $buffer = 'liveNow(';
        $arg0 =\var_export($input->liveNow, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->preload !== false) {
    
        
    $buffer = 'preload(';
        $arg0 =\var_export($input->preload, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->revision !== null) {
    
        
    $buffer = 'revision(';
        $arg0 =\var_export($input->revision, true);
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
    
    
            if ($input->title !== "") {
    
        
    $buffer = 'title(';
        $arg0 =\var_export($input->title, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if (count($input->variables) >= 1) {
    
        
    $buffer = 'variables(';
        $tmparg0 = [];
        foreach ($input->variables as $arg1) {
        switch (true) {
            case $arg1 instanceof \Grafana\Foundation\Dashboardv2beta1\QueryVariableKind:
                $disjunctionarg1 = \Grafana\Foundation\Dashboardv2beta1\QueryVariableConverter::convert($arg1);
                $tmpvariablesarg1 = $disjunctionarg1;
                break;
            case $arg1 instanceof \Grafana\Foundation\Dashboardv2beta1\TextVariableKind:
                $disjunctionarg1 = \Grafana\Foundation\Dashboardv2beta1\TextVariableConverter::convert($arg1);
                $tmpvariablesarg1 = $disjunctionarg1;
                break;
            case $arg1 instanceof \Grafana\Foundation\Dashboardv2beta1\ConstantVariableKind:
                $disjunctionarg1 = \Grafana\Foundation\Dashboardv2beta1\ConstantVariableConverter::convert($arg1);
                $tmpvariablesarg1 = $disjunctionarg1;
                break;
            case $arg1 instanceof \Grafana\Foundation\Dashboardv2beta1\DatasourceVariableKind:
                $disjunctionarg1 = \Grafana\Foundation\Dashboardv2beta1\DatasourceVariableConverter::convert($arg1);
                $tmpvariablesarg1 = $disjunctionarg1;
                break;
            case $arg1 instanceof \Grafana\Foundation\Dashboardv2beta1\IntervalVariableKind:
                $disjunctionarg1 = \Grafana\Foundation\Dashboardv2beta1\IntervalVariableConverter::convert($arg1);
                $tmpvariablesarg1 = $disjunctionarg1;
                break;
            case $arg1 instanceof \Grafana\Foundation\Dashboardv2beta1\CustomVariableKind:
                $disjunctionarg1 = \Grafana\Foundation\Dashboardv2beta1\CustomVariableConverter::convert($arg1);
                $tmpvariablesarg1 = $disjunctionarg1;
                break;
            case $arg1 instanceof \Grafana\Foundation\Dashboardv2beta1\GroupByVariableKind:
                $disjunctionarg1 = \Grafana\Foundation\Dashboardv2beta1\GroupByVariableConverter::convert($arg1);
                $tmpvariablesarg1 = $disjunctionarg1;
                break;
            case $arg1 instanceof \Grafana\Foundation\Dashboardv2beta1\AdhocVariableKind:
                $disjunctionarg1 = \Grafana\Foundation\Dashboardv2beta1\AdhocVariableConverter::convert($arg1);
                $tmpvariablesarg1 = $disjunctionarg1;
                break;
            case $arg1 instanceof \Grafana\Foundation\Dashboardv2beta1\SwitchVariableKind:
                $disjunctionarg1 = \Grafana\Foundation\Dashboardv2beta1\SwitchVariableConverter::convert($arg1);
                $tmpvariablesarg1 = $disjunctionarg1;
                break;
            default:
                throw new \ValueError('disjunction branch not handled');
        }
        $tmparg0[] = $tmpvariablesarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

