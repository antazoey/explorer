import { readable } from "svelte/store";
import { tendermintUrl } from "../config";

let timer

export function initialValue() {
    return {
        peerCount: 0,
        peers: new Map(),
    }
}

export function makeNetworkStore(args) {
    let initial = initialValue();
    return readable(initial, makeSubscribe(initial, args));
}

function unsubscribe() {
    return
}

function makeSubscribe(data, _args) {
    return set => {
        if (!timer) {
            timer = setInterval(() => {
                fetchNetworkData(data, set)
            }, 5000)
        }

        fetchNetworkData(data, set);

        return unsubscribe;
    };
}

async function fetchNetworkData(data, set) {
    try {
        const response = await fetch(tendermintUrl('validators'));

        if(response.ok) {
            const res = await response.json();

            data.peerCount = res.result.total

            res.result.validators.forEach(validator => {
                data.peers.set(validator.address, validator)
            })

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

export const store = makeNetworkStore()
