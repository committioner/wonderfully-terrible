unnamed gopher (almost but didnt make this h1)

an attempt at showing some cool go stuff ive picked up over the years. the channels primitves for ~lockless/async impl are particularly cool, among other facets.

so far, this doesnt do much except show the overhead associated with `JSON over HTTP` activites in go syntax

## helpful commands:
- /bin/bash ~> `curl -LOs "https://interview-data.herokuapp.com/{surveys,survey-questions,survey-responses}"`

## perhaps-interesting output block:

### exhibit a:

``` bash
# fs caching benefit purely on gopher ~instruction availability/no network contention to wait for etc..... can vc data artifacts etc. eg `git add data/*`
e@coip:~/lab/lattice$ go run main.go 
DEBUG: priming...INFO: loading from cached file ./data/surveys
DEBUG: primed (177µs)
surveys in! get em while theyre hot :-)
15 surveys found
e@coip:~/lab/lattice$ rm data/surveys 
e@coip:~/lab/lattice$ go run main.go 
DEBUG: priming...INFO: couldnt load from cached files on fs, os.Open("./data/surveys")
err when trying to load from FS cache: open ./data/surveys: no such file or directory
DEBUG: primed (354.7523ms)
surveys in! get em while theyre hot :-)
15 surveys found
e@coip:~/lab/lattice$ go run main.go 
DEBUG: priming...INFO: loading from cached file ./data/surveys
DEBUG: primed (167.3µs)
surveys in! get em while theyre hot :-)
15 surveys found

```