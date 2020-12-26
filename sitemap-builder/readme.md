### Description

Ex 5 of the gophercises to develop a sitemap builder. The main goal is to 

1. Get all links for a webpage
2. Iterate over all links that are in same domain
3. Repeat step 2 and 3 until all links are visited once
4. Generate an XML 

### Instructions

```
go run pkg/main.go

// will fetch and get all links for https://gophercises.com

// Output
<?xml version="1.0" encoding="UTF-8"?>
<Urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9"><url><loc>https://gophercises.com/demos/cyoa/debate</loc></url><url><loc>https://gophercises.com/demos/cyoa/sean-kelly</loc></url><url><loc>https://gophercises.com/demos/cyoa/mark-bates</loc></url><url><loc>https://gophercises.com/demos/cyoa/denver</loc></url><url><loc>https://gophercises.com/</loc></url><url><loc>https://gophercises.com/demos/cyoa/</loc></url><url><loc>https://gophercises.com/demos/cyoa/new-york</loc></url><url><loc>https://gophercises.com/demos/cyoa/home</loc></url></Urlset>
```