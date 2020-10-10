# httpreq
Simple http client which make http request to input url. Does not use any external package. Built purely based on socket.

Table of Content

## Usage
### **--url**: Specifies url to fetch
  - httpreq --url https://red-tree-f56a.someshkoli.workers.dev/links
  ```json
  [
    {
      "name": "Facebook",
      "url": "https://www.facebook.com/"
    },
    {
      "name": "Whatsapp",
      "url": "https://web.whatsapp.com/"
    },
    {
      "name": "Instagram",
      "url": "https://www.instagram.com/"
    }
  ]
  ```
### **--profile**: Specifies no of request to make
  - httpreq --url https://red-tree-f56a.someshkoli.workers.dev/links --profile 5
  ```txt
  Mean of all response time: 303.200000 ms
  Median of all response time: 302414358 ms
  Percentage of successfull response: 100%
  Request with fastest response: 308 
  Request with slowest response: 301 
  Request with biggest size: 220 
  Request with smallest size: 220 
  ```
### **--full**: Specifies wether to print full response or only body
  - httpreq --url https://red-tree-f56a.someshkoli.workers.dev/links --full
  ```txt
  HTTP/1.1 200 OK
  Date: Sat, 10 Oct 2020 01:06:46 GMT
  Content-Type: application/json;charset=UTF-8
  Content-Length: 220
  Connection: close
  Set-Cookie: __cfduid=d2df8cbd9614bd076cb48b51759be5f321602292006; expires=Mon, 09-Nov-20 01:06:46 GMT; path=/; domain=.someshkoli.workers.dev; HttpOnly; SameSite=Lax
  cf-request-id: 05b1a4d4b70000f427601e2200000001
  Report-To: {"endpoints":[{"url":"https:\/\/a.nel.cloudflare.com\/report?lkg-colo=21&lkg-time=1602292006"}],"group":"cf-nel","max_age":604800}
  NEL: {"report_to":"cf-nel","max_age":604800}
  Server: cloudflare
  CF-RAY: 5dfc70cdff49f427-LHR

  [
    {
      "name": "Facebook",
      "url": "https://www.facebook.com/"
    },
    {
      "name": "Whatsapp",
      "url": "https://web.whatsapp.com/"
    },
    {
      "name": "Instagram",
      "url": "https://www.instagram.com/"
    }
  ] 
  ```

## Comparision
### With popular domain
- Google
![Google](https://cdn.discordapp.com/attachments/708998065305944075/764294630921207818/httpreqgoogle.png)

- Cloudflare worker
![Cloudflare worker](https://cdn.discordapp.com/attachments/708998065305944075/764294615406739496/httpreqcloudflare.png)

- Facebook
![Facebook](https://cdn.discordapp.com/attachments/708998065305944075/764295589449957386/httpreqfacebook.png)

### Criteria
- No of requests - Since there can be some sort of caching mechanish this can be usefull to determine the breaking/sweet point for no of request that can be handled
- payload - Response time may also depend on type and amount of payload in the response
- Success Rate - Since successfull request is the main thing here, success rate should is compared with each domain
- Status code - Since faliur can be caused due to many reasons, the main cause should be known so as to find proper analytics
