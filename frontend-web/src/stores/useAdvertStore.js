import { create } from "zustand";

export const useAdvertStore = create((set) => ({
  filledFormData: null,
  setFilledFormData: (data) => set({ filledFormData: data }),
  clearFilledFormData: () => set({ filledFormData: null }),
}));
