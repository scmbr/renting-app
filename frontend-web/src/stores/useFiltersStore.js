import { create } from "zustand";

export const useFiltersStore = create((set) => ({
  filters: {},
  setFilters: (newFilters) => set({ filters: newFilters }),
 updateFilter: (key, value) =>
  set((state) => {
    const newFilters = { ...state.filters };
    newFilters[key] = value;
    return { filters: newFilters };
  }),
  resetFilters: () => set({ filters: {} }),
}));
