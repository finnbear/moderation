.PHONY=all

all: dictionary_filtered.txt ../wordlists.go

dictionary.txt:
	wget -O dictionary.txt https://raw.githubusercontent.com/dwyl/english-words/master/words_alpha.txt

dictionary_filtered.txt ../wordlists.go: filter_dictionary.go dictionary.txt profanity.txt dictionary_blacklist.txt dictionary_extra.txt
	go run filter_dictionary.go dictionary.txt profanity.txt dictionary_blacklist.txt dictionary_extra.txt dictionary_filtered.txt ../wordlists.go
