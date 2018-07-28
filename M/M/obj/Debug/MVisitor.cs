//------------------------------------------------------------------------------
// <auto-generated>
//     This code was generated by a tool.
//     ANTLR Version: 4.6.5-SNAPSHOT
//
//     Changes to this file may cause incorrect behavior and will be lost if
//     the code is regenerated.
// </auto-generated>
//------------------------------------------------------------------------------

// Generated from E:\Repos\MCPC\M\M\Grammar\M.g4 by ANTLR 4.6.5-SNAPSHOT

// Unreachable code detected
#pragma warning disable 0162
// The variable '...' is assigned but its value is never used
#pragma warning disable 0219
// Missing XML comment for publicly visible type or member '...'
#pragma warning disable 1591
// Ambiguous reference in cref attribute
#pragma warning disable 419

namespace M {
using Antlr4.Runtime.Misc;
using Antlr4.Runtime.Tree;
using IToken = Antlr4.Runtime.IToken;

/// <summary>
/// This interface defines a complete generic visitor for a parse tree produced
/// by <see cref="MParser"/>.
/// </summary>
/// <typeparam name="Result">The return type of the visit operation.</typeparam>
[System.CodeDom.Compiler.GeneratedCode("ANTLR", "4.6.5-SNAPSHOT")]
[System.CLSCompliant(false)]
public interface IMVisitor<Result> : IParseTreeVisitor<Result> {
	/// <summary>
	/// Visit a parse tree produced by <see cref="MParser.calcExpression"/>.
	/// </summary>
	/// <param name="context">The parse tree.</param>
	/// <return>The visitor result.</return>
	Result VisitCalcExpression([NotNull] MParser.CalcExpressionContext context);

	/// <summary>
	/// Visit a parse tree produced by <see cref="MParser.calcOperator"/>.
	/// </summary>
	/// <param name="context">The parse tree.</param>
	/// <return>The visitor result.</return>
	Result VisitCalcOperator([NotNull] MParser.CalcOperatorContext context);

	/// <summary>
	/// Visit a parse tree produced by <see cref="MParser.unaryCalcOperator"/>.
	/// </summary>
	/// <param name="context">The parse tree.</param>
	/// <return>The visitor result.</return>
	Result VisitUnaryCalcOperator([NotNull] MParser.UnaryCalcOperatorContext context);

	/// <summary>
	/// Visit a parse tree produced by <see cref="MParser.initializerList"/>.
	/// </summary>
	/// <param name="context">The parse tree.</param>
	/// <return>The visitor result.</return>
	Result VisitInitializerList([NotNull] MParser.InitializerListContext context);

	/// <summary>
	/// Visit a parse tree produced by <see cref="MParser.assignmentExpression"/>.
	/// </summary>
	/// <param name="context">The parse tree.</param>
	/// <return>The visitor result.</return>
	Result VisitAssignmentExpression([NotNull] MParser.AssignmentExpressionContext context);

	/// <summary>
	/// Visit a parse tree produced by <see cref="MParser.topLevelAssignmentExpression"/>.
	/// </summary>
	/// <param name="context">The parse tree.</param>
	/// <return>The visitor result.</return>
	Result VisitTopLevelAssignmentExpression([NotNull] MParser.TopLevelAssignmentExpressionContext context);

	/// <summary>
	/// Visit a parse tree produced by <see cref="MParser.typeSpecifier"/>.
	/// </summary>
	/// <param name="context">The parse tree.</param>
	/// <return>The visitor result.</return>
	Result VisitTypeSpecifier([NotNull] MParser.TypeSpecifierContext context);

	/// <summary>
	/// Visit a parse tree produced by <see cref="MParser.nestedParenthesesBlock"/>.
	/// </summary>
	/// <param name="context">The parse tree.</param>
	/// <return>The visitor result.</return>
	Result VisitNestedParenthesesBlock([NotNull] MParser.NestedParenthesesBlockContext context);

	/// <summary>
	/// Visit a parse tree produced by <see cref="MParser.parameterDeclarationList"/>.
	/// </summary>
	/// <param name="context">The parse tree.</param>
	/// <return>The visitor result.</return>
	Result VisitParameterDeclarationList([NotNull] MParser.ParameterDeclarationListContext context);

	/// <summary>
	/// Visit a parse tree produced by <see cref="MParser.parameterDeclaration"/>.
	/// </summary>
	/// <param name="context">The parse tree.</param>
	/// <return>The visitor result.</return>
	Result VisitParameterDeclaration([NotNull] MParser.ParameterDeclarationContext context);

	/// <summary>
	/// Visit a parse tree produced by <see cref="MParser.paramterPassList"/>.
	/// </summary>
	/// <param name="context">The parse tree.</param>
	/// <return>The visitor result.</return>
	Result VisitParamterPassList([NotNull] MParser.ParamterPassListContext context);

	/// <summary>
	/// Visit a parse tree produced by <see cref="MParser.parameterCallList"/>.
	/// </summary>
	/// <param name="context">The parse tree.</param>
	/// <return>The visitor result.</return>
	Result VisitParameterCallList([NotNull] MParser.ParameterCallListContext context);

