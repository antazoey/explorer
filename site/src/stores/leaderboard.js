import { readable } from "svelte/store";

let initialVal = new Map();
let Leaderboard = readable(initialVal, (set) => {
  fetch('https://topgun-service-testnet.ops.vega.xyz/leaderboard').then(res => {
    set(res.data.traders)
  })
 
  return () => {
      get: () => 'Edd'
  };
});

export {Leaderboard}
