# Various Go Unique ID Token count

This small project's purpose is to provide an overview of various unique ID formats, lengths and OpenAI GPT token length.

Valiate with https://platform.openai.com/tokenizer

```go
$ make get
$ make
RL8Fhr2zp (len=9) (t=6)
kpN3hrlzp (len=9) (t=6)
5D7Fhr2zM (len=9) (t=7)
T-NFhr2-M (len=9) (t=6)
dnN3O9lzp (len=9) (t=7)
-- shortid (avg token: 7)

1769274228579766272 (len=19) (t=7)
1769274228919504896 (len=19) (t=7)
1769274229242466304 (len=19) (t=7)
1769274229598982144 (len=19) (t=7)
1769274229934526464 (len=19) (t=7)
-- snowflake (avg token: 7)

cnra99jv48hovi87enbg (len=20) (t=11)
cnra99jv48hovi87enc0 (len=20) (t=11)
cnra99jv48hovi87encg (len=20) (t=11)
cnra99jv48hovi87end0 (len=20) (t=11)
cnra99jv48hovi87endg (len=20) (t=11)
-- xid (avg token: 11)

yYfFSh-5jI2F1K9iRqT-D (len=21) (t=19)
Rf2s0byHiYJjIW6XwNIQA (len=21) (t=16)
Vp5lwfjff1g2S7uLPXArE (len=21) (t=17)
psz-6lsEPcQIThZaLj9fd (len=21) (t=15)
g0km3kzOt9Th1IIFTFAAu (len=21) (t=14)
-- nanoid (avg token: 17)

2do7xRmBTNL3BgJWS9I7IoD6rkH (len=27) (t=20)
2do7xRyagoIfwnK4VPdgd96Vyh8 (len=27) (t=18)
2do7xTjyWnY2p9XUkx4OZ0wBM4y (len=27) (t=25)
2do7xTNbrU0iWJecH5GepsfediZ (len=27) (t=19)
2do7xSCBeyLvpUOtMWq3I8NCWd5 (len=27) (t=20)
-- ksuid (avg token: 21)

229074e1-e9f1-4d7e-b861-08bb41ec6894 (len=36) (t=22)
b0ea697a-115b-4ce9-9383-381657feb258 (len=36) (t=20)
38d7fc11-4b4e-46cd-ba23-00a0cfd829bd (len=36) (t=24)
e43d6c67-6d5c-4f88-a64a-a8441dba5342 (len=36) (t=24)
caba2a41-e0f3-46c1-9a69-cde8c5ee8380 (len=36) (t=25)
-- uuid (avg token: 23)
```
