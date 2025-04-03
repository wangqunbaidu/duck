import { defineStore } from "pinia";

export const useDisplayStore = defineStore("display", {
  state: () => ({
    config: {
      navbarActive: false,
    },
  }),
  getters: {
    // doubleCounter: state => state.counter * 2,
    // doubleCounterPlusOne(): number {
    //   return this.doubleCounter + 1
    // },
    isEnable(state) {
      return state.config.navbarActive;
    },
  },
  actions: {
    increment() {
      this.config.navbarActive = !this.config.navbarActive;
    },
  },
});
