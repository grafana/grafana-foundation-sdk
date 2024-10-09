<?php

namespace Grafana\Foundation\Grafanapyroscope;

final class DataqueryConverter
{
    public static function convert(\Grafana\Foundation\Cog\Dataquery $input): string
    {
        assert($input instanceof \Grafana\Foundation\Grafanapyroscope\Dataquery);
        $calls = [
            '(new \Grafana\Foundation\Grafanapyroscope\DataqueryBuilder())',
        ];
            if ($input->labelSelector !== "" && $input->labelSelector !== "{}") {
    
        
    $buffer = 'labelSelector(';
        $arg0 =\var_export($input->labelSelector, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->profileTypeId !== "") {
    
        
    $buffer = 'profileTypeId(';
        $arg0 =\var_export($input->profileTypeId, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if (count($input->groupBy) >= 1) {
    
        
    $buffer = 'groupBy(';
        $tmparg0 = [];
        foreach ($input->groupBy as $arg1) {
        $tmpgroupByarg1 =\var_export($arg1, true);
        $tmparg0[] = $tmpgroupByarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->maxNodes !== null) {
    
        
    $buffer = 'maxNodes(';
        $arg0 =\var_export($input->maxNodes, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->refId !== "") {
    
        
    $buffer = 'refId(';
        $arg0 =\var_export($input->refId, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->hide !== null) {
    
        
    $buffer = 'hide(';
        $arg0 =\var_export($input->hide, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->queryType !== null && $input->queryType !== "") {
    
        
    $buffer = 'queryType(';
        $arg0 =\var_export($input->queryType, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->datasource !== null) {
    
        
    $buffer = 'datasource(';
        $arg0 ='(new \Grafana\Foundation\Dashboard\DataSourceRef('.(($input->datasource->type !== null) ? 'type: '.\var_export($input->datasource->type, true).', ' : '').''.(($input->datasource->uid !== null) ? 'uid: '.\var_export($input->datasource->uid, true).', ' : '').'))';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

