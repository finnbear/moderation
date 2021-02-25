.PHONY=all

all: ../wordlists.go

dictionary.txt:
	wget -O dictionary.txt https://raw.githubusercontent.com/dwyl/english-words/master/words_alpha.txt

dictionary_common.txt:
	wget -O dictionary_common.txt https://raw.githubusercontent.com/first20hours/google-10000-english/master/google-10000-english.txt

../wordlists.go: generate.go dictionary.txt dictionary_common.txt profanity.csv dictionary_blacklist.txt dictionary_extra.txt
	go run generate.go dictionary.txt dictionary_common.txt profanity.csv dictionary_blacklist.txt dictionary_extra.txt ../wordlists.go ../wordlists.csv
