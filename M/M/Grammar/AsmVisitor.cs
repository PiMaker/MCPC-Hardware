// File: BuildASTVisitor.cs
// Created: 23.03.2018
// 
// See <summary> tags for more information.

using System;
using System.Text;
using Antlr4.Runtime.Misc;

namespace M.Grammar
{
    internal class AsmVisitor : MBaseVisitor<string>
    {
        public StringBuilder Asm { get; private set; }

        public override string VisitAssignmentExpression([NotNull] MParser.AssignmentExpressionContext context)
        {
            Console.WriteLine("AssignmentExp: " + context.GetText());
            return base.VisitAssignmentExpression(context);
        }

        public override string VisitAssignmentStatement([NotNull] MParser.AssignmentStatementContext context)
        {
            Console.WriteLine("Assignment: " + context.GetText());
            return base.VisitAssignmentStatement(context);
        }

        public override string VisitCalcExpression([NotNull] MParser.CalcExpressionContext context)
        {
            Console.WriteLine("Calc: " + context.GetText());
            context.children.Clear();
            return base.VisitCalcExpression(context);
        }

        public override string VisitCallStatement([NotNull] MParser.CallStatementContext context)
        {
            Console.WriteLine("Function Call: " + context.GetText());
            return base.VisitCallStatement(context);
        }

        public override string VisitCompilationUnit([NotNull] MParser.CompilationUnitContext context)
        {
            Console.WriteLine("Base unit");
            return base.VisitCompilationUnit(context);
        }

        public override string VisitCompoundStatement([NotNull] MParser.CompoundStatementContext context)
        {
            return base.VisitCompoundStatement(context);
        }

        public override string VisitForCondition([NotNull] MParser.ForConditionContext context)
        {
            return base.VisitForCondition(context);
        }

        public override string VisitFunctionDefinition([NotNull] MParser.FunctionDefinitionContext context)
        {
            Console.WriteLine("Function definition: Identifier=" + context.Identifier().GetText());
            return base.VisitFunctionDefinition(context);
        }

        public override string VisitInitializerList([NotNull] MParser.InitializerListContext context)
        {
            return base.VisitInitializerList(context);
        }

        public override string VisitIterationStatement([NotNull] MParser.IterationStatementContext context)
        {
            return base.VisitIterationStatement(context);
        }

        public override string VisitJumpStatement([NotNull] MParser.JumpStatementContext context)
        {
            return base.VisitJumpStatement(context);
        }

        public override string VisitLabeledStatement([NotNull] MParser.LabeledStatementContext context)
        {
            return base.VisitLabeledStatement(context);
        }

        public override string VisitNestedParenthesesBlock([NotNull] MParser.NestedParenthesesBlockContext context)
        {
            return base.VisitNestedParenthesesBlock(context);
        }

        public override string VisitParameterCallList([NotNull] MParser.ParameterCallListContext context)
        {
            return base.VisitParameterCallList(context);
        }

        public override string VisitParameterDeclaration([NotNull] MParser.ParameterDeclarationContext context)
        {
            return base.VisitParameterDeclaration(context);
        }

        public override string VisitParameterDeclarationList([NotNull] MParser.ParameterDeclarationListContext context)
        {
            return base.VisitParameterDeclarationList(context);
        }

        public override string VisitParamterPassList([NotNull] MParser.ParamterPassListContext context)
        {
            return base.VisitParamterPassList(context);
        }

        public override string VisitSelectionStatement([NotNull] MParser.SelectionStatementContext context)
        {
            return base.VisitSelectionStatement(context);
        }

        public override string VisitStatement([NotNull] MParser.StatementContext context)
        {
            return base.VisitStatement(context);
        }

        public override string VisitTopLevelAssignmentExpression([NotNull] MParser.TopLevelAssignmentExpressionContext context)
        {
            return base.VisitTopLevelAssignmentExpression(context);
        }

        public override string VisitTopLevelDeclaration([NotNull] MParser.TopLevelDeclarationContext context)
        {
            Console.WriteLine("TLD: " + context.GetText());
            return base.VisitTopLevelDeclaration(context);
        }

        public override string VisitAsmStatement(MParser.AsmStatementContext context)
        {
            Console.WriteLine("Asm: " + context.GetText());
            return base.VisitAsmStatement(context);
        }

        public override string VisitPreprocessorDirective(MParser.PreprocessorDirectiveContext context)
        {
            Console.WriteLine("Preprocessor Directive: Include file " + context.Filename());
            return base.VisitPreprocessorDirective(context);
        }
    }
}