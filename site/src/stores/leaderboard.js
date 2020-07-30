import { readable } from "svelte/store";
import { topgunUrl } from "../config";

export function initialValue() {
    return {
        leaderboard: new Map(),
    }
}

export function makeLeaderboardStore(args) {
    let initial = initialValue();
    return readable(initial, makeSubscribe(initial, args));
}

function unsubscribe() {
    return
}

function makeSubscribe(data, _args) {
    return set => {
        fetchLeaderboardData(data, set);

        return unsubscribe;
    };
}

async function fetchLeaderboardData(data, set) {
    try {
        // 5. Dispatch the request for the users
        const response = await fetch(topgunUrl());

        if(response.ok) {
            const res = await response.json();

            let i = 1;
            for (const trader of res.traders) {
                data.leaderboard.set(i, {
                    trader
                });
                i++;
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

export const store = makeLeaderboardStore()
