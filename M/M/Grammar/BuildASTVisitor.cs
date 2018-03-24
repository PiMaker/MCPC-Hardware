// File: BuildASTVisitor.cs
// Created: 23.03.2018
// 
// See <summary> tags for more information.

using System;
using System.Runtime.InteropServices;
using Antlr4.Runtime.Misc;
using M.AST;

namespace M.Grammar
{
    internal class BuildASTVisitor : MBaseVisitor<ASTNode>
    {
        public override ASTNode VisitAssignmentExpression([NotNull] MParser.AssignmentExpressionContext context)
        {
            Console.WriteLine("AssignmentExp: " + context.GetText());
            return base.VisitAssignmentExpression(context);
        }

        public override ASTNode VisitAssignmentStatement([NotNull] MParser.AssignmentStatementContext context)
        {
            Console.WriteLine("Assignment: " + context.GetText());
            return base.VisitAssignmentStatement(context);
        }

        public override ASTNode VisitCalcExpression([NotNull] MParser.CalcExpressionContext context)
        {
            Console.WriteLine("Calc: " + context.GetText());
            context.children.Clear();
            return base.VisitCalcExpression(context);
        }

        public override ASTNode VisitCallStatement([NotNull] MParser.CallStatementContext context)
        {
            Console.WriteLine("Function Call: " + context.GetText());
            return base.VisitCallStatement(context);
        }

        public override ASTNode VisitCompilationUnit([NotNull] MParser.CompilationUnitContext context)
        {
            Console.WriteLine("Base unit");
            return base.VisitCompilationUnit(context);
        }

        public override ASTNode VisitCompoundStatement([NotNull] MParser.CompoundStatementContext context)
        {
            return base.VisitCompoundStatement(context);
        }

        public override ASTNode VisitForCondition([NotNull] MParser.ForConditionContext context)
        {
            return base.VisitForCondition(context);
        }

        public override ASTNode VisitFunctionDefinition([NotNull] MParser.FunctionDefinitionContext context)
        {
            Console.WriteLine("Function definition: Identifier=" + context.Identifier().GetText());
            return base.VisitFunctionDefinition(context);
        }

        public override ASTNode VisitInitializerList([NotNull] MParser.InitializerListContext context)
        {
            return base.VisitInitializerList(context);
        }

        public override ASTNode VisitIterationStatement([NotNull] MParser.IterationStatementContext context)
        {
            return base.VisitIterationStatement(context);
        }

        public override ASTNode VisitJumpStatement([NotNull] MParser.JumpStatementContext context)
        {
            return base.VisitJumpStatement(context);
        }

        public override ASTNode VisitLabeledStatement([NotNull] MParser.LabeledStatementContext context)
        {
            return base.VisitLabeledStatement(context);
        }

        public override ASTNode VisitNestedParenthesesBlock([NotNull] MParser.NestedParenthesesBlockContext context)
        {
            return base.VisitNestedParenthesesBlock(context);
        }

        public override ASTNode VisitParameterCallList([NotNull] MParser.ParameterCallListContext context)
        {
            return base.VisitParameterCallList(context);
        }

        public override ASTNode VisitParameterDeclaration([NotNull] MParser.ParameterDeclarationContext context)
        {
            return base.VisitParameterDeclaration(context);
        }

        public override ASTNode VisitParameterDeclarationList([NotNull] MParser.ParameterDeclarationListContext context)
        {
            return base.VisitParameterDeclarationList(context);
        }

        public override ASTNode VisitParamterPassList([NotNull] MParser.ParamterPassListContext context)
        {
            return base.VisitParamterPassList(context);
        }

        public override ASTNode VisitSelectionStatement([NotNull] MParser.SelectionStatementContext context)
        {
            return base.VisitSelectionStatement(context);
        }

        public override ASTNode VisitStatement([NotNull] MParser.StatementContext context)
        {
            return base.VisitStatement(context);
        }

        public override ASTNode VisitTopLevelAssignmentExpression([NotNull] MParser.TopLevelAssignmentExpressionContext context)
        {
            return base.VisitTopLevelAssignmentExpression(context);
        }

        public override ASTNode VisitTopLevelDeclaration([NotNull] MParser.TopLevelDeclarationContext context)
        {
            Console.WriteLine("TLD: " + context.GetText());
            return base.VisitTopLevelDeclaration(context);
        }

        public override ASTNode VisitAsmStatement(MParser.AsmStatementContext context)
        {
            Console.WriteLine("Asm: " + context.GetText());
            return base.VisitAsmStatement(context);
        }

        public override ASTNode VisitPreprocessorDirective(MParser.PreprocessorDirectiveContext context)
        {
            Console.WriteLine("Preprocessor Directive: Include file " + context.Filename());
            return base.VisitPreprocessorDirective(context);
        }
    }
}