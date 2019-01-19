.PHONY: default
default: doc;

doc:
	pandoc -t latex -o doc.pdf -s --template eisvogel --listings  --number-sections *.md
