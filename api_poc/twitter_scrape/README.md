# Twitter Scrape

### Prerequisites
1. python version >= 3.8 and < 3.10
2. `install pip3`

### Set up
1. install flask `pip install Flask`
2. install pandas `pip install pandas`
3. install snsscrape `pip install snsscrape`

### Walkthrough
**Part I :** TwitterScrape function

	1.1) Using TwitterSearchScraper to scrape data and append tweets to list.

```python
def twitterScrape(req):

    attributes_container = []
    input = 'from'+req

    for i,tweet in enumerate(sntwitter.TwitterSearchScraper(input).get_items()):
		# limit to 100 tweets
        if i>100:
            break

        attributes_container.append([tweet.date, tweet.likeCount, tweet.sourceLabel, tweet.content])
```

	1.2) Creating dataframe from the tweets list above.

```python
    tweets_df = pd.DataFrame(attributes_container, columns=["Date Created", "Number of Likes", "Source of Tweet", "Tweets"])

    df_json = tweets_df.to_json(orient = 'records')
    # df_json = tweets_df.to_json(orient = 'columns')

    return df_json
```

**Part 2:** Create search API using flask

	2.1) routing 'GET' method with query params 'name'
```python
app = Flask(__name__)

@app.route('/')
def index():
    return json.dumps({'health': 'OK'})

@app.route('/search')
def search():
    req = request.args.get('name')
    res = twitterScrape(req)
    return json.dumps(res)

app.run(debug=True)
```

#### Sample curl

```bash
curl --location --request GET 'http://localhost:5000/search?name='elonmusk''
```


#### Example Response

```python
# orient = 'records'
[
    {
        "Date Created": 1669198001000,
        "Number of Likes": 202664,
        "Source of Tweet": "Twitter for iPhone",
        "Tweets": "More and more over time, as we hew closer to the truth, Twitter will earn the trust of the people"
    },
    {
        "Date Created": 1669197642000,
        "Number of Likes": 82755,
        "Source of Tweet": "Twitter for iPhone",
        "Tweets": "I hope you get that which truly makes you happy"
    },
    {
        "Date Created": 1669195155000,
        "Numer of Likes": 13881,
        "Source of Tweet": "Twitter for iPhone",
        "Tweets": "@MattWallace888 @semafor Inmates are running the asylum at WaPo while Jeff parties in his hot tub"
    }
 ]
```

```python
# orient = 'columns'
{
    "Date Created": {
        "0": 1669198001000,
        "1": 1669197642000,
        "2": 1669195155000,
        "3": 1669194698000
    },
    "Number of Likes": {
        "0": 199818,
        "1": 81629,
        "2": 13594,
        "3": 6044
    },
    "Source of Tweet": {
        "0": "Twitter for iPhone",
        "1": "Twitter for iPhone",
        "2": "Twitter for iPhone",
        "3": "Twitter for iPhone"
    },
    "Tweets": {
        "0": "More and more over time, as we hew closer to the truth, Twitter will earn the trust of the people",
        "1": "I hope you get that which truly makes you happy",
        "2": "@MattWallace888 @semafor Inmates are running the asylum at WaPo while Jeff parties in his hot tub",
        "3": "@realchasegeiser \ud83e\udd23\ud83e\udd23 worth it"
    }
}
```

### What's next ?
	1. Enhance function to be able to scrape on more website.
	2. Add catch keyword feature.
	3. Add notification feature when found the keyword what we looking for. Eg. sending notification via Line API.
