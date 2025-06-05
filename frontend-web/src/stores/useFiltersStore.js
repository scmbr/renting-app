import { create } from "zustand";

export const useFiltersStore = create((set) => ({
  filters: {},
  setFilters: (newFilters) => set({ filters: newFilters }),
  updateFilter: (key, value) =>
    set((state) => ({
      filters: { ...state.filters, [key]: value },
    })),
  resetFilters: () => set({ filters: {} }),
}));
