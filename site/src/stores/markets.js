import { readable } from "svelte/store";
let initialVal = new Map();
let Markets = readable(initialVal, async (set) => {

  return () => {
      get: () => 'Edd'
  };
});

export {Markets}
