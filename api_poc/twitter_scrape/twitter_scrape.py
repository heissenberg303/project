import json
from flask import Flask, request
import pandas as pd
import snscrape.modules.twitter as sntwitter


# Using TwitterSearchScraper to scrape data and append tweets to list
def twitterScrape(req):
    attributes_container = []
    input = 'from'+req
    for i,tweet in enumerate(sntwitter.TwitterSearchScraper(input).get_items()):
        if i>100:
            break
        attributes_container.append([tweet.date, tweet.likeCount, tweet.sourceLabel, tweet.content])

    # Creating a dataframe from the tweets list above 
    tweets_df = pd.DataFrame(attributes_container, columns=["Date Created", "Number of Likes", "Source of Tweet", "Tweets"])
    df_json = tweets_df.to_json(orient = 'records')
    return df_json
    # df_json = tweets_df.to_json(orient = 'columns')


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
