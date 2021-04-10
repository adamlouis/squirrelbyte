const fetch = require('node-fetch')

const sleep = (ms) => new Promise(resolve => setTimeout(resolve, ms))

const AbortController = require('abort-controller');

const main  = async () => {
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
            res = await fetch(url, {signal: controller.signal})
            j = await res.json();

            if (res.status !== 200) {
                console.log(res)
                console.log(res.json())
                process.exit(1)
            }
            
            res = await fetch(`http://localhost:8888/api/documents`, {
                method: 'POST',
                body: JSON.stringify({
                    id:docID, 
                    body:j, 
                    header:{ url },
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

main();
