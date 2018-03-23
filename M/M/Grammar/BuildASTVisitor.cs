using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using Antlr4.Runtime.Tree;

namespace M.Grammar
{
    internal class BuildASTVisitor : MBaseVisitor<AbstractAssemblerNode>
    {
        public override AbstractAssemblerNode VisitCompilationUnit(MParser.CompilationUnitContext context)
        {
            return base.VisitCompilationUnit(context);
        }

        public override AbstractAssemblerNode VisitCallStatement(MParser.CallStatementContext context)
        {
            Console.WriteLine(context.parameterCallList().GetText());
            return base.VisitCallStatement(context);
        }
    }
}
