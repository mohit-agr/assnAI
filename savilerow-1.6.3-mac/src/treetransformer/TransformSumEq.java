package savilerow.treetransformer;
/*

    Savile Row http://savilerow.cs.st-andrews.ac.uk/
    Copyright (C) 2014, 2015 Peter Nightingale
    
    This file is part of Savile Row.
    
    Savile Row is free software: you can redistribute it and/or modify
    it under the terms of the GNU General Public License as published by
    the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.
    
    Savile Row is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU General Public License for more details.
    
    You should have received a copy of the GNU General Public License
    along with Savile Row.  If not, see <http://www.gnu.org/licenses/>.

*/


import savilerow.expression.*;
import savilerow.model.*;

import java.util.ArrayList;
import java.util.HashMap;

//  Transform a tovariable(sum, aux) into and(sumleq, sumgeq) ahead of Minion output. 

public class TransformSumEq extends TreeTransformerBottomUpNoWrapper
{
    public TransformSumEq() { super(null); }
    protected NodeReplacement processNode(ASTNode curnode)
	{
	    if(curnode instanceof ToVariable && curnode.getChild(0) instanceof WeightedSum)
        {
            if(curnode.getChild(0).numChildren()==1 && ((WeightedSum)curnode.getChild(0)).getWeight(0)==-1) {
                // Special case, need to output a minuseq. Leave the sum intact. 
                return null;
            }
            else {
                return new NodeReplacement(new And(
                    new LessEqual(curnode.getChild(0), curnode.getChild(1)), 
                    new LessEqual(curnode.getChild(1), curnode.getChild(0))
                    ));
            }
        }
        return null;
    }
}

