// Generated from e:\Repos\MCPC\M\M\Grammar\M.g4 by ANTLR 4.7.1
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class MParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.7.1", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		Break=1, Char=2, Continue=3, Do=4, Else=5, For=6, Goto=7, If=8, Inline=9, 
		Int=10, Return=11, Void=12, While=13, LeftParen=14, RightParen=15, LeftBracket=16, 
		RightBracket=17, LeftBrace=18, RightBrace=19, Less=20, LessEqual=21, Greater=22, 
		GreaterEqual=23, LeftShift=24, RightShift=25, Plus=26, PlusPlus=27, Minus=28, 
		MinusMinus=29, Star=30, Div=31, Mod=32, And=33, Or=34, AndAnd=35, OrOr=36, 
		Xor=37, Not=38, Tilde=39, Question=40, Colon=41, Semi=42, Comma=43, Assign=44, 
		Equal=45, NotEqual=46, Dot=47, Identifier=48, Constant=49, DigitSequence=50, 
		StringLiteral=51, ComplexDefine=52, AsmBlock=53, PreprocessorDirective=54, 
		Whitespace=55, Newline=56, BlockComment=57, LineComment=58;
	public static final int
		RULE_calcExpression = 0, RULE_calcOperator = 1, RULE_unaryCalcOperator = 2, 
		RULE_initializerList = 3, RULE_assignmentExpression = 4, RULE_topLevelAssignmentExpression = 5, 
		RULE_typeSpecifier = 6, RULE_nestedParenthesesBlock = 7, RULE_pointer = 8, 
		RULE_parameterDeclarationList = 9, RULE_parameterDeclaration = 10, RULE_paramterPassList = 11, 
		RULE_parameterCallList = 12, RULE_statement = 13, RULE_assignmentStatement = 14, 
		RULE_callStatement = 15, RULE_labeledStatement = 16, RULE_compoundStatement = 17, 
		RULE_statementList = 18, RULE_selectionStatement = 19, RULE_iterationStatement = 20, 
		RULE_forCondition = 21, RULE_jumpStatement = 22, RULE_compilationUnit = 23, 
		RULE_translationUnit = 24, RULE_topLevelDeclaration = 25, RULE_functionDefinition = 26;
	public static final String[] ruleNames = {
		"calcExpression", "calcOperator", "unaryCalcOperator", "initializerList", 
		"assignmentExpression", "topLevelAssignmentExpression", "typeSpecifier", 
		"nestedParenthesesBlock", "pointer", "parameterDeclarationList", "parameterDeclaration", 
		"paramterPassList", "parameterCallList", "statement", "assignmentStatement", 
		"callStatement", "labeledStatement", "compoundStatement", "statementList", 
		"selectionStatement", "iterationStatement", "forCondition", "jumpStatement", 
		"compilationUnit", "translationUnit", "topLevelDeclaration", "functionDefinition"
	};

	private static final String[] _LITERAL_NAMES = {
		null, "'break'", "'char'", "'continue'", "'do'", "'else'", "'for'", "'goto'", 
		"'if'", "'inline'", "'int'", "'return'", "'void'", "'while'", "'('", "')'", 
		"'['", "']'", "'{'", "'}'", "'<'", "'<='", "'>'", "'>='", "'<<'", "'>>'", 
		"'+'", "'++'", "'-'", "'--'", "'*'", "'/'", "'%'", "'&'", "'|'", "'&&'", 
		"'||'", "'^'", "'!'", "'~'", "'?'", "':'", "';'", "','", "'='", "'=='", 
		"'!='", "'.'"
	};
	private static final String[] _SYMBOLIC_NAMES = {
		null, "Break", "Char", "Continue", "Do", "Else", "For", "Goto", "If", 
		"Inline", "Int", "Return", "Void", "While", "LeftParen", "RightParen", 
		"LeftBracket", "RightBracket", "LeftBrace", "RightBrace", "Less", "LessEqual", 
		"Greater", "GreaterEqual", "LeftShift", "RightShift", "Plus", "PlusPlus", 
		"Minus", "MinusMinus", "Star", "Div", "Mod", "And", "Or", "AndAnd", "OrOr", 
		"Xor", "Not", "Tilde", "Question", "Colon", "Semi", "Comma", "Assign", 
		"Equal", "NotEqual", "Dot", "Identifier", "Constant", "DigitSequence", 
		"StringLiteral", "ComplexDefine", "AsmBlock", "PreprocessorDirective", 
		"Whitespace", "Newline", "BlockComment", "LineComment"
	};
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}

	@Override
	public String getGrammarFileName() { return "M.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public MParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}
	public static class CalcExpressionContext extends ParserRuleContext {
		public TerminalNode Identifier() { return getToken(MParser.Identifier, 0); }
		public TerminalNode StringLiteral() { return getToken(MParser.StringLiteral, 0); }
		public InitializerListContext initializerList() {
			return getRuleContext(InitializerListContext.class,0);
		}
		public UnaryCalcOperatorContext unaryCalcOperator() {
			return getRuleContext(UnaryCalcOperatorContext.class,0);
		}
		public List<CalcExpressionContext> calcExpression() {
			return getRuleContexts(CalcExpressionContext.class);
		}
		public CalcExpressionContext calcExpression(int i) {
			return getRuleContext(CalcExpressionContext.class,i);
		}
		public TypeSpecifierContext typeSpecifier() {
			return getRuleContext(TypeSpecifierContext.class,0);
		}
		public CalcOperatorContext calcOperator() {
			return getRuleContext(CalcOperatorContext.class,0);
		}
		public CalcExpressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_calcExpression; }
	}

	public final CalcExpressionContext calcExpression() throws RecognitionException {
		return calcExpression(0);
	}

	private CalcExpressionContext calcExpression(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		CalcExpressionContext _localctx = new CalcExpressionContext(_ctx, _parentState);
		CalcExpressionContext _prevctx = _localctx;
		int _startState = 0;
		enterRecursionRule(_localctx, 0, RULE_calcExpression, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(73);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,0,_ctx) ) {
			case 1:
				{
				setState(55);
				match(Identifier);
				}
				break;
			case 2:
				{
				setState(56);
				match(StringLiteral);
				}
				break;
			case 3:
				{
				setState(57);
				match(LeftBrace);
				setState(58);
				initializerList(0);
				setState(59);
				match(RightBrace);
				}
				break;
			case 4:
				{
				setState(61);
				unaryCalcOperator();
				setState(62);
				calcExpression(4);
				}
				break;
			case 5:
				{
				setState(64);
				match(LeftParen);
				setState(65);
				typeSpecifier(0);
				setState(66);
				match(RightParen);
				setState(67);
				calcExpression(2);
				}
				break;
			case 6:
				{
				setState(69);
				match(LeftParen);
				setState(70);
				calcExpression(0);
				setState(71);
				match(RightParen);
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(83);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,2,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(81);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,1,_ctx) ) {
					case 1:
						{
						_localctx = new CalcExpressionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_calcExpression);
						setState(75);
						if (!(precpred(_ctx, 5))) throw new FailedPredicateException(this, "precpred(_ctx, 5)");
						setState(76);
						calcOperator();
						setState(77);
						calcExpression(6);
						}
						break;
					case 2:
						{
						_localctx = new CalcExpressionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_calcExpression);
						setState(79);
						if (!(precpred(_ctx, 3))) throw new FailedPredicateException(this, "precpred(_ctx, 3)");
						setState(80);
						unaryCalcOperator();
						}
						break;
					}
					} 
				}
				setState(85);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,2,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public static class CalcOperatorContext extends ParserRuleContext {
		public TerminalNode Less() { return getToken(MParser.Less, 0); }
		public TerminalNode LessEqual() { return getToken(MParser.LessEqual, 0); }
		public TerminalNode Greater() { return getToken(MParser.Greater, 0); }
		public TerminalNode GreaterEqual() { return getToken(MParser.GreaterEqual, 0); }
		public TerminalNode LeftShift() { return getToken(MParser.LeftShift, 0); }
		public TerminalNode RightShift() { return getToken(MParser.RightShift, 0); }
		public TerminalNode Plus() { return getToken(MParser.Plus, 0); }
		public TerminalNode Minus() { return getToken(MParser.Minus, 0); }
		public TerminalNode Star() { return getToken(MParser.Star, 0); }
		public TerminalNode Div() { return getToken(MParser.Div, 0); }
		public TerminalNode Mod() { return getToken(MParser.Mod, 0); }
		public TerminalNode And() { return getToken(MParser.And, 0); }
		public TerminalNode Or() { return getToken(MParser.Or, 0); }
		public TerminalNode AndAnd() { return getToken(MParser.AndAnd, 0); }
		public TerminalNode OrOr() { return getToken(MParser.OrOr, 0); }
		public TerminalNode Xor() { return getToken(MParser.Xor, 0); }
		public CalcOperatorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_calcOperator; }
	}

	public final CalcOperatorContext calcOperator() throws RecognitionException {
		CalcOperatorContext _localctx = new CalcOperatorContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_calcOperator);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(86);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << Less) | (1L << LessEqual) | (1L << Greater) | (1L << GreaterEqual) | (1L << LeftShift) | (1L << RightShift) | (1L << Plus) | (1L << Minus) | (1L << Star) | (1L << Div) | (1L << Mod) | (1L << And) | (1L << Or) | (1L << AndAnd) | (1L << OrOr) | (1L << Xor))) != 0)) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class UnaryCalcOperatorContext extends ParserRuleContext {
		public TerminalNode PlusPlus() { return getToken(MParser.PlusPlus, 0); }
		public TerminalNode MinusMinus() { return getToken(MParser.MinusMinus, 0); }
		public TerminalNode Tilde() { return getToken(MParser.Tilde, 0); }
		public TerminalNode Not() { return getToken(MParser.Not, 0); }
		public TerminalNode Star() { return getToken(MParser.Star, 0); }
		public TerminalNode And() { return getToken(MParser.And, 0); }
		public UnaryCalcOperatorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_unaryCalcOperator; }
	}

	public final UnaryCalcOperatorContext unaryCalcOperator() throws RecognitionException {
		UnaryCalcOperatorContext _localctx = new UnaryCalcOperatorContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_unaryCalcOperator);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(88);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << PlusPlus) | (1L << MinusMinus) | (1L << Star) | (1L << And) | (1L << Not) | (1L << Tilde))) != 0)) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class InitializerListContext extends ParserRuleContext {
		public TerminalNode Constant() { return getToken(MParser.Constant, 0); }
		public InitializerListContext initializerList() {
			return getRuleContext(InitializerListContext.class,0);
		}
		public InitializerListContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_initializerList; }
	}

	public final InitializerListContext initializerList() throws RecognitionException {
		return initializerList(0);
	}

	private InitializerListContext initializerList(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		InitializerListContext _localctx = new InitializerListContext(_ctx, _parentState);
		InitializerListContext _prevctx = _localctx;
		int _startState = 6;
		enterRecursionRule(_localctx, 6, RULE_initializerList, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(91);
			match(Constant);
			}
			_ctx.stop = _input.LT(-1);
			setState(98);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,3,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new InitializerListContext(_parentctx, _parentState);
					pushNewRecursionContext(_localctx, _startState, RULE_initializerList);
					setState(93);
					if (!(precpred(_ctx, 1))) throw new FailedPredicateException(this, "precpred(_ctx, 1)");
					setState(94);
					match(Comma);
					setState(95);
					match(Constant);
					}
					} 
				}
				setState(100);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,3,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public static class AssignmentExpressionContext extends ParserRuleContext {
		public TerminalNode Identifier() { return getToken(MParser.Identifier, 0); }
		public CalcExpressionContext calcExpression() {
			return getRuleContext(CalcExpressionContext.class,0);
		}
		public TypeSpecifierContext typeSpecifier() {
			return getRuleContext(TypeSpecifierContext.class,0);
		}
		public AssignmentExpressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_assignmentExpression; }
	}

	public final AssignmentExpressionContext assignmentExpression() throws RecognitionException {
		AssignmentExpressionContext _localctx = new AssignmentExpressionContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_assignmentExpression);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(102);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << Char) | (1L << Int) | (1L << Void))) != 0)) {
				{
				setState(101);
				typeSpecifier(0);
				}
			}

			setState(104);
			match(Identifier);
			setState(105);
			match(Assign);
			setState(106);
			calcExpression(0);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class TopLevelAssignmentExpressionContext extends ParserRuleContext {
		public TypeSpecifierContext typeSpecifier() {
			return getRuleContext(TypeSpecifierContext.class,0);
		}
		public TerminalNode Identifier() { return getToken(MParser.Identifier, 0); }
		public CalcExpressionContext calcExpression() {
			return getRuleContext(CalcExpressionContext.class,0);
		}
		public TopLevelAssignmentExpressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_topLevelAssignmentExpression; }
	}

	public final TopLevelAssignmentExpressionContext topLevelAssignmentExpression() throws RecognitionException {
		TopLevelAssignmentExpressionContext _localctx = new TopLevelAssignmentExpressionContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_topLevelAssignmentExpression);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(108);
			typeSpecifier(0);
			setState(109);
			match(Identifier);
			setState(110);
			match(Assign);
			setState(111);
			calcExpression(0);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class TypeSpecifierContext extends ParserRuleContext {
		public TypeSpecifierContext typeSpecifier() {
			return getRuleContext(TypeSpecifierContext.class,0);
		}
		public PointerContext pointer() {
			return getRuleContext(PointerContext.class,0);
		}
		public TypeSpecifierContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_typeSpecifier; }
	}

	public final TypeSpecifierContext typeSpecifier() throws RecognitionException {
		return typeSpecifier(0);
	}

	private TypeSpecifierContext typeSpecifier(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		TypeSpecifierContext _localctx = new TypeSpecifierContext(_ctx, _parentState);
		TypeSpecifierContext _prevctx = _localctx;
		int _startState = 12;
		enterRecursionRule(_localctx, 12, RULE_typeSpecifier, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(114);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << Char) | (1L << Int) | (1L << Void))) != 0)) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
			_ctx.stop = _input.LT(-1);
			setState(120);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,5,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new TypeSpecifierContext(_parentctx, _parentState);
					pushNewRecursionContext(_localctx, _startState, RULE_typeSpecifier);
					setState(116);
					if (!(precpred(_ctx, 1))) throw new FailedPredicateException(this, "precpred(_ctx, 1)");
					setState(117);
					pointer();
					}
					} 
				}
				setState(122);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,5,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public static class NestedParenthesesBlockContext extends ParserRuleContext {
		public List<NestedParenthesesBlockContext> nestedParenthesesBlock() {
			return getRuleContexts(NestedParenthesesBlockContext.class);
		}
		public NestedParenthesesBlockContext nestedParenthesesBlock(int i) {
			return getRuleContext(NestedParenthesesBlockContext.class,i);
		}
		public NestedParenthesesBlockContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_nestedParenthesesBlock; }
	}

	public final NestedParenthesesBlockContext nestedParenthesesBlock() throws RecognitionException {
		NestedParenthesesBlockContext _localctx = new NestedParenthesesBlockContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_nestedParenthesesBlock);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(130);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << Break) | (1L << Char) | (1L << Continue) | (1L << Do) | (1L << Else) | (1L << For) | (1L << Goto) | (1L << If) | (1L << Inline) | (1L << Int) | (1L << Return) | (1L << Void) | (1L << While) | (1L << LeftParen) | (1L << LeftBracket) | (1L << RightBracket) | (1L << LeftBrace) | (1L << RightBrace) | (1L << Less) | (1L << LessEqual) | (1L << Greater) | (1L << GreaterEqual) | (1L << LeftShift) | (1L << RightShift) | (1L << Plus) | (1L << PlusPlus) | (1L << Minus) | (1L << MinusMinus) | (1L << Star) | (1L << Div) | (1L << Mod) | (1L << And) | (1L << Or) | (1L << AndAnd) | (1L << OrOr) | (1L << Xor) | (1L << Not) | (1L << Tilde) | (1L << Question) | (1L << Colon) | (1L << Semi) | (1L << Comma) | (1L << Assign) | (1L << Equal) | (1L << NotEqual) | (1L << Dot) | (1L << Identifier) | (1L << Constant) | (1L << DigitSequence) | (1L << StringLiteral) | (1L << ComplexDefine) | (1L << AsmBlock) | (1L << PreprocessorDirective) | (1L << Whitespace) | (1L << Newline) | (1L << BlockComment) | (1L << LineComment))) != 0)) {
				{
				setState(128);
				_errHandler.sync(this);
				switch (_input.LA(1)) {
				case Break:
				case Char:
				case Continue:
				case Do:
				case Else:
				case For:
				case Goto:
				case If:
				case Inline:
				case Int:
				case Return:
				case Void:
				case While:
				case LeftBracket:
				case RightBracket:
				case LeftBrace:
				case RightBrace:
				case Less:
				case LessEqual:
				case Greater:
				case GreaterEqual:
				case LeftShift:
				case RightShift:
				case Plus:
				case PlusPlus:
				case Minus:
				case MinusMinus:
				case Star:
				case Div:
				case Mod:
				case And:
				case Or:
				case AndAnd:
				case OrOr:
				case Xor:
				case Not:
				case Tilde:
				case Question:
				case Colon:
				case Semi:
				case Comma:
				case Assign:
				case Equal:
				case NotEqual:
				case Dot:
				case Identifier:
				case Constant:
				case DigitSequence:
				case StringLiteral:
				case ComplexDefine:
				case AsmBlock:
				case PreprocessorDirective:
				case Whitespace:
				case Newline:
				case BlockComment:
				case LineComment:
					{
					setState(123);
					_la = _input.LA(1);
					if ( _la <= 0 || (_la==LeftParen || _la==RightParen) ) {
					_errHandler.recoverInline(this);
					}
					else {
						if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
						_errHandler.reportMatch(this);
						consume();
					}
					}
					break;
				case LeftParen:
					{
					setState(124);
					match(LeftParen);
					setState(125);
					nestedParenthesesBlock();
					setState(126);
					match(RightParen);
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				}
				setState(132);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class PointerContext extends ParserRuleContext {
		public PointerContext pointer() {
			return getRuleContext(PointerContext.class,0);
		}
		public PointerContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_pointer; }
	}

	public final PointerContext pointer() throws RecognitionException {
		PointerContext _localctx = new PointerContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_pointer);
		try {
			setState(136);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,8,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(133);
				match(Star);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(134);
				match(Star);
				setState(135);
				pointer();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ParameterDeclarationListContext extends ParserRuleContext {
		public ParameterDeclarationContext parameterDeclaration() {
			return getRuleContext(ParameterDeclarationContext.class,0);
		}
		public ParameterDeclarationListContext parameterDeclarationList() {
			return getRuleContext(ParameterDeclarationListContext.class,0);
		}
		public ParameterDeclarationListContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_parameterDeclarationList; }
	}

	public final ParameterDeclarationListContext parameterDeclarationList() throws RecognitionException {
		return parameterDeclarationList(0);
	}

	private ParameterDeclarationListContext parameterDeclarationList(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ParameterDeclarationListContext _localctx = new ParameterDeclarationListContext(_ctx, _parentState);
		ParameterDeclarationListContext _prevctx = _localctx;
		int _startState = 18;
		enterRecursionRule(_localctx, 18, RULE_parameterDeclarationList, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(139);
			parameterDeclaration();
			}
			_ctx.stop = _input.LT(-1);
			setState(146);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,9,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new ParameterDeclarationListContext(_parentctx, _parentState);
					pushNewRecursionContext(_localctx, _startState, RULE_parameterDeclarationList);
					setState(141);
					if (!(precpred(_ctx, 1))) throw new FailedPredicateException(this, "precpred(_ctx, 1)");
					setState(142);
					match(Comma);
					setState(143);
					parameterDeclaration();
					}
					} 
				}
				setState(148);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,9,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public static class ParameterDeclarationContext extends ParserRuleContext {
		public TypeSpecifierContext typeSpecifier() {
			return getRuleContext(TypeSpecifierContext.class,0);
		}
		public TerminalNode Identifier() { return getToken(MParser.Identifier, 0); }
		public ParameterDeclarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_parameterDeclaration; }
	}

	public final ParameterDeclarationContext parameterDeclaration() throws RecognitionException {
		ParameterDeclarationContext _localctx = new ParameterDeclarationContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_parameterDeclaration);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(149);
			typeSpecifier(0);
			setState(150);
			match(Identifier);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ParamterPassListContext extends ParserRuleContext {
		public CalcExpressionContext calcExpression() {
			return getRuleContext(CalcExpressionContext.class,0);
		}
		public ParamterPassListContext paramterPassList() {
			return getRuleContext(ParamterPassListContext.class,0);
		}
		public ParamterPassListContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_paramterPassList; }
	}

	public final ParamterPassListContext paramterPassList() throws RecognitionException {
		return paramterPassList(0);
	}

	private ParamterPassListContext paramterPassList(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ParamterPassListContext _localctx = new ParamterPassListContext(_ctx, _parentState);
		ParamterPassListContext _prevctx = _localctx;
		int _startState = 22;
		enterRecursionRule(_localctx, 22, RULE_paramterPassList, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(153);
			calcExpression(0);
			}
			_ctx.stop = _input.LT(-1);
			setState(160);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,10,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new ParamterPassListContext(_parentctx, _parentState);
					pushNewRecursionContext(_localctx, _startState, RULE_paramterPassList);
					setState(155);
					if (!(precpred(_ctx, 1))) throw new FailedPredicateException(this, "precpred(_ctx, 1)");
					setState(156);
					match(Comma);
					setState(157);
					calcExpression(0);
					}
					} 
				}
				setState(162);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,10,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public static class ParameterCallListContext extends ParserRuleContext {
		public CalcExpressionContext calcExpression() {
			return getRuleContext(CalcExpressionContext.class,0);
		}
		public ParameterCallListContext parameterCallList() {
			return getRuleContext(ParameterCallListContext.class,0);
		}
		public ParameterCallListContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_parameterCallList; }
	}

	public final ParameterCallListContext parameterCallList() throws RecognitionException {
		return parameterCallList(0);
	}

	private ParameterCallListContext parameterCallList(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ParameterCallListContext _localctx = new ParameterCallListContext(_ctx, _parentState);
		ParameterCallListContext _prevctx = _localctx;
		int _startState = 24;
		enterRecursionRule(_localctx, 24, RULE_parameterCallList, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(164);
			calcExpression(0);
			}
			_ctx.stop = _input.LT(-1);
			setState(171);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,11,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new ParameterCallListContext(_parentctx, _parentState);
					pushNewRecursionContext(_localctx, _startState, RULE_parameterCallList);
					setState(166);
					if (!(precpred(_ctx, 1))) throw new FailedPredicateException(this, "precpred(_ctx, 1)");
					setState(167);
					match(Comma);
					setState(168);
					calcExpression(0);
					}
					} 
				}
				setState(173);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,11,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public static class StatementContext extends ParserRuleContext {
		public LabeledStatementContext labeledStatement() {
			return getRuleContext(LabeledStatementContext.class,0);
		}
		public CompoundStatementContext compoundStatement() {
			return getRuleContext(CompoundStatementContext.class,0);
		}
		public AssignmentStatementContext assignmentStatement() {
			return getRuleContext(AssignmentStatementContext.class,0);
		}
		public SelectionStatementContext selectionStatement() {
			return getRuleContext(SelectionStatementContext.class,0);
		}
		public IterationStatementContext iterationStatement() {
			return getRuleContext(IterationStatementContext.class,0);
		}
		public CallStatementContext callStatement() {
			return getRuleContext(CallStatementContext.class,0);
		}
		public JumpStatementContext jumpStatement() {
			return getRuleContext(JumpStatementContext.class,0);
		}
		public StatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_statement; }
	}

	public final StatementContext statement() throws RecognitionException {
		StatementContext _localctx = new StatementContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_statement);
		try {
			setState(181);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,12,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(174);
				labeledStatement();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(175);
				compoundStatement();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(176);
				assignmentStatement();
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(177);
				selectionStatement();
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(178);
				iterationStatement();
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(179);
				callStatement();
				}
				break;
			case 7:
				enterOuterAlt(_localctx, 7);
				{
				setState(180);
				jumpStatement();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class AssignmentStatementContext extends ParserRuleContext {
		public AssignmentExpressionContext assignmentExpression() {
			return getRuleContext(AssignmentExpressionContext.class,0);
		}
		public AssignmentStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_assignmentStatement; }
	}

	public final AssignmentStatementContext assignmentStatement() throws RecognitionException {
		AssignmentStatementContext _localctx = new AssignmentStatementContext(_ctx, getState());
		enterRule(_localctx, 28, RULE_assignmentStatement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(183);
			assignmentExpression();
			setState(184);
			match(Semi);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class CallStatementContext extends ParserRuleContext {
		public TerminalNode Identifier() { return getToken(MParser.Identifier, 0); }
		public ParameterCallListContext parameterCallList() {
			return getRuleContext(ParameterCallListContext.class,0);
		}
		public CallStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_callStatement; }
	}

	public final CallStatementContext callStatement() throws RecognitionException {
		CallStatementContext _localctx = new CallStatementContext(_ctx, getState());
		enterRule(_localctx, 30, RULE_callStatement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(186);
			match(Identifier);
			setState(187);
			match(LeftParen);
			setState(188);
			parameterCallList(0);
			setState(189);
			match(RightParen);
			setState(190);
			match(Semi);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class LabeledStatementContext extends ParserRuleContext {
		public TerminalNode Identifier() { return getToken(MParser.Identifier, 0); }
		public StatementContext statement() {
			return getRuleContext(StatementContext.class,0);
		}
		public LabeledStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_labeledStatement; }
	}

	public final LabeledStatementContext labeledStatement() throws RecognitionException {
		LabeledStatementContext _localctx = new LabeledStatementContext(_ctx, getState());
		enterRule(_localctx, 32, RULE_labeledStatement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(192);
			match(Identifier);
			setState(193);
			match(Colon);
			setState(194);
			statement();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class CompoundStatementContext extends ParserRuleContext {
		public StatementListContext statementList() {
			return getRuleContext(StatementListContext.class,0);
		}
		public CompoundStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_compoundStatement; }
	}

	public final CompoundStatementContext compoundStatement() throws RecognitionException {
		CompoundStatementContext _localctx = new CompoundStatementContext(_ctx, getState());
		enterRule(_localctx, 34, RULE_compoundStatement);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(196);
			match(LeftBrace);
			setState(198);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << Break) | (1L << Char) | (1L << Continue) | (1L << Do) | (1L << For) | (1L << Goto) | (1L << If) | (1L << Int) | (1L << Return) | (1L << Void) | (1L << While) | (1L << LeftBrace) | (1L << Identifier))) != 0)) {
				{
				setState(197);
				statementList(0);
				}
			}

			setState(200);
			match(RightBrace);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class StatementListContext extends ParserRuleContext {
		public StatementContext statement() {
			return getRuleContext(StatementContext.class,0);
		}
		public StatementListContext statementList() {
			return getRuleContext(StatementListContext.class,0);
		}
		public StatementListContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_statementList; }
	}

	public final StatementListContext statementList() throws RecognitionException {
		return statementList(0);
	}

	private StatementListContext statementList(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		StatementListContext _localctx = new StatementListContext(_ctx, _parentState);
		StatementListContext _prevctx = _localctx;
		int _startState = 36;
		enterRecursionRule(_localctx, 36, RULE_statementList, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(203);
			statement();
			}
			_ctx.stop = _input.LT(-1);
			setState(209);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,14,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new StatementListContext(_parentctx, _parentState);
					pushNewRecursionContext(_localctx, _startState, RULE_statementList);
					setState(205);
					if (!(precpred(_ctx, 1))) throw new FailedPredicateException(this, "precpred(_ctx, 1)");
					setState(206);
					statement();
					}
					} 
				}
				setState(211);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,14,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public static class SelectionStatementContext extends ParserRuleContext {
		public CalcExpressionContext calcExpression() {
			return getRuleContext(CalcExpressionContext.class,0);
		}
		public List<StatementContext> statement() {
			return getRuleContexts(StatementContext.class);
		}
		public StatementContext statement(int i) {
			return getRuleContext(StatementContext.class,i);
		}
		public SelectionStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_selectionStatement; }
	}

	public final SelectionStatementContext selectionStatement() throws RecognitionException {
		SelectionStatementContext _localctx = new SelectionStatementContext(_ctx, getState());
		enterRule(_localctx, 38, RULE_selectionStatement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(212);
			match(If);
			setState(213);
			match(LeftParen);
			setState(214);
			calcExpression(0);
			setState(215);
			match(RightParen);
			setState(216);
			statement();
			setState(219);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,15,_ctx) ) {
			case 1:
				{
				setState(217);
				match(Else);
				setState(218);
				statement();
				}
				break;
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class IterationStatementContext extends ParserRuleContext {
		public TerminalNode While() { return getToken(MParser.While, 0); }
		public CalcExpressionContext calcExpression() {
			return getRuleContext(CalcExpressionContext.class,0);
		}
		public StatementContext statement() {
			return getRuleContext(StatementContext.class,0);
		}
		public TerminalNode Do() { return getToken(MParser.Do, 0); }
		public TerminalNode For() { return getToken(MParser.For, 0); }
		public ForConditionContext forCondition() {
			return getRuleContext(ForConditionContext.class,0);
		}
		public IterationStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_iterationStatement; }
	}

	public final IterationStatementContext iterationStatement() throws RecognitionException {
		IterationStatementContext _localctx = new IterationStatementContext(_ctx, getState());
		enterRule(_localctx, 40, RULE_iterationStatement);
		try {
			setState(241);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case While:
				enterOuterAlt(_localctx, 1);
				{
				setState(221);
				match(While);
				setState(222);
				match(LeftParen);
				setState(223);
				calcExpression(0);
				setState(224);
				match(RightParen);
				setState(225);
				statement();
				}
				break;
			case Do:
				enterOuterAlt(_localctx, 2);
				{
				setState(227);
				match(Do);
				setState(228);
				statement();
				setState(229);
				match(While);
				setState(230);
				match(LeftParen);
				setState(231);
				calcExpression(0);
				setState(232);
				match(RightParen);
				setState(233);
				match(Semi);
				}
				break;
			case For:
				enterOuterAlt(_localctx, 3);
				{
				setState(235);
				match(For);
				setState(236);
				match(LeftParen);
				setState(237);
				forCondition();
				setState(238);
				match(RightParen);
				setState(239);
				statement();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ForConditionContext extends ParserRuleContext {
		public List<AssignmentExpressionContext> assignmentExpression() {
			return getRuleContexts(AssignmentExpressionContext.class);
		}
		public AssignmentExpressionContext assignmentExpression(int i) {
			return getRuleContext(AssignmentExpressionContext.class,i);
		}
		public CalcExpressionContext calcExpression() {
			return getRuleContext(CalcExpressionContext.class,0);
		}
		public ForConditionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_forCondition; }
	}

	public final ForConditionContext forCondition() throws RecognitionException {
		ForConditionContext _localctx = new ForConditionContext(_ctx, getState());
		enterRule(_localctx, 42, RULE_forCondition);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(243);
			assignmentExpression();
			setState(244);
			match(Semi);
			setState(246);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << LeftParen) | (1L << LeftBrace) | (1L << PlusPlus) | (1L << MinusMinus) | (1L << Star) | (1L << And) | (1L << Not) | (1L << Tilde) | (1L << Identifier) | (1L << StringLiteral))) != 0)) {
				{
				setState(245);
				calcExpression(0);
				}
			}

			setState(248);
			match(Semi);
			setState(250);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << Char) | (1L << Int) | (1L << Void) | (1L << Identifier))) != 0)) {
				{
				setState(249);
				assignmentExpression();
				}
			}

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class JumpStatementContext extends ParserRuleContext {
		public TerminalNode Identifier() { return getToken(MParser.Identifier, 0); }
		public CalcExpressionContext calcExpression() {
			return getRuleContext(CalcExpressionContext.class,0);
		}
		public JumpStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_jumpStatement; }
	}

	public final JumpStatementContext jumpStatement() throws RecognitionException {
		JumpStatementContext _localctx = new JumpStatementContext(_ctx, getState());
		enterRule(_localctx, 44, RULE_jumpStatement);
		int _la;
		try {
			setState(264);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case Goto:
				enterOuterAlt(_localctx, 1);
				{
				setState(252);
				match(Goto);
				setState(253);
				match(Identifier);
				setState(254);
				match(Semi);
				}
				break;
			case Continue:
				enterOuterAlt(_localctx, 2);
				{
				setState(255);
				match(Continue);
				setState(256);
				match(Semi);
				}
				break;
			case Break:
				enterOuterAlt(_localctx, 3);
				{
				setState(257);
				match(Break);
				setState(258);
				match(Semi);
				}
				break;
			case Return:
				enterOuterAlt(_localctx, 4);
				{
				setState(259);
				match(Return);
				setState(261);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << LeftParen) | (1L << LeftBrace) | (1L << PlusPlus) | (1L << MinusMinus) | (1L << Star) | (1L << And) | (1L << Not) | (1L << Tilde) | (1L << Identifier) | (1L << StringLiteral))) != 0)) {
					{
					setState(260);
					calcExpression(0);
					}
				}

				setState(263);
				match(Semi);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class CompilationUnitContext extends ParserRuleContext {
		public TranslationUnitContext translationUnit() {
			return getRuleContext(TranslationUnitContext.class,0);
		}
		public TerminalNode EOF() { return getToken(MParser.EOF, 0); }
		public CompilationUnitContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_compilationUnit; }
	}

	public final CompilationUnitContext compilationUnit() throws RecognitionException {
		CompilationUnitContext _localctx = new CompilationUnitContext(_ctx, getState());
		enterRule(_localctx, 46, RULE_compilationUnit);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(266);
			translationUnit(0);
			setState(267);
			match(EOF);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class TranslationUnitContext extends ParserRuleContext {
		public TopLevelDeclarationContext topLevelDeclaration() {
			return getRuleContext(TopLevelDeclarationContext.class,0);
		}
		public TranslationUnitContext translationUnit() {
			return getRuleContext(TranslationUnitContext.class,0);
		}
		public TranslationUnitContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_translationUnit; }
	}

	public final TranslationUnitContext translationUnit() throws RecognitionException {
		return translationUnit(0);
	}

	private TranslationUnitContext translationUnit(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		TranslationUnitContext _localctx = new TranslationUnitContext(_ctx, _parentState);
		TranslationUnitContext _prevctx = _localctx;
		int _startState = 48;
		enterRecursionRule(_localctx, 48, RULE_translationUnit, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(270);
			topLevelDeclaration();
			}
			_ctx.stop = _input.LT(-1);
			setState(276);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,21,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new TranslationUnitContext(_parentctx, _parentState);
					pushNewRecursionContext(_localctx, _startState, RULE_translationUnit);
					setState(272);
					if (!(precpred(_ctx, 1))) throw new FailedPredicateException(this, "precpred(_ctx, 1)");
					setState(273);
					topLevelDeclaration();
					}
					} 
				}
				setState(278);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,21,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public static class TopLevelDeclarationContext extends ParserRuleContext {
		public FunctionDefinitionContext functionDefinition() {
			return getRuleContext(FunctionDefinitionContext.class,0);
		}
		public TopLevelAssignmentExpressionContext topLevelAssignmentExpression() {
			return getRuleContext(TopLevelAssignmentExpressionContext.class,0);
		}
		public TopLevelDeclarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_topLevelDeclaration; }
	}

	public final TopLevelDeclarationContext topLevelDeclaration() throws RecognitionException {
		TopLevelDeclarationContext _localctx = new TopLevelDeclarationContext(_ctx, getState());
		enterRule(_localctx, 50, RULE_topLevelDeclaration);
		try {
			setState(281);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,22,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(279);
				functionDefinition();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(280);
				topLevelAssignmentExpression();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class FunctionDefinitionContext extends ParserRuleContext {
		public TypeSpecifierContext typeSpecifier() {
			return getRuleContext(TypeSpecifierContext.class,0);
		}
		public TerminalNode Identifier() { return getToken(MParser.Identifier, 0); }
		public StatementContext statement() {
			return getRuleContext(StatementContext.class,0);
		}
		public PointerContext pointer() {
			return getRuleContext(PointerContext.class,0);
		}
		public ParameterDeclarationListContext parameterDeclarationList() {
			return getRuleContext(ParameterDeclarationListContext.class,0);
		}
		public FunctionDefinitionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_functionDefinition; }
	}

	public final FunctionDefinitionContext functionDefinition() throws RecognitionException {
		FunctionDefinitionContext _localctx = new FunctionDefinitionContext(_ctx, getState());
		enterRule(_localctx, 52, RULE_functionDefinition);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(284);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==Inline) {
				{
				setState(283);
				match(Inline);
				}
			}

			setState(286);
			typeSpecifier(0);
			setState(288);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==Star) {
				{
				setState(287);
				pointer();
				}
			}

			setState(290);
			match(Identifier);
			setState(291);
			match(LeftParen);
			setState(293);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << Char) | (1L << Int) | (1L << Void))) != 0)) {
				{
				setState(292);
				parameterDeclarationList(0);
				}
			}

			setState(295);
			match(RightParen);
			setState(296);
			statement();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public boolean sempred(RuleContext _localctx, int ruleIndex, int predIndex) {
		switch (ruleIndex) {
		case 0:
			return calcExpression_sempred((CalcExpressionContext)_localctx, predIndex);
		case 3:
			return initializerList_sempred((InitializerListContext)_localctx, predIndex);
		case 6:
			return typeSpecifier_sempred((TypeSpecifierContext)_localctx, predIndex);
		case 9:
			return parameterDeclarationList_sempred((ParameterDeclarationListContext)_localctx, predIndex);
		case 11:
			return paramterPassList_sempred((ParamterPassListContext)_localctx, predIndex);
		case 12:
			return parameterCallList_sempred((ParameterCallListContext)_localctx, predIndex);
		case 18:
			return statementList_sempred((StatementListContext)_localctx, predIndex);
		case 24:
			return translationUnit_sempred((TranslationUnitContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean calcExpression_sempred(CalcExpressionContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 5);
		case 1:
			return precpred(_ctx, 3);
		}
		return true;
	}
	private boolean initializerList_sempred(InitializerListContext _localctx, int predIndex) {
		switch (predIndex) {
		case 2:
			return precpred(_ctx, 1);
		}
		return true;
	}
	private boolean typeSpecifier_sempred(TypeSpecifierContext _localctx, int predIndex) {
		switch (predIndex) {
		case 3:
			return precpred(_ctx, 1);
		}
		return true;
	}
	private boolean parameterDeclarationList_sempred(ParameterDeclarationListContext _localctx, int predIndex) {
		switch (predIndex) {
		case 4:
			return precpred(_ctx, 1);
		}
		return true;
	}
	private boolean paramterPassList_sempred(ParamterPassListContext _localctx, int predIndex) {
		switch (predIndex) {
		case 5:
			return precpred(_ctx, 1);
		}
		return true;
	}
	private boolean parameterCallList_sempred(ParameterCallListContext _localctx, int predIndex) {
		switch (predIndex) {
		case 6:
			return precpred(_ctx, 1);
		}
		return true;
	}
	private boolean statementList_sempred(StatementListContext _localctx, int predIndex) {
		switch (predIndex) {
		case 7:
			return precpred(_ctx, 1);
		}
		return true;
	}
	private boolean translationUnit_sempred(TranslationUnitContext _localctx, int predIndex) {
		switch (predIndex) {
		case 8:
			return precpred(_ctx, 1);
		}
		return true;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3<\u012d\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\3\2\3\2\3\2\3\2\3\2\3\2\3\2\3\2\3\2\3\2"+
		"\3\2\3\2\3\2\3\2\3\2\3\2\3\2\3\2\3\2\5\2L\n\2\3\2\3\2\3\2\3\2\3\2\3\2"+
		"\7\2T\n\2\f\2\16\2W\13\2\3\3\3\3\3\4\3\4\3\5\3\5\3\5\3\5\3\5\3\5\7\5c"+
		"\n\5\f\5\16\5f\13\5\3\6\5\6i\n\6\3\6\3\6\3\6\3\6\3\7\3\7\3\7\3\7\3\7\3"+
		"\b\3\b\3\b\3\b\3\b\7\by\n\b\f\b\16\b|\13\b\3\t\3\t\3\t\3\t\3\t\7\t\u0083"+
		"\n\t\f\t\16\t\u0086\13\t\3\n\3\n\3\n\5\n\u008b\n\n\3\13\3\13\3\13\3\13"+
		"\3\13\3\13\7\13\u0093\n\13\f\13\16\13\u0096\13\13\3\f\3\f\3\f\3\r\3\r"+
		"\3\r\3\r\3\r\3\r\7\r\u00a1\n\r\f\r\16\r\u00a4\13\r\3\16\3\16\3\16\3\16"+
		"\3\16\3\16\7\16\u00ac\n\16\f\16\16\16\u00af\13\16\3\17\3\17\3\17\3\17"+
		"\3\17\3\17\3\17\5\17\u00b8\n\17\3\20\3\20\3\20\3\21\3\21\3\21\3\21\3\21"+
		"\3\21\3\22\3\22\3\22\3\22\3\23\3\23\5\23\u00c9\n\23\3\23\3\23\3\24\3\24"+
		"\3\24\3\24\3\24\7\24\u00d2\n\24\f\24\16\24\u00d5\13\24\3\25\3\25\3\25"+
		"\3\25\3\25\3\25\3\25\5\25\u00de\n\25\3\26\3\26\3\26\3\26\3\26\3\26\3\26"+
		"\3\26\3\26\3\26\3\26\3\26\3\26\3\26\3\26\3\26\3\26\3\26\3\26\3\26\5\26"+
		"\u00f4\n\26\3\27\3\27\3\27\5\27\u00f9\n\27\3\27\3\27\5\27\u00fd\n\27\3"+
		"\30\3\30\3\30\3\30\3\30\3\30\3\30\3\30\3\30\5\30\u0108\n\30\3\30\5\30"+
		"\u010b\n\30\3\31\3\31\3\31\3\32\3\32\3\32\3\32\3\32\7\32\u0115\n\32\f"+
		"\32\16\32\u0118\13\32\3\33\3\33\5\33\u011c\n\33\3\34\5\34\u011f\n\34\3"+
		"\34\3\34\5\34\u0123\n\34\3\34\3\34\3\34\5\34\u0128\n\34\3\34\3\34\3\34"+
		"\3\34\2\n\2\b\16\24\30\32&\62\35\2\4\6\b\n\f\16\20\22\24\26\30\32\34\36"+
		" \"$&(*,.\60\62\64\66\2\6\5\2\26\34\36\36 \'\6\2\35\35\37 ##()\5\2\4\4"+
		"\f\f\16\16\3\2\20\21\2\u0137\2K\3\2\2\2\4X\3\2\2\2\6Z\3\2\2\2\b\\\3\2"+
		"\2\2\nh\3\2\2\2\fn\3\2\2\2\16s\3\2\2\2\20\u0084\3\2\2\2\22\u008a\3\2\2"+
		"\2\24\u008c\3\2\2\2\26\u0097\3\2\2\2\30\u009a\3\2\2\2\32\u00a5\3\2\2\2"+
		"\34\u00b7\3\2\2\2\36\u00b9\3\2\2\2 \u00bc\3\2\2\2\"\u00c2\3\2\2\2$\u00c6"+
		"\3\2\2\2&\u00cc\3\2\2\2(\u00d6\3\2\2\2*\u00f3\3\2\2\2,\u00f5\3\2\2\2."+
		"\u010a\3\2\2\2\60\u010c\3\2\2\2\62\u010f\3\2\2\2\64\u011b\3\2\2\2\66\u011e"+
		"\3\2\2\289\b\2\1\29L\7\62\2\2:L\7\65\2\2;<\7\24\2\2<=\5\b\5\2=>\7\25\2"+
		"\2>L\3\2\2\2?@\5\6\4\2@A\5\2\2\6AL\3\2\2\2BC\7\20\2\2CD\5\16\b\2DE\7\21"+
		"\2\2EF\5\2\2\4FL\3\2\2\2GH\7\20\2\2HI\5\2\2\2IJ\7\21\2\2JL\3\2\2\2K8\3"+
		"\2\2\2K:\3\2\2\2K;\3\2\2\2K?\3\2\2\2KB\3\2\2\2KG\3\2\2\2LU\3\2\2\2MN\f"+
		"\7\2\2NO\5\4\3\2OP\5\2\2\bPT\3\2\2\2QR\f\5\2\2RT\5\6\4\2SM\3\2\2\2SQ\3"+
		"\2\2\2TW\3\2\2\2US\3\2\2\2UV\3\2\2\2V\3\3\2\2\2WU\3\2\2\2XY\t\2\2\2Y\5"+
		"\3\2\2\2Z[\t\3\2\2[\7\3\2\2\2\\]\b\5\1\2]^\7\63\2\2^d\3\2\2\2_`\f\3\2"+
		"\2`a\7-\2\2ac\7\63\2\2b_\3\2\2\2cf\3\2\2\2db\3\2\2\2de\3\2\2\2e\t\3\2"+
		"\2\2fd\3\2\2\2gi\5\16\b\2hg\3\2\2\2hi\3\2\2\2ij\3\2\2\2jk\7\62\2\2kl\7"+
		".\2\2lm\5\2\2\2m\13\3\2\2\2no\5\16\b\2op\7\62\2\2pq\7.\2\2qr\5\2\2\2r"+
		"\r\3\2\2\2st\b\b\1\2tu\t\4\2\2uz\3\2\2\2vw\f\3\2\2wy\5\22\n\2xv\3\2\2"+
		"\2y|\3\2\2\2zx\3\2\2\2z{\3\2\2\2{\17\3\2\2\2|z\3\2\2\2}\u0083\n\5\2\2"+
		"~\177\7\20\2\2\177\u0080\5\20\t\2\u0080\u0081\7\21\2\2\u0081\u0083\3\2"+
		"\2\2\u0082}\3\2\2\2\u0082~\3\2\2\2\u0083\u0086\3\2\2\2\u0084\u0082\3\2"+
		"\2\2\u0084\u0085\3\2\2\2\u0085\21\3\2\2\2\u0086\u0084\3\2\2\2\u0087\u008b"+
		"\7 \2\2\u0088\u0089\7 \2\2\u0089\u008b\5\22\n\2\u008a\u0087\3\2\2\2\u008a"+
		"\u0088\3\2\2\2\u008b\23\3\2\2\2\u008c\u008d\b\13\1\2\u008d\u008e\5\26"+
		"\f\2\u008e\u0094\3\2\2\2\u008f\u0090\f\3\2\2\u0090\u0091\7-\2\2\u0091"+
		"\u0093\5\26\f\2\u0092\u008f\3\2\2\2\u0093\u0096\3\2\2\2\u0094\u0092\3"+
		"\2\2\2\u0094\u0095\3\2\2\2\u0095\25\3\2\2\2\u0096\u0094\3\2\2\2\u0097"+
		"\u0098\5\16\b\2\u0098\u0099\7\62\2\2\u0099\27\3\2\2\2\u009a\u009b\b\r"+
		"\1\2\u009b\u009c\5\2\2\2\u009c\u00a2\3\2\2\2\u009d\u009e\f\3\2\2\u009e"+
		"\u009f\7-\2\2\u009f\u00a1\5\2\2\2\u00a0\u009d\3\2\2\2\u00a1\u00a4\3\2"+
		"\2\2\u00a2\u00a0\3\2\2\2\u00a2\u00a3\3\2\2\2\u00a3\31\3\2\2\2\u00a4\u00a2"+
		"\3\2\2\2\u00a5\u00a6\b\16\1\2\u00a6\u00a7\5\2\2\2\u00a7\u00ad\3\2\2\2"+
		"\u00a8\u00a9\f\3\2\2\u00a9\u00aa\7-\2\2\u00aa\u00ac\5\2\2\2\u00ab\u00a8"+
		"\3\2\2\2\u00ac\u00af\3\2\2\2\u00ad\u00ab\3\2\2\2\u00ad\u00ae\3\2\2\2\u00ae"+
		"\33\3\2\2\2\u00af\u00ad\3\2\2\2\u00b0\u00b8\5\"\22\2\u00b1\u00b8\5$\23"+
		"\2\u00b2\u00b8\5\36\20\2\u00b3\u00b8\5(\25\2\u00b4\u00b8\5*\26\2\u00b5"+
		"\u00b8\5 \21\2\u00b6\u00b8\5.\30\2\u00b7\u00b0\3\2\2\2\u00b7\u00b1\3\2"+
		"\2\2\u00b7\u00b2\3\2\2\2\u00b7\u00b3\3\2\2\2\u00b7\u00b4\3\2\2\2\u00b7"+
		"\u00b5\3\2\2\2\u00b7\u00b6\3\2\2\2\u00b8\35\3\2\2\2\u00b9\u00ba\5\n\6"+
		"\2\u00ba\u00bb\7,\2\2\u00bb\37\3\2\2\2\u00bc\u00bd\7\62\2\2\u00bd\u00be"+
		"\7\20\2\2\u00be\u00bf\5\32\16\2\u00bf\u00c0\7\21\2\2\u00c0\u00c1\7,\2"+
		"\2\u00c1!\3\2\2\2\u00c2\u00c3\7\62\2\2\u00c3\u00c4\7+\2\2\u00c4\u00c5"+
		"\5\34\17\2\u00c5#\3\2\2\2\u00c6\u00c8\7\24\2\2\u00c7\u00c9\5&\24\2\u00c8"+
		"\u00c7\3\2\2\2\u00c8\u00c9\3\2\2\2\u00c9\u00ca\3\2\2\2\u00ca\u00cb\7\25"+
		"\2\2\u00cb%\3\2\2\2\u00cc\u00cd\b\24\1\2\u00cd\u00ce\5\34\17\2\u00ce\u00d3"+
		"\3\2\2\2\u00cf\u00d0\f\3\2\2\u00d0\u00d2\5\34\17\2\u00d1\u00cf\3\2\2\2"+
		"\u00d2\u00d5\3\2\2\2\u00d3\u00d1\3\2\2\2\u00d3\u00d4\3\2\2\2\u00d4\'\3"+
		"\2\2\2\u00d5\u00d3\3\2\2\2\u00d6\u00d7\7\n\2\2\u00d7\u00d8\7\20\2\2\u00d8"+
		"\u00d9\5\2\2\2\u00d9\u00da\7\21\2\2\u00da\u00dd\5\34\17\2\u00db\u00dc"+
		"\7\7\2\2\u00dc\u00de\5\34\17\2\u00dd\u00db\3\2\2\2\u00dd\u00de\3\2\2\2"+
		"\u00de)\3\2\2\2\u00df\u00e0\7\17\2\2\u00e0\u00e1\7\20\2\2\u00e1\u00e2"+
		"\5\2\2\2\u00e2\u00e3\7\21\2\2\u00e3\u00e4\5\34\17\2\u00e4\u00f4\3\2\2"+
		"\2\u00e5\u00e6\7\6\2\2\u00e6\u00e7\5\34\17\2\u00e7\u00e8\7\17\2\2\u00e8"+
		"\u00e9\7\20\2\2\u00e9\u00ea\5\2\2\2\u00ea\u00eb\7\21\2\2\u00eb\u00ec\7"+
		",\2\2\u00ec\u00f4\3\2\2\2\u00ed\u00ee\7\b\2\2\u00ee\u00ef\7\20\2\2\u00ef"+
		"\u00f0\5,\27\2\u00f0\u00f1\7\21\2\2\u00f1\u00f2\5\34\17\2\u00f2\u00f4"+
		"\3\2\2\2\u00f3\u00df\3\2\2\2\u00f3\u00e5\3\2\2\2\u00f3\u00ed\3\2\2\2\u00f4"+
		"+\3\2\2\2\u00f5\u00f6\5\n\6\2\u00f6\u00f8\7,\2\2\u00f7\u00f9\5\2\2\2\u00f8"+
		"\u00f7\3\2\2\2\u00f8\u00f9\3\2\2\2\u00f9\u00fa\3\2\2\2\u00fa\u00fc\7,"+
		"\2\2\u00fb\u00fd\5\n\6\2\u00fc\u00fb\3\2\2\2\u00fc\u00fd\3\2\2\2\u00fd"+
		"-\3\2\2\2\u00fe\u00ff\7\t\2\2\u00ff\u0100\7\62\2\2\u0100\u010b\7,\2\2"+
		"\u0101\u0102\7\5\2\2\u0102\u010b\7,\2\2\u0103\u0104\7\3\2\2\u0104\u010b"+
		"\7,\2\2\u0105\u0107\7\r\2\2\u0106\u0108\5\2\2\2\u0107\u0106\3\2\2\2\u0107"+
		"\u0108\3\2\2\2\u0108\u0109\3\2\2\2\u0109\u010b\7,\2\2\u010a\u00fe\3\2"+
		"\2\2\u010a\u0101\3\2\2\2\u010a\u0103\3\2\2\2\u010a\u0105\3\2\2\2\u010b"+
		"/\3\2\2\2\u010c\u010d\5\62\32\2\u010d\u010e\7\2\2\3\u010e\61\3\2\2\2\u010f"+
		"\u0110\b\32\1\2\u0110\u0111\5\64\33\2\u0111\u0116\3\2\2\2\u0112\u0113"+
		"\f\3\2\2\u0113\u0115\5\64\33\2\u0114\u0112\3\2\2\2\u0115\u0118\3\2\2\2"+
		"\u0116\u0114\3\2\2\2\u0116\u0117\3\2\2\2\u0117\63\3\2\2\2\u0118\u0116"+
		"\3\2\2\2\u0119\u011c\5\66\34\2\u011a\u011c\5\f\7\2\u011b\u0119\3\2\2\2"+
		"\u011b\u011a\3\2\2\2\u011c\65\3\2\2\2\u011d\u011f\7\13\2\2\u011e\u011d"+
		"\3\2\2\2\u011e\u011f\3\2\2\2\u011f\u0120\3\2\2\2\u0120\u0122\5\16\b\2"+
		"\u0121\u0123\5\22\n\2\u0122\u0121\3\2\2\2\u0122\u0123\3\2\2\2\u0123\u0124"+
		"\3\2\2\2\u0124\u0125\7\62\2\2\u0125\u0127\7\20\2\2\u0126\u0128\5\24\13"+
		"\2\u0127\u0126\3\2\2\2\u0127\u0128\3\2\2\2\u0128\u0129\3\2\2\2\u0129\u012a"+
		"\7\21\2\2\u012a\u012b\5\34\17\2\u012b\67\3\2\2\2\34KSUdhz\u0082\u0084"+
		"\u008a\u0094\u00a2\u00ad\u00b7\u00c8\u00d3\u00dd\u00f3\u00f8\u00fc\u0107"+
		"\u010a\u0116\u011b\u011e\u0122\u0127";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}