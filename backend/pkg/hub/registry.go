package hub

type APIMetadata struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Category    string `json:"category"`
	URL         string `json:"url"`
	Description string `json:"description"`
}

var APIRegistry = []APIMetadata{
	// Politics
	{ID: "rss-cna", Name: "CNA", Category: "Politics", URL: "https://www.channelnewsasia.com/api/v1/rss-outbound-feed?_format=xml", Description: "Politics in Asia"},
	{ID: "rss-lemonde", Name: "Le Monde", Category: "Politics", URL: "https://www.lemonde.fr/rss/une.xml", Description: "French News"},

	// Tech
	{ID: "rss-techcrunch", Name: "TechCrunch", Category: "Tech", URL: "https://techcrunch.com/feed/", Description: "Technology news"},
	{ID: "rss-theverge", Name: "The Verge", Category: "Tech", URL: "https://www.theverge.com/rss/index.xml", Description: "Tech & culture"},

	// AI
	{ID: "rss-openai", Name: "OpenAI Blog", Category: "AI", URL: "https://openai.com/news/rss.xml", Description: "AI developments"},

	// Finance
	{ID: "finnhub-stocks", Name: "Finnhub", Category: "Finance", URL: "https://finnhub.io/api/v1", Description: "Stock quotes"},
	{ID: "coingecko-crypto", Name: "CoinGecko", Category: "Finance", URL: "https://api.coingecko.com/api/v3", Description: "Crypto data"},

	// Military
	{ID: "opensky-flights", Name: "OpenSky", Category: "Military", URL: "https://opensky-network.org/api", Description: "Flight tracking"},
	{ID: "ais-stream", Name: "AISStream", Category: "Military", URL: "wss://stream.aisstream.io", Description: "Vessel tracking"},
}

func GetRegistry() []APIMetadata {
	return APIRegistry
}
