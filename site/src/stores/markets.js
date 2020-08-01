import { readable } from "svelte/store";
import {apiUrl} from "../config";

export function initialValue() {
    return {
        markets: new Map(),
    }
}

export function makeMarketStore(args) {
    let initial = initialValue();
    return readable(initial, makeSubscribe(initial, args));
}

function unsubscribe() {
    return
}

function makeSubscribe(data, _args) {
    return set => {
        fetchMarketData(data, set);

        return unsubscribe;
    };
}

async function fetchMarketData(data, set) {
    try {
        // 5. Dispatch the request for the users
        const response = await fetch(apiUrl('markets'));

        if(response.ok) {
            const res = await response.json();

            for (const market of res.markets) {
                data.markets.set(market.id,
                    market
                );
            }
            set(data);

        } else {
            const text = response.text();
            throw new Error(text);
        }

    } catch(error) {
        data.error = error;
        set(data);
    }
}

export const store = makeMarketStore()
