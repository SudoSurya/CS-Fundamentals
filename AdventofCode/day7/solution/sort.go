package main


func sortCards(cards *[]Card)  {
    qs(*cards, 0, len(*cards)-1)
}

func qs(cards []Card, low, high int) {
    if low >= high {
        return
    }
    p := partition(cards, low, high)
    qs(cards, low, p-1)
    qs(cards, p+1, high)
}

func partition(cards []Card, low, high int) int {
    pivot := cards[high]
    i := low - 1
    for j := low; j < high; j++ {
        if compareCards(cards[j], pivot) {
            i++
            cards[i], cards[j] = cards[j], cards[i]
        }
    }
    i++
    cards[high] = cards[i]
    cards[i] = pivot
    return i
}
