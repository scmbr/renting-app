import { create } from "zustand";

export const useNotificationStore = create((set) => ({
  notifications: [],
  addNotification: (notification) =>
    set((state) => ({
      notifications: [notification, ...state.notifications],
    })),
  setNotifications: (notifications) => set({ notifications }),
}));