import { writable } from 'svelte/store';
import { apiUrl } from '../config/'

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
                const res = get(apiUrl(`/parties/${id}/accounts`))
                update(initialValue.set(id, res.data))
                return res.data
            }
       },
		reset: () => set(0)
	};
}

export {Parties}
