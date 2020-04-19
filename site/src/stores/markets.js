import get from 'axios'
import { readable } from "svelte/store";

let initialVal = new Map();
let Markets = readable(initialVal, (set) => {
  get('https://geo.n.vega.xyz/markets').then(res => {
    let MarketsMap = new Map()

    res.data.markets.forEach(m => {
        MarketsMap.set(m.id, m)
    })

    set(MarketsMap)
  })
 
  return () => {
      get: () => 'Edd'
  };
});

export {Markets}