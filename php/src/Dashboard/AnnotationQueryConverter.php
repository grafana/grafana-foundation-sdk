<?php

namespace Grafana\Foundation\Dashboard;

final class AnnotationQueryConverter
{
    public static function convert(\Grafana\Foundation\Dashboard\AnnotationQuery $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboard\AnnotationQueryBuilder())',
        ];
            if ($input->name !== "") {
    
        
    $buffer = 'name(';
        $arg0 =\var_export($input->name, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            
    
        {
    $buffer = 'datasource(';
        $arg0 ='(new \Grafana\Foundation\Dashboard\DataSourceRef('.(($input->datasource->type !== null) ? 'type: '.\var_export($input->datasource->type, true).', ' : '').''.(($input->datasource->uid !== null) ? 'uid: '.\var_export($input->datasource->uid, true).', ' : '').'))';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            if ($input->enable !== true) {
    
        
    $buffer = 'enable(';
        $arg0 =\var_export($input->enable, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->hide !== null && $input->hide !== false) {
    
        
    $buffer = 'hide(';
        $arg0 =\var_export($input->hide, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->iconColor !== "") {
    
        
    $buffer = 'iconColor(';
        $arg0 =\var_export($input->iconColor, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->filter !== null) {
    
        
    $buffer = 'filter(';
        $arg0 = \Grafana\Foundation\Dashboard\AnnotationPanelFilterConverter::convert($input->filter);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->target !== null) {
    
        
    $buffer = 'target(';
        $arg0 = \Grafana\Foundation\Dashboard\AnnotationTargetConverter::convert($input->target);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->type !== null && $input->type !== "") {
    
        
    $buffer = 'type(';
        $arg0 =\var_export($input->type, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->builtIn !== null && $input->builtIn !== 0) {
    
        
    $buffer = 'builtIn(';
        $arg0 =\var_export($input->builtIn, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->expr !== null && $input->expr !== "") {
    
        
    $buffer = 'expr(';
        $arg0 =\var_export($input->expr, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

