import { create } from "zustand";

export const useMapStore = create((set) => ({
  mapState: null,

  setMapState: (city, coords, zoom) => {
    set({ mapState: { city, coords, zoom } });
  },
}));
