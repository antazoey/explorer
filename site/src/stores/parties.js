import { writable } from 'svelte/store';
import { apiUrl } from '../config/'

const initialValue = new Map()

function Parties() {
	const { subscribe, set, update } = writable(initialValue);

	return {
        subscribe,
        get: async (id) => {
            const value = initialValue.get(id)
            if (value) {
                return value
            } else {
                const res = await fetch(apiUrl('query'), {
                    "headers": {
                        "content-type": "application/json",
                    },
                    "body": "{\"operationName\":null,\"variables\":{},\"query\":\"{\\n  party(id: \\\"" +id+ "\\\") {\\n    id\\n    accounts {\\n      type\\n      asset\\n      market {\\n        id\\n        decimalPlaces\\n      }\\n      balance\\n    }\\n    orders(last: 50) {\\n      id\\n      size\\n      price\\n      side\\n      market {\\n        id\\n        decimalPlaces\\n      }\\n      remaining\\n      trades {\\n        id\\n        size\\n        aggressor\\n        size\\n        seller {\\n          id\\n        }\\n        buyer {\\n          id\\n        }\\n      }\\n    }\\n  }\\n}\\n\"}",
                    "method": "POST",
                    "mode": "cors",
                    "credentials": "omit"
                });

                const { data } = await res.json()
                    //                update(initialValue.set(id, data.party))
                return data.party
            }
       },
		reset: () => set(0)
	};
}

export {Parties}
