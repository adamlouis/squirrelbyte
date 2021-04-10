const fetch = require('node-fetch')

const sleep = (ms) => new Promise(resolve => setTimeout(resolve, ms))

const AbortController = require('abort-controller');

const fixHeaders = async () => {
    let pt = ''
    while (true) {
        const q = `http://localhost:8888/api/documents?page_token=${pt}`
        const res = await fetch(q)
        const j = await res.json()
        pt = j.next_page_token

        for (let d of j.documents) {
            d.header = {
                'api_url': `https://hacker-news.firebaseio.com/v0/item/${d.body.id}.json`,
                'hn_url': "https://news.ycombinator.com/item?id=" + d.body.id,
            }

            const r = await await fetch(`http://localhost:8888/api/documents`, {
                method: 'POST',
                body: JSON.stringify(d),
            })

            if (r.status !== 200) {
                console.log(r)
                process.exit(1)
            }
        }
        if (!pt) {
            process.exit(0)
        }
    }
}

const main = async () => {
    let itemID = parseInt(process.argv[2])
    let toItemID = parseInt(process.argv[3])

    if (!isFinite(itemID)) {
        process.exit(1)
    }
    if (!toItemID(itemID)) {
        process.exit(1)
    }

    while (itemID > toItemID) {
        try {
            const docID = `hn.item.${itemID}`
            res = await fetch(`http://localhost:8888/api/documents/${docID}`)
            if (res.status === 200) {
                itemID -= 1
                continue
            }

            const controller = new AbortController();
            setTimeout(() => {
                controller.abort();
            }, 10000);

            const url = `https://hacker-news.firebaseio.com/v0/item/${itemID}.json`;
            res = await fetch(url, { signal: controller.signal })
            j = await res.json();

            if (res.status !== 200) {
                console.log(res)
                console.log(res.json())
                process.exit(1)
            }

            res = await fetch(`http://localhost:8888/api/documents`, {
                method: 'POST',
                body: JSON.stringify({
                    id: docID,
                    body: j,
                    header: { url },
                }),
                headers: {
                    'Content-Type': 'application/json'
                }
            })
            j = await res.json();
            console.log(JSON.stringify(j))
        } catch (e) {
            console.warn(e)
        }
        itemID -= 1
        await sleep(3000)
    }
}

// fixHeaders();
main();
