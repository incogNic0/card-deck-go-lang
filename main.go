package main

func main() {
	myDeck := deckFromFile("my-deck")
	myDeck.shuffle()
	myDeck.print()
}