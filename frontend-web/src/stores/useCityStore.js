import { create } from "zustand";
import { slugToName, nameToSlug } from "@/shared/constants/cities";

const savedCity = localStorage.getItem("city");

export const useCityStore = create((set) => ({
  city: savedCity || "Москва",
  setCity: (city) => {
    localStorage.setItem("city", city);
    set({ city });
  },
  setCityFromSlug: (slug) => {
    const city = slugToName(slug);
    localStorage.setItem("city", city);
    set({ city });
  },
  getSlug: () => nameToSlug(useCityStore.getState().city),
}));
