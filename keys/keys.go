package keys

import (
	. "os"

	"github.com/ChimeraCoder/anaconda"
)

/* TwitterAPIのポインタ変換 public関数 */
func GetTwitterApi() *anaconda.TwitterApi {

	anaconda.SetConsumerKey(Getenv("TWITTER_CONSUMER_KEY"))
	anaconda.SetConsumerSecret(Getenv("TWITTER_CONSUMER_SECRET"))
	api := anaconda.NewTwitterApi(Getenv("TWITTER_ACCESS_TOKEN"), Getenv("TWITTER_ACCESS_TOKEN_SECRET"))

	return api
}