	/// <summary>
	/// Visit a parse tree produced by <see cref="MParser.statement"/>.
	/// </summary>
	/// <param name="context">The parse tree.</param>
	/// <return>The visitor result.</return>
	Result VisitStatement([NotNull] MParser.StatementContext context);

	/// <summary>
	/// Visit a parse tree produced by <see cref="MParser.shiftStatement"/>.
	/// </summary>
	/// <param name="context">The parse tree.</param>
	/// <return>The visitor result.</return>
	Result VisitShiftStatement([NotNull] MParser.ShiftStatementContext context);

	/// <summary>
	/// Visit a parse tree produced by <see cref="MParser.assignmentStatement"/>.
	/// </summary>
	/// <param name="context">The parse tree.</param>
	/// <return>The visitor result.</return>
	Result VisitAssignmentStatement([NotNull] MParser.AssignmentStatementContext context);

	/// <summary>
	/// Visit a parse tree produced by <see cref="MParser.callStatement"/>.
	/// </summary>
	/// <param name="context">The parse tree.</param>
	/// <return>The visitor result.</return>
	Result VisitCallStatement([NotNull] MParser.CallStatementContext context);

	/// <summary>
	/// Visit a parse tree produced by <see cref="MParser.labeledStatement"/>.
	/// </summary>
	/// <param name="context">The parse tree.</param>
	/// <return>The visitor result.</return>
	Result VisitLabeledStatement([NotNull] MParser.LabeledStatementContext context);

	/// <summary>
	/// Visit a parse tree produced by <see cref="MParser.compoundStatement"/>.
	/// </summary>
	/// <param name="context">The parse tree.</param>
	/// <return>The visitor result.</return>
	Result VisitCompoundStatement([NotNull] MParser.CompoundStatementContext context);

	/// <summary>
	/// Visit a parse tree produced by <see cref="MParser.statementList"/>.
	/// </summary>
	/// <param name="context">The parse tree.</param>
	/// <return>The visitor result.</return>
	Result VisitStatementList([NotNull] MParser.StatementListContext context);

	/// <summary>
	/// Visit a parse tree produced by <see cref="MParser.selectionStatement"/>.
	/// </summary>
	/// <param name="context">The parse tree.</param>
	/// <return>The visitor result.</return>
	Result VisitSelectionStatement([NotNull] MParser.SelectionStatementContext context);

	/// <summary>
	/// Visit a parse tree produced by <see cref="MParser.iterationStatement"/>.
	/// </summary>
	/// <param name="context">The parse tree.</param>
	/// <return>The visitor result.</return>
	Result VisitIterationStatement([NotNull] MParser.IterationStatementContext context);

	/// <summary>
	/// Visit a parse tree produced by <see cref="MParser.forCondition"/>.
	/// </summary>
	/// <param name="context">The parse tree.</param>
	/// <return>The visitor result.</return>
	Result VisitForCondition([NotNull] MParser.ForConditionContext context);

	/// <summary>
	/// Visit a parse tree produced by <see cref="MParser.jumpStatement"/>.
	/// </summary>
	/// <param name="context">The parse tree.</param>
	/// <return>The visitor result.</return>
	Result VisitJumpStatement([NotNull] MParser.JumpStatementContext context);

	/// <summary>
	/// Visit a parse tree produced by <see cref="MParser.compilationUnit"/>.
	/// </summary>
	/// <param name="context">The parse tree.</param>
	/// <return>The visitor result.</return>
	Result VisitCompilationUnit([NotNull] MParser.CompilationUnitContext context);

	/// <summary>
	/// Visit a parse tree produced by <see cref="MParser.translationUnit"/>.
	/// </summary>
	/// <param name="context">The parse tree.</param>
	/// <return>The visitor result.</return>
	Result VisitTranslationUnit([NotNull] MParser.TranslationUnitContext context);

	/// <summary>
	/// Visit a parse tree produced by <see cref="MParser.topLevelDeclaration"/>.
	/// </summary>
	/// <param name="context">The parse tree.</param>
	/// <return>The visitor result.</return>
	Result VisitTopLevelDeclaration([NotNull] MParser.TopLevelDeclarationContext context);

	/// <summary>
	/// Visit a parse tree produced by <see cref="MParser.functionDefinition"/>.
	/// </summary>
	/// <param name="context">The parse tree.</param>
	/// <return>The visitor result.</return>
	Result VisitFunctionDefinition([NotNull] MParser.FunctionDefinitionContext context);

	/// <summary>
	/// Visit a parse tree produced by <see cref="MParser.asmStatement"/>.
	/// </summary>
	/// <param name="context">The parse tree.</param>
	/// <return>The visitor result.</return>
	Result VisitAsmStatement([NotNull] MParser.AsmStatementContext context);

	/// <summary>
	/// Visit a parse tree produced by <see cref="MParser.preprocessorDirective"/>.
	/// </summary>
	/// <param name="context">The parse tree.</param>
	/// <return>The visitor result.</return>
	Result VisitPreprocessorDirective([NotNull] MParser.PreprocessorDirectiveContext context);
}
} // namespace M
