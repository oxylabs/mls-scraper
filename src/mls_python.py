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
