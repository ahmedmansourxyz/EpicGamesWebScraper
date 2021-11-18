package scraper

import (
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
    "strings"
)

type Game struct{
    Title string 
    Price string 
    Status string 
    Start string 
    End string 
}

func Games() []Game {
	response, err := http.Get("https://store-site-backend-static-ipv4.ak.epicgames.com/freeGamesPromotions?locale=pl&country=PL&allowCountries=PL")
	if err != nil {
		panic(err.Error())
	}
	defer response.Body.Close()
	byteValue, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err.Error())
	}
    titles := gjson.Get(string(byteValue), `data.Catalog.searchStore.elements.#.title`)
    prices := gjson.Get(string(byteValue), `data.Catalog.searchStore.elements.#.price.totalPrice.fmtPrice.originalPrice`)

    var indices []int 
    promotions := gjson.Get(string(byteValue), `data.Catalog.searchStore.elements.#.promotions`)
    for index, value := range promotions.Array() {
        if value.String() != "" {
            indices = append(indices, index) 
        }
    }

    starts := gjson.Get(string(byteValue), `data.Catalog.searchStore.elements.#.promotions.promotionalOffers.#.promotionalOffers.#.startDate`)
    var indexes []int 
    var this []int 
    for index, value := range starts.Array() {
        if len(value.Array()) > 0 {
            indexes = append(indexes, indices[index])
            this = append(this, index)
        }
    }
    ends := gjson.Get(string(byteValue), `data.Catalog.searchStore.elements.#.promotions.promotionalOffers.#.promotionalOffers.#.endDate`)
    Games := make([]Game, len(indexes))
    for index := range indexes {
        title := titles.Array()[indexes[this[index]]].String()
        price := prices.Array()[indexes[this[index]]].String()
        start := strings.Split(starts.Array()[this[index]].Array()[0].Array()[0].String(), "T")[0]
        end :=  strings.Split(ends.Array()[this[index]].Array()[0].Array()[0].String(), "T")[0]
        Games[index] = Game{Title:title, Price:price, Start:start, End:end, Status: "Free"}
    }
    return Games 
}
