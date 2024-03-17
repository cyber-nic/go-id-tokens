# Various Go Unique ID Token count

This small project's purpose is to provide an overview of various unique ID formats, lengths and OpenAI GPT token length.

Valiate with https://platform.openai.com/tokenizer

```go
$ make get
$ make
-ww8s9l-M (len=9) (t=7)
7k65sr2zp (len=9) (t=6)
ST18s92-M (len=9) (t=5)
T4Q8b9l-p (len=9) (t=8)
1z15srl-p (len=9) (t=6)
-- shortid (avg token: 7)

dWCfi4T8tN (len=10) (t=8)
LYvjUzeF38 (len=10) (t=6)
zPEzKI6K8p (len=10) (t=8)
FAadNgf25n (len=10) (t=6)
GVIGPYVQSm (len=10) (t=6)
-- nanoid (avg token: 7)

1769279216483831808 (len=19) (t=7)
1769279216815181824 (len=19) (t=7)
1769279217138143232 (len=19) (t=7)
1769279217452716032 (len=19) (t=7)
1769279217779871744 (len=19) (t=7)
-- snowflake (avg token: 7)

cnraiirv48hp8dn50e4g (len=20) (t=12)
cnraiirv48hp8dn50e50 (len=20) (t=11)
cnraij3v48hp8dn50e5g (len=20) (t=13)
cnraij3v48hp8dn50e60 (len=20) (t=12)
cnraij3v48hp8dn50e6g (len=20) (t=13)
-- xid (avg token: 13)

2doAMyeZYugycGz8CmlgQjDqKCW (len=27) (t=19)
2doAMs7ayjCdB9aM8tRmhlm1q89 (len=27) (t=20)
2doAMvFxx8RFcS58KgdP3w4BQCI (len=27) (t=20)
2doAMtzKtwjlJEUdArrurj8jHBj (len=27) (t=17)
2doAMvdGnSqWA9Z9yQElAxMw5Vs (len=27) (t=19)
-- ksuid (avg token: 19)

a9938e75-310f-4deb-b89d-1e18b9978084 (len=36) (t=22)
dfdca538-eb32-4c8b-a5db-06bed92e9ef2 (len=36) (t=22)
2b314cdf-fa7f-427f-8c2c-3f5963677f32 (len=36) (t=24)
9ad6df6c-b98e-4b89-b582-08aeadb9d717 (len=36) (t=23)
8a6f88af-99ca-4d03-81b7-955ba210f8eb (len=36) (t=24)
-- uuid (avg token: 23)
```
