import { writable } from 'svelte/store';

const initialValue = new Map()

function Parties() {
	const { subscribe, set, update } = writable(initialValue);

	return {
        subscribe,
        get: (id) => {
            const value = initialValue.get(id)

            if (value) {
                return value
            } else {
                const res = get(`https://geo.n.vega.xyz/parties/${id}/accounts`)
                update(initialValue.set(id, res.data))
                return res.data
            }
       },
		reset: () => set(0)
	};
}

export {Parties}
