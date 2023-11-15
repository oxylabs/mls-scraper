# MLS Scraper

[<u>MLS
Scraper</u>](https://oxylabs.io/products/scraper-api/real-estate/mls) is
an advanced scraping solution for public MLS data extraction in
real-time. Follow this guide to see how to easily scrape MLS using
Oxylabs’ Scraper API.

## How it works

MLS results can be retrieved by sending a request to our API. Simply
provide the URLs you want to scrape, and the API will return the HTML of
any MLS page.

### Python code example

The code below illustrates how to scrape MLS listing page and get its
HTML file:

```python
import requests
from pprint import pprint

# Structure payload.
payload = {
   'source': 'universal',
   'url': 'https://mls.foreclosure.com/listing/search.html?ci=bay%20shore&st=ny&utm_source=internal&utm_medium=link&utm_campaign=MLS_top_links'
}

# Get a response.
response = requests.request(
    'POST',
    'https://realtime.oxylabs.io/v1/queries',
    auth=('USERNAME', 'PASSWORD'), #Your credentials go here
    json=payload,
)

# Instead of a response with job status and results URL, this will return the
# JSON response with results.
pprint(response.json())
```

Visit our
[<u>documentation</u>](https://developers.oxylabs.io/scraper-apis/web-scraper-api)
for more details.

### Output sample

```json
{
    "results": [
        {
            "content":"<!doctype html>\n<html lang=\"en\">\n<head>
            ...
            </script></body>\n</html>\n",
            "created_at": "2023-11-14 15:41:18",
            "updated_at": "2023-11-14 15:41:34",
            "page": 1,
            "url": "https://mls.foreclosure.com/listing/search.html?ci=bay%20shore&st=ny&utm_source=internal&utm_medium=link&utm_campaign=MLS_top_links",
            "job_id": "7130218141837191169",
            "status_code": 200,
        }
    ]
}
```

Oxylabs’ MLS Scraper API simplifies your scraping efforts by providing a
way to extract public MLS data without blocks and with speed. You can
extract details about properties, agents, locations, rental price
information, and more. Feel free to reach us via [<u>live
chat</u>](https://oxylabs.io/) or
[<u>email</u>](mailto:support@oxylabs.io) in case you have any
questions.
