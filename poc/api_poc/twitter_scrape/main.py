
import pandas as pd
import snscrape.modules.twitter as sntwitter

# Created a list to append all tweet attributes(data)

attributes_container = []

# Using TwitterSearchScraper to scrape data and append tweets to list
for i,tweet in enumerate(sntwitter.TwitterSearchScraper('from:elonmusk').get_items()):
    if i>100:
        break
    attributes_container.append([tweet.date, tweet.likeCount, tweet.sourceLabel, tweet.content])
    
# Creating a dataframe from the tweets list above 
tweets_df = pd.DataFrame(attributes_container, columns=["Date Created", "Number of Likes", "Source of Tweet", "Tweets"])
df_json = tweets_df.to_json(orient = 'records')
# df_json = tweets_df.to_json(orient = 'columns')

print(tweets_df)
print(df_json)
print("Hello World")