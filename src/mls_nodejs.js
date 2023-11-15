import fetch from 'node-fetch';

const username = 'USERNAME';
const password = 'PASSWORD';
const body = {
 'source': 'universal',
 'url': 'https://mls.foreclosure.com/listing/search.html?ci=bay%20shore&st=ny&utm_source=internal&utm_medium=link&utm_campaign=MLS_top_links',
};
const response = await fetch('https://realtime.oxylabs.io/v1/queries', {
  method: 'post',
  body: JSON.stringify(body),
  headers: {
      'Content-Type': 'application/json',
      'Authorization': 'Basic ' + Buffer.from(`${username}:${password}`).toString('base64'),
  }
});

console.log(await response.json());
