// Generated from e:\Repos\MCPC\M\M\Grammar\M.g4 by ANTLR 4.7.1
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class MLexer extends Lexer {
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
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	public static final String[] ruleNames = {
		"Break", "Char", "Continue", "Do", "Else", "For", "Goto", "If", "Inline", 
		"Int", "Return", "Void", "While", "LeftParen", "RightParen", "LeftBracket", 
		"RightBracket", "LeftBrace", "RightBrace", "Less", "LessEqual", "Greater", 
		"GreaterEqual", "LeftShift", "RightShift", "Plus", "PlusPlus", "Minus", 
		"MinusMinus", "Star", "Div", "Mod", "And", "Or", "AndAnd", "OrOr", "Xor", 
		"Not", "Tilde", "Question", "Colon", "Semi", "Comma", "Assign", "Equal", 
		"NotEqual", "Dot", "Identifier", "IdentifierNondigit", "Nondigit", "Digit", 
		"UniversalCharacterName", "HexQuad", "Constant", "IntegerConstant", "BinaryConstant", 
		"DecimalConstant", "OctalConstant", "HexadecimalConstant", "HexadecimalPrefix", 
		"NonzeroDigit", "OctalDigit", "HexadecimalDigit", "ExponentPart", "Sign", 
		"DigitSequence", "BinaryExponentPart", "HexadecimalDigitSequence", "CharacterConstant", 
		"CCharSequence", "CChar", "EscapeSequence", "SimpleEscapeSequence", "OctalEscapeSequence", 
		"HexadecimalEscapeSequence", "StringLiteral", "SCharSequence", "SChar", 
		"ComplexDefine", "AsmBlock", "PreprocessorDirective", "Whitespace", "Newline", 
		"BlockComment", "LineComment"
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


	public MLexer(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "M.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public String[] getChannelNames() { return channelNames; }

	@Override
	public String[] getModeNames() { return modeNames; }

	@Override
	public ATN getATN() { return _ATN; }

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2<\u0269\b\1\4\2\t"+
		"\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13"+
		"\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\4*\t*\4+\t+\4"+
		",\t,\4-\t-\4.\t.\4/\t/\4\60\t\60\4\61\t\61\4\62\t\62\4\63\t\63\4\64\t"+
		"\64\4\65\t\65\4\66\t\66\4\67\t\67\48\t8\49\t9\4:\t:\4;\t;\4<\t<\4=\t="+
		"\4>\t>\4?\t?\4@\t@\4A\tA\4B\tB\4C\tC\4D\tD\4E\tE\4F\tF\4G\tG\4H\tH\4I"+
		"\tI\4J\tJ\4K\tK\4L\tL\4M\tM\4N\tN\4O\tO\4P\tP\4Q\tQ\4R\tR\4S\tS\4T\tT"+
		"\4U\tU\4V\tV\3\2\3\2\3\2\3\2\3\2\3\2\3\3\3\3\3\3\3\3\3\3\3\4\3\4\3\4\3"+
		"\4\3\4\3\4\3\4\3\4\3\4\3\5\3\5\3\5\3\6\3\6\3\6\3\6\3\6\3\7\3\7\3\7\3\7"+
		"\3\b\3\b\3\b\3\b\3\b\3\t\3\t\3\t\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\13\3\13"+
		"\3\13\3\13\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\r\3\r\3\r\3\r\3\r\3\16\3\16\3"+
		"\16\3\16\3\16\3\16\3\17\3\17\3\20\3\20\3\21\3\21\3\22\3\22\3\23\3\23\3"+
		"\24\3\24\3\25\3\25\3\26\3\26\3\26\3\27\3\27\3\30\3\30\3\30\3\31\3\31\3"+
		"\31\3\32\3\32\3\32\3\33\3\33\3\34\3\34\3\34\3\35\3\35\3\36\3\36\3\36\3"+
		"\37\3\37\3 \3 \3!\3!\3\"\3\"\3#\3#\3$\3$\3$\3%\3%\3%\3&\3&\3\'\3\'\3("+
		"\3(\3)\3)\3*\3*\3+\3+\3,\3,\3-\3-\3.\3.\3.\3/\3/\3/\3\60\3\60\3\61\3\61"+
		"\3\61\7\61\u0144\n\61\f\61\16\61\u0147\13\61\3\62\3\62\5\62\u014b\n\62"+
		"\3\63\3\63\3\64\3\64\3\65\3\65\3\65\3\65\3\65\3\65\3\65\3\65\3\65\3\65"+
		"\5\65\u015b\n\65\3\66\3\66\3\66\3\66\3\66\3\67\3\67\5\67\u0164\n\67\3"+
		"8\38\38\38\58\u016a\n8\39\39\39\69\u016f\n9\r9\169\u0170\3:\3:\7:\u0175"+
		"\n:\f:\16:\u0178\13:\3;\3;\7;\u017c\n;\f;\16;\u017f\13;\3<\3<\6<\u0183"+
		"\n<\r<\16<\u0184\3=\3=\3=\3>\3>\3?\3?\3@\3@\3A\3A\5A\u0192\nA\3A\3A\3"+
		"A\5A\u0197\nA\3A\5A\u019a\nA\3B\3B\3C\6C\u019f\nC\rC\16C\u01a0\3D\3D\5"+
		"D\u01a5\nD\3D\3D\3D\5D\u01aa\nD\3D\5D\u01ad\nD\3E\6E\u01b0\nE\rE\16E\u01b1"+
		"\3F\3F\3F\3F\3F\3F\3F\3F\3F\3F\3F\3F\3F\3F\3F\3F\3F\3F\3F\3F\3F\3F\5F"+
		"\u01ca\nF\3G\6G\u01cd\nG\rG\16G\u01ce\3H\3H\5H\u01d3\nH\3I\3I\3I\3I\5"+
		"I\u01d9\nI\3J\3J\3J\3K\3K\3K\3K\3K\3K\3K\3K\3K\3K\3K\5K\u01e9\nK\3L\3"+
		"L\3L\3L\6L\u01ef\nL\rL\16L\u01f0\3M\3M\5M\u01f5\nM\3M\3M\3N\6N\u01fa\n"+
		"N\rN\16N\u01fb\3O\3O\3O\3O\3O\3O\3O\5O\u0205\nO\3P\3P\5P\u0209\nP\3P\3"+
		"P\3P\3P\3P\3P\3P\3P\7P\u0213\nP\fP\16P\u0216\13P\3P\3P\3Q\3Q\3Q\3Q\3Q"+
		"\5Q\u021f\nQ\3Q\7Q\u0222\nQ\fQ\16Q\u0225\13Q\3Q\3Q\7Q\u0229\nQ\fQ\16Q"+
		"\u022c\13Q\3Q\3Q\3Q\3Q\3R\5R\u0233\nR\3R\3R\5R\u0237\nR\3R\7R\u023a\n"+
		"R\fR\16R\u023d\13R\3R\3R\3S\6S\u0242\nS\rS\16S\u0243\3S\3S\3T\3T\5T\u024a"+
		"\nT\3T\5T\u024d\nT\3T\3T\3U\3U\3U\3U\7U\u0255\nU\fU\16U\u0258\13U\3U\3"+
		"U\3U\3U\3U\3V\3V\3V\3V\7V\u0263\nV\fV\16V\u0266\13V\3V\3V\4\u023b\u0256"+
		"\2W\3\3\5\4\7\5\t\6\13\7\r\b\17\t\21\n\23\13\25\f\27\r\31\16\33\17\35"+
		"\20\37\21!\22#\23%\24\'\25)\26+\27-\30/\31\61\32\63\33\65\34\67\359\36"+
		";\37= ?!A\"C#E$G%I&K\'M(O)Q*S+U,W-Y.[/]\60_\61a\62c\2e\2g\2i\2k\2m\63"+
		"o\2q\2s\2u\2w\2y\2{\2}\2\177\2\u0081\2\u0083\2\u0085\64\u0087\2\u0089"+
		"\2\u008b\2\u008d\2\u008f\2\u0091\2\u0093\2\u0095\2\u0097\2\u0099\65\u009b"+
		"\2\u009d\2\u009f\66\u00a1\67\u00a38\u00a59\u00a7:\u00a9;\u00ab<\3\2\22"+
		"\5\2C\\aac|\3\2\62;\4\2DDdd\3\2\62\63\4\2ZZzz\3\2\63;\3\2\629\5\2\62;"+
		"CHch\4\2--//\6\2\f\f\17\17))^^\f\2$$))AA^^cdhhppttvvxx\6\2\f\f\17\17$"+
		"$^^\3\2%%\3\2\177\177\5\2\13\f\17\17\"\"\4\2\f\f\17\17\2\u027e\2\3\3\2"+
		"\2\2\2\5\3\2\2\2\2\7\3\2\2\2\2\t\3\2\2\2\2\13\3\2\2\2\2\r\3\2\2\2\2\17"+
		"\3\2\2\2\2\21\3\2\2\2\2\23\3\2\2\2\2\25\3\2\2\2\2\27\3\2\2\2\2\31\3\2"+
		"\2\2\2\33\3\2\2\2\2\35\3\2\2\2\2\37\3\2\2\2\2!\3\2\2\2\2#\3\2\2\2\2%\3"+
		"\2\2\2\2\'\3\2\2\2\2)\3\2\2\2\2+\3\2\2\2\2-\3\2\2\2\2/\3\2\2\2\2\61\3"+
		"\2\2\2\2\63\3\2\2\2\2\65\3\2\2\2\2\67\3\2\2\2\29\3\2\2\2\2;\3\2\2\2\2"+
		"=\3\2\2\2\2?\3\2\2\2\2A\3\2\2\2\2C\3\2\2\2\2E\3\2\2\2\2G\3\2\2\2\2I\3"+
		"\2\2\2\2K\3\2\2\2\2M\3\2\2\2\2O\3\2\2\2\2Q\3\2\2\2\2S\3\2\2\2\2U\3\2\2"+
		"\2\2W\3\2\2\2\2Y\3\2\2\2\2[\3\2\2\2\2]\3\2\2\2\2_\3\2\2\2\2a\3\2\2\2\2"+
		"m\3\2\2\2\2\u0085\3\2\2\2\2\u0099\3\2\2\2\2\u009f\3\2\2\2\2\u00a1\3\2"+
		"\2\2\2\u00a3\3\2\2\2\2\u00a5\3\2\2\2\2\u00a7\3\2\2\2\2\u00a9\3\2\2\2\2"+
		"\u00ab\3\2\2\2\3\u00ad\3\2\2\2\5\u00b3\3\2\2\2\7\u00b8\3\2\2\2\t\u00c1"+
		"\3\2\2\2\13\u00c4\3\2\2\2\r\u00c9\3\2\2\2\17\u00cd\3\2\2\2\21\u00d2\3"+
		"\2\2\2\23\u00d5\3\2\2\2\25\u00dc\3\2\2\2\27\u00e0\3\2\2\2\31\u00e7\3\2"+
		"\2\2\33\u00ec\3\2\2\2\35\u00f2\3\2\2\2\37\u00f4\3\2\2\2!\u00f6\3\2\2\2"+
		"#\u00f8\3\2\2\2%\u00fa\3\2\2\2\'\u00fc\3\2\2\2)\u00fe\3\2\2\2+\u0100\3"+
		"\2\2\2-\u0103\3\2\2\2/\u0105\3\2\2\2\61\u0108\3\2\2\2\63\u010b\3\2\2\2"+
		"\65\u010e\3\2\2\2\67\u0110\3\2\2\29\u0113\3\2\2\2;\u0115\3\2\2\2=\u0118"+
		"\3\2\2\2?\u011a\3\2\2\2A\u011c\3\2\2\2C\u011e\3\2\2\2E\u0120\3\2\2\2G"+
		"\u0122\3\2\2\2I\u0125\3\2\2\2K\u0128\3\2\2\2M\u012a\3\2\2\2O\u012c\3\2"+
		"\2\2Q\u012e\3\2\2\2S\u0130\3\2\2\2U\u0132\3\2\2\2W\u0134\3\2\2\2Y\u0136"+
		"\3\2\2\2[\u0138\3\2\2\2]\u013b\3\2\2\2_\u013e\3\2\2\2a\u0140\3\2\2\2c"+
		"\u014a\3\2\2\2e\u014c\3\2\2\2g\u014e\3\2\2\2i\u015a\3\2\2\2k\u015c\3\2"+
		"\2\2m\u0163\3\2\2\2o\u0169\3\2\2\2q\u016b\3\2\2\2s\u0172\3\2\2\2u\u0179"+
		"\3\2\2\2w\u0180\3\2\2\2y\u0186\3\2\2\2{\u0189\3\2\2\2}\u018b\3\2\2\2\177"+
		"\u018d\3\2\2\2\u0081\u0199\3\2\2\2\u0083\u019b\3\2\2\2\u0085\u019e\3\2"+
		"\2\2\u0087\u01ac\3\2\2\2\u0089\u01af\3\2\2\2\u008b\u01c9\3\2\2\2\u008d"+
		"\u01cc\3\2\2\2\u008f\u01d2\3\2\2\2\u0091\u01d8\3\2\2\2\u0093\u01da\3\2"+
		"\2\2\u0095\u01e8\3\2\2\2\u0097\u01ea\3\2\2\2\u0099\u01f2\3\2\2\2\u009b"+
		"\u01f9\3\2\2\2\u009d\u0204\3\2\2\2\u009f\u0206\3\2\2\2\u00a1\u0219\3\2"+
		"\2\2\u00a3\u0232\3\2\2\2\u00a5\u0241\3\2\2\2\u00a7\u024c\3\2\2\2\u00a9"+
		"\u0250\3\2\2\2\u00ab\u025e\3\2\2\2\u00ad\u00ae\7d\2\2\u00ae\u00af\7t\2"+
		"\2\u00af\u00b0\7g\2\2\u00b0\u00b1\7c\2\2\u00b1\u00b2\7m\2\2\u00b2\4\3"+
		"\2\2\2\u00b3\u00b4\7e\2\2\u00b4\u00b5\7j\2\2\u00b5\u00b6\7c\2\2\u00b6"+
		"\u00b7\7t\2\2\u00b7\6\3\2\2\2\u00b8\u00b9\7e\2\2\u00b9\u00ba\7q\2\2\u00ba"+
		"\u00bb\7p\2\2\u00bb\u00bc\7v\2\2\u00bc\u00bd\7k\2\2\u00bd\u00be\7p\2\2"+
		"\u00be\u00bf\7w\2\2\u00bf\u00c0\7g\2\2\u00c0\b\3\2\2\2\u00c1\u00c2\7f"+
		"\2\2\u00c2\u00c3\7q\2\2\u00c3\n\3\2\2\2\u00c4\u00c5\7g\2\2\u00c5\u00c6"+
		"\7n\2\2\u00c6\u00c7\7u\2\2\u00c7\u00c8\7g\2\2\u00c8\f\3\2\2\2\u00c9\u00ca"+
		"\7h\2\2\u00ca\u00cb\7q\2\2\u00cb\u00cc\7t\2\2\u00cc\16\3\2\2\2\u00cd\u00ce"+
		"\7i\2\2\u00ce\u00cf\7q\2\2\u00cf\u00d0\7v\2\2\u00d0\u00d1\7q\2\2\u00d1"+
		"\20\3\2\2\2\u00d2\u00d3\7k\2\2\u00d3\u00d4\7h\2\2\u00d4\22\3\2\2\2\u00d5"+
		"\u00d6\7k\2\2\u00d6\u00d7\7p\2\2\u00d7\u00d8\7n\2\2\u00d8\u00d9\7k\2\2"+
		"\u00d9\u00da\7p\2\2\u00da\u00db\7g\2\2\u00db\24\3\2\2\2\u00dc\u00dd\7"+
		"k\2\2\u00dd\u00de\7p\2\2\u00de\u00df\7v\2\2\u00df\26\3\2\2\2\u00e0\u00e1"+
		"\7t\2\2\u00e1\u00e2\7g\2\2\u00e2\u00e3\7v\2\2\u00e3\u00e4\7w\2\2\u00e4"+
		"\u00e5\7t\2\2\u00e5\u00e6\7p\2\2\u00e6\30\3\2\2\2\u00e7\u00e8\7x\2\2\u00e8"+
		"\u00e9\7q\2\2\u00e9\u00ea\7k\2\2\u00ea\u00eb\7f\2\2\u00eb\32\3\2\2\2\u00ec"+
		"\u00ed\7y\2\2\u00ed\u00ee\7j\2\2\u00ee\u00ef\7k\2\2\u00ef\u00f0\7n\2\2"+
		"\u00f0\u00f1\7g\2\2\u00f1\34\3\2\2\2\u00f2\u00f3\7*\2\2\u00f3\36\3\2\2"+
		"\2\u00f4\u00f5\7+\2\2\u00f5 \3\2\2\2\u00f6\u00f7\7]\2\2\u00f7\"\3\2\2"+
		"\2\u00f8\u00f9\7_\2\2\u00f9$\3\2\2\2\u00fa\u00fb\7}\2\2\u00fb&\3\2\2\2"+
		"\u00fc\u00fd\7\177\2\2\u00fd(\3\2\2\2\u00fe\u00ff\7>\2\2\u00ff*\3\2\2"+
		"\2\u0100\u0101\7>\2\2\u0101\u0102\7?\2\2\u0102,\3\2\2\2\u0103\u0104\7"+
		"@\2\2\u0104.\3\2\2\2\u0105\u0106\7@\2\2\u0106\u0107\7?\2\2\u0107\60\3"+
		"\2\2\2\u0108\u0109\7>\2\2\u0109\u010a\7>\2\2\u010a\62\3\2\2\2\u010b\u010c"+
		"\7@\2\2\u010c\u010d\7@\2\2\u010d\64\3\2\2\2\u010e\u010f\7-\2\2\u010f\66"+
		"\3\2\2\2\u0110\u0111\7-\2\2\u0111\u0112\7-\2\2\u01128\3\2\2\2\u0113\u0114"+
		"\7/\2\2\u0114:\3\2\2\2\u0115\u0116\7/\2\2\u0116\u0117\7/\2\2\u0117<\3"+
		"\2\2\2\u0118\u0119\7,\2\2\u0119>\3\2\2\2\u011a\u011b\7\61\2\2\u011b@\3"+
		"\2\2\2\u011c\u011d\7\'\2\2\u011dB\3\2\2\2\u011e\u011f\7(\2\2\u011fD\3"+
		"\2\2\2\u0120\u0121\7~\2\2\u0121F\3\2\2\2\u0122\u0123\7(\2\2\u0123\u0124"+
		"\7(\2\2\u0124H\3\2\2\2\u0125\u0126\7~\2\2\u0126\u0127\7~\2\2\u0127J\3"+
		"\2\2\2\u0128\u0129\7`\2\2\u0129L\3\2\2\2\u012a\u012b\7#\2\2\u012bN\3\2"+
		"\2\2\u012c\u012d\7\u0080\2\2\u012dP\3\2\2\2\u012e\u012f\7A\2\2\u012fR"+
		"\3\2\2\2\u0130\u0131\7<\2\2\u0131T\3\2\2\2\u0132\u0133\7=\2\2\u0133V\3"+
		"\2\2\2\u0134\u0135\7.\2\2\u0135X\3\2\2\2\u0136\u0137\7?\2\2\u0137Z\3\2"+
		"\2\2\u0138\u0139\7?\2\2\u0139\u013a\7?\2\2\u013a\\\3\2\2\2\u013b\u013c"+
		"\7#\2\2\u013c\u013d\7?\2\2\u013d^\3\2\2\2\u013e\u013f\7\60\2\2\u013f`"+
		"\3\2\2\2\u0140\u0145\5c\62\2\u0141\u0144\5c\62\2\u0142\u0144\5g\64\2\u0143"+
		"\u0141\3\2\2\2\u0143\u0142\3\2\2\2\u0144\u0147\3\2\2\2\u0145\u0143\3\2"+
		"\2\2\u0145\u0146\3\2\2\2\u0146b\3\2\2\2\u0147\u0145\3\2\2\2\u0148\u014b"+
		"\5e\63\2\u0149\u014b\5i\65\2\u014a\u0148\3\2\2\2\u014a\u0149\3\2\2\2\u014b"+
		"d\3\2\2\2\u014c\u014d\t\2\2\2\u014df\3\2\2\2\u014e\u014f\t\3\2\2\u014f"+
		"h\3\2\2\2\u0150\u0151\7^\2\2\u0151\u0152\7w\2\2\u0152\u0153\3\2\2\2\u0153"+
		"\u015b\5k\66\2\u0154\u0155\7^\2\2\u0155\u0156\7W\2\2\u0156\u0157\3\2\2"+
		"\2\u0157\u0158\5k\66\2\u0158\u0159\5k\66\2\u0159\u015b\3\2\2\2\u015a\u0150"+
		"\3\2\2\2\u015a\u0154\3\2\2\2\u015bj\3\2\2\2\u015c\u015d\5\177@\2\u015d"+
		"\u015e\5\177@\2\u015e\u015f\5\177@\2\u015f\u0160\5\177@\2\u0160l\3\2\2"+
		"\2\u0161\u0164\5o8\2\u0162\u0164\5\u008bF\2\u0163\u0161\3\2\2\2\u0163"+
		"\u0162\3\2\2\2\u0164n\3\2\2\2\u0165\u016a\5s:\2\u0166\u016a\5u;\2\u0167"+
		"\u016a\5w<\2\u0168\u016a\5q9\2\u0169\u0165\3\2\2\2\u0169\u0166\3\2\2\2"+
		"\u0169\u0167\3\2\2\2\u0169\u0168\3\2\2\2\u016ap\3\2\2\2\u016b\u016c\7"+
		"\62\2\2\u016c\u016e\t\4\2\2\u016d\u016f\t\5\2\2\u016e\u016d\3\2\2\2\u016f"+
		"\u0170\3\2\2\2\u0170\u016e\3\2\2\2\u0170\u0171\3\2\2\2\u0171r\3\2\2\2"+
		"\u0172\u0176\5{>\2\u0173\u0175\5g\64\2\u0174\u0173\3\2\2\2\u0175\u0178"+
		"\3\2\2\2\u0176\u0174\3\2\2\2\u0176\u0177\3\2\2\2\u0177t\3\2\2\2\u0178"+
		"\u0176\3\2\2\2\u0179\u017d\7\62\2\2\u017a\u017c\5}?\2\u017b\u017a\3\2"+
		"\2\2\u017c\u017f\3\2\2\2\u017d\u017b\3\2\2\2\u017d\u017e\3\2\2\2\u017e"+
		"v\3\2\2\2\u017f\u017d\3\2\2\2\u0180\u0182\5y=\2\u0181\u0183\5\177@\2\u0182"+
		"\u0181\3\2\2\2\u0183\u0184\3\2\2\2\u0184\u0182\3\2\2\2\u0184\u0185\3\2"+
		"\2\2\u0185x\3\2\2\2\u0186\u0187\7\62\2\2\u0187\u0188\t\6\2\2\u0188z\3"+
		"\2\2\2\u0189\u018a\t\7\2\2\u018a|\3\2\2\2\u018b\u018c\t\b\2\2\u018c~\3"+
		"\2\2\2\u018d\u018e\t\t\2\2\u018e\u0080\3\2\2\2\u018f\u0191\7g\2\2\u0190"+
		"\u0192\5\u0083B\2\u0191\u0190\3\2\2\2\u0191\u0192\3\2\2\2\u0192\u0193"+
		"\3\2\2\2\u0193\u019a\5\u0085C\2\u0194\u0196\7G\2\2\u0195\u0197\5\u0083"+
		"B\2\u0196\u0195\3\2\2\2\u0196\u0197\3\2\2\2\u0197\u0198\3\2\2\2\u0198"+
		"\u019a\5\u0085C\2\u0199\u018f\3\2\2\2\u0199\u0194\3\2\2\2\u019a\u0082"+
		"\3\2\2\2\u019b\u019c\t\n\2\2\u019c\u0084\3\2\2\2\u019d\u019f\5g\64\2\u019e"+
		"\u019d\3\2\2\2\u019f\u01a0\3\2\2\2\u01a0\u019e\3\2\2\2\u01a0\u01a1\3\2"+
		"\2\2\u01a1\u0086\3\2\2\2\u01a2\u01a4\7r\2\2\u01a3\u01a5\5\u0083B\2\u01a4"+
		"\u01a3\3\2\2\2\u01a4\u01a5\3\2\2\2\u01a5\u01a6\3\2\2\2\u01a6\u01ad\5\u0085"+
		"C\2\u01a7\u01a9\7R\2\2\u01a8\u01aa\5\u0083B\2\u01a9\u01a8\3\2\2\2\u01a9"+
		"\u01aa\3\2\2\2\u01aa\u01ab\3\2\2\2\u01ab\u01ad\5\u0085C\2\u01ac\u01a2"+
		"\3\2\2\2\u01ac\u01a7\3\2\2\2\u01ad\u0088\3\2\2\2\u01ae\u01b0\5\177@\2"+
		"\u01af\u01ae\3\2\2\2\u01b0\u01b1\3\2\2\2\u01b1\u01af\3\2\2\2\u01b1\u01b2"+
		"\3\2\2\2\u01b2\u008a\3\2\2\2\u01b3\u01b4\7)\2\2\u01b4\u01b5\5\u008dG\2"+
		"\u01b5\u01b6\7)\2\2\u01b6\u01ca\3\2\2\2\u01b7\u01b8\7N\2\2\u01b8\u01b9"+
		"\7)\2\2\u01b9\u01ba\3\2\2\2\u01ba\u01bb\5\u008dG\2\u01bb\u01bc\7)\2\2"+
		"\u01bc\u01ca\3\2\2\2\u01bd\u01be\7w\2\2\u01be\u01bf\7)\2\2\u01bf\u01c0"+
		"\3\2\2\2\u01c0\u01c1\5\u008dG\2\u01c1\u01c2\7)\2\2\u01c2\u01ca\3\2\2\2"+
		"\u01c3\u01c4\7W\2\2\u01c4\u01c5\7)\2\2\u01c5\u01c6\3\2\2\2\u01c6\u01c7"+
		"\5\u008dG\2\u01c7\u01c8\7)\2\2\u01c8\u01ca\3\2\2\2\u01c9\u01b3\3\2\2\2"+
		"\u01c9\u01b7\3\2\2\2\u01c9\u01bd\3\2\2\2\u01c9\u01c3\3\2\2\2\u01ca\u008c"+
		"\3\2\2\2\u01cb\u01cd\5\u008fH\2\u01cc\u01cb\3\2\2\2\u01cd\u01ce\3\2\2"+
		"\2\u01ce\u01cc\3\2\2\2\u01ce\u01cf\3\2\2\2\u01cf\u008e\3\2\2\2\u01d0\u01d3"+
		"\n\13\2\2\u01d1\u01d3\5\u0091I\2\u01d2\u01d0\3\2\2\2\u01d2\u01d1\3\2\2"+
		"\2\u01d3\u0090\3\2\2\2\u01d4\u01d9\5\u0093J\2\u01d5\u01d9\5\u0095K\2\u01d6"+
		"\u01d9\5\u0097L\2\u01d7\u01d9\5i\65\2\u01d8\u01d4\3\2\2\2\u01d8\u01d5"+
		"\3\2\2\2\u01d8\u01d6\3\2\2\2\u01d8\u01d7\3\2\2\2\u01d9\u0092\3\2\2\2\u01da"+
		"\u01db\7^\2\2\u01db\u01dc\t\f\2\2\u01dc\u0094\3\2\2\2\u01dd\u01de\7^\2"+
		"\2\u01de\u01e9\5}?\2\u01df\u01e0\7^\2\2\u01e0\u01e1\5}?\2\u01e1\u01e2"+
		"\5}?\2\u01e2\u01e9\3\2\2\2\u01e3\u01e4\7^\2\2\u01e4\u01e5\5}?\2\u01e5"+
		"\u01e6\5}?\2\u01e6\u01e7\5}?\2\u01e7\u01e9\3\2\2\2\u01e8\u01dd\3\2\2\2"+
		"\u01e8\u01df\3\2\2\2\u01e8\u01e3\3\2\2\2\u01e9\u0096\3\2\2\2\u01ea\u01eb"+
		"\7^\2\2\u01eb\u01ec\7z\2\2\u01ec\u01ee\3\2\2\2\u01ed\u01ef\5\177@\2\u01ee"+
		"\u01ed\3\2\2\2\u01ef\u01f0\3\2\2\2\u01f0\u01ee\3\2\2\2\u01f0\u01f1\3\2"+
		"\2\2\u01f1\u0098\3\2\2\2\u01f2\u01f4\7$\2\2\u01f3\u01f5\5\u009bN\2\u01f4"+
		"\u01f3\3\2\2\2\u01f4\u01f5\3\2\2\2\u01f5\u01f6\3\2\2\2\u01f6\u01f7\7$"+
		"\2\2\u01f7\u009a\3\2\2\2\u01f8\u01fa\5\u009dO\2\u01f9\u01f8\3\2\2\2\u01fa"+
		"\u01fb\3\2\2\2\u01fb\u01f9\3\2\2\2\u01fb\u01fc\3\2\2\2\u01fc\u009c\3\2"+
		"\2\2\u01fd\u0205\n\r\2\2\u01fe\u0205\5\u0091I\2\u01ff\u0200\7^\2\2\u0200"+
		"\u0205\7\f\2\2\u0201\u0202\7^\2\2\u0202\u0203\7\17\2\2\u0203\u0205\7\f"+
		"\2\2\u0204\u01fd\3\2\2\2\u0204\u01fe\3\2\2\2\u0204\u01ff\3\2\2\2\u0204"+
		"\u0201\3\2\2\2\u0205\u009e\3\2\2\2\u0206\u0208\7%\2\2\u0207\u0209\5\u00a5"+
		"S\2\u0208\u0207\3\2\2\2\u0208\u0209\3\2\2\2\u0209\u020a\3\2\2\2\u020a"+
		"\u020b\7f\2\2\u020b\u020c\7g\2\2\u020c\u020d\7h\2\2\u020d\u020e\7k\2\2"+
		"\u020e\u020f\7p\2\2\u020f\u0210\7g\2\2\u0210\u0214\3\2\2\2\u0211\u0213"+
		"\n\16\2\2\u0212\u0211\3\2\2\2\u0213\u0216\3\2\2\2\u0214\u0212\3\2\2\2"+
		"\u0214\u0215\3\2\2\2\u0215\u0217\3\2\2\2\u0216\u0214\3\2\2\2\u0217\u0218"+
		"\bP\2\2\u0218\u00a0\3\2\2\2\u0219\u021a\7c\2\2\u021a\u021b\7u\2\2\u021b"+
		"\u021c\7o\2\2\u021c\u021e\3\2\2\2\u021d\u021f\5\u00a5S\2\u021e\u021d\3"+
		"\2\2\2\u021e\u021f\3\2\2\2\u021f\u0223\3\2\2\2\u0220\u0222\7}\2\2\u0221"+
		"\u0220\3\2\2\2\u0222\u0225\3\2\2\2\u0223\u0221\3\2\2\2\u0223\u0224\3\2"+
		"\2\2\u0224\u0226\3\2\2\2\u0225\u0223\3\2\2\2\u0226\u022a\7}\2\2\u0227"+
		"\u0229\n\17\2\2\u0228\u0227\3\2\2\2\u0229\u022c\3\2\2\2\u022a\u0228\3"+
		"\2\2\2\u022a\u022b\3\2\2\2\u022b\u022d\3\2\2\2\u022c\u022a\3\2\2\2\u022d"+
		"\u022e\7\177\2\2\u022e\u022f\3\2\2\2\u022f\u0230\bQ\2\2\u0230\u00a2\3"+
		"\2\2\2\u0231\u0233\5\u00a5S\2\u0232\u0231\3\2\2\2\u0232\u0233\3\2\2\2"+
		"\u0233\u0234\3\2\2\2\u0234\u0236\7%\2\2\u0235\u0237\5\u00a5S\2\u0236\u0235"+
		"\3\2\2\2\u0236\u0237\3\2\2\2\u0237\u023b\3\2\2\2\u0238\u023a\13\2\2\2"+
		"\u0239\u0238\3\2\2\2\u023a\u023d\3\2\2\2\u023b\u023c\3\2\2\2\u023b\u0239"+
		"\3\2\2\2\u023c\u023e\3\2\2\2\u023d\u023b\3\2\2\2\u023e\u023f\5\u00a7T"+
		"\2\u023f\u00a4\3\2\2\2\u0240\u0242\t\20\2\2\u0241\u0240\3\2\2\2\u0242"+
		"\u0243\3\2\2\2\u0243\u0241\3\2\2\2\u0243\u0244\3\2\2\2\u0244\u0245\3\2"+
		"\2\2\u0245\u0246\bS\2\2\u0246\u00a6\3\2\2\2\u0247\u0249\7\17\2\2\u0248"+
		"\u024a\7\f\2\2\u0249\u0248\3\2\2\2\u0249\u024a\3\2\2\2\u024a\u024d\3\2"+
		"\2\2\u024b\u024d\7\f\2\2\u024c\u0247\3\2\2\2\u024c\u024b\3\2\2\2\u024d"+
		"\u024e\3\2\2\2\u024e\u024f\bT\2\2\u024f\u00a8\3\2\2\2\u0250\u0251\7\61"+
		"\2\2\u0251\u0252\7,\2\2\u0252\u0256\3\2\2\2\u0253\u0255\13\2\2\2\u0254"+
		"\u0253\3\2\2\2\u0255\u0258\3\2\2\2\u0256\u0257\3\2\2\2\u0256\u0254\3\2"+
		"\2\2\u0257\u0259\3\2\2\2\u0258\u0256\3\2\2\2\u0259\u025a\7,\2\2\u025a"+
		"\u025b\7\61\2\2\u025b\u025c\3\2\2\2\u025c\u025d\bU\2\2\u025d\u00aa\3\2"+
		"\2\2\u025e\u025f\7\61\2\2\u025f\u0260\7\61\2\2\u0260\u0264\3\2\2\2\u0261"+
		"\u0263\n\21\2\2\u0262\u0261\3\2\2\2\u0263\u0266\3\2\2\2\u0264\u0262\3"+
		"\2\2\2\u0264\u0265\3\2\2\2\u0265\u0267\3\2\2\2\u0266\u0264\3\2\2\2\u0267"+
		"\u0268\bV\2\2\u0268\u00ac\3\2\2\2+\2\u0143\u0145\u014a\u015a\u0163\u0169"+
		"\u0170\u0176\u017d\u0184\u0191\u0196\u0199\u01a0\u01a4\u01a9\u01ac\u01b1"+
		"\u01c9\u01ce\u01d2\u01d8\u01e8\u01f0\u01f4\u01fb\u0204\u0208\u0214\u021e"+
		"\u0223\u022a\u0232\u0236\u023b\u0243\u0249\u024c\u0256\u0264\3\b\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}